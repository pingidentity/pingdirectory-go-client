# \LocationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLocation**](LocationApi.md#AddLocation) | **Post** /locations | Add a new Location to the config
[**DeleteLocation**](LocationApi.md#DeleteLocation) | **Delete** /locations/{location-name} | Delete a Location
[**GetLocation**](LocationApi.md#GetLocation) | **Get** /locations/{location-name} | Returns a single Location
[**UpdateLocation**](LocationApi.md#UpdateLocation) | **Patch** /locations/{location-name} | Update an existing Location by name



## AddLocation

> LocationResponse AddLocation(ctx).AddLocationRequest(addLocationRequest).Execute()

Add a new Location to the config

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
    addLocationRequest := *openapiclient.NewAddLocationRequest("LocationName_example") // AddLocationRequest | Create a new Location in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationApi.AddLocation(context.Background()).AddLocationRequest(addLocationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationApi.AddLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLocation`: LocationResponse
    fmt.Fprintf(os.Stdout, "Response from `LocationApi.AddLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLocationRequest** | [**AddLocationRequest**](AddLocationRequest.md) | Create a new Location in the config | 

### Return type

[**LocationResponse**](LocationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLocation

> DeleteLocation(ctx, locationName).Execute()

Delete a Location

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
    locationName := "locationName_example" // string | Name of the Location to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationApi.DeleteLocation(context.Background(), locationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationApi.DeleteLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationName** | **string** | Name of the Location to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocationRequest struct via the builder pattern


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


## GetLocation

> LocationResponse GetLocation(ctx, locationName).Execute()

Returns a single Location

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
    locationName := "locationName_example" // string | Name of the Location to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationApi.GetLocation(context.Background(), locationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationApi.GetLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocation`: LocationResponse
    fmt.Fprintf(os.Stdout, "Response from `LocationApi.GetLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationName** | **string** | Name of the Location to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LocationResponse**](LocationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocation

> LocationResponse UpdateLocation(ctx, locationName).UpdateRequest(updateRequest).Execute()

Update an existing Location by name

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
    locationName := "locationName_example" // string | Name of the Location to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Location

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationApi.UpdateLocation(context.Background(), locationName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationApi.UpdateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLocation`: LocationResponse
    fmt.Fprintf(os.Stdout, "Response from `LocationApi.UpdateLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationName** | **string** | Name of the Location to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Location | 

### Return type

[**LocationResponse**](LocationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

