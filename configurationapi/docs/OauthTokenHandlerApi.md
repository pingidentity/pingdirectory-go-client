# \OauthTokenHandlerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOauthTokenHandler**](OauthTokenHandlerAPI.md#AddOauthTokenHandler) | **Post** /oauth-token-handlers | Add a new OAuth Token Handler to the config
[**DeleteOauthTokenHandler**](OauthTokenHandlerAPI.md#DeleteOauthTokenHandler) | **Delete** /oauth-token-handlers/{oauth-token-handler-name} | Delete a OAuth Token Handler
[**GetOauthTokenHandler**](OauthTokenHandlerAPI.md#GetOauthTokenHandler) | **Get** /oauth-token-handlers/{oauth-token-handler-name} | Returns a single OAuth Token Handler
[**ListOauthTokenHandlers**](OauthTokenHandlerAPI.md#ListOauthTokenHandlers) | **Get** /oauth-token-handlers | Returns a list of all OAuth Token Handler objects
[**UpdateOauthTokenHandler**](OauthTokenHandlerAPI.md#UpdateOauthTokenHandler) | **Patch** /oauth-token-handlers/{oauth-token-handler-name} | Update an existing OAuth Token Handler by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addOauthTokenHandlerRequest := openapiclient.add_oauth_token_handler_request{AddGroovyScriptedOauthTokenHandlerRequest: openapiclient.NewAddGroovyScriptedOauthTokenHandlerRequest([]openapiclient.EnumgroovyScriptedOauthTokenHandlerSchemaUrn{openapiclient.Enumgroovy-scripted-oauth-token-handlerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:oauth-token-handler:groovy-scripted")}, "ScriptClass_example", "HandlerName_example")} // AddOauthTokenHandlerRequest | Create a new OAuth Token Handler in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenHandlerAPI.AddOauthTokenHandler(context.Background()).AddOauthTokenHandlerRequest(addOauthTokenHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenHandlerAPI.AddOauthTokenHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOauthTokenHandler`: AddOauthTokenHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenHandlerAPI.AddOauthTokenHandler`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    oauthTokenHandlerName := "oauthTokenHandlerName_example" // string | Name of the OAuth Token Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OauthTokenHandlerAPI.DeleteOauthTokenHandler(context.Background(), oauthTokenHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenHandlerAPI.DeleteOauthTokenHandler``: %v\n", err)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    oauthTokenHandlerName := "oauthTokenHandlerName_example" // string | Name of the OAuth Token Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenHandlerAPI.GetOauthTokenHandler(context.Background(), oauthTokenHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenHandlerAPI.GetOauthTokenHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthTokenHandler`: AddOauthTokenHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenHandlerAPI.GetOauthTokenHandler`: %v\n", resp)
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


## ListOauthTokenHandlers

> OauthTokenHandlerListResponse ListOauthTokenHandlers(ctx).Filter(filter).Execute()

Returns a list of all OAuth Token Handler objects

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
    resp, r, err := apiClient.OauthTokenHandlerAPI.ListOauthTokenHandlers(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenHandlerAPI.ListOauthTokenHandlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOauthTokenHandlers`: OauthTokenHandlerListResponse
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenHandlerAPI.ListOauthTokenHandlers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOauthTokenHandlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**OauthTokenHandlerListResponse**](OauthTokenHandlerListResponse.md)

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    oauthTokenHandlerName := "oauthTokenHandlerName_example" // string | Name of the OAuth Token Handler
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing OAuth Token Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthTokenHandlerAPI.UpdateOauthTokenHandler(context.Background(), oauthTokenHandlerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthTokenHandlerAPI.UpdateOauthTokenHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOauthTokenHandler`: AddOauthTokenHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `OauthTokenHandlerAPI.UpdateOauthTokenHandler`: %v\n", resp)
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

