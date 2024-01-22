# \ResultCriteriaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResultCriteria**](ResultCriteriaAPI.md#AddResultCriteria) | **Post** /result-criteria | Add a new Result Criteria to the config
[**DeleteResultCriteria**](ResultCriteriaAPI.md#DeleteResultCriteria) | **Delete** /result-criteria/{result-criteria-name} | Delete a Result Criteria
[**GetResultCriteria**](ResultCriteriaAPI.md#GetResultCriteria) | **Get** /result-criteria/{result-criteria-name} | Returns a single Result Criteria
[**ListResultCriteria**](ResultCriteriaAPI.md#ListResultCriteria) | **Get** /result-criteria | Returns a list of all Result Criteria objects
[**UpdateResultCriteria**](ResultCriteriaAPI.md#UpdateResultCriteria) | **Patch** /result-criteria/{result-criteria-name} | Update an existing Result Criteria by name



## AddResultCriteria

> AddResultCriteria200Response AddResultCriteria(ctx).AddResultCriteriaRequest(addResultCriteriaRequest).Execute()

Add a new Result Criteria to the config

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
    addResultCriteriaRequest := openapiclient.add_result_criteria_request{AddAggregateResultCriteriaRequest: openapiclient.NewAddAggregateResultCriteriaRequest([]openapiclient.EnumaggregateResultCriteriaSchemaUrn{openapiclient.Enumaggregate-result-criteriaSchemaUrn("urn:pingidentity:schemas:configuration:2.0:result-criteria:aggregate")}, "CriteriaName_example")} // AddResultCriteriaRequest | Create a new Result Criteria in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResultCriteriaAPI.AddResultCriteria(context.Background()).AddResultCriteriaRequest(addResultCriteriaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultCriteriaAPI.AddResultCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddResultCriteria`: AddResultCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `ResultCriteriaAPI.AddResultCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddResultCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addResultCriteriaRequest** | [**AddResultCriteriaRequest**](AddResultCriteriaRequest.md) | Create a new Result Criteria in the config | 

### Return type

[**AddResultCriteria200Response**](AddResultCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResultCriteria

> DeleteResultCriteria(ctx, resultCriteriaName).Execute()

Delete a Result Criteria

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
    resultCriteriaName := "resultCriteriaName_example" // string | Name of the Result Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResultCriteriaAPI.DeleteResultCriteria(context.Background(), resultCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultCriteriaAPI.DeleteResultCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultCriteriaName** | **string** | Name of the Result Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResultCriteriaRequest struct via the builder pattern


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


## GetResultCriteria

> AddResultCriteria200Response GetResultCriteria(ctx, resultCriteriaName).Execute()

Returns a single Result Criteria

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
    resultCriteriaName := "resultCriteriaName_example" // string | Name of the Result Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResultCriteriaAPI.GetResultCriteria(context.Background(), resultCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultCriteriaAPI.GetResultCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResultCriteria`: AddResultCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `ResultCriteriaAPI.GetResultCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultCriteriaName** | **string** | Name of the Result Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResultCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddResultCriteria200Response**](AddResultCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResultCriteria

> ResultCriteriaListResponse ListResultCriteria(ctx).Filter(filter).Execute()

Returns a list of all Result Criteria objects

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
    resp, r, err := apiClient.ResultCriteriaAPI.ListResultCriteria(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultCriteriaAPI.ListResultCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResultCriteria`: ResultCriteriaListResponse
    fmt.Fprintf(os.Stdout, "Response from `ResultCriteriaAPI.ListResultCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResultCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ResultCriteriaListResponse**](ResultCriteriaListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResultCriteria

> AddResultCriteria200Response UpdateResultCriteria(ctx, resultCriteriaName).UpdateRequest(updateRequest).Execute()

Update an existing Result Criteria by name

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
    resultCriteriaName := "resultCriteriaName_example" // string | Name of the Result Criteria
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Result Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResultCriteriaAPI.UpdateResultCriteria(context.Background(), resultCriteriaName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResultCriteriaAPI.UpdateResultCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResultCriteria`: AddResultCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `ResultCriteriaAPI.UpdateResultCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resultCriteriaName** | **string** | Name of the Result Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResultCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Result Criteria | 

### Return type

[**AddResultCriteria200Response**](AddResultCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

