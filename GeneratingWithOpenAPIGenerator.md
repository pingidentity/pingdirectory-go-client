# Generating this client with [OpenAPI Generator](https://openapi-generator.tech/)
This client is created with OpenAPI Generator. This document shows the command used to regenerate the client when the OpenAPI spec has been updated.

The following command should be run from the root of the repository to regenerate the client:
```
openapi-generator generate -i api/openapi.yaml -g go --git-host github.com --git-repo-id pingdata-config-api-go-client --git-user-id pingidentity
```

This generates a `test/` directory that, so far, I haven't been able to get working correctly. So just remove that directory (for now):

```
rm -r test/
```
