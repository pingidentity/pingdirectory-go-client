# \KeyManagerProviderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKeyManagerProvider**](KeyManagerProviderApi.md#AddKeyManagerProvider) | **Post** /key-manager-providers | Add a new Key Manager Provider to the config
[**DeleteKeyManagerProvider**](KeyManagerProviderApi.md#DeleteKeyManagerProvider) | **Delete** /key-manager-providers/{key-manager-provider-name} | Delete a Key Manager Provider
[**GetKeyManagerProvider**](KeyManagerProviderApi.md#GetKeyManagerProvider) | **Get** /key-manager-providers/{key-manager-provider-name} | Returns a single Key Manager Provider
[**UpdateKeyManagerProvider**](KeyManagerProviderApi.md#UpdateKeyManagerProvider) | **Patch** /key-manager-providers/{key-manager-provider-name} | Update an existing Key Manager Provider by name



## AddKeyManagerProvider

> AddKeyManagerProvider200Response AddKeyManagerProvider(ctx).AddKeyManagerProviderRequest(addKeyManagerProviderRequest).Execute()

Add a new Key Manager Provider to the config

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
    addKeyManagerProviderRequest := openapiclient.add_key_manager_provider_request{AddFileBasedKeyManagerProviderRequest: openapiclient.NewAddFileBasedKeyManagerProviderRequest("ProviderName_example", []openapiclient.EnumfileBasedKeyManagerProviderSchemaUrn{openapiclient.Enumfile-based-key-manager-providerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:key-manager-provider:file-based")}, "KeyStoreFile_example", false)} // AddKeyManagerProviderRequest | Create a new Key Manager Provider in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagerProviderApi.AddKeyManagerProvider(context.Background()).AddKeyManagerProviderRequest(addKeyManagerProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagerProviderApi.AddKeyManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddKeyManagerProvider`: AddKeyManagerProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `KeyManagerProviderApi.AddKeyManagerProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddKeyManagerProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addKeyManagerProviderRequest** | [**AddKeyManagerProviderRequest**](AddKeyManagerProviderRequest.md) | Create a new Key Manager Provider in the config | 

### Return type

[**AddKeyManagerProvider200Response**](AddKeyManagerProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeyManagerProvider

> DeleteKeyManagerProvider(ctx, keyManagerProviderName).Execute()

Delete a Key Manager Provider

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
    keyManagerProviderName := "keyManagerProviderName_example" // string | Name of the Key Manager Provider to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagerProviderApi.DeleteKeyManagerProvider(context.Background(), keyManagerProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagerProviderApi.DeleteKeyManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyManagerProviderName** | **string** | Name of the Key Manager Provider to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyManagerProviderRequest struct via the builder pattern


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


## GetKeyManagerProvider

> AddKeyManagerProvider200Response GetKeyManagerProvider(ctx, keyManagerProviderName).Execute()

Returns a single Key Manager Provider

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
    keyManagerProviderName := "keyManagerProviderName_example" // string | Name of the Key Manager Provider to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagerProviderApi.GetKeyManagerProvider(context.Background(), keyManagerProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagerProviderApi.GetKeyManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyManagerProvider`: AddKeyManagerProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `KeyManagerProviderApi.GetKeyManagerProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyManagerProviderName** | **string** | Name of the Key Manager Provider to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyManagerProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddKeyManagerProvider200Response**](AddKeyManagerProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeyManagerProvider

> AddKeyManagerProvider200Response UpdateKeyManagerProvider(ctx, keyManagerProviderName).UpdateRequest(updateRequest).Execute()

Update an existing Key Manager Provider by name

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
    keyManagerProviderName := "keyManagerProviderName_example" // string | Name of the Key Manager Provider to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Key Manager Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagerProviderApi.UpdateKeyManagerProvider(context.Background(), keyManagerProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagerProviderApi.UpdateKeyManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKeyManagerProvider`: AddKeyManagerProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `KeyManagerProviderApi.UpdateKeyManagerProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyManagerProviderName** | **string** | Name of the Key Manager Provider to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyManagerProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Key Manager Provider | 

### Return type

[**AddKeyManagerProvider200Response**](AddKeyManagerProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

