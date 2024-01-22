# \RestResourceTypeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRestResourceType**](RestResourceTypeAPI.md#AddRestResourceType) | **Post** /rest-resource-types | Add a new REST Resource Type to the config
[**DeleteRestResourceType**](RestResourceTypeAPI.md#DeleteRestResourceType) | **Delete** /rest-resource-types/{rest-resource-type-name} | Delete a REST Resource Type
[**GetRestResourceType**](RestResourceTypeAPI.md#GetRestResourceType) | **Get** /rest-resource-types/{rest-resource-type-name} | Returns a single REST Resource Type
[**ListRestResourceTypes**](RestResourceTypeAPI.md#ListRestResourceTypes) | **Get** /rest-resource-types | Returns a list of all REST Resource Type objects
[**UpdateRestResourceType**](RestResourceTypeAPI.md#UpdateRestResourceType) | **Patch** /rest-resource-types/{rest-resource-type-name} | Update an existing REST Resource Type by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addRestResourceTypeRequest := openapiclient.add_rest_resource_type_request{AddGenericRestResourceTypeRequest: openapiclient.NewAddGenericRestResourceTypeRequest([]openapiclient.EnumgenericRestResourceTypeSchemaUrn{openapiclient.Enumgeneric-rest-resource-typeSchemaUrn("urn:pingidentity:schemas:configuration:2.0:rest-resource-type:generic")}, false, "ResourceEndpoint_example", "StructuralLDAPObjectclass_example", "SearchBaseDN_example", "TypeName_example")} // AddRestResourceTypeRequest | Create a new REST Resource Type in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestResourceTypeAPI.AddRestResourceType(context.Background()).AddRestResourceTypeRequest(addRestResourceTypeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestResourceTypeAPI.AddRestResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRestResourceType`: AddRestResourceType200Response
    fmt.Fprintf(os.Stdout, "Response from `RestResourceTypeAPI.AddRestResourceType`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RestResourceTypeAPI.DeleteRestResourceType(context.Background(), restResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestResourceTypeAPI.DeleteRestResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestResourceTypeAPI.GetRestResourceType(context.Background(), restResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestResourceTypeAPI.GetRestResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestResourceType`: AddRestResourceType200Response
    fmt.Fprintf(os.Stdout, "Response from `RestResourceTypeAPI.GetRestResourceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

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


## ListRestResourceTypes

> RestResourceTypeListResponse ListRestResourceTypes(ctx).Filter(filter).Execute()

Returns a list of all REST Resource Type objects

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
    resp, r, err := apiClient.RestResourceTypeAPI.ListRestResourceTypes(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestResourceTypeAPI.ListRestResourceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRestResourceTypes`: RestResourceTypeListResponse
    fmt.Fprintf(os.Stdout, "Response from `RestResourceTypeAPI.ListRestResourceTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRestResourceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**RestResourceTypeListResponse**](RestResourceTypeListResponse.md)

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing REST Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestResourceTypeAPI.UpdateRestResourceType(context.Background(), restResourceTypeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestResourceTypeAPI.UpdateRestResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRestResourceType`: AddRestResourceType200Response
    fmt.Fprintf(os.Stdout, "Response from `RestResourceTypeAPI.UpdateRestResourceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

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

