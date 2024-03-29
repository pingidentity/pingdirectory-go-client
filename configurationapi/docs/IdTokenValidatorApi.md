# \IdTokenValidatorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIdTokenValidator**](IdTokenValidatorAPI.md#AddIdTokenValidator) | **Post** /id-token-validators | Add a new ID Token Validator to the config
[**DeleteIdTokenValidator**](IdTokenValidatorAPI.md#DeleteIdTokenValidator) | **Delete** /id-token-validators/{id-token-validator-name} | Delete a ID Token Validator
[**GetIdTokenValidator**](IdTokenValidatorAPI.md#GetIdTokenValidator) | **Get** /id-token-validators/{id-token-validator-name} | Returns a single ID Token Validator
[**ListIdTokenValidators**](IdTokenValidatorAPI.md#ListIdTokenValidators) | **Get** /id-token-validators | Returns a list of all ID Token Validator objects
[**UpdateIdTokenValidator**](IdTokenValidatorAPI.md#UpdateIdTokenValidator) | **Patch** /id-token-validators/{id-token-validator-name} | Update an existing ID Token Validator by name



## AddIdTokenValidator

> AddIdTokenValidator200Response AddIdTokenValidator(ctx).AddIdTokenValidatorRequest(addIdTokenValidatorRequest).Execute()

Add a new ID Token Validator to the config

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addIdTokenValidatorRequest := openapiclient.add_id_token_validator_request{AddOpenidConnectIdTokenValidatorRequest: openapiclient.NewAddOpenidConnectIdTokenValidatorRequest([]openapiclient.EnumopenidConnectIdTokenValidatorSchemaUrn{openapiclient.Enumopenid-connect-id-token-validatorSchemaUrn("urn:pingidentity:schemas:configuration:2.0:id-token-validator:openid-connect")}, []openapiclient.EnumidTokenValidatorAllowedSigningAlgorithmProp{openapiclient.Enumid-token-validator-allowedSigningAlgorithmProp("RS256")}, false, "IdentityMapper_example", "IssuerURL_example", int64(123), "ValidatorName_example")} // AddIdTokenValidatorRequest | Create a new ID Token Validator in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdTokenValidatorAPI.AddIdTokenValidator(context.Background()).AddIdTokenValidatorRequest(addIdTokenValidatorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdTokenValidatorAPI.AddIdTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIdTokenValidator`: AddIdTokenValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `IdTokenValidatorAPI.AddIdTokenValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIdTokenValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addIdTokenValidatorRequest** | [**AddIdTokenValidatorRequest**](AddIdTokenValidatorRequest.md) | Create a new ID Token Validator in the config | 

### Return type

[**AddIdTokenValidator200Response**](AddIdTokenValidator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdTokenValidator

> DeleteIdTokenValidator(ctx, idTokenValidatorName).Execute()

Delete a ID Token Validator

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    idTokenValidatorName := "idTokenValidatorName_example" // string | Name of the ID Token Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdTokenValidatorAPI.DeleteIdTokenValidator(context.Background(), idTokenValidatorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdTokenValidatorAPI.DeleteIdTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idTokenValidatorName** | **string** | Name of the ID Token Validator | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdTokenValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdTokenValidator

> AddIdTokenValidator200Response GetIdTokenValidator(ctx, idTokenValidatorName).Execute()

Returns a single ID Token Validator

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    idTokenValidatorName := "idTokenValidatorName_example" // string | Name of the ID Token Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdTokenValidatorAPI.GetIdTokenValidator(context.Background(), idTokenValidatorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdTokenValidatorAPI.GetIdTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdTokenValidator`: AddIdTokenValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `IdTokenValidatorAPI.GetIdTokenValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idTokenValidatorName** | **string** | Name of the ID Token Validator | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdTokenValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddIdTokenValidator200Response**](AddIdTokenValidator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdTokenValidators

> IdTokenValidatorListResponse ListIdTokenValidators(ctx).Filter(filter).Execute()

Returns a list of all ID Token Validator objects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdTokenValidatorAPI.ListIdTokenValidators(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdTokenValidatorAPI.ListIdTokenValidators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdTokenValidators`: IdTokenValidatorListResponse
    fmt.Fprintf(os.Stdout, "Response from `IdTokenValidatorAPI.ListIdTokenValidators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdTokenValidatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**IdTokenValidatorListResponse**](IdTokenValidatorListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdTokenValidator

> AddIdTokenValidator200Response UpdateIdTokenValidator(ctx, idTokenValidatorName).UpdateRequest(updateRequest).Execute()

Update an existing ID Token Validator by name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    idTokenValidatorName := "idTokenValidatorName_example" // string | Name of the ID Token Validator
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing ID Token Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdTokenValidatorAPI.UpdateIdTokenValidator(context.Background(), idTokenValidatorName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdTokenValidatorAPI.UpdateIdTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdTokenValidator`: AddIdTokenValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `IdTokenValidatorAPI.UpdateIdTokenValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idTokenValidatorName** | **string** | Name of the ID Token Validator | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdTokenValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing ID Token Validator | 

### Return type

[**AddIdTokenValidator200Response**](AddIdTokenValidator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

