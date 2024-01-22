# \JsonFieldConstraintsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddJsonFieldConstraints**](JsonFieldConstraintsAPI.md#AddJsonFieldConstraints) | **Post** /json-attribute-constraints/{json-attribute-constraints-name}/json-field-constraints | Add a new JSON Field Constraints to the config
[**DeleteJsonFieldConstraints**](JsonFieldConstraintsAPI.md#DeleteJsonFieldConstraints) | **Delete** /json-attribute-constraints/{json-attribute-constraints-name}/json-field-constraints/{json-field-constraints-name} | Delete a JSON Field Constraints
[**GetJsonFieldConstraints**](JsonFieldConstraintsAPI.md#GetJsonFieldConstraints) | **Get** /json-attribute-constraints/{json-attribute-constraints-name}/json-field-constraints/{json-field-constraints-name} | Returns a single JSON Field Constraints
[**ListJsonFieldConstraints**](JsonFieldConstraintsAPI.md#ListJsonFieldConstraints) | **Get** /json-attribute-constraints/{json-attribute-constraints-name}/json-field-constraints | Returns a list of all JSON Field Constraints objects
[**UpdateJsonFieldConstraints**](JsonFieldConstraintsAPI.md#UpdateJsonFieldConstraints) | **Patch** /json-attribute-constraints/{json-attribute-constraints-name}/json-field-constraints/{json-field-constraints-name} | Update an existing JSON Field Constraints by name



## AddJsonFieldConstraints

> JsonFieldConstraintsResponse AddJsonFieldConstraints(ctx, jsonAttributeConstraintsName).AddJsonFieldConstraintsRequest(addJsonFieldConstraintsRequest).Execute()

Add a new JSON Field Constraints to the config

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
    addJsonFieldConstraintsRequest := *openapiclient.NewAddJsonFieldConstraintsRequest("JsonField_example", openapiclient.Enumjson-field-constraints-valueTypeProp("any")) // AddJsonFieldConstraintsRequest | Create a new JSON Field Constraints in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JsonFieldConstraintsAPI.AddJsonFieldConstraints(context.Background(), jsonAttributeConstraintsName).AddJsonFieldConstraintsRequest(addJsonFieldConstraintsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonFieldConstraintsAPI.AddJsonFieldConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddJsonFieldConstraints`: JsonFieldConstraintsResponse
    fmt.Fprintf(os.Stdout, "Response from `JsonFieldConstraintsAPI.AddJsonFieldConstraints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jsonAttributeConstraintsName** | **string** | Name of the JSON Attribute Constraints | 

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

> DeleteJsonFieldConstraints(ctx, jsonFieldConstraintsName, jsonAttributeConstraintsName).Execute()

Delete a JSON Field Constraints

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
    jsonFieldConstraintsName := "jsonFieldConstraintsName_example" // string | Name of the JSON Field Constraints
    jsonAttributeConstraintsName := "jsonAttributeConstraintsName_example" // string | Name of the JSON Attribute Constraints

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.JsonFieldConstraintsAPI.DeleteJsonFieldConstraints(context.Background(), jsonFieldConstraintsName, jsonAttributeConstraintsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonFieldConstraintsAPI.DeleteJsonFieldConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jsonFieldConstraintsName** | **string** | Name of the JSON Field Constraints | 
**jsonAttributeConstraintsName** | **string** | Name of the JSON Attribute Constraints | 

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

> JsonFieldConstraintsResponse GetJsonFieldConstraints(ctx, jsonFieldConstraintsName, jsonAttributeConstraintsName).Execute()

Returns a single JSON Field Constraints

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
    jsonFieldConstraintsName := "jsonFieldConstraintsName_example" // string | Name of the JSON Field Constraints
    jsonAttributeConstraintsName := "jsonAttributeConstraintsName_example" // string | Name of the JSON Attribute Constraints

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JsonFieldConstraintsAPI.GetJsonFieldConstraints(context.Background(), jsonFieldConstraintsName, jsonAttributeConstraintsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonFieldConstraintsAPI.GetJsonFieldConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsonFieldConstraints`: JsonFieldConstraintsResponse
    fmt.Fprintf(os.Stdout, "Response from `JsonFieldConstraintsAPI.GetJsonFieldConstraints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jsonFieldConstraintsName** | **string** | Name of the JSON Field Constraints | 
**jsonAttributeConstraintsName** | **string** | Name of the JSON Attribute Constraints | 

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


## ListJsonFieldConstraints

> JsonFieldConstraintsListResponse ListJsonFieldConstraints(ctx, jsonAttributeConstraintsName).Filter(filter).Execute()

Returns a list of all JSON Field Constraints objects

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
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JsonFieldConstraintsAPI.ListJsonFieldConstraints(context.Background(), jsonAttributeConstraintsName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonFieldConstraintsAPI.ListJsonFieldConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJsonFieldConstraints`: JsonFieldConstraintsListResponse
    fmt.Fprintf(os.Stdout, "Response from `JsonFieldConstraintsAPI.ListJsonFieldConstraints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jsonAttributeConstraintsName** | **string** | Name of the JSON Attribute Constraints | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJsonFieldConstraintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**JsonFieldConstraintsListResponse**](JsonFieldConstraintsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJsonFieldConstraints

> JsonFieldConstraintsResponse UpdateJsonFieldConstraints(ctx, jsonFieldConstraintsName, jsonAttributeConstraintsName).UpdateRequest(updateRequest).Execute()

Update an existing JSON Field Constraints by name

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
    jsonFieldConstraintsName := "jsonFieldConstraintsName_example" // string | Name of the JSON Field Constraints
    jsonAttributeConstraintsName := "jsonAttributeConstraintsName_example" // string | Name of the JSON Attribute Constraints
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing JSON Field Constraints

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JsonFieldConstraintsAPI.UpdateJsonFieldConstraints(context.Background(), jsonFieldConstraintsName, jsonAttributeConstraintsName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JsonFieldConstraintsAPI.UpdateJsonFieldConstraints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJsonFieldConstraints`: JsonFieldConstraintsResponse
    fmt.Fprintf(os.Stdout, "Response from `JsonFieldConstraintsAPI.UpdateJsonFieldConstraints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jsonFieldConstraintsName** | **string** | Name of the JSON Field Constraints | 
**jsonAttributeConstraintsName** | **string** | Name of the JSON Attribute Constraints | 

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

