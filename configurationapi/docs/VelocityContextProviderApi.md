# \VelocityContextProviderAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVelocityContextProvider**](VelocityContextProviderAPI.md#AddVelocityContextProvider) | **Post** /http-servlet-extensions/{http-servlet-extension-name}/velocity-context-providers | Add a new Velocity Context Provider to the config
[**DeleteVelocityContextProvider**](VelocityContextProviderAPI.md#DeleteVelocityContextProvider) | **Delete** /http-servlet-extensions/{http-servlet-extension-name}/velocity-context-providers/{velocity-context-provider-name} | Delete a Velocity Context Provider
[**GetVelocityContextProvider**](VelocityContextProviderAPI.md#GetVelocityContextProvider) | **Get** /http-servlet-extensions/{http-servlet-extension-name}/velocity-context-providers/{velocity-context-provider-name} | Returns a single Velocity Context Provider
[**ListVelocityContextProviders**](VelocityContextProviderAPI.md#ListVelocityContextProviders) | **Get** /http-servlet-extensions/{http-servlet-extension-name}/velocity-context-providers | Returns a list of all Velocity Context Provider objects
[**UpdateVelocityContextProvider**](VelocityContextProviderAPI.md#UpdateVelocityContextProvider) | **Patch** /http-servlet-extensions/{http-servlet-extension-name}/velocity-context-providers/{velocity-context-provider-name} | Update an existing Velocity Context Provider by name



## AddVelocityContextProvider

> AddVelocityContextProvider200Response AddVelocityContextProvider(ctx, httpServletExtensionName).AddVelocityContextProviderRequest(addVelocityContextProviderRequest).Execute()

Add a new Velocity Context Provider to the config

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
    addVelocityContextProviderRequest := openapiclient.add_velocity_context_provider_request{AddThirdPartyVelocityContextProviderRequest: openapiclient.NewAddThirdPartyVelocityContextProviderRequest([]openapiclient.EnumthirdPartyVelocityContextProviderSchemaUrn{openapiclient.Enumthird-party-velocity-context-providerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:velocity-context-provider:third-party")}, "ExtensionClass_example", "ProviderName_example")} // AddVelocityContextProviderRequest | Create a new Velocity Context Provider in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityContextProviderAPI.AddVelocityContextProvider(context.Background(), httpServletExtensionName).AddVelocityContextProviderRequest(addVelocityContextProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityContextProviderAPI.AddVelocityContextProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVelocityContextProvider`: AddVelocityContextProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `VelocityContextProviderAPI.AddVelocityContextProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVelocityContextProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addVelocityContextProviderRequest** | [**AddVelocityContextProviderRequest**](AddVelocityContextProviderRequest.md) | Create a new Velocity Context Provider in the config | 

### Return type

[**AddVelocityContextProvider200Response**](AddVelocityContextProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVelocityContextProvider

> DeleteVelocityContextProvider(ctx, velocityContextProviderName, httpServletExtensionName).Execute()

Delete a Velocity Context Provider

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
    velocityContextProviderName := "velocityContextProviderName_example" // string | Name of the Velocity Context Provider
    httpServletExtensionName := "httpServletExtensionName_example" // string | Name of the HTTP Servlet Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VelocityContextProviderAPI.DeleteVelocityContextProvider(context.Background(), velocityContextProviderName, httpServletExtensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityContextProviderAPI.DeleteVelocityContextProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityContextProviderName** | **string** | Name of the Velocity Context Provider | 
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVelocityContextProviderRequest struct via the builder pattern


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


## GetVelocityContextProvider

> GetVelocityContextProvider200Response GetVelocityContextProvider(ctx, velocityContextProviderName, httpServletExtensionName).Execute()

Returns a single Velocity Context Provider

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
    velocityContextProviderName := "velocityContextProviderName_example" // string | Name of the Velocity Context Provider
    httpServletExtensionName := "httpServletExtensionName_example" // string | Name of the HTTP Servlet Extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityContextProviderAPI.GetVelocityContextProvider(context.Background(), velocityContextProviderName, httpServletExtensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityContextProviderAPI.GetVelocityContextProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVelocityContextProvider`: GetVelocityContextProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `VelocityContextProviderAPI.GetVelocityContextProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityContextProviderName** | **string** | Name of the Velocity Context Provider | 
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVelocityContextProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVelocityContextProvider200Response**](GetVelocityContextProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVelocityContextProviders

> VelocityContextProviderListResponse ListVelocityContextProviders(ctx, httpServletExtensionName).Filter(filter).Execute()

Returns a list of all Velocity Context Provider objects

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
    resp, r, err := apiClient.VelocityContextProviderAPI.ListVelocityContextProviders(context.Background(), httpServletExtensionName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityContextProviderAPI.ListVelocityContextProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVelocityContextProviders`: VelocityContextProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `VelocityContextProviderAPI.ListVelocityContextProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVelocityContextProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**VelocityContextProviderListResponse**](VelocityContextProviderListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVelocityContextProvider

> GetVelocityContextProvider200Response UpdateVelocityContextProvider(ctx, velocityContextProviderName, httpServletExtensionName).UpdateRequest(updateRequest).Execute()

Update an existing Velocity Context Provider by name

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
    velocityContextProviderName := "velocityContextProviderName_example" // string | Name of the Velocity Context Provider
    httpServletExtensionName := "httpServletExtensionName_example" // string | Name of the HTTP Servlet Extension
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Velocity Context Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityContextProviderAPI.UpdateVelocityContextProvider(context.Background(), velocityContextProviderName, httpServletExtensionName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityContextProviderAPI.UpdateVelocityContextProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVelocityContextProvider`: GetVelocityContextProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `VelocityContextProviderAPI.UpdateVelocityContextProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityContextProviderName** | **string** | Name of the Velocity Context Provider | 
**httpServletExtensionName** | **string** | Name of the HTTP Servlet Extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVelocityContextProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Velocity Context Provider | 

### Return type

[**GetVelocityContextProvider200Response**](GetVelocityContextProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

