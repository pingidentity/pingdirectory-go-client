# Generating this client with [OpenAPI Generator](https://openapi-generator.tech/)
This client is created with OpenAPI Generator. This document shows the command used to regenerate the client when the OpenAPI spec has been updated.

Prior to generating the client, the openapi-generator tool will need to be installed. Instructions are included [here](https://openapi-generator.tech/docs/installation)

Use the `scripts/generateClient.sh` script to regenerate the client. Be sure to run this command from the root of this repository.

The generator uses the `api/openapi.yaml` file as the basis for the client. The files under `complete-api-specs/` are not used by the generator. They are just available to view as an easier way to look through the spec for an individual endpoint.
