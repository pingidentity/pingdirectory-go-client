# \PluginRootApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPluginRoot**](PluginRootApi.md#GetPluginRoot) | **Get** /plugin-root | Returns a single Plugin Root
[**UpdatePluginRoot**](PluginRootApi.md#UpdatePluginRoot) | **Patch** /plugin-root | Update an existing Plugin Root by name



## GetPluginRoot

> PluginRootResponse GetPluginRoot(ctx).Execute()

Returns a single Plugin Root

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PluginRootApi.GetPluginRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginRootApi.GetPluginRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginRoot`: PluginRootResponse
    fmt.Fprintf(os.Stdout, "Response from `PluginRootApi.GetPluginRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginRootRequest struct via the builder pattern


### Return type

[**PluginRootResponse**](PluginRootResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePluginRoot

> PluginRootResponse UpdatePluginRoot(ctx).UpdateRequest(updateRequest).Execute()

Update an existing Plugin Root by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Plugin Root

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PluginRootApi.UpdatePluginRoot(context.Background()).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PluginRootApi.UpdatePluginRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePluginRoot`: PluginRootResponse
    fmt.Fprintf(os.Stdout, "Response from `PluginRootApi.UpdatePluginRoot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePluginRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Plugin Root | 

### Return type

[**PluginRootResponse**](PluginRootResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

