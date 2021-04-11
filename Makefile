### Scripts to automate the build/deployment processes.

################################################################################

### configurables
# anything in this section can be overridden with environment variables

# AWS configuration
AWS_REGION?=us-east-1
AWS_PROFILE?=default

# app configuration
# "ch" for convex-hull
APP_PREFIX?=ch

# build directory for all intermediate files
BUILDDIR?=target

# creating bucket for hosting website (note: has to be universally unique)
HOST_BUCKET_NAME?=$(APP_PREFIX)hostbucket
WEBSITE_SRCDIR?=src/$(APP_PREFIX)frontend

# creating upload bucket (note: has to be universally unique)
UPLOAD_BUCKET_NAME?=$(APP_PREFIX)uploadsbucket

# deploying lambda
LAMBDA_NAME?=$(APP_PREFIX)_function
LAMBDA_DESC?=Lambda for convex hull application
LAMBDA_ROLE?=$(APP_PREFIX)_role

# compiling and packaging lambda
GO_SOURCES?=$(shell find src -name *.go)
GO_PACKAGE?=$(APP_PREFIX)lambda
GO_BINARY?=$(APP_PREFIX)lambda
GO_ENVVAR?=GOOS=linux GOARCH=amd64 CGO_ENABLED=0
GO_LDFLAGS?=-ldflags="-X main.awsRegion=$(AWS_REGION)\
	-X main.uploadBucketName=$(UPLOAD_BUCKET_NAME)"

# api gateway
API_NAME?=$(APP_PREFIX)_api
API_STAGE?=dev

################################################################################

### non-configurables
# everything past here is predetermined by the configurables; do not modify
AWS:=aws --region $(AWS_REGION) --profile $(AWS_PROFILE)
AWS_ID:=$(shell $(AWS) sts get-caller-identity|jq -r '.Account')

.PHONY:
all: host-bucket-create\
	upload-bucket-create\
	upload-bucket-policy-create\
	iam-create sleep\
	loggroup-create\
	lambda-create\
	api-create

.PHONY:
clean: target-clean\
	api-delete\
	lambda-delete\
	loggroup-delete\
	iam-delete\
	upload-bucket-policy-delete\
	upload-bucket-delete\
	host-bucket-delete

################################################################################

### creating an s3 bucket (for hosting the static webpage)
# ref: https://docs.aws.amazon.com/cli/latest/userguide/cli-services-s3-commands.html
HOST_BUCKET_URI:=s3://$(HOST_BUCKET_NAME)
HOST_BUCKET_ARN:=arn:aws:s3:::$(HOST_BUCKET_NAME)/*
HOST_BUCKET_POLICY:=$(shell cat aws_res/host_bucket_policy.json|\
	sed 's|ARN|$(HOST_BUCKET_ARN)|'|tr -d '\t')

.PHONY:
host-bucket-create:
	-$(AWS) s3 mb $(HOST_BUCKET_URI)
	-$(AWS) s3 sync $(WEBSITE_SRCDIR) $(HOST_BUCKET_URI)
	-$(AWS) s3api put-bucket-policy \
		--bucket $(HOST_BUCKET_NAME) \
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
# cors ref: https://docs.aws.amazon.com/AmazonS3/latest/userguide/ManageCorsUsing.html
UPLOAD_BUCKET_URI:=s3://$(UPLOAD_BUCKET_NAME)
UPLOAD_BUCKET_ARN:=arn:aws:s3:::$(UPLOAD_BUCKET_NAME)
UPLOAD_BUCKET_CORS_POLICY_FILE:=$(shell cat \
	aws_res/upload_bucket_cors_policy.json|tr -d '\t')

.PHONY:
upload-bucket-create:
	-$(AWS) s3 mb $(UPLOAD_BUCKET_URI)
	-$(AWS) s3api put-bucket-cors \
		--bucket $(UPLOAD_BUCKET_NAME) \
		--cors-configuration '$(UPLOAD_BUCKET_CORS_POLICY_FILE)'

.PHONY:
upload-bucket-delete:
	-$(AWS) s3 rb $(UPLOAD_BUCKET_URI) \
		--force

################################################################################

### packaging golang for aws
# ref: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html
GO_BINARY_PATH:=$(BUILDDIR)/$(GO_BINARY)
GO_ZIP_PATH:=$(GO_BINARY_PATH).zip
GO_AWS_ZIP_PATH:=fileb://$(GO_ZIP_PATH)

$(GO_BINARY_PATH): $(GO_SOURCES)
	$(GO_ENVVAR) go build -o $@ $(GO_LDFLAGS) $(GO_PACKAGE)

# for testing the build
.PHONY:
target-build: $(GO_BINARY_PATH)

$(GO_ZIP_PATH): $(GO_BINARY_PATH)
	zip -j $@ $<

.PHONY:
target-clean:
	rm -rf $(BUILDDIR)

.PHONY:
lint:
	golint src/$(GO_PACKAGE)/**
	go fmt $(GO_PACKAGE)

################################################################################

### creating iam policy for accessing upload bucket
# ref: https://docs.aws.amazon.com/cli/latest/reference/iam/create-policy.html
UPLOAD_BUCKET_POLICY:=$(UPLOAD_BUCKET_NAME)policy
UPLOAD_BUCKET_POLICY_ARN:=arn:aws:iam::$(AWS_ID):policy/$(UPLOAD_BUCKET_POLICY)
UPLOAD_BUCKET_POLICY_FILE:=$(shell cat aws_res/upload_bucket_policy.json|\
	sed 's|ARN|$(UPLOAD_BUCKET_ARN)|'|tr -d '\t')

.PHONY:
upload-bucket-policy-create:
	-$(AWS) iam create-policy \
		--policy-name $(UPLOAD_BUCKET_POLICY) \
		--policy-document '$(UPLOAD_BUCKET_POLICY_FILE)'|jq .

.PHONY:
upload-bucket-policy-delete:
	-$(AWS) iam delete-policy \
		--policy-arn $(UPLOAD_BUCKET_POLICY_ARN)

################################################################################

### creating iam role for lambda execution
# ref: https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html
LAMBDA_POLICY:=$(shell cat aws_res/lambda_policy.json|\
	sed 's|ARN|$(UPLOAD_BUCKET_ARN)|'|tr -d '\t')
LAMBDA_CLOUDWATCH_POLICY_ARN:=arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
LAMBDA_ROLE_ARN:=arn:aws:iam::$(AWS_ID):role/$(LAMBDA_ROLE)

.PHONY:
iam-create:
	-$(AWS) iam create-role \
		--role-name $(LAMBDA_ROLE) \
		--assume-role-policy-document '$(LAMBDA_POLICY)'|jq .

	@# allow lambda access upload bucket
	-$(AWS) iam attach-role-policy \
		--role-name $(LAMBDA_ROLE) \
		--policy-arn $(UPLOAD_BUCKET_POLICY_ARN)|jq .

	@# enable logging to cloudwatch
	-$(AWS) iam attach-role-policy \
		--role-name $(LAMBDA_ROLE) \
		--policy-arn $(LAMBDA_CLOUDWATCH_POLICY_ARN)|jq .

.PHONY:
sleep:
	@echo "Sleeping to let IAM provision before creating lambda..."
	@sleep 5

.PHONY:
iam-delete:
	-$(AWS) iam detach-role-policy \
		--role-name $(LAMBDA_ROLE) \
		--policy-arn $(UPLOAD_BUCKET_POLICY_ARN)
	-$(AWS) iam detach-role-policy \
		--role-name $(LAMBDA_ROLE) \
		--policy-arn $(LAMBDA_CLOUDWATCH_POLICY_ARN)
	-$(AWS) iam delete-role \
		--role $(LAMBDA_ROLE)

################################################################################

### deploying lambda
# ref: (see packaging golang for lambda)
LAMBDA_ARN:=arn:aws:lambda:$(AWS_REGION):$(AWS_ID):$(LAMBDA_NAME)

.PHONY:
lambda-create: $(GO_ZIP_PATH)
	-$(AWS) lambda create-function \
		--function-name $(LAMBDA_NAME) \
		--runtime go1.x \
		--zip-file $(GO_AWS_ZIP_PATH) \
		--handler $(GO_BINARY) \
		--role $(LAMBDA_ROLE_ARN) \
		--description "$(LAMBDA_DESC)"|jq .

.PHONY:
lambda-update: $(GO_ZIP_PATH)
	-$(AWS) lambda update-function-code \
		--function-name $(LAMBDA_NAME) \
		--zip-file $(GO_AWS_ZIP_PATH)|jq .

.PHONY:
lambda-delete:
	-$(AWS) lambda delete-function \
		--function-name $(LAMBDA_NAME)

################################################################################

### creating the appropriate (CloudWatch) log groups
LOGGROUP_LAMBDA:=/aws/lambda/$(LAMBDA_NAME)
LOGGROUP_API:=/aws/apigatewayv2/$(API_NAME)
LOGGROUP_API_ARN:=arn:aws:logs:$(AWS_REGION):$(AWS_ID):log-group:$(LOGGROUP_API)

.PHONY:
loggroup-create:
	-$(AWS) logs create-log-group \
		--log-group-name $(LOGGROUP_LAMBDA)
	-$(AWS) logs create-log-group \
		--log-group-name $(LOGGROUP_API)

.PHONY:
loggroup-delete:
	-$(AWS) logs delete-log-group \
		--log-group-name $(LOGGROUP_LAMBDA)
	-$(AWS) logs delete-log-group \
		--log-group-name $(LOGGROUP_API)

# watch logs
.PHONY:
loggroup-lambda-watch:
	-$(AWS) logs tail $(LOGGROUP_LAMBDA) \
		--follow

.PHONY:
loggroup-api-watch:
	-$(AWS) logs tail $(LOGGROUP_API) \
		--follow

################################################################################

### creating an API gateway (proxy for s3 and lambda dispatcher)
# ref: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop.html

# this is ugly because APIs don't go by their names, like lambdas or buckets,
# but rather by auto-generated ID's
CREATE_API_CMD:=$(AWS) apigatewayv2 create-api \
	--name $(API_NAME) \
	--protocol-type HTTP
CREATE_INTEGRATION_CMD=$(AWS) apigatewayv2 create-integration \
	--api-id $(API_ID) \
	--integration-type AWS_PROXY \
	--integration-uri $(LAMBDA_ARN) \
	--payload-format-version 2.0
API_LOGGROUP_FORMAT_FILE:=$(shell cat aws_res/api_loggroup_format.json|\
	sed 's|ARN|$(LOGGROUP_API_ARN)|'|tr -d '\t')
LAMBDA_API_PERMISSION_SID=$(API_NAME)_invoke

# this calls api-delete beforehand to make sure there are not multiple copies
# of the API with the same name
.PHONY:
api-create: api-delete
	@# get and store output of create api
	@# TODO: convert this to a function
	@echo $(CREATE_API_CMD)
	$(eval RET='$(shell $(CREATE_API_CMD))')
	@echo $(RET)|jq .
	$(eval API_ID=$(shell echo $(RET)|jq -r .ApiId))
	$(eval API_ARN=arn:aws:execute-api:$(AWS_REGION):$(AWS_ID):$(API_ID)/*/POST/test)

	@echo $(CREATE_INTEGRATION_CMD)
	$(eval RET='$(shell $(CREATE_INTEGRATION_CMD))')
	@echo $(RET)|jq .
	$(eval INTEGRATION_ID=$(shell echo $(RET)|jq -r .IntegrationId))

	-$(AWS) apigatewayv2 create-route \
		--api-id $(API_ID) \
		--route-key 'POST /test' \
		--target "integrations/$(INTEGRATION_ID)"|jq .

	-$(AWS) apigatewayv2 create-stage \
		--access-log-settings '$(API_LOGGROUP_FORMAT_FILE)' \
		--api-id $(API_ID) \
		--stage-name $(API_STAGE)|jq .

	-$(AWS) apigatewayv2 create-deployment \
		--api-id $(API_ID) \
		--stage-name $(API_STAGE)|jq .

	-$(AWS) lambda add-permission \
		--function-name $(LAMBDA_ARN) \
		--statement-id $(LAMBDA_API_PERMISSION_SID) \
		--action lambda:InvokeFunction \
		--source-arn $(API_ARN) \
		--principal apigateway.amazonaws.com|jq .

.PHONY:
api-delete:
	@# in case of multiple apis with same name, delete them all
	-$(foreach API_ID,\
		$(shell $(AWS) apigatewayv2 get-apis|\
			jq -r '.["Items"][]|select(.Name=="$(API_NAME)")|\
			.ApiId'),\
		$(AWS) apigatewayv2 delete-api \
			--api-id $(API_ID);)

	-$(AWS) lambda remove-permission \
		--function-name $(LAMBDA_ARN) \
		--statement-id $(LAMBDA_API_PERMISSION_SID)