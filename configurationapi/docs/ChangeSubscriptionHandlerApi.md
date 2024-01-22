# \ChangeSubscriptionHandlerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddChangeSubscriptionHandler**](ChangeSubscriptionHandlerAPI.md#AddChangeSubscriptionHandler) | **Post** /change-subscription-handlers | Add a new Change Subscription Handler to the config
[**DeleteChangeSubscriptionHandler**](ChangeSubscriptionHandlerAPI.md#DeleteChangeSubscriptionHandler) | **Delete** /change-subscription-handlers/{change-subscription-handler-name} | Delete a Change Subscription Handler
[**GetChangeSubscriptionHandler**](ChangeSubscriptionHandlerAPI.md#GetChangeSubscriptionHandler) | **Get** /change-subscription-handlers/{change-subscription-handler-name} | Returns a single Change Subscription Handler
[**ListChangeSubscriptionHandlers**](ChangeSubscriptionHandlerAPI.md#ListChangeSubscriptionHandlers) | **Get** /change-subscription-handlers | Returns a list of all Change Subscription Handler objects
[**UpdateChangeSubscriptionHandler**](ChangeSubscriptionHandlerAPI.md#UpdateChangeSubscriptionHandler) | **Patch** /change-subscription-handlers/{change-subscription-handler-name} | Update an existing Change Subscription Handler by name



## AddChangeSubscriptionHandler

> AddChangeSubscriptionHandler200Response AddChangeSubscriptionHandler(ctx).AddChangeSubscriptionHandlerRequest(addChangeSubscriptionHandlerRequest).Execute()

Add a new Change Subscription Handler to the config

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
    addChangeSubscriptionHandlerRequest := openapiclient.add_change_subscription_handler_request{AddGroovyScriptedChangeSubscriptionHandlerRequest: openapiclient.NewAddGroovyScriptedChangeSubscriptionHandlerRequest([]openapiclient.EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn{openapiclient.Enumgroovy-scripted-change-subscription-handlerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:change-subscription-handler:groovy-scripted")}, "ScriptClass_example", false, "HandlerName_example")} // AddChangeSubscriptionHandlerRequest | Create a new Change Subscription Handler in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChangeSubscriptionHandlerAPI.AddChangeSubscriptionHandler(context.Background()).AddChangeSubscriptionHandlerRequest(addChangeSubscriptionHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChangeSubscriptionHandlerAPI.AddChangeSubscriptionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddChangeSubscriptionHandler`: AddChangeSubscriptionHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ChangeSubscriptionHandlerAPI.AddChangeSubscriptionHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddChangeSubscriptionHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addChangeSubscriptionHandlerRequest** | [**AddChangeSubscriptionHandlerRequest**](AddChangeSubscriptionHandlerRequest.md) | Create a new Change Subscription Handler in the config | 

### Return type

[**AddChangeSubscriptionHandler200Response**](AddChangeSubscriptionHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChangeSubscriptionHandler

> DeleteChangeSubscriptionHandler(ctx, changeSubscriptionHandlerName).Execute()

Delete a Change Subscription Handler

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
    changeSubscriptionHandlerName := "changeSubscriptionHandlerName_example" // string | Name of the Change Subscription Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ChangeSubscriptionHandlerAPI.DeleteChangeSubscriptionHandler(context.Background(), changeSubscriptionHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChangeSubscriptionHandlerAPI.DeleteChangeSubscriptionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeSubscriptionHandlerName** | **string** | Name of the Change Subscription Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChangeSubscriptionHandlerRequest struct via the builder pattern


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


## GetChangeSubscriptionHandler

> AddChangeSubscriptionHandler200Response GetChangeSubscriptionHandler(ctx, changeSubscriptionHandlerName).Execute()

Returns a single Change Subscription Handler

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
    changeSubscriptionHandlerName := "changeSubscriptionHandlerName_example" // string | Name of the Change Subscription Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChangeSubscriptionHandlerAPI.GetChangeSubscriptionHandler(context.Background(), changeSubscriptionHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChangeSubscriptionHandlerAPI.GetChangeSubscriptionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChangeSubscriptionHandler`: AddChangeSubscriptionHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ChangeSubscriptionHandlerAPI.GetChangeSubscriptionHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeSubscriptionHandlerName** | **string** | Name of the Change Subscription Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeSubscriptionHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddChangeSubscriptionHandler200Response**](AddChangeSubscriptionHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChangeSubscriptionHandlers

> ChangeSubscriptionHandlerListResponse ListChangeSubscriptionHandlers(ctx).Filter(filter).Execute()

Returns a list of all Change Subscription Handler objects

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
    resp, r, err := apiClient.ChangeSubscriptionHandlerAPI.ListChangeSubscriptionHandlers(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChangeSubscriptionHandlerAPI.ListChangeSubscriptionHandlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChangeSubscriptionHandlers`: ChangeSubscriptionHandlerListResponse
    fmt.Fprintf(os.Stdout, "Response from `ChangeSubscriptionHandlerAPI.ListChangeSubscriptionHandlers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListChangeSubscriptionHandlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ChangeSubscriptionHandlerListResponse**](ChangeSubscriptionHandlerListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChangeSubscriptionHandler

> AddChangeSubscriptionHandler200Response UpdateChangeSubscriptionHandler(ctx, changeSubscriptionHandlerName).UpdateRequest(updateRequest).Execute()

Update an existing Change Subscription Handler by name

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
    changeSubscriptionHandlerName := "changeSubscriptionHandlerName_example" // string | Name of the Change Subscription Handler
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Change Subscription Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChangeSubscriptionHandlerAPI.UpdateChangeSubscriptionHandler(context.Background(), changeSubscriptionHandlerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChangeSubscriptionHandlerAPI.UpdateChangeSubscriptionHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChangeSubscriptionHandler`: AddChangeSubscriptionHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `ChangeSubscriptionHandlerAPI.UpdateChangeSubscriptionHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeSubscriptionHandlerName** | **string** | Name of the Change Subscription Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChangeSubscriptionHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Change Subscription Handler | 

### Return type

[**AddChangeSubscriptionHandler200Response**](AddChangeSubscriptionHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

