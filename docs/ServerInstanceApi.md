# \ServerInstanceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServerInstance**](ServerInstanceApi.md#GetServerInstance) | **Get** /server-instances/{server-instance-name} | Returns a single Server Instance
[**UpdateServerInstance**](ServerInstanceApi.md#UpdateServerInstance) | **Patch** /server-instances/{server-instance-name} | Update an existing Server Instance by name



## GetServerInstance

> GetServerInstance200Response GetServerInstance(ctx, serverInstanceName).Execute()

Returns a single Server Instance

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
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerInstanceApi.GetServerInstance(context.Background(), serverInstanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceApi.GetServerInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerInstance`: GetServerInstance200Response
    fmt.Fprintf(os.Stdout, "Response from `ServerInstanceApi.GetServerInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceName** | **string** | Name of the Server Instance to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServerInstance200Response**](GetServerInstance200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerInstance

> GetServerInstance200Response UpdateServerInstance(ctx, serverInstanceName).UpdateRequest(updateRequest).Execute()

Update an existing Server Instance by name

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
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Server Instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerInstanceApi.UpdateServerInstance(context.Background(), serverInstanceName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceApi.UpdateServerInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServerInstance`: GetServerInstance200Response
    fmt.Fprintf(os.Stdout, "Response from `ServerInstanceApi.UpdateServerInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceName** | **string** | Name of the Server Instance to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Server Instance | 

### Return type

[**GetServerInstance200Response**](GetServerInstance200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

