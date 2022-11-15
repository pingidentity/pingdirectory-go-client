# \RestResourceTypeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRestResourceType**](RestResourceTypeApi.md#AddRestResourceType) | **Post** /rest-resource-types | Add a new REST Resource Type to the config
[**DeleteRestResourceType**](RestResourceTypeApi.md#DeleteRestResourceType) | **Delete** /rest-resource-types/{rest-resource-type-name} | Delete a REST Resource Type
[**GetRestResourceType**](RestResourceTypeApi.md#GetRestResourceType) | **Get** /rest-resource-types/{rest-resource-type-name} | Returns a single REST Resource Type
[**UpdateRestResourceType**](RestResourceTypeApi.md#UpdateRestResourceType) | **Patch** /rest-resource-types/{rest-resource-type-name} | Update an existing REST Resource Type by name



## AddRestResourceType

> AddRestResourceType200Response AddRestResourceType(ctx).AddRestResourceTypeRequest(addRestResourceTypeRequest).Execute()

Add a new REST Resource Type to the config

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
    addRestResourceTypeRequest := openapiclient.add_rest_resource_type_request{AddGroupRestResourceTypeRequest: openapiclient.NewAddGroupRestResourceTypeRequest("TypeName_example", []openapiclient.EnumgroupRestResourceTypeSchemaUrn{openapiclient.Enumgroup-rest-resource-typeSchemaUrn("urn:pingidentity:schemas:configuration:2.0:rest-resource-type:group")}, false, "ResourceEndpoint_example", "StructuralLDAPObjectclass_example", "SearchBaseDN_example")} // AddRestResourceTypeRequest | Create a new REST Resource Type in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestResourceTypeApi.AddRestResourceType(context.Background()).AddRestResourceTypeRequest(addRestResourceTypeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestResourceTypeApi.AddRestResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRestResourceType`: AddRestResourceType200Response
    fmt.Fprintf(os.Stdout, "Response from `RestResourceTypeApi.AddRestResourceType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRestResourceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addRestResourceTypeRequest** | [**AddRestResourceTypeRequest**](AddRestResourceTypeRequest.md) | Create a new REST Resource Type in the config | 

### Return type

[**AddRestResourceType200Response**](AddRestResourceType200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRestResourceType

> DeleteRestResourceType(ctx, restResourceTypeName).Execute()

Delete a REST Resource Type

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
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestResourceTypeApi.DeleteRestResourceType(context.Background(), restResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestResourceTypeApi.DeleteRestResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restResourceTypeName** | **string** | Name of the REST Resource Type to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRestResourceTypeRequest struct via the builder pattern


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


## GetRestResourceType

> AddRestResourceType200Response GetRestResourceType(ctx, restResourceTypeName).Execute()

Returns a single REST Resource Type

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
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestResourceTypeApi.GetRestResourceType(context.Background(), restResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestResourceTypeApi.GetRestResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestResourceType`: AddRestResourceType200Response
    fmt.Fprintf(os.Stdout, "Response from `RestResourceTypeApi.GetRestResourceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restResourceTypeName** | **string** | Name of the REST Resource Type to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestResourceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddRestResourceType200Response**](AddRestResourceType200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRestResourceType

> AddRestResourceType200Response UpdateRestResourceType(ctx, restResourceTypeName).UpdateRequest(updateRequest).Execute()

Update an existing REST Resource Type by name

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
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing REST Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestResourceTypeApi.UpdateRestResourceType(context.Background(), restResourceTypeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestResourceTypeApi.UpdateRestResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRestResourceType`: AddRestResourceType200Response
    fmt.Fprintf(os.Stdout, "Response from `RestResourceTypeApi.UpdateRestResourceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restResourceTypeName** | **string** | Name of the REST Resource Type to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRestResourceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing REST Resource Type | 

### Return type

[**AddRestResourceType200Response**](AddRestResourceType200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

