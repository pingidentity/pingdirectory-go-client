# \GroupImplementationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupImplementation**](GroupImplementationApi.md#GetGroupImplementation) | **Get** /group-implementations/{group-implementation-name} | Returns a single Group Implementation
[**UpdateGroupImplementation**](GroupImplementationApi.md#UpdateGroupImplementation) | **Patch** /group-implementations/{group-implementation-name} | Update an existing Group Implementation by name



## GetGroupImplementation

> GetGroupImplementation200Response GetGroupImplementation(ctx, groupImplementationName).Execute()

Returns a single Group Implementation

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
    groupImplementationName := "groupImplementationName_example" // string | Name of the Group Implementation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupImplementationApi.GetGroupImplementation(context.Background(), groupImplementationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupImplementationApi.GetGroupImplementation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupImplementation`: GetGroupImplementation200Response
    fmt.Fprintf(os.Stdout, "Response from `GroupImplementationApi.GetGroupImplementation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupImplementationName** | **string** | Name of the Group Implementation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupImplementationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetGroupImplementation200Response**](GetGroupImplementation200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupImplementation

> GetGroupImplementation200Response UpdateGroupImplementation(ctx, groupImplementationName).UpdateRequest(updateRequest).Execute()

Update an existing Group Implementation by name

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
    groupImplementationName := "groupImplementationName_example" // string | Name of the Group Implementation
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Group Implementation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupImplementationApi.UpdateGroupImplementation(context.Background(), groupImplementationName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupImplementationApi.UpdateGroupImplementation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupImplementation`: GetGroupImplementation200Response
    fmt.Fprintf(os.Stdout, "Response from `GroupImplementationApi.UpdateGroupImplementation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupImplementationName** | **string** | Name of the Group Implementation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupImplementationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Group Implementation | 

### Return type

[**GetGroupImplementation200Response**](GetGroupImplementation200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

