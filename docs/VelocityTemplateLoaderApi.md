# \VelocityTemplateLoaderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVelocityTemplateLoader**](VelocityTemplateLoaderApi.md#AddVelocityTemplateLoader) | **Post** /velocity-template-loaders | Add a new Velocity Template Loader to the config
[**DeleteVelocityTemplateLoader**](VelocityTemplateLoaderApi.md#DeleteVelocityTemplateLoader) | **Delete** /velocity-template-loaders/{velocity-template-loader-name} | Delete a Velocity Template Loader
[**GetVelocityTemplateLoader**](VelocityTemplateLoaderApi.md#GetVelocityTemplateLoader) | **Get** /velocity-template-loaders/{velocity-template-loader-name} | Returns a single Velocity Template Loader
[**UpdateVelocityTemplateLoader**](VelocityTemplateLoaderApi.md#UpdateVelocityTemplateLoader) | **Patch** /velocity-template-loaders/{velocity-template-loader-name} | Update an existing Velocity Template Loader by name



## AddVelocityTemplateLoader

> VelocityTemplateLoaderResponse AddVelocityTemplateLoader(ctx).AddVelocityTemplateLoaderRequest(addVelocityTemplateLoaderRequest).Execute()

Add a new Velocity Template Loader to the config

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
    addVelocityTemplateLoaderRequest := *openapiclient.NewAddVelocityTemplateLoaderRequest("LoaderName_example", int32(123), "MimeTypeMatcher_example") // AddVelocityTemplateLoaderRequest | Create a new Velocity Template Loader in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityTemplateLoaderApi.AddVelocityTemplateLoader(context.Background()).AddVelocityTemplateLoaderRequest(addVelocityTemplateLoaderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityTemplateLoaderApi.AddVelocityTemplateLoader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVelocityTemplateLoader`: VelocityTemplateLoaderResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityTemplateLoaderApi.AddVelocityTemplateLoader`: %v\n", resp)
}
```

### Path Parameters



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

> DeleteVelocityTemplateLoader(ctx, velocityTemplateLoaderName).Execute()

Delete a Velocity Template Loader

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
    velocityTemplateLoaderName := "velocityTemplateLoaderName_example" // string | Name of the Velocity Template Loader to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityTemplateLoaderApi.DeleteVelocityTemplateLoader(context.Background(), velocityTemplateLoaderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityTemplateLoaderApi.DeleteVelocityTemplateLoader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityTemplateLoaderName** | **string** | Name of the Velocity Template Loader to be deleted | 

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

> VelocityTemplateLoaderResponse GetVelocityTemplateLoader(ctx, velocityTemplateLoaderName).Execute()

Returns a single Velocity Template Loader

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
    velocityTemplateLoaderName := "velocityTemplateLoaderName_example" // string | Name of the Velocity Template Loader to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityTemplateLoaderApi.GetVelocityTemplateLoader(context.Background(), velocityTemplateLoaderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityTemplateLoaderApi.GetVelocityTemplateLoader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVelocityTemplateLoader`: VelocityTemplateLoaderResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityTemplateLoaderApi.GetVelocityTemplateLoader`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityTemplateLoaderName** | **string** | Name of the Velocity Template Loader to be read | 

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


## UpdateVelocityTemplateLoader

> VelocityTemplateLoaderResponse UpdateVelocityTemplateLoader(ctx, velocityTemplateLoaderName).UpdateRequest(updateRequest).Execute()

Update an existing Velocity Template Loader by name

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
    velocityTemplateLoaderName := "velocityTemplateLoaderName_example" // string | Name of the Velocity Template Loader to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Velocity Template Loader

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityTemplateLoaderApi.UpdateVelocityTemplateLoader(context.Background(), velocityTemplateLoaderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityTemplateLoaderApi.UpdateVelocityTemplateLoader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVelocityTemplateLoader`: VelocityTemplateLoaderResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityTemplateLoaderApi.UpdateVelocityTemplateLoader`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityTemplateLoaderName** | **string** | Name of the Velocity Template Loader to be updated | 

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

