### creating an s3 bucket (for hosting the static webpage)
# ref: https://docs.aws.amazon.com/cli/latest/userguide/cli-services-s3-commands.html
HOST_BUCKET_URI:=s3://$(HOST_BUCKET_NAME)
HOST_BUCKET_ARN:=$(call S3ARN,$(HOST_BUCKET_NAME)/*)
HOST_BUCKET_POLICY:=$(shell cat aws_res/host_bucket_policy.json|\
	sed 's|ARN|$(HOST_BUCKET_ARN)|'|tr -d '\t')

.PHONY:
host-bucket-create:
	-$(AWS) s3 mb $(HOST_BUCKET_URI)
	-$(AWS) s3 sync $(WEBSITE_SRCDIR) $(HOST_BUCKET_URI)
	-$(AWS) s3api put-bucket-policy \
		--bucket $(HOST_BUCKET_NAME) \
		--policy '$(HOST_BUCKET_POLICY)'
	-$(AWS) s3 website $(HOST_BUCKET_URI) \
		--index-document index.html

.PHONY:
host-bucket-delete:
	-$(AWS) s3 rb $(HOST_BUCKET_URI) \
		--force