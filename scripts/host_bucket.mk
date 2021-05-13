### creating an s3 bucket (for hosting the static webpage)
# ref: https://docs.aws.amazon.com/cli/latest/userguide/cli-services-s3-commands.html
HOST_BUCKET_URI:=s3://$(HOST_BUCKET_NAME)
HOST_BUCKET_ARN:=$(call S3ARN,$(HOST_BUCKET_NAME)/*)
HOST_BUCKET_POLICY:=$(shell cat aws_res/host_bucket_policy.json|\
	sed 's|ARN|$(HOST_BUCKET_ARN)|'|tr -d '\t')
HOST_BUCKET_WEBSITE:=http://$(HOST_BUCKET_NAME).s3-website-$(AWS_REGION).amazonaws.com

.PHONY:
build-website:
	@# update env with api endpoint
	-echo "VITE_API_URL=$(shell $(AWS) apigatewayv2 get-apis|jq -r \
		'[."Items"[]|select(.Name=="$(API_NAME)")][0].ApiEndpoint')/dev/" \
		>src/chfrontend/.env

	-$(WEBSITE_BUILD)

.PHONY:
host-bucket-sync:
	-$(AWS) s3 sync $(WEBSITE_DISTDIR) $(HOST_BUCKET_URI)

.PHONY:
host-bucket-create:
	-$(AWS) s3 mb $(HOST_BUCKET_URI)
	-$(AWS) s3api put-bucket-policy \
		--bucket $(HOST_BUCKET_NAME) \
		--policy '$(HOST_BUCKET_POLICY)'
	-$(AWS) s3 website $(HOST_BUCKET_URI) \
		--index-document index.html \
		--error-document index.html

.PHONY:
host-bucket-delete:
	-$(AWS) s3 rb $(HOST_BUCKET_URI) \
		--force