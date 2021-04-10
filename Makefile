# Scripts to automate the build/deployment processes.
#
# Most of the variables are set with default values (?=) be overwritten with
# environment variables. E.g., to override the build directory:
# $ BUILDDIR=build make lambda-create

################################################################################

### Configurables

# AWS primitives
AWS_REGION?=us-east-1
AWS_PROFILE?=default
AWS?=aws --region $(AWS_REGION) --profile $(AWS_PROFILE)

# get AWS_ID
AWS_ID?=$(shell $(AWS) sts get-caller-identity|jq -r '.Account')

# "ch" for convex-hull
APP_PREFIX?=ch

# build directory for all intermediate files
BUILDDIR?=target

# creating bucket for hosting website (note: has to be universally unique)
HOST_BUCKET?=$(APP_PREFIX)hostbucket
WEBSITE_SRCDIR?=src/$(APP_PREFIX)frontend

# creating upload bucket (note: has to be universally unique)
UPLOAD_BUCKET?=$(APP_PREFIX)uploadsbucket

# upload bucket role
UPLOAD_BUCKET_POLICY?=$(UPLOAD_BUCKET)policy

# deploying lambda
LAMBDA_DESC?=Lambda for convex hull application
AWS_FUNCTION?=$(APP_PREFIX)function
LAMBDA_EXEC_ROLE?=$(APP_PREFIX)role

# compiling and packaging lambda
GO_SOURCES?=$(shell find src -name *.go)
GO_PACKAGE?=$(APP_PREFIX)lambda
GO_BINARY?=$(APP_PREFIX)lambda
GO_ENVVAR?=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
GO_LDFLAGS?=-ldflags="-X main.awsRegion=$(AWS_REGION)\
	-X main.uploadBucketName=$(UPLOAD_BUCKET)"

# api gateway
API?=$(APP_PREFIX)_api

################################################################################

.PHONY:
all: host-bucket-create\
	upload-bucket-create\
	upload-bucket-policy-create\
	iam-create sleep\
	lambda-create

.PHONY:
clean: target-clean\
	lambda-delete\
	iam-delete\
	upload-bucket-policy-delete\
	upload-bucket-delete\
	host-bucket-delete

################################################################################

### deploying lambda
# ref: (see packaging golang for lambda)
LAMBDA_ARN:=arn:aws:lambda::$(AWS_ID):function:$(AWS_FUNCTION)

.PHONY:
lambda-create: $(BUILDDIR)/$(GO_BINARY).zip
	-$(AWS) lambda create-function \
		--function-name $(AWS_FUNCTION) \
		--runtime go1.x \
		--zip-file "fileb://$(BUILDDIR)/$(GO_BINARY).zip" \
		--handler $(GO_BINARY) \
		--role 'arn:aws:iam::$(AWS_ID):role/$(LAMBDA_EXEC_ROLE)' \
		--description "$(LAMBDA_DESC)"

.PHONY:
lambda-delete:
	-$(AWS) lambda delete-function \
		--function-name $(AWS_FUNCTION)

################################################################################

### creating an s3 bucket (for hosting the static webpage)
# ref: https://docs.aws.amazon.com/cli/latest/userguide/cli-services-s3-commands.html
HOST_BUCKET_URI:=s3://$(HOST_BUCKET)
HOST_BUCKET_ARN:=arn:aws:s3:::$(HOST_BUCKET)/*
HOST_BUCKET_POLICY:=$(shell cat aws_res/host_bucket_policy.json|\
	sed 's|ARN|$(HOST_BUCKET_ARN)|'|tr -d '\t')

.PHONY:
host-bucket-create:
	-$(AWS) s3 mb $(HOST_BUCKET_URI)
	-$(AWS) s3 sync $(WEBSITE_SRCDIR) $(HOST_BUCKET_URI)
	-$(AWS) s3api put-bucket-policy \
		--bucket $(HOST_BUCKET) \
		--policy '$(HOST_BUCKET_POLICY)'
	-$(AWS) s3 website $(HOST_BUCKET_URI) \
		--index-document index.html

.PHONY:
host-bucket-delete:
	-$(AWS) s3 rb $(HOST_BUCKET_URI) \
		--force

################################################################################

### creating an s3 bucket (for uploads)
# ref: see above
UPLOAD_BUCKET_URI:=s3://$(UPLOAD_BUCKET)
UPLOAD_BUCKET_ARN:=arn:aws:s3:::$(UPLOAD_BUCKET)

.PHONY:
upload-bucket-create:
	-$(AWS) s3 mb $(UPLOAD_BUCKET_URI)

.PHONY:
upload-bucket-delete:
	-$(AWS) s3 rb $(UPLOAD_BUCKET_URI) \
		--force

################################################################################

### creating an API gateway (proxy for s3 and lambda dispatcher)
# ref: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop.html

.PHONY:
api-create:
	-$(AWS) apigatewayv2 create-api \
		--name $(API) \
		--protocol-type HTTP

.PHONY:
api-delete:
	-$(AWS) apigatewayv2 delete-api \
		--name $(API)

################################################################################

### packaging golang for aws
# ref: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html

$(BUILDDIR)/$(GO_BINARY): $(GO_SOURCES)
	$(GO_ENVVAR) go build -o $@ $(GO_LDFLAGS) $(GO_PACKAGE)

# for testing the build
.PHONY:
target-build: $(BUILDDIR)/$(GO_BINARY)

$(BUILDDIR)/$(GO_BINARY).zip: $(BUILDDIR)/$(GO_BINARY)
	zip -j $@ $<

.PHONY:
target-clean:
	rm -rf $(BUILDDIR)

################################################################################

### creating iam policy for accessing upload bucket
# ref: https://docs.aws.amazon.com/cli/latest/reference/iam/create-policy.html
UPLOAD_BUCKET_POLICY_ARN:=arn:aws:iam::$(AWS_ID):policy/$(UPLOAD_BUCKET_POLICY)
UPLOAD_BUCKET_POLICY_FILE:=$(shell cat aws_res/upload_bucket_policy.json|\
	sed 's|ARN|$(UPLOAD_BUCKET_ARN)|'|tr -d '\t')

.PHONY:
upload-bucket-policy-create:
	-$(AWS) iam create-policy \
		--policy-name $(UPLOAD_BUCKET_POLICY) \
		--policy-document '$(UPLOAD_BUCKET_POLICY_FILE)'

.PHONY:
upload-bucket-policy-delete:
	-$(AWS) iam delete-policy \
		--policy-arn $(UPLOAD_BUCKET_POLICY_ARN)

################################################################################

### creating iam role for lambda execution
# ref: https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html
LAMBDA_POLICY:=$(shell cat aws_res/lambda_policy.json|\
	sed 's|ARN|$(UPLOAD_BUCKET_ARN)|'|tr -d '\t')

.PHONY:
iam-create:
	-$(AWS) iam create-role \
		--role-name $(LAMBDA_EXEC_ROLE) \
		--assume-role-policy-document '$(LAMBDA_POLICY)'
	-$(AWS) iam attach-role-policy \
		--role-name $(LAMBDA_EXEC_ROLE) \
		--policy-arn $(UPLOAD_BUCKET_POLICY_ARN)

.PHONY:
sleep:
	@echo "Pausing to let IAM provision before creating lambda..."
	@sleep 5

.PHONY:
iam-delete:
	-$(AWS) iam detach-role-policy \
		--role-name $(LAMBDA_EXEC_ROLE) \
		--policy-arn $(UPLOAD_BUCKET_POLICY_ARN)
	-$(AWS) iam delete-role \
		--role $(LAMBDA_EXEC_ROLE)