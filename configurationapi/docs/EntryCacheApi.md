# \EntryCacheApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEntryCache**](EntryCacheApi.md#AddEntryCache) | **Post** /entry-caches | Add a new Entry Cache to the config
[**DeleteEntryCache**](EntryCacheApi.md#DeleteEntryCache) | **Delete** /entry-caches/{entry-cache-name} | Delete a Entry Cache
[**GetEntryCache**](EntryCacheApi.md#GetEntryCache) | **Get** /entry-caches/{entry-cache-name} | Returns a single Entry Cache
[**UpdateEntryCache**](EntryCacheApi.md#UpdateEntryCache) | **Patch** /entry-caches/{entry-cache-name} | Update an existing Entry Cache by name



## AddEntryCache

> FifoEntryCacheResponse AddEntryCache(ctx).AddFifoEntryCacheRequest(addFifoEntryCacheRequest).Execute()

Add a new Entry Cache to the config

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
    addFifoEntryCacheRequest := *openapiclient.NewAddFifoEntryCacheRequest("CacheName_example", false, int64(123)) // AddFifoEntryCacheRequest | Create a new Entry Cache in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryCacheApi.AddEntryCache(context.Background()).AddFifoEntryCacheRequest(addFifoEntryCacheRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryCacheApi.AddEntryCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEntryCache`: FifoEntryCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `EntryCacheApi.AddEntryCache`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddEntryCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addFifoEntryCacheRequest** | [**AddFifoEntryCacheRequest**](AddFifoEntryCacheRequest.md) | Create a new Entry Cache in the config | 

### Return type

[**FifoEntryCacheResponse**](FifoEntryCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntryCache

> DeleteEntryCache(ctx, entryCacheName).Execute()

Delete a Entry Cache

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
    entryCacheName := "entryCacheName_example" // string | Name of the Entry Cache

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EntryCacheApi.DeleteEntryCache(context.Background(), entryCacheName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryCacheApi.DeleteEntryCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryCacheName** | **string** | Name of the Entry Cache | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntryCacheRequest struct via the builder pattern


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


## GetEntryCache

> FifoEntryCacheResponse GetEntryCache(ctx, entryCacheName).Execute()

Returns a single Entry Cache

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
    entryCacheName := "entryCacheName_example" // string | Name of the Entry Cache

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryCacheApi.GetEntryCache(context.Background(), entryCacheName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryCacheApi.GetEntryCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntryCache`: FifoEntryCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `EntryCacheApi.GetEntryCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryCacheName** | **string** | Name of the Entry Cache | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntryCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FifoEntryCacheResponse**](FifoEntryCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntryCache

> FifoEntryCacheResponse UpdateEntryCache(ctx, entryCacheName).UpdateRequest(updateRequest).Execute()

Update an existing Entry Cache by name

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
    entryCacheName := "entryCacheName_example" // string | Name of the Entry Cache
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Entry Cache

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryCacheApi.UpdateEntryCache(context.Background(), entryCacheName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryCacheApi.UpdateEntryCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntryCache`: FifoEntryCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `EntryCacheApi.UpdateEntryCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryCacheName** | **string** | Name of the Entry Cache | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntryCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Entry Cache | 

### Return type

[**FifoEntryCacheResponse**](FifoEntryCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

