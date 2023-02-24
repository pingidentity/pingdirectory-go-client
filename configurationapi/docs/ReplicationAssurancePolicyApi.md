# \ReplicationAssurancePolicyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddReplicationAssurancePolicy**](ReplicationAssurancePolicyApi.md#AddReplicationAssurancePolicy) | **Post** /replication-assurance-policies | Add a new Replication Assurance Policy to the config
[**DeleteReplicationAssurancePolicy**](ReplicationAssurancePolicyApi.md#DeleteReplicationAssurancePolicy) | **Delete** /replication-assurance-policies/{replication-assurance-policy-name} | Delete a Replication Assurance Policy
[**GetReplicationAssurancePolicy**](ReplicationAssurancePolicyApi.md#GetReplicationAssurancePolicy) | **Get** /replication-assurance-policies/{replication-assurance-policy-name} | Returns a single Replication Assurance Policy
[**UpdateReplicationAssurancePolicy**](ReplicationAssurancePolicyApi.md#UpdateReplicationAssurancePolicy) | **Patch** /replication-assurance-policies/{replication-assurance-policy-name} | Update an existing Replication Assurance Policy by name



## AddReplicationAssurancePolicy

> ReplicationAssurancePolicyResponse AddReplicationAssurancePolicy(ctx).AddReplicationAssurancePolicyRequest(addReplicationAssurancePolicyRequest).Execute()

Add a new Replication Assurance Policy to the config

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
    addReplicationAssurancePolicyRequest := *openapiclient.NewAddReplicationAssurancePolicyRequest("PolicyName_example", int32(123), "Timeout_example") // AddReplicationAssurancePolicyRequest | Create a new Replication Assurance Policy in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicationAssurancePolicyApi.AddReplicationAssurancePolicy(context.Background()).AddReplicationAssurancePolicyRequest(addReplicationAssurancePolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAssurancePolicyApi.AddReplicationAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddReplicationAssurancePolicy`: ReplicationAssurancePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ReplicationAssurancePolicyApi.AddReplicationAssurancePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddReplicationAssurancePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addReplicationAssurancePolicyRequest** | [**AddReplicationAssurancePolicyRequest**](AddReplicationAssurancePolicyRequest.md) | Create a new Replication Assurance Policy in the config | 

### Return type

[**ReplicationAssurancePolicyResponse**](ReplicationAssurancePolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReplicationAssurancePolicy

> DeleteReplicationAssurancePolicy(ctx, replicationAssurancePolicyName).Execute()

Delete a Replication Assurance Policy

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
    replicationAssurancePolicyName := "replicationAssurancePolicyName_example" // string | Name of the Replication Assurance Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReplicationAssurancePolicyApi.DeleteReplicationAssurancePolicy(context.Background(), replicationAssurancePolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAssurancePolicyApi.DeleteReplicationAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicationAssurancePolicyName** | **string** | Name of the Replication Assurance Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReplicationAssurancePolicyRequest struct via the builder pattern


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


## GetReplicationAssurancePolicy

> ReplicationAssurancePolicyResponse GetReplicationAssurancePolicy(ctx, replicationAssurancePolicyName).Execute()

Returns a single Replication Assurance Policy

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
    replicationAssurancePolicyName := "replicationAssurancePolicyName_example" // string | Name of the Replication Assurance Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicationAssurancePolicyApi.GetReplicationAssurancePolicy(context.Background(), replicationAssurancePolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAssurancePolicyApi.GetReplicationAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplicationAssurancePolicy`: ReplicationAssurancePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ReplicationAssurancePolicyApi.GetReplicationAssurancePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicationAssurancePolicyName** | **string** | Name of the Replication Assurance Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplicationAssurancePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReplicationAssurancePolicyResponse**](ReplicationAssurancePolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReplicationAssurancePolicy

> ReplicationAssurancePolicyResponse UpdateReplicationAssurancePolicy(ctx, replicationAssurancePolicyName).UpdateRequest(updateRequest).Execute()

Update an existing Replication Assurance Policy by name

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
    replicationAssurancePolicyName := "replicationAssurancePolicyName_example" // string | Name of the Replication Assurance Policy
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Replication Assurance Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicationAssurancePolicyApi.UpdateReplicationAssurancePolicy(context.Background(), replicationAssurancePolicyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationAssurancePolicyApi.UpdateReplicationAssurancePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReplicationAssurancePolicy`: ReplicationAssurancePolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ReplicationAssurancePolicyApi.UpdateReplicationAssurancePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicationAssurancePolicyName** | **string** | Name of the Replication Assurance Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReplicationAssurancePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Replication Assurance Policy | 

### Return type

[**ReplicationAssurancePolicyResponse**](ReplicationAssurancePolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

