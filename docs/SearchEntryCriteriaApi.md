# \SearchEntryCriteriaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSearchEntryCriteria**](SearchEntryCriteriaApi.md#AddSearchEntryCriteria) | **Post** /search-entry-criteria | Add a new Search Entry Criteria to the config
[**DeleteSearchEntryCriteria**](SearchEntryCriteriaApi.md#DeleteSearchEntryCriteria) | **Delete** /search-entry-criteria/{search-entry-criteria-name} | Delete a Search Entry Criteria
[**GetSearchEntryCriteria**](SearchEntryCriteriaApi.md#GetSearchEntryCriteria) | **Get** /search-entry-criteria/{search-entry-criteria-name} | Returns a single Search Entry Criteria
[**UpdateSearchEntryCriteria**](SearchEntryCriteriaApi.md#UpdateSearchEntryCriteria) | **Patch** /search-entry-criteria/{search-entry-criteria-name} | Update an existing Search Entry Criteria by name



## AddSearchEntryCriteria

> AddSearchEntryCriteria200Response AddSearchEntryCriteria(ctx).AddSearchEntryCriteriaRequest(addSearchEntryCriteriaRequest).Execute()

Add a new Search Entry Criteria to the config

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
    addSearchEntryCriteriaRequest := openapiclient.add_search_entry_criteria_request{AddAggregateSearchEntryCriteriaRequest: openapiclient.NewAddAggregateSearchEntryCriteriaRequest("CriteriaName_example", []openapiclient.EnumaggregateSearchEntryCriteriaSchemaUrn{openapiclient.Enumaggregate-search-entry-criteriaSchemaUrn("urn:pingidentity:schemas:configuration:2.0:search-entry-criteria:aggregate")})} // AddSearchEntryCriteriaRequest | Create a new Search Entry Criteria in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchEntryCriteriaApi.AddSearchEntryCriteria(context.Background()).AddSearchEntryCriteriaRequest(addSearchEntryCriteriaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchEntryCriteriaApi.AddSearchEntryCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSearchEntryCriteria`: AddSearchEntryCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchEntryCriteriaApi.AddSearchEntryCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSearchEntryCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addSearchEntryCriteriaRequest** | [**AddSearchEntryCriteriaRequest**](AddSearchEntryCriteriaRequest.md) | Create a new Search Entry Criteria in the config | 

### Return type

[**AddSearchEntryCriteria200Response**](AddSearchEntryCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSearchEntryCriteria

> DeleteSearchEntryCriteria(ctx, searchEntryCriteriaName).Execute()

Delete a Search Entry Criteria

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
    searchEntryCriteriaName := "searchEntryCriteriaName_example" // string | Name of the Search Entry Criteria to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchEntryCriteriaApi.DeleteSearchEntryCriteria(context.Background(), searchEntryCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchEntryCriteriaApi.DeleteSearchEntryCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchEntryCriteriaName** | **string** | Name of the Search Entry Criteria to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSearchEntryCriteriaRequest struct via the builder pattern


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


## GetSearchEntryCriteria

> AddSearchEntryCriteria200Response GetSearchEntryCriteria(ctx, searchEntryCriteriaName).Execute()

Returns a single Search Entry Criteria

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
    searchEntryCriteriaName := "searchEntryCriteriaName_example" // string | Name of the Search Entry Criteria to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchEntryCriteriaApi.GetSearchEntryCriteria(context.Background(), searchEntryCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchEntryCriteriaApi.GetSearchEntryCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchEntryCriteria`: AddSearchEntryCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchEntryCriteriaApi.GetSearchEntryCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchEntryCriteriaName** | **string** | Name of the Search Entry Criteria to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchEntryCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddSearchEntryCriteria200Response**](AddSearchEntryCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSearchEntryCriteria

> AddSearchEntryCriteria200Response UpdateSearchEntryCriteria(ctx, searchEntryCriteriaName).UpdateRequest(updateRequest).Execute()

Update an existing Search Entry Criteria by name

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
    searchEntryCriteriaName := "searchEntryCriteriaName_example" // string | Name of the Search Entry Criteria to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Search Entry Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchEntryCriteriaApi.UpdateSearchEntryCriteria(context.Background(), searchEntryCriteriaName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchEntryCriteriaApi.UpdateSearchEntryCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSearchEntryCriteria`: AddSearchEntryCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchEntryCriteriaApi.UpdateSearchEntryCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchEntryCriteriaName** | **string** | Name of the Search Entry Criteria to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSearchEntryCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Search Entry Criteria | 

### Return type

[**AddSearchEntryCriteria200Response**](AddSearchEntryCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

