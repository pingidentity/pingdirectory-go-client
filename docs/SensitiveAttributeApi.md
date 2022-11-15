# \SensitiveAttributeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSensitiveAttribute**](SensitiveAttributeApi.md#AddSensitiveAttribute) | **Post** /sensitive-attributes | Add a new Sensitive Attribute to the config
[**DeleteSensitiveAttribute**](SensitiveAttributeApi.md#DeleteSensitiveAttribute) | **Delete** /sensitive-attributes/{sensitive-attribute-name} | Delete a Sensitive Attribute
[**GetSensitiveAttribute**](SensitiveAttributeApi.md#GetSensitiveAttribute) | **Get** /sensitive-attributes/{sensitive-attribute-name} | Returns a single Sensitive Attribute
[**UpdateSensitiveAttribute**](SensitiveAttributeApi.md#UpdateSensitiveAttribute) | **Patch** /sensitive-attributes/{sensitive-attribute-name} | Update an existing Sensitive Attribute by name



## AddSensitiveAttribute

> SensitiveAttributeResponse AddSensitiveAttribute(ctx).AddSensitiveAttributeRequest(addSensitiveAttributeRequest).Execute()

Add a new Sensitive Attribute to the config

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
    addSensitiveAttributeRequest := *openapiclient.NewAddSensitiveAttributeRequest("AttributeName_example", []string{"AttributeType_example"}) // AddSensitiveAttributeRequest | Create a new Sensitive Attribute in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SensitiveAttributeApi.AddSensitiveAttribute(context.Background()).AddSensitiveAttributeRequest(addSensitiveAttributeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensitiveAttributeApi.AddSensitiveAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSensitiveAttribute`: SensitiveAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `SensitiveAttributeApi.AddSensitiveAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSensitiveAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addSensitiveAttributeRequest** | [**AddSensitiveAttributeRequest**](AddSensitiveAttributeRequest.md) | Create a new Sensitive Attribute in the config | 

### Return type

[**SensitiveAttributeResponse**](SensitiveAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSensitiveAttribute

> DeleteSensitiveAttribute(ctx, sensitiveAttributeName).Execute()

Delete a Sensitive Attribute

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
    sensitiveAttributeName := "sensitiveAttributeName_example" // string | Name of the Sensitive Attribute to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SensitiveAttributeApi.DeleteSensitiveAttribute(context.Background(), sensitiveAttributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensitiveAttributeApi.DeleteSensitiveAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sensitiveAttributeName** | **string** | Name of the Sensitive Attribute to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSensitiveAttributeRequest struct via the builder pattern


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


## GetSensitiveAttribute

> SensitiveAttributeResponse GetSensitiveAttribute(ctx, sensitiveAttributeName).Execute()

Returns a single Sensitive Attribute

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
    sensitiveAttributeName := "sensitiveAttributeName_example" // string | Name of the Sensitive Attribute to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SensitiveAttributeApi.GetSensitiveAttribute(context.Background(), sensitiveAttributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensitiveAttributeApi.GetSensitiveAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSensitiveAttribute`: SensitiveAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `SensitiveAttributeApi.GetSensitiveAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sensitiveAttributeName** | **string** | Name of the Sensitive Attribute to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSensitiveAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SensitiveAttributeResponse**](SensitiveAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSensitiveAttribute

> SensitiveAttributeResponse UpdateSensitiveAttribute(ctx, sensitiveAttributeName).UpdateRequest(updateRequest).Execute()

Update an existing Sensitive Attribute by name

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
    sensitiveAttributeName := "sensitiveAttributeName_example" // string | Name of the Sensitive Attribute to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Sensitive Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SensitiveAttributeApi.UpdateSensitiveAttribute(context.Background(), sensitiveAttributeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensitiveAttributeApi.UpdateSensitiveAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSensitiveAttribute`: SensitiveAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `SensitiveAttributeApi.UpdateSensitiveAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sensitiveAttributeName** | **string** | Name of the Sensitive Attribute to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSensitiveAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Sensitive Attribute | 

### Return type

[**SensitiveAttributeResponse**](SensitiveAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

