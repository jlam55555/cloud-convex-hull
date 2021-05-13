### creating the appropriate (CloudWatch) log groups
LOGGROUP_LAMBDA:=/aws/lambda/$(PRESIGN_LAMBDA_NAME)
CH_LOGGROUP_LAMBDA:=/aws/lambda/$(CH_LAMBDA_NAME)
LOGGROUP_API:=/aws/apigatewayv2/$(API_NAME)
LOGGROUP_API_ARN:=$(call ARN,logs,log-group:$(LOGGROUP_API))

.PHONY:
loggroup-create:
	-$(AWS) logs create-log-group \
		--log-group-name $(LOGGROUP_LAMBDA)
	-$(AWS) logs create-log-group \
		--log-group-name $(CH_LOGGROUP_LAMBDA)
	-$(AWS) logs create-log-group \
		--log-group-name $(LOGGROUP_API)

.PHONY:
loggroup-delete:
	-$(AWS) logs delete-log-group \
		--log-group-name $(LOGGROUP_LAMBDA)
	-$(AWS) logs delete-log-group \
		--log-group-name $(CH_LOGGROUP_LAMBDA)
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
