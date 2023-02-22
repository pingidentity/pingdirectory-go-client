# \AlertHandlerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAlertHandler**](AlertHandlerApi.md#AddAlertHandler) | **Post** /alert-handlers | Add a new Alert Handler to the config
[**DeleteAlertHandler**](AlertHandlerApi.md#DeleteAlertHandler) | **Delete** /alert-handlers/{alert-handler-name} | Delete a Alert Handler
[**GetAlertHandler**](AlertHandlerApi.md#GetAlertHandler) | **Get** /alert-handlers/{alert-handler-name} | Returns a single Alert Handler
[**UpdateAlertHandler**](AlertHandlerApi.md#UpdateAlertHandler) | **Patch** /alert-handlers/{alert-handler-name} | Update an existing Alert Handler by name



## AddAlertHandler

> AddAlertHandler200Response AddAlertHandler(ctx).AddAlertHandlerRequest(addAlertHandlerRequest).Execute()

Add a new Alert Handler to the config

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
    addAlertHandlerRequest := openapiclient.add_alert_handler_request{AddErrorLogAlertHandlerRequest: openapiclient.NewAddErrorLogAlertHandlerRequest("HandlerName_example", []openapiclient.EnumerrorLogAlertHandlerSchemaUrn{openapiclient.Enumerror-log-alert-handlerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:alert-handler:error-log")}, false)} // AddAlertHandlerRequest | Create a new Alert Handler in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertHandlerApi.AddAlertHandler(context.Background()).AddAlertHandlerRequest(addAlertHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertHandlerApi.AddAlertHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAlertHandler`: AddAlertHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `AlertHandlerApi.AddAlertHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAlertHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addAlertHandlerRequest** | [**AddAlertHandlerRequest**](AddAlertHandlerRequest.md) | Create a new Alert Handler in the config | 

### Return type

[**AddAlertHandler200Response**](AddAlertHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertHandler

> DeleteAlertHandler(ctx, alertHandlerName).Execute()

Delete a Alert Handler

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
    alertHandlerName := "alertHandlerName_example" // string | Name of the Alert Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertHandlerApi.DeleteAlertHandler(context.Background(), alertHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertHandlerApi.DeleteAlertHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertHandlerName** | **string** | Name of the Alert Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertHandlerRequest struct via the builder pattern


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


## GetAlertHandler

> GetAlertHandler200Response GetAlertHandler(ctx, alertHandlerName).Execute()

Returns a single Alert Handler

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
    alertHandlerName := "alertHandlerName_example" // string | Name of the Alert Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertHandlerApi.GetAlertHandler(context.Background(), alertHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertHandlerApi.GetAlertHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertHandler`: GetAlertHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `AlertHandlerApi.GetAlertHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertHandlerName** | **string** | Name of the Alert Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAlertHandler200Response**](GetAlertHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertHandler

> GetAlertHandler200Response UpdateAlertHandler(ctx, alertHandlerName).UpdateRequest(updateRequest).Execute()

Update an existing Alert Handler by name

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
    alertHandlerName := "alertHandlerName_example" // string | Name of the Alert Handler
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Alert Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertHandlerApi.UpdateAlertHandler(context.Background(), alertHandlerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertHandlerApi.UpdateAlertHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertHandler`: GetAlertHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `AlertHandlerApi.UpdateAlertHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertHandlerName** | **string** | Name of the Alert Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Alert Handler | 

### Return type

[**GetAlertHandler200Response**](GetAlertHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

