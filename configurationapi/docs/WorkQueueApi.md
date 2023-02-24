# \WorkQueueApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkQueue**](WorkQueueApi.md#GetWorkQueue) | **Get** /work-queue | Returns a single Work Queue
[**UpdateWorkQueue**](WorkQueueApi.md#UpdateWorkQueue) | **Patch** /work-queue | Update an existing Work Queue by name



## GetWorkQueue

> HighThroughputWorkQueueResponse GetWorkQueue(ctx).Execute()

Returns a single Work Queue

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkQueueApi.GetWorkQueue(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkQueueApi.GetWorkQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkQueue`: HighThroughputWorkQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkQueueApi.GetWorkQueue`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkQueueRequest struct via the builder pattern


### Return type

[**HighThroughputWorkQueueResponse**](HighThroughputWorkQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkQueue

> HighThroughputWorkQueueResponse UpdateWorkQueue(ctx).UpdateRequest(updateRequest).Execute()

Update an existing Work Queue by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Work Queue

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkQueueApi.UpdateWorkQueue(context.Background()).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkQueueApi.UpdateWorkQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkQueue`: HighThroughputWorkQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkQueueApi.UpdateWorkQueue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Work Queue | 

### Return type

[**HighThroughputWorkQueueResponse**](HighThroughputWorkQueueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

