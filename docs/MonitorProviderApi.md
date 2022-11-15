# \MonitorProviderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMonitorProvider**](MonitorProviderApi.md#AddMonitorProvider) | **Post** /monitor-providers | Add a new Monitor Provider to the config
[**DeleteMonitorProvider**](MonitorProviderApi.md#DeleteMonitorProvider) | **Delete** /monitor-providers/{monitor-provider-name} | Delete a Monitor Provider
[**GetMonitorProvider**](MonitorProviderApi.md#GetMonitorProvider) | **Get** /monitor-providers/{monitor-provider-name} | Returns a single Monitor Provider
[**UpdateMonitorProvider**](MonitorProviderApi.md#UpdateMonitorProvider) | **Patch** /monitor-providers/{monitor-provider-name} | Update an existing Monitor Provider by name



## AddMonitorProvider

> AddMonitorProvider200Response AddMonitorProvider(ctx).AddThirdPartyMonitorProviderRequest(addThirdPartyMonitorProviderRequest).Execute()

Add a new Monitor Provider to the config

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
    addThirdPartyMonitorProviderRequest := *openapiclient.NewAddThirdPartyMonitorProviderRequest("ProviderName_example", []openapiclient.EnumthirdPartyMonitorProviderSchemaUrn{openapiclient.Enumthird-party-monitor-providerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:monitor-provider:third-party")}, "ExtensionClass_example", false) // AddThirdPartyMonitorProviderRequest | Create a new Monitor Provider in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorProviderApi.AddMonitorProvider(context.Background()).AddThirdPartyMonitorProviderRequest(addThirdPartyMonitorProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorProviderApi.AddMonitorProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMonitorProvider`: AddMonitorProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `MonitorProviderApi.AddMonitorProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMonitorProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addThirdPartyMonitorProviderRequest** | [**AddThirdPartyMonitorProviderRequest**](AddThirdPartyMonitorProviderRequest.md) | Create a new Monitor Provider in the config | 

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
    openapiclient "./openapi"
)

func main() {
    monitorProviderName := "monitorProviderName_example" // string | Name of the Monitor Provider to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorProviderApi.DeleteMonitorProvider(context.Background(), monitorProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorProviderApi.DeleteMonitorProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorProviderName** | **string** | Name of the Monitor Provider to be deleted | 

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
    openapiclient "./openapi"
)

func main() {
    monitorProviderName := "monitorProviderName_example" // string | Name of the Monitor Provider to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorProviderApi.GetMonitorProvider(context.Background(), monitorProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorProviderApi.GetMonitorProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorProvider`: GetMonitorProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `MonitorProviderApi.GetMonitorProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorProviderName** | **string** | Name of the Monitor Provider to be read | 

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
    openapiclient "./openapi"
)

func main() {
    monitorProviderName := "monitorProviderName_example" // string | Name of the Monitor Provider to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Monitor Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorProviderApi.UpdateMonitorProvider(context.Background(), monitorProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorProviderApi.UpdateMonitorProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitorProvider`: GetMonitorProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `MonitorProviderApi.UpdateMonitorProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorProviderName** | **string** | Name of the Monitor Provider to be updated | 

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

