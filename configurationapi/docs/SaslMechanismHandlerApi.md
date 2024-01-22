# \SaslMechanismHandlerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSaslMechanismHandler**](SaslMechanismHandlerAPI.md#AddSaslMechanismHandler) | **Post** /sasl-mechanism-handlers | Add a new SASL Mechanism Handler to the config
[**DeleteSaslMechanismHandler**](SaslMechanismHandlerAPI.md#DeleteSaslMechanismHandler) | **Delete** /sasl-mechanism-handlers/{sasl-mechanism-handler-name} | Delete a SASL Mechanism Handler
[**GetSaslMechanismHandler**](SaslMechanismHandlerAPI.md#GetSaslMechanismHandler) | **Get** /sasl-mechanism-handlers/{sasl-mechanism-handler-name} | Returns a single SASL Mechanism Handler
[**ListSaslMechanismHandlers**](SaslMechanismHandlerAPI.md#ListSaslMechanismHandlers) | **Get** /sasl-mechanism-handlers | Returns a list of all SASL Mechanism Handler objects
[**UpdateSaslMechanismHandler**](SaslMechanismHandlerAPI.md#UpdateSaslMechanismHandler) | **Patch** /sasl-mechanism-handlers/{sasl-mechanism-handler-name} | Update an existing SASL Mechanism Handler by name



## AddSaslMechanismHandler

> AddSaslMechanismHandler200Response AddSaslMechanismHandler(ctx).AddSaslMechanismHandlerRequest(addSaslMechanismHandlerRequest).Execute()

Add a new SASL Mechanism Handler to the config

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
    addSaslMechanismHandlerRequest := openapiclient.add_sasl_mechanism_handler_request{AddOauthBearerSaslMechanismHandlerRequest: openapiclient.NewAddOauthBearerSaslMechanismHandlerRequest([]openapiclient.EnumoauthBearerSaslMechanismHandlerSchemaUrn{openapiclient.Enumoauth-bearer-sasl-mechanism-handlerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:oauth-bearer")}, false, "HandlerName_example")} // AddSaslMechanismHandlerRequest | Create a new SASL Mechanism Handler in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaslMechanismHandlerAPI.AddSaslMechanismHandler(context.Background()).AddSaslMechanismHandlerRequest(addSaslMechanismHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaslMechanismHandlerAPI.AddSaslMechanismHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSaslMechanismHandler`: AddSaslMechanismHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `SaslMechanismHandlerAPI.AddSaslMechanismHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSaslMechanismHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addSaslMechanismHandlerRequest** | [**AddSaslMechanismHandlerRequest**](AddSaslMechanismHandlerRequest.md) | Create a new SASL Mechanism Handler in the config | 

### Return type

[**AddSaslMechanismHandler200Response**](AddSaslMechanismHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSaslMechanismHandler

> DeleteSaslMechanismHandler(ctx, saslMechanismHandlerName).Execute()

Delete a SASL Mechanism Handler

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
    saslMechanismHandlerName := "saslMechanismHandlerName_example" // string | Name of the SASL Mechanism Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SaslMechanismHandlerAPI.DeleteSaslMechanismHandler(context.Background(), saslMechanismHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaslMechanismHandlerAPI.DeleteSaslMechanismHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saslMechanismHandlerName** | **string** | Name of the SASL Mechanism Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSaslMechanismHandlerRequest struct via the builder pattern


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


## GetSaslMechanismHandler

> GetSaslMechanismHandler200Response GetSaslMechanismHandler(ctx, saslMechanismHandlerName).Execute()

Returns a single SASL Mechanism Handler

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
    saslMechanismHandlerName := "saslMechanismHandlerName_example" // string | Name of the SASL Mechanism Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaslMechanismHandlerAPI.GetSaslMechanismHandler(context.Background(), saslMechanismHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaslMechanismHandlerAPI.GetSaslMechanismHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSaslMechanismHandler`: GetSaslMechanismHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `SaslMechanismHandlerAPI.GetSaslMechanismHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saslMechanismHandlerName** | **string** | Name of the SASL Mechanism Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSaslMechanismHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSaslMechanismHandler200Response**](GetSaslMechanismHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSaslMechanismHandlers

> SaslMechanismHandlerListResponse ListSaslMechanismHandlers(ctx).Filter(filter).Execute()

Returns a list of all SASL Mechanism Handler objects

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
    resp, r, err := apiClient.SaslMechanismHandlerAPI.ListSaslMechanismHandlers(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaslMechanismHandlerAPI.ListSaslMechanismHandlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSaslMechanismHandlers`: SaslMechanismHandlerListResponse
    fmt.Fprintf(os.Stdout, "Response from `SaslMechanismHandlerAPI.ListSaslMechanismHandlers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSaslMechanismHandlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**SaslMechanismHandlerListResponse**](SaslMechanismHandlerListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSaslMechanismHandler

> GetSaslMechanismHandler200Response UpdateSaslMechanismHandler(ctx, saslMechanismHandlerName).UpdateRequest(updateRequest).Execute()

Update an existing SASL Mechanism Handler by name

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
    saslMechanismHandlerName := "saslMechanismHandlerName_example" // string | Name of the SASL Mechanism Handler
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing SASL Mechanism Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaslMechanismHandlerAPI.UpdateSaslMechanismHandler(context.Background(), saslMechanismHandlerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaslMechanismHandlerAPI.UpdateSaslMechanismHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSaslMechanismHandler`: GetSaslMechanismHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `SaslMechanismHandlerAPI.UpdateSaslMechanismHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saslMechanismHandlerName** | **string** | Name of the SASL Mechanism Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSaslMechanismHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing SASL Mechanism Handler | 

### Return type

[**GetSaslMechanismHandler200Response**](GetSaslMechanismHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

