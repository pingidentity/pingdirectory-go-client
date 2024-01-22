# \RequestCriteriaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRequestCriteria**](RequestCriteriaAPI.md#AddRequestCriteria) | **Post** /request-criteria | Add a new Request Criteria to the config
[**DeleteRequestCriteria**](RequestCriteriaAPI.md#DeleteRequestCriteria) | **Delete** /request-criteria/{request-criteria-name} | Delete a Request Criteria
[**GetRequestCriteria**](RequestCriteriaAPI.md#GetRequestCriteria) | **Get** /request-criteria/{request-criteria-name} | Returns a single Request Criteria
[**ListRequestCriteria**](RequestCriteriaAPI.md#ListRequestCriteria) | **Get** /request-criteria | Returns a list of all Request Criteria objects
[**UpdateRequestCriteria**](RequestCriteriaAPI.md#UpdateRequestCriteria) | **Patch** /request-criteria/{request-criteria-name} | Update an existing Request Criteria by name



## AddRequestCriteria

> AddRequestCriteria200Response AddRequestCriteria(ctx).AddRequestCriteriaRequest(addRequestCriteriaRequest).Execute()

Add a new Request Criteria to the config

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
    addRequestCriteriaRequest := openapiclient.add_request_criteria_request{AddAggregateRequestCriteriaRequest: openapiclient.NewAddAggregateRequestCriteriaRequest([]openapiclient.EnumaggregateRequestCriteriaSchemaUrn{openapiclient.Enumaggregate-request-criteriaSchemaUrn("urn:pingidentity:schemas:configuration:2.0:request-criteria:aggregate")}, "CriteriaName_example")} // AddRequestCriteriaRequest | Create a new Request Criteria in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestCriteriaAPI.AddRequestCriteria(context.Background()).AddRequestCriteriaRequest(addRequestCriteriaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestCriteriaAPI.AddRequestCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRequestCriteria`: AddRequestCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `RequestCriteriaAPI.AddRequestCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRequestCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addRequestCriteriaRequest** | [**AddRequestCriteriaRequest**](AddRequestCriteriaRequest.md) | Create a new Request Criteria in the config | 

### Return type

[**AddRequestCriteria200Response**](AddRequestCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRequestCriteria

> DeleteRequestCriteria(ctx, requestCriteriaName).Execute()

Delete a Request Criteria

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
    requestCriteriaName := "requestCriteriaName_example" // string | Name of the Request Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RequestCriteriaAPI.DeleteRequestCriteria(context.Background(), requestCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestCriteriaAPI.DeleteRequestCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestCriteriaName** | **string** | Name of the Request Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestCriteriaRequest struct via the builder pattern


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


## GetRequestCriteria

> AddRequestCriteria200Response GetRequestCriteria(ctx, requestCriteriaName).Execute()

Returns a single Request Criteria

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
    requestCriteriaName := "requestCriteriaName_example" // string | Name of the Request Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestCriteriaAPI.GetRequestCriteria(context.Background(), requestCriteriaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestCriteriaAPI.GetRequestCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestCriteria`: AddRequestCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `RequestCriteriaAPI.GetRequestCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestCriteriaName** | **string** | Name of the Request Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddRequestCriteria200Response**](AddRequestCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequestCriteria

> RequestCriteriaListResponse ListRequestCriteria(ctx).Filter(filter).Execute()

Returns a list of all Request Criteria objects

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
    resp, r, err := apiClient.RequestCriteriaAPI.ListRequestCriteria(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestCriteriaAPI.ListRequestCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestCriteria`: RequestCriteriaListResponse
    fmt.Fprintf(os.Stdout, "Response from `RequestCriteriaAPI.ListRequestCriteria`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequestCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**RequestCriteriaListResponse**](RequestCriteriaListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRequestCriteria

> AddRequestCriteria200Response UpdateRequestCriteria(ctx, requestCriteriaName).UpdateRequest(updateRequest).Execute()

Update an existing Request Criteria by name

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
    requestCriteriaName := "requestCriteriaName_example" // string | Name of the Request Criteria
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Request Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestCriteriaAPI.UpdateRequestCriteria(context.Background(), requestCriteriaName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestCriteriaAPI.UpdateRequestCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRequestCriteria`: AddRequestCriteria200Response
    fmt.Fprintf(os.Stdout, "Response from `RequestCriteriaAPI.UpdateRequestCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestCriteriaName** | **string** | Name of the Request Criteria | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequestCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Request Criteria | 

### Return type

[**AddRequestCriteria200Response**](AddRequestCriteria200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

