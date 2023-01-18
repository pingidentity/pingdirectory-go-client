# \RecurringTaskChainApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRecurringTaskChain**](RecurringTaskChainApi.md#AddRecurringTaskChain) | **Post** /recurring-task-chains | Add a new Recurring Task Chain to the config
[**DeleteRecurringTaskChain**](RecurringTaskChainApi.md#DeleteRecurringTaskChain) | **Delete** /recurring-task-chains/{recurring-task-chain-name} | Delete a Recurring Task Chain
[**GetRecurringTaskChain**](RecurringTaskChainApi.md#GetRecurringTaskChain) | **Get** /recurring-task-chains/{recurring-task-chain-name} | Returns a single Recurring Task Chain
[**UpdateRecurringTaskChain**](RecurringTaskChainApi.md#UpdateRecurringTaskChain) | **Patch** /recurring-task-chains/{recurring-task-chain-name} | Update an existing Recurring Task Chain by name



## AddRecurringTaskChain

> RecurringTaskChainResponse AddRecurringTaskChain(ctx).AddRecurringTaskChainRequest(addRecurringTaskChainRequest).Execute()

Add a new Recurring Task Chain to the config

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
    addRecurringTaskChainRequest := *openapiclient.NewAddRecurringTaskChainRequest("ChainName_example", false, []string{"RecurringTask_example"}, []openapiclient.EnumrecurringTaskChainScheduledMonthProp{openapiclient.Enumrecurring-task-chain-scheduledMonthProp("january")}, openapiclient.Enumrecurring-task-chain-scheduledDateSelectionTypeProp("every-day"), []string{"ScheduledTimeOfDay_example"}) // AddRecurringTaskChainRequest | Create a new Recurring Task Chain in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringTaskChainApi.AddRecurringTaskChain(context.Background()).AddRecurringTaskChainRequest(addRecurringTaskChainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskChainApi.AddRecurringTaskChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRecurringTaskChain`: RecurringTaskChainResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskChainApi.AddRecurringTaskChain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRecurringTaskChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addRecurringTaskChainRequest** | [**AddRecurringTaskChainRequest**](AddRecurringTaskChainRequest.md) | Create a new Recurring Task Chain in the config | 

### Return type

[**RecurringTaskChainResponse**](RecurringTaskChainResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecurringTaskChain

> DeleteRecurringTaskChain(ctx, recurringTaskChainName).Execute()

Delete a Recurring Task Chain

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
    recurringTaskChainName := "recurringTaskChainName_example" // string | Name of the Recurring Task Chain to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringTaskChainApi.DeleteRecurringTaskChain(context.Background(), recurringTaskChainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskChainApi.DeleteRecurringTaskChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringTaskChainName** | **string** | Name of the Recurring Task Chain to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecurringTaskChainRequest struct via the builder pattern


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


## GetRecurringTaskChain

> RecurringTaskChainResponse GetRecurringTaskChain(ctx, recurringTaskChainName).Execute()

Returns a single Recurring Task Chain

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
    recurringTaskChainName := "recurringTaskChainName_example" // string | Name of the Recurring Task Chain to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringTaskChainApi.GetRecurringTaskChain(context.Background(), recurringTaskChainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskChainApi.GetRecurringTaskChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecurringTaskChain`: RecurringTaskChainResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskChainApi.GetRecurringTaskChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringTaskChainName** | **string** | Name of the Recurring Task Chain to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecurringTaskChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecurringTaskChainResponse**](RecurringTaskChainResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecurringTaskChain

> RecurringTaskChainResponse UpdateRecurringTaskChain(ctx, recurringTaskChainName).UpdateRequest(updateRequest).Execute()

Update an existing Recurring Task Chain by name

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
    recurringTaskChainName := "recurringTaskChainName_example" // string | Name of the Recurring Task Chain to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Recurring Task Chain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringTaskChainApi.UpdateRecurringTaskChain(context.Background(), recurringTaskChainName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskChainApi.UpdateRecurringTaskChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecurringTaskChain`: RecurringTaskChainResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskChainApi.UpdateRecurringTaskChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringTaskChainName** | **string** | Name of the Recurring Task Chain to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecurringTaskChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Recurring Task Chain | 

### Return type

[**RecurringTaskChainResponse**](RecurringTaskChainResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
