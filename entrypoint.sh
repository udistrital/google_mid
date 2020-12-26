#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
export GOOGLE_MID_DRIVE_PRIVATE_KEY="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/google_mid/apigoogle/auth --output text --query Parameter.Value)"
fi

exec ./main "$@"
