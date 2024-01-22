# \UncachedAttributeCriteriaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUncachedAttributeCriteria**](UncachedAttributeCriteriaAPI.md#AddUncachedAttributeCriteria) | **Post** /uncached-attribute-criteria | Add a new Uncached Attribute Criteria to the config
[**DeleteUncachedAttributeCriteria**](UncachedAttributeCriteriaAPI.md#DeleteUncachedAttributeCriteria) | **Delete** /uncached-attribute-criteria/{uncached-attribute-criteria-name} | Delete a Uncached Attribute Criteria
[**GetUncachedAttributeCriteria**](UncachedAttributeCriteriaAPI.md#GetUncachedAttributeCriteria) | **Get** /uncached-attribute-criteria/{uncached-attribute-criteria-name} | Returns a single Uncached Attribute Criteria
[**ListUncachedAttributeCriteria**](UncachedAttributeCriteriaAPI.md#ListUncachedAttributeCriteria) | **Get** /uncached-attribute-criteria | Returns a list of all Uncached Attribute Criteria objects
[**UpdateUncachedAttributeCriteria**](UncachedAttributeCriteriaAPI.md#UpdateUncachedAttributeCriteria) | **Patch** /uncached-attribute-criteria/{uncached-attribute-criteria-name} | Update an existing Uncached Attribute Criteria by name



## AddUncachedAttributeCriteria

> AddUncachedAttributeCriteria200Response AddUncachedAttributeCriteria(ctx).AddUncachedAttributeCriteriaRequest(addUncachedAttributeCriteriaRequest).Execute()

Add a new Uncached Attribute Criteria to the config

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
    addUncachedAttributeCriteriaRequest := openapiclient.add_uncached_attribute_criteria_request{AddDefaultUncachedAttributeCriteriaRequest: openapiclient.NewAddDefaultUncachedAttributeCriteriaRequest([]openapiclient.EnumdefaultUncachedAttributeCriteriaSchemaUrn{openapiclient.Enumdefault-uncached-attribute-criteriaSchemaUrn("urn:pingidentity:schemas:configuration:2.0:uncached-attribute-criteria:default")}, false, "CriteriaName_example")} // AddUncachedAttributeCriteriaRequest | Create a new Uncached Attribute Criteria in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UncachedAttributeCriteriaAPI.AddUncachedAttributeCriteria(context.Background()).AddUncachedAttributeCriteriaRequest(addUncachedAttributeCriteriaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UncachedAttributeCriteriaAPI.AddUncachedAttributeCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUncachedAttributeCriteria`: AddUncachedAttributeCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `UncachedAttributeCriteriaAPI.AddUncachedAttributeCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUncachedAttributeCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addUncachedAttributeCriteriaRequest** | [**AddUncachedAttributeCriteriaRequest**](AddUncachedAttributeCriteriaRequest.md) | Create a new Uncached Attribute Criteria in the config | 

### Return type

[**AddUncachedAttributeCriteria200Response**](AddUncachedAttributeCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUncachedAttributeCriteria

> DeleteUncachedAttributeCriteria(ctx, uncachedAttributeCriteriaName).Execute()

Delete a Uncached Attribute Criteria

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
    uncachedAttributeCriteriaName := "uncachedAttributeCriteriaName_example" // string | Name of the Uncached Attribute Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UncachedAttributeCriteriaAPI.DeleteUncachedAttributeCriteria(context.Background(), uncachedAttributeCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UncachedAttributeCriteriaAPI.DeleteUncachedAttributeCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uncachedAttributeCriteriaName** | **string** | Name of the Uncached Attribute Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUncachedAttributeCriteriaRequest struct via the builder pattern


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


## GetUncachedAttributeCriteria

> AddUncachedAttributeCriteria200Response GetUncachedAttributeCriteria(ctx, uncachedAttributeCriteriaName).Execute()

Returns a single Uncached Attribute Criteria

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
    uncachedAttributeCriteriaName := "uncachedAttributeCriteriaName_example" // string | Name of the Uncached Attribute Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UncachedAttributeCriteriaAPI.GetUncachedAttributeCriteria(context.Background(), uncachedAttributeCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UncachedAttributeCriteriaAPI.GetUncachedAttributeCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUncachedAttributeCriteria`: AddUncachedAttributeCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `UncachedAttributeCriteriaAPI.GetUncachedAttributeCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uncachedAttributeCriteriaName** | **string** | Name of the Uncached Attribute Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUncachedAttributeCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddUncachedAttributeCriteria200Response**](AddUncachedAttributeCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUncachedAttributeCriteria

> UncachedAttributeCriteriaListResponse ListUncachedAttributeCriteria(ctx).Filter(filter).Execute()

Returns a list of all Uncached Attribute Criteria objects

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
    resp, r, err := apiClient.UncachedAttributeCriteriaAPI.ListUncachedAttributeCriteria(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UncachedAttributeCriteriaAPI.ListUncachedAttributeCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUncachedAttributeCriteria`: UncachedAttributeCriteriaListResponse
    fmt.Fprintf(os.Stdout, "Response from `UncachedAttributeCriteriaAPI.ListUncachedAttributeCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUncachedAttributeCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**UncachedAttributeCriteriaListResponse**](UncachedAttributeCriteriaListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUncachedAttributeCriteria

> AddUncachedAttributeCriteria200Response UpdateUncachedAttributeCriteria(ctx, uncachedAttributeCriteriaName).UpdateRequest(updateRequest).Execute()

Update an existing Uncached Attribute Criteria by name

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
    uncachedAttributeCriteriaName := "uncachedAttributeCriteriaName_example" // string | Name of the Uncached Attribute Criteria
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Uncached Attribute Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UncachedAttributeCriteriaAPI.UpdateUncachedAttributeCriteria(context.Background(), uncachedAttributeCriteriaName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UncachedAttributeCriteriaAPI.UpdateUncachedAttributeCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUncachedAttributeCriteria`: AddUncachedAttributeCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `UncachedAttributeCriteriaAPI.UpdateUncachedAttributeCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uncachedAttributeCriteriaName** | **string** | Name of the Uncached Attribute Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUncachedAttributeCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Uncached Attribute Criteria | 

### Return type

[**AddUncachedAttributeCriteria200Response**](AddUncachedAttributeCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

