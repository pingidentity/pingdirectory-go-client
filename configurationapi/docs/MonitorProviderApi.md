# \MonitorProviderAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMonitorProvider**](MonitorProviderAPI.md#AddMonitorProvider) | **Post** /monitor-providers | Add a new Monitor Provider to the config
[**DeleteMonitorProvider**](MonitorProviderAPI.md#DeleteMonitorProvider) | **Delete** /monitor-providers/{monitor-provider-name} | Delete a Monitor Provider
[**GetMonitorProvider**](MonitorProviderAPI.md#GetMonitorProvider) | **Get** /monitor-providers/{monitor-provider-name} | Returns a single Monitor Provider
[**ListMonitorProviders**](MonitorProviderAPI.md#ListMonitorProviders) | **Get** /monitor-providers | Returns a list of all Monitor Provider objects
[**UpdateMonitorProvider**](MonitorProviderAPI.md#UpdateMonitorProvider) | **Patch** /monitor-providers/{monitor-provider-name} | Update an existing Monitor Provider by name



## AddMonitorProvider

> AddMonitorProvider200Response AddMonitorProvider(ctx).AddMonitorProviderRequest(addMonitorProviderRequest).Execute()

Add a new Monitor Provider to the config

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
    addMonitorProviderRequest := openapiclient.add_monitor_provider_request{AddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest: openapiclient.NewAddEncryptionSettingsDatabaseAccessibilityMonitorProviderRequest([]openapiclient.EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn{openapiclient.Enumencryption-settings-database-accessibility-monitor-providerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:monitor-provider:encryption-settings-database-accessibility")}, false, "ProviderName_example")} // AddMonitorProviderRequest | Create a new Monitor Provider in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorProviderAPI.AddMonitorProvider(context.Background()).AddMonitorProviderRequest(addMonitorProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorProviderAPI.AddMonitorProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMonitorProvider`: AddMonitorProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `MonitorProviderAPI.AddMonitorProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMonitorProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addMonitorProviderRequest** | [**AddMonitorProviderRequest**](AddMonitorProviderRequest.md) | Create a new Monitor Provider in the config | 

### Return type

[**AddMonitorProvider200Response**](AddMonitorProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMonitorProvider

> DeleteMonitorProvider(ctx, monitorProviderName).Execute()

Delete a Monitor Provider

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
    monitorProviderName := "monitorProviderName_example" // string | Name of the Monitor Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MonitorProviderAPI.DeleteMonitorProvider(context.Background(), monitorProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorProviderAPI.DeleteMonitorProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorProviderName** | **string** | Name of the Monitor Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMonitorProviderRequest struct via the builder pattern


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


## GetMonitorProvider

> GetMonitorProvider200Response GetMonitorProvider(ctx, monitorProviderName).Execute()

Returns a single Monitor Provider

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
    monitorProviderName := "monitorProviderName_example" // string | Name of the Monitor Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorProviderAPI.GetMonitorProvider(context.Background(), monitorProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorProviderAPI.GetMonitorProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorProvider`: GetMonitorProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `MonitorProviderAPI.GetMonitorProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorProviderName** | **string** | Name of the Monitor Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMonitorProvider200Response**](GetMonitorProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMonitorProviders

> MonitorProviderListResponse ListMonitorProviders(ctx).Filter(filter).Execute()

Returns a list of all Monitor Provider objects

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
    resp, r, err := apiClient.MonitorProviderAPI.ListMonitorProviders(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorProviderAPI.ListMonitorProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMonitorProviders`: MonitorProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorProviderAPI.ListMonitorProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMonitorProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**MonitorProviderListResponse**](MonitorProviderListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitorProvider

> GetMonitorProvider200Response UpdateMonitorProvider(ctx, monitorProviderName).UpdateRequest(updateRequest).Execute()

Update an existing Monitor Provider by name

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
    monitorProviderName := "monitorProviderName_example" // string | Name of the Monitor Provider
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Monitor Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorProviderAPI.UpdateMonitorProvider(context.Background(), monitorProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorProviderAPI.UpdateMonitorProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitorProvider`: GetMonitorProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `MonitorProviderAPI.UpdateMonitorProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorProviderName** | **string** | Name of the Monitor Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitorProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Monitor Provider | 

### Return type

[**GetMonitorProvider200Response**](GetMonitorProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

