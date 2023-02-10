# \CipherSecretKeyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCipherSecretKey**](CipherSecretKeyApi.md#GetCipherSecretKey) | **Get** /cipher-secret-keys/{cipher-secret-key-name} | Returns a single Cipher Secret Key
[**UpdateCipherSecretKey**](CipherSecretKeyApi.md#UpdateCipherSecretKey) | **Patch** /cipher-secret-keys/{cipher-secret-key-name} | Update an existing Cipher Secret Key by name



## GetCipherSecretKey

> CipherSecretKeyResponse GetCipherSecretKey(ctx, cipherSecretKeyName).Execute()

Returns a single Cipher Secret Key

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
    cipherSecretKeyName := "cipherSecretKeyName_example" // string | Name of the Cipher Secret Key to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherSecretKeyApi.GetCipherSecretKey(context.Background(), cipherSecretKeyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherSecretKeyApi.GetCipherSecretKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCipherSecretKey`: CipherSecretKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `CipherSecretKeyApi.GetCipherSecretKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cipherSecretKeyName** | **string** | Name of the Cipher Secret Key to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCipherSecretKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CipherSecretKeyResponse**](CipherSecretKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCipherSecretKey

> CipherSecretKeyResponse UpdateCipherSecretKey(ctx, cipherSecretKeyName).UpdateRequest(updateRequest).Execute()

Update an existing Cipher Secret Key by name

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
    cipherSecretKeyName := "cipherSecretKeyName_example" // string | Name of the Cipher Secret Key to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Cipher Secret Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherSecretKeyApi.UpdateCipherSecretKey(context.Background(), cipherSecretKeyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherSecretKeyApi.UpdateCipherSecretKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCipherSecretKey`: CipherSecretKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `CipherSecretKeyApi.UpdateCipherSecretKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cipherSecretKeyName** | **string** | Name of the Cipher Secret Key to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCipherSecretKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Cipher Secret Key | 

### Return type

[**CipherSecretKeyResponse**](CipherSecretKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

