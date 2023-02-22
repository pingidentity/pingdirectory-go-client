# \GlobalConfigurationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalConfiguration**](GlobalConfigurationApi.md#GetGlobalConfiguration) | **Get** /global-configuration | Returns a single Global Configuration
[**UpdateGlobalConfiguration**](GlobalConfigurationApi.md#UpdateGlobalConfiguration) | **Patch** /global-configuration | Update an existing Global Configuration by name



## GetGlobalConfiguration

> GlobalConfigurationResponse GetGlobalConfiguration(ctx).Execute()

Returns a single Global Configuration

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalConfigurationApi.GetGlobalConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationApi.GetGlobalConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalConfiguration`: GlobalConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `GlobalConfigurationApi.GetGlobalConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalConfigurationRequest struct via the builder pattern


### Return type

[**GlobalConfigurationResponse**](GlobalConfigurationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalConfiguration

> GlobalConfigurationResponse UpdateGlobalConfiguration(ctx).UpdateRequest(updateRequest).Execute()

Update an existing Global Configuration by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Global Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalConfigurationApi.UpdateGlobalConfiguration(context.Background()).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationApi.UpdateGlobalConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGlobalConfiguration`: GlobalConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `GlobalConfigurationApi.UpdateGlobalConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Global Configuration | 

### Return type

[**GlobalConfigurationResponse**](GlobalConfigurationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

