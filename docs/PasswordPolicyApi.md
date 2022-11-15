# \PasswordPolicyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPasswordPolicy**](PasswordPolicyApi.md#AddPasswordPolicy) | **Post** /password-policies | Add a new Password Policy to the config
[**DeletePasswordPolicy**](PasswordPolicyApi.md#DeletePasswordPolicy) | **Delete** /password-policies/{password-policy-name} | Delete a Password Policy
[**GetPasswordPolicy**](PasswordPolicyApi.md#GetPasswordPolicy) | **Get** /password-policies/{password-policy-name} | Returns a single Password Policy
[**UpdatePasswordPolicy**](PasswordPolicyApi.md#UpdatePasswordPolicy) | **Patch** /password-policies/{password-policy-name} | Update an existing Password Policy by name



## AddPasswordPolicy

> PasswordPolicyResponse AddPasswordPolicy(ctx).AddPasswordPolicyRequest(addPasswordPolicyRequest).Execute()

Add a new Password Policy to the config

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
    addPasswordPolicyRequest := *openapiclient.NewAddPasswordPolicyRequest("PolicyName_example", "PasswordAttribute_example", []string{"DefaultPasswordStorageScheme_example"}) // AddPasswordPolicyRequest | Create a new Password Policy in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPolicyApi.AddPasswordPolicy(context.Background()).AddPasswordPolicyRequest(addPasswordPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyApi.AddPasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPasswordPolicy`: PasswordPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyApi.AddPasswordPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPasswordPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPasswordPolicyRequest** | [**AddPasswordPolicyRequest**](AddPasswordPolicyRequest.md) | Create a new Password Policy in the config | 

### Return type

[**PasswordPolicyResponse**](PasswordPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePasswordPolicy

> DeletePasswordPolicy(ctx, passwordPolicyName).Execute()

Delete a Password Policy

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
    passwordPolicyName := "passwordPolicyName_example" // string | Name of the Password Policy to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPolicyApi.DeletePasswordPolicy(context.Background(), passwordPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyApi.DeletePasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordPolicyName** | **string** | Name of the Password Policy to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePasswordPolicyRequest struct via the builder pattern


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


## GetPasswordPolicy

> PasswordPolicyResponse GetPasswordPolicy(ctx, passwordPolicyName).Execute()

Returns a single Password Policy

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
    passwordPolicyName := "passwordPolicyName_example" // string | Name of the Password Policy to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPolicyApi.GetPasswordPolicy(context.Background(), passwordPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyApi.GetPasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordPolicy`: PasswordPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyApi.GetPasswordPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordPolicyName** | **string** | Name of the Password Policy to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordPolicyResponse**](PasswordPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordPolicy

> PasswordPolicyResponse UpdatePasswordPolicy(ctx, passwordPolicyName).UpdateRequest(updateRequest).Execute()

Update an existing Password Policy by name

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
    passwordPolicyName := "passwordPolicyName_example" // string | Name of the Password Policy to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Password Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPolicyApi.UpdatePasswordPolicy(context.Background(), passwordPolicyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyApi.UpdatePasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordPolicy`: PasswordPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyApi.UpdatePasswordPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordPolicyName** | **string** | Name of the Password Policy to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Password Policy | 

### Return type

[**PasswordPolicyResponse**](PasswordPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

