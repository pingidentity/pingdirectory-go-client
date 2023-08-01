# \UncachedEntryCriteriaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUncachedEntryCriteria**](UncachedEntryCriteriaApi.md#AddUncachedEntryCriteria) | **Post** /uncached-entry-criteria | Add a new Uncached Entry Criteria to the config
[**DeleteUncachedEntryCriteria**](UncachedEntryCriteriaApi.md#DeleteUncachedEntryCriteria) | **Delete** /uncached-entry-criteria/{uncached-entry-criteria-name} | Delete a Uncached Entry Criteria
[**GetUncachedEntryCriteria**](UncachedEntryCriteriaApi.md#GetUncachedEntryCriteria) | **Get** /uncached-entry-criteria/{uncached-entry-criteria-name} | Returns a single Uncached Entry Criteria
[**ListUncachedEntryCriteria**](UncachedEntryCriteriaApi.md#ListUncachedEntryCriteria) | **Get** /uncached-entry-criteria | Returns a list of all Uncached Entry Criteria objects
[**UpdateUncachedEntryCriteria**](UncachedEntryCriteriaApi.md#UpdateUncachedEntryCriteria) | **Patch** /uncached-entry-criteria/{uncached-entry-criteria-name} | Update an existing Uncached Entry Criteria by name



## AddUncachedEntryCriteria

> AddUncachedEntryCriteria200Response AddUncachedEntryCriteria(ctx).AddUncachedEntryCriteriaRequest(addUncachedEntryCriteriaRequest).Execute()

Add a new Uncached Entry Criteria to the config

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
    addUncachedEntryCriteriaRequest := openapiclient.add_uncached_entry_criteria_request{AddDefaultUncachedEntryCriteriaRequest: openapiclient.NewAddDefaultUncachedEntryCriteriaRequest("CriteriaName_example", []openapiclient.EnumdefaultUncachedEntryCriteriaSchemaUrn{openapiclient.Enumdefault-uncached-entry-criteriaSchemaUrn("urn:pingidentity:schemas:configuration:2.0:uncached-entry-criteria:default")}, false)} // AddUncachedEntryCriteriaRequest | Create a new Uncached Entry Criteria in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UncachedEntryCriteriaApi.AddUncachedEntryCriteria(context.Background()).AddUncachedEntryCriteriaRequest(addUncachedEntryCriteriaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UncachedEntryCriteriaApi.AddUncachedEntryCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUncachedEntryCriteria`: AddUncachedEntryCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `UncachedEntryCriteriaApi.AddUncachedEntryCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUncachedEntryCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addUncachedEntryCriteriaRequest** | [**AddUncachedEntryCriteriaRequest**](AddUncachedEntryCriteriaRequest.md) | Create a new Uncached Entry Criteria in the config | 

### Return type

[**AddUncachedEntryCriteria200Response**](AddUncachedEntryCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUncachedEntryCriteria

> DeleteUncachedEntryCriteria(ctx, uncachedEntryCriteriaName).Execute()

Delete a Uncached Entry Criteria

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
    uncachedEntryCriteriaName := "uncachedEntryCriteriaName_example" // string | Name of the Uncached Entry Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UncachedEntryCriteriaApi.DeleteUncachedEntryCriteria(context.Background(), uncachedEntryCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UncachedEntryCriteriaApi.DeleteUncachedEntryCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uncachedEntryCriteriaName** | **string** | Name of the Uncached Entry Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUncachedEntryCriteriaRequest struct via the builder pattern


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


## GetUncachedEntryCriteria

> AddUncachedEntryCriteria200Response GetUncachedEntryCriteria(ctx, uncachedEntryCriteriaName).Execute()

Returns a single Uncached Entry Criteria

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
    uncachedEntryCriteriaName := "uncachedEntryCriteriaName_example" // string | Name of the Uncached Entry Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UncachedEntryCriteriaApi.GetUncachedEntryCriteria(context.Background(), uncachedEntryCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UncachedEntryCriteriaApi.GetUncachedEntryCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUncachedEntryCriteria`: AddUncachedEntryCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `UncachedEntryCriteriaApi.GetUncachedEntryCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uncachedEntryCriteriaName** | **string** | Name of the Uncached Entry Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUncachedEntryCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddUncachedEntryCriteria200Response**](AddUncachedEntryCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUncachedEntryCriteria

> UncachedEntryCriteriaListResponse ListUncachedEntryCriteria(ctx).Filter(filter).Execute()

Returns a list of all Uncached Entry Criteria objects

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
    resp, r, err := apiClient.UncachedEntryCriteriaApi.ListUncachedEntryCriteria(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UncachedEntryCriteriaApi.ListUncachedEntryCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUncachedEntryCriteria`: UncachedEntryCriteriaListResponse
    fmt.Fprintf(os.Stdout, "Response from `UncachedEntryCriteriaApi.ListUncachedEntryCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUncachedEntryCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**UncachedEntryCriteriaListResponse**](UncachedEntryCriteriaListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUncachedEntryCriteria

> AddUncachedEntryCriteria200Response UpdateUncachedEntryCriteria(ctx, uncachedEntryCriteriaName).UpdateRequest(updateRequest).Execute()

Update an existing Uncached Entry Criteria by name

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
    uncachedEntryCriteriaName := "uncachedEntryCriteriaName_example" // string | Name of the Uncached Entry Criteria
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Uncached Entry Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UncachedEntryCriteriaApi.UpdateUncachedEntryCriteria(context.Background(), uncachedEntryCriteriaName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UncachedEntryCriteriaApi.UpdateUncachedEntryCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUncachedEntryCriteria`: AddUncachedEntryCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `UncachedEntryCriteriaApi.UpdateUncachedEntryCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uncachedEntryCriteriaName** | **string** | Name of the Uncached Entry Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUncachedEntryCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Uncached Entry Criteria | 

### Return type

[**AddUncachedEntryCriteria200Response**](AddUncachedEntryCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

