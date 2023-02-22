# \CipherStreamProviderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCipherStreamProvider**](CipherStreamProviderApi.md#AddCipherStreamProvider) | **Post** /cipher-stream-providers | Add a new Cipher Stream Provider to the config
[**DeleteCipherStreamProvider**](CipherStreamProviderApi.md#DeleteCipherStreamProvider) | **Delete** /cipher-stream-providers/{cipher-stream-provider-name} | Delete a Cipher Stream Provider
[**GetCipherStreamProvider**](CipherStreamProviderApi.md#GetCipherStreamProvider) | **Get** /cipher-stream-providers/{cipher-stream-provider-name} | Returns a single Cipher Stream Provider
[**UpdateCipherStreamProvider**](CipherStreamProviderApi.md#UpdateCipherStreamProvider) | **Patch** /cipher-stream-providers/{cipher-stream-provider-name} | Update an existing Cipher Stream Provider by name



## AddCipherStreamProvider

> AddCipherStreamProvider200Response AddCipherStreamProvider(ctx).AddCipherStreamProviderRequest(addCipherStreamProviderRequest).Execute()

Add a new Cipher Stream Provider to the config

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
    addCipherStreamProviderRequest := openapiclient.add_cipher_stream_provider_request{AddAmazonKeyManagementServiceCipherStreamProviderRequest: openapiclient.NewAddAmazonKeyManagementServiceCipherStreamProviderRequest("ProviderName_example", []openapiclient.EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn{openapiclient.Enumamazon-key-management-service-cipher-stream-providerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:cipher-stream-provider:amazon-key-management-service")}, "EncryptedPassphraseFile_example", "KmsEncryptionKeyArn_example", false)} // AddCipherStreamProviderRequest | Create a new Cipher Stream Provider in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherStreamProviderApi.AddCipherStreamProvider(context.Background()).AddCipherStreamProviderRequest(addCipherStreamProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherStreamProviderApi.AddCipherStreamProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCipherStreamProvider`: AddCipherStreamProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `CipherStreamProviderApi.AddCipherStreamProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCipherStreamProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addCipherStreamProviderRequest** | [**AddCipherStreamProviderRequest**](AddCipherStreamProviderRequest.md) | Create a new Cipher Stream Provider in the config | 

### Return type

[**AddCipherStreamProvider200Response**](AddCipherStreamProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCipherStreamProvider

> DeleteCipherStreamProvider(ctx, cipherStreamProviderName).Execute()

Delete a Cipher Stream Provider

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
    cipherStreamProviderName := "cipherStreamProviderName_example" // string | Name of the Cipher Stream Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherStreamProviderApi.DeleteCipherStreamProvider(context.Background(), cipherStreamProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherStreamProviderApi.DeleteCipherStreamProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cipherStreamProviderName** | **string** | Name of the Cipher Stream Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCipherStreamProviderRequest struct via the builder pattern


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


## GetCipherStreamProvider

> AddCipherStreamProvider200Response GetCipherStreamProvider(ctx, cipherStreamProviderName).Execute()

Returns a single Cipher Stream Provider

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
    cipherStreamProviderName := "cipherStreamProviderName_example" // string | Name of the Cipher Stream Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherStreamProviderApi.GetCipherStreamProvider(context.Background(), cipherStreamProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherStreamProviderApi.GetCipherStreamProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCipherStreamProvider`: AddCipherStreamProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `CipherStreamProviderApi.GetCipherStreamProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cipherStreamProviderName** | **string** | Name of the Cipher Stream Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCipherStreamProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddCipherStreamProvider200Response**](AddCipherStreamProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCipherStreamProvider

> AddCipherStreamProvider200Response UpdateCipherStreamProvider(ctx, cipherStreamProviderName).UpdateRequest(updateRequest).Execute()

Update an existing Cipher Stream Provider by name

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
    cipherStreamProviderName := "cipherStreamProviderName_example" // string | Name of the Cipher Stream Provider
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Cipher Stream Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherStreamProviderApi.UpdateCipherStreamProvider(context.Background(), cipherStreamProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherStreamProviderApi.UpdateCipherStreamProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCipherStreamProvider`: AddCipherStreamProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `CipherStreamProviderApi.UpdateCipherStreamProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cipherStreamProviderName** | **string** | Name of the Cipher Stream Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCipherStreamProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Cipher Stream Provider | 

### Return type

[**AddCipherStreamProvider200Response**](AddCipherStreamProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

