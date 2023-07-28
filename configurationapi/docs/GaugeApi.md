# \GaugeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGauge**](GaugeApi.md#AddGauge) | **Post** /gauges | Add a new Gauge to the config
[**DeleteGauge**](GaugeApi.md#DeleteGauge) | **Delete** /gauges/{gauge-name} | Delete a Gauge
[**GetGauge**](GaugeApi.md#GetGauge) | **Get** /gauges/{gauge-name} | Returns a single Gauge
[**ListGauges**](GaugeApi.md#ListGauges) | **Get** /gauges | Returns a list of all Gauge objects
[**UpdateGauge**](GaugeApi.md#UpdateGauge) | **Patch** /gauges/{gauge-name} | Update an existing Gauge by name



## AddGauge

> AddGauge200Response AddGauge(ctx).AddGaugeRequest(addGaugeRequest).Execute()

Add a new Gauge to the config

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
    addGaugeRequest := openapiclient.add_gauge_request{AddIndicatorGaugeRequest: openapiclient.NewAddIndicatorGaugeRequest("GaugeName_example", []openapiclient.EnumindicatorGaugeSchemaUrn{openapiclient.Enumindicator-gaugeSchemaUrn("urn:pingidentity:schemas:configuration:2.0:gauge:indicator")}, "GaugeDataSource_example")} // AddGaugeRequest | Create a new Gauge in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GaugeApi.AddGauge(context.Background()).AddGaugeRequest(addGaugeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaugeApi.AddGauge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGauge`: AddGauge200Response
    fmt.Fprintf(os.Stdout, "Response from `GaugeApi.AddGauge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGaugeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addGaugeRequest** | [**AddGaugeRequest**](AddGaugeRequest.md) | Create a new Gauge in the config | 

### Return type

[**AddGauge200Response**](AddGauge200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGauge

> DeleteGauge(ctx, gaugeName).Execute()

Delete a Gauge

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
    gaugeName := "gaugeName_example" // string | Name of the Gauge

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GaugeApi.DeleteGauge(context.Background(), gaugeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaugeApi.DeleteGauge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gaugeName** | **string** | Name of the Gauge | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGaugeRequest struct via the builder pattern


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


## GetGauge

> AddGauge200Response GetGauge(ctx, gaugeName).Execute()

Returns a single Gauge

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
    gaugeName := "gaugeName_example" // string | Name of the Gauge

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GaugeApi.GetGauge(context.Background(), gaugeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaugeApi.GetGauge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGauge`: AddGauge200Response
    fmt.Fprintf(os.Stdout, "Response from `GaugeApi.GetGauge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gaugeName** | **string** | Name of the Gauge | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGaugeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddGauge200Response**](AddGauge200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGauges

> GaugeListResponse ListGauges(ctx).Filter(filter).Execute()

Returns a list of all Gauge objects

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
    resp, r, err := apiClient.GaugeApi.ListGauges(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaugeApi.ListGauges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGauges`: GaugeListResponse
    fmt.Fprintf(os.Stdout, "Response from `GaugeApi.ListGauges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGaugesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**GaugeListResponse**](GaugeListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGauge

> AddGauge200Response UpdateGauge(ctx, gaugeName).UpdateRequest(updateRequest).Execute()

Update an existing Gauge by name

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
    gaugeName := "gaugeName_example" // string | Name of the Gauge
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Gauge

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GaugeApi.UpdateGauge(context.Background(), gaugeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaugeApi.UpdateGauge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGauge`: AddGauge200Response
    fmt.Fprintf(os.Stdout, "Response from `GaugeApi.UpdateGauge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gaugeName** | **string** | Name of the Gauge | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGaugeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Gauge | 

### Return type

[**AddGauge200Response**](AddGauge200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

