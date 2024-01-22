# \RecurringTaskChainAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRecurringTaskChain**](RecurringTaskChainAPI.md#AddRecurringTaskChain) | **Post** /recurring-task-chains | Add a new Recurring Task Chain to the config
[**DeleteRecurringTaskChain**](RecurringTaskChainAPI.md#DeleteRecurringTaskChain) | **Delete** /recurring-task-chains/{recurring-task-chain-name} | Delete a Recurring Task Chain
[**GetRecurringTaskChain**](RecurringTaskChainAPI.md#GetRecurringTaskChain) | **Get** /recurring-task-chains/{recurring-task-chain-name} | Returns a single Recurring Task Chain
[**ListRecurringTaskChains**](RecurringTaskChainAPI.md#ListRecurringTaskChains) | **Get** /recurring-task-chains | Returns a list of all Recurring Task Chain objects
[**UpdateRecurringTaskChain**](RecurringTaskChainAPI.md#UpdateRecurringTaskChain) | **Patch** /recurring-task-chains/{recurring-task-chain-name} | Update an existing Recurring Task Chain by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addRecurringTaskChainRequest := *openapiclient.NewAddRecurringTaskChainRequest([]string{"RecurringTask_example"}, openapiclient.Enumrecurring-task-chain-scheduledDateSelectionTypeProp("every-day"), []string{"ScheduledTimeOfDay_example"}, "ChainName_example") // AddRecurringTaskChainRequest | Create a new Recurring Task Chain in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringTaskChainAPI.AddRecurringTaskChain(context.Background()).AddRecurringTaskChainRequest(addRecurringTaskChainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskChainAPI.AddRecurringTaskChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRecurringTaskChain`: RecurringTaskChainResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskChainAPI.AddRecurringTaskChain`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    recurringTaskChainName := "recurringTaskChainName_example" // string | Name of the Recurring Task Chain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RecurringTaskChainAPI.DeleteRecurringTaskChain(context.Background(), recurringTaskChainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskChainAPI.DeleteRecurringTaskChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringTaskChainName** | **string** | Name of the Recurring Task Chain | 

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    recurringTaskChainName := "recurringTaskChainName_example" // string | Name of the Recurring Task Chain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringTaskChainAPI.GetRecurringTaskChain(context.Background(), recurringTaskChainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskChainAPI.GetRecurringTaskChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecurringTaskChain`: RecurringTaskChainResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskChainAPI.GetRecurringTaskChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringTaskChainName** | **string** | Name of the Recurring Task Chain | 

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


## ListRecurringTaskChains

> RecurringTaskChainListResponse ListRecurringTaskChains(ctx).Filter(filter).Execute()

Returns a list of all Recurring Task Chain objects

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
    resp, r, err := apiClient.RecurringTaskChainAPI.ListRecurringTaskChains(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskChainAPI.ListRecurringTaskChains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecurringTaskChains`: RecurringTaskChainListResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskChainAPI.ListRecurringTaskChains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRecurringTaskChainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**RecurringTaskChainListResponse**](RecurringTaskChainListResponse.md)

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    recurringTaskChainName := "recurringTaskChainName_example" // string | Name of the Recurring Task Chain
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Recurring Task Chain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecurringTaskChainAPI.UpdateRecurringTaskChain(context.Background(), recurringTaskChainName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurringTaskChainAPI.UpdateRecurringTaskChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecurringTaskChain`: RecurringTaskChainResponse
    fmt.Fprintf(os.Stdout, "Response from `RecurringTaskChainAPI.UpdateRecurringTaskChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recurringTaskChainName** | **string** | Name of the Recurring Task Chain | 

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

