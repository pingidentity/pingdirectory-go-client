# \SearchReferenceCriteriaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSearchReferenceCriteria**](SearchReferenceCriteriaAPI.md#AddSearchReferenceCriteria) | **Post** /search-reference-criteria | Add a new Search Reference Criteria to the config
[**DeleteSearchReferenceCriteria**](SearchReferenceCriteriaAPI.md#DeleteSearchReferenceCriteria) | **Delete** /search-reference-criteria/{search-reference-criteria-name} | Delete a Search Reference Criteria
[**GetSearchReferenceCriteria**](SearchReferenceCriteriaAPI.md#GetSearchReferenceCriteria) | **Get** /search-reference-criteria/{search-reference-criteria-name} | Returns a single Search Reference Criteria
[**ListSearchReferenceCriteria**](SearchReferenceCriteriaAPI.md#ListSearchReferenceCriteria) | **Get** /search-reference-criteria | Returns a list of all Search Reference Criteria objects
[**UpdateSearchReferenceCriteria**](SearchReferenceCriteriaAPI.md#UpdateSearchReferenceCriteria) | **Patch** /search-reference-criteria/{search-reference-criteria-name} | Update an existing Search Reference Criteria by name



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
    addSearchReferenceCriteriaRequest := openapiclient.add_search_reference_criteria_request{AddAggregateSearchReferenceCriteriaRequest: openapiclient.NewAddAggregateSearchReferenceCriteriaRequest([]openapiclient.EnumaggregateSearchReferenceCriteriaSchemaUrn{openapiclient.Enumaggregate-search-reference-criteriaSchemaUrn("urn:pingidentity:schemas:configuration:2.0:search-reference-criteria:aggregate")}, "CriteriaName_example")} // AddSearchReferenceCriteriaRequest | Create a new Search Reference Criteria in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchReferenceCriteriaAPI.AddSearchReferenceCriteria(context.Background()).AddSearchReferenceCriteriaRequest(addSearchReferenceCriteriaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchReferenceCriteriaAPI.AddSearchReferenceCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSearchReferenceCriteria`: AddSearchReferenceCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchReferenceCriteriaAPI.AddSearchReferenceCriteria`: %v\n", resp)
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
    r, err := apiClient.SearchReferenceCriteriaAPI.DeleteSearchReferenceCriteria(context.Background(), searchReferenceCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchReferenceCriteriaAPI.DeleteSearchReferenceCriteria``: %v\n", err)
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
    resp, r, err := apiClient.SearchReferenceCriteriaAPI.GetSearchReferenceCriteria(context.Background(), searchReferenceCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchReferenceCriteriaAPI.GetSearchReferenceCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchReferenceCriteria`: AddSearchReferenceCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchReferenceCriteriaAPI.GetSearchReferenceCriteria`: %v\n", resp)
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


## ListSearchReferenceCriteria

> SearchReferenceCriteriaListResponse ListSearchReferenceCriteria(ctx).Filter(filter).Execute()

Returns a list of all Search Reference Criteria objects

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
    resp, r, err := apiClient.SearchReferenceCriteriaAPI.ListSearchReferenceCriteria(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchReferenceCriteriaAPI.ListSearchReferenceCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSearchReferenceCriteria`: SearchReferenceCriteriaListResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchReferenceCriteriaAPI.ListSearchReferenceCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSearchReferenceCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**SearchReferenceCriteriaListResponse**](SearchReferenceCriteriaListResponse.md)

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
    resp, r, err := apiClient.SearchReferenceCriteriaAPI.UpdateSearchReferenceCriteria(context.Background(), searchReferenceCriteriaName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchReferenceCriteriaAPI.UpdateSearchReferenceCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSearchReferenceCriteria`: AddSearchReferenceCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchReferenceCriteriaAPI.UpdateSearchReferenceCriteria`: %v\n", resp)
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

