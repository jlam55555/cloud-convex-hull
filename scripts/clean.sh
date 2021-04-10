#!/bin/sh

# cleanup binaries and AWS resources

# source macros
. scripts/common.sh

# remove build target
echo "Removing build target directory..."
rm -rf "${BUILDDIR}"

# delete aws lambda
echo "Deleting lambda $AWS_FUNCTION..."
$AWS lambda delete-function \
	--function-name "$AWS_FUNCTION"

# delete aws role
echo "Deleting role $AWS_ROLE..."
$AWS iam delete-role \
	--role "$AWS_ROLE"

echo "Done."