# \MacSecretKeyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMacSecretKey**](MacSecretKeyApi.md#GetMacSecretKey) | **Get** /server-instances/{server-instance-name}/mac-secret-keys/{mac-secret-key-name} | Returns a single Mac Secret Key
[**ListMacSecretKeys**](MacSecretKeyApi.md#ListMacSecretKeys) | **Get** /server-instances/{server-instance-name}/mac-secret-keys | Returns a list of all Mac Secret Key objects
[**UpdateMacSecretKey**](MacSecretKeyApi.md#UpdateMacSecretKey) | **Patch** /server-instances/{server-instance-name}/mac-secret-keys/{mac-secret-key-name} | Update an existing Mac Secret Key by name



## GetMacSecretKey

> MacSecretKeyResponse GetMacSecretKey(ctx, macSecretKeyName, serverInstanceName).Execute()

Returns a single Mac Secret Key

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
    macSecretKeyName := "macSecretKeyName_example" // string | Name of the Mac Secret Key
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MacSecretKeyApi.GetMacSecretKey(context.Background(), macSecretKeyName, serverInstanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacSecretKeyApi.GetMacSecretKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMacSecretKey`: MacSecretKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `MacSecretKeyApi.GetMacSecretKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**macSecretKeyName** | **string** | Name of the Mac Secret Key | 
**serverInstanceName** | **string** | Name of the Server Instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMacSecretKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MacSecretKeyResponse**](MacSecretKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMacSecretKeys

> MacSecretKeyListResponse ListMacSecretKeys(ctx, serverInstanceName).Filter(filter).Execute()

Returns a list of all Mac Secret Key objects

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
    resp, r, err := apiClient.MacSecretKeyApi.ListMacSecretKeys(context.Background(), serverInstanceName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacSecretKeyApi.ListMacSecretKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMacSecretKeys`: MacSecretKeyListResponse
    fmt.Fprintf(os.Stdout, "Response from `MacSecretKeyApi.ListMacSecretKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceName** | **string** | Name of the Server Instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMacSecretKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**MacSecretKeyListResponse**](MacSecretKeyListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMacSecretKey

> MacSecretKeyResponse UpdateMacSecretKey(ctx, macSecretKeyName, serverInstanceName).UpdateRequest(updateRequest).Execute()

Update an existing Mac Secret Key by name

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
    macSecretKeyName := "macSecretKeyName_example" // string | Name of the Mac Secret Key
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Mac Secret Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MacSecretKeyApi.UpdateMacSecretKey(context.Background(), macSecretKeyName, serverInstanceName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacSecretKeyApi.UpdateMacSecretKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMacSecretKey`: MacSecretKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `MacSecretKeyApi.UpdateMacSecretKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**macSecretKeyName** | **string** | Name of the Mac Secret Key | 
**serverInstanceName** | **string** | Name of the Server Instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMacSecretKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Mac Secret Key | 

### Return type

[**MacSecretKeyResponse**](MacSecretKeyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

