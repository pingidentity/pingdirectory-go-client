# \AccessControlHandlerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessControlHandler**](AccessControlHandlerApi.md#GetAccessControlHandler) | **Get** /access-control-handler | Returns a single Access Control Handler
[**UpdateAccessControlHandler**](AccessControlHandlerApi.md#UpdateAccessControlHandler) | **Patch** /access-control-handler | Update an existing Access Control Handler by name



## GetAccessControlHandler

> DseeCompatAccessControlHandlerResponse GetAccessControlHandler(ctx).Execute()

Returns a single Access Control Handler

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
    resp, r, err := apiClient.AccessControlHandlerApi.GetAccessControlHandler(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessControlHandlerApi.GetAccessControlHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessControlHandler`: DseeCompatAccessControlHandlerResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessControlHandlerApi.GetAccessControlHandler`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessControlHandlerRequest struct via the builder pattern


### Return type

[**DseeCompatAccessControlHandlerResponse**](DseeCompatAccessControlHandlerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessControlHandler

> DseeCompatAccessControlHandlerResponse UpdateAccessControlHandler(ctx).UpdateRequest(updateRequest).Execute()

Update an existing Access Control Handler by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Access Control Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessControlHandlerApi.UpdateAccessControlHandler(context.Background()).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessControlHandlerApi.UpdateAccessControlHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessControlHandler`: DseeCompatAccessControlHandlerResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessControlHandlerApi.UpdateAccessControlHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessControlHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Access Control Handler | 

### Return type

[**DseeCompatAccessControlHandlerResponse**](DseeCompatAccessControlHandlerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

