# \LocalDbVlvIndexApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLocalDbVlvIndex**](LocalDbVlvIndexApi.md#AddLocalDbVlvIndex) | **Post** /backends/{backend-name}/local-db-vlv-indexes | Add a new Local DB VLV Index to the config
[**DeleteLocalDbVlvIndex**](LocalDbVlvIndexApi.md#DeleteLocalDbVlvIndex) | **Delete** /backends/{backend-name}/local-db-vlv-indexes/{local-db-vlv-index-name} | Delete a Local DB VLV Index
[**GetLocalDbVlvIndex**](LocalDbVlvIndexApi.md#GetLocalDbVlvIndex) | **Get** /backends/{backend-name}/local-db-vlv-indexes/{local-db-vlv-index-name} | Returns a single Local DB VLV Index
[**ListLocalDbVlvIndexes**](LocalDbVlvIndexApi.md#ListLocalDbVlvIndexes) | **Get** /backends/{backend-name}/local-db-vlv-indexes | Returns a list of all Local DB VLV Index objects
[**UpdateLocalDbVlvIndex**](LocalDbVlvIndexApi.md#UpdateLocalDbVlvIndex) | **Patch** /backends/{backend-name}/local-db-vlv-indexes/{local-db-vlv-index-name} | Update an existing Local DB VLV Index by name



## AddLocalDbVlvIndex

> LocalDbVlvIndexResponse AddLocalDbVlvIndex(ctx, backendName).AddLocalDbVlvIndexRequest(addLocalDbVlvIndexRequest).Execute()

Add a new Local DB VLV Index to the config

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
    addLocalDbVlvIndexRequest := *openapiclient.NewAddLocalDbVlvIndexRequest("IndexName_example", "BaseDN_example", openapiclient.Enumlocal-db-vlv-index-scopeProp("base-object"), "Filter_example", "SortOrder_example", "Name_example") // AddLocalDbVlvIndexRequest | Create a new Local DB VLV Index in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalDbVlvIndexApi.AddLocalDbVlvIndex(context.Background(), backendName).AddLocalDbVlvIndexRequest(addLocalDbVlvIndexRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbVlvIndexApi.AddLocalDbVlvIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLocalDbVlvIndex`: LocalDbVlvIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbVlvIndexApi.AddLocalDbVlvIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLocalDbVlvIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addLocalDbVlvIndexRequest** | [**AddLocalDbVlvIndexRequest**](AddLocalDbVlvIndexRequest.md) | Create a new Local DB VLV Index in the config | 

### Return type

[**LocalDbVlvIndexResponse**](LocalDbVlvIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLocalDbVlvIndex

> DeleteLocalDbVlvIndex(ctx, localDbVlvIndexName, backendName).Execute()

Delete a Local DB VLV Index

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
    localDbVlvIndexName := "localDbVlvIndexName_example" // string | Name of the Local DB VLV Index
    backendName := "backendName_example" // string | Name of the Backend

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LocalDbVlvIndexApi.DeleteLocalDbVlvIndex(context.Background(), localDbVlvIndexName, backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbVlvIndexApi.DeleteLocalDbVlvIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localDbVlvIndexName** | **string** | Name of the Local DB VLV Index | 
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocalDbVlvIndexRequest struct via the builder pattern


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


## GetLocalDbVlvIndex

> LocalDbVlvIndexResponse GetLocalDbVlvIndex(ctx, localDbVlvIndexName, backendName).Execute()

Returns a single Local DB VLV Index

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
    localDbVlvIndexName := "localDbVlvIndexName_example" // string | Name of the Local DB VLV Index
    backendName := "backendName_example" // string | Name of the Backend

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalDbVlvIndexApi.GetLocalDbVlvIndex(context.Background(), localDbVlvIndexName, backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbVlvIndexApi.GetLocalDbVlvIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocalDbVlvIndex`: LocalDbVlvIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbVlvIndexApi.GetLocalDbVlvIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localDbVlvIndexName** | **string** | Name of the Local DB VLV Index | 
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalDbVlvIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LocalDbVlvIndexResponse**](LocalDbVlvIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocalDbVlvIndexes

> LocalDbVlvIndexListResponse ListLocalDbVlvIndexes(ctx, backendName).Filter(filter).Execute()

Returns a list of all Local DB VLV Index objects

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
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalDbVlvIndexApi.ListLocalDbVlvIndexes(context.Background(), backendName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbVlvIndexApi.ListLocalDbVlvIndexes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocalDbVlvIndexes`: LocalDbVlvIndexListResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbVlvIndexApi.ListLocalDbVlvIndexes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocalDbVlvIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**LocalDbVlvIndexListResponse**](LocalDbVlvIndexListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocalDbVlvIndex

> LocalDbVlvIndexResponse UpdateLocalDbVlvIndex(ctx, localDbVlvIndexName, backendName).UpdateRequest(updateRequest).Execute()

Update an existing Local DB VLV Index by name

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
    localDbVlvIndexName := "localDbVlvIndexName_example" // string | Name of the Local DB VLV Index
    backendName := "backendName_example" // string | Name of the Backend
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Local DB VLV Index

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalDbVlvIndexApi.UpdateLocalDbVlvIndex(context.Background(), localDbVlvIndexName, backendName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbVlvIndexApi.UpdateLocalDbVlvIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLocalDbVlvIndex`: LocalDbVlvIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbVlvIndexApi.UpdateLocalDbVlvIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localDbVlvIndexName** | **string** | Name of the Local DB VLV Index | 
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocalDbVlvIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Local DB VLV Index | 

### Return type

[**LocalDbVlvIndexResponse**](LocalDbVlvIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

