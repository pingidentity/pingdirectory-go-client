#!/usr/bin/env bash

set -x

# First remove any existing files that were generated. If we want to edit files post-generation, this section will need to be removed
rm ./api_*.go
rm ./model_*.go
rm -r docs/

openapi-generator generate \
    -i complete-api-spec/openapi.yaml \
    -g go \
    --git-host github.com \
    --git-repo-id pingdata-config-api-go-client \
    --git-user-id pingidentity \
    --additional-properties=enumClassPrefix=true

rm -r test/
