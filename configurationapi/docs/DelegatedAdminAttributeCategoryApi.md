# \DelegatedAdminAttributeCategoryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDelegatedAdminAttributeCategory**](DelegatedAdminAttributeCategoryAPI.md#AddDelegatedAdminAttributeCategory) | **Post** /delegated-admin-attribute-categories | Add a new Delegated Admin Attribute Category to the config
[**DeleteDelegatedAdminAttributeCategory**](DelegatedAdminAttributeCategoryAPI.md#DeleteDelegatedAdminAttributeCategory) | **Delete** /delegated-admin-attribute-categories/{delegated-admin-attribute-category-name} | Delete a Delegated Admin Attribute Category
[**GetDelegatedAdminAttributeCategory**](DelegatedAdminAttributeCategoryAPI.md#GetDelegatedAdminAttributeCategory) | **Get** /delegated-admin-attribute-categories/{delegated-admin-attribute-category-name} | Returns a single Delegated Admin Attribute Category
[**ListDelegatedAdminAttributeCategories**](DelegatedAdminAttributeCategoryAPI.md#ListDelegatedAdminAttributeCategories) | **Get** /delegated-admin-attribute-categories | Returns a list of all Delegated Admin Attribute Category objects
[**UpdateDelegatedAdminAttributeCategory**](DelegatedAdminAttributeCategoryAPI.md#UpdateDelegatedAdminAttributeCategory) | **Patch** /delegated-admin-attribute-categories/{delegated-admin-attribute-category-name} | Update an existing Delegated Admin Attribute Category by name



## AddDelegatedAdminAttributeCategory

> DelegatedAdminAttributeCategoryResponse AddDelegatedAdminAttributeCategory(ctx).AddDelegatedAdminAttributeCategoryRequest(addDelegatedAdminAttributeCategoryRequest).Execute()

Add a new Delegated Admin Attribute Category to the config

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
    addDelegatedAdminAttributeCategoryRequest := *openapiclient.NewAddDelegatedAdminAttributeCategoryRequest("DisplayName_example", int64(123)) // AddDelegatedAdminAttributeCategoryRequest | Create a new Delegated Admin Attribute Category in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminAttributeCategoryAPI.AddDelegatedAdminAttributeCategory(context.Background()).AddDelegatedAdminAttributeCategoryRequest(addDelegatedAdminAttributeCategoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminAttributeCategoryAPI.AddDelegatedAdminAttributeCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDelegatedAdminAttributeCategory`: DelegatedAdminAttributeCategoryResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminAttributeCategoryAPI.AddDelegatedAdminAttributeCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDelegatedAdminAttributeCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addDelegatedAdminAttributeCategoryRequest** | [**AddDelegatedAdminAttributeCategoryRequest**](AddDelegatedAdminAttributeCategoryRequest.md) | Create a new Delegated Admin Attribute Category in the config | 

### Return type

[**DelegatedAdminAttributeCategoryResponse**](DelegatedAdminAttributeCategoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDelegatedAdminAttributeCategory

> DeleteDelegatedAdminAttributeCategory(ctx, delegatedAdminAttributeCategoryName).Execute()

Delete a Delegated Admin Attribute Category

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
    delegatedAdminAttributeCategoryName := "delegatedAdminAttributeCategoryName_example" // string | Name of the Delegated Admin Attribute Category

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DelegatedAdminAttributeCategoryAPI.DeleteDelegatedAdminAttributeCategory(context.Background(), delegatedAdminAttributeCategoryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminAttributeCategoryAPI.DeleteDelegatedAdminAttributeCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminAttributeCategoryName** | **string** | Name of the Delegated Admin Attribute Category | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDelegatedAdminAttributeCategoryRequest struct via the builder pattern


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


## GetDelegatedAdminAttributeCategory

> DelegatedAdminAttributeCategoryResponse GetDelegatedAdminAttributeCategory(ctx, delegatedAdminAttributeCategoryName).Execute()

Returns a single Delegated Admin Attribute Category

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
    delegatedAdminAttributeCategoryName := "delegatedAdminAttributeCategoryName_example" // string | Name of the Delegated Admin Attribute Category

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminAttributeCategoryAPI.GetDelegatedAdminAttributeCategory(context.Background(), delegatedAdminAttributeCategoryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminAttributeCategoryAPI.GetDelegatedAdminAttributeCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDelegatedAdminAttributeCategory`: DelegatedAdminAttributeCategoryResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminAttributeCategoryAPI.GetDelegatedAdminAttributeCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminAttributeCategoryName** | **string** | Name of the Delegated Admin Attribute Category | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDelegatedAdminAttributeCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DelegatedAdminAttributeCategoryResponse**](DelegatedAdminAttributeCategoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDelegatedAdminAttributeCategories

> DelegatedAdminAttributeCategoryListResponse ListDelegatedAdminAttributeCategories(ctx).Filter(filter).Execute()

Returns a list of all Delegated Admin Attribute Category objects

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
    resp, r, err := apiClient.DelegatedAdminAttributeCategoryAPI.ListDelegatedAdminAttributeCategories(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminAttributeCategoryAPI.ListDelegatedAdminAttributeCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDelegatedAdminAttributeCategories`: DelegatedAdminAttributeCategoryListResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminAttributeCategoryAPI.ListDelegatedAdminAttributeCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDelegatedAdminAttributeCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**DelegatedAdminAttributeCategoryListResponse**](DelegatedAdminAttributeCategoryListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDelegatedAdminAttributeCategory

> DelegatedAdminAttributeCategoryResponse UpdateDelegatedAdminAttributeCategory(ctx, delegatedAdminAttributeCategoryName).UpdateRequest(updateRequest).Execute()

Update an existing Delegated Admin Attribute Category by name

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
    delegatedAdminAttributeCategoryName := "delegatedAdminAttributeCategoryName_example" // string | Name of the Delegated Admin Attribute Category
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Delegated Admin Attribute Category

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminAttributeCategoryAPI.UpdateDelegatedAdminAttributeCategory(context.Background(), delegatedAdminAttributeCategoryName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminAttributeCategoryAPI.UpdateDelegatedAdminAttributeCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDelegatedAdminAttributeCategory`: DelegatedAdminAttributeCategoryResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminAttributeCategoryAPI.UpdateDelegatedAdminAttributeCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminAttributeCategoryName** | **string** | Name of the Delegated Admin Attribute Category | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDelegatedAdminAttributeCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Delegated Admin Attribute Category | 

### Return type

[**DelegatedAdminAttributeCategoryResponse**](DelegatedAdminAttributeCategoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

