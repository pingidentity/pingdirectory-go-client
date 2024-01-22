# \ConnectionHandlerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConnectionHandler**](ConnectionHandlerAPI.md#AddConnectionHandler) | **Post** /connection-handlers | Add a new Connection Handler to the config
[**DeleteConnectionHandler**](ConnectionHandlerAPI.md#DeleteConnectionHandler) | **Delete** /connection-handlers/{connection-handler-name} | Delete a Connection Handler
[**GetConnectionHandler**](ConnectionHandlerAPI.md#GetConnectionHandler) | **Get** /connection-handlers/{connection-handler-name} | Returns a single Connection Handler
[**ListConnectionHandlers**](ConnectionHandlerAPI.md#ListConnectionHandlers) | **Get** /connection-handlers | Returns a list of all Connection Handler objects
[**UpdateConnectionHandler**](ConnectionHandlerAPI.md#UpdateConnectionHandler) | **Patch** /connection-handlers/{connection-handler-name} | Update an existing Connection Handler by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addConnectionHandlerRequest := openapiclient.add_connection_handler_request{AddHttpConnectionHandlerRequest: openapiclient.NewAddHttpConnectionHandlerRequest([]openapiclient.EnumhttpConnectionHandlerSchemaUrn{openapiclient.Enumhttp-connection-handlerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:connection-handler:http")}, int64(123), false, "HandlerName_example")} // AddConnectionHandlerRequest | Create a new Connection Handler in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionHandlerAPI.AddConnectionHandler(context.Background()).AddConnectionHandlerRequest(addConnectionHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionHandlerAPI.AddConnectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConnectionHandler`: AddConnectionHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionHandlerAPI.AddConnectionHandler`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    connectionHandlerName := "connectionHandlerName_example" // string | Name of the Connection Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConnectionHandlerAPI.DeleteConnectionHandler(context.Background(), connectionHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionHandlerAPI.DeleteConnectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandlerName** | **string** | Name of the Connection Handler | 

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    connectionHandlerName := "connectionHandlerName_example" // string | Name of the Connection Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionHandlerAPI.GetConnectionHandler(context.Background(), connectionHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionHandlerAPI.GetConnectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionHandler`: AddConnectionHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionHandlerAPI.GetConnectionHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandlerName** | **string** | Name of the Connection Handler | 

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


## ListConnectionHandlers

> ConnectionHandlerListResponse ListConnectionHandlers(ctx).Filter(filter).Execute()

Returns a list of all Connection Handler objects

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
    resp, r, err := apiClient.ConnectionHandlerAPI.ListConnectionHandlers(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionHandlerAPI.ListConnectionHandlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectionHandlers`: ConnectionHandlerListResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionHandlerAPI.ListConnectionHandlers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionHandlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ConnectionHandlerListResponse**](ConnectionHandlerListResponse.md)

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    connectionHandlerName := "connectionHandlerName_example" // string | Name of the Connection Handler
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Connection Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionHandlerAPI.UpdateConnectionHandler(context.Background(), connectionHandlerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionHandlerAPI.UpdateConnectionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectionHandler`: AddConnectionHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ConnectionHandlerAPI.UpdateConnectionHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandlerName** | **string** | Name of the Connection Handler | 

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

