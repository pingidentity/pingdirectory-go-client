# \PasswordValidatorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPasswordValidator**](PasswordValidatorAPI.md#AddPasswordValidator) | **Post** /password-validators | Add a new Password Validator to the config
[**DeletePasswordValidator**](PasswordValidatorAPI.md#DeletePasswordValidator) | **Delete** /password-validators/{password-validator-name} | Delete a Password Validator
[**GetPasswordValidator**](PasswordValidatorAPI.md#GetPasswordValidator) | **Get** /password-validators/{password-validator-name} | Returns a single Password Validator
[**ListPasswordValidators**](PasswordValidatorAPI.md#ListPasswordValidators) | **Get** /password-validators | Returns a list of all Password Validator objects
[**UpdatePasswordValidator**](PasswordValidatorAPI.md#UpdatePasswordValidator) | **Patch** /password-validators/{password-validator-name} | Update an existing Password Validator by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addPasswordValidatorRequest := openapiclient.add_password_validator_request{AddAttributeValuePasswordValidatorRequest: openapiclient.NewAddAttributeValuePasswordValidatorRequest([]openapiclient.EnumattributeValuePasswordValidatorSchemaUrn{openapiclient.Enumattribute-value-password-validatorSchemaUrn("urn:pingidentity:schemas:configuration:2.0:password-validator:attribute-value")}, false, false, "ValidatorName_example")} // AddPasswordValidatorRequest | Create a new Password Validator in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordValidatorAPI.AddPasswordValidator(context.Background()).AddPasswordValidatorRequest(addPasswordValidatorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordValidatorAPI.AddPasswordValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPasswordValidator`: AddPasswordValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordValidatorAPI.AddPasswordValidator`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    passwordValidatorName := "passwordValidatorName_example" // string | Name of the Password Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PasswordValidatorAPI.DeletePasswordValidator(context.Background(), passwordValidatorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordValidatorAPI.DeletePasswordValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordValidatorName** | **string** | Name of the Password Validator | 

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

> GetPasswordValidator200Response GetPasswordValidator(ctx, passwordValidatorName).Execute()

Returns a single Password Validator

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
    passwordValidatorName := "passwordValidatorName_example" // string | Name of the Password Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordValidatorAPI.GetPasswordValidator(context.Background(), passwordValidatorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordValidatorAPI.GetPasswordValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordValidator`: GetPasswordValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordValidatorAPI.GetPasswordValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordValidatorName** | **string** | Name of the Password Validator | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPasswordValidator200Response**](GetPasswordValidator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPasswordValidators

> PasswordValidatorListResponse ListPasswordValidators(ctx).Filter(filter).Execute()

Returns a list of all Password Validator objects

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
    resp, r, err := apiClient.PasswordValidatorAPI.ListPasswordValidators(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordValidatorAPI.ListPasswordValidators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPasswordValidators`: PasswordValidatorListResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordValidatorAPI.ListPasswordValidators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPasswordValidatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**PasswordValidatorListResponse**](PasswordValidatorListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordValidator

> GetPasswordValidator200Response UpdatePasswordValidator(ctx, passwordValidatorName).UpdateRequest(updateRequest).Execute()

Update an existing Password Validator by name

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
    passwordValidatorName := "passwordValidatorName_example" // string | Name of the Password Validator
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Password Validator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordValidatorAPI.UpdatePasswordValidator(context.Background(), passwordValidatorName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordValidatorAPI.UpdatePasswordValidator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordValidator`: GetPasswordValidator200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordValidatorAPI.UpdatePasswordValidator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordValidatorName** | **string** | Name of the Password Validator | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Password Validator | 

### Return type

[**GetPasswordValidator200Response**](GetPasswordValidator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

