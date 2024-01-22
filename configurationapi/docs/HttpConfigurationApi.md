# \HttpConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHttpConfiguration**](HttpConfigurationAPI.md#GetHttpConfiguration) | **Get** /http-configuration | Returns a single HTTP Configuration
[**UpdateHttpConfiguration**](HttpConfigurationAPI.md#UpdateHttpConfiguration) | **Patch** /http-configuration | Update an existing HTTP Configuration by name



## GetHttpConfiguration

> HttpConfigurationResponse GetHttpConfiguration(ctx).Execute()

Returns a single HTTP Configuration

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
    resp, r, err := apiClient.HttpConfigurationAPI.GetHttpConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigurationAPI.GetHttpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHttpConfiguration`: HttpConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `HttpConfigurationAPI.GetHttpConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHttpConfigurationRequest struct via the builder pattern


### Return type

[**HttpConfigurationResponse**](HttpConfigurationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHttpConfiguration

> HttpConfigurationResponse UpdateHttpConfiguration(ctx).UpdateRequest(updateRequest).Execute()

Update an existing HTTP Configuration by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing HTTP Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpConfigurationAPI.UpdateHttpConfiguration(context.Background()).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpConfigurationAPI.UpdateHttpConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHttpConfiguration`: HttpConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `HttpConfigurationAPI.UpdateHttpConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHttpConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing HTTP Configuration | 

### Return type

[**HttpConfigurationResponse**](HttpConfigurationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

