# \PassThroughAuthenticationHandlerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPassThroughAuthenticationHandler**](PassThroughAuthenticationHandlerAPI.md#AddPassThroughAuthenticationHandler) | **Post** /pass-through-authentication-handlers | Add a new Pass Through Authentication Handler to the config
[**DeletePassThroughAuthenticationHandler**](PassThroughAuthenticationHandlerAPI.md#DeletePassThroughAuthenticationHandler) | **Delete** /pass-through-authentication-handlers/{pass-through-authentication-handler-name} | Delete a Pass Through Authentication Handler
[**GetPassThroughAuthenticationHandler**](PassThroughAuthenticationHandlerAPI.md#GetPassThroughAuthenticationHandler) | **Get** /pass-through-authentication-handlers/{pass-through-authentication-handler-name} | Returns a single Pass Through Authentication Handler
[**ListPassThroughAuthenticationHandlers**](PassThroughAuthenticationHandlerAPI.md#ListPassThroughAuthenticationHandlers) | **Get** /pass-through-authentication-handlers | Returns a list of all Pass Through Authentication Handler objects
[**UpdatePassThroughAuthenticationHandler**](PassThroughAuthenticationHandlerAPI.md#UpdatePassThroughAuthenticationHandler) | **Patch** /pass-through-authentication-handlers/{pass-through-authentication-handler-name} | Update an existing Pass Through Authentication Handler by name



## AddPassThroughAuthenticationHandler

> AddPassThroughAuthenticationHandler200Response AddPassThroughAuthenticationHandler(ctx).AddPassThroughAuthenticationHandlerRequest(addPassThroughAuthenticationHandlerRequest).Execute()

Add a new Pass Through Authentication Handler to the config

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
    addPassThroughAuthenticationHandlerRequest := openapiclient.add_pass_through_authentication_handler_request{AddAggregatePassThroughAuthenticationHandlerRequest: openapiclient.NewAddAggregatePassThroughAuthenticationHandlerRequest([]openapiclient.EnumaggregatePassThroughAuthenticationHandlerSchemaUrn{openapiclient.Enumaggregate-pass-through-authentication-handlerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:pass-through-authentication-handler:aggregate")}, []string{"SubordinatePassThroughAuthenticationHandler_example"}, "HandlerName_example")} // AddPassThroughAuthenticationHandlerRequest | Create a new Pass Through Authentication Handler in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PassThroughAuthenticationHandlerAPI.AddPassThroughAuthenticationHandler(context.Background()).AddPassThroughAuthenticationHandlerRequest(addPassThroughAuthenticationHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassThroughAuthenticationHandlerAPI.AddPassThroughAuthenticationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPassThroughAuthenticationHandler`: AddPassThroughAuthenticationHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `PassThroughAuthenticationHandlerAPI.AddPassThroughAuthenticationHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPassThroughAuthenticationHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPassThroughAuthenticationHandlerRequest** | [**AddPassThroughAuthenticationHandlerRequest**](AddPassThroughAuthenticationHandlerRequest.md) | Create a new Pass Through Authentication Handler in the config | 

### Return type

[**AddPassThroughAuthenticationHandler200Response**](AddPassThroughAuthenticationHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePassThroughAuthenticationHandler

> DeletePassThroughAuthenticationHandler(ctx, passThroughAuthenticationHandlerName).Execute()

Delete a Pass Through Authentication Handler

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
    passThroughAuthenticationHandlerName := "passThroughAuthenticationHandlerName_example" // string | Name of the Pass Through Authentication Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PassThroughAuthenticationHandlerAPI.DeletePassThroughAuthenticationHandler(context.Background(), passThroughAuthenticationHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassThroughAuthenticationHandlerAPI.DeletePassThroughAuthenticationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passThroughAuthenticationHandlerName** | **string** | Name of the Pass Through Authentication Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePassThroughAuthenticationHandlerRequest struct via the builder pattern


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


## GetPassThroughAuthenticationHandler

> AddPassThroughAuthenticationHandler200Response GetPassThroughAuthenticationHandler(ctx, passThroughAuthenticationHandlerName).Execute()

Returns a single Pass Through Authentication Handler

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
    passThroughAuthenticationHandlerName := "passThroughAuthenticationHandlerName_example" // string | Name of the Pass Through Authentication Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PassThroughAuthenticationHandlerAPI.GetPassThroughAuthenticationHandler(context.Background(), passThroughAuthenticationHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassThroughAuthenticationHandlerAPI.GetPassThroughAuthenticationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPassThroughAuthenticationHandler`: AddPassThroughAuthenticationHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `PassThroughAuthenticationHandlerAPI.GetPassThroughAuthenticationHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passThroughAuthenticationHandlerName** | **string** | Name of the Pass Through Authentication Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPassThroughAuthenticationHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddPassThroughAuthenticationHandler200Response**](AddPassThroughAuthenticationHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPassThroughAuthenticationHandlers

> PassThroughAuthenticationHandlerListResponse ListPassThroughAuthenticationHandlers(ctx).Filter(filter).Execute()

Returns a list of all Pass Through Authentication Handler objects

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
    resp, r, err := apiClient.PassThroughAuthenticationHandlerAPI.ListPassThroughAuthenticationHandlers(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassThroughAuthenticationHandlerAPI.ListPassThroughAuthenticationHandlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPassThroughAuthenticationHandlers`: PassThroughAuthenticationHandlerListResponse
    fmt.Fprintf(os.Stdout, "Response from `PassThroughAuthenticationHandlerAPI.ListPassThroughAuthenticationHandlers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPassThroughAuthenticationHandlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**PassThroughAuthenticationHandlerListResponse**](PassThroughAuthenticationHandlerListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePassThroughAuthenticationHandler

> AddPassThroughAuthenticationHandler200Response UpdatePassThroughAuthenticationHandler(ctx, passThroughAuthenticationHandlerName).UpdateRequest(updateRequest).Execute()

Update an existing Pass Through Authentication Handler by name

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
    passThroughAuthenticationHandlerName := "passThroughAuthenticationHandlerName_example" // string | Name of the Pass Through Authentication Handler
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Pass Through Authentication Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PassThroughAuthenticationHandlerAPI.UpdatePassThroughAuthenticationHandler(context.Background(), passThroughAuthenticationHandlerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassThroughAuthenticationHandlerAPI.UpdatePassThroughAuthenticationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePassThroughAuthenticationHandler`: AddPassThroughAuthenticationHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `PassThroughAuthenticationHandlerAPI.UpdatePassThroughAuthenticationHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passThroughAuthenticationHandlerName** | **string** | Name of the Pass Through Authentication Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePassThroughAuthenticationHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Pass Through Authentication Handler | 

### Return type

[**AddPassThroughAuthenticationHandler200Response**](AddPassThroughAuthenticationHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

