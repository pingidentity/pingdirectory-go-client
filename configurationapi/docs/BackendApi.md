# \BackendAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBackend**](BackendAPI.md#AddBackend) | **Post** /backends | Add a new Backend to the config
[**DeleteBackend**](BackendAPI.md#DeleteBackend) | **Delete** /backends/{backend-name} | Delete a Backend
[**GetBackend**](BackendAPI.md#GetBackend) | **Get** /backends/{backend-name} | Returns a single Backend
[**ListBackends**](BackendAPI.md#ListBackends) | **Get** /backends | Returns a list of all Backend objects
[**UpdateBackend**](BackendAPI.md#UpdateBackend) | **Patch** /backends/{backend-name} | Update an existing Backend by name



## AddBackend

> AddBackend200Response AddBackend(ctx).AddLocalDbBackendRequest(addLocalDbBackendRequest).Execute()

Add a new Backend to the config

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
    addLocalDbBackendRequest := *openapiclient.NewAddLocalDbBackendRequest([]openapiclient.EnumlocalDbBackendSchemaUrn{openapiclient.Enumlocal-db-backendSchemaUrn("urn:pingidentity:schemas:configuration:2.0:backend:local-db")}, "BackendID_example", false, []string{"BaseDN_example"}, "BackendName_example") // AddLocalDbBackendRequest | Create a new Backend in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackendAPI.AddBackend(context.Background()).AddLocalDbBackendRequest(addLocalDbBackendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackendAPI.AddBackend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddBackend`: AddBackend200Response
    fmt.Fprintf(os.Stdout, "Response from `BackendAPI.AddBackend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLocalDbBackendRequest** | [**AddLocalDbBackendRequest**](AddLocalDbBackendRequest.md) | Create a new Backend in the config | 

### Return type

[**AddBackend200Response**](AddBackend200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackend

> DeleteBackend(ctx, backendName).Execute()

Delete a Backend

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
    backendName := "backendName_example" // string | Name of the Backend

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BackendAPI.DeleteBackend(context.Background(), backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackendAPI.DeleteBackend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackendRequest struct via the builder pattern


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


## GetBackend

> GetBackend200Response GetBackend(ctx, backendName).Execute()

Returns a single Backend

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
    backendName := "backendName_example" // string | Name of the Backend

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackendAPI.GetBackend(context.Background(), backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackendAPI.GetBackend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackend`: GetBackend200Response
    fmt.Fprintf(os.Stdout, "Response from `BackendAPI.GetBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBackend200Response**](GetBackend200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackends

> BackendListResponse ListBackends(ctx).Filter(filter).Execute()

Returns a list of all Backend objects

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
    resp, r, err := apiClient.BackendAPI.ListBackends(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackendAPI.ListBackends``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBackends`: BackendListResponse
    fmt.Fprintf(os.Stdout, "Response from `BackendAPI.ListBackends`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**BackendListResponse**](BackendListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBackend

> GetBackend200Response UpdateBackend(ctx, backendName).UpdateRequest(updateRequest).Execute()

Update an existing Backend by name

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
    backendName := "backendName_example" // string | Name of the Backend
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Backend

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackendAPI.UpdateBackend(context.Background(), backendName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackendAPI.UpdateBackend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBackend`: GetBackend200Response
    fmt.Fprintf(os.Stdout, "Response from `BackendAPI.UpdateBackend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Backend | 

### Return type

[**GetBackend200Response**](GetBackend200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

