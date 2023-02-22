#!/usr/bin/env bash

set -x
set -e

# First remove any existing files that were generated. If we want to edit files post-generation, this section will need to be removed
# rm ./api_*.go
# rm ./model_*.go
# rm -r docs/

cd configurationapi

openapi-generator generate \
    -i ../complete-api-spec/openapi.yaml \
    -g go \
    --git-host github.com \
    --git-repo-id pingdirectory-go-client \
    --git-user-id pingidentity \
    --additional-properties=enumClassPrefix=true,packageName=configurationapi

rm -r test/

# Run any code generators
go mod tidy
go generate ./...

# Clean things up
go fmt ./...
