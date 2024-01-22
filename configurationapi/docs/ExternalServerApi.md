# \ExternalServerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExternalServer**](ExternalServerAPI.md#AddExternalServer) | **Post** /external-servers | Add a new External Server to the config
[**DeleteExternalServer**](ExternalServerAPI.md#DeleteExternalServer) | **Delete** /external-servers/{external-server-name} | Delete a External Server
[**GetExternalServer**](ExternalServerAPI.md#GetExternalServer) | **Get** /external-servers/{external-server-name} | Returns a single External Server
[**ListExternalServers**](ExternalServerAPI.md#ListExternalServers) | **Get** /external-servers | Returns a list of all External Server objects
[**UpdateExternalServer**](ExternalServerAPI.md#UpdateExternalServer) | **Patch** /external-servers/{external-server-name} | Update an existing External Server by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addExternalServerRequest := openapiclient.add_external_server_request{AddActiveDirectoryExternalServerRequest: openapiclient.NewAddActiveDirectoryExternalServerRequest([]openapiclient.EnumactiveDirectoryExternalServerSchemaUrn{openapiclient.Enumactive-directory-external-serverSchemaUrn("urn:pingidentity:schemas:configuration:2.0:external-server:active-directory")}, "ServerHostName_example", "ServerName_example")} // AddExternalServerRequest | Create a new External Server in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalServerAPI.AddExternalServer(context.Background()).AddExternalServerRequest(addExternalServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalServerAPI.AddExternalServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExternalServer`: AddExternalServer200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalServerAPI.AddExternalServer`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    externalServerName := "externalServerName_example" // string | Name of the External Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExternalServerAPI.DeleteExternalServer(context.Background(), externalServerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalServerAPI.DeleteExternalServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalServerName** | **string** | Name of the External Server | 

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    externalServerName := "externalServerName_example" // string | Name of the External Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalServerAPI.GetExternalServer(context.Background(), externalServerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalServerAPI.GetExternalServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalServer`: AddExternalServer200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalServerAPI.GetExternalServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalServerName** | **string** | Name of the External Server | 

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


## ListExternalServers

> ExternalServerListResponse ListExternalServers(ctx).Filter(filter).Execute()

Returns a list of all External Server objects

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
    resp, r, err := apiClient.ExternalServerAPI.ListExternalServers(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalServerAPI.ListExternalServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExternalServers`: ExternalServerListResponse
    fmt.Fprintf(os.Stdout, "Response from `ExternalServerAPI.ListExternalServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExternalServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ExternalServerListResponse**](ExternalServerListResponse.md)

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    externalServerName := "externalServerName_example" // string | Name of the External Server
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing External Server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalServerAPI.UpdateExternalServer(context.Background(), externalServerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalServerAPI.UpdateExternalServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExternalServer`: AddExternalServer200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalServerAPI.UpdateExternalServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalServerName** | **string** | Name of the External Server | 

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

