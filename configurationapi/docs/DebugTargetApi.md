# \DebugTargetApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDebugTarget**](DebugTargetApi.md#AddDebugTarget) | **Post** /log-publishers/{log-publisher-name}/debug-targets | Add a new Debug Target to the config
[**DeleteDebugTarget**](DebugTargetApi.md#DeleteDebugTarget) | **Delete** /log-publishers/{log-publisher-name}/debug-targets/{debug-target-name} | Delete a Debug Target
[**GetDebugTarget**](DebugTargetApi.md#GetDebugTarget) | **Get** /log-publishers/{log-publisher-name}/debug-targets/{debug-target-name} | Returns a single Debug Target
[**ListDebugTargets**](DebugTargetApi.md#ListDebugTargets) | **Get** /log-publishers/{log-publisher-name}/debug-targets | Returns a list of all Debug Target objects
[**UpdateDebugTarget**](DebugTargetApi.md#UpdateDebugTarget) | **Patch** /log-publishers/{log-publisher-name}/debug-targets/{debug-target-name} | Update an existing Debug Target by name



## AddDebugTarget

> DebugTargetResponse AddDebugTarget(ctx, logPublisherName).AddDebugTargetRequest(addDebugTargetRequest).Execute()

Add a new Debug Target to the config

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
    logPublisherName := "logPublisherName_example" // string | Name of the Log Publisher
    addDebugTargetRequest := *openapiclient.NewAddDebugTargetRequest("TargetName_example", "DebugScope_example", openapiclient.Enumdebug-target-debugLevelProp("disabled")) // AddDebugTargetRequest | Create a new Debug Target in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DebugTargetApi.AddDebugTarget(context.Background(), logPublisherName).AddDebugTargetRequest(addDebugTargetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DebugTargetApi.AddDebugTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDebugTarget`: DebugTargetResponse
    fmt.Fprintf(os.Stdout, "Response from `DebugTargetApi.AddDebugTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logPublisherName** | **string** | Name of the Log Publisher | 

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

> DeleteDebugTarget(ctx, debugTargetName, logPublisherName).Execute()

Delete a Debug Target

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
    debugTargetName := "debugTargetName_example" // string | Name of the Debug Target
    logPublisherName := "logPublisherName_example" // string | Name of the Log Publisher

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DebugTargetApi.DeleteDebugTarget(context.Background(), debugTargetName, logPublisherName).Execute()
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
**debugTargetName** | **string** | Name of the Debug Target | 
**logPublisherName** | **string** | Name of the Log Publisher | 

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

> DebugTargetResponse GetDebugTarget(ctx, debugTargetName, logPublisherName).Execute()

Returns a single Debug Target

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
    debugTargetName := "debugTargetName_example" // string | Name of the Debug Target
    logPublisherName := "logPublisherName_example" // string | Name of the Log Publisher

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DebugTargetApi.GetDebugTarget(context.Background(), debugTargetName, logPublisherName).Execute()
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
**debugTargetName** | **string** | Name of the Debug Target | 
**logPublisherName** | **string** | Name of the Log Publisher | 

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


## ListDebugTargets

> DebugTargetListResponse ListDebugTargets(ctx, logPublisherName).Filter(filter).Execute()

Returns a list of all Debug Target objects

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
    logPublisherName := "logPublisherName_example" // string | Name of the Log Publisher
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DebugTargetApi.ListDebugTargets(context.Background(), logPublisherName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DebugTargetApi.ListDebugTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDebugTargets`: DebugTargetListResponse
    fmt.Fprintf(os.Stdout, "Response from `DebugTargetApi.ListDebugTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logPublisherName** | **string** | Name of the Log Publisher | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDebugTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**DebugTargetListResponse**](DebugTargetListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDebugTarget

> DebugTargetResponse UpdateDebugTarget(ctx, debugTargetName, logPublisherName).UpdateRequest(updateRequest).Execute()

Update an existing Debug Target by name

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
    debugTargetName := "debugTargetName_example" // string | Name of the Debug Target
    logPublisherName := "logPublisherName_example" // string | Name of the Log Publisher
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Debug Target

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DebugTargetApi.UpdateDebugTarget(context.Background(), debugTargetName, logPublisherName).UpdateRequest(updateRequest).Execute()
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
**debugTargetName** | **string** | Name of the Debug Target | 
**logPublisherName** | **string** | Name of the Log Publisher | 

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

