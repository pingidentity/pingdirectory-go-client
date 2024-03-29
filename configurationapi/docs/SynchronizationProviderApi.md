# \SynchronizationProviderAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSynchronizationProvider**](SynchronizationProviderAPI.md#GetSynchronizationProvider) | **Get** /synchronization-providers/{synchronization-provider-name} | Returns a single Synchronization Provider
[**ListSynchronizationProviders**](SynchronizationProviderAPI.md#ListSynchronizationProviders) | **Get** /synchronization-providers | Returns a list of all Synchronization Provider objects
[**UpdateSynchronizationProvider**](SynchronizationProviderAPI.md#UpdateSynchronizationProvider) | **Patch** /synchronization-providers/{synchronization-provider-name} | Update an existing Synchronization Provider by name



## GetSynchronizationProvider

> GetSynchronizationProvider200Response GetSynchronizationProvider(ctx, synchronizationProviderName).Execute()

Returns a single Synchronization Provider

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
    synchronizationProviderName := "synchronizationProviderName_example" // string | Name of the Synchronization Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SynchronizationProviderAPI.GetSynchronizationProvider(context.Background(), synchronizationProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationProviderAPI.GetSynchronizationProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSynchronizationProvider`: GetSynchronizationProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `SynchronizationProviderAPI.GetSynchronizationProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**synchronizationProviderName** | **string** | Name of the Synchronization Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSynchronizationProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSynchronizationProvider200Response**](GetSynchronizationProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSynchronizationProviders

> SynchronizationProviderListResponse ListSynchronizationProviders(ctx).Filter(filter).Execute()

Returns a list of all Synchronization Provider objects

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
    resp, r, err := apiClient.SynchronizationProviderAPI.ListSynchronizationProviders(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationProviderAPI.ListSynchronizationProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSynchronizationProviders`: SynchronizationProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `SynchronizationProviderAPI.ListSynchronizationProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSynchronizationProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**SynchronizationProviderListResponse**](SynchronizationProviderListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSynchronizationProvider

> GetSynchronizationProvider200Response UpdateSynchronizationProvider(ctx, synchronizationProviderName).UpdateRequest(updateRequest).Execute()

Update an existing Synchronization Provider by name

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
    synchronizationProviderName := "synchronizationProviderName_example" // string | Name of the Synchronization Provider
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Synchronization Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SynchronizationProviderAPI.UpdateSynchronizationProvider(context.Background(), synchronizationProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationProviderAPI.UpdateSynchronizationProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSynchronizationProvider`: GetSynchronizationProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `SynchronizationProviderAPI.UpdateSynchronizationProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**synchronizationProviderName** | **string** | Name of the Synchronization Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSynchronizationProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Synchronization Provider | 

### Return type

[**GetSynchronizationProvider200Response**](GetSynchronizationProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

