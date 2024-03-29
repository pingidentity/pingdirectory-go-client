# \MonitoringEndpointAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMonitoringEndpoint**](MonitoringEndpointAPI.md#AddMonitoringEndpoint) | **Post** /monitoring-endpoints | Add a new Monitoring Endpoint to the config
[**DeleteMonitoringEndpoint**](MonitoringEndpointAPI.md#DeleteMonitoringEndpoint) | **Delete** /monitoring-endpoints/{monitoring-endpoint-name} | Delete a Monitoring Endpoint
[**GetMonitoringEndpoint**](MonitoringEndpointAPI.md#GetMonitoringEndpoint) | **Get** /monitoring-endpoints/{monitoring-endpoint-name} | Returns a single Monitoring Endpoint
[**ListMonitoringEndpoints**](MonitoringEndpointAPI.md#ListMonitoringEndpoints) | **Get** /monitoring-endpoints | Returns a list of all Monitoring Endpoint objects
[**UpdateMonitoringEndpoint**](MonitoringEndpointAPI.md#UpdateMonitoringEndpoint) | **Patch** /monitoring-endpoints/{monitoring-endpoint-name} | Update an existing Monitoring Endpoint by name



## AddMonitoringEndpoint

> StatsdMonitoringEndpointResponse AddMonitoringEndpoint(ctx).AddStatsdMonitoringEndpointRequest(addStatsdMonitoringEndpointRequest).Execute()

Add a new Monitoring Endpoint to the config

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
    addStatsdMonitoringEndpointRequest := *openapiclient.NewAddStatsdMonitoringEndpointRequest([]openapiclient.EnumstatsdMonitoringEndpointSchemaUrn{openapiclient.Enumstatsd-monitoring-endpointSchemaUrn("urn:pingidentity:schemas:configuration:2.0:monitoring-endpoint:statsd")}, "Hostname_example", false, "EndpointName_example") // AddStatsdMonitoringEndpointRequest | Create a new Monitoring Endpoint in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringEndpointAPI.AddMonitoringEndpoint(context.Background()).AddStatsdMonitoringEndpointRequest(addStatsdMonitoringEndpointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringEndpointAPI.AddMonitoringEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMonitoringEndpoint`: StatsdMonitoringEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringEndpointAPI.AddMonitoringEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMonitoringEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addStatsdMonitoringEndpointRequest** | [**AddStatsdMonitoringEndpointRequest**](AddStatsdMonitoringEndpointRequest.md) | Create a new Monitoring Endpoint in the config | 

### Return type

[**StatsdMonitoringEndpointResponse**](StatsdMonitoringEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitoringEndpoint

> DeleteMonitoringEndpoint(ctx, monitoringEndpointName).Execute()

Delete a Monitoring Endpoint

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
    monitoringEndpointName := "monitoringEndpointName_example" // string | Name of the Monitoring Endpoint

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MonitoringEndpointAPI.DeleteMonitoringEndpoint(context.Background(), monitoringEndpointName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringEndpointAPI.DeleteMonitoringEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitoringEndpointName** | **string** | Name of the Monitoring Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitoringEndpointRequest struct via the builder pattern


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


## GetMonitoringEndpoint

> StatsdMonitoringEndpointResponse GetMonitoringEndpoint(ctx, monitoringEndpointName).Execute()

Returns a single Monitoring Endpoint

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
    monitoringEndpointName := "monitoringEndpointName_example" // string | Name of the Monitoring Endpoint

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringEndpointAPI.GetMonitoringEndpoint(context.Background(), monitoringEndpointName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringEndpointAPI.GetMonitoringEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitoringEndpoint`: StatsdMonitoringEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringEndpointAPI.GetMonitoringEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitoringEndpointName** | **string** | Name of the Monitoring Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitoringEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatsdMonitoringEndpointResponse**](StatsdMonitoringEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMonitoringEndpoints

> MonitoringEndpointListResponse ListMonitoringEndpoints(ctx).Filter(filter).Execute()

Returns a list of all Monitoring Endpoint objects

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
    resp, r, err := apiClient.MonitoringEndpointAPI.ListMonitoringEndpoints(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringEndpointAPI.ListMonitoringEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMonitoringEndpoints`: MonitoringEndpointListResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringEndpointAPI.ListMonitoringEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMonitoringEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**MonitoringEndpointListResponse**](MonitoringEndpointListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitoringEndpoint

> StatsdMonitoringEndpointResponse UpdateMonitoringEndpoint(ctx, monitoringEndpointName).UpdateRequest(updateRequest).Execute()

Update an existing Monitoring Endpoint by name

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
    monitoringEndpointName := "monitoringEndpointName_example" // string | Name of the Monitoring Endpoint
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Monitoring Endpoint

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringEndpointAPI.UpdateMonitoringEndpoint(context.Background(), monitoringEndpointName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringEndpointAPI.UpdateMonitoringEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitoringEndpoint`: StatsdMonitoringEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringEndpointAPI.UpdateMonitoringEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitoringEndpointName** | **string** | Name of the Monitoring Endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitoringEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Monitoring Endpoint | 

### Return type

[**StatsdMonitoringEndpointResponse**](StatsdMonitoringEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

