# \SaslMechanismHandlerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSaslMechanismHandler**](SaslMechanismHandlerApi.md#AddSaslMechanismHandler) | **Post** /sasl-mechanism-handlers | Add a new SASL Mechanism Handler to the config
[**DeleteSaslMechanismHandler**](SaslMechanismHandlerApi.md#DeleteSaslMechanismHandler) | **Delete** /sasl-mechanism-handlers/{sasl-mechanism-handler-name} | Delete a SASL Mechanism Handler
[**GetSaslMechanismHandler**](SaslMechanismHandlerApi.md#GetSaslMechanismHandler) | **Get** /sasl-mechanism-handlers/{sasl-mechanism-handler-name} | Returns a single SASL Mechanism Handler
[**UpdateSaslMechanismHandler**](SaslMechanismHandlerApi.md#UpdateSaslMechanismHandler) | **Patch** /sasl-mechanism-handlers/{sasl-mechanism-handler-name} | Update an existing SASL Mechanism Handler by name



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
    openapiclient "./openapi"
)

func main() {
    addSaslMechanismHandlerRequest := openapiclient.add_sasl_mechanism_handler_request{AddOauthBearerSaslMechanismHandlerRequest: openapiclient.NewAddOauthBearerSaslMechanismHandlerRequest("HandlerName_example", []openapiclient.EnumoauthBearerSaslMechanismHandlerSchemaUrn{openapiclient.Enumoauth-bearer-sasl-mechanism-handlerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:sasl-mechanism-handler:oauth-bearer")}, false)} // AddSaslMechanismHandlerRequest | Create a new SASL Mechanism Handler in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaslMechanismHandlerApi.AddSaslMechanismHandler(context.Background()).AddSaslMechanismHandlerRequest(addSaslMechanismHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaslMechanismHandlerApi.AddSaslMechanismHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSaslMechanismHandler`: AddSaslMechanismHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `SaslMechanismHandlerApi.AddSaslMechanismHandler`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    saslMechanismHandlerName := "saslMechanismHandlerName_example" // string | Name of the SASL Mechanism Handler to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaslMechanismHandlerApi.DeleteSaslMechanismHandler(context.Background(), saslMechanismHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaslMechanismHandlerApi.DeleteSaslMechanismHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saslMechanismHandlerName** | **string** | Name of the SASL Mechanism Handler to be deleted | 

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
    openapiclient "./openapi"
)

func main() {
    saslMechanismHandlerName := "saslMechanismHandlerName_example" // string | Name of the SASL Mechanism Handler to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaslMechanismHandlerApi.GetSaslMechanismHandler(context.Background(), saslMechanismHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaslMechanismHandlerApi.GetSaslMechanismHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSaslMechanismHandler`: GetSaslMechanismHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `SaslMechanismHandlerApi.GetSaslMechanismHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saslMechanismHandlerName** | **string** | Name of the SASL Mechanism Handler to be read | 

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
    openapiclient "./openapi"
)

func main() {
    saslMechanismHandlerName := "saslMechanismHandlerName_example" // string | Name of the SASL Mechanism Handler to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing SASL Mechanism Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaslMechanismHandlerApi.UpdateSaslMechanismHandler(context.Background(), saslMechanismHandlerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaslMechanismHandlerApi.UpdateSaslMechanismHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSaslMechanismHandler`: GetSaslMechanismHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `SaslMechanismHandlerApi.UpdateSaslMechanismHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saslMechanismHandlerName** | **string** | Name of the SASL Mechanism Handler to be updated | 

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

