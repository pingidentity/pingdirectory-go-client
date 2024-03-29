# \TokenClaimValidationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTokenClaimValidation**](TokenClaimValidationAPI.md#AddTokenClaimValidation) | **Post** /id-token-validators/{id-token-validator-name}/token-claim-validations | Add a new Token Claim Validation to the config
[**DeleteTokenClaimValidation**](TokenClaimValidationAPI.md#DeleteTokenClaimValidation) | **Delete** /id-token-validators/{id-token-validator-name}/token-claim-validations/{token-claim-validation-name} | Delete a Token Claim Validation
[**GetTokenClaimValidation**](TokenClaimValidationAPI.md#GetTokenClaimValidation) | **Get** /id-token-validators/{id-token-validator-name}/token-claim-validations/{token-claim-validation-name} | Returns a single Token Claim Validation
[**ListTokenClaimValidations**](TokenClaimValidationAPI.md#ListTokenClaimValidations) | **Get** /id-token-validators/{id-token-validator-name}/token-claim-validations | Returns a list of all Token Claim Validation objects
[**UpdateTokenClaimValidation**](TokenClaimValidationAPI.md#UpdateTokenClaimValidation) | **Patch** /id-token-validators/{id-token-validator-name}/token-claim-validations/{token-claim-validation-name} | Update an existing Token Claim Validation by name



## AddTokenClaimValidation

> AddTokenClaimValidation200Response AddTokenClaimValidation(ctx, idTokenValidatorName).AddTokenClaimValidationRequest(addTokenClaimValidationRequest).Execute()

Add a new Token Claim Validation to the config

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
    addTokenClaimValidationRequest := openapiclient.add_token_claim_validation_request{AddBooleanTokenClaimValidationRequest: openapiclient.NewAddBooleanTokenClaimValidationRequest([]openapiclient.EnumbooleanTokenClaimValidationSchemaUrn{openapiclient.Enumboolean-token-claim-validationSchemaUrn("urn:pingidentity:schemas:configuration:2.0:token-claim-validation:boolean")}, openapiclient.Enumtoken-claim-validation-requiredValueProp("true"), "ClaimName_example", "ValidationName_example")} // AddTokenClaimValidationRequest | Create a new Token Claim Validation in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenClaimValidationAPI.AddTokenClaimValidation(context.Background(), idTokenValidatorName).AddTokenClaimValidationRequest(addTokenClaimValidationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenClaimValidationAPI.AddTokenClaimValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTokenClaimValidation`: AddTokenClaimValidation200Response
    fmt.Fprintf(os.Stdout, "Response from `TokenClaimValidationAPI.AddTokenClaimValidation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idTokenValidatorName** | **string** | Name of the ID Token Validator | 

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

> DeleteTokenClaimValidation(ctx, tokenClaimValidationName, idTokenValidatorName).Execute()

Delete a Token Claim Validation

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
    tokenClaimValidationName := "tokenClaimValidationName_example" // string | Name of the Token Claim Validation
    idTokenValidatorName := "idTokenValidatorName_example" // string | Name of the ID Token Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TokenClaimValidationAPI.DeleteTokenClaimValidation(context.Background(), tokenClaimValidationName, idTokenValidatorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenClaimValidationAPI.DeleteTokenClaimValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenClaimValidationName** | **string** | Name of the Token Claim Validation | 
**idTokenValidatorName** | **string** | Name of the ID Token Validator | 

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

> AddTokenClaimValidation200Response GetTokenClaimValidation(ctx, tokenClaimValidationName, idTokenValidatorName).Execute()

Returns a single Token Claim Validation

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
    tokenClaimValidationName := "tokenClaimValidationName_example" // string | Name of the Token Claim Validation
    idTokenValidatorName := "idTokenValidatorName_example" // string | Name of the ID Token Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenClaimValidationAPI.GetTokenClaimValidation(context.Background(), tokenClaimValidationName, idTokenValidatorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenClaimValidationAPI.GetTokenClaimValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenClaimValidation`: AddTokenClaimValidation200Response
    fmt.Fprintf(os.Stdout, "Response from `TokenClaimValidationAPI.GetTokenClaimValidation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenClaimValidationName** | **string** | Name of the Token Claim Validation | 
**idTokenValidatorName** | **string** | Name of the ID Token Validator | 

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


## ListTokenClaimValidations

> TokenClaimValidationListResponse ListTokenClaimValidations(ctx, idTokenValidatorName).Filter(filter).Execute()

Returns a list of all Token Claim Validation objects

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
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenClaimValidationAPI.ListTokenClaimValidations(context.Background(), idTokenValidatorName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenClaimValidationAPI.ListTokenClaimValidations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokenClaimValidations`: TokenClaimValidationListResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenClaimValidationAPI.ListTokenClaimValidations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idTokenValidatorName** | **string** | Name of the ID Token Validator | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTokenClaimValidationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**TokenClaimValidationListResponse**](TokenClaimValidationListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTokenClaimValidation

> AddTokenClaimValidation200Response UpdateTokenClaimValidation(ctx, tokenClaimValidationName, idTokenValidatorName).UpdateRequest(updateRequest).Execute()

Update an existing Token Claim Validation by name

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
    tokenClaimValidationName := "tokenClaimValidationName_example" // string | Name of the Token Claim Validation
    idTokenValidatorName := "idTokenValidatorName_example" // string | Name of the ID Token Validator
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Token Claim Validation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenClaimValidationAPI.UpdateTokenClaimValidation(context.Background(), tokenClaimValidationName, idTokenValidatorName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenClaimValidationAPI.UpdateTokenClaimValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTokenClaimValidation`: AddTokenClaimValidation200Response
    fmt.Fprintf(os.Stdout, "Response from `TokenClaimValidationAPI.UpdateTokenClaimValidation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenClaimValidationName** | **string** | Name of the Token Claim Validation | 
**idTokenValidatorName** | **string** | Name of the ID Token Validator | 

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

