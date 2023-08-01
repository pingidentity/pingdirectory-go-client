# \LogFileRotationListenerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogFileRotationListener**](LogFileRotationListenerApi.md#AddLogFileRotationListener) | **Post** /log-file-rotation-listeners | Add a new Log File Rotation Listener to the config
[**DeleteLogFileRotationListener**](LogFileRotationListenerApi.md#DeleteLogFileRotationListener) | **Delete** /log-file-rotation-listeners/{log-file-rotation-listener-name} | Delete a Log File Rotation Listener
[**GetLogFileRotationListener**](LogFileRotationListenerApi.md#GetLogFileRotationListener) | **Get** /log-file-rotation-listeners/{log-file-rotation-listener-name} | Returns a single Log File Rotation Listener
[**ListLogFileRotationListeners**](LogFileRotationListenerApi.md#ListLogFileRotationListeners) | **Get** /log-file-rotation-listeners | Returns a list of all Log File Rotation Listener objects
[**UpdateLogFileRotationListener**](LogFileRotationListenerApi.md#UpdateLogFileRotationListener) | **Patch** /log-file-rotation-listeners/{log-file-rotation-listener-name} | Update an existing Log File Rotation Listener by name



## AddLogFileRotationListener

> AddLogFileRotationListener200Response AddLogFileRotationListener(ctx).AddLogFileRotationListenerRequest(addLogFileRotationListenerRequest).Execute()

Add a new Log File Rotation Listener to the config

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
    addLogFileRotationListenerRequest := openapiclient.add_log_file_rotation_listener_request{AddCopyLogFileRotationListenerRequest: openapiclient.NewAddCopyLogFileRotationListenerRequest("ListenerName_example", []openapiclient.EnumcopyLogFileRotationListenerSchemaUrn{openapiclient.Enumcopy-log-file-rotation-listenerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:log-file-rotation-listener:copy")}, "CopyToDirectory_example", false)} // AddLogFileRotationListenerRequest | Create a new Log File Rotation Listener in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFileRotationListenerApi.AddLogFileRotationListener(context.Background()).AddLogFileRotationListenerRequest(addLogFileRotationListenerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFileRotationListenerApi.AddLogFileRotationListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLogFileRotationListener`: AddLogFileRotationListener200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFileRotationListenerApi.AddLogFileRotationListener`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLogFileRotationListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLogFileRotationListenerRequest** | [**AddLogFileRotationListenerRequest**](AddLogFileRotationListenerRequest.md) | Create a new Log File Rotation Listener in the config | 

### Return type

[**AddLogFileRotationListener200Response**](AddLogFileRotationListener200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogFileRotationListener

> DeleteLogFileRotationListener(ctx, logFileRotationListenerName).Execute()

Delete a Log File Rotation Listener

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
    logFileRotationListenerName := "logFileRotationListenerName_example" // string | Name of the Log File Rotation Listener

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LogFileRotationListenerApi.DeleteLogFileRotationListener(context.Background(), logFileRotationListenerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFileRotationListenerApi.DeleteLogFileRotationListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logFileRotationListenerName** | **string** | Name of the Log File Rotation Listener | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogFileRotationListenerRequest struct via the builder pattern


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


## GetLogFileRotationListener

> AddLogFileRotationListener200Response GetLogFileRotationListener(ctx, logFileRotationListenerName).Execute()

Returns a single Log File Rotation Listener

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
    logFileRotationListenerName := "logFileRotationListenerName_example" // string | Name of the Log File Rotation Listener

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFileRotationListenerApi.GetLogFileRotationListener(context.Background(), logFileRotationListenerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFileRotationListenerApi.GetLogFileRotationListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogFileRotationListener`: AddLogFileRotationListener200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFileRotationListenerApi.GetLogFileRotationListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logFileRotationListenerName** | **string** | Name of the Log File Rotation Listener | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogFileRotationListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddLogFileRotationListener200Response**](AddLogFileRotationListener200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogFileRotationListeners

> LogFileRotationListenerListResponse ListLogFileRotationListeners(ctx).Filter(filter).Execute()

Returns a list of all Log File Rotation Listener objects

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
    resp, r, err := apiClient.LogFileRotationListenerApi.ListLogFileRotationListeners(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFileRotationListenerApi.ListLogFileRotationListeners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogFileRotationListeners`: LogFileRotationListenerListResponse
    fmt.Fprintf(os.Stdout, "Response from `LogFileRotationListenerApi.ListLogFileRotationListeners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogFileRotationListenersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**LogFileRotationListenerListResponse**](LogFileRotationListenerListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogFileRotationListener

> AddLogFileRotationListener200Response UpdateLogFileRotationListener(ctx, logFileRotationListenerName).UpdateRequest(updateRequest).Execute()

Update an existing Log File Rotation Listener by name

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
    logFileRotationListenerName := "logFileRotationListenerName_example" // string | Name of the Log File Rotation Listener
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Log File Rotation Listener

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFileRotationListenerApi.UpdateLogFileRotationListener(context.Background(), logFileRotationListenerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFileRotationListenerApi.UpdateLogFileRotationListener``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogFileRotationListener`: AddLogFileRotationListener200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFileRotationListenerApi.UpdateLogFileRotationListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logFileRotationListenerName** | **string** | Name of the Log File Rotation Listener | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogFileRotationListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Log File Rotation Listener | 

### Return type

[**AddLogFileRotationListener200Response**](AddLogFileRotationListener200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

