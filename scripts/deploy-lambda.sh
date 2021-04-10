#!/bin/sh

# deploy lambda in src/chlambda to aws
# ref: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html

PACKAGE=${PACKAGE:-chlambda}
BUILDDIR=${BUILDDIR:-target}
BINARY=${BINARY:-chlambda}
HANDLER=${HANDLER:-main}

if [ -z "$AWS_FUNCTION" ] || [ -z "$AWS_ID" ] || [ -z "$AWS_EXECUTION_ROLE" ]
then
	echo -e "error: the AWS_FUNCTION, AWS_ID, and AWS_EXECUTION_ROLE"
	echo -e "\tenvironment variables must be set"
	exit
fi

# build flags
export GOOS=linux
export CGO_ENABLED=0

# build
mkdir -p "$BUILDDIR"
go build -o "$BUILDDIR/$BINARY" "$PACKAGE"

# package for aws
zip "$BUILDDIR/$BINARY.zip" "$BUILDDIR/$BINARY"
aws lambda create-function \
	--function-name "$AWS_FUNCTION" \
	--runtime go1.x \
	--zip-file "$BUILDDIR/$BINARY.zip" \
	--handler "$HANDLER" \
	--role "arn:aws:iam::$AWS_ID:role/$AWS_EXECUTION_ROLE"

echo "Done."