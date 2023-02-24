# \ConnectionCriteriaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConnectionCriteria**](ConnectionCriteriaApi.md#AddConnectionCriteria) | **Post** /connection-criteria | Add a new Connection Criteria to the config
[**DeleteConnectionCriteria**](ConnectionCriteriaApi.md#DeleteConnectionCriteria) | **Delete** /connection-criteria/{connection-criteria-name} | Delete a Connection Criteria
[**GetConnectionCriteria**](ConnectionCriteriaApi.md#GetConnectionCriteria) | **Get** /connection-criteria/{connection-criteria-name} | Returns a single Connection Criteria
[**UpdateConnectionCriteria**](ConnectionCriteriaApi.md#UpdateConnectionCriteria) | **Patch** /connection-criteria/{connection-criteria-name} | Update an existing Connection Criteria by name



## AddConnectionCriteria

> AddConnectionCriteria200Response AddConnectionCriteria(ctx).AddConnectionCriteriaRequest(addConnectionCriteriaRequest).Execute()

Add a new Connection Criteria to the config

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
    addConnectionCriteriaRequest := openapiclient.add_connection_criteria_request{AddAggregateConnectionCriteriaRequest: openapiclient.NewAddAggregateConnectionCriteriaRequest("CriteriaName_example", []openapiclient.EnumaggregateConnectionCriteriaSchemaUrn{openapiclient.Enumaggregate-connection-criteriaSchemaUrn("urn:pingidentity:schemas:configuration:2.0:connection-criteria:aggregate")})} // AddConnectionCriteriaRequest | Create a new Connection Criteria in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionCriteriaApi.AddConnectionCriteria(context.Background()).AddConnectionCriteriaRequest(addConnectionCriteriaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionCriteriaApi.AddConnectionCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConnectionCriteria`: AddConnectionCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionCriteriaApi.AddConnectionCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddConnectionCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addConnectionCriteriaRequest** | [**AddConnectionCriteriaRequest**](AddConnectionCriteriaRequest.md) | Create a new Connection Criteria in the config | 

### Return type

[**AddConnectionCriteria200Response**](AddConnectionCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectionCriteria

> DeleteConnectionCriteria(ctx, connectionCriteriaName).Execute()

Delete a Connection Criteria

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
    connectionCriteriaName := "connectionCriteriaName_example" // string | Name of the Connection Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConnectionCriteriaApi.DeleteConnectionCriteria(context.Background(), connectionCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionCriteriaApi.DeleteConnectionCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionCriteriaName** | **string** | Name of the Connection Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionCriteriaRequest struct via the builder pattern


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


## GetConnectionCriteria

> AddConnectionCriteria200Response GetConnectionCriteria(ctx, connectionCriteriaName).Execute()

Returns a single Connection Criteria

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
    connectionCriteriaName := "connectionCriteriaName_example" // string | Name of the Connection Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionCriteriaApi.GetConnectionCriteria(context.Background(), connectionCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionCriteriaApi.GetConnectionCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionCriteria`: AddConnectionCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionCriteriaApi.GetConnectionCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionCriteriaName** | **string** | Name of the Connection Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddConnectionCriteria200Response**](AddConnectionCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionCriteria

> AddConnectionCriteria200Response UpdateConnectionCriteria(ctx, connectionCriteriaName).UpdateRequest(updateRequest).Execute()

Update an existing Connection Criteria by name

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
    connectionCriteriaName := "connectionCriteriaName_example" // string | Name of the Connection Criteria
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Connection Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionCriteriaApi.UpdateConnectionCriteria(context.Background(), connectionCriteriaName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionCriteriaApi.UpdateConnectionCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectionCriteria`: AddConnectionCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionCriteriaApi.UpdateConnectionCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionCriteriaName** | **string** | Name of the Connection Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Connection Criteria | 

### Return type

[**AddConnectionCriteria200Response**](AddConnectionCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

