# Generating this client with [OpenAPI Generator](https://openapi-generator.tech/)
This client is created with OpenAPI Generator. This document shows the command used to regenerate the client when the OpenAPI spec has been updated.

Use the `scripts/generateClient.sh` script to regenerate the client. Be sure to run this command from the root of this repository.

The generator uses the `api/openapi.yaml` file as the basis for the client. The files under `complete-api-specs/` are not used by the generator. They are just available to view as an easier way to look through the spec for an individual endpoint.

This generates a `test/` directory that, so far, I haven't been able to get working correctly. So the script just removes that directory (for now).