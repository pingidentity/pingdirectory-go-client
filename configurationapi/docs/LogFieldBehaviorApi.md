# \LogFieldBehaviorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogFieldBehavior**](LogFieldBehaviorAPI.md#AddLogFieldBehavior) | **Post** /log-field-behaviors | Add a new Log Field Behavior to the config
[**DeleteLogFieldBehavior**](LogFieldBehaviorAPI.md#DeleteLogFieldBehavior) | **Delete** /log-field-behaviors/{log-field-behavior-name} | Delete a Log Field Behavior
[**GetLogFieldBehavior**](LogFieldBehaviorAPI.md#GetLogFieldBehavior) | **Get** /log-field-behaviors/{log-field-behavior-name} | Returns a single Log Field Behavior
[**ListLogFieldBehaviors**](LogFieldBehaviorAPI.md#ListLogFieldBehaviors) | **Get** /log-field-behaviors | Returns a list of all Log Field Behavior objects
[**UpdateLogFieldBehavior**](LogFieldBehaviorAPI.md#UpdateLogFieldBehavior) | **Patch** /log-field-behaviors/{log-field-behavior-name} | Update an existing Log Field Behavior by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addLogFieldBehaviorRequest := openapiclient.add_log_field_behavior_request{AddJsonFormattedAccessLogFieldBehaviorRequest: openapiclient.NewAddJsonFormattedAccessLogFieldBehaviorRequest([]openapiclient.EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn{openapiclient.Enumjson-formatted-access-log-field-behaviorSchemaUrn("urn:pingidentity:schemas:configuration:2.0:log-field-behavior:json-formatted-access")}, "BehaviorName_example")} // AddLogFieldBehaviorRequest | Create a new Log Field Behavior in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldBehaviorAPI.AddLogFieldBehavior(context.Background()).AddLogFieldBehaviorRequest(addLogFieldBehaviorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldBehaviorAPI.AddLogFieldBehavior``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLogFieldBehavior`: AddLogFieldBehavior200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFieldBehaviorAPI.AddLogFieldBehavior`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    logFieldBehaviorName := "logFieldBehaviorName_example" // string | Name of the Log Field Behavior

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LogFieldBehaviorAPI.DeleteLogFieldBehavior(context.Background(), logFieldBehaviorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldBehaviorAPI.DeleteLogFieldBehavior``: %v\n", err)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    logFieldBehaviorName := "logFieldBehaviorName_example" // string | Name of the Log Field Behavior

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldBehaviorAPI.GetLogFieldBehavior(context.Background(), logFieldBehaviorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldBehaviorAPI.GetLogFieldBehavior``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogFieldBehavior`: AddLogFieldBehavior200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFieldBehaviorAPI.GetLogFieldBehavior`: %v\n", resp)
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


## ListLogFieldBehaviors

> LogFieldBehaviorListResponse ListLogFieldBehaviors(ctx).Filter(filter).Execute()

Returns a list of all Log Field Behavior objects

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
    resp, r, err := apiClient.LogFieldBehaviorAPI.ListLogFieldBehaviors(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldBehaviorAPI.ListLogFieldBehaviors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogFieldBehaviors`: LogFieldBehaviorListResponse
    fmt.Fprintf(os.Stdout, "Response from `LogFieldBehaviorAPI.ListLogFieldBehaviors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogFieldBehaviorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**LogFieldBehaviorListResponse**](LogFieldBehaviorListResponse.md)

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    logFieldBehaviorName := "logFieldBehaviorName_example" // string | Name of the Log Field Behavior
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Log Field Behavior

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldBehaviorAPI.UpdateLogFieldBehavior(context.Background(), logFieldBehaviorName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldBehaviorAPI.UpdateLogFieldBehavior``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogFieldBehavior`: AddLogFieldBehavior200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFieldBehaviorAPI.UpdateLogFieldBehavior`: %v\n", resp)
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

