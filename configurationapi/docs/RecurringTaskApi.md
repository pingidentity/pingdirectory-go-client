# \RecurringTaskAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRecurringTask**](RecurringTaskAPI.md#AddRecurringTask) | **Post** /recurring-tasks | Add a new Recurring Task to the config
[**DeleteRecurringTask**](RecurringTaskAPI.md#DeleteRecurringTask) | **Delete** /recurring-tasks/{recurring-task-name} | Delete a Recurring Task
[**GetRecurringTask**](RecurringTaskAPI.md#GetRecurringTask) | **Get** /recurring-tasks/{recurring-task-name} | Returns a single Recurring Task
[**ListRecurringTasks**](RecurringTaskAPI.md#ListRecurringTasks) | **Get** /recurring-tasks | Returns a list of all Recurring Task objects
[**UpdateRecurringTask**](RecurringTaskAPI.md#UpdateRecurringTask) | **Patch** /recurring-tasks/{recurring-task-name} | Update an existing Recurring Task by name



## AddRecurringTask

> AddRecurringTask200Response AddRecurringTask(ctx).AddRecurringTaskRequest(addRecurringTaskRequest).Execute()

Add a new Recurring Task to the config

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
    addRecurringTaskRequest := openapiclient.add_recurring_task_request{AddAuditDataSecurityRecurringTaskRequest: openapiclient.NewAddAuditDataSecurityRecurringTaskRequest([]openapiclient.EnumauditDataSecurityRecurringTaskSchemaUrn{openapiclient.Enumaudit-data-security-recurring-taskSchemaUrn("urn:pingidentity:schemas:configuration:2.0:recurring-task:audit-data-security")}, "TaskName_example")} // AddRecurringTaskRequest | Create a new Recurring Task in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringTaskAPI.AddRecurringTask(context.Background()).AddRecurringTaskRequest(addRecurringTaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskAPI.AddRecurringTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRecurringTask`: AddRecurringTask200Response
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskAPI.AddRecurringTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRecurringTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addRecurringTaskRequest** | [**AddRecurringTaskRequest**](AddRecurringTaskRequest.md) | Create a new Recurring Task in the config | 

### Return type

[**AddRecurringTask200Response**](AddRecurringTask200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecurringTask

> DeleteRecurringTask(ctx, recurringTaskName).Execute()

Delete a Recurring Task

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
    recurringTaskName := "recurringTaskName_example" // string | Name of the Recurring Task

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RecurringTaskAPI.DeleteRecurringTask(context.Background(), recurringTaskName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskAPI.DeleteRecurringTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringTaskName** | **string** | Name of the Recurring Task | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecurringTaskRequest struct via the builder pattern


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


## GetRecurringTask

> AddRecurringTask200Response GetRecurringTask(ctx, recurringTaskName).Execute()

Returns a single Recurring Task

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
    recurringTaskName := "recurringTaskName_example" // string | Name of the Recurring Task

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringTaskAPI.GetRecurringTask(context.Background(), recurringTaskName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskAPI.GetRecurringTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecurringTask`: AddRecurringTask200Response
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskAPI.GetRecurringTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringTaskName** | **string** | Name of the Recurring Task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecurringTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddRecurringTask200Response**](AddRecurringTask200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecurringTasks

> RecurringTaskListResponse ListRecurringTasks(ctx).Filter(filter).Execute()

Returns a list of all Recurring Task objects

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
    resp, r, err := apiClient.RecurringTaskAPI.ListRecurringTasks(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskAPI.ListRecurringTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecurringTasks`: RecurringTaskListResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskAPI.ListRecurringTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRecurringTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**RecurringTaskListResponse**](RecurringTaskListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecurringTask

> AddRecurringTask200Response UpdateRecurringTask(ctx, recurringTaskName).UpdateRequest(updateRequest).Execute()

Update an existing Recurring Task by name

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
    recurringTaskName := "recurringTaskName_example" // string | Name of the Recurring Task
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Recurring Task

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringTaskAPI.UpdateRecurringTask(context.Background(), recurringTaskName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskAPI.UpdateRecurringTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecurringTask`: AddRecurringTask200Response
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskAPI.UpdateRecurringTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringTaskName** | **string** | Name of the Recurring Task | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecurringTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Recurring Task | 

### Return type

[**AddRecurringTask200Response**](AddRecurringTask200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

