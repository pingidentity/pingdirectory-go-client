# \LdapCorrelationAttributePairApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLdapCorrelationAttributePair**](LdapCorrelationAttributePairApi.md#AddLdapCorrelationAttributePair) | **Post** /scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}/ldap-correlation-attribute-pairs | Add a new LDAP Correlation Attribute Pair to the config
[**DeleteLdapCorrelationAttributePair**](LdapCorrelationAttributePairApi.md#DeleteLdapCorrelationAttributePair) | **Delete** /scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}/ldap-correlation-attribute-pairs/{ldap-correlation-attribute-pair-name} | Delete a LDAP Correlation Attribute Pair
[**GetLdapCorrelationAttributePair**](LdapCorrelationAttributePairApi.md#GetLdapCorrelationAttributePair) | **Get** /scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}/ldap-correlation-attribute-pairs/{ldap-correlation-attribute-pair-name} | Returns a single LDAP Correlation Attribute Pair
[**ListLdapCorrelationAttributePairs**](LdapCorrelationAttributePairApi.md#ListLdapCorrelationAttributePairs) | **Get** /scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}/ldap-correlation-attribute-pairs | Returns a list of all LDAP Correlation Attribute Pair objects
[**UpdateLdapCorrelationAttributePair**](LdapCorrelationAttributePairApi.md#UpdateLdapCorrelationAttributePair) | **Patch** /scim-resource-types/{scim-resource-type-name}/correlated-ldap-data-views/{correlated-ldap-data-view-name}/ldap-correlation-attribute-pairs/{ldap-correlation-attribute-pair-name} | Update an existing LDAP Correlation Attribute Pair by name



## AddLdapCorrelationAttributePair

> LdapCorrelationAttributePairResponse AddLdapCorrelationAttributePair(ctx, correlatedLdapDataViewName, scimResourceTypeName).AddLdapCorrelationAttributePairRequest(addLdapCorrelationAttributePairRequest).Execute()

Add a new LDAP Correlation Attribute Pair to the config

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
    addLdapCorrelationAttributePairRequest := *openapiclient.NewAddLdapCorrelationAttributePairRequest("PairName_example", "PrimaryCorrelationAttribute_example", "SecondaryCorrelationAttribute_example") // AddLdapCorrelationAttributePairRequest | Create a new LDAP Correlation Attribute Pair in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LdapCorrelationAttributePairApi.AddLdapCorrelationAttributePair(context.Background(), correlatedLdapDataViewName, scimResourceTypeName).AddLdapCorrelationAttributePairRequest(addLdapCorrelationAttributePairRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapCorrelationAttributePairApi.AddLdapCorrelationAttributePair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLdapCorrelationAttributePair`: LdapCorrelationAttributePairResponse
    fmt.Fprintf(os.Stdout, "Response from `LdapCorrelationAttributePairApi.AddLdapCorrelationAttributePair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**correlatedLdapDataViewName** | **string** | Name of the Correlated LDAP Data View | 
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLdapCorrelationAttributePairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addLdapCorrelationAttributePairRequest** | [**AddLdapCorrelationAttributePairRequest**](AddLdapCorrelationAttributePairRequest.md) | Create a new LDAP Correlation Attribute Pair in the config | 

### Return type

[**LdapCorrelationAttributePairResponse**](LdapCorrelationAttributePairResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLdapCorrelationAttributePair

> DeleteLdapCorrelationAttributePair(ctx, ldapCorrelationAttributePairName, correlatedLdapDataViewName, scimResourceTypeName).Execute()

Delete a LDAP Correlation Attribute Pair

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
    ldapCorrelationAttributePairName := "ldapCorrelationAttributePairName_example" // string | Name of the LDAP Correlation Attribute Pair
    correlatedLdapDataViewName := "correlatedLdapDataViewName_example" // string | Name of the Correlated LDAP Data View
    scimResourceTypeName := "scimResourceTypeName_example" // string | Name of the SCIM Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LdapCorrelationAttributePairApi.DeleteLdapCorrelationAttributePair(context.Background(), ldapCorrelationAttributePairName, correlatedLdapDataViewName, scimResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapCorrelationAttributePairApi.DeleteLdapCorrelationAttributePair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapCorrelationAttributePairName** | **string** | Name of the LDAP Correlation Attribute Pair | 
**correlatedLdapDataViewName** | **string** | Name of the Correlated LDAP Data View | 
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLdapCorrelationAttributePairRequest struct via the builder pattern


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


## GetLdapCorrelationAttributePair

> LdapCorrelationAttributePairResponse GetLdapCorrelationAttributePair(ctx, ldapCorrelationAttributePairName, correlatedLdapDataViewName, scimResourceTypeName).Execute()

Returns a single LDAP Correlation Attribute Pair

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
    ldapCorrelationAttributePairName := "ldapCorrelationAttributePairName_example" // string | Name of the LDAP Correlation Attribute Pair
    correlatedLdapDataViewName := "correlatedLdapDataViewName_example" // string | Name of the Correlated LDAP Data View
    scimResourceTypeName := "scimResourceTypeName_example" // string | Name of the SCIM Resource Type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LdapCorrelationAttributePairApi.GetLdapCorrelationAttributePair(context.Background(), ldapCorrelationAttributePairName, correlatedLdapDataViewName, scimResourceTypeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapCorrelationAttributePairApi.GetLdapCorrelationAttributePair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapCorrelationAttributePair`: LdapCorrelationAttributePairResponse
    fmt.Fprintf(os.Stdout, "Response from `LdapCorrelationAttributePairApi.GetLdapCorrelationAttributePair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapCorrelationAttributePairName** | **string** | Name of the LDAP Correlation Attribute Pair | 
**correlatedLdapDataViewName** | **string** | Name of the Correlated LDAP Data View | 
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapCorrelationAttributePairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**LdapCorrelationAttributePairResponse**](LdapCorrelationAttributePairResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLdapCorrelationAttributePairs

> LdapCorrelationAttributePairListResponse ListLdapCorrelationAttributePairs(ctx, correlatedLdapDataViewName, scimResourceTypeName).Filter(filter).Execute()

Returns a list of all LDAP Correlation Attribute Pair objects

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
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LdapCorrelationAttributePairApi.ListLdapCorrelationAttributePairs(context.Background(), correlatedLdapDataViewName, scimResourceTypeName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapCorrelationAttributePairApi.ListLdapCorrelationAttributePairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLdapCorrelationAttributePairs`: LdapCorrelationAttributePairListResponse
    fmt.Fprintf(os.Stdout, "Response from `LdapCorrelationAttributePairApi.ListLdapCorrelationAttributePairs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**correlatedLdapDataViewName** | **string** | Name of the Correlated LDAP Data View | 
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLdapCorrelationAttributePairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | SCIM filter | 

### Return type

[**LdapCorrelationAttributePairListResponse**](LdapCorrelationAttributePairListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLdapCorrelationAttributePair

> LdapCorrelationAttributePairResponse UpdateLdapCorrelationAttributePair(ctx, ldapCorrelationAttributePairName, correlatedLdapDataViewName, scimResourceTypeName).UpdateRequest(updateRequest).Execute()

Update an existing LDAP Correlation Attribute Pair by name

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
    ldapCorrelationAttributePairName := "ldapCorrelationAttributePairName_example" // string | Name of the LDAP Correlation Attribute Pair
    correlatedLdapDataViewName := "correlatedLdapDataViewName_example" // string | Name of the Correlated LDAP Data View
    scimResourceTypeName := "scimResourceTypeName_example" // string | Name of the SCIM Resource Type
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing LDAP Correlation Attribute Pair

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LdapCorrelationAttributePairApi.UpdateLdapCorrelationAttributePair(context.Background(), ldapCorrelationAttributePairName, correlatedLdapDataViewName, scimResourceTypeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapCorrelationAttributePairApi.UpdateLdapCorrelationAttributePair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLdapCorrelationAttributePair`: LdapCorrelationAttributePairResponse
    fmt.Fprintf(os.Stdout, "Response from `LdapCorrelationAttributePairApi.UpdateLdapCorrelationAttributePair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ldapCorrelationAttributePairName** | **string** | Name of the LDAP Correlation Attribute Pair | 
**correlatedLdapDataViewName** | **string** | Name of the Correlated LDAP Data View | 
**scimResourceTypeName** | **string** | Name of the SCIM Resource Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapCorrelationAttributePairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing LDAP Correlation Attribute Pair | 

### Return type

[**LdapCorrelationAttributePairResponse**](LdapCorrelationAttributePairResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

