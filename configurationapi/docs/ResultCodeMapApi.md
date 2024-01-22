# \ResultCodeMapAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResultCodeMap**](ResultCodeMapAPI.md#AddResultCodeMap) | **Post** /result-code-maps | Add a new Result Code Map to the config
[**DeleteResultCodeMap**](ResultCodeMapAPI.md#DeleteResultCodeMap) | **Delete** /result-code-maps/{result-code-map-name} | Delete a Result Code Map
[**GetResultCodeMap**](ResultCodeMapAPI.md#GetResultCodeMap) | **Get** /result-code-maps/{result-code-map-name} | Returns a single Result Code Map
[**ListResultCodeMaps**](ResultCodeMapAPI.md#ListResultCodeMaps) | **Get** /result-code-maps | Returns a list of all Result Code Map objects
[**UpdateResultCodeMap**](ResultCodeMapAPI.md#UpdateResultCodeMap) | **Patch** /result-code-maps/{result-code-map-name} | Update an existing Result Code Map by name



## AddResultCodeMap

> ResultCodeMapResponse AddResultCodeMap(ctx).AddResultCodeMapRequest(addResultCodeMapRequest).Execute()

Add a new Result Code Map to the config

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
    addResultCodeMapRequest := *openapiclient.NewAddResultCodeMapRequest("MapName_example") // AddResultCodeMapRequest | Create a new Result Code Map in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResultCodeMapAPI.AddResultCodeMap(context.Background()).AddResultCodeMapRequest(addResultCodeMapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultCodeMapAPI.AddResultCodeMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddResultCodeMap`: ResultCodeMapResponse
    fmt.Fprintf(os.Stdout, "Response from `ResultCodeMapAPI.AddResultCodeMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddResultCodeMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addResultCodeMapRequest** | [**AddResultCodeMapRequest**](AddResultCodeMapRequest.md) | Create a new Result Code Map in the config | 

### Return type

[**ResultCodeMapResponse**](ResultCodeMapResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResultCodeMap

> DeleteResultCodeMap(ctx, resultCodeMapName).Execute()

Delete a Result Code Map

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
    resultCodeMapName := "resultCodeMapName_example" // string | Name of the Result Code Map

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResultCodeMapAPI.DeleteResultCodeMap(context.Background(), resultCodeMapName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultCodeMapAPI.DeleteResultCodeMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultCodeMapName** | **string** | Name of the Result Code Map | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResultCodeMapRequest struct via the builder pattern


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


## GetResultCodeMap

> ResultCodeMapResponse GetResultCodeMap(ctx, resultCodeMapName).Execute()

Returns a single Result Code Map

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
    resultCodeMapName := "resultCodeMapName_example" // string | Name of the Result Code Map

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResultCodeMapAPI.GetResultCodeMap(context.Background(), resultCodeMapName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultCodeMapAPI.GetResultCodeMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResultCodeMap`: ResultCodeMapResponse
    fmt.Fprintf(os.Stdout, "Response from `ResultCodeMapAPI.GetResultCodeMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultCodeMapName** | **string** | Name of the Result Code Map | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResultCodeMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultCodeMapResponse**](ResultCodeMapResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResultCodeMaps

> ResultCodeMapListResponse ListResultCodeMaps(ctx).Filter(filter).Execute()

Returns a list of all Result Code Map objects

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
    resp, r, err := apiClient.ResultCodeMapAPI.ListResultCodeMaps(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultCodeMapAPI.ListResultCodeMaps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResultCodeMaps`: ResultCodeMapListResponse
    fmt.Fprintf(os.Stdout, "Response from `ResultCodeMapAPI.ListResultCodeMaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResultCodeMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ResultCodeMapListResponse**](ResultCodeMapListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResultCodeMap

> ResultCodeMapResponse UpdateResultCodeMap(ctx, resultCodeMapName).UpdateRequest(updateRequest).Execute()

Update an existing Result Code Map by name

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
    resultCodeMapName := "resultCodeMapName_example" // string | Name of the Result Code Map
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Result Code Map

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResultCodeMapAPI.UpdateResultCodeMap(context.Background(), resultCodeMapName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultCodeMapAPI.UpdateResultCodeMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResultCodeMap`: ResultCodeMapResponse
    fmt.Fprintf(os.Stdout, "Response from `ResultCodeMapAPI.UpdateResultCodeMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultCodeMapName** | **string** | Name of the Result Code Map | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResultCodeMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Result Code Map | 

### Return type

[**ResultCodeMapResponse**](ResultCodeMapResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

