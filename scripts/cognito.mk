### creating Cognito user pool

.PHONY:
user-pool-create:
	-$(AWS) cognito-idp create-user-pool \
		--pool-name $(USERPOOL_NAME)|jq

.PHONY:
user-pool-delete:
	@# in case of multiple apis with same name, delete them all
	-$(foreach USERPOOL_ID,\
		$(shell $(AWS) cognito-idp list-user-pools \
			--max-results 60| \
			jq '.UserPools[]| \
				select(.Name=="$(USERPOOL_NAME)")|.Id'),\
		$(AWS) cognito-idp delete-user-pool \
			--user-pool-id $(USERPOOL_ID);)
