# \VelocityContextProviderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVelocityContextProvider**](VelocityContextProviderApi.md#AddVelocityContextProvider) | **Post** /velocity-context-providers | Add a new Velocity Context Provider to the config
[**DeleteVelocityContextProvider**](VelocityContextProviderApi.md#DeleteVelocityContextProvider) | **Delete** /velocity-context-providers/{velocity-context-provider-name} | Delete a Velocity Context Provider
[**GetVelocityContextProvider**](VelocityContextProviderApi.md#GetVelocityContextProvider) | **Get** /velocity-context-providers/{velocity-context-provider-name} | Returns a single Velocity Context Provider
[**UpdateVelocityContextProvider**](VelocityContextProviderApi.md#UpdateVelocityContextProvider) | **Patch** /velocity-context-providers/{velocity-context-provider-name} | Update an existing Velocity Context Provider by name



## AddVelocityContextProvider

> AddVelocityContextProvider200Response AddVelocityContextProvider(ctx).AddVelocityContextProviderRequest(addVelocityContextProviderRequest).Execute()

Add a new Velocity Context Provider to the config

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
    addVelocityContextProviderRequest := openapiclient.add_velocity_context_provider_request{AddThirdPartyVelocityContextProviderRequest: openapiclient.NewAddThirdPartyVelocityContextProviderRequest("ProviderName_example", []openapiclient.EnumthirdPartyVelocityContextProviderSchemaUrn{openapiclient.Enumthird-party-velocity-context-providerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:velocity-context-provider:third-party")}, "ExtensionClass_example")} // AddVelocityContextProviderRequest | Create a new Velocity Context Provider in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityContextProviderApi.AddVelocityContextProvider(context.Background()).AddVelocityContextProviderRequest(addVelocityContextProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityContextProviderApi.AddVelocityContextProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVelocityContextProvider`: AddVelocityContextProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `VelocityContextProviderApi.AddVelocityContextProvider`: %v\n", resp)
}
```

### Path Parameters



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

> DeleteVelocityContextProvider(ctx, velocityContextProviderName).Execute()

Delete a Velocity Context Provider

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
    velocityContextProviderName := "velocityContextProviderName_example" // string | Name of the Velocity Context Provider to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityContextProviderApi.DeleteVelocityContextProvider(context.Background(), velocityContextProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityContextProviderApi.DeleteVelocityContextProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityContextProviderName** | **string** | Name of the Velocity Context Provider to be deleted | 

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

> AddVelocityContextProvider200Response GetVelocityContextProvider(ctx, velocityContextProviderName).Execute()

Returns a single Velocity Context Provider

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
    velocityContextProviderName := "velocityContextProviderName_example" // string | Name of the Velocity Context Provider to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityContextProviderApi.GetVelocityContextProvider(context.Background(), velocityContextProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityContextProviderApi.GetVelocityContextProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVelocityContextProvider`: AddVelocityContextProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `VelocityContextProviderApi.GetVelocityContextProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityContextProviderName** | **string** | Name of the Velocity Context Provider to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVelocityContextProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddVelocityContextProvider200Response**](AddVelocityContextProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVelocityContextProvider

> AddVelocityContextProvider200Response UpdateVelocityContextProvider(ctx, velocityContextProviderName).UpdateRequest(updateRequest).Execute()

Update an existing Velocity Context Provider by name

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
    velocityContextProviderName := "velocityContextProviderName_example" // string | Name of the Velocity Context Provider to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Velocity Context Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VelocityContextProviderApi.UpdateVelocityContextProvider(context.Background(), velocityContextProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VelocityContextProviderApi.UpdateVelocityContextProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVelocityContextProvider`: AddVelocityContextProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `VelocityContextProviderApi.UpdateVelocityContextProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**velocityContextProviderName** | **string** | Name of the Velocity Context Provider to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVelocityContextProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Velocity Context Provider | 

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

