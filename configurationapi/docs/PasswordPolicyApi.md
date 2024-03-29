# \PasswordPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPasswordPolicy**](PasswordPolicyAPI.md#AddPasswordPolicy) | **Post** /password-policies | Add a new Password Policy to the config
[**DeletePasswordPolicy**](PasswordPolicyAPI.md#DeletePasswordPolicy) | **Delete** /password-policies/{password-policy-name} | Delete a Password Policy
[**GetPasswordPolicy**](PasswordPolicyAPI.md#GetPasswordPolicy) | **Get** /password-policies/{password-policy-name} | Returns a single Password Policy
[**ListPasswordPolicies**](PasswordPolicyAPI.md#ListPasswordPolicies) | **Get** /password-policies | Returns a list of all Password Policy objects
[**UpdatePasswordPolicy**](PasswordPolicyAPI.md#UpdatePasswordPolicy) | **Patch** /password-policies/{password-policy-name} | Update an existing Password Policy by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addPasswordPolicyRequest := *openapiclient.NewAddPasswordPolicyRequest("PasswordAttribute_example", []string{"DefaultPasswordStorageScheme_example"}, "PolicyName_example") // AddPasswordPolicyRequest | Create a new Password Policy in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPolicyAPI.AddPasswordPolicy(context.Background()).AddPasswordPolicyRequest(addPasswordPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyAPI.AddPasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPasswordPolicy`: PasswordPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyAPI.AddPasswordPolicy`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    passwordPolicyName := "passwordPolicyName_example" // string | Name of the Password Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PasswordPolicyAPI.DeletePasswordPolicy(context.Background(), passwordPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyAPI.DeletePasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordPolicyName** | **string** | Name of the Password Policy | 

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    passwordPolicyName := "passwordPolicyName_example" // string | Name of the Password Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPolicyAPI.GetPasswordPolicy(context.Background(), passwordPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyAPI.GetPasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordPolicy`: PasswordPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyAPI.GetPasswordPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordPolicyName** | **string** | Name of the Password Policy | 

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


## ListPasswordPolicies

> PasswordPolicyListResponse ListPasswordPolicies(ctx).Filter(filter).Execute()

Returns a list of all Password Policy objects

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
    resp, r, err := apiClient.PasswordPolicyAPI.ListPasswordPolicies(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyAPI.ListPasswordPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPasswordPolicies`: PasswordPolicyListResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyAPI.ListPasswordPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPasswordPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**PasswordPolicyListResponse**](PasswordPolicyListResponse.md)

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    passwordPolicyName := "passwordPolicyName_example" // string | Name of the Password Policy
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Password Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPolicyAPI.UpdatePasswordPolicy(context.Background(), passwordPolicyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyAPI.UpdatePasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordPolicy`: PasswordPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyAPI.UpdatePasswordPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordPolicyName** | **string** | Name of the Password Policy | 

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

