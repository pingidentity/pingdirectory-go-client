# \CryptoManagerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCryptoManager**](CryptoManagerApi.md#GetCryptoManager) | **Get** /crypto-manager | Returns a single Crypto Manager
[**UpdateCryptoManager**](CryptoManagerApi.md#UpdateCryptoManager) | **Patch** /crypto-manager | Update an existing Crypto Manager by name



## GetCryptoManager

> CryptoManagerResponse GetCryptoManager(ctx).Execute()

Returns a single Crypto Manager

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoManagerApi.GetCryptoManager(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoManagerApi.GetCryptoManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCryptoManager`: CryptoManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `CryptoManagerApi.GetCryptoManager`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCryptoManagerRequest struct via the builder pattern


### Return type

[**CryptoManagerResponse**](CryptoManagerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCryptoManager

> CryptoManagerResponse UpdateCryptoManager(ctx).UpdateRequest(updateRequest).Execute()

Update an existing Crypto Manager by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Crypto Manager

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CryptoManagerApi.UpdateCryptoManager(context.Background()).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CryptoManagerApi.UpdateCryptoManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCryptoManager`: CryptoManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `CryptoManagerApi.UpdateCryptoManager`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCryptoManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Crypto Manager | 

### Return type

[**CryptoManagerResponse**](CryptoManagerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

