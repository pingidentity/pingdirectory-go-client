# \EntryCounterPluginCriteriaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEntryCounterPluginCriteria**](EntryCounterPluginCriteriaAPI.md#AddEntryCounterPluginCriteria) | **Post** /plugin-root/plugins/{plugin-name}/entry-counter-criteria | Add a new Entry Counter Plugin Criteria to the config
[**DeleteEntryCounterPluginCriteria**](EntryCounterPluginCriteriaAPI.md#DeleteEntryCounterPluginCriteria) | **Delete** /plugin-root/plugins/{plugin-name}/entry-counter-criteria/{entry-counter-plugin-criteria-name} | Delete a Entry Counter Plugin Criteria
[**GetEntryCounterPluginCriteria**](EntryCounterPluginCriteriaAPI.md#GetEntryCounterPluginCriteria) | **Get** /plugin-root/plugins/{plugin-name}/entry-counter-criteria/{entry-counter-plugin-criteria-name} | Returns a single Entry Counter Plugin Criteria
[**ListEntryCounterCriteria**](EntryCounterPluginCriteriaAPI.md#ListEntryCounterCriteria) | **Get** /plugin-root/plugins/{plugin-name}/entry-counter-criteria | Returns a list of all Entry Counter Plugin Criteria objects
[**UpdateEntryCounterPluginCriteria**](EntryCounterPluginCriteriaAPI.md#UpdateEntryCounterPluginCriteria) | **Patch** /plugin-root/plugins/{plugin-name}/entry-counter-criteria/{entry-counter-plugin-criteria-name} | Update an existing Entry Counter Plugin Criteria by name



## AddEntryCounterPluginCriteria

> EntryCounterPluginCriteriaResponse AddEntryCounterPluginCriteria(ctx, pluginName).AddEntryCounterPluginCriteriaRequest(addEntryCounterPluginCriteriaRequest).Execute()

Add a new Entry Counter Plugin Criteria to the config

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
    addEntryCounterPluginCriteriaRequest := *openapiclient.NewAddEntryCounterPluginCriteriaRequest("Filter_example", "CriteriaName_example") // AddEntryCounterPluginCriteriaRequest | Create a new Entry Counter Plugin Criteria in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryCounterPluginCriteriaAPI.AddEntryCounterPluginCriteria(context.Background(), pluginName).AddEntryCounterPluginCriteriaRequest(addEntryCounterPluginCriteriaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryCounterPluginCriteriaAPI.AddEntryCounterPluginCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEntryCounterPluginCriteria`: EntryCounterPluginCriteriaResponse
    fmt.Fprintf(os.Stdout, "Response from `EntryCounterPluginCriteriaAPI.AddEntryCounterPluginCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEntryCounterPluginCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addEntryCounterPluginCriteriaRequest** | [**AddEntryCounterPluginCriteriaRequest**](AddEntryCounterPluginCriteriaRequest.md) | Create a new Entry Counter Plugin Criteria in the config | 

### Return type

[**EntryCounterPluginCriteriaResponse**](EntryCounterPluginCriteriaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntryCounterPluginCriteria

> DeleteEntryCounterPluginCriteria(ctx, entryCounterPluginCriteriaName, pluginName).Execute()

Delete a Entry Counter Plugin Criteria

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
    entryCounterPluginCriteriaName := "entryCounterPluginCriteriaName_example" // string | Name of the Entry Counter Plugin Criteria
    pluginName := "pluginName_example" // string | Name of the Plugin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EntryCounterPluginCriteriaAPI.DeleteEntryCounterPluginCriteria(context.Background(), entryCounterPluginCriteriaName, pluginName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryCounterPluginCriteriaAPI.DeleteEntryCounterPluginCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryCounterPluginCriteriaName** | **string** | Name of the Entry Counter Plugin Criteria | 
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntryCounterPluginCriteriaRequest struct via the builder pattern


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


## GetEntryCounterPluginCriteria

> EntryCounterPluginCriteriaResponse GetEntryCounterPluginCriteria(ctx, entryCounterPluginCriteriaName, pluginName).Execute()

Returns a single Entry Counter Plugin Criteria

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
    entryCounterPluginCriteriaName := "entryCounterPluginCriteriaName_example" // string | Name of the Entry Counter Plugin Criteria
    pluginName := "pluginName_example" // string | Name of the Plugin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryCounterPluginCriteriaAPI.GetEntryCounterPluginCriteria(context.Background(), entryCounterPluginCriteriaName, pluginName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryCounterPluginCriteriaAPI.GetEntryCounterPluginCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntryCounterPluginCriteria`: EntryCounterPluginCriteriaResponse
    fmt.Fprintf(os.Stdout, "Response from `EntryCounterPluginCriteriaAPI.GetEntryCounterPluginCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryCounterPluginCriteriaName** | **string** | Name of the Entry Counter Plugin Criteria | 
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntryCounterPluginCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntryCounterPluginCriteriaResponse**](EntryCounterPluginCriteriaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntryCounterCriteria

> EntryCounterPluginCriteriaListResponse ListEntryCounterCriteria(ctx, pluginName).Filter(filter).Execute()

Returns a list of all Entry Counter Plugin Criteria objects

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
    resp, r, err := apiClient.EntryCounterPluginCriteriaAPI.ListEntryCounterCriteria(context.Background(), pluginName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryCounterPluginCriteriaAPI.ListEntryCounterCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntryCounterCriteria`: EntryCounterPluginCriteriaListResponse
    fmt.Fprintf(os.Stdout, "Response from `EntryCounterPluginCriteriaAPI.ListEntryCounterCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntryCounterCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**EntryCounterPluginCriteriaListResponse**](EntryCounterPluginCriteriaListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntryCounterPluginCriteria

> EntryCounterPluginCriteriaResponse UpdateEntryCounterPluginCriteria(ctx, entryCounterPluginCriteriaName, pluginName).UpdateRequest(updateRequest).Execute()

Update an existing Entry Counter Plugin Criteria by name

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
    entryCounterPluginCriteriaName := "entryCounterPluginCriteriaName_example" // string | Name of the Entry Counter Plugin Criteria
    pluginName := "pluginName_example" // string | Name of the Plugin
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Entry Counter Plugin Criteria

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryCounterPluginCriteriaAPI.UpdateEntryCounterPluginCriteria(context.Background(), entryCounterPluginCriteriaName, pluginName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryCounterPluginCriteriaAPI.UpdateEntryCounterPluginCriteria``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntryCounterPluginCriteria`: EntryCounterPluginCriteriaResponse
    fmt.Fprintf(os.Stdout, "Response from `EntryCounterPluginCriteriaAPI.UpdateEntryCounterPluginCriteria`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryCounterPluginCriteriaName** | **string** | Name of the Entry Counter Plugin Criteria | 
**pluginName** | **string** | Name of the Plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntryCounterPluginCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Entry Counter Plugin Criteria | 

### Return type

[**EntryCounterPluginCriteriaResponse**](EntryCounterPluginCriteriaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

