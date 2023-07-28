# \PrometheusMonitorAttributeMetricApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrometheusMonitorAttributeMetric**](PrometheusMonitorAttributeMetricApi.md#AddPrometheusMonitorAttributeMetric) | **Post** /http-servlet-extensions/{http-servlet-extension-name}/prometheus-monitor-attribute-metrics | Add a new Prometheus Monitor Attribute Metric to the config
[**DeletePrometheusMonitorAttributeMetric**](PrometheusMonitorAttributeMetricApi.md#DeletePrometheusMonitorAttributeMetric) | **Delete** /http-servlet-extensions/{http-servlet-extension-name}/prometheus-monitor-attribute-metrics/{prometheus-monitor-attribute-metric-name} | Delete a Prometheus Monitor Attribute Metric
[**GetPrometheusMonitorAttributeMetric**](PrometheusMonitorAttributeMetricApi.md#GetPrometheusMonitorAttributeMetric) | **Get** /http-servlet-extensions/{http-servlet-extension-name}/prometheus-monitor-attribute-metrics/{prometheus-monitor-attribute-metric-name} | Returns a single Prometheus Monitor Attribute Metric
[**ListPrometheusMonitorAttributeMetrics**](PrometheusMonitorAttributeMetricApi.md#ListPrometheusMonitorAttributeMetrics) | **Get** /http-servlet-extensions/{http-servlet-extension-name}/prometheus-monitor-attribute-metrics | Returns a list of all Prometheus Monitor Attribute Metric objects
[**UpdatePrometheusMonitorAttributeMetric**](PrometheusMonitorAttributeMetricApi.md#UpdatePrometheusMonitorAttributeMetric) | **Patch** /http-servlet-extensions/{http-servlet-extension-name}/prometheus-monitor-attribute-metrics/{prometheus-monitor-attribute-metric-name} | Update an existing Prometheus Monitor Attribute Metric by name



## AddPrometheusMonitorAttributeMetric

> PrometheusMonitorAttributeMetricResponse AddPrometheusMonitorAttributeMetric(ctx, httpServletExtensionName).AddPrometheusMonitorAttributeMetricRequest(addPrometheusMonitorAttributeMetricRequest).Execute()

Add a new Prometheus Monitor Attribute Metric to the config

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
    httpServletExtensionName := "httpServletExtensionName_example" // string | Name of the HTTP Servlet Extension
    addPrometheusMonitorAttributeMetricRequest := *openapiclient.NewAddPrometheusMonitorAttributeMetricRequest("MetricName_example", "MonitorAttributeName_example", "MonitorObjectClassName_example", openapiclient.Enumprometheus-monitor-attribute-metric-metricTypeProp("counter")) // AddPrometheusMonitorAttributeMetricRequest | Create a new Prometheus Monitor Attribute Metric in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrometheusMonitorAttributeMetricApi.AddPrometheusMonitorAttributeMetric(context.Background(), httpServletExtensionName).AddPrometheusMonitorAttributeMetricRequest(addPrometheusMonitorAttributeMetricRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusMonitorAttributeMetricApi.AddPrometheusMonitorAttributeMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPrometheusMonitorAttributeMetric`: PrometheusMonitorAttributeMetricResponse
    fmt.Fprintf(os.Stdout, "Response from `PrometheusMonitorAttributeMetricApi.AddPrometheusMonitorAttributeMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPrometheusMonitorAttributeMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addPrometheusMonitorAttributeMetricRequest** | [**AddPrometheusMonitorAttributeMetricRequest**](AddPrometheusMonitorAttributeMetricRequest.md) | Create a new Prometheus Monitor Attribute Metric in the config | 

### Return type

[**PrometheusMonitorAttributeMetricResponse**](PrometheusMonitorAttributeMetricResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrometheusMonitorAttributeMetric

> DeletePrometheusMonitorAttributeMetric(ctx, prometheusMonitorAttributeMetricName, httpServletExtensionName).Execute()

Delete a Prometheus Monitor Attribute Metric

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
    prometheusMonitorAttributeMetricName := "prometheusMonitorAttributeMetricName_example" // string | Name of the Prometheus Monitor Attribute Metric
    httpServletExtensionName := "httpServletExtensionName_example" // string | Name of the HTTP Servlet Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PrometheusMonitorAttributeMetricApi.DeletePrometheusMonitorAttributeMetric(context.Background(), prometheusMonitorAttributeMetricName, httpServletExtensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusMonitorAttributeMetricApi.DeletePrometheusMonitorAttributeMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prometheusMonitorAttributeMetricName** | **string** | Name of the Prometheus Monitor Attribute Metric | 
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrometheusMonitorAttributeMetricRequest struct via the builder pattern


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


## GetPrometheusMonitorAttributeMetric

> PrometheusMonitorAttributeMetricResponse GetPrometheusMonitorAttributeMetric(ctx, prometheusMonitorAttributeMetricName, httpServletExtensionName).Execute()

Returns a single Prometheus Monitor Attribute Metric

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
    prometheusMonitorAttributeMetricName := "prometheusMonitorAttributeMetricName_example" // string | Name of the Prometheus Monitor Attribute Metric
    httpServletExtensionName := "httpServletExtensionName_example" // string | Name of the HTTP Servlet Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrometheusMonitorAttributeMetricApi.GetPrometheusMonitorAttributeMetric(context.Background(), prometheusMonitorAttributeMetricName, httpServletExtensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusMonitorAttributeMetricApi.GetPrometheusMonitorAttributeMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrometheusMonitorAttributeMetric`: PrometheusMonitorAttributeMetricResponse
    fmt.Fprintf(os.Stdout, "Response from `PrometheusMonitorAttributeMetricApi.GetPrometheusMonitorAttributeMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prometheusMonitorAttributeMetricName** | **string** | Name of the Prometheus Monitor Attribute Metric | 
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrometheusMonitorAttributeMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PrometheusMonitorAttributeMetricResponse**](PrometheusMonitorAttributeMetricResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrometheusMonitorAttributeMetrics

> PrometheusMonitorAttributeMetricListResponse ListPrometheusMonitorAttributeMetrics(ctx, httpServletExtensionName).Filter(filter).Execute()

Returns a list of all Prometheus Monitor Attribute Metric objects

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
    httpServletExtensionName := "httpServletExtensionName_example" // string | Name of the HTTP Servlet Extension
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrometheusMonitorAttributeMetricApi.ListPrometheusMonitorAttributeMetrics(context.Background(), httpServletExtensionName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusMonitorAttributeMetricApi.ListPrometheusMonitorAttributeMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrometheusMonitorAttributeMetrics`: PrometheusMonitorAttributeMetricListResponse
    fmt.Fprintf(os.Stdout, "Response from `PrometheusMonitorAttributeMetricApi.ListPrometheusMonitorAttributeMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrometheusMonitorAttributeMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**PrometheusMonitorAttributeMetricListResponse**](PrometheusMonitorAttributeMetricListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrometheusMonitorAttributeMetric

> PrometheusMonitorAttributeMetricResponse UpdatePrometheusMonitorAttributeMetric(ctx, prometheusMonitorAttributeMetricName, httpServletExtensionName).UpdateRequest(updateRequest).Execute()

Update an existing Prometheus Monitor Attribute Metric by name

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
    prometheusMonitorAttributeMetricName := "prometheusMonitorAttributeMetricName_example" // string | Name of the Prometheus Monitor Attribute Metric
    httpServletExtensionName := "httpServletExtensionName_example" // string | Name of the HTTP Servlet Extension
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Prometheus Monitor Attribute Metric

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrometheusMonitorAttributeMetricApi.UpdatePrometheusMonitorAttributeMetric(context.Background(), prometheusMonitorAttributeMetricName, httpServletExtensionName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrometheusMonitorAttributeMetricApi.UpdatePrometheusMonitorAttributeMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePrometheusMonitorAttributeMetric`: PrometheusMonitorAttributeMetricResponse
    fmt.Fprintf(os.Stdout, "Response from `PrometheusMonitorAttributeMetricApi.UpdatePrometheusMonitorAttributeMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prometheusMonitorAttributeMetricName** | **string** | Name of the Prometheus Monitor Attribute Metric | 
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrometheusMonitorAttributeMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Prometheus Monitor Attribute Metric | 

### Return type

[**PrometheusMonitorAttributeMetricResponse**](PrometheusMonitorAttributeMetricResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

