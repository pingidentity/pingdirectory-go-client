# \LocalDbCompositeIndexAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLocalDbCompositeIndex**](LocalDbCompositeIndexAPI.md#AddLocalDbCompositeIndex) | **Post** /backends/{backend-name}/local-db-composite-indexes | Add a new Local DB Composite Index to the config
[**DeleteLocalDbCompositeIndex**](LocalDbCompositeIndexAPI.md#DeleteLocalDbCompositeIndex) | **Delete** /backends/{backend-name}/local-db-composite-indexes/{local-db-composite-index-name} | Delete a Local DB Composite Index
[**GetLocalDbCompositeIndex**](LocalDbCompositeIndexAPI.md#GetLocalDbCompositeIndex) | **Get** /backends/{backend-name}/local-db-composite-indexes/{local-db-composite-index-name} | Returns a single Local DB Composite Index
[**ListLocalDbCompositeIndexes**](LocalDbCompositeIndexAPI.md#ListLocalDbCompositeIndexes) | **Get** /backends/{backend-name}/local-db-composite-indexes | Returns a list of all Local DB Composite Index objects
[**UpdateLocalDbCompositeIndex**](LocalDbCompositeIndexAPI.md#UpdateLocalDbCompositeIndex) | **Patch** /backends/{backend-name}/local-db-composite-indexes/{local-db-composite-index-name} | Update an existing Local DB Composite Index by name



## AddLocalDbCompositeIndex

> LocalDbCompositeIndexResponse AddLocalDbCompositeIndex(ctx, backendName).AddLocalDbCompositeIndexRequest(addLocalDbCompositeIndexRequest).Execute()

Add a new Local DB Composite Index to the config

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
    addLocalDbCompositeIndexRequest := *openapiclient.NewAddLocalDbCompositeIndexRequest("IndexFilterPattern_example", "IndexName_example") // AddLocalDbCompositeIndexRequest | Create a new Local DB Composite Index in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalDbCompositeIndexAPI.AddLocalDbCompositeIndex(context.Background(), backendName).AddLocalDbCompositeIndexRequest(addLocalDbCompositeIndexRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbCompositeIndexAPI.AddLocalDbCompositeIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLocalDbCompositeIndex`: LocalDbCompositeIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbCompositeIndexAPI.AddLocalDbCompositeIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLocalDbCompositeIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addLocalDbCompositeIndexRequest** | [**AddLocalDbCompositeIndexRequest**](AddLocalDbCompositeIndexRequest.md) | Create a new Local DB Composite Index in the config | 

### Return type

[**LocalDbCompositeIndexResponse**](LocalDbCompositeIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLocalDbCompositeIndex

> DeleteLocalDbCompositeIndex(ctx, localDbCompositeIndexName, backendName).Execute()

Delete a Local DB Composite Index

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
    localDbCompositeIndexName := "localDbCompositeIndexName_example" // string | Name of the Local DB Composite Index
    backendName := "backendName_example" // string | Name of the Backend

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LocalDbCompositeIndexAPI.DeleteLocalDbCompositeIndex(context.Background(), localDbCompositeIndexName, backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbCompositeIndexAPI.DeleteLocalDbCompositeIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localDbCompositeIndexName** | **string** | Name of the Local DB Composite Index | 
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocalDbCompositeIndexRequest struct via the builder pattern


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


## GetLocalDbCompositeIndex

> LocalDbCompositeIndexResponse GetLocalDbCompositeIndex(ctx, localDbCompositeIndexName, backendName).Execute()

Returns a single Local DB Composite Index

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
    localDbCompositeIndexName := "localDbCompositeIndexName_example" // string | Name of the Local DB Composite Index
    backendName := "backendName_example" // string | Name of the Backend

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalDbCompositeIndexAPI.GetLocalDbCompositeIndex(context.Background(), localDbCompositeIndexName, backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbCompositeIndexAPI.GetLocalDbCompositeIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocalDbCompositeIndex`: LocalDbCompositeIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbCompositeIndexAPI.GetLocalDbCompositeIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localDbCompositeIndexName** | **string** | Name of the Local DB Composite Index | 
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalDbCompositeIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LocalDbCompositeIndexResponse**](LocalDbCompositeIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocalDbCompositeIndexes

> LocalDbCompositeIndexListResponse ListLocalDbCompositeIndexes(ctx, backendName).Filter(filter).Execute()

Returns a list of all Local DB Composite Index objects

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
    resp, r, err := apiClient.LocalDbCompositeIndexAPI.ListLocalDbCompositeIndexes(context.Background(), backendName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbCompositeIndexAPI.ListLocalDbCompositeIndexes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocalDbCompositeIndexes`: LocalDbCompositeIndexListResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbCompositeIndexAPI.ListLocalDbCompositeIndexes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocalDbCompositeIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**LocalDbCompositeIndexListResponse**](LocalDbCompositeIndexListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocalDbCompositeIndex

> LocalDbCompositeIndexResponse UpdateLocalDbCompositeIndex(ctx, localDbCompositeIndexName, backendName).UpdateRequest(updateRequest).Execute()

Update an existing Local DB Composite Index by name

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
    localDbCompositeIndexName := "localDbCompositeIndexName_example" // string | Name of the Local DB Composite Index
    backendName := "backendName_example" // string | Name of the Backend
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Local DB Composite Index

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalDbCompositeIndexAPI.UpdateLocalDbCompositeIndex(context.Background(), localDbCompositeIndexName, backendName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbCompositeIndexAPI.UpdateLocalDbCompositeIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLocalDbCompositeIndex`: LocalDbCompositeIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbCompositeIndexAPI.UpdateLocalDbCompositeIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localDbCompositeIndexName** | **string** | Name of the Local DB Composite Index | 
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocalDbCompositeIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Local DB Composite Index | 

### Return type

[**LocalDbCompositeIndexResponse**](LocalDbCompositeIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

