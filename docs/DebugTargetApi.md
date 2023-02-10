# \DebugTargetApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDebugTarget**](DebugTargetApi.md#AddDebugTarget) | **Post** /debug-targets | Add a new Debug Target to the config
[**DeleteDebugTarget**](DebugTargetApi.md#DeleteDebugTarget) | **Delete** /debug-targets/{debug-target-name} | Delete a Debug Target
[**GetDebugTarget**](DebugTargetApi.md#GetDebugTarget) | **Get** /debug-targets/{debug-target-name} | Returns a single Debug Target
[**UpdateDebugTarget**](DebugTargetApi.md#UpdateDebugTarget) | **Patch** /debug-targets/{debug-target-name} | Update an existing Debug Target by name



## AddDebugTarget

> DebugTargetResponse AddDebugTarget(ctx).AddDebugTargetRequest(addDebugTargetRequest).Execute()

Add a new Debug Target to the config

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
    addDebugTargetRequest := *openapiclient.NewAddDebugTargetRequest("TargetName_example", "DebugScope_example", openapiclient.Enumdebug-target-debugLevelProp("disabled")) // AddDebugTargetRequest | Create a new Debug Target in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DebugTargetApi.AddDebugTarget(context.Background()).AddDebugTargetRequest(addDebugTargetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DebugTargetApi.AddDebugTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDebugTarget`: DebugTargetResponse
    fmt.Fprintf(os.Stdout, "Response from `DebugTargetApi.AddDebugTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDebugTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addDebugTargetRequest** | [**AddDebugTargetRequest**](AddDebugTargetRequest.md) | Create a new Debug Target in the config | 

### Return type

[**DebugTargetResponse**](DebugTargetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDebugTarget

> DeleteDebugTarget(ctx, debugTargetName).Execute()

Delete a Debug Target

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
    debugTargetName := "debugTargetName_example" // string | Name of the Debug Target to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DebugTargetApi.DeleteDebugTarget(context.Background(), debugTargetName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DebugTargetApi.DeleteDebugTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debugTargetName** | **string** | Name of the Debug Target to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDebugTargetRequest struct via the builder pattern


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


## GetDebugTarget

> DebugTargetResponse GetDebugTarget(ctx, debugTargetName).Execute()

Returns a single Debug Target

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
    debugTargetName := "debugTargetName_example" // string | Name of the Debug Target to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DebugTargetApi.GetDebugTarget(context.Background(), debugTargetName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DebugTargetApi.GetDebugTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDebugTarget`: DebugTargetResponse
    fmt.Fprintf(os.Stdout, "Response from `DebugTargetApi.GetDebugTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debugTargetName** | **string** | Name of the Debug Target to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDebugTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DebugTargetResponse**](DebugTargetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDebugTarget

> DebugTargetResponse UpdateDebugTarget(ctx, debugTargetName).UpdateRequest(updateRequest).Execute()

Update an existing Debug Target by name

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
    debugTargetName := "debugTargetName_example" // string | Name of the Debug Target to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Debug Target

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DebugTargetApi.UpdateDebugTarget(context.Background(), debugTargetName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DebugTargetApi.UpdateDebugTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDebugTarget`: DebugTargetResponse
    fmt.Fprintf(os.Stdout, "Response from `DebugTargetApi.UpdateDebugTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**debugTargetName** | **string** | Name of the Debug Target to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDebugTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Debug Target | 

### Return type

[**DebugTargetResponse**](DebugTargetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

