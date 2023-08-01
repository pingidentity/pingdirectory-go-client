# \ObscuredValueApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddObscuredValue**](ObscuredValueApi.md#AddObscuredValue) | **Post** /obscured-values | Add a new Obscured Value to the config
[**DeleteObscuredValue**](ObscuredValueApi.md#DeleteObscuredValue) | **Delete** /obscured-values/{obscured-value-name} | Delete a Obscured Value
[**GetObscuredValue**](ObscuredValueApi.md#GetObscuredValue) | **Get** /obscured-values/{obscured-value-name} | Returns a single Obscured Value
[**ListObscuredValues**](ObscuredValueApi.md#ListObscuredValues) | **Get** /obscured-values | Returns a list of all Obscured Value objects
[**UpdateObscuredValue**](ObscuredValueApi.md#UpdateObscuredValue) | **Patch** /obscured-values/{obscured-value-name} | Update an existing Obscured Value by name



## AddObscuredValue

> ObscuredValueResponse AddObscuredValue(ctx).AddObscuredValueRequest(addObscuredValueRequest).Execute()

Add a new Obscured Value to the config

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
    addObscuredValueRequest := *openapiclient.NewAddObscuredValueRequest("ValueName_example", "ObscuredValue_example") // AddObscuredValueRequest | Create a new Obscured Value in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObscuredValueApi.AddObscuredValue(context.Background()).AddObscuredValueRequest(addObscuredValueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObscuredValueApi.AddObscuredValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddObscuredValue`: ObscuredValueResponse
    fmt.Fprintf(os.Stdout, "Response from `ObscuredValueApi.AddObscuredValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddObscuredValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addObscuredValueRequest** | [**AddObscuredValueRequest**](AddObscuredValueRequest.md) | Create a new Obscured Value in the config | 

### Return type

[**ObscuredValueResponse**](ObscuredValueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObscuredValue

> DeleteObscuredValue(ctx, obscuredValueName).Execute()

Delete a Obscured Value

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
    obscuredValueName := "obscuredValueName_example" // string | Name of the Obscured Value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ObscuredValueApi.DeleteObscuredValue(context.Background(), obscuredValueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObscuredValueApi.DeleteObscuredValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**obscuredValueName** | **string** | Name of the Obscured Value | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObscuredValueRequest struct via the builder pattern


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


## GetObscuredValue

> ObscuredValueResponse GetObscuredValue(ctx, obscuredValueName).Execute()

Returns a single Obscured Value

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
    obscuredValueName := "obscuredValueName_example" // string | Name of the Obscured Value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObscuredValueApi.GetObscuredValue(context.Background(), obscuredValueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObscuredValueApi.GetObscuredValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObscuredValue`: ObscuredValueResponse
    fmt.Fprintf(os.Stdout, "Response from `ObscuredValueApi.GetObscuredValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**obscuredValueName** | **string** | Name of the Obscured Value | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObscuredValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObscuredValueResponse**](ObscuredValueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObscuredValues

> ObscuredValueListResponse ListObscuredValues(ctx).Filter(filter).Execute()

Returns a list of all Obscured Value objects

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
    resp, r, err := apiClient.ObscuredValueApi.ListObscuredValues(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObscuredValueApi.ListObscuredValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListObscuredValues`: ObscuredValueListResponse
    fmt.Fprintf(os.Stdout, "Response from `ObscuredValueApi.ListObscuredValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListObscuredValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ObscuredValueListResponse**](ObscuredValueListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObscuredValue

> ObscuredValueResponse UpdateObscuredValue(ctx, obscuredValueName).UpdateRequest(updateRequest).Execute()

Update an existing Obscured Value by name

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
    obscuredValueName := "obscuredValueName_example" // string | Name of the Obscured Value
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Obscured Value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObscuredValueApi.UpdateObscuredValue(context.Background(), obscuredValueName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObscuredValueApi.UpdateObscuredValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateObscuredValue`: ObscuredValueResponse
    fmt.Fprintf(os.Stdout, "Response from `ObscuredValueApi.UpdateObscuredValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**obscuredValueName** | **string** | Name of the Obscured Value | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObscuredValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Obscured Value | 

### Return type

[**ObscuredValueResponse**](ObscuredValueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

