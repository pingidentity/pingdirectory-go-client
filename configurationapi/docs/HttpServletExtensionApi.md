# \HttpServletExtensionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHttpServletExtension**](HttpServletExtensionApi.md#AddHttpServletExtension) | **Post** /http-servlet-extensions | Add a new HTTP Servlet Extension to the config
[**DeleteHttpServletExtension**](HttpServletExtensionApi.md#DeleteHttpServletExtension) | **Delete** /http-servlet-extensions/{http-servlet-extension-name} | Delete a HTTP Servlet Extension
[**GetHttpServletExtension**](HttpServletExtensionApi.md#GetHttpServletExtension) | **Get** /http-servlet-extensions/{http-servlet-extension-name} | Returns a single HTTP Servlet Extension
[**UpdateHttpServletExtension**](HttpServletExtensionApi.md#UpdateHttpServletExtension) | **Patch** /http-servlet-extensions/{http-servlet-extension-name} | Update an existing HTTP Servlet Extension by name



## AddHttpServletExtension

> AddHttpServletExtension200Response AddHttpServletExtension(ctx).AddHttpServletExtensionRequest(addHttpServletExtensionRequest).Execute()

Add a new HTTP Servlet Extension to the config

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
    addHttpServletExtensionRequest := openapiclient.add_http_servlet_extension_request{AddAvailabilityStateHttpServletExtensionRequest: openapiclient.NewAddAvailabilityStateHttpServletExtensionRequest("ExtensionName_example", []openapiclient.EnumavailabilityStateHttpServletExtensionSchemaUrn{openapiclient.Enumavailability-state-http-servlet-extensionSchemaUrn("urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:availability-state")}, "BaseContextPath_example", int64(123), int64(123), int64(123))} // AddHttpServletExtensionRequest | Create a new HTTP Servlet Extension in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpServletExtensionApi.AddHttpServletExtension(context.Background()).AddHttpServletExtensionRequest(addHttpServletExtensionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServletExtensionApi.AddHttpServletExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddHttpServletExtension`: AddHttpServletExtension200Response
    fmt.Fprintf(os.Stdout, "Response from `HttpServletExtensionApi.AddHttpServletExtension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHttpServletExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addHttpServletExtensionRequest** | [**AddHttpServletExtensionRequest**](AddHttpServletExtensionRequest.md) | Create a new HTTP Servlet Extension in the config | 

### Return type

[**AddHttpServletExtension200Response**](AddHttpServletExtension200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHttpServletExtension

> DeleteHttpServletExtension(ctx, httpServletExtensionName).Execute()

Delete a HTTP Servlet Extension

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HttpServletExtensionApi.DeleteHttpServletExtension(context.Background(), httpServletExtensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServletExtensionApi.DeleteHttpServletExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHttpServletExtensionRequest struct via the builder pattern


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


## GetHttpServletExtension

> GetHttpServletExtension200Response GetHttpServletExtension(ctx, httpServletExtensionName).Execute()

Returns a single HTTP Servlet Extension

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpServletExtensionApi.GetHttpServletExtension(context.Background(), httpServletExtensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServletExtensionApi.GetHttpServletExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHttpServletExtension`: GetHttpServletExtension200Response
    fmt.Fprintf(os.Stdout, "Response from `HttpServletExtensionApi.GetHttpServletExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHttpServletExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHttpServletExtension200Response**](GetHttpServletExtension200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHttpServletExtension

> GetHttpServletExtension200Response UpdateHttpServletExtension(ctx, httpServletExtensionName).UpdateRequest(updateRequest).Execute()

Update an existing HTTP Servlet Extension by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing HTTP Servlet Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpServletExtensionApi.UpdateHttpServletExtension(context.Background(), httpServletExtensionName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServletExtensionApi.UpdateHttpServletExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHttpServletExtension`: GetHttpServletExtension200Response
    fmt.Fprintf(os.Stdout, "Response from `HttpServletExtensionApi.UpdateHttpServletExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHttpServletExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing HTTP Servlet Extension | 

### Return type

[**GetHttpServletExtension200Response**](GetHttpServletExtension200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

