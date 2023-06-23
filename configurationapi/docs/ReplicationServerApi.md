# \ReplicationServerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReplicationServer**](ReplicationServerApi.md#GetReplicationServer) | **Get** /synchronization-providers/{synchronization-provider-name}/replication-server | Returns a single Replication Server
[**UpdateReplicationServer**](ReplicationServerApi.md#UpdateReplicationServer) | **Patch** /synchronization-providers/{synchronization-provider-name}/replication-server | Update an existing Replication Server by name



## GetReplicationServer

> ReplicationServerResponse GetReplicationServer(ctx, synchronizationProviderName).Execute()

Returns a single Replication Server

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
    resp, r, err := apiClient.ReplicationServerApi.GetReplicationServer(context.Background(), synchronizationProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationServerApi.GetReplicationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplicationServer`: ReplicationServerResponse
    fmt.Fprintf(os.Stdout, "Response from `ReplicationServerApi.GetReplicationServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**synchronizationProviderName** | **string** | Name of the Synchronization Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplicationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReplicationServerResponse**](ReplicationServerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReplicationServer

> ReplicationServerResponse UpdateReplicationServer(ctx, synchronizationProviderName).UpdateRequest(updateRequest).Execute()

Update an existing Replication Server by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Replication Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicationServerApi.UpdateReplicationServer(context.Background(), synchronizationProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationServerApi.UpdateReplicationServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReplicationServer`: ReplicationServerResponse
    fmt.Fprintf(os.Stdout, "Response from `ReplicationServerApi.UpdateReplicationServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**synchronizationProviderName** | **string** | Name of the Synchronization Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReplicationServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Replication Server | 

### Return type

[**ReplicationServerResponse**](ReplicationServerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

