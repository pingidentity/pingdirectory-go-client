# \CipherStreamProviderAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCipherStreamProvider**](CipherStreamProviderAPI.md#AddCipherStreamProvider) | **Post** /cipher-stream-providers | Add a new Cipher Stream Provider to the config
[**DeleteCipherStreamProvider**](CipherStreamProviderAPI.md#DeleteCipherStreamProvider) | **Delete** /cipher-stream-providers/{cipher-stream-provider-name} | Delete a Cipher Stream Provider
[**GetCipherStreamProvider**](CipherStreamProviderAPI.md#GetCipherStreamProvider) | **Get** /cipher-stream-providers/{cipher-stream-provider-name} | Returns a single Cipher Stream Provider
[**ListCipherStreamProviders**](CipherStreamProviderAPI.md#ListCipherStreamProviders) | **Get** /cipher-stream-providers | Returns a list of all Cipher Stream Provider objects
[**UpdateCipherStreamProvider**](CipherStreamProviderAPI.md#UpdateCipherStreamProvider) | **Patch** /cipher-stream-providers/{cipher-stream-provider-name} | Update an existing Cipher Stream Provider by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addCipherStreamProviderRequest := openapiclient.add_cipher_stream_provider_request{AddAmazonKeyManagementServiceCipherStreamProviderRequest: openapiclient.NewAddAmazonKeyManagementServiceCipherStreamProviderRequest([]openapiclient.EnumamazonKeyManagementServiceCipherStreamProviderSchemaUrn{openapiclient.Enumamazon-key-management-service-cipher-stream-providerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:cipher-stream-provider:amazon-key-management-service")}, "KmsEncryptionKeyArn_example", false, "ProviderName_example")} // AddCipherStreamProviderRequest | Create a new Cipher Stream Provider in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherStreamProviderAPI.AddCipherStreamProvider(context.Background()).AddCipherStreamProviderRequest(addCipherStreamProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherStreamProviderAPI.AddCipherStreamProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCipherStreamProvider`: AddCipherStreamProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `CipherStreamProviderAPI.AddCipherStreamProvider`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    cipherStreamProviderName := "cipherStreamProviderName_example" // string | Name of the Cipher Stream Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CipherStreamProviderAPI.DeleteCipherStreamProvider(context.Background(), cipherStreamProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherStreamProviderAPI.DeleteCipherStreamProvider``: %v\n", err)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    cipherStreamProviderName := "cipherStreamProviderName_example" // string | Name of the Cipher Stream Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherStreamProviderAPI.GetCipherStreamProvider(context.Background(), cipherStreamProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherStreamProviderAPI.GetCipherStreamProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCipherStreamProvider`: AddCipherStreamProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `CipherStreamProviderAPI.GetCipherStreamProvider`: %v\n", resp)
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


## ListCipherStreamProviders

> CipherStreamProviderListResponse ListCipherStreamProviders(ctx).Filter(filter).Execute()

Returns a list of all Cipher Stream Provider objects

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
    resp, r, err := apiClient.CipherStreamProviderAPI.ListCipherStreamProviders(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherStreamProviderAPI.ListCipherStreamProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCipherStreamProviders`: CipherStreamProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `CipherStreamProviderAPI.ListCipherStreamProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCipherStreamProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**CipherStreamProviderListResponse**](CipherStreamProviderListResponse.md)

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    cipherStreamProviderName := "cipherStreamProviderName_example" // string | Name of the Cipher Stream Provider
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Cipher Stream Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherStreamProviderAPI.UpdateCipherStreamProvider(context.Background(), cipherStreamProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherStreamProviderAPI.UpdateCipherStreamProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCipherStreamProvider`: AddCipherStreamProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `CipherStreamProviderAPI.UpdateCipherStreamProvider`: %v\n", resp)
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

