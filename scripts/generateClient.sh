#!/usr/bin/env bash

set -x

openapi-generator generate -i complete-api-spec/openapi.yaml -g go --git-host github.com --git-repo-id pingdata-config-api-go-client --git-user-id pingidentity

rm -r test/
