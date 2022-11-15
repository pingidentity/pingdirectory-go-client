# \PasswordValidatorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPasswordValidator**](PasswordValidatorApi.md#AddPasswordValidator) | **Post** /password-validators | Add a new Password Validator to the config
[**DeletePasswordValidator**](PasswordValidatorApi.md#DeletePasswordValidator) | **Delete** /password-validators/{password-validator-name} | Delete a Password Validator
[**GetPasswordValidator**](PasswordValidatorApi.md#GetPasswordValidator) | **Get** /password-validators/{password-validator-name} | Returns a single Password Validator
[**UpdatePasswordValidator**](PasswordValidatorApi.md#UpdatePasswordValidator) | **Patch** /password-validators/{password-validator-name} | Update an existing Password Validator by name



## AddPasswordValidator

> AddPasswordValidator200Response AddPasswordValidator(ctx).AddPasswordValidatorRequest(addPasswordValidatorRequest).Execute()

Add a new Password Validator to the config

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
    addPasswordValidatorRequest := openapiclient.add_password_validator_request{AddAttributeValuePasswordValidatorRequest: openapiclient.NewAddAttributeValuePasswordValidatorRequest("ValidatorName_example", []openapiclient.EnumattributeValuePasswordValidatorSchemaUrn{openapiclient.Enumattribute-value-password-validatorSchemaUrn("urn:pingidentity:schemas:configuration:2.0:password-validator:attribute-value")}, false, false)} // AddPasswordValidatorRequest | Create a new Password Validator in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordValidatorApi.AddPasswordValidator(context.Background()).AddPasswordValidatorRequest(addPasswordValidatorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordValidatorApi.AddPasswordValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPasswordValidator`: AddPasswordValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordValidatorApi.AddPasswordValidator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPasswordValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPasswordValidatorRequest** | [**AddPasswordValidatorRequest**](AddPasswordValidatorRequest.md) | Create a new Password Validator in the config | 

### Return type

[**AddPasswordValidator200Response**](AddPasswordValidator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePasswordValidator

> DeletePasswordValidator(ctx, passwordValidatorName).Execute()

Delete a Password Validator

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
    passwordValidatorName := "passwordValidatorName_example" // string | Name of the Password Validator to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordValidatorApi.DeletePasswordValidator(context.Background(), passwordValidatorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordValidatorApi.DeletePasswordValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordValidatorName** | **string** | Name of the Password Validator to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePasswordValidatorRequest struct via the builder pattern


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


## GetPasswordValidator

> AddPasswordValidator200Response GetPasswordValidator(ctx, passwordValidatorName).Execute()

Returns a single Password Validator

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
    passwordValidatorName := "passwordValidatorName_example" // string | Name of the Password Validator to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordValidatorApi.GetPasswordValidator(context.Background(), passwordValidatorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordValidatorApi.GetPasswordValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordValidator`: AddPasswordValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordValidatorApi.GetPasswordValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordValidatorName** | **string** | Name of the Password Validator to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddPasswordValidator200Response**](AddPasswordValidator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordValidator

> AddPasswordValidator200Response UpdatePasswordValidator(ctx, passwordValidatorName).UpdateRequest(updateRequest).Execute()

Update an existing Password Validator by name

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
    passwordValidatorName := "passwordValidatorName_example" // string | Name of the Password Validator to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Password Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordValidatorApi.UpdatePasswordValidator(context.Background(), passwordValidatorName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordValidatorApi.UpdatePasswordValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordValidator`: AddPasswordValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordValidatorApi.UpdatePasswordValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordValidatorName** | **string** | Name of the Password Validator to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Password Validator | 

### Return type

[**AddPasswordValidator200Response**](AddPasswordValidator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

