# \ReplicationDomainApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReplicationDomain**](ReplicationDomainApi.md#GetReplicationDomain) | **Get** /replication-domains/{replication-domain-name} | Returns a single Replication Domain
[**UpdateReplicationDomain**](ReplicationDomainApi.md#UpdateReplicationDomain) | **Patch** /replication-domains/{replication-domain-name} | Update an existing Replication Domain by name



## GetReplicationDomain

> ReplicationDomainResponse GetReplicationDomain(ctx, replicationDomainName).Execute()

Returns a single Replication Domain

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
    replicationDomainName := "replicationDomainName_example" // string | Name of the Replication Domain to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicationDomainApi.GetReplicationDomain(context.Background(), replicationDomainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationDomainApi.GetReplicationDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplicationDomain`: ReplicationDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `ReplicationDomainApi.GetReplicationDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicationDomainName** | **string** | Name of the Replication Domain to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplicationDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReplicationDomainResponse**](ReplicationDomainResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReplicationDomain

> ReplicationDomainResponse UpdateReplicationDomain(ctx, replicationDomainName).UpdateRequest(updateRequest).Execute()

Update an existing Replication Domain by name

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
    replicationDomainName := "replicationDomainName_example" // string | Name of the Replication Domain to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Replication Domain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicationDomainApi.UpdateReplicationDomain(context.Background(), replicationDomainName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationDomainApi.UpdateReplicationDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReplicationDomain`: ReplicationDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `ReplicationDomainApi.UpdateReplicationDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicationDomainName** | **string** | Name of the Replication Domain to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReplicationDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Replication Domain | 

### Return type

[**ReplicationDomainResponse**](ReplicationDomainResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

