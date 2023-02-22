# \DelegatedAdminResourceRightsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDelegatedAdminResourceRights**](DelegatedAdminResourceRightsApi.md#AddDelegatedAdminResourceRights) | **Post** /delegated-admin-rights/{delegated-admin-rights-name}/delegated-admin-resource-rights | Add a new Delegated Admin Resource Rights to the config
[**DeleteDelegatedAdminResourceRights**](DelegatedAdminResourceRightsApi.md#DeleteDelegatedAdminResourceRights) | **Delete** /delegated-admin-rights/{delegated-admin-rights-name}/delegated-admin-resource-rights/{delegated-admin-resource-rights-name} | Delete a Delegated Admin Resource Rights
[**GetDelegatedAdminResourceRights**](DelegatedAdminResourceRightsApi.md#GetDelegatedAdminResourceRights) | **Get** /delegated-admin-rights/{delegated-admin-rights-name}/delegated-admin-resource-rights/{delegated-admin-resource-rights-name} | Returns a single Delegated Admin Resource Rights
[**UpdateDelegatedAdminResourceRights**](DelegatedAdminResourceRightsApi.md#UpdateDelegatedAdminResourceRights) | **Patch** /delegated-admin-rights/{delegated-admin-rights-name}/delegated-admin-resource-rights/{delegated-admin-resource-rights-name} | Update an existing Delegated Admin Resource Rights by name



## AddDelegatedAdminResourceRights

> DelegatedAdminResourceRightsResponse AddDelegatedAdminResourceRights(ctx, delegatedAdminRightsName).AddDelegatedAdminResourceRightsRequest(addDelegatedAdminResourceRightsRequest).Execute()

Add a new Delegated Admin Resource Rights to the config

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
    delegatedAdminRightsName := "delegatedAdminRightsName_example" // string | Name of the Delegated Admin Rights
    addDelegatedAdminResourceRightsRequest := *openapiclient.NewAddDelegatedAdminResourceRightsRequest("RestResourceType_example", false) // AddDelegatedAdminResourceRightsRequest | Create a new Delegated Admin Resource Rights in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminResourceRightsApi.AddDelegatedAdminResourceRights(context.Background(), delegatedAdminRightsName).AddDelegatedAdminResourceRightsRequest(addDelegatedAdminResourceRightsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminResourceRightsApi.AddDelegatedAdminResourceRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDelegatedAdminResourceRights`: DelegatedAdminResourceRightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminResourceRightsApi.AddDelegatedAdminResourceRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDelegatedAdminResourceRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addDelegatedAdminResourceRightsRequest** | [**AddDelegatedAdminResourceRightsRequest**](AddDelegatedAdminResourceRightsRequest.md) | Create a new Delegated Admin Resource Rights in the config | 

### Return type

[**DelegatedAdminResourceRightsResponse**](DelegatedAdminResourceRightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDelegatedAdminResourceRights

> DeleteDelegatedAdminResourceRights(ctx, delegatedAdminResourceRightsName, delegatedAdminRightsName).Execute()

Delete a Delegated Admin Resource Rights

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
    delegatedAdminResourceRightsName := "delegatedAdminResourceRightsName_example" // string | Name of the Delegated Admin Resource Rights
    delegatedAdminRightsName := "delegatedAdminRightsName_example" // string | Name of the Delegated Admin Rights

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminResourceRightsApi.DeleteDelegatedAdminResourceRights(context.Background(), delegatedAdminResourceRightsName, delegatedAdminRightsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminResourceRightsApi.DeleteDelegatedAdminResourceRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminResourceRightsName** | **string** | Name of the Delegated Admin Resource Rights | 
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDelegatedAdminResourceRightsRequest struct via the builder pattern


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


## GetDelegatedAdminResourceRights

> DelegatedAdminResourceRightsResponse GetDelegatedAdminResourceRights(ctx, delegatedAdminResourceRightsName, delegatedAdminRightsName).Execute()

Returns a single Delegated Admin Resource Rights

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
    delegatedAdminResourceRightsName := "delegatedAdminResourceRightsName_example" // string | Name of the Delegated Admin Resource Rights
    delegatedAdminRightsName := "delegatedAdminRightsName_example" // string | Name of the Delegated Admin Rights

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminResourceRightsApi.GetDelegatedAdminResourceRights(context.Background(), delegatedAdminResourceRightsName, delegatedAdminRightsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminResourceRightsApi.GetDelegatedAdminResourceRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDelegatedAdminResourceRights`: DelegatedAdminResourceRightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminResourceRightsApi.GetDelegatedAdminResourceRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminResourceRightsName** | **string** | Name of the Delegated Admin Resource Rights | 
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDelegatedAdminResourceRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DelegatedAdminResourceRightsResponse**](DelegatedAdminResourceRightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDelegatedAdminResourceRights

> DelegatedAdminResourceRightsResponse UpdateDelegatedAdminResourceRights(ctx, delegatedAdminResourceRightsName, delegatedAdminRightsName).UpdateRequest(updateRequest).Execute()

Update an existing Delegated Admin Resource Rights by name

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
    delegatedAdminResourceRightsName := "delegatedAdminResourceRightsName_example" // string | Name of the Delegated Admin Resource Rights
    delegatedAdminRightsName := "delegatedAdminRightsName_example" // string | Name of the Delegated Admin Rights
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Delegated Admin Resource Rights

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminResourceRightsApi.UpdateDelegatedAdminResourceRights(context.Background(), delegatedAdminResourceRightsName, delegatedAdminRightsName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminResourceRightsApi.UpdateDelegatedAdminResourceRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDelegatedAdminResourceRights`: DelegatedAdminResourceRightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminResourceRightsApi.UpdateDelegatedAdminResourceRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminResourceRightsName** | **string** | Name of the Delegated Admin Resource Rights | 
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDelegatedAdminResourceRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Delegated Admin Resource Rights | 

### Return type

[**DelegatedAdminResourceRightsResponse**](DelegatedAdminResourceRightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

