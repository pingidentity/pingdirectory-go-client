# \JsonFieldConstraintsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddJsonFieldConstraints**](JsonFieldConstraintsApi.md#AddJsonFieldConstraints) | **Post** /json-field-constraints | Add a new JSON Field Constraints to the config
[**DeleteJsonFieldConstraints**](JsonFieldConstraintsApi.md#DeleteJsonFieldConstraints) | **Delete** /json-field-constraints/{json-field-constraints-name} | Delete a JSON Field Constraints
[**GetJsonFieldConstraints**](JsonFieldConstraintsApi.md#GetJsonFieldConstraints) | **Get** /json-field-constraints/{json-field-constraints-name} | Returns a single JSON Field Constraints
[**UpdateJsonFieldConstraints**](JsonFieldConstraintsApi.md#UpdateJsonFieldConstraints) | **Patch** /json-field-constraints/{json-field-constraints-name} | Update an existing JSON Field Constraints by name



## AddJsonFieldConstraints

> JsonFieldConstraintsResponse AddJsonFieldConstraints(ctx).AddJsonFieldConstraintsRequest(addJsonFieldConstraintsRequest).Execute()

Add a new JSON Field Constraints to the config

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
    addJsonFieldConstraintsRequest := *openapiclient.NewAddJsonFieldConstraintsRequest("JsonField_example", openapiclient.Enumjson-field-constraints-valueTypeProp("any")) // AddJsonFieldConstraintsRequest | Create a new JSON Field Constraints in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JsonFieldConstraintsApi.AddJsonFieldConstraints(context.Background()).AddJsonFieldConstraintsRequest(addJsonFieldConstraintsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonFieldConstraintsApi.AddJsonFieldConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddJsonFieldConstraints`: JsonFieldConstraintsResponse
    fmt.Fprintf(os.Stdout, "Response from `JsonFieldConstraintsApi.AddJsonFieldConstraints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddJsonFieldConstraintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addJsonFieldConstraintsRequest** | [**AddJsonFieldConstraintsRequest**](AddJsonFieldConstraintsRequest.md) | Create a new JSON Field Constraints in the config | 

### Return type

[**JsonFieldConstraintsResponse**](JsonFieldConstraintsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJsonFieldConstraints

> DeleteJsonFieldConstraints(ctx, jsonFieldConstraintsName).Execute()

Delete a JSON Field Constraints

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
    jsonFieldConstraintsName := "jsonFieldConstraintsName_example" // string | Name of the JSON Field Constraints to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JsonFieldConstraintsApi.DeleteJsonFieldConstraints(context.Background(), jsonFieldConstraintsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonFieldConstraintsApi.DeleteJsonFieldConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jsonFieldConstraintsName** | **string** | Name of the JSON Field Constraints to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJsonFieldConstraintsRequest struct via the builder pattern


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


## GetJsonFieldConstraints

> JsonFieldConstraintsResponse GetJsonFieldConstraints(ctx, jsonFieldConstraintsName).Execute()

Returns a single JSON Field Constraints

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
    jsonFieldConstraintsName := "jsonFieldConstraintsName_example" // string | Name of the JSON Field Constraints to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JsonFieldConstraintsApi.GetJsonFieldConstraints(context.Background(), jsonFieldConstraintsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonFieldConstraintsApi.GetJsonFieldConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsonFieldConstraints`: JsonFieldConstraintsResponse
    fmt.Fprintf(os.Stdout, "Response from `JsonFieldConstraintsApi.GetJsonFieldConstraints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jsonFieldConstraintsName** | **string** | Name of the JSON Field Constraints to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsonFieldConstraintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JsonFieldConstraintsResponse**](JsonFieldConstraintsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJsonFieldConstraints

> JsonFieldConstraintsResponse UpdateJsonFieldConstraints(ctx, jsonFieldConstraintsName).UpdateRequest(updateRequest).Execute()

Update an existing JSON Field Constraints by name

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
    jsonFieldConstraintsName := "jsonFieldConstraintsName_example" // string | Name of the JSON Field Constraints to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing JSON Field Constraints

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JsonFieldConstraintsApi.UpdateJsonFieldConstraints(context.Background(), jsonFieldConstraintsName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonFieldConstraintsApi.UpdateJsonFieldConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJsonFieldConstraints`: JsonFieldConstraintsResponse
    fmt.Fprintf(os.Stdout, "Response from `JsonFieldConstraintsApi.UpdateJsonFieldConstraints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jsonFieldConstraintsName** | **string** | Name of the JSON Field Constraints to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJsonFieldConstraintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing JSON Field Constraints | 

### Return type

[**JsonFieldConstraintsResponse**](JsonFieldConstraintsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

