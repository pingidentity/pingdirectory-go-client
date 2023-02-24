# \LogPublisherApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogPublisher**](LogPublisherApi.md#AddLogPublisher) | **Post** /log-publishers | Add a new Log Publisher to the config
[**DeleteLogPublisher**](LogPublisherApi.md#DeleteLogPublisher) | **Delete** /log-publishers/{log-publisher-name} | Delete a Log Publisher
[**GetLogPublisher**](LogPublisherApi.md#GetLogPublisher) | **Get** /log-publishers/{log-publisher-name} | Returns a single Log Publisher
[**UpdateLogPublisher**](LogPublisherApi.md#UpdateLogPublisher) | **Patch** /log-publishers/{log-publisher-name} | Update an existing Log Publisher by name



## AddLogPublisher

> AddLogPublisher200Response AddLogPublisher(ctx).AddLogPublisherRequest(addLogPublisherRequest).Execute()

Add a new Log Publisher to the config

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
    addLogPublisherRequest := openapiclient.add_log_publisher_request{AddAdminAlertAccessLogPublisherRequest: openapiclient.NewAddAdminAlertAccessLogPublisherRequest("PublisherName_example", []openapiclient.EnumadminAlertAccessLogPublisherSchemaUrn{openapiclient.Enumadmin-alert-access-log-publisherSchemaUrn("urn:pingidentity:schemas:configuration:2.0:log-publisher:admin-alert-access")}, false)} // AddLogPublisherRequest | Create a new Log Publisher in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogPublisherApi.AddLogPublisher(context.Background()).AddLogPublisherRequest(addLogPublisherRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogPublisherApi.AddLogPublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLogPublisher`: AddLogPublisher200Response
    fmt.Fprintf(os.Stdout, "Response from `LogPublisherApi.AddLogPublisher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLogPublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLogPublisherRequest** | [**AddLogPublisherRequest**](AddLogPublisherRequest.md) | Create a new Log Publisher in the config | 

### Return type

[**AddLogPublisher200Response**](AddLogPublisher200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogPublisher

> DeleteLogPublisher(ctx, logPublisherName).Execute()

Delete a Log Publisher

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
    logPublisherName := "logPublisherName_example" // string | Name of the Log Publisher

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogPublisherApi.DeleteLogPublisher(context.Background(), logPublisherName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogPublisherApi.DeleteLogPublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logPublisherName** | **string** | Name of the Log Publisher | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogPublisherRequest struct via the builder pattern


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


## GetLogPublisher

> GetLogPublisher200Response GetLogPublisher(ctx, logPublisherName).Execute()

Returns a single Log Publisher

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
    logPublisherName := "logPublisherName_example" // string | Name of the Log Publisher

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogPublisherApi.GetLogPublisher(context.Background(), logPublisherName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogPublisherApi.GetLogPublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogPublisher`: GetLogPublisher200Response
    fmt.Fprintf(os.Stdout, "Response from `LogPublisherApi.GetLogPublisher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logPublisherName** | **string** | Name of the Log Publisher | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogPublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLogPublisher200Response**](GetLogPublisher200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogPublisher

> GetLogPublisher200Response UpdateLogPublisher(ctx, logPublisherName).UpdateRequest(updateRequest).Execute()

Update an existing Log Publisher by name

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
    logPublisherName := "logPublisherName_example" // string | Name of the Log Publisher
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Log Publisher

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogPublisherApi.UpdateLogPublisher(context.Background(), logPublisherName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogPublisherApi.UpdateLogPublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogPublisher`: GetLogPublisher200Response
    fmt.Fprintf(os.Stdout, "Response from `LogPublisherApi.UpdateLogPublisher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logPublisherName** | **string** | Name of the Log Publisher | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogPublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Log Publisher | 

### Return type

[**GetLogPublisher200Response**](GetLogPublisher200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

