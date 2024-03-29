# \DelegatedAdminCorrelatedRestResourceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDelegatedAdminCorrelatedRestResource**](DelegatedAdminCorrelatedRestResourceAPI.md#AddDelegatedAdminCorrelatedRestResource) | **Post** /rest-resource-types/{rest-resource-type-name}/delegated-admin-correlated-rest-resources | Add a new Delegated Admin Correlated REST Resource to the config
[**DeleteDelegatedAdminCorrelatedRestResource**](DelegatedAdminCorrelatedRestResourceAPI.md#DeleteDelegatedAdminCorrelatedRestResource) | **Delete** /rest-resource-types/{rest-resource-type-name}/delegated-admin-correlated-rest-resources/{delegated-admin-correlated-rest-resource-name} | Delete a Delegated Admin Correlated REST Resource
[**GetDelegatedAdminCorrelatedRestResource**](DelegatedAdminCorrelatedRestResourceAPI.md#GetDelegatedAdminCorrelatedRestResource) | **Get** /rest-resource-types/{rest-resource-type-name}/delegated-admin-correlated-rest-resources/{delegated-admin-correlated-rest-resource-name} | Returns a single Delegated Admin Correlated REST Resource
[**ListDelegatedAdminCorrelatedRestResources**](DelegatedAdminCorrelatedRestResourceAPI.md#ListDelegatedAdminCorrelatedRestResources) | **Get** /rest-resource-types/{rest-resource-type-name}/delegated-admin-correlated-rest-resources | Returns a list of all Delegated Admin Correlated REST Resource objects
[**UpdateDelegatedAdminCorrelatedRestResource**](DelegatedAdminCorrelatedRestResourceAPI.md#UpdateDelegatedAdminCorrelatedRestResource) | **Patch** /rest-resource-types/{rest-resource-type-name}/delegated-admin-correlated-rest-resources/{delegated-admin-correlated-rest-resource-name} | Update an existing Delegated Admin Correlated REST Resource by name



## AddDelegatedAdminCorrelatedRestResource

> DelegatedAdminCorrelatedRestResourceResponse AddDelegatedAdminCorrelatedRestResource(ctx, restResourceTypeName).AddDelegatedAdminCorrelatedRestResourceRequest(addDelegatedAdminCorrelatedRestResourceRequest).Execute()

Add a new Delegated Admin Correlated REST Resource to the config

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
    addDelegatedAdminCorrelatedRestResourceRequest := *openapiclient.NewAddDelegatedAdminCorrelatedRestResourceRequest("DisplayName_example", "CorrelatedRESTResource_example", "PrimaryRESTResourceCorrelationAttribute_example", "SecondaryRESTResourceCorrelationAttribute_example", "ResourceName_example") // AddDelegatedAdminCorrelatedRestResourceRequest | Create a new Delegated Admin Correlated REST Resource in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminCorrelatedRestResourceAPI.AddDelegatedAdminCorrelatedRestResource(context.Background(), restResourceTypeName).AddDelegatedAdminCorrelatedRestResourceRequest(addDelegatedAdminCorrelatedRestResourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminCorrelatedRestResourceAPI.AddDelegatedAdminCorrelatedRestResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDelegatedAdminCorrelatedRestResource`: DelegatedAdminCorrelatedRestResourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminCorrelatedRestResourceAPI.AddDelegatedAdminCorrelatedRestResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDelegatedAdminCorrelatedRestResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addDelegatedAdminCorrelatedRestResourceRequest** | [**AddDelegatedAdminCorrelatedRestResourceRequest**](AddDelegatedAdminCorrelatedRestResourceRequest.md) | Create a new Delegated Admin Correlated REST Resource in the config | 

### Return type

[**DelegatedAdminCorrelatedRestResourceResponse**](DelegatedAdminCorrelatedRestResourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDelegatedAdminCorrelatedRestResource

> DeleteDelegatedAdminCorrelatedRestResource(ctx, delegatedAdminCorrelatedRestResourceName, restResourceTypeName).Execute()

Delete a Delegated Admin Correlated REST Resource

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
    delegatedAdminCorrelatedRestResourceName := "delegatedAdminCorrelatedRestResourceName_example" // string | Name of the Delegated Admin Correlated REST Resource
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DelegatedAdminCorrelatedRestResourceAPI.DeleteDelegatedAdminCorrelatedRestResource(context.Background(), delegatedAdminCorrelatedRestResourceName, restResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminCorrelatedRestResourceAPI.DeleteDelegatedAdminCorrelatedRestResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminCorrelatedRestResourceName** | **string** | Name of the Delegated Admin Correlated REST Resource | 
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDelegatedAdminCorrelatedRestResourceRequest struct via the builder pattern


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


## GetDelegatedAdminCorrelatedRestResource

> DelegatedAdminCorrelatedRestResourceResponse GetDelegatedAdminCorrelatedRestResource(ctx, delegatedAdminCorrelatedRestResourceName, restResourceTypeName).Execute()

Returns a single Delegated Admin Correlated REST Resource

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
    delegatedAdminCorrelatedRestResourceName := "delegatedAdminCorrelatedRestResourceName_example" // string | Name of the Delegated Admin Correlated REST Resource
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminCorrelatedRestResourceAPI.GetDelegatedAdminCorrelatedRestResource(context.Background(), delegatedAdminCorrelatedRestResourceName, restResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminCorrelatedRestResourceAPI.GetDelegatedAdminCorrelatedRestResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDelegatedAdminCorrelatedRestResource`: DelegatedAdminCorrelatedRestResourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminCorrelatedRestResourceAPI.GetDelegatedAdminCorrelatedRestResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminCorrelatedRestResourceName** | **string** | Name of the Delegated Admin Correlated REST Resource | 
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDelegatedAdminCorrelatedRestResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DelegatedAdminCorrelatedRestResourceResponse**](DelegatedAdminCorrelatedRestResourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDelegatedAdminCorrelatedRestResources

> DelegatedAdminCorrelatedRestResourceListResponse ListDelegatedAdminCorrelatedRestResources(ctx, restResourceTypeName).Filter(filter).Execute()

Returns a list of all Delegated Admin Correlated REST Resource objects

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
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminCorrelatedRestResourceAPI.ListDelegatedAdminCorrelatedRestResources(context.Background(), restResourceTypeName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminCorrelatedRestResourceAPI.ListDelegatedAdminCorrelatedRestResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDelegatedAdminCorrelatedRestResources`: DelegatedAdminCorrelatedRestResourceListResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminCorrelatedRestResourceAPI.ListDelegatedAdminCorrelatedRestResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDelegatedAdminCorrelatedRestResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**DelegatedAdminCorrelatedRestResourceListResponse**](DelegatedAdminCorrelatedRestResourceListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDelegatedAdminCorrelatedRestResource

> DelegatedAdminCorrelatedRestResourceResponse UpdateDelegatedAdminCorrelatedRestResource(ctx, delegatedAdminCorrelatedRestResourceName, restResourceTypeName).UpdateRequest(updateRequest).Execute()

Update an existing Delegated Admin Correlated REST Resource by name

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
    delegatedAdminCorrelatedRestResourceName := "delegatedAdminCorrelatedRestResourceName_example" // string | Name of the Delegated Admin Correlated REST Resource
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Delegated Admin Correlated REST Resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminCorrelatedRestResourceAPI.UpdateDelegatedAdminCorrelatedRestResource(context.Background(), delegatedAdminCorrelatedRestResourceName, restResourceTypeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminCorrelatedRestResourceAPI.UpdateDelegatedAdminCorrelatedRestResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDelegatedAdminCorrelatedRestResource`: DelegatedAdminCorrelatedRestResourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminCorrelatedRestResourceAPI.UpdateDelegatedAdminCorrelatedRestResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminCorrelatedRestResourceName** | **string** | Name of the Delegated Admin Correlated REST Resource | 
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDelegatedAdminCorrelatedRestResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Delegated Admin Correlated REST Resource | 

### Return type

[**DelegatedAdminCorrelatedRestResourceResponse**](DelegatedAdminCorrelatedRestResourceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

