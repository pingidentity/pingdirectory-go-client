# \CustomLoggedStatsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomLoggedStats**](CustomLoggedStatsAPI.md#AddCustomLoggedStats) | **Post** /plugin-root/plugins/{plugin-name}/custom-logged-stats | Add a new Custom Logged Stats to the config
[**DeleteCustomLoggedStats**](CustomLoggedStatsAPI.md#DeleteCustomLoggedStats) | **Delete** /plugin-root/plugins/{plugin-name}/custom-logged-stats/{custom-logged-stats-name} | Delete a Custom Logged Stats
[**GetCustomLoggedStats**](CustomLoggedStatsAPI.md#GetCustomLoggedStats) | **Get** /plugin-root/plugins/{plugin-name}/custom-logged-stats/{custom-logged-stats-name} | Returns a single Custom Logged Stats
[**ListCustomLoggedStats**](CustomLoggedStatsAPI.md#ListCustomLoggedStats) | **Get** /plugin-root/plugins/{plugin-name}/custom-logged-stats | Returns a list of all Custom Logged Stats objects
[**UpdateCustomLoggedStats**](CustomLoggedStatsAPI.md#UpdateCustomLoggedStats) | **Patch** /plugin-root/plugins/{plugin-name}/custom-logged-stats/{custom-logged-stats-name} | Update an existing Custom Logged Stats by name



## AddCustomLoggedStats

> CustomLoggedStatsResponse AddCustomLoggedStats(ctx, pluginName).AddCustomLoggedStatsRequest(addCustomLoggedStatsRequest).Execute()

Add a new Custom Logged Stats to the config

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
    pluginName := "pluginName_example" // string | Name of the Plugin
    addCustomLoggedStatsRequest := *openapiclient.NewAddCustomLoggedStatsRequest([]openapiclient.EnumcustomLoggedStatsSchemaUrn{openapiclient.Enumcustom-logged-statsSchemaUrn("urn:pingidentity:schemas:configuration:2.0:custom-logged-stats:custom")}, "MonitorObjectclass_example", []string{"AttributeToLog_example"}, []openapiclient.EnumcustomLoggedStatsStatisticTypeProp{openapiclient.Enumcustom-logged-stats-statisticTypeProp("raw")}, "StatsName_example") // AddCustomLoggedStatsRequest | Create a new Custom Logged Stats in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomLoggedStatsAPI.AddCustomLoggedStats(context.Background(), pluginName).AddCustomLoggedStatsRequest(addCustomLoggedStatsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomLoggedStatsAPI.AddCustomLoggedStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCustomLoggedStats`: CustomLoggedStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomLoggedStatsAPI.AddCustomLoggedStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCustomLoggedStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addCustomLoggedStatsRequest** | [**AddCustomLoggedStatsRequest**](AddCustomLoggedStatsRequest.md) | Create a new Custom Logged Stats in the config | 

### Return type

[**CustomLoggedStatsResponse**](CustomLoggedStatsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomLoggedStats

> DeleteCustomLoggedStats(ctx, customLoggedStatsName, pluginName).Execute()

Delete a Custom Logged Stats

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
    customLoggedStatsName := "customLoggedStatsName_example" // string | Name of the Custom Logged Stats
    pluginName := "pluginName_example" // string | Name of the Plugin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomLoggedStatsAPI.DeleteCustomLoggedStats(context.Background(), customLoggedStatsName, pluginName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomLoggedStatsAPI.DeleteCustomLoggedStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customLoggedStatsName** | **string** | Name of the Custom Logged Stats | 
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomLoggedStatsRequest struct via the builder pattern


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


## GetCustomLoggedStats

> CustomLoggedStatsResponse GetCustomLoggedStats(ctx, customLoggedStatsName, pluginName).Execute()

Returns a single Custom Logged Stats

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
    customLoggedStatsName := "customLoggedStatsName_example" // string | Name of the Custom Logged Stats
    pluginName := "pluginName_example" // string | Name of the Plugin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomLoggedStatsAPI.GetCustomLoggedStats(context.Background(), customLoggedStatsName, pluginName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomLoggedStatsAPI.GetCustomLoggedStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomLoggedStats`: CustomLoggedStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomLoggedStatsAPI.GetCustomLoggedStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customLoggedStatsName** | **string** | Name of the Custom Logged Stats | 
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomLoggedStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomLoggedStatsResponse**](CustomLoggedStatsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomLoggedStats

> CustomLoggedStatsListResponse ListCustomLoggedStats(ctx, pluginName).Filter(filter).Execute()

Returns a list of all Custom Logged Stats objects

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
    pluginName := "pluginName_example" // string | Name of the Plugin
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomLoggedStatsAPI.ListCustomLoggedStats(context.Background(), pluginName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomLoggedStatsAPI.ListCustomLoggedStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomLoggedStats`: CustomLoggedStatsListResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomLoggedStatsAPI.ListCustomLoggedStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomLoggedStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**CustomLoggedStatsListResponse**](CustomLoggedStatsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomLoggedStats

> CustomLoggedStatsResponse UpdateCustomLoggedStats(ctx, customLoggedStatsName, pluginName).UpdateRequest(updateRequest).Execute()

Update an existing Custom Logged Stats by name

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
    customLoggedStatsName := "customLoggedStatsName_example" // string | Name of the Custom Logged Stats
    pluginName := "pluginName_example" // string | Name of the Plugin
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Custom Logged Stats

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomLoggedStatsAPI.UpdateCustomLoggedStats(context.Background(), customLoggedStatsName, pluginName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomLoggedStatsAPI.UpdateCustomLoggedStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomLoggedStats`: CustomLoggedStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomLoggedStatsAPI.UpdateCustomLoggedStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customLoggedStatsName** | **string** | Name of the Custom Logged Stats | 
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomLoggedStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Custom Logged Stats | 

### Return type

[**CustomLoggedStatsResponse**](CustomLoggedStatsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

