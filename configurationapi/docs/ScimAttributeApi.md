# \ScimAttributeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScimAttribute**](ScimAttributeApi.md#AddScimAttribute) | **Post** /scim-schemas/{scim-schema-name}/scim-attributes | Add a new SCIM Attribute to the config
[**DeleteScimAttribute**](ScimAttributeApi.md#DeleteScimAttribute) | **Delete** /scim-schemas/{scim-schema-name}/scim-attributes/{scim-attribute-name} | Delete a SCIM Attribute
[**GetScimAttribute**](ScimAttributeApi.md#GetScimAttribute) | **Get** /scim-schemas/{scim-schema-name}/scim-attributes/{scim-attribute-name} | Returns a single SCIM Attribute
[**UpdateScimAttribute**](ScimAttributeApi.md#UpdateScimAttribute) | **Patch** /scim-schemas/{scim-schema-name}/scim-attributes/{scim-attribute-name} | Update an existing SCIM Attribute by name



## AddScimAttribute

> ScimAttributeResponse AddScimAttribute(ctx, scimSchemaName).AddScimAttributeRequest(addScimAttributeRequest).Execute()

Add a new SCIM Attribute to the config

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
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema
    addScimAttributeRequest := *openapiclient.NewAddScimAttributeRequest("AttributeName_example", "Name_example") // AddScimAttributeRequest | Create a new SCIM Attribute in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeApi.AddScimAttribute(context.Background(), scimSchemaName).AddScimAttributeRequest(addScimAttributeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeApi.AddScimAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddScimAttribute`: ScimAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimAttributeApi.AddScimAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddScimAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addScimAttributeRequest** | [**AddScimAttributeRequest**](AddScimAttributeRequest.md) | Create a new SCIM Attribute in the config | 

### Return type

[**ScimAttributeResponse**](ScimAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScimAttribute

> DeleteScimAttribute(ctx, scimAttributeName, scimSchemaName).Execute()

Delete a SCIM Attribute

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
    scimAttributeName := "scimAttributeName_example" // string | Name of the SCIM Attribute
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeApi.DeleteScimAttribute(context.Background(), scimAttributeName, scimSchemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeApi.DeleteScimAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimAttributeName** | **string** | Name of the SCIM Attribute | 
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScimAttributeRequest struct via the builder pattern


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


## GetScimAttribute

> ScimAttributeResponse GetScimAttribute(ctx, scimAttributeName, scimSchemaName).Execute()

Returns a single SCIM Attribute

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
    scimAttributeName := "scimAttributeName_example" // string | Name of the SCIM Attribute
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeApi.GetScimAttribute(context.Background(), scimAttributeName, scimSchemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeApi.GetScimAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScimAttribute`: ScimAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimAttributeApi.GetScimAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimAttributeName** | **string** | Name of the SCIM Attribute | 
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScimAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScimAttributeResponse**](ScimAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScimAttribute

> ScimAttributeResponse UpdateScimAttribute(ctx, scimAttributeName, scimSchemaName).UpdateRequest(updateRequest).Execute()

Update an existing SCIM Attribute by name

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
    scimAttributeName := "scimAttributeName_example" // string | Name of the SCIM Attribute
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing SCIM Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeApi.UpdateScimAttribute(context.Background(), scimAttributeName, scimSchemaName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeApi.UpdateScimAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScimAttribute`: ScimAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimAttributeApi.UpdateScimAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimAttributeName** | **string** | Name of the SCIM Attribute | 
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScimAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing SCIM Attribute | 

### Return type

[**ScimAttributeResponse**](ScimAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

