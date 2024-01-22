# \SoftDeletePolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSoftDeletePolicy**](SoftDeletePolicyAPI.md#AddSoftDeletePolicy) | **Post** /soft-delete-policies | Add a new Soft Delete Policy to the config
[**DeleteSoftDeletePolicy**](SoftDeletePolicyAPI.md#DeleteSoftDeletePolicy) | **Delete** /soft-delete-policies/{soft-delete-policy-name} | Delete a Soft Delete Policy
[**GetSoftDeletePolicy**](SoftDeletePolicyAPI.md#GetSoftDeletePolicy) | **Get** /soft-delete-policies/{soft-delete-policy-name} | Returns a single Soft Delete Policy
[**ListSoftDeletePolicies**](SoftDeletePolicyAPI.md#ListSoftDeletePolicies) | **Get** /soft-delete-policies | Returns a list of all Soft Delete Policy objects
[**UpdateSoftDeletePolicy**](SoftDeletePolicyAPI.md#UpdateSoftDeletePolicy) | **Patch** /soft-delete-policies/{soft-delete-policy-name} | Update an existing Soft Delete Policy by name



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
    resp, r, err := apiClient.SoftDeletePolicyAPI.AddSoftDeletePolicy(context.Background()).AddSoftDeletePolicyRequest(addSoftDeletePolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftDeletePolicyAPI.AddSoftDeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSoftDeletePolicy`: SoftDeletePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftDeletePolicyAPI.AddSoftDeletePolicy`: %v\n", resp)
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
    r, err := apiClient.SoftDeletePolicyAPI.DeleteSoftDeletePolicy(context.Background(), softDeletePolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftDeletePolicyAPI.DeleteSoftDeletePolicy``: %v\n", err)
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
    resp, r, err := apiClient.SoftDeletePolicyAPI.GetSoftDeletePolicy(context.Background(), softDeletePolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftDeletePolicyAPI.GetSoftDeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoftDeletePolicy`: SoftDeletePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftDeletePolicyAPI.GetSoftDeletePolicy`: %v\n", resp)
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


## ListSoftDeletePolicies

> SoftDeletePolicyListResponse ListSoftDeletePolicies(ctx).Filter(filter).Execute()

Returns a list of all Soft Delete Policy objects

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
    resp, r, err := apiClient.SoftDeletePolicyAPI.ListSoftDeletePolicies(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftDeletePolicyAPI.ListSoftDeletePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSoftDeletePolicies`: SoftDeletePolicyListResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftDeletePolicyAPI.ListSoftDeletePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSoftDeletePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**SoftDeletePolicyListResponse**](SoftDeletePolicyListResponse.md)

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
    resp, r, err := apiClient.SoftDeletePolicyAPI.UpdateSoftDeletePolicy(context.Background(), softDeletePolicyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftDeletePolicyAPI.UpdateSoftDeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSoftDeletePolicy`: SoftDeletePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `SoftDeletePolicyAPI.UpdateSoftDeletePolicy`: %v\n", resp)
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

