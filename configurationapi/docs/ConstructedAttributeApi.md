# \ConstructedAttributeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConstructedAttribute**](ConstructedAttributeAPI.md#AddConstructedAttribute) | **Post** /constructed-attributes | Add a new Constructed Attribute to the config
[**DeleteConstructedAttribute**](ConstructedAttributeAPI.md#DeleteConstructedAttribute) | **Delete** /constructed-attributes/{constructed-attribute-name} | Delete a Constructed Attribute
[**GetConstructedAttribute**](ConstructedAttributeAPI.md#GetConstructedAttribute) | **Get** /constructed-attributes/{constructed-attribute-name} | Returns a single Constructed Attribute
[**ListConstructedAttributes**](ConstructedAttributeAPI.md#ListConstructedAttributes) | **Get** /constructed-attributes | Returns a list of all Constructed Attribute objects
[**UpdateConstructedAttribute**](ConstructedAttributeAPI.md#UpdateConstructedAttribute) | **Patch** /constructed-attributes/{constructed-attribute-name} | Update an existing Constructed Attribute by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addConstructedAttributeRequest := *openapiclient.NewAddConstructedAttributeRequest("AttributeType_example", []string{"ValuePattern_example"}, "AttributeName_example") // AddConstructedAttributeRequest | Create a new Constructed Attribute in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConstructedAttributeAPI.AddConstructedAttribute(context.Background()).AddConstructedAttributeRequest(addConstructedAttributeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConstructedAttributeAPI.AddConstructedAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConstructedAttribute`: ConstructedAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ConstructedAttributeAPI.AddConstructedAttribute`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    constructedAttributeName := "constructedAttributeName_example" // string | Name of the Constructed Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConstructedAttributeAPI.DeleteConstructedAttribute(context.Background(), constructedAttributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConstructedAttributeAPI.DeleteConstructedAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**constructedAttributeName** | **string** | Name of the Constructed Attribute | 

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    constructedAttributeName := "constructedAttributeName_example" // string | Name of the Constructed Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConstructedAttributeAPI.GetConstructedAttribute(context.Background(), constructedAttributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConstructedAttributeAPI.GetConstructedAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConstructedAttribute`: ConstructedAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ConstructedAttributeAPI.GetConstructedAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**constructedAttributeName** | **string** | Name of the Constructed Attribute | 

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


## ListConstructedAttributes

> ConstructedAttributeListResponse ListConstructedAttributes(ctx).Filter(filter).Execute()

Returns a list of all Constructed Attribute objects

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
    resp, r, err := apiClient.ConstructedAttributeAPI.ListConstructedAttributes(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConstructedAttributeAPI.ListConstructedAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConstructedAttributes`: ConstructedAttributeListResponse
    fmt.Fprintf(os.Stdout, "Response from `ConstructedAttributeAPI.ListConstructedAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConstructedAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ConstructedAttributeListResponse**](ConstructedAttributeListResponse.md)

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    constructedAttributeName := "constructedAttributeName_example" // string | Name of the Constructed Attribute
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Constructed Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConstructedAttributeAPI.UpdateConstructedAttribute(context.Background(), constructedAttributeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConstructedAttributeAPI.UpdateConstructedAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConstructedAttribute`: ConstructedAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ConstructedAttributeAPI.UpdateConstructedAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**constructedAttributeName** | **string** | Name of the Constructed Attribute | 

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

