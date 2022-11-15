# \ConstructedAttributeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConstructedAttribute**](ConstructedAttributeApi.md#AddConstructedAttribute) | **Post** /constructed-attributes | Add a new Constructed Attribute to the config
[**DeleteConstructedAttribute**](ConstructedAttributeApi.md#DeleteConstructedAttribute) | **Delete** /constructed-attributes/{constructed-attribute-name} | Delete a Constructed Attribute
[**GetConstructedAttribute**](ConstructedAttributeApi.md#GetConstructedAttribute) | **Get** /constructed-attributes/{constructed-attribute-name} | Returns a single Constructed Attribute
[**UpdateConstructedAttribute**](ConstructedAttributeApi.md#UpdateConstructedAttribute) | **Patch** /constructed-attributes/{constructed-attribute-name} | Update an existing Constructed Attribute by name



## AddConstructedAttribute

> ConstructedAttributeResponse AddConstructedAttribute(ctx).AddConstructedAttributeRequest(addConstructedAttributeRequest).Execute()

Add a new Constructed Attribute to the config

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
    addConstructedAttributeRequest := *openapiclient.NewAddConstructedAttributeRequest("AttributeName_example", "AttributeType_example", []string{"ValuePattern_example"}) // AddConstructedAttributeRequest | Create a new Constructed Attribute in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConstructedAttributeApi.AddConstructedAttribute(context.Background()).AddConstructedAttributeRequest(addConstructedAttributeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConstructedAttributeApi.AddConstructedAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConstructedAttribute`: ConstructedAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ConstructedAttributeApi.AddConstructedAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddConstructedAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addConstructedAttributeRequest** | [**AddConstructedAttributeRequest**](AddConstructedAttributeRequest.md) | Create a new Constructed Attribute in the config | 

### Return type

[**ConstructedAttributeResponse**](ConstructedAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConstructedAttribute

> DeleteConstructedAttribute(ctx, constructedAttributeName).Execute()

Delete a Constructed Attribute

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
    constructedAttributeName := "constructedAttributeName_example" // string | Name of the Constructed Attribute to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConstructedAttributeApi.DeleteConstructedAttribute(context.Background(), constructedAttributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConstructedAttributeApi.DeleteConstructedAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**constructedAttributeName** | **string** | Name of the Constructed Attribute to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConstructedAttributeRequest struct via the builder pattern


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


## GetConstructedAttribute

> ConstructedAttributeResponse GetConstructedAttribute(ctx, constructedAttributeName).Execute()

Returns a single Constructed Attribute

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
    constructedAttributeName := "constructedAttributeName_example" // string | Name of the Constructed Attribute to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConstructedAttributeApi.GetConstructedAttribute(context.Background(), constructedAttributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConstructedAttributeApi.GetConstructedAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConstructedAttribute`: ConstructedAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ConstructedAttributeApi.GetConstructedAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**constructedAttributeName** | **string** | Name of the Constructed Attribute to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConstructedAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConstructedAttributeResponse**](ConstructedAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConstructedAttribute

> ConstructedAttributeResponse UpdateConstructedAttribute(ctx, constructedAttributeName).UpdateRequest(updateRequest).Execute()

Update an existing Constructed Attribute by name

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
    constructedAttributeName := "constructedAttributeName_example" // string | Name of the Constructed Attribute to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Constructed Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConstructedAttributeApi.UpdateConstructedAttribute(context.Background(), constructedAttributeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConstructedAttributeApi.UpdateConstructedAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConstructedAttribute`: ConstructedAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ConstructedAttributeApi.UpdateConstructedAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**constructedAttributeName** | **string** | Name of the Constructed Attribute to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConstructedAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Constructed Attribute | 

### Return type

[**ConstructedAttributeResponse**](ConstructedAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

