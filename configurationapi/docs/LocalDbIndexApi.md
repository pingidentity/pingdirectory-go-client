# \LocalDbIndexAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLocalDbIndex**](LocalDbIndexAPI.md#AddLocalDbIndex) | **Post** /backends/{backend-name}/local-db-indexes | Add a new Local DB Index to the config
[**DeleteLocalDbIndex**](LocalDbIndexAPI.md#DeleteLocalDbIndex) | **Delete** /backends/{backend-name}/local-db-indexes/{local-db-index-name} | Delete a Local DB Index
[**GetLocalDbIndex**](LocalDbIndexAPI.md#GetLocalDbIndex) | **Get** /backends/{backend-name}/local-db-indexes/{local-db-index-name} | Returns a single Local DB Index
[**ListLocalDbIndexes**](LocalDbIndexAPI.md#ListLocalDbIndexes) | **Get** /backends/{backend-name}/local-db-indexes | Returns a list of all Local DB Index objects
[**UpdateLocalDbIndex**](LocalDbIndexAPI.md#UpdateLocalDbIndex) | **Patch** /backends/{backend-name}/local-db-indexes/{local-db-index-name} | Update an existing Local DB Index by name



## AddLocalDbIndex

> LocalDbIndexResponse AddLocalDbIndex(ctx, backendName).AddLocalDbIndexRequest(addLocalDbIndexRequest).Execute()

Add a new Local DB Index to the config

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
    addLocalDbIndexRequest := *openapiclient.NewAddLocalDbIndexRequest("Attribute_example", []openapiclient.EnumlocalDbIndexIndexTypeProp{openapiclient.Enumlocal-db-index-indexTypeProp("equality")}, "IndexName_example") // AddLocalDbIndexRequest | Create a new Local DB Index in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalDbIndexAPI.AddLocalDbIndex(context.Background(), backendName).AddLocalDbIndexRequest(addLocalDbIndexRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbIndexAPI.AddLocalDbIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLocalDbIndex`: LocalDbIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbIndexAPI.AddLocalDbIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLocalDbIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addLocalDbIndexRequest** | [**AddLocalDbIndexRequest**](AddLocalDbIndexRequest.md) | Create a new Local DB Index in the config | 

### Return type

[**LocalDbIndexResponse**](LocalDbIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLocalDbIndex

> DeleteLocalDbIndex(ctx, localDbIndexName, backendName).Execute()

Delete a Local DB Index

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
    localDbIndexName := "localDbIndexName_example" // string | Name of the Local DB Index
    backendName := "backendName_example" // string | Name of the Backend

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LocalDbIndexAPI.DeleteLocalDbIndex(context.Background(), localDbIndexName, backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbIndexAPI.DeleteLocalDbIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localDbIndexName** | **string** | Name of the Local DB Index | 
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocalDbIndexRequest struct via the builder pattern


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


## GetLocalDbIndex

> LocalDbIndexResponse GetLocalDbIndex(ctx, localDbIndexName, backendName).Execute()

Returns a single Local DB Index

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
    localDbIndexName := "localDbIndexName_example" // string | Name of the Local DB Index
    backendName := "backendName_example" // string | Name of the Backend

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalDbIndexAPI.GetLocalDbIndex(context.Background(), localDbIndexName, backendName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbIndexAPI.GetLocalDbIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocalDbIndex`: LocalDbIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbIndexAPI.GetLocalDbIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localDbIndexName** | **string** | Name of the Local DB Index | 
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalDbIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LocalDbIndexResponse**](LocalDbIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocalDbIndexes

> LocalDbIndexListResponse ListLocalDbIndexes(ctx, backendName).Filter(filter).Execute()

Returns a list of all Local DB Index objects

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
    resp, r, err := apiClient.LocalDbIndexAPI.ListLocalDbIndexes(context.Background(), backendName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbIndexAPI.ListLocalDbIndexes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocalDbIndexes`: LocalDbIndexListResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbIndexAPI.ListLocalDbIndexes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocalDbIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**LocalDbIndexListResponse**](LocalDbIndexListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocalDbIndex

> LocalDbIndexResponse UpdateLocalDbIndex(ctx, localDbIndexName, backendName).UpdateRequest(updateRequest).Execute()

Update an existing Local DB Index by name

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
    localDbIndexName := "localDbIndexName_example" // string | Name of the Local DB Index
    backendName := "backendName_example" // string | Name of the Backend
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Local DB Index

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalDbIndexAPI.UpdateLocalDbIndex(context.Background(), localDbIndexName, backendName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalDbIndexAPI.UpdateLocalDbIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLocalDbIndex`: LocalDbIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `LocalDbIndexAPI.UpdateLocalDbIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localDbIndexName** | **string** | Name of the Local DB Index | 
**backendName** | **string** | Name of the Backend | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocalDbIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Local DB Index | 

### Return type

[**LocalDbIndexResponse**](LocalDbIndexResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

