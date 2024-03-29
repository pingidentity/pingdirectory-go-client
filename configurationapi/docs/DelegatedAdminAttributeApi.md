# \DelegatedAdminAttributeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDelegatedAdminAttribute**](DelegatedAdminAttributeAPI.md#AddDelegatedAdminAttribute) | **Post** /rest-resource-types/{rest-resource-type-name}/delegated-admin-attributes | Add a new Delegated Admin Attribute to the config
[**DeleteDelegatedAdminAttribute**](DelegatedAdminAttributeAPI.md#DeleteDelegatedAdminAttribute) | **Delete** /rest-resource-types/{rest-resource-type-name}/delegated-admin-attributes/{delegated-admin-attribute-name} | Delete a Delegated Admin Attribute
[**GetDelegatedAdminAttribute**](DelegatedAdminAttributeAPI.md#GetDelegatedAdminAttribute) | **Get** /rest-resource-types/{rest-resource-type-name}/delegated-admin-attributes/{delegated-admin-attribute-name} | Returns a single Delegated Admin Attribute
[**ListDelegatedAdminAttributes**](DelegatedAdminAttributeAPI.md#ListDelegatedAdminAttributes) | **Get** /rest-resource-types/{rest-resource-type-name}/delegated-admin-attributes | Returns a list of all Delegated Admin Attribute objects
[**UpdateDelegatedAdminAttribute**](DelegatedAdminAttributeAPI.md#UpdateDelegatedAdminAttribute) | **Patch** /rest-resource-types/{rest-resource-type-name}/delegated-admin-attributes/{delegated-admin-attribute-name} | Update an existing Delegated Admin Attribute by name



## AddDelegatedAdminAttribute

> AddDelegatedAdminAttribute200Response AddDelegatedAdminAttribute(ctx, restResourceTypeName).AddDelegatedAdminAttributeRequest(addDelegatedAdminAttributeRequest).Execute()

Add a new Delegated Admin Attribute to the config

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
    addDelegatedAdminAttributeRequest := openapiclient.add_delegated_admin_attribute_request{AddCertificateDelegatedAdminAttributeRequest: openapiclient.NewAddCertificateDelegatedAdminAttributeRequest([]openapiclient.EnumcertificateDelegatedAdminAttributeSchemaUrn{openapiclient.Enumcertificate-delegated-admin-attributeSchemaUrn("urn:pingidentity:schemas:configuration:2.0:delegated-admin-attribute:certificate")}, "AttributeType_example", "DisplayName_example")} // AddDelegatedAdminAttributeRequest | Create a new Delegated Admin Attribute in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminAttributeAPI.AddDelegatedAdminAttribute(context.Background(), restResourceTypeName).AddDelegatedAdminAttributeRequest(addDelegatedAdminAttributeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminAttributeAPI.AddDelegatedAdminAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDelegatedAdminAttribute`: AddDelegatedAdminAttribute200Response
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminAttributeAPI.AddDelegatedAdminAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDelegatedAdminAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addDelegatedAdminAttributeRequest** | [**AddDelegatedAdminAttributeRequest**](AddDelegatedAdminAttributeRequest.md) | Create a new Delegated Admin Attribute in the config | 

### Return type

[**AddDelegatedAdminAttribute200Response**](AddDelegatedAdminAttribute200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDelegatedAdminAttribute

> DeleteDelegatedAdminAttribute(ctx, delegatedAdminAttributeName, restResourceTypeName).Execute()

Delete a Delegated Admin Attribute

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
    delegatedAdminAttributeName := "delegatedAdminAttributeName_example" // string | Name of the Delegated Admin Attribute
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DelegatedAdminAttributeAPI.DeleteDelegatedAdminAttribute(context.Background(), delegatedAdminAttributeName, restResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminAttributeAPI.DeleteDelegatedAdminAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminAttributeName** | **string** | Name of the Delegated Admin Attribute | 
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDelegatedAdminAttributeRequest struct via the builder pattern


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


## GetDelegatedAdminAttribute

> AddDelegatedAdminAttribute200Response GetDelegatedAdminAttribute(ctx, delegatedAdminAttributeName, restResourceTypeName).Execute()

Returns a single Delegated Admin Attribute

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
    delegatedAdminAttributeName := "delegatedAdminAttributeName_example" // string | Name of the Delegated Admin Attribute
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminAttributeAPI.GetDelegatedAdminAttribute(context.Background(), delegatedAdminAttributeName, restResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminAttributeAPI.GetDelegatedAdminAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDelegatedAdminAttribute`: AddDelegatedAdminAttribute200Response
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminAttributeAPI.GetDelegatedAdminAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminAttributeName** | **string** | Name of the Delegated Admin Attribute | 
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDelegatedAdminAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddDelegatedAdminAttribute200Response**](AddDelegatedAdminAttribute200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDelegatedAdminAttributes

> DelegatedAdminAttributeListResponse ListDelegatedAdminAttributes(ctx, restResourceTypeName).Filter(filter).Execute()

Returns a list of all Delegated Admin Attribute objects

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
    resp, r, err := apiClient.DelegatedAdminAttributeAPI.ListDelegatedAdminAttributes(context.Background(), restResourceTypeName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminAttributeAPI.ListDelegatedAdminAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDelegatedAdminAttributes`: DelegatedAdminAttributeListResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminAttributeAPI.ListDelegatedAdminAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDelegatedAdminAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**DelegatedAdminAttributeListResponse**](DelegatedAdminAttributeListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDelegatedAdminAttribute

> AddDelegatedAdminAttribute200Response UpdateDelegatedAdminAttribute(ctx, delegatedAdminAttributeName, restResourceTypeName).UpdateRequest(updateRequest).Execute()

Update an existing Delegated Admin Attribute by name

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
    delegatedAdminAttributeName := "delegatedAdminAttributeName_example" // string | Name of the Delegated Admin Attribute
    restResourceTypeName := "restResourceTypeName_example" // string | Name of the REST Resource Type
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Delegated Admin Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminAttributeAPI.UpdateDelegatedAdminAttribute(context.Background(), delegatedAdminAttributeName, restResourceTypeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminAttributeAPI.UpdateDelegatedAdminAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDelegatedAdminAttribute`: AddDelegatedAdminAttribute200Response
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminAttributeAPI.UpdateDelegatedAdminAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminAttributeName** | **string** | Name of the Delegated Admin Attribute | 
**restResourceTypeName** | **string** | Name of the REST Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDelegatedAdminAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Delegated Admin Attribute | 

### Return type

[**AddDelegatedAdminAttribute200Response**](AddDelegatedAdminAttribute200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

