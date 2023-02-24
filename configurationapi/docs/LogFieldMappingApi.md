# \LogFieldMappingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogFieldMapping**](LogFieldMappingApi.md#AddLogFieldMapping) | **Post** /log-field-mappings | Add a new Log Field Mapping to the config
[**DeleteLogFieldMapping**](LogFieldMappingApi.md#DeleteLogFieldMapping) | **Delete** /log-field-mappings/{log-field-mapping-name} | Delete a Log Field Mapping
[**GetLogFieldMapping**](LogFieldMappingApi.md#GetLogFieldMapping) | **Get** /log-field-mappings/{log-field-mapping-name} | Returns a single Log Field Mapping
[**UpdateLogFieldMapping**](LogFieldMappingApi.md#UpdateLogFieldMapping) | **Patch** /log-field-mappings/{log-field-mapping-name} | Update an existing Log Field Mapping by name



## AddLogFieldMapping

> AddLogFieldMapping200Response AddLogFieldMapping(ctx).AddLogFieldMappingRequest(addLogFieldMappingRequest).Execute()

Add a new Log Field Mapping to the config

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
    addLogFieldMappingRequest := openapiclient.add_log_field_mapping_request{AddAccessLogFieldMappingRequest: openapiclient.NewAddAccessLogFieldMappingRequest("MappingName_example", []openapiclient.EnumaccessLogFieldMappingSchemaUrn{openapiclient.Enumaccess-log-field-mappingSchemaUrn("urn:pingidentity:schemas:configuration:2.0:log-field-mapping:access")})} // AddLogFieldMappingRequest | Create a new Log Field Mapping in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldMappingApi.AddLogFieldMapping(context.Background()).AddLogFieldMappingRequest(addLogFieldMappingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldMappingApi.AddLogFieldMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLogFieldMapping`: AddLogFieldMapping200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFieldMappingApi.AddLogFieldMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLogFieldMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLogFieldMappingRequest** | [**AddLogFieldMappingRequest**](AddLogFieldMappingRequest.md) | Create a new Log Field Mapping in the config | 

### Return type

[**AddLogFieldMapping200Response**](AddLogFieldMapping200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogFieldMapping

> DeleteLogFieldMapping(ctx, logFieldMappingName).Execute()

Delete a Log Field Mapping

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
    logFieldMappingName := "logFieldMappingName_example" // string | Name of the Log Field Mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LogFieldMappingApi.DeleteLogFieldMapping(context.Background(), logFieldMappingName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldMappingApi.DeleteLogFieldMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logFieldMappingName** | **string** | Name of the Log Field Mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogFieldMappingRequest struct via the builder pattern


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


## GetLogFieldMapping

> AddLogFieldMapping200Response GetLogFieldMapping(ctx, logFieldMappingName).Execute()

Returns a single Log Field Mapping

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
    logFieldMappingName := "logFieldMappingName_example" // string | Name of the Log Field Mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldMappingApi.GetLogFieldMapping(context.Background(), logFieldMappingName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldMappingApi.GetLogFieldMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogFieldMapping`: AddLogFieldMapping200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFieldMappingApi.GetLogFieldMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logFieldMappingName** | **string** | Name of the Log Field Mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogFieldMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddLogFieldMapping200Response**](AddLogFieldMapping200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogFieldMapping

> AddLogFieldMapping200Response UpdateLogFieldMapping(ctx, logFieldMappingName).UpdateRequest(updateRequest).Execute()

Update an existing Log Field Mapping by name

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
    logFieldMappingName := "logFieldMappingName_example" // string | Name of the Log Field Mapping
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Log Field Mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldMappingApi.UpdateLogFieldMapping(context.Background(), logFieldMappingName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldMappingApi.UpdateLogFieldMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogFieldMapping`: AddLogFieldMapping200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFieldMappingApi.UpdateLogFieldMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logFieldMappingName** | **string** | Name of the Log Field Mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogFieldMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Log Field Mapping | 

### Return type

[**AddLogFieldMapping200Response**](AddLogFieldMapping200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

