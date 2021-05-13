### deploying lambda
# ref: (see packaging golang for lambda)
PRESIGN_LAMBDA_ARN:=$(call ARN,lambda,$(PRESIGN_LAMBDA_NAME))
CH_LAMBDA_ARN:=$(call ARN,lambda,$(CH_LAMBDA_NAME))

.PHONY:
lambda-create: $(PRESIGN_GO_ZIP_PATH) $(CH_GO_ZIP_PATH)
	@# presign lambda
	-$(AWS) lambda create-function \
		--function-name $(PRESIGN_LAMBDA_NAME) \
		--runtime go1.x \
		--zip-file $(PRESIGN_GO_AWS_ZIP_PATH) \
		--handler $(PRESIGN_GO_BINARY) \
		--role $(PRESIGN_LAMBDA_ROLE_ARN) \
		--description "$(PRESIGN_LAMBDA_DESC)"|jq .

	@# convex hull lambda
	-$(AWS) lambda create-function \
		--function-name $(CH_LAMBDA_NAME) \
		--runtime go1.x \
		--zip-file $(CH_GO_AWS_ZIP_PATH) \
		--handler $(CH_GO_BINARY) \
		--role $(PRESIGN_LAMBDA_ROLE_ARN) \
		--description "$(CH_LAMBDA_DESC)"|jq .

.PHONY:
lambda-update: $(PRESIGN_GO_ZIP_PATH) $(CH_GO_ZIP_PATH)
	-$(AWS) lambda update-function-code \
		--function-name $(PRESIGN_LAMBDA_NAME) \
		--zip-file $(PRESIGN_GO_AWS_ZIP_PATH)|jq .
	-$(AWS) lambda update-function-code \
		--function-name $(CH_LAMBDA_NAME) \
		--zip-file $(CH_GO_AWS_ZIP_PATH)|jq .

.PHONY:
lambda-delete:
	-$(AWS) lambda delete-function \
		--function-name $(PRESIGN_LAMBDA_NAME)
	-$(AWS) lambda delete-function \
		--function-name $(CH_LAMBDA_NAME)
