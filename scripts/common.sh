#!/bin/sh

# provides macros to be sourced from other files (. common.sh)
# any of these macros can be overwritten by environment variables

AWS_REGION=${AWS_REGION:-us-east-1}
AWS_PROFILE=${AWS_PROFILE:-default}
AWS="aws --region $AWS_REGION --profile $AWS_PROFILE"

# get AWS_ID
AWS_ID=${AWS_ID:-$($AWS sts get-caller-identity|jq -r '.Account')}

# set default AWS_ROLE, AWS_FUNCTION
AWS_PREFIX=${AWS_PREFIX:-"ch"}		# for convex-hull
AWS_FUNCTION=${AWS_FUNCTION:-"${AWS_PREFIX}_function"}
AWS_ROLE=${AWS_ROLE:-"${AWS_PREFIX}_role"}

# default build directory
BUILDDIR=${BUILDDIR:-target}
