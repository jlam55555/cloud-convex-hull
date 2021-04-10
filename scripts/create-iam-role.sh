#!/bin/sh

# create AWS IAM role for use with lambda
# ref: https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html

# source macros
. scripts/common.sh

# lambda execution role trust policy; see ref
TRUST_POLICY='{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Principal": {
				"Service": "lambda.amazonaws.com"
			},
			"Action": "sts:AssumeRole"
		}
	]
}'

echo "Creating IAM lambda role..."
$AWS iam create-role \
	--role-name "$AWS_ROLE" \
	--assume-role-policy-document "$TRUST_POLICY"

echo "Done."