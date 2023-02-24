# \SearchReferenceCriteriaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSearchReferenceCriteria**](SearchReferenceCriteriaApi.md#AddSearchReferenceCriteria) | **Post** /search-reference-criteria | Add a new Search Reference Criteria to the config
[**DeleteSearchReferenceCriteria**](SearchReferenceCriteriaApi.md#DeleteSearchReferenceCriteria) | **Delete** /search-reference-criteria/{search-reference-criteria-name} | Delete a Search Reference Criteria
[**GetSearchReferenceCriteria**](SearchReferenceCriteriaApi.md#GetSearchReferenceCriteria) | **Get** /search-reference-criteria/{search-reference-criteria-name} | Returns a single Search Reference Criteria
[**UpdateSearchReferenceCriteria**](SearchReferenceCriteriaApi.md#UpdateSearchReferenceCriteria) | **Patch** /search-reference-criteria/{search-reference-criteria-name} | Update an existing Search Reference Criteria by name



## AddSearchReferenceCriteria

> AddSearchReferenceCriteria200Response AddSearchReferenceCriteria(ctx).AddSearchReferenceCriteriaRequest(addSearchReferenceCriteriaRequest).Execute()

Add a new Search Reference Criteria to the config

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
    addSearchReferenceCriteriaRequest := openapiclient.add_search_reference_criteria_request{AddAggregateSearchReferenceCriteriaRequest: openapiclient.NewAddAggregateSearchReferenceCriteriaRequest("CriteriaName_example", []openapiclient.EnumaggregateSearchReferenceCriteriaSchemaUrn{openapiclient.Enumaggregate-search-reference-criteriaSchemaUrn("urn:pingidentity:schemas:configuration:2.0:search-reference-criteria:aggregate")})} // AddSearchReferenceCriteriaRequest | Create a new Search Reference Criteria in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchReferenceCriteriaApi.AddSearchReferenceCriteria(context.Background()).AddSearchReferenceCriteriaRequest(addSearchReferenceCriteriaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchReferenceCriteriaApi.AddSearchReferenceCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSearchReferenceCriteria`: AddSearchReferenceCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchReferenceCriteriaApi.AddSearchReferenceCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSearchReferenceCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addSearchReferenceCriteriaRequest** | [**AddSearchReferenceCriteriaRequest**](AddSearchReferenceCriteriaRequest.md) | Create a new Search Reference Criteria in the config | 

### Return type

[**AddSearchReferenceCriteria200Response**](AddSearchReferenceCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSearchReferenceCriteria

> DeleteSearchReferenceCriteria(ctx, searchReferenceCriteriaName).Execute()

Delete a Search Reference Criteria

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
    searchReferenceCriteriaName := "searchReferenceCriteriaName_example" // string | Name of the Search Reference Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SearchReferenceCriteriaApi.DeleteSearchReferenceCriteria(context.Background(), searchReferenceCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchReferenceCriteriaApi.DeleteSearchReferenceCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchReferenceCriteriaName** | **string** | Name of the Search Reference Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSearchReferenceCriteriaRequest struct via the builder pattern


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


## GetSearchReferenceCriteria

> AddSearchReferenceCriteria200Response GetSearchReferenceCriteria(ctx, searchReferenceCriteriaName).Execute()

Returns a single Search Reference Criteria

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
    searchReferenceCriteriaName := "searchReferenceCriteriaName_example" // string | Name of the Search Reference Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchReferenceCriteriaApi.GetSearchReferenceCriteria(context.Background(), searchReferenceCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchReferenceCriteriaApi.GetSearchReferenceCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchReferenceCriteria`: AddSearchReferenceCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchReferenceCriteriaApi.GetSearchReferenceCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchReferenceCriteriaName** | **string** | Name of the Search Reference Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchReferenceCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddSearchReferenceCriteria200Response**](AddSearchReferenceCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSearchReferenceCriteria

> AddSearchReferenceCriteria200Response UpdateSearchReferenceCriteria(ctx, searchReferenceCriteriaName).UpdateRequest(updateRequest).Execute()

Update an existing Search Reference Criteria by name

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
    searchReferenceCriteriaName := "searchReferenceCriteriaName_example" // string | Name of the Search Reference Criteria
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Search Reference Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchReferenceCriteriaApi.UpdateSearchReferenceCriteria(context.Background(), searchReferenceCriteriaName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchReferenceCriteriaApi.UpdateSearchReferenceCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSearchReferenceCriteria`: AddSearchReferenceCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchReferenceCriteriaApi.UpdateSearchReferenceCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchReferenceCriteriaName** | **string** | Name of the Search Reference Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSearchReferenceCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Search Reference Criteria | 

### Return type

[**AddSearchReferenceCriteria200Response**](AddSearchReferenceCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

