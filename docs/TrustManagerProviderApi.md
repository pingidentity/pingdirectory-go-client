# \TrustManagerProviderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTrustManagerProvider**](TrustManagerProviderApi.md#AddTrustManagerProvider) | **Post** /trust-manager-providers | Add a new Trust Manager Provider to the config
[**DeleteTrustManagerProvider**](TrustManagerProviderApi.md#DeleteTrustManagerProvider) | **Delete** /trust-manager-providers/{trust-manager-provider-name} | Delete a Trust Manager Provider
[**GetTrustManagerProvider**](TrustManagerProviderApi.md#GetTrustManagerProvider) | **Get** /trust-manager-providers/{trust-manager-provider-name} | Returns a single Trust Manager Provider
[**UpdateTrustManagerProvider**](TrustManagerProviderApi.md#UpdateTrustManagerProvider) | **Patch** /trust-manager-providers/{trust-manager-provider-name} | Update an existing Trust Manager Provider by name



## AddTrustManagerProvider

> AddTrustManagerProvider200Response AddTrustManagerProvider(ctx).AddTrustManagerProviderRequest(addTrustManagerProviderRequest).Execute()

Add a new Trust Manager Provider to the config

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
    addTrustManagerProviderRequest := openapiclient.add_trust_manager_provider_request{AddBlindTrustManagerProviderRequest: openapiclient.NewAddBlindTrustManagerProviderRequest("ProviderName_example", []openapiclient.EnumblindTrustManagerProviderSchemaUrn{openapiclient.Enumblind-trust-manager-providerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:trust-manager-provider:blind")}, false)} // AddTrustManagerProviderRequest | Create a new Trust Manager Provider in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustManagerProviderApi.AddTrustManagerProvider(context.Background()).AddTrustManagerProviderRequest(addTrustManagerProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustManagerProviderApi.AddTrustManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTrustManagerProvider`: AddTrustManagerProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `TrustManagerProviderApi.AddTrustManagerProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTrustManagerProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addTrustManagerProviderRequest** | [**AddTrustManagerProviderRequest**](AddTrustManagerProviderRequest.md) | Create a new Trust Manager Provider in the config | 

### Return type

[**AddTrustManagerProvider200Response**](AddTrustManagerProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrustManagerProvider

> DeleteTrustManagerProvider(ctx, trustManagerProviderName).Execute()

Delete a Trust Manager Provider

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
    trustManagerProviderName := "trustManagerProviderName_example" // string | Name of the Trust Manager Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustManagerProviderApi.DeleteTrustManagerProvider(context.Background(), trustManagerProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustManagerProviderApi.DeleteTrustManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustManagerProviderName** | **string** | Name of the Trust Manager Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrustManagerProviderRequest struct via the builder pattern


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


## GetTrustManagerProvider

> AddTrustManagerProvider200Response GetTrustManagerProvider(ctx, trustManagerProviderName).Execute()

Returns a single Trust Manager Provider

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
    trustManagerProviderName := "trustManagerProviderName_example" // string | Name of the Trust Manager Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustManagerProviderApi.GetTrustManagerProvider(context.Background(), trustManagerProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustManagerProviderApi.GetTrustManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrustManagerProvider`: AddTrustManagerProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `TrustManagerProviderApi.GetTrustManagerProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustManagerProviderName** | **string** | Name of the Trust Manager Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustManagerProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddTrustManagerProvider200Response**](AddTrustManagerProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrustManagerProvider

> AddTrustManagerProvider200Response UpdateTrustManagerProvider(ctx, trustManagerProviderName).UpdateRequest(updateRequest).Execute()

Update an existing Trust Manager Provider by name

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
    trustManagerProviderName := "trustManagerProviderName_example" // string | Name of the Trust Manager Provider
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Trust Manager Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustManagerProviderApi.UpdateTrustManagerProvider(context.Background(), trustManagerProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustManagerProviderApi.UpdateTrustManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrustManagerProvider`: AddTrustManagerProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `TrustManagerProviderApi.UpdateTrustManagerProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustManagerProviderName** | **string** | Name of the Trust Manager Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrustManagerProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Trust Manager Provider | 

### Return type

[**AddTrustManagerProvider200Response**](AddTrustManagerProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

