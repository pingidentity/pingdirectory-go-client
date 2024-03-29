# \JsonAttributeConstraintsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddJsonAttributeConstraints**](JsonAttributeConstraintsAPI.md#AddJsonAttributeConstraints) | **Post** /json-attribute-constraints | Add a new JSON Attribute Constraints to the config
[**DeleteJsonAttributeConstraints**](JsonAttributeConstraintsAPI.md#DeleteJsonAttributeConstraints) | **Delete** /json-attribute-constraints/{json-attribute-constraints-name} | Delete a JSON Attribute Constraints
[**GetJsonAttributeConstraints**](JsonAttributeConstraintsAPI.md#GetJsonAttributeConstraints) | **Get** /json-attribute-constraints/{json-attribute-constraints-name} | Returns a single JSON Attribute Constraints
[**ListJsonAttributeConstraints**](JsonAttributeConstraintsAPI.md#ListJsonAttributeConstraints) | **Get** /json-attribute-constraints | Returns a list of all JSON Attribute Constraints objects
[**UpdateJsonAttributeConstraints**](JsonAttributeConstraintsAPI.md#UpdateJsonAttributeConstraints) | **Patch** /json-attribute-constraints/{json-attribute-constraints-name} | Update an existing JSON Attribute Constraints by name



## AddJsonAttributeConstraints

> JsonAttributeConstraintsResponse AddJsonAttributeConstraints(ctx).AddJsonAttributeConstraintsRequest(addJsonAttributeConstraintsRequest).Execute()

Add a new JSON Attribute Constraints to the config

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
    addJsonAttributeConstraintsRequest := *openapiclient.NewAddJsonAttributeConstraintsRequest("AttributeType_example") // AddJsonAttributeConstraintsRequest | Create a new JSON Attribute Constraints in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JsonAttributeConstraintsAPI.AddJsonAttributeConstraints(context.Background()).AddJsonAttributeConstraintsRequest(addJsonAttributeConstraintsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonAttributeConstraintsAPI.AddJsonAttributeConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddJsonAttributeConstraints`: JsonAttributeConstraintsResponse
    fmt.Fprintf(os.Stdout, "Response from `JsonAttributeConstraintsAPI.AddJsonAttributeConstraints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddJsonAttributeConstraintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addJsonAttributeConstraintsRequest** | [**AddJsonAttributeConstraintsRequest**](AddJsonAttributeConstraintsRequest.md) | Create a new JSON Attribute Constraints in the config | 

### Return type

[**JsonAttributeConstraintsResponse**](JsonAttributeConstraintsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJsonAttributeConstraints

> DeleteJsonAttributeConstraints(ctx, jsonAttributeConstraintsName).Execute()

Delete a JSON Attribute Constraints

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
    jsonAttributeConstraintsName := "jsonAttributeConstraintsName_example" // string | Name of the JSON Attribute Constraints

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.JsonAttributeConstraintsAPI.DeleteJsonAttributeConstraints(context.Background(), jsonAttributeConstraintsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonAttributeConstraintsAPI.DeleteJsonAttributeConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jsonAttributeConstraintsName** | **string** | Name of the JSON Attribute Constraints | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJsonAttributeConstraintsRequest struct via the builder pattern


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


## GetJsonAttributeConstraints

> JsonAttributeConstraintsResponse GetJsonAttributeConstraints(ctx, jsonAttributeConstraintsName).Execute()

Returns a single JSON Attribute Constraints

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
    jsonAttributeConstraintsName := "jsonAttributeConstraintsName_example" // string | Name of the JSON Attribute Constraints

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JsonAttributeConstraintsAPI.GetJsonAttributeConstraints(context.Background(), jsonAttributeConstraintsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonAttributeConstraintsAPI.GetJsonAttributeConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsonAttributeConstraints`: JsonAttributeConstraintsResponse
    fmt.Fprintf(os.Stdout, "Response from `JsonAttributeConstraintsAPI.GetJsonAttributeConstraints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jsonAttributeConstraintsName** | **string** | Name of the JSON Attribute Constraints | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsonAttributeConstraintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JsonAttributeConstraintsResponse**](JsonAttributeConstraintsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJsonAttributeConstraints

> JsonAttributeConstraintsListResponse ListJsonAttributeConstraints(ctx).Filter(filter).Execute()

Returns a list of all JSON Attribute Constraints objects

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
    resp, r, err := apiClient.JsonAttributeConstraintsAPI.ListJsonAttributeConstraints(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonAttributeConstraintsAPI.ListJsonAttributeConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJsonAttributeConstraints`: JsonAttributeConstraintsListResponse
    fmt.Fprintf(os.Stdout, "Response from `JsonAttributeConstraintsAPI.ListJsonAttributeConstraints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListJsonAttributeConstraintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**JsonAttributeConstraintsListResponse**](JsonAttributeConstraintsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJsonAttributeConstraints

> JsonAttributeConstraintsResponse UpdateJsonAttributeConstraints(ctx, jsonAttributeConstraintsName).UpdateRequest(updateRequest).Execute()

Update an existing JSON Attribute Constraints by name

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
    jsonAttributeConstraintsName := "jsonAttributeConstraintsName_example" // string | Name of the JSON Attribute Constraints
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing JSON Attribute Constraints

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JsonAttributeConstraintsAPI.UpdateJsonAttributeConstraints(context.Background(), jsonAttributeConstraintsName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonAttributeConstraintsAPI.UpdateJsonAttributeConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJsonAttributeConstraints`: JsonAttributeConstraintsResponse
    fmt.Fprintf(os.Stdout, "Response from `JsonAttributeConstraintsAPI.UpdateJsonAttributeConstraints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jsonAttributeConstraintsName** | **string** | Name of the JSON Attribute Constraints | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJsonAttributeConstraintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing JSON Attribute Constraints | 

### Return type

[**JsonAttributeConstraintsResponse**](JsonAttributeConstraintsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

