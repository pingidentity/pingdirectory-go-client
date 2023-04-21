# \ScimResourceTypeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScimResourceType**](ScimResourceTypeApi.md#AddScimResourceType) | **Post** /scim-resource-types | Add a new SCIM Resource Type to the config
[**DeleteScimResourceType**](ScimResourceTypeApi.md#DeleteScimResourceType) | **Delete** /scim-resource-types/{scim-resource-type-name} | Delete a SCIM Resource Type
[**GetScimResourceType**](ScimResourceTypeApi.md#GetScimResourceType) | **Get** /scim-resource-types/{scim-resource-type-name} | Returns a single SCIM Resource Type
[**UpdateScimResourceType**](ScimResourceTypeApi.md#UpdateScimResourceType) | **Patch** /scim-resource-types/{scim-resource-type-name} | Update an existing SCIM Resource Type by name



## AddScimResourceType

> AddScimResourceType200Response AddScimResourceType(ctx).AddScimResourceTypeRequest(addScimResourceTypeRequest).Execute()

Add a new SCIM Resource Type to the config

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
    addScimResourceTypeRequest := openapiclient.add_scim_resource_type_request{AddLdapMappingScimResourceTypeRequest: openapiclient.NewAddLdapMappingScimResourceTypeRequest("TypeName_example", []openapiclient.EnumldapMappingScimResourceTypeSchemaUrn{openapiclient.Enumldap-mapping-scim-resource-typeSchemaUrn("urn:pingidentity:schemas:configuration:2.0:scim-resource-type:ldap-mapping")}, "CoreSchema_example", false, "Endpoint_example")} // AddScimResourceTypeRequest | Create a new SCIM Resource Type in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimResourceTypeApi.AddScimResourceType(context.Background()).AddScimResourceTypeRequest(addScimResourceTypeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimResourceTypeApi.AddScimResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddScimResourceType`: AddScimResourceType200Response
    fmt.Fprintf(os.Stdout, "Response from `ScimResourceTypeApi.AddScimResourceType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddScimResourceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addScimResourceTypeRequest** | [**AddScimResourceTypeRequest**](AddScimResourceTypeRequest.md) | Create a new SCIM Resource Type in the config | 

### Return type

[**AddScimResourceType200Response**](AddScimResourceType200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScimResourceType

> DeleteScimResourceType(ctx, scimResourceTypeName).Execute()

Delete a SCIM Resource Type

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScimResourceTypeApi.DeleteScimResourceType(context.Background(), scimResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimResourceTypeApi.DeleteScimResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScimResourceTypeRequest struct via the builder pattern


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


## GetScimResourceType

> AddScimResourceType200Response GetScimResourceType(ctx, scimResourceTypeName).Execute()

Returns a single SCIM Resource Type

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimResourceTypeApi.GetScimResourceType(context.Background(), scimResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimResourceTypeApi.GetScimResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScimResourceType`: AddScimResourceType200Response
    fmt.Fprintf(os.Stdout, "Response from `ScimResourceTypeApi.GetScimResourceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScimResourceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddScimResourceType200Response**](AddScimResourceType200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScimResourceType

> AddScimResourceType200Response UpdateScimResourceType(ctx, scimResourceTypeName).UpdateRequest(updateRequest).Execute()

Update an existing SCIM Resource Type by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing SCIM Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimResourceTypeApi.UpdateScimResourceType(context.Background(), scimResourceTypeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimResourceTypeApi.UpdateScimResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScimResourceType`: AddScimResourceType200Response
    fmt.Fprintf(os.Stdout, "Response from `ScimResourceTypeApi.UpdateScimResourceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScimResourceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing SCIM Resource Type | 

### Return type

[**AddScimResourceType200Response**](AddScimResourceType200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

