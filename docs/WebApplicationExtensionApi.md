# \WebApplicationExtensionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebApplicationExtension**](WebApplicationExtensionApi.md#GetWebApplicationExtension) | **Get** /web-application-extensions/{web-application-extension-name} | Returns a single Web Application Extension
[**UpdateWebApplicationExtension**](WebApplicationExtensionApi.md#UpdateWebApplicationExtension) | **Patch** /web-application-extensions/{web-application-extension-name} | Update an existing Web Application Extension by name



## GetWebApplicationExtension

> ConsoleWebApplicationExtensionResponse GetWebApplicationExtension(ctx, webApplicationExtensionName).Execute()

Returns a single Web Application Extension

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
    webApplicationExtensionName := "webApplicationExtensionName_example" // string | Name of the Web Application Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebApplicationExtensionApi.GetWebApplicationExtension(context.Background(), webApplicationExtensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebApplicationExtensionApi.GetWebApplicationExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebApplicationExtension`: ConsoleWebApplicationExtensionResponse
    fmt.Fprintf(os.Stdout, "Response from `WebApplicationExtensionApi.GetWebApplicationExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webApplicationExtensionName** | **string** | Name of the Web Application Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebApplicationExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleWebApplicationExtensionResponse**](ConsoleWebApplicationExtensionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebApplicationExtension

> ConsoleWebApplicationExtensionResponse UpdateWebApplicationExtension(ctx, webApplicationExtensionName).UpdateRequest(updateRequest).Execute()

Update an existing Web Application Extension by name

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
    webApplicationExtensionName := "webApplicationExtensionName_example" // string | Name of the Web Application Extension
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Web Application Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebApplicationExtensionApi.UpdateWebApplicationExtension(context.Background(), webApplicationExtensionName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebApplicationExtensionApi.UpdateWebApplicationExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebApplicationExtension`: ConsoleWebApplicationExtensionResponse
    fmt.Fprintf(os.Stdout, "Response from `WebApplicationExtensionApi.UpdateWebApplicationExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webApplicationExtensionName** | **string** | Name of the Web Application Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebApplicationExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Web Application Extension | 

### Return type

[**ConsoleWebApplicationExtensionResponse**](ConsoleWebApplicationExtensionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

