# \SoftDeletePolicyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSoftDeletePolicy**](SoftDeletePolicyApi.md#AddSoftDeletePolicy) | **Post** /soft-delete-policies | Add a new Soft Delete Policy to the config
[**DeleteSoftDeletePolicy**](SoftDeletePolicyApi.md#DeleteSoftDeletePolicy) | **Delete** /soft-delete-policies/{soft-delete-policy-name} | Delete a Soft Delete Policy
[**GetSoftDeletePolicy**](SoftDeletePolicyApi.md#GetSoftDeletePolicy) | **Get** /soft-delete-policies/{soft-delete-policy-name} | Returns a single Soft Delete Policy
[**UpdateSoftDeletePolicy**](SoftDeletePolicyApi.md#UpdateSoftDeletePolicy) | **Patch** /soft-delete-policies/{soft-delete-policy-name} | Update an existing Soft Delete Policy by name



## AddSoftDeletePolicy

> SoftDeletePolicyResponse AddSoftDeletePolicy(ctx).AddSoftDeletePolicyRequest(addSoftDeletePolicyRequest).Execute()

Add a new Soft Delete Policy to the config

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
    addSoftDeletePolicyRequest := *openapiclient.NewAddSoftDeletePolicyRequest("PolicyName_example") // AddSoftDeletePolicyRequest | Create a new Soft Delete Policy in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftDeletePolicyApi.AddSoftDeletePolicy(context.Background()).AddSoftDeletePolicyRequest(addSoftDeletePolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftDeletePolicyApi.AddSoftDeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSoftDeletePolicy`: SoftDeletePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftDeletePolicyApi.AddSoftDeletePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSoftDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addSoftDeletePolicyRequest** | [**AddSoftDeletePolicyRequest**](AddSoftDeletePolicyRequest.md) | Create a new Soft Delete Policy in the config | 

### Return type

[**SoftDeletePolicyResponse**](SoftDeletePolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSoftDeletePolicy

> DeleteSoftDeletePolicy(ctx, softDeletePolicyName).Execute()

Delete a Soft Delete Policy

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
    softDeletePolicyName := "softDeletePolicyName_example" // string | Name of the Soft Delete Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SoftDeletePolicyApi.DeleteSoftDeletePolicy(context.Background(), softDeletePolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftDeletePolicyApi.DeleteSoftDeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**softDeletePolicyName** | **string** | Name of the Soft Delete Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSoftDeletePolicyRequest struct via the builder pattern


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


## GetSoftDeletePolicy

> SoftDeletePolicyResponse GetSoftDeletePolicy(ctx, softDeletePolicyName).Execute()

Returns a single Soft Delete Policy

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
    softDeletePolicyName := "softDeletePolicyName_example" // string | Name of the Soft Delete Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftDeletePolicyApi.GetSoftDeletePolicy(context.Background(), softDeletePolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftDeletePolicyApi.GetSoftDeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftDeletePolicy`: SoftDeletePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftDeletePolicyApi.GetSoftDeletePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**softDeletePolicyName** | **string** | Name of the Soft Delete Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SoftDeletePolicyResponse**](SoftDeletePolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSoftDeletePolicy

> SoftDeletePolicyResponse UpdateSoftDeletePolicy(ctx, softDeletePolicyName).UpdateRequest(updateRequest).Execute()

Update an existing Soft Delete Policy by name

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
    softDeletePolicyName := "softDeletePolicyName_example" // string | Name of the Soft Delete Policy
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Soft Delete Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftDeletePolicyApi.UpdateSoftDeletePolicy(context.Background(), softDeletePolicyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftDeletePolicyApi.UpdateSoftDeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSoftDeletePolicy`: SoftDeletePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftDeletePolicyApi.UpdateSoftDeletePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**softDeletePolicyName** | **string** | Name of the Soft Delete Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSoftDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Soft Delete Policy | 

### Return type

[**SoftDeletePolicyResponse**](SoftDeletePolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

