#!/usr/bin/env bash

set -x

fileArguments=()

for yaml in complete-api-spec/*.yaml; do
    fileArguments+=("-i")
    fileArguments+=($yaml)
done

openapi-generator generate \
    -g go \
    --git-host github.com \
    --git-repo-id pingdata-config-api-go-client \
    --git-user-id pingidentity \
    "${fileArguments[@]}"
