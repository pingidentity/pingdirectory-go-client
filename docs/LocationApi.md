# \LocationApi

All URIs are relative to *https://localhost:1443/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLocation**](LocationApi.md#AddLocation) | **Post** /locations | Add a new location to the config
[**DeleteLocation**](LocationApi.md#DeleteLocation) | **Delete** /locations/{locationName} | Deletes a location
[**GetLocation**](LocationApi.md#GetLocation) | **Get** /locations/{locationName} | Find location by name
[**UpdateLocation**](LocationApi.md#UpdateLocation) | **Patch** /locations/{locationName} | Update an existing location



## AddLocation

> AddLocation200Response AddLocation(ctx).AddLocationRequest(addLocationRequest).Execute()

Add a new location to the config



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
    addLocationRequest := *openapiclient.NewAddLocationRequest("LocationName_example") // AddLocationRequest | Create a new location in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationApi.AddLocation(context.Background()).AddLocationRequest(addLocationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationApi.AddLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLocation`: AddLocation200Response
    fmt.Fprintf(os.Stdout, "Response from `LocationApi.AddLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLocationRequest** | [**AddLocationRequest**](AddLocationRequest.md) | Create a new location in the config | 

### Return type

[**AddLocation200Response**](AddLocation200Response.md)

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

Deletes a location



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
    locationName := "locationName_example" // string | Name of the location to be read

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
**locationName** | **string** | Name of the location to be read | 

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

> GetLocation200Response GetLocation(ctx, locationName).Execute()

Find location by name



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
    locationName := "locationName_example" // string | Name of the location to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationApi.GetLocation(context.Background(), locationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationApi.GetLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocation`: GetLocation200Response
    fmt.Fprintf(os.Stdout, "Response from `LocationApi.GetLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationName** | **string** | Name of the location to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLocation200Response**](GetLocation200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocation

> AddLocation200Response UpdateLocation(ctx, locationName).UpdateLocationRequest(updateLocationRequest).Execute()

Update an existing location



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
    locationName := "locationName_example" // string | Name of the location that needs to be updated
    updateLocationRequest := *openapiclient.NewUpdateLocationRequest([]openapiclient.Operation{*openapiclient.NewOperation()}) // UpdateLocationRequest | Update an existing location

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationApi.UpdateLocation(context.Background(), locationName).UpdateLocationRequest(updateLocationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationApi.UpdateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLocation`: AddLocation200Response
    fmt.Fprintf(os.Stdout, "Response from `LocationApi.UpdateLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationName** | **string** | Name of the location that needs to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLocationRequest** | [**UpdateLocationRequest**](UpdateLocationRequest.md) | Update an existing location | 

### Return type

[**AddLocation200Response**](AddLocation200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

