# \AccessTokenValidatorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccessTokenValidator**](AccessTokenValidatorAPI.md#AddAccessTokenValidator) | **Post** /access-token-validators | Add a new Access Token Validator to the config
[**DeleteAccessTokenValidator**](AccessTokenValidatorAPI.md#DeleteAccessTokenValidator) | **Delete** /access-token-validators/{access-token-validator-name} | Delete a Access Token Validator
[**GetAccessTokenValidator**](AccessTokenValidatorAPI.md#GetAccessTokenValidator) | **Get** /access-token-validators/{access-token-validator-name} | Returns a single Access Token Validator
[**ListAccessTokenValidators**](AccessTokenValidatorAPI.md#ListAccessTokenValidators) | **Get** /access-token-validators | Returns a list of all Access Token Validator objects
[**UpdateAccessTokenValidator**](AccessTokenValidatorAPI.md#UpdateAccessTokenValidator) | **Patch** /access-token-validators/{access-token-validator-name} | Update an existing Access Token Validator by name



## AddAccessTokenValidator

> AddAccessTokenValidator200Response AddAccessTokenValidator(ctx).AddAccessTokenValidatorRequest(addAccessTokenValidatorRequest).Execute()

Add a new Access Token Validator to the config

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
    addAccessTokenValidatorRequest := openapiclient.add_access_token_validator_request{AddJwtAccessTokenValidatorRequest: openapiclient.NewAddJwtAccessTokenValidatorRequest([]openapiclient.EnumjwtAccessTokenValidatorSchemaUrn{openapiclient.Enumjwt-access-token-validatorSchemaUrn("urn:pingidentity:schemas:configuration:2.0:access-token-validator:jwt")}, false, "ValidatorName_example")} // AddAccessTokenValidatorRequest | Create a new Access Token Validator in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTokenValidatorAPI.AddAccessTokenValidator(context.Background()).AddAccessTokenValidatorRequest(addAccessTokenValidatorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenValidatorAPI.AddAccessTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccessTokenValidator`: AddAccessTokenValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenValidatorAPI.AddAccessTokenValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessTokenValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addAccessTokenValidatorRequest** | [**AddAccessTokenValidatorRequest**](AddAccessTokenValidatorRequest.md) | Create a new Access Token Validator in the config | 

### Return type

[**AddAccessTokenValidator200Response**](AddAccessTokenValidator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessTokenValidator

> DeleteAccessTokenValidator(ctx, accessTokenValidatorName).Execute()

Delete a Access Token Validator

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
    accessTokenValidatorName := "accessTokenValidatorName_example" // string | Name of the Access Token Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessTokenValidatorAPI.DeleteAccessTokenValidator(context.Background(), accessTokenValidatorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenValidatorAPI.DeleteAccessTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessTokenValidatorName** | **string** | Name of the Access Token Validator | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessTokenValidatorRequest struct via the builder pattern


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


## GetAccessTokenValidator

> GetAccessTokenValidator200Response GetAccessTokenValidator(ctx, accessTokenValidatorName).Execute()

Returns a single Access Token Validator

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
    accessTokenValidatorName := "accessTokenValidatorName_example" // string | Name of the Access Token Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTokenValidatorAPI.GetAccessTokenValidator(context.Background(), accessTokenValidatorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenValidatorAPI.GetAccessTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessTokenValidator`: GetAccessTokenValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenValidatorAPI.GetAccessTokenValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessTokenValidatorName** | **string** | Name of the Access Token Validator | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessTokenValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAccessTokenValidator200Response**](GetAccessTokenValidator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessTokenValidators

> AccessTokenValidatorListResponse ListAccessTokenValidators(ctx).Filter(filter).Execute()

Returns a list of all Access Token Validator objects

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
    resp, r, err := apiClient.AccessTokenValidatorAPI.ListAccessTokenValidators(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenValidatorAPI.ListAccessTokenValidators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessTokenValidators`: AccessTokenValidatorListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenValidatorAPI.ListAccessTokenValidators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessTokenValidatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**AccessTokenValidatorListResponse**](AccessTokenValidatorListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessTokenValidator

> GetAccessTokenValidator200Response UpdateAccessTokenValidator(ctx, accessTokenValidatorName).UpdateRequest(updateRequest).Execute()

Update an existing Access Token Validator by name

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
    accessTokenValidatorName := "accessTokenValidatorName_example" // string | Name of the Access Token Validator
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Access Token Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTokenValidatorAPI.UpdateAccessTokenValidator(context.Background(), accessTokenValidatorName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenValidatorAPI.UpdateAccessTokenValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessTokenValidator`: GetAccessTokenValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenValidatorAPI.UpdateAccessTokenValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessTokenValidatorName** | **string** | Name of the Access Token Validator | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessTokenValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Access Token Validator | 

### Return type

[**GetAccessTokenValidator200Response**](GetAccessTokenValidator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

