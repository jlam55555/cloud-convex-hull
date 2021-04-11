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
