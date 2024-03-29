# \ReplicationDomainAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReplicationDomain**](ReplicationDomainAPI.md#GetReplicationDomain) | **Get** /synchronization-providers/{synchronization-provider-name}/replication-domains/{replication-domain-name} | Returns a single Replication Domain
[**ListReplicationDomains**](ReplicationDomainAPI.md#ListReplicationDomains) | **Get** /synchronization-providers/{synchronization-provider-name}/replication-domains | Returns a list of all Replication Domain objects
[**UpdateReplicationDomain**](ReplicationDomainAPI.md#UpdateReplicationDomain) | **Patch** /synchronization-providers/{synchronization-provider-name}/replication-domains/{replication-domain-name} | Update an existing Replication Domain by name



## GetReplicationDomain

> ReplicationDomainResponse GetReplicationDomain(ctx, replicationDomainName, synchronizationProviderName).Execute()

Returns a single Replication Domain

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
    replicationDomainName := "replicationDomainName_example" // string | Name of the Replication Domain
    synchronizationProviderName := "synchronizationProviderName_example" // string | Name of the Synchronization Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicationDomainAPI.GetReplicationDomain(context.Background(), replicationDomainName, synchronizationProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationDomainAPI.GetReplicationDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplicationDomain`: ReplicationDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `ReplicationDomainAPI.GetReplicationDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicationDomainName** | **string** | Name of the Replication Domain | 
**synchronizationProviderName** | **string** | Name of the Synchronization Provider | 

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


## ListReplicationDomains

> ReplicationDomainListResponse ListReplicationDomains(ctx, synchronizationProviderName).Filter(filter).Execute()

Returns a list of all Replication Domain objects

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
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicationDomainAPI.ListReplicationDomains(context.Background(), synchronizationProviderName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationDomainAPI.ListReplicationDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReplicationDomains`: ReplicationDomainListResponse
    fmt.Fprintf(os.Stdout, "Response from `ReplicationDomainAPI.ListReplicationDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**synchronizationProviderName** | **string** | Name of the Synchronization Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReplicationDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**ReplicationDomainListResponse**](ReplicationDomainListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReplicationDomain

> ReplicationDomainResponse UpdateReplicationDomain(ctx, replicationDomainName, synchronizationProviderName).UpdateRequest(updateRequest).Execute()

Update an existing Replication Domain by name

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
    replicationDomainName := "replicationDomainName_example" // string | Name of the Replication Domain
    synchronizationProviderName := "synchronizationProviderName_example" // string | Name of the Synchronization Provider
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Replication Domain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicationDomainAPI.UpdateReplicationDomain(context.Background(), replicationDomainName, synchronizationProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationDomainAPI.UpdateReplicationDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReplicationDomain`: ReplicationDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `ReplicationDomainAPI.UpdateReplicationDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicationDomainName** | **string** | Name of the Replication Domain | 
**synchronizationProviderName** | **string** | Name of the Synchronization Provider | 

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

