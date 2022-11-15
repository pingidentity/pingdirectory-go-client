# \ExternalServerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExternalServer**](ExternalServerApi.md#AddExternalServer) | **Post** /external-servers | Add a new External Server to the config
[**DeleteExternalServer**](ExternalServerApi.md#DeleteExternalServer) | **Delete** /external-servers/{external-server-name} | Delete a External Server
[**GetExternalServer**](ExternalServerApi.md#GetExternalServer) | **Get** /external-servers/{external-server-name} | Returns a single External Server
[**UpdateExternalServer**](ExternalServerApi.md#UpdateExternalServer) | **Patch** /external-servers/{external-server-name} | Update an existing External Server by name



## AddExternalServer

> AddExternalServer200Response AddExternalServer(ctx).AddExternalServerRequest(addExternalServerRequest).Execute()

Add a new External Server to the config

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
    addExternalServerRequest := openapiclient.add_external_server_request{AddActiveDirectoryExternalServerRequest: openapiclient.NewAddActiveDirectoryExternalServerRequest("ServerName_example", []openapiclient.EnumactiveDirectoryExternalServerSchemaUrn{openapiclient.Enumactive-directory-external-serverSchemaUrn("urn:pingidentity:schemas:configuration:2.0:external-server:active-directory")}, "ServerHostName_example", int32(123), openapiclient.Enumexternal-server-connectionSecurityProp("none"), openapiclient.Enumexternal-server-authenticationMethodProp("none"), openapiclient.Enumexternal-server-verifyCredentialsMethodProp("separate-connections"), "MaxConnectionAge_example", "ConnectTimeout_example", "MaxResponseSize_example")} // AddExternalServerRequest | Create a new External Server in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalServerApi.AddExternalServer(context.Background()).AddExternalServerRequest(addExternalServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalServerApi.AddExternalServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExternalServer`: AddExternalServer200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalServerApi.AddExternalServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddExternalServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addExternalServerRequest** | [**AddExternalServerRequest**](AddExternalServerRequest.md) | Create a new External Server in the config | 

### Return type

[**AddExternalServer200Response**](AddExternalServer200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalServer

> DeleteExternalServer(ctx, externalServerName).Execute()

Delete a External Server

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
    externalServerName := "externalServerName_example" // string | Name of the External Server to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalServerApi.DeleteExternalServer(context.Background(), externalServerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalServerApi.DeleteExternalServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalServerName** | **string** | Name of the External Server to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalServerRequest struct via the builder pattern


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


## GetExternalServer

> AddExternalServer200Response GetExternalServer(ctx, externalServerName).Execute()

Returns a single External Server

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
    externalServerName := "externalServerName_example" // string | Name of the External Server to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalServerApi.GetExternalServer(context.Background(), externalServerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalServerApi.GetExternalServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalServer`: AddExternalServer200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalServerApi.GetExternalServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalServerName** | **string** | Name of the External Server to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddExternalServer200Response**](AddExternalServer200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalServer

> AddExternalServer200Response UpdateExternalServer(ctx, externalServerName).UpdateRequest(updateRequest).Execute()

Update an existing External Server by name

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
    externalServerName := "externalServerName_example" // string | Name of the External Server to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing External Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalServerApi.UpdateExternalServer(context.Background(), externalServerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalServerApi.UpdateExternalServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExternalServer`: AddExternalServer200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalServerApi.UpdateExternalServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalServerName** | **string** | Name of the External Server to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing External Server | 

### Return type

[**AddExternalServer200Response**](AddExternalServer200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

