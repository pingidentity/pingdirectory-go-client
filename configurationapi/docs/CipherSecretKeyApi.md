# \CipherSecretKeyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCipherSecretKey**](CipherSecretKeyAPI.md#GetCipherSecretKey) | **Get** /server-instances/{server-instance-name}/cipher-secret-keys/{cipher-secret-key-name} | Returns a single Cipher Secret Key
[**ListCipherSecretKeys**](CipherSecretKeyAPI.md#ListCipherSecretKeys) | **Get** /server-instances/{server-instance-name}/cipher-secret-keys | Returns a list of all Cipher Secret Key objects
[**UpdateCipherSecretKey**](CipherSecretKeyAPI.md#UpdateCipherSecretKey) | **Patch** /server-instances/{server-instance-name}/cipher-secret-keys/{cipher-secret-key-name} | Update an existing Cipher Secret Key by name



## GetCipherSecretKey

> CipherSecretKeyResponse GetCipherSecretKey(ctx, cipherSecretKeyName, serverInstanceName).Execute()

Returns a single Cipher Secret Key

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
    cipherSecretKeyName := "cipherSecretKeyName_example" // string | Name of the Cipher Secret Key
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherSecretKeyAPI.GetCipherSecretKey(context.Background(), cipherSecretKeyName, serverInstanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherSecretKeyAPI.GetCipherSecretKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCipherSecretKey`: CipherSecretKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `CipherSecretKeyAPI.GetCipherSecretKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cipherSecretKeyName** | **string** | Name of the Cipher Secret Key | 
**serverInstanceName** | **string** | Name of the Server Instance | 

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


## ListCipherSecretKeys

> CipherSecretKeyListResponse ListCipherSecretKeys(ctx, serverInstanceName).Filter(filter).Execute()

Returns a list of all Cipher Secret Key objects

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
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherSecretKeyAPI.ListCipherSecretKeys(context.Background(), serverInstanceName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherSecretKeyAPI.ListCipherSecretKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCipherSecretKeys`: CipherSecretKeyListResponse
    fmt.Fprintf(os.Stdout, "Response from `CipherSecretKeyAPI.ListCipherSecretKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceName** | **string** | Name of the Server Instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCipherSecretKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**CipherSecretKeyListResponse**](CipherSecretKeyListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCipherSecretKey

> CipherSecretKeyResponse UpdateCipherSecretKey(ctx, cipherSecretKeyName, serverInstanceName).UpdateRequest(updateRequest).Execute()

Update an existing Cipher Secret Key by name

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
    cipherSecretKeyName := "cipherSecretKeyName_example" // string | Name of the Cipher Secret Key
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Cipher Secret Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CipherSecretKeyAPI.UpdateCipherSecretKey(context.Background(), cipherSecretKeyName, serverInstanceName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CipherSecretKeyAPI.UpdateCipherSecretKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCipherSecretKey`: CipherSecretKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `CipherSecretKeyAPI.UpdateCipherSecretKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cipherSecretKeyName** | **string** | Name of the Cipher Secret Key | 
**serverInstanceName** | **string** | Name of the Server Instance | 

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

