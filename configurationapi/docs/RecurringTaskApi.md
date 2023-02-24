# \RecurringTaskApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRecurringTask**](RecurringTaskApi.md#AddRecurringTask) | **Post** /recurring-tasks | Add a new Recurring Task to the config
[**DeleteRecurringTask**](RecurringTaskApi.md#DeleteRecurringTask) | **Delete** /recurring-tasks/{recurring-task-name} | Delete a Recurring Task
[**GetRecurringTask**](RecurringTaskApi.md#GetRecurringTask) | **Get** /recurring-tasks/{recurring-task-name} | Returns a single Recurring Task
[**UpdateRecurringTask**](RecurringTaskApi.md#UpdateRecurringTask) | **Patch** /recurring-tasks/{recurring-task-name} | Update an existing Recurring Task by name



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
    addRecurringTaskRequest := openapiclient.add_recurring_task_request{AddBackupRecurringTaskRequest: openapiclient.NewAddBackupRecurringTaskRequest("TaskName_example", []openapiclient.EnumbackupRecurringTaskSchemaUrn{openapiclient.Enumbackup-recurring-taskSchemaUrn("urn:pingidentity:schemas:configuration:2.0:recurring-task:backup")})} // AddRecurringTaskRequest | Create a new Recurring Task in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringTaskApi.AddRecurringTask(context.Background()).AddRecurringTaskRequest(addRecurringTaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskApi.AddRecurringTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRecurringTask`: AddRecurringTask200Response
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskApi.AddRecurringTask`: %v\n", resp)
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
    r, err := apiClient.RecurringTaskApi.DeleteRecurringTask(context.Background(), recurringTaskName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskApi.DeleteRecurringTask``: %v\n", err)
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
    resp, r, err := apiClient.RecurringTaskApi.GetRecurringTask(context.Background(), recurringTaskName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskApi.GetRecurringTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecurringTask`: AddRecurringTask200Response
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskApi.GetRecurringTask`: %v\n", resp)
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
    resp, r, err := apiClient.RecurringTaskApi.UpdateRecurringTask(context.Background(), recurringTaskName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskApi.UpdateRecurringTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecurringTask`: AddRecurringTask200Response
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskApi.UpdateRecurringTask`: %v\n", resp)
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

