### deploying lambda
# ref: (see packaging golang for lambda)


.PHONY:
lambda-create: $(PRESIGN_GO_ZIP_PATH)
	@# presign lambda
	-$(AWS) lambda create-function \
		--function-name $(PRESIGN_LAMBDA_NAME) \
		--runtime go1.x \
		--zip-file $(PRESIGN_GO_AWS_ZIP_PATH) \
		--handler $(PRESIGN_GO_BINARY) \
		--role $(PRESIGN_LAMBDA_ROLE_ARN) \
		--description "$(PRESIGN_LAMBDA_DESC)"|jq .

.PHONY:
lambda-update: $(PRESIGN_GO_ZIP_PATH)
	@# presign lambda
	-$(AWS) lambda update-function-code \
		--function-name $(PRESIGN_LAMBDA_NAME) \
		--zip-file $(PRESIGN_GO_AWS_ZIP_PATH)|jq .

.PHONY:
lambda-delete:
	@# presign lambda
	-$(AWS) lambda delete-function \
		--function-name $(PRESIGN_LAMBDA_NAME)

