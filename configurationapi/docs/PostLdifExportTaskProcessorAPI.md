# \PostLdifExportTaskProcessorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPostLdifExportTaskProcessor**](PostLdifExportTaskProcessorAPI.md#AddPostLdifExportTaskProcessor) | **Post** /post-ldif-export-task-processors | Add a new Post LDIF Export Task Processor to the config
[**DeletePostLdifExportTaskProcessor**](PostLdifExportTaskProcessorAPI.md#DeletePostLdifExportTaskProcessor) | **Delete** /post-ldif-export-task-processors/{post-ldif-export-task-processor-name} | Delete a Post LDIF Export Task Processor
[**GetPostLdifExportTaskProcessor**](PostLdifExportTaskProcessorAPI.md#GetPostLdifExportTaskProcessor) | **Get** /post-ldif-export-task-processors/{post-ldif-export-task-processor-name} | Returns a single Post LDIF Export Task Processor
[**ListPostLdifExportTaskProcessors**](PostLdifExportTaskProcessorAPI.md#ListPostLdifExportTaskProcessors) | **Get** /post-ldif-export-task-processors | Returns a list of all Post LDIF Export Task Processor objects
[**UpdatePostLdifExportTaskProcessor**](PostLdifExportTaskProcessorAPI.md#UpdatePostLdifExportTaskProcessor) | **Patch** /post-ldif-export-task-processors/{post-ldif-export-task-processor-name} | Update an existing Post LDIF Export Task Processor by name



## AddPostLdifExportTaskProcessor

> AddPostLdifExportTaskProcessor200Response AddPostLdifExportTaskProcessor(ctx).AddPostLdifExportTaskProcessorRequest(addPostLdifExportTaskProcessorRequest).Execute()

Add a new Post LDIF Export Task Processor to the config

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
    addPostLdifExportTaskProcessorRequest := openapiclient.add_post_ldif_export_task_processor_request{AddThirdPartyPostLdifExportTaskProcessorRequest: openapiclient.NewAddThirdPartyPostLdifExportTaskProcessorRequest([]openapiclient.EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn{openapiclient.Enumthird-party-post-ldif-export-task-processorSchemaUrn("urn:pingidentity:schemas:configuration:2.0:post-ldif-export-task-processor:third-party")}, "ExtensionClass_example", false, "ProcessorName_example")} // AddPostLdifExportTaskProcessorRequest | Create a new Post LDIF Export Task Processor in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostLdifExportTaskProcessorAPI.AddPostLdifExportTaskProcessor(context.Background()).AddPostLdifExportTaskProcessorRequest(addPostLdifExportTaskProcessorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostLdifExportTaskProcessorAPI.AddPostLdifExportTaskProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPostLdifExportTaskProcessor`: AddPostLdifExportTaskProcessor200Response
    fmt.Fprintf(os.Stdout, "Response from `PostLdifExportTaskProcessorAPI.AddPostLdifExportTaskProcessor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPostLdifExportTaskProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPostLdifExportTaskProcessorRequest** | [**AddPostLdifExportTaskProcessorRequest**](AddPostLdifExportTaskProcessorRequest.md) | Create a new Post LDIF Export Task Processor in the config | 

### Return type

[**AddPostLdifExportTaskProcessor200Response**](AddPostLdifExportTaskProcessor200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePostLdifExportTaskProcessor

> DeletePostLdifExportTaskProcessor(ctx, postLdifExportTaskProcessorName).Execute()

Delete a Post LDIF Export Task Processor

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
    postLdifExportTaskProcessorName := "postLdifExportTaskProcessorName_example" // string | Name of the Post LDIF Export Task Processor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PostLdifExportTaskProcessorAPI.DeletePostLdifExportTaskProcessor(context.Background(), postLdifExportTaskProcessorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostLdifExportTaskProcessorAPI.DeletePostLdifExportTaskProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postLdifExportTaskProcessorName** | **string** | Name of the Post LDIF Export Task Processor | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostLdifExportTaskProcessorRequest struct via the builder pattern


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


## GetPostLdifExportTaskProcessor

> AddPostLdifExportTaskProcessor200Response GetPostLdifExportTaskProcessor(ctx, postLdifExportTaskProcessorName).Execute()

Returns a single Post LDIF Export Task Processor

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
    postLdifExportTaskProcessorName := "postLdifExportTaskProcessorName_example" // string | Name of the Post LDIF Export Task Processor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostLdifExportTaskProcessorAPI.GetPostLdifExportTaskProcessor(context.Background(), postLdifExportTaskProcessorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostLdifExportTaskProcessorAPI.GetPostLdifExportTaskProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostLdifExportTaskProcessor`: AddPostLdifExportTaskProcessor200Response
    fmt.Fprintf(os.Stdout, "Response from `PostLdifExportTaskProcessorAPI.GetPostLdifExportTaskProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postLdifExportTaskProcessorName** | **string** | Name of the Post LDIF Export Task Processor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostLdifExportTaskProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddPostLdifExportTaskProcessor200Response**](AddPostLdifExportTaskProcessor200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostLdifExportTaskProcessors

> PostLdifExportTaskProcessorListResponse ListPostLdifExportTaskProcessors(ctx).Filter(filter).Execute()

Returns a list of all Post LDIF Export Task Processor objects

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
    resp, r, err := apiClient.PostLdifExportTaskProcessorAPI.ListPostLdifExportTaskProcessors(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostLdifExportTaskProcessorAPI.ListPostLdifExportTaskProcessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPostLdifExportTaskProcessors`: PostLdifExportTaskProcessorListResponse
    fmt.Fprintf(os.Stdout, "Response from `PostLdifExportTaskProcessorAPI.ListPostLdifExportTaskProcessors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPostLdifExportTaskProcessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**PostLdifExportTaskProcessorListResponse**](PostLdifExportTaskProcessorListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePostLdifExportTaskProcessor

> AddPostLdifExportTaskProcessor200Response UpdatePostLdifExportTaskProcessor(ctx, postLdifExportTaskProcessorName).UpdateRequest(updateRequest).Execute()

Update an existing Post LDIF Export Task Processor by name

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
    postLdifExportTaskProcessorName := "postLdifExportTaskProcessorName_example" // string | Name of the Post LDIF Export Task Processor
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Post LDIF Export Task Processor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostLdifExportTaskProcessorAPI.UpdatePostLdifExportTaskProcessor(context.Background(), postLdifExportTaskProcessorName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostLdifExportTaskProcessorAPI.UpdatePostLdifExportTaskProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePostLdifExportTaskProcessor`: AddPostLdifExportTaskProcessor200Response
    fmt.Fprintf(os.Stdout, "Response from `PostLdifExportTaskProcessorAPI.UpdatePostLdifExportTaskProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postLdifExportTaskProcessorName** | **string** | Name of the Post LDIF Export Task Processor | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostLdifExportTaskProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Post LDIF Export Task Processor | 

### Return type

[**AddPostLdifExportTaskProcessor200Response**](AddPostLdifExportTaskProcessor200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

