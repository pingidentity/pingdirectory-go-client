# \GaugeDataSourceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGaugeDataSource**](GaugeDataSourceApi.md#AddGaugeDataSource) | **Post** /gauge-data-sources | Add a new Gauge Data Source to the config
[**DeleteGaugeDataSource**](GaugeDataSourceApi.md#DeleteGaugeDataSource) | **Delete** /gauge-data-sources/{gauge-data-source-name} | Delete a Gauge Data Source
[**GetGaugeDataSource**](GaugeDataSourceApi.md#GetGaugeDataSource) | **Get** /gauge-data-sources/{gauge-data-source-name} | Returns a single Gauge Data Source
[**UpdateGaugeDataSource**](GaugeDataSourceApi.md#UpdateGaugeDataSource) | **Patch** /gauge-data-sources/{gauge-data-source-name} | Update an existing Gauge Data Source by name



## AddGaugeDataSource

> AddGaugeDataSource200Response AddGaugeDataSource(ctx).AddGaugeDataSourceRequest(addGaugeDataSourceRequest).Execute()

Add a new Gauge Data Source to the config

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
    addGaugeDataSourceRequest := openapiclient.add_gauge_data_source_request{AddIndicatorGaugeDataSourceRequest: openapiclient.NewAddIndicatorGaugeDataSourceRequest("SourceName_example", []openapiclient.EnumindicatorGaugeDataSourceSchemaUrn{openapiclient.Enumindicator-gauge-data-sourceSchemaUrn("urn:pingidentity:schemas:configuration:2.0:gauge-data-source:indicator")}, "MonitorObjectclass_example", "MonitorAttribute_example")} // AddGaugeDataSourceRequest | Create a new Gauge Data Source in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GaugeDataSourceApi.AddGaugeDataSource(context.Background()).AddGaugeDataSourceRequest(addGaugeDataSourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaugeDataSourceApi.AddGaugeDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGaugeDataSource`: AddGaugeDataSource200Response
    fmt.Fprintf(os.Stdout, "Response from `GaugeDataSourceApi.AddGaugeDataSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGaugeDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addGaugeDataSourceRequest** | [**AddGaugeDataSourceRequest**](AddGaugeDataSourceRequest.md) | Create a new Gauge Data Source in the config | 

### Return type

[**AddGaugeDataSource200Response**](AddGaugeDataSource200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGaugeDataSource

> DeleteGaugeDataSource(ctx, gaugeDataSourceName).Execute()

Delete a Gauge Data Source

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
    gaugeDataSourceName := "gaugeDataSourceName_example" // string | Name of the Gauge Data Source to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GaugeDataSourceApi.DeleteGaugeDataSource(context.Background(), gaugeDataSourceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaugeDataSourceApi.DeleteGaugeDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gaugeDataSourceName** | **string** | Name of the Gauge Data Source to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGaugeDataSourceRequest struct via the builder pattern


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


## GetGaugeDataSource

> AddGaugeDataSource200Response GetGaugeDataSource(ctx, gaugeDataSourceName).Execute()

Returns a single Gauge Data Source

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
    gaugeDataSourceName := "gaugeDataSourceName_example" // string | Name of the Gauge Data Source to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GaugeDataSourceApi.GetGaugeDataSource(context.Background(), gaugeDataSourceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaugeDataSourceApi.GetGaugeDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGaugeDataSource`: AddGaugeDataSource200Response
    fmt.Fprintf(os.Stdout, "Response from `GaugeDataSourceApi.GetGaugeDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gaugeDataSourceName** | **string** | Name of the Gauge Data Source to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGaugeDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddGaugeDataSource200Response**](AddGaugeDataSource200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGaugeDataSource

> AddGaugeDataSource200Response UpdateGaugeDataSource(ctx, gaugeDataSourceName).UpdateRequest(updateRequest).Execute()

Update an existing Gauge Data Source by name

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
    gaugeDataSourceName := "gaugeDataSourceName_example" // string | Name of the Gauge Data Source to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Gauge Data Source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GaugeDataSourceApi.UpdateGaugeDataSource(context.Background(), gaugeDataSourceName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaugeDataSourceApi.UpdateGaugeDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGaugeDataSource`: AddGaugeDataSource200Response
    fmt.Fprintf(os.Stdout, "Response from `GaugeDataSourceApi.UpdateGaugeDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gaugeDataSourceName** | **string** | Name of the Gauge Data Source to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGaugeDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Gauge Data Source | 

### Return type

[**AddGaugeDataSource200Response**](AddGaugeDataSource200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

