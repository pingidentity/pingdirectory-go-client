# \ScimAttributeMappingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScimAttributeMapping**](ScimAttributeMappingAPI.md#AddScimAttributeMapping) | **Post** /scim-resource-types/{scim-resource-type-name}/scim-attribute-mappings | Add a new SCIM Attribute Mapping to the config
[**DeleteScimAttributeMapping**](ScimAttributeMappingAPI.md#DeleteScimAttributeMapping) | **Delete** /scim-resource-types/{scim-resource-type-name}/scim-attribute-mappings/{scim-attribute-mapping-name} | Delete a SCIM Attribute Mapping
[**GetScimAttributeMapping**](ScimAttributeMappingAPI.md#GetScimAttributeMapping) | **Get** /scim-resource-types/{scim-resource-type-name}/scim-attribute-mappings/{scim-attribute-mapping-name} | Returns a single SCIM Attribute Mapping
[**ListScimAttributeMappings**](ScimAttributeMappingAPI.md#ListScimAttributeMappings) | **Get** /scim-resource-types/{scim-resource-type-name}/scim-attribute-mappings | Returns a list of all SCIM Attribute Mapping objects
[**UpdateScimAttributeMapping**](ScimAttributeMappingAPI.md#UpdateScimAttributeMapping) | **Patch** /scim-resource-types/{scim-resource-type-name}/scim-attribute-mappings/{scim-attribute-mapping-name} | Update an existing SCIM Attribute Mapping by name



## AddScimAttributeMapping

> ScimAttributeMappingResponse AddScimAttributeMapping(ctx, scimResourceTypeName).AddScimAttributeMappingRequest(addScimAttributeMappingRequest).Execute()

Add a new SCIM Attribute Mapping to the config

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
    scimResourceTypeName := "scimResourceTypeName_example" // string | Name of the SCIM Resource Type
    addScimAttributeMappingRequest := *openapiclient.NewAddScimAttributeMappingRequest("ScimResourceTypeAttribute_example", "LdapAttribute_example", "MappingName_example") // AddScimAttributeMappingRequest | Create a new SCIM Attribute Mapping in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeMappingAPI.AddScimAttributeMapping(context.Background(), scimResourceTypeName).AddScimAttributeMappingRequest(addScimAttributeMappingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeMappingAPI.AddScimAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddScimAttributeMapping`: ScimAttributeMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimAttributeMappingAPI.AddScimAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

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

> DeleteScimAttributeMapping(ctx, scimAttributeMappingName, scimResourceTypeName).Execute()

Delete a SCIM Attribute Mapping

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
    scimAttributeMappingName := "scimAttributeMappingName_example" // string | Name of the SCIM Attribute Mapping
    scimResourceTypeName := "scimResourceTypeName_example" // string | Name of the SCIM Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScimAttributeMappingAPI.DeleteScimAttributeMapping(context.Background(), scimAttributeMappingName, scimResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeMappingAPI.DeleteScimAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimAttributeMappingName** | **string** | Name of the SCIM Attribute Mapping | 
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

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

> ScimAttributeMappingResponse GetScimAttributeMapping(ctx, scimAttributeMappingName, scimResourceTypeName).Execute()

Returns a single SCIM Attribute Mapping

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
    scimAttributeMappingName := "scimAttributeMappingName_example" // string | Name of the SCIM Attribute Mapping
    scimResourceTypeName := "scimResourceTypeName_example" // string | Name of the SCIM Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeMappingAPI.GetScimAttributeMapping(context.Background(), scimAttributeMappingName, scimResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeMappingAPI.GetScimAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScimAttributeMapping`: ScimAttributeMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimAttributeMappingAPI.GetScimAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimAttributeMappingName** | **string** | Name of the SCIM Attribute Mapping | 
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

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


## ListScimAttributeMappings

> ScimAttributeMappingListResponse ListScimAttributeMappings(ctx, scimResourceTypeName).Filter(filter).Execute()

Returns a list of all SCIM Attribute Mapping objects

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
    scimResourceTypeName := "scimResourceTypeName_example" // string | Name of the SCIM Resource Type
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeMappingAPI.ListScimAttributeMappings(context.Background(), scimResourceTypeName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeMappingAPI.ListScimAttributeMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScimAttributeMappings`: ScimAttributeMappingListResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimAttributeMappingAPI.ListScimAttributeMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListScimAttributeMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**ScimAttributeMappingListResponse**](ScimAttributeMappingListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScimAttributeMapping

> ScimAttributeMappingResponse UpdateScimAttributeMapping(ctx, scimAttributeMappingName, scimResourceTypeName).UpdateRequest(updateRequest).Execute()

Update an existing SCIM Attribute Mapping by name

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
    scimAttributeMappingName := "scimAttributeMappingName_example" // string | Name of the SCIM Attribute Mapping
    scimResourceTypeName := "scimResourceTypeName_example" // string | Name of the SCIM Resource Type
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing SCIM Attribute Mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimAttributeMappingAPI.UpdateScimAttributeMapping(context.Background(), scimAttributeMappingName, scimResourceTypeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimAttributeMappingAPI.UpdateScimAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScimAttributeMapping`: ScimAttributeMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimAttributeMappingAPI.UpdateScimAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimAttributeMappingName** | **string** | Name of the SCIM Attribute Mapping | 
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

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

