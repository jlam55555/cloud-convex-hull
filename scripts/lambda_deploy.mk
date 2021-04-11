### deploying lambda
# ref: (see packaging golang for lambda)
LAMBDA_ARN:=$(call ARN,lambda,$(LAMBDA_NAME))

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

