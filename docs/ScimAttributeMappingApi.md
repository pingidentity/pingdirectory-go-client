# \ScimAttributeMappingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScimAttributeMapping**](ScimAttributeMappingApi.md#AddScimAttributeMapping) | **Post** /scim-attribute-mappings | Add a new SCIM Attribute Mapping to the config
[**DeleteScimAttributeMapping**](ScimAttributeMappingApi.md#DeleteScimAttributeMapping) | **Delete** /scim-attribute-mappings/{scim-attribute-mapping-name} | Delete a SCIM Attribute Mapping
[**GetScimAttributeMapping**](ScimAttributeMappingApi.md#GetScimAttributeMapping) | **Get** /scim-attribute-mappings/{scim-attribute-mapping-name} | Returns a single SCIM Attribute Mapping
[**UpdateScimAttributeMapping**](ScimAttributeMappingApi.md#UpdateScimAttributeMapping) | **Patch** /scim-attribute-mappings/{scim-attribute-mapping-name} | Update an existing SCIM Attribute Mapping by name



## AddScimAttributeMapping

> ScimAttributeMappingResponse AddScimAttributeMapping(ctx).AddScimAttributeMappingRequest(addScimAttributeMappingRequest).Execute()

Add a new SCIM Attribute Mapping to the config

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
    addScimAttributeMappingRequest := *openapiclient.NewAddScimAttributeMappingRequest("MappingName_example", "ScimResourceTypeAttribute_example", "LdapAttribute_example") // AddScimAttributeMappingRequest | Create a new SCIM Attribute Mapping in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeMappingApi.AddScimAttributeMapping(context.Background()).AddScimAttributeMappingRequest(addScimAttributeMappingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeMappingApi.AddScimAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddScimAttributeMapping`: ScimAttributeMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimAttributeMappingApi.AddScimAttributeMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddScimAttributeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addScimAttributeMappingRequest** | [**AddScimAttributeMappingRequest**](AddScimAttributeMappingRequest.md) | Create a new SCIM Attribute Mapping in the config | 

### Return type

[**ScimAttributeMappingResponse**](ScimAttributeMappingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScimAttributeMapping

> DeleteScimAttributeMapping(ctx, scimAttributeMappingName).Execute()

Delete a SCIM Attribute Mapping

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
    scimAttributeMappingName := "scimAttributeMappingName_example" // string | Name of the SCIM Attribute Mapping to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeMappingApi.DeleteScimAttributeMapping(context.Background(), scimAttributeMappingName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeMappingApi.DeleteScimAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimAttributeMappingName** | **string** | Name of the SCIM Attribute Mapping to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScimAttributeMappingRequest struct via the builder pattern


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


## GetScimAttributeMapping

> ScimAttributeMappingResponse GetScimAttributeMapping(ctx, scimAttributeMappingName).Execute()

Returns a single SCIM Attribute Mapping

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
    scimAttributeMappingName := "scimAttributeMappingName_example" // string | Name of the SCIM Attribute Mapping to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeMappingApi.GetScimAttributeMapping(context.Background(), scimAttributeMappingName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeMappingApi.GetScimAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScimAttributeMapping`: ScimAttributeMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimAttributeMappingApi.GetScimAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimAttributeMappingName** | **string** | Name of the SCIM Attribute Mapping to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScimAttributeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScimAttributeMappingResponse**](ScimAttributeMappingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScimAttributeMapping

> ScimAttributeMappingResponse UpdateScimAttributeMapping(ctx, scimAttributeMappingName).UpdateRequest(updateRequest).Execute()

Update an existing SCIM Attribute Mapping by name

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
    scimAttributeMappingName := "scimAttributeMappingName_example" // string | Name of the SCIM Attribute Mapping to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing SCIM Attribute Mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeMappingApi.UpdateScimAttributeMapping(context.Background(), scimAttributeMappingName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeMappingApi.UpdateScimAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScimAttributeMapping`: ScimAttributeMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimAttributeMappingApi.UpdateScimAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimAttributeMappingName** | **string** | Name of the SCIM Attribute Mapping to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScimAttributeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing SCIM Attribute Mapping | 

### Return type

[**ScimAttributeMappingResponse**](ScimAttributeMappingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

