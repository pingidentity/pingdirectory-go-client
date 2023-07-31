# \ServerInstanceListenerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServerInstanceListener**](ServerInstanceListenerApi.md#GetServerInstanceListener) | **Get** /server-instances/{server-instance-name}/server-instance-listeners/{server-instance-listener-name} | Returns a single Server Instance Listener
[**ListServerInstanceListeners**](ServerInstanceListenerApi.md#ListServerInstanceListeners) | **Get** /server-instances/{server-instance-name}/server-instance-listeners | Returns a list of all Server Instance Listener objects
[**UpdateServerInstanceListener**](ServerInstanceListenerApi.md#UpdateServerInstanceListener) | **Patch** /server-instances/{server-instance-name}/server-instance-listeners/{server-instance-listener-name} | Update an existing Server Instance Listener by name



## GetServerInstanceListener

> GetServerInstanceListener200Response GetServerInstanceListener(ctx, serverInstanceListenerName, serverInstanceName).Execute()

Returns a single Server Instance Listener

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
    serverInstanceListenerName := "serverInstanceListenerName_example" // string | Name of the Server Instance Listener
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerInstanceListenerApi.GetServerInstanceListener(context.Background(), serverInstanceListenerName, serverInstanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceListenerApi.GetServerInstanceListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerInstanceListener`: GetServerInstanceListener200Response
    fmt.Fprintf(os.Stdout, "Response from `ServerInstanceListenerApi.GetServerInstanceListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceListenerName** | **string** | Name of the Server Instance Listener | 
**serverInstanceName** | **string** | Name of the Server Instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerInstanceListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetServerInstanceListener200Response**](GetServerInstanceListener200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerInstanceListeners

> ServerInstanceListenerListResponse ListServerInstanceListeners(ctx, serverInstanceName).Filter(filter).Execute()

Returns a list of all Server Instance Listener objects

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
    resp, r, err := apiClient.ServerInstanceListenerApi.ListServerInstanceListeners(context.Background(), serverInstanceName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceListenerApi.ListServerInstanceListeners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServerInstanceListeners`: ServerInstanceListenerListResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerInstanceListenerApi.ListServerInstanceListeners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceName** | **string** | Name of the Server Instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerInstanceListenersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**ServerInstanceListenerListResponse**](ServerInstanceListenerListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerInstanceListener

> GetServerInstanceListener200Response UpdateServerInstanceListener(ctx, serverInstanceListenerName, serverInstanceName).UpdateRequest(updateRequest).Execute()

Update an existing Server Instance Listener by name

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
    serverInstanceListenerName := "serverInstanceListenerName_example" // string | Name of the Server Instance Listener
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Server Instance Listener

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerInstanceListenerApi.UpdateServerInstanceListener(context.Background(), serverInstanceListenerName, serverInstanceName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerInstanceListenerApi.UpdateServerInstanceListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServerInstanceListener`: GetServerInstanceListener200Response
    fmt.Fprintf(os.Stdout, "Response from `ServerInstanceListenerApi.UpdateServerInstanceListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceListenerName** | **string** | Name of the Server Instance Listener | 
**serverInstanceName** | **string** | Name of the Server Instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerInstanceListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Server Instance Listener | 

### Return type

[**GetServerInstanceListener200Response**](GetServerInstanceListener200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

