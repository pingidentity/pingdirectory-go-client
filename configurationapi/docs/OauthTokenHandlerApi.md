# \OauthTokenHandlerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOauthTokenHandler**](OauthTokenHandlerApi.md#AddOauthTokenHandler) | **Post** /oauth-token-handlers | Add a new OAuth Token Handler to the config
[**DeleteOauthTokenHandler**](OauthTokenHandlerApi.md#DeleteOauthTokenHandler) | **Delete** /oauth-token-handlers/{oauth-token-handler-name} | Delete a OAuth Token Handler
[**GetOauthTokenHandler**](OauthTokenHandlerApi.md#GetOauthTokenHandler) | **Get** /oauth-token-handlers/{oauth-token-handler-name} | Returns a single OAuth Token Handler
[**UpdateOauthTokenHandler**](OauthTokenHandlerApi.md#UpdateOauthTokenHandler) | **Patch** /oauth-token-handlers/{oauth-token-handler-name} | Update an existing OAuth Token Handler by name



## AddOauthTokenHandler

> AddOauthTokenHandler200Response AddOauthTokenHandler(ctx).AddOauthTokenHandlerRequest(addOauthTokenHandlerRequest).Execute()

Add a new OAuth Token Handler to the config

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
    addOauthTokenHandlerRequest := openapiclient.add_oauth_token_handler_request{AddGroovyScriptedOauthTokenHandlerRequest: openapiclient.NewAddGroovyScriptedOauthTokenHandlerRequest("HandlerName_example", []openapiclient.EnumgroovyScriptedOauthTokenHandlerSchemaUrn{openapiclient.Enumgroovy-scripted-oauth-token-handlerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:oauth-token-handler:groovy-scripted")}, "ScriptClass_example")} // AddOauthTokenHandlerRequest | Create a new OAuth Token Handler in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenHandlerApi.AddOauthTokenHandler(context.Background()).AddOauthTokenHandlerRequest(addOauthTokenHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenHandlerApi.AddOauthTokenHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOauthTokenHandler`: AddOauthTokenHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenHandlerApi.AddOauthTokenHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOauthTokenHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addOauthTokenHandlerRequest** | [**AddOauthTokenHandlerRequest**](AddOauthTokenHandlerRequest.md) | Create a new OAuth Token Handler in the config | 

### Return type

[**AddOauthTokenHandler200Response**](AddOauthTokenHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOauthTokenHandler

> DeleteOauthTokenHandler(ctx, oauthTokenHandlerName).Execute()

Delete a OAuth Token Handler

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
    oauthTokenHandlerName := "oauthTokenHandlerName_example" // string | Name of the OAuth Token Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenHandlerApi.DeleteOauthTokenHandler(context.Background(), oauthTokenHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenHandlerApi.DeleteOauthTokenHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oauthTokenHandlerName** | **string** | Name of the OAuth Token Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOauthTokenHandlerRequest struct via the builder pattern


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


## GetOauthTokenHandler

> AddOauthTokenHandler200Response GetOauthTokenHandler(ctx, oauthTokenHandlerName).Execute()

Returns a single OAuth Token Handler

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
    oauthTokenHandlerName := "oauthTokenHandlerName_example" // string | Name of the OAuth Token Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenHandlerApi.GetOauthTokenHandler(context.Background(), oauthTokenHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenHandlerApi.GetOauthTokenHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthTokenHandler`: AddOauthTokenHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenHandlerApi.GetOauthTokenHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oauthTokenHandlerName** | **string** | Name of the OAuth Token Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthTokenHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddOauthTokenHandler200Response**](AddOauthTokenHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOauthTokenHandler

> AddOauthTokenHandler200Response UpdateOauthTokenHandler(ctx, oauthTokenHandlerName).UpdateRequest(updateRequest).Execute()

Update an existing OAuth Token Handler by name

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
    oauthTokenHandlerName := "oauthTokenHandlerName_example" // string | Name of the OAuth Token Handler
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing OAuth Token Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenHandlerApi.UpdateOauthTokenHandler(context.Background(), oauthTokenHandlerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenHandlerApi.UpdateOauthTokenHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOauthTokenHandler`: AddOauthTokenHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenHandlerApi.UpdateOauthTokenHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oauthTokenHandlerName** | **string** | Name of the OAuth Token Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOauthTokenHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing OAuth Token Handler | 

### Return type

[**AddOauthTokenHandler200Response**](AddOauthTokenHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

