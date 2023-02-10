# \TokenClaimValidationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTokenClaimValidation**](TokenClaimValidationApi.md#AddTokenClaimValidation) | **Post** /token-claim-validations | Add a new Token Claim Validation to the config
[**DeleteTokenClaimValidation**](TokenClaimValidationApi.md#DeleteTokenClaimValidation) | **Delete** /token-claim-validations/{token-claim-validation-name} | Delete a Token Claim Validation
[**GetTokenClaimValidation**](TokenClaimValidationApi.md#GetTokenClaimValidation) | **Get** /token-claim-validations/{token-claim-validation-name} | Returns a single Token Claim Validation
[**UpdateTokenClaimValidation**](TokenClaimValidationApi.md#UpdateTokenClaimValidation) | **Patch** /token-claim-validations/{token-claim-validation-name} | Update an existing Token Claim Validation by name



## AddTokenClaimValidation

> AddTokenClaimValidation200Response AddTokenClaimValidation(ctx).AddTokenClaimValidationRequest(addTokenClaimValidationRequest).Execute()

Add a new Token Claim Validation to the config

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    addTokenClaimValidationRequest := openapiclient.add_token_claim_validation_request{AddBooleanTokenClaimValidationRequest: openapiclient.NewAddBooleanTokenClaimValidationRequest("ValidationName_example", []openapiclient.EnumbooleanTokenClaimValidationSchemaUrn{openapiclient.Enumboolean-token-claim-validationSchemaUrn("urn:pingidentity:schemas:configuration:2.0:token-claim-validation:boolean")}, openapiclient.Enumtoken-claim-validation-requiredValueProp("true"), "ClaimName_example")} // AddTokenClaimValidationRequest | Create a new Token Claim Validation in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenClaimValidationApi.AddTokenClaimValidation(context.Background()).AddTokenClaimValidationRequest(addTokenClaimValidationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenClaimValidationApi.AddTokenClaimValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTokenClaimValidation`: AddTokenClaimValidation200Response
    fmt.Fprintf(os.Stdout, "Response from `TokenClaimValidationApi.AddTokenClaimValidation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTokenClaimValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addTokenClaimValidationRequest** | [**AddTokenClaimValidationRequest**](AddTokenClaimValidationRequest.md) | Create a new Token Claim Validation in the config | 

### Return type

[**AddTokenClaimValidation200Response**](AddTokenClaimValidation200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTokenClaimValidation

> DeleteTokenClaimValidation(ctx, tokenClaimValidationName).Execute()

Delete a Token Claim Validation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenClaimValidationName := "tokenClaimValidationName_example" // string | Name of the Token Claim Validation to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenClaimValidationApi.DeleteTokenClaimValidation(context.Background(), tokenClaimValidationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenClaimValidationApi.DeleteTokenClaimValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenClaimValidationName** | **string** | Name of the Token Claim Validation to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenClaimValidationRequest struct via the builder pattern


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


## GetTokenClaimValidation

> AddTokenClaimValidation200Response GetTokenClaimValidation(ctx, tokenClaimValidationName).Execute()

Returns a single Token Claim Validation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenClaimValidationName := "tokenClaimValidationName_example" // string | Name of the Token Claim Validation to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenClaimValidationApi.GetTokenClaimValidation(context.Background(), tokenClaimValidationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenClaimValidationApi.GetTokenClaimValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenClaimValidation`: AddTokenClaimValidation200Response
    fmt.Fprintf(os.Stdout, "Response from `TokenClaimValidationApi.GetTokenClaimValidation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenClaimValidationName** | **string** | Name of the Token Claim Validation to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenClaimValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddTokenClaimValidation200Response**](AddTokenClaimValidation200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenClaimValidation

> AddTokenClaimValidation200Response UpdateTokenClaimValidation(ctx, tokenClaimValidationName).UpdateRequest(updateRequest).Execute()

Update an existing Token Claim Validation by name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenClaimValidationName := "tokenClaimValidationName_example" // string | Name of the Token Claim Validation to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Token Claim Validation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenClaimValidationApi.UpdateTokenClaimValidation(context.Background(), tokenClaimValidationName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenClaimValidationApi.UpdateTokenClaimValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTokenClaimValidation`: AddTokenClaimValidation200Response
    fmt.Fprintf(os.Stdout, "Response from `TokenClaimValidationApi.UpdateTokenClaimValidation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenClaimValidationName** | **string** | Name of the Token Claim Validation to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenClaimValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Token Claim Validation | 

### Return type

[**AddTokenClaimValidation200Response**](AddTokenClaimValidation200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

