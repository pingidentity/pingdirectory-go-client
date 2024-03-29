# \PluginAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPlugin**](PluginAPI.md#AddPlugin) | **Post** /plugin-root/plugins | Add a new Plugin to the config
[**DeletePlugin**](PluginAPI.md#DeletePlugin) | **Delete** /plugin-root/plugins/{plugin-name} | Delete a Plugin
[**GetPlugin**](PluginAPI.md#GetPlugin) | **Get** /plugin-root/plugins/{plugin-name} | Returns a single Plugin
[**ListPlugins**](PluginAPI.md#ListPlugins) | **Get** /plugin-root/plugins | Returns a list of all Plugin objects
[**UpdatePlugin**](PluginAPI.md#UpdatePlugin) | **Patch** /plugin-root/plugins/{plugin-name} | Update an existing Plugin by name



## AddPlugin

> AddPlugin200Response AddPlugin(ctx).AddPluginRequest(addPluginRequest).Execute()

Add a new Plugin to the config

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
    addPluginRequest := openapiclient.add_plugin_request{AddAttributeMapperPluginRequest: openapiclient.NewAddAttributeMapperPluginRequest([]openapiclient.EnumattributeMapperPluginSchemaUrn{openapiclient.Enumattribute-mapper-pluginSchemaUrn("urn:pingidentity:schemas:configuration:2.0:plugin:attribute-mapper")}, "SourceAttribute_example", "TargetAttribute_example", false, "PluginName_example")} // AddPluginRequest | Create a new Plugin in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PluginAPI.AddPlugin(context.Background()).AddPluginRequest(addPluginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.AddPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPlugin`: AddPlugin200Response
    fmt.Fprintf(os.Stdout, "Response from `PluginAPI.AddPlugin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPluginRequest** | [**AddPluginRequest**](AddPluginRequest.md) | Create a new Plugin in the config | 

### Return type

[**AddPlugin200Response**](AddPlugin200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlugin

> DeletePlugin(ctx, pluginName).Execute()

Delete a Plugin

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
    pluginName := "pluginName_example" // string | Name of the Plugin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PluginAPI.DeletePlugin(context.Background(), pluginName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.DeletePlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePluginRequest struct via the builder pattern


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


## GetPlugin

> GetPlugin200Response GetPlugin(ctx, pluginName).Execute()

Returns a single Plugin

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
    pluginName := "pluginName_example" // string | Name of the Plugin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PluginAPI.GetPlugin(context.Background(), pluginName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlugin`: GetPlugin200Response
    fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPlugin200Response**](GetPlugin200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlugins

> PluginListResponse ListPlugins(ctx).Filter(filter).Execute()

Returns a list of all Plugin objects

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
    resp, r, err := apiClient.PluginAPI.ListPlugins(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.ListPlugins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlugins`: PluginListResponse
    fmt.Fprintf(os.Stdout, "Response from `PluginAPI.ListPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**PluginListResponse**](PluginListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlugin

> GetPlugin200Response UpdatePlugin(ctx, pluginName).UpdateRequest(updateRequest).Execute()

Update an existing Plugin by name

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
    pluginName := "pluginName_example" // string | Name of the Plugin
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Plugin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PluginAPI.UpdatePlugin(context.Background(), pluginName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.UpdatePlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlugin`: GetPlugin200Response
    fmt.Fprintf(os.Stdout, "Response from `PluginAPI.UpdatePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Plugin | 

### Return type

[**GetPlugin200Response**](GetPlugin200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

