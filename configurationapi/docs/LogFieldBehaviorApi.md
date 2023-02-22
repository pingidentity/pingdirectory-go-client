# \LogFieldBehaviorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogFieldBehavior**](LogFieldBehaviorApi.md#AddLogFieldBehavior) | **Post** /log-field-behaviors | Add a new Log Field Behavior to the config
[**DeleteLogFieldBehavior**](LogFieldBehaviorApi.md#DeleteLogFieldBehavior) | **Delete** /log-field-behaviors/{log-field-behavior-name} | Delete a Log Field Behavior
[**GetLogFieldBehavior**](LogFieldBehaviorApi.md#GetLogFieldBehavior) | **Get** /log-field-behaviors/{log-field-behavior-name} | Returns a single Log Field Behavior
[**UpdateLogFieldBehavior**](LogFieldBehaviorApi.md#UpdateLogFieldBehavior) | **Patch** /log-field-behaviors/{log-field-behavior-name} | Update an existing Log Field Behavior by name



## AddLogFieldBehavior

> AddLogFieldBehavior200Response AddLogFieldBehavior(ctx).AddLogFieldBehaviorRequest(addLogFieldBehaviorRequest).Execute()

Add a new Log Field Behavior to the config

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
    addLogFieldBehaviorRequest := openapiclient.add_log_field_behavior_request{AddJsonFormattedAccessLogFieldBehaviorRequest: openapiclient.NewAddJsonFormattedAccessLogFieldBehaviorRequest("BehaviorName_example", []openapiclient.EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn{openapiclient.Enumjson-formatted-access-log-field-behaviorSchemaUrn("urn:pingidentity:schemas:configuration:2.0:log-field-behavior:json-formatted-access")})} // AddLogFieldBehaviorRequest | Create a new Log Field Behavior in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldBehaviorApi.AddLogFieldBehavior(context.Background()).AddLogFieldBehaviorRequest(addLogFieldBehaviorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldBehaviorApi.AddLogFieldBehavior``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLogFieldBehavior`: AddLogFieldBehavior200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFieldBehaviorApi.AddLogFieldBehavior`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLogFieldBehaviorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLogFieldBehaviorRequest** | [**AddLogFieldBehaviorRequest**](AddLogFieldBehaviorRequest.md) | Create a new Log Field Behavior in the config | 

### Return type

[**AddLogFieldBehavior200Response**](AddLogFieldBehavior200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogFieldBehavior

> DeleteLogFieldBehavior(ctx, logFieldBehaviorName).Execute()

Delete a Log Field Behavior

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
    logFieldBehaviorName := "logFieldBehaviorName_example" // string | Name of the Log Field Behavior

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldBehaviorApi.DeleteLogFieldBehavior(context.Background(), logFieldBehaviorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldBehaviorApi.DeleteLogFieldBehavior``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logFieldBehaviorName** | **string** | Name of the Log Field Behavior | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogFieldBehaviorRequest struct via the builder pattern


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


## GetLogFieldBehavior

> AddLogFieldBehavior200Response GetLogFieldBehavior(ctx, logFieldBehaviorName).Execute()

Returns a single Log Field Behavior

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
    logFieldBehaviorName := "logFieldBehaviorName_example" // string | Name of the Log Field Behavior

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldBehaviorApi.GetLogFieldBehavior(context.Background(), logFieldBehaviorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldBehaviorApi.GetLogFieldBehavior``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogFieldBehavior`: AddLogFieldBehavior200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFieldBehaviorApi.GetLogFieldBehavior`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logFieldBehaviorName** | **string** | Name of the Log Field Behavior | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogFieldBehaviorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddLogFieldBehavior200Response**](AddLogFieldBehavior200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogFieldBehavior

> AddLogFieldBehavior200Response UpdateLogFieldBehavior(ctx, logFieldBehaviorName).UpdateRequest(updateRequest).Execute()

Update an existing Log Field Behavior by name

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
    logFieldBehaviorName := "logFieldBehaviorName_example" // string | Name of the Log Field Behavior
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Log Field Behavior

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldBehaviorApi.UpdateLogFieldBehavior(context.Background(), logFieldBehaviorName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldBehaviorApi.UpdateLogFieldBehavior``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogFieldBehavior`: AddLogFieldBehavior200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFieldBehaviorApi.UpdateLogFieldBehavior`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logFieldBehaviorName** | **string** | Name of the Log Field Behavior | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogFieldBehaviorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Log Field Behavior | 

### Return type

[**AddLogFieldBehavior200Response**](AddLogFieldBehavior200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

