# \CorrelatedLdapDataViewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCorrelatedLdapDataView**](CorrelatedLdapDataViewApi.md#AddCorrelatedLdapDataView) | **Post** /scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views | Add a new Correlated LDAP Data View to the config
[**DeleteCorrelatedLdapDataView**](CorrelatedLdapDataViewApi.md#DeleteCorrelatedLdapDataView) | **Delete** /scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name} | Delete a Correlated LDAP Data View
[**GetCorrelatedLdapDataView**](CorrelatedLdapDataViewApi.md#GetCorrelatedLdapDataView) | **Get** /scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name} | Returns a single Correlated LDAP Data View
[**UpdateCorrelatedLdapDataView**](CorrelatedLdapDataViewApi.md#UpdateCorrelatedLdapDataView) | **Patch** /scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name} | Update an existing Correlated LDAP Data View by name



## AddCorrelatedLdapDataView

> CorrelatedLdapDataViewResponse AddCorrelatedLdapDataView(ctx, scimResourceTypeName).AddCorrelatedLdapDataViewRequest(addCorrelatedLdapDataViewRequest).Execute()

Add a new Correlated LDAP Data View to the config

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
    addCorrelatedLdapDataViewRequest := *openapiclient.NewAddCorrelatedLdapDataViewRequest("ViewName_example", "StructuralLDAPObjectclass_example", "IncludeBaseDN_example", "PrimaryCorrelationAttribute_example", "SecondaryCorrelationAttribute_example") // AddCorrelatedLdapDataViewRequest | Create a new Correlated LDAP Data View in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorrelatedLdapDataViewApi.AddCorrelatedLdapDataView(context.Background(), scimResourceTypeName).AddCorrelatedLdapDataViewRequest(addCorrelatedLdapDataViewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorrelatedLdapDataViewApi.AddCorrelatedLdapDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCorrelatedLdapDataView`: CorrelatedLdapDataViewResponse
    fmt.Fprintf(os.Stdout, "Response from `CorrelatedLdapDataViewApi.AddCorrelatedLdapDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCorrelatedLdapDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addCorrelatedLdapDataViewRequest** | [**AddCorrelatedLdapDataViewRequest**](AddCorrelatedLdapDataViewRequest.md) | Create a new Correlated LDAP Data View in the config | 

### Return type

[**CorrelatedLdapDataViewResponse**](CorrelatedLdapDataViewResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCorrelatedLdapDataView

> DeleteCorrelatedLdapDataView(ctx, correlatedLdapDataViewName, scimResourceTypeName).Execute()

Delete a Correlated LDAP Data View

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
    correlatedLdapDataViewName := "correlatedLdapDataViewName_example" // string | Name of the Correlated LDAP Data View
    scimResourceTypeName := "scimResourceTypeName_example" // string | Name of the SCIM Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorrelatedLdapDataViewApi.DeleteCorrelatedLdapDataView(context.Background(), correlatedLdapDataViewName, scimResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorrelatedLdapDataViewApi.DeleteCorrelatedLdapDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**correlatedLdapDataViewName** | **string** | Name of the Correlated LDAP Data View | 
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCorrelatedLdapDataViewRequest struct via the builder pattern


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


## GetCorrelatedLdapDataView

> CorrelatedLdapDataViewResponse GetCorrelatedLdapDataView(ctx, correlatedLdapDataViewName, scimResourceTypeName).Execute()

Returns a single Correlated LDAP Data View

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
    correlatedLdapDataViewName := "correlatedLdapDataViewName_example" // string | Name of the Correlated LDAP Data View
    scimResourceTypeName := "scimResourceTypeName_example" // string | Name of the SCIM Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorrelatedLdapDataViewApi.GetCorrelatedLdapDataView(context.Background(), correlatedLdapDataViewName, scimResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorrelatedLdapDataViewApi.GetCorrelatedLdapDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCorrelatedLdapDataView`: CorrelatedLdapDataViewResponse
    fmt.Fprintf(os.Stdout, "Response from `CorrelatedLdapDataViewApi.GetCorrelatedLdapDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**correlatedLdapDataViewName** | **string** | Name of the Correlated LDAP Data View | 
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorrelatedLdapDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CorrelatedLdapDataViewResponse**](CorrelatedLdapDataViewResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCorrelatedLdapDataView

> CorrelatedLdapDataViewResponse UpdateCorrelatedLdapDataView(ctx, correlatedLdapDataViewName, scimResourceTypeName).UpdateRequest(updateRequest).Execute()

Update an existing Correlated LDAP Data View by name

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
    correlatedLdapDataViewName := "correlatedLdapDataViewName_example" // string | Name of the Correlated LDAP Data View
    scimResourceTypeName := "scimResourceTypeName_example" // string | Name of the SCIM Resource Type
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Correlated LDAP Data View

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorrelatedLdapDataViewApi.UpdateCorrelatedLdapDataView(context.Background(), correlatedLdapDataViewName, scimResourceTypeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorrelatedLdapDataViewApi.UpdateCorrelatedLdapDataView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCorrelatedLdapDataView`: CorrelatedLdapDataViewResponse
    fmt.Fprintf(os.Stdout, "Response from `CorrelatedLdapDataViewApi.UpdateCorrelatedLdapDataView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**correlatedLdapDataViewName** | **string** | Name of the Correlated LDAP Data View | 
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCorrelatedLdapDataViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Correlated LDAP Data View | 

### Return type

[**CorrelatedLdapDataViewResponse**](CorrelatedLdapDataViewResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

