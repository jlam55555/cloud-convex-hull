#!/bin/sh

# deploy lambda in src/chlambda to aws
# ref: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html

# source macros
. scripts/common.sh

PACKAGE=${PACKAGE:-chlambda}
BINARY=${BINARY:-chlambda}
HANDLER=${HANDLER:-chlambda}
LAMBDA_DESC=${LAMBDA_DESC:-"Lambda for convex hull application"}

# build flags
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

# build
echo "Creating build directory..."
mkdir -p "$BUILDDIR"

echo "Building lambda binary..."
go build -o "$BUILDDIR/$BINARY" "$PACKAGE"

# package for aws
echo "Zipping binary..."
zip -j "$BUILDDIR/$BINARY.zip" "$BUILDDIR/$BINARY"

echo "Creating function on lambda..."
$AWS lambda create-function \
	--function-name "$AWS_FUNCTION" \
	--runtime go1.x \
	--zip-file "fileb://$BUILDDIR/$BINARY.zip" \
	--handler "$HANDLER" \
	--role "arn:aws:iam::$AWS_ID:role/$AWS_ROLE" \
	--description "$LAMBDA_DESC"

echo "Done."