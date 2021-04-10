# Scripts to automate the build/deployment processes

# AWS primitives
AWS_REGION ?=us-east-1
AWS_PROFILE ?=default
AWS ?=aws --region $(AWS_REGION) --profile $(AWS_PROFILE)

# get AWS_ID
AWS_ID ?=$(shell $(AWS) sts get-caller-identity|jq -r '.Account')

# set default AWS_ROLE, AWS_FUNCTION
# "ch" for convex-hull
AWS_PREFIX ?=ch
AWS_FUNCTION ?=$(AWS_PREFIX)_function
AWS_ROLE ?=$(AWS_PREFIX)_role

################################################################################

.PHONY:
all: iam-create sleep lambda-create

.PHONY:
clean: target-clean iam-delete lambda-delete

################################################################################

### creating iam role for lambda execution
# ref: https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html
TRUST_POLICY:=$(shell cat aws_res/trust_policy.txt|tr -d '\t')

.PHONY:
iam-create:
	-$(AWS) iam create-role \
		--role-name $(AWS_ROLE) \
		--assume-role-policy-document '$(TRUST_POLICY)'

.PHONY:
sleep:
	@echo "Pausing to let IAM provision before creating lambda..."
	@sleep 5

.PHONY:
iam-delete:
	-$(AWS) iam delete-role \
		--role $(AWS_ROLE)


################################################################################

### packaging golang for aws
# ref: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html
BUILDDIR ?=target
SOURCES ?=$(shell find src -name *.go)
PACKAGE ?=$(AWS_PREFIX)lambda
BINARY ?=$(AWS_PREFIX)lambda
GOFLAGS ?=GOOS=linux GOARCH=amd64 CGO_ENABLED=0

$(BUILDDIR)/$(BINARY): $(SOURCES)
	$(GOFLAGS) go build -o $@ $(PACKAGE)

$(BUILDDIR)/$(BINARY).zip: $(BUILDDIR)/$(BINARY)
	zip -j $@ $<

.PHONY:
target-clean:
	rm -rf $(BUILDDIR)

################################################################################

### deploying lambda
# ref: (see packaging golang for lambda)
LAMBDA_DESC ?=Lambda for convex hull application

.PHONY:
lambda-create: $(BUILDDIR)/$(BINARY).zip
	-$(AWS) lambda create-function \
		--function-name $(AWS_FUNCTION) \
		--runtime go1.x \
		--zip-file "fileb://$(BUILDDIR)/$(BINARY).zip" \
		--handler $(BINARY) \
		--role 'arn:aws:iam::$(AWS_ID):role/$(AWS_ROLE)' \
		--description "$(LAMBDA_DESC)"

.PHONY:
lambda-delete:
	-$(AWS) lambda delete-function \
		--function-name $(AWS_FUNCTION)