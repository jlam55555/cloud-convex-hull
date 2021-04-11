### Scripts to set up the AWS infrastructure

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
	lambda-iam-create\
	loggroup-create\
	lambda-create\
	api-create

.PHONY:
clean: target-clean\
	api-delete\
	lambda-delete\
	loggroup-delete\
	lambda-iam-delete\
	upload-bucket-policy-delete\
	upload-bucket-delete\
	host-bucket-delete

-include scripts/host_bucket.mk
-include scripts/upload_bucket.mk
-include scripts/go_compile.mk
-include scripts/upload_bucket_policy.mk
-include scripts/lambda_iam.mk
-include scripts/lambda_deploy.mk
-include scripts/cloudwatch_loggroups.mk
-include scripts/apigateway.mk
