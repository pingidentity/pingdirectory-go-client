# \VelocityTemplateLoaderAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVelocityTemplateLoader**](VelocityTemplateLoaderAPI.md#AddVelocityTemplateLoader) | **Post** /http-servlet-extensions/{http-servlet-extension-name}/velocity-template-loaders | Add a new Velocity Template Loader to the config
[**DeleteVelocityTemplateLoader**](VelocityTemplateLoaderAPI.md#DeleteVelocityTemplateLoader) | **Delete** /http-servlet-extensions/{http-servlet-extension-name}/velocity-template-loaders/{velocity-template-loader-name} | Delete a Velocity Template Loader
[**GetVelocityTemplateLoader**](VelocityTemplateLoaderAPI.md#GetVelocityTemplateLoader) | **Get** /http-servlet-extensions/{http-servlet-extension-name}/velocity-template-loaders/{velocity-template-loader-name} | Returns a single Velocity Template Loader
[**ListVelocityTemplateLoaders**](VelocityTemplateLoaderAPI.md#ListVelocityTemplateLoaders) | **Get** /http-servlet-extensions/{http-servlet-extension-name}/velocity-template-loaders | Returns a list of all Velocity Template Loader objects
[**UpdateVelocityTemplateLoader**](VelocityTemplateLoaderAPI.md#UpdateVelocityTemplateLoader) | **Patch** /http-servlet-extensions/{http-servlet-extension-name}/velocity-template-loaders/{velocity-template-loader-name} | Update an existing Velocity Template Loader by name



## AddVelocityTemplateLoader

> VelocityTemplateLoaderResponse AddVelocityTemplateLoader(ctx, httpServletExtensionName).AddVelocityTemplateLoaderRequest(addVelocityTemplateLoaderRequest).Execute()

Add a new Velocity Template Loader to the config

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
    addVelocityTemplateLoaderRequest := *openapiclient.NewAddVelocityTemplateLoaderRequest("MimeTypeMatcher_example", "LoaderName_example") // AddVelocityTemplateLoaderRequest | Create a new Velocity Template Loader in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityTemplateLoaderAPI.AddVelocityTemplateLoader(context.Background(), httpServletExtensionName).AddVelocityTemplateLoaderRequest(addVelocityTemplateLoaderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityTemplateLoaderAPI.AddVelocityTemplateLoader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVelocityTemplateLoader`: VelocityTemplateLoaderResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityTemplateLoaderAPI.AddVelocityTemplateLoader`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVelocityTemplateLoaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addVelocityTemplateLoaderRequest** | [**AddVelocityTemplateLoaderRequest**](AddVelocityTemplateLoaderRequest.md) | Create a new Velocity Template Loader in the config | 

### Return type

[**VelocityTemplateLoaderResponse**](VelocityTemplateLoaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVelocityTemplateLoader

> DeleteVelocityTemplateLoader(ctx, velocityTemplateLoaderName, httpServletExtensionName).Execute()

Delete a Velocity Template Loader

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
    velocityTemplateLoaderName := "velocityTemplateLoaderName_example" // string | Name of the Velocity Template Loader
    httpServletExtensionName := "httpServletExtensionName_example" // string | Name of the HTTP Servlet Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VelocityTemplateLoaderAPI.DeleteVelocityTemplateLoader(context.Background(), velocityTemplateLoaderName, httpServletExtensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityTemplateLoaderAPI.DeleteVelocityTemplateLoader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityTemplateLoaderName** | **string** | Name of the Velocity Template Loader | 
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVelocityTemplateLoaderRequest struct via the builder pattern


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


## GetVelocityTemplateLoader

> VelocityTemplateLoaderResponse GetVelocityTemplateLoader(ctx, velocityTemplateLoaderName, httpServletExtensionName).Execute()

Returns a single Velocity Template Loader

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
    velocityTemplateLoaderName := "velocityTemplateLoaderName_example" // string | Name of the Velocity Template Loader
    httpServletExtensionName := "httpServletExtensionName_example" // string | Name of the HTTP Servlet Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityTemplateLoaderAPI.GetVelocityTemplateLoader(context.Background(), velocityTemplateLoaderName, httpServletExtensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityTemplateLoaderAPI.GetVelocityTemplateLoader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVelocityTemplateLoader`: VelocityTemplateLoaderResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityTemplateLoaderAPI.GetVelocityTemplateLoader`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityTemplateLoaderName** | **string** | Name of the Velocity Template Loader | 
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVelocityTemplateLoaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VelocityTemplateLoaderResponse**](VelocityTemplateLoaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVelocityTemplateLoaders

> VelocityTemplateLoaderListResponse ListVelocityTemplateLoaders(ctx, httpServletExtensionName).Filter(filter).Execute()

Returns a list of all Velocity Template Loader objects

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
    resp, r, err := apiClient.VelocityTemplateLoaderAPI.ListVelocityTemplateLoaders(context.Background(), httpServletExtensionName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityTemplateLoaderAPI.ListVelocityTemplateLoaders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVelocityTemplateLoaders`: VelocityTemplateLoaderListResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityTemplateLoaderAPI.ListVelocityTemplateLoaders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVelocityTemplateLoadersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**VelocityTemplateLoaderListResponse**](VelocityTemplateLoaderListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVelocityTemplateLoader

> VelocityTemplateLoaderResponse UpdateVelocityTemplateLoader(ctx, velocityTemplateLoaderName, httpServletExtensionName).UpdateRequest(updateRequest).Execute()

Update an existing Velocity Template Loader by name

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
    velocityTemplateLoaderName := "velocityTemplateLoaderName_example" // string | Name of the Velocity Template Loader
    httpServletExtensionName := "httpServletExtensionName_example" // string | Name of the HTTP Servlet Extension
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Velocity Template Loader

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityTemplateLoaderAPI.UpdateVelocityTemplateLoader(context.Background(), velocityTemplateLoaderName, httpServletExtensionName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityTemplateLoaderAPI.UpdateVelocityTemplateLoader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVelocityTemplateLoader`: VelocityTemplateLoaderResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityTemplateLoaderAPI.UpdateVelocityTemplateLoader`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityTemplateLoaderName** | **string** | Name of the Velocity Template Loader | 
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVelocityTemplateLoaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Velocity Template Loader | 

### Return type

[**VelocityTemplateLoaderResponse**](VelocityTemplateLoaderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

