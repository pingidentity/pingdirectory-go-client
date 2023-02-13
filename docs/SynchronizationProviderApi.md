# \SynchronizationProviderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSynchronizationProvider**](SynchronizationProviderApi.md#GetSynchronizationProvider) | **Get** /synchronization-providers/{synchronization-provider-name} | Returns a single Synchronization Provider
[**UpdateSynchronizationProvider**](SynchronizationProviderApi.md#UpdateSynchronizationProvider) | **Patch** /synchronization-providers/{synchronization-provider-name} | Update an existing Synchronization Provider by name



## GetSynchronizationProvider

> ReplicationSynchronizationProviderResponse GetSynchronizationProvider(ctx, synchronizationProviderName).Execute()

Returns a single Synchronization Provider

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
    synchronizationProviderName := "synchronizationProviderName_example" // string | Name of the Synchronization Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SynchronizationProviderApi.GetSynchronizationProvider(context.Background(), synchronizationProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationProviderApi.GetSynchronizationProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSynchronizationProvider`: ReplicationSynchronizationProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `SynchronizationProviderApi.GetSynchronizationProvider`: %v\n", resp)
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

[**ReplicationSynchronizationProviderResponse**](ReplicationSynchronizationProviderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSynchronizationProvider

> ReplicationSynchronizationProviderResponse UpdateSynchronizationProvider(ctx, synchronizationProviderName).UpdateRequest(updateRequest).Execute()

Update an existing Synchronization Provider by name

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
    synchronizationProviderName := "synchronizationProviderName_example" // string | Name of the Synchronization Provider
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Synchronization Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SynchronizationProviderApi.UpdateSynchronizationProvider(context.Background(), synchronizationProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationProviderApi.UpdateSynchronizationProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSynchronizationProvider`: ReplicationSynchronizationProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `SynchronizationProviderApi.UpdateSynchronizationProvider`: %v\n", resp)
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

[**ReplicationSynchronizationProviderResponse**](ReplicationSynchronizationProviderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

