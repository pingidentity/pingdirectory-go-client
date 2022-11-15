# \ConnectionHandlerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConnectionHandler**](ConnectionHandlerApi.md#AddConnectionHandler) | **Post** /connection-handlers | Add a new Connection Handler to the config
[**DeleteConnectionHandler**](ConnectionHandlerApi.md#DeleteConnectionHandler) | **Delete** /connection-handlers/{connection-handler-name} | Delete a Connection Handler
[**GetConnectionHandler**](ConnectionHandlerApi.md#GetConnectionHandler) | **Get** /connection-handlers/{connection-handler-name} | Returns a single Connection Handler
[**UpdateConnectionHandler**](ConnectionHandlerApi.md#UpdateConnectionHandler) | **Patch** /connection-handlers/{connection-handler-name} | Update an existing Connection Handler by name



## AddConnectionHandler

> AddConnectionHandler200Response AddConnectionHandler(ctx).AddConnectionHandlerRequest(addConnectionHandlerRequest).Execute()

Add a new Connection Handler to the config

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
    addConnectionHandlerRequest := openapiclient.add_connection_handler_request{AddHttpConnectionHandlerRequest: openapiclient.NewAddHttpConnectionHandlerRequest("HandlerName_example", []openapiclient.EnumhttpConnectionHandlerSchemaUrn{openapiclient.Enumhttp-connection-handlerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:connection-handler:http")}, int32(123), false)} // AddConnectionHandlerRequest | Create a new Connection Handler in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionHandlerApi.AddConnectionHandler(context.Background()).AddConnectionHandlerRequest(addConnectionHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionHandlerApi.AddConnectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConnectionHandler`: AddConnectionHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionHandlerApi.AddConnectionHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddConnectionHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addConnectionHandlerRequest** | [**AddConnectionHandlerRequest**](AddConnectionHandlerRequest.md) | Create a new Connection Handler in the config | 

### Return type

[**AddConnectionHandler200Response**](AddConnectionHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectionHandler

> DeleteConnectionHandler(ctx, connectionHandlerName).Execute()

Delete a Connection Handler

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
    connectionHandlerName := "connectionHandlerName_example" // string | Name of the Connection Handler to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionHandlerApi.DeleteConnectionHandler(context.Background(), connectionHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionHandlerApi.DeleteConnectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandlerName** | **string** | Name of the Connection Handler to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionHandlerRequest struct via the builder pattern


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


## GetConnectionHandler

> AddConnectionHandler200Response GetConnectionHandler(ctx, connectionHandlerName).Execute()

Returns a single Connection Handler

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
    connectionHandlerName := "connectionHandlerName_example" // string | Name of the Connection Handler to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionHandlerApi.GetConnectionHandler(context.Background(), connectionHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionHandlerApi.GetConnectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionHandler`: AddConnectionHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionHandlerApi.GetConnectionHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandlerName** | **string** | Name of the Connection Handler to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddConnectionHandler200Response**](AddConnectionHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionHandler

> AddConnectionHandler200Response UpdateConnectionHandler(ctx, connectionHandlerName).UpdateRequest(updateRequest).Execute()

Update an existing Connection Handler by name

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
    connectionHandlerName := "connectionHandlerName_example" // string | Name of the Connection Handler to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Connection Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionHandlerApi.UpdateConnectionHandler(context.Background(), connectionHandlerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionHandlerApi.UpdateConnectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectionHandler`: AddConnectionHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionHandlerApi.UpdateConnectionHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandlerName** | **string** | Name of the Connection Handler to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Connection Handler | 

### Return type

[**AddConnectionHandler200Response**](AddConnectionHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

