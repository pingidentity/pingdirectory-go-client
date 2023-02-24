# \WebApplicationExtensionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWebApplicationExtension**](WebApplicationExtensionApi.md#AddWebApplicationExtension) | **Post** /web-application-extensions | Add a new Web Application Extension to the config
[**DeleteWebApplicationExtension**](WebApplicationExtensionApi.md#DeleteWebApplicationExtension) | **Delete** /web-application-extensions/{web-application-extension-name} | Delete a Web Application Extension
[**GetWebApplicationExtension**](WebApplicationExtensionApi.md#GetWebApplicationExtension) | **Get** /web-application-extensions/{web-application-extension-name} | Returns a single Web Application Extension
[**UpdateWebApplicationExtension**](WebApplicationExtensionApi.md#UpdateWebApplicationExtension) | **Patch** /web-application-extensions/{web-application-extension-name} | Update an existing Web Application Extension by name



## AddWebApplicationExtension

> AddWebApplicationExtension200Response AddWebApplicationExtension(ctx).AddGenericWebApplicationExtensionRequest(addGenericWebApplicationExtensionRequest).Execute()

Add a new Web Application Extension to the config

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
    addGenericWebApplicationExtensionRequest := *openapiclient.NewAddGenericWebApplicationExtensionRequest("ExtensionName_example", []openapiclient.EnumgenericWebApplicationExtensionSchemaUrn{openapiclient.Enumgeneric-web-application-extensionSchemaUrn("urn:pingidentity:schemas:configuration:2.0:web-application-extension:generic")}, "BaseContextPath_example") // AddGenericWebApplicationExtensionRequest | Create a new Web Application Extension in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebApplicationExtensionApi.AddWebApplicationExtension(context.Background()).AddGenericWebApplicationExtensionRequest(addGenericWebApplicationExtensionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebApplicationExtensionApi.AddWebApplicationExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWebApplicationExtension`: AddWebApplicationExtension200Response
    fmt.Fprintf(os.Stdout, "Response from `WebApplicationExtensionApi.AddWebApplicationExtension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddWebApplicationExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addGenericWebApplicationExtensionRequest** | [**AddGenericWebApplicationExtensionRequest**](AddGenericWebApplicationExtensionRequest.md) | Create a new Web Application Extension in the config | 

### Return type

[**AddWebApplicationExtension200Response**](AddWebApplicationExtension200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebApplicationExtension

> DeleteWebApplicationExtension(ctx, webApplicationExtensionName).Execute()

Delete a Web Application Extension

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
    webApplicationExtensionName := "webApplicationExtensionName_example" // string | Name of the Web Application Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebApplicationExtensionApi.DeleteWebApplicationExtension(context.Background(), webApplicationExtensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebApplicationExtensionApi.DeleteWebApplicationExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webApplicationExtensionName** | **string** | Name of the Web Application Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebApplicationExtensionRequest struct via the builder pattern


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


## GetWebApplicationExtension

> GetWebApplicationExtension200Response GetWebApplicationExtension(ctx, webApplicationExtensionName).Execute()

Returns a single Web Application Extension

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
    webApplicationExtensionName := "webApplicationExtensionName_example" // string | Name of the Web Application Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebApplicationExtensionApi.GetWebApplicationExtension(context.Background(), webApplicationExtensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebApplicationExtensionApi.GetWebApplicationExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebApplicationExtension`: GetWebApplicationExtension200Response
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

[**GetWebApplicationExtension200Response**](GetWebApplicationExtension200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebApplicationExtension

> GetWebApplicationExtension200Response UpdateWebApplicationExtension(ctx, webApplicationExtensionName).UpdateRequest(updateRequest).Execute()

Update an existing Web Application Extension by name

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
    webApplicationExtensionName := "webApplicationExtensionName_example" // string | Name of the Web Application Extension
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Web Application Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebApplicationExtensionApi.UpdateWebApplicationExtension(context.Background(), webApplicationExtensionName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebApplicationExtensionApi.UpdateWebApplicationExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebApplicationExtension`: GetWebApplicationExtension200Response
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

[**GetWebApplicationExtension200Response**](GetWebApplicationExtension200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

