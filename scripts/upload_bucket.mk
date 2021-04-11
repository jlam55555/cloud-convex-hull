### creating an s3 bucket (for uploads)
# ref: see above
# cors ref: https://docs.aws.amazon.com/AmazonS3/latest/userguide/ManageCorsUsing.html
UPLOAD_BUCKET_URI:=s3://$(UPLOAD_BUCKET_NAME)
UPLOAD_BUCKET_ARN:=$(call S3ARN,$(UPLOAD_BUCKET_NAME))
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
