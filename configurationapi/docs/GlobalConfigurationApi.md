# \GlobalConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalConfiguration**](GlobalConfigurationAPI.md#GetGlobalConfiguration) | **Get** /global-configuration | Returns a single Global Configuration
[**UpdateGlobalConfiguration**](GlobalConfigurationAPI.md#UpdateGlobalConfiguration) | **Patch** /global-configuration | Update an existing Global Configuration by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalConfigurationAPI.GetGlobalConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.GetGlobalConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalConfiguration`: GlobalConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `GlobalConfigurationAPI.GetGlobalConfiguration`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Global Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalConfigurationAPI.UpdateGlobalConfiguration(context.Background()).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalConfigurationAPI.UpdateGlobalConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGlobalConfiguration`: GlobalConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `GlobalConfigurationAPI.UpdateGlobalConfiguration`: %v\n", resp)
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

