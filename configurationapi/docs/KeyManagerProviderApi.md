# \KeyManagerProviderAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKeyManagerProvider**](KeyManagerProviderAPI.md#AddKeyManagerProvider) | **Post** /key-manager-providers | Add a new Key Manager Provider to the config
[**DeleteKeyManagerProvider**](KeyManagerProviderAPI.md#DeleteKeyManagerProvider) | **Delete** /key-manager-providers/{key-manager-provider-name} | Delete a Key Manager Provider
[**GetKeyManagerProvider**](KeyManagerProviderAPI.md#GetKeyManagerProvider) | **Get** /key-manager-providers/{key-manager-provider-name} | Returns a single Key Manager Provider
[**ListKeyManagerProviders**](KeyManagerProviderAPI.md#ListKeyManagerProviders) | **Get** /key-manager-providers | Returns a list of all Key Manager Provider objects
[**UpdateKeyManagerProvider**](KeyManagerProviderAPI.md#UpdateKeyManagerProvider) | **Patch** /key-manager-providers/{key-manager-provider-name} | Update an existing Key Manager Provider by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addKeyManagerProviderRequest := openapiclient.add_key_manager_provider_request{AddFileBasedKeyManagerProviderRequest: openapiclient.NewAddFileBasedKeyManagerProviderRequest([]openapiclient.EnumfileBasedKeyManagerProviderSchemaUrn{openapiclient.Enumfile-based-key-manager-providerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:key-manager-provider:file-based")}, "KeyStoreFile_example", false, "ProviderName_example")} // AddKeyManagerProviderRequest | Create a new Key Manager Provider in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagerProviderAPI.AddKeyManagerProvider(context.Background()).AddKeyManagerProviderRequest(addKeyManagerProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagerProviderAPI.AddKeyManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddKeyManagerProvider`: AddKeyManagerProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `KeyManagerProviderAPI.AddKeyManagerProvider`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    keyManagerProviderName := "keyManagerProviderName_example" // string | Name of the Key Manager Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyManagerProviderAPI.DeleteKeyManagerProvider(context.Background(), keyManagerProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagerProviderAPI.DeleteKeyManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyManagerProviderName** | **string** | Name of the Key Manager Provider | 

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

> GetKeyManagerProvider200Response GetKeyManagerProvider(ctx, keyManagerProviderName).Execute()

Returns a single Key Manager Provider

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
    keyManagerProviderName := "keyManagerProviderName_example" // string | Name of the Key Manager Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagerProviderAPI.GetKeyManagerProvider(context.Background(), keyManagerProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagerProviderAPI.GetKeyManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyManagerProvider`: GetKeyManagerProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `KeyManagerProviderAPI.GetKeyManagerProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyManagerProviderName** | **string** | Name of the Key Manager Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyManagerProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetKeyManagerProvider200Response**](GetKeyManagerProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeyManagerProviders

> KeyManagerProviderListResponse ListKeyManagerProviders(ctx).Filter(filter).Execute()

Returns a list of all Key Manager Provider objects

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
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagerProviderAPI.ListKeyManagerProviders(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagerProviderAPI.ListKeyManagerProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKeyManagerProviders`: KeyManagerProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyManagerProviderAPI.ListKeyManagerProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKeyManagerProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**KeyManagerProviderListResponse**](KeyManagerProviderListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeyManagerProvider

> GetKeyManagerProvider200Response UpdateKeyManagerProvider(ctx, keyManagerProviderName).UpdateRequest(updateRequest).Execute()

Update an existing Key Manager Provider by name

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
    keyManagerProviderName := "keyManagerProviderName_example" // string | Name of the Key Manager Provider
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Key Manager Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyManagerProviderAPI.UpdateKeyManagerProvider(context.Background(), keyManagerProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyManagerProviderAPI.UpdateKeyManagerProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKeyManagerProvider`: GetKeyManagerProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `KeyManagerProviderAPI.UpdateKeyManagerProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyManagerProviderName** | **string** | Name of the Key Manager Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyManagerProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Key Manager Provider | 

### Return type

[**GetKeyManagerProvider200Response**](GetKeyManagerProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

