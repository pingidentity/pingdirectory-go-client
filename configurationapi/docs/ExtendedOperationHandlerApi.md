# \ExtendedOperationHandlerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExtendedOperationHandler**](ExtendedOperationHandlerAPI.md#AddExtendedOperationHandler) | **Post** /extended-operation-handlers | Add a new Extended Operation Handler to the config
[**DeleteExtendedOperationHandler**](ExtendedOperationHandlerAPI.md#DeleteExtendedOperationHandler) | **Delete** /extended-operation-handlers/{extended-operation-handler-name} | Delete a Extended Operation Handler
[**GetExtendedOperationHandler**](ExtendedOperationHandlerAPI.md#GetExtendedOperationHandler) | **Get** /extended-operation-handlers/{extended-operation-handler-name} | Returns a single Extended Operation Handler
[**ListExtendedOperationHandlers**](ExtendedOperationHandlerAPI.md#ListExtendedOperationHandlers) | **Get** /extended-operation-handlers | Returns a list of all Extended Operation Handler objects
[**UpdateExtendedOperationHandler**](ExtendedOperationHandlerAPI.md#UpdateExtendedOperationHandler) | **Patch** /extended-operation-handlers/{extended-operation-handler-name} | Update an existing Extended Operation Handler by name



## AddExtendedOperationHandler

> AddExtendedOperationHandler200Response AddExtendedOperationHandler(ctx).AddExtendedOperationHandlerRequest(addExtendedOperationHandlerRequest).Execute()

Add a new Extended Operation Handler to the config

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
    addExtendedOperationHandlerRequest := openapiclient.add_extended_operation_handler_request{AddCollectSupportDataExtendedOperationHandlerRequest: openapiclient.NewAddCollectSupportDataExtendedOperationHandlerRequest([]openapiclient.EnumcollectSupportDataExtendedOperationHandlerSchemaUrn{openapiclient.Enumcollect-support-data-extended-operation-handlerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:collect-support-data")}, false, "HandlerName_example")} // AddExtendedOperationHandlerRequest | Create a new Extended Operation Handler in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtendedOperationHandlerAPI.AddExtendedOperationHandler(context.Background()).AddExtendedOperationHandlerRequest(addExtendedOperationHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtendedOperationHandlerAPI.AddExtendedOperationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExtendedOperationHandler`: AddExtendedOperationHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ExtendedOperationHandlerAPI.AddExtendedOperationHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddExtendedOperationHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addExtendedOperationHandlerRequest** | [**AddExtendedOperationHandlerRequest**](AddExtendedOperationHandlerRequest.md) | Create a new Extended Operation Handler in the config | 

### Return type

[**AddExtendedOperationHandler200Response**](AddExtendedOperationHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExtendedOperationHandler

> DeleteExtendedOperationHandler(ctx, extendedOperationHandlerName).Execute()

Delete a Extended Operation Handler

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
    extendedOperationHandlerName := "extendedOperationHandlerName_example" // string | Name of the Extended Operation Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExtendedOperationHandlerAPI.DeleteExtendedOperationHandler(context.Background(), extendedOperationHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtendedOperationHandlerAPI.DeleteExtendedOperationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extendedOperationHandlerName** | **string** | Name of the Extended Operation Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExtendedOperationHandlerRequest struct via the builder pattern


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


## GetExtendedOperationHandler

> GetExtendedOperationHandler200Response GetExtendedOperationHandler(ctx, extendedOperationHandlerName).Execute()

Returns a single Extended Operation Handler

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
    extendedOperationHandlerName := "extendedOperationHandlerName_example" // string | Name of the Extended Operation Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtendedOperationHandlerAPI.GetExtendedOperationHandler(context.Background(), extendedOperationHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtendedOperationHandlerAPI.GetExtendedOperationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtendedOperationHandler`: GetExtendedOperationHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ExtendedOperationHandlerAPI.GetExtendedOperationHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extendedOperationHandlerName** | **string** | Name of the Extended Operation Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedOperationHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetExtendedOperationHandler200Response**](GetExtendedOperationHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtendedOperationHandlers

> ExtendedOperationHandlerListResponse ListExtendedOperationHandlers(ctx).Filter(filter).Execute()

Returns a list of all Extended Operation Handler objects

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
    resp, r, err := apiClient.ExtendedOperationHandlerAPI.ListExtendedOperationHandlers(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtendedOperationHandlerAPI.ListExtendedOperationHandlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExtendedOperationHandlers`: ExtendedOperationHandlerListResponse
    fmt.Fprintf(os.Stdout, "Response from `ExtendedOperationHandlerAPI.ListExtendedOperationHandlers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExtendedOperationHandlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ExtendedOperationHandlerListResponse**](ExtendedOperationHandlerListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtendedOperationHandler

> GetExtendedOperationHandler200Response UpdateExtendedOperationHandler(ctx, extendedOperationHandlerName).UpdateRequest(updateRequest).Execute()

Update an existing Extended Operation Handler by name

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
    extendedOperationHandlerName := "extendedOperationHandlerName_example" // string | Name of the Extended Operation Handler
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Extended Operation Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtendedOperationHandlerAPI.UpdateExtendedOperationHandler(context.Background(), extendedOperationHandlerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtendedOperationHandlerAPI.UpdateExtendedOperationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExtendedOperationHandler`: GetExtendedOperationHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ExtendedOperationHandlerAPI.UpdateExtendedOperationHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extendedOperationHandlerName** | **string** | Name of the Extended Operation Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtendedOperationHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Extended Operation Handler | 

### Return type

[**GetExtendedOperationHandler200Response**](GetExtendedOperationHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

