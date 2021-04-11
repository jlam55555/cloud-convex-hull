### creating an API gateway (proxy for s3 and lambda dispatcher)
# ref: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop.html

# this is ugly because APIs don't go by their names, like lambdas or buckets,
# but rather by auto-generated ID's
API_LOGGROUP_FORMAT_FILE:=$(shell cat aws_res/api_loggroup_format.json|\
	sed 's|ARN|$(LOGGROUP_API_ARN)|'|tr -d '\t')
LAMBDA_API_PERMISSION_SID=$(API_NAME)_invoke

# this calls api-delete beforehand to make sure there are not multiple copies
# of the API with the same name
.PHONY:
api-create:
	@# get and store output of create api
	$(call ECHO_SAVE,$(AWS) apigatewayv2 create-api \
		--name $(API_NAME) \
		--protocol-type HTTP,API_ID,'.ApiId')
	$(eval API_ARN=$(call ARN,execute-api,$(API_ID)/*/POST/test))

	$(call ECHO_SAVE,$(AWS) apigatewayv2 create-integration \
		--api-id $(API_ID) \
		--integration-type AWS_PROXY \
		--integration-uri $(LAMBDA_ARN) \
		--payload-format-version 2.0,INTEGRATION_ID,".IntegrationId")

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

.PHONY:
api-get-endpoint:
	-$(AWS) apigatewayv2 get-apis|jq -r \
		'[."Items"[]|select(.Name=="$(API_NAME)")][0].ApiEndpoint'