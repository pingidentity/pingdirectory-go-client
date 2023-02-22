# \DataSecurityAuditorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDataSecurityAuditor**](DataSecurityAuditorApi.md#AddDataSecurityAuditor) | **Post** /data-security-auditors | Add a new Data Security Auditor to the config
[**DeleteDataSecurityAuditor**](DataSecurityAuditorApi.md#DeleteDataSecurityAuditor) | **Delete** /data-security-auditors/{data-security-auditor-name} | Delete a Data Security Auditor
[**GetDataSecurityAuditor**](DataSecurityAuditorApi.md#GetDataSecurityAuditor) | **Get** /data-security-auditors/{data-security-auditor-name} | Returns a single Data Security Auditor
[**UpdateDataSecurityAuditor**](DataSecurityAuditorApi.md#UpdateDataSecurityAuditor) | **Patch** /data-security-auditors/{data-security-auditor-name} | Update an existing Data Security Auditor by name



## AddDataSecurityAuditor

> AddDataSecurityAuditor200Response AddDataSecurityAuditor(ctx).AddDataSecurityAuditorRequest(addDataSecurityAuditorRequest).Execute()

Add a new Data Security Auditor to the config

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
    addDataSecurityAuditorRequest := openapiclient.add_data_security_auditor_request{AddAccessControlDataSecurityAuditorRequest: openapiclient.NewAddAccessControlDataSecurityAuditorRequest("AuditorName_example", []openapiclient.EnumaccessControlDataSecurityAuditorSchemaUrn{openapiclient.Enumaccess-control-data-security-auditorSchemaUrn("urn:pingidentity:schemas:configuration:2.0:data-security-auditor:access-control")}, "ReportFile_example", false)} // AddDataSecurityAuditorRequest | Create a new Data Security Auditor in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataSecurityAuditorApi.AddDataSecurityAuditor(context.Background()).AddDataSecurityAuditorRequest(addDataSecurityAuditorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSecurityAuditorApi.AddDataSecurityAuditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDataSecurityAuditor`: AddDataSecurityAuditor200Response
    fmt.Fprintf(os.Stdout, "Response from `DataSecurityAuditorApi.AddDataSecurityAuditor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDataSecurityAuditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addDataSecurityAuditorRequest** | [**AddDataSecurityAuditorRequest**](AddDataSecurityAuditorRequest.md) | Create a new Data Security Auditor in the config | 

### Return type

[**AddDataSecurityAuditor200Response**](AddDataSecurityAuditor200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataSecurityAuditor

> DeleteDataSecurityAuditor(ctx, dataSecurityAuditorName).Execute()

Delete a Data Security Auditor

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
    dataSecurityAuditorName := "dataSecurityAuditorName_example" // string | Name of the Data Security Auditor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataSecurityAuditorApi.DeleteDataSecurityAuditor(context.Background(), dataSecurityAuditorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSecurityAuditorApi.DeleteDataSecurityAuditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSecurityAuditorName** | **string** | Name of the Data Security Auditor | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataSecurityAuditorRequest struct via the builder pattern


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


## GetDataSecurityAuditor

> AddDataSecurityAuditor200Response GetDataSecurityAuditor(ctx, dataSecurityAuditorName).Execute()

Returns a single Data Security Auditor

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
    dataSecurityAuditorName := "dataSecurityAuditorName_example" // string | Name of the Data Security Auditor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataSecurityAuditorApi.GetDataSecurityAuditor(context.Background(), dataSecurityAuditorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSecurityAuditorApi.GetDataSecurityAuditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataSecurityAuditor`: AddDataSecurityAuditor200Response
    fmt.Fprintf(os.Stdout, "Response from `DataSecurityAuditorApi.GetDataSecurityAuditor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSecurityAuditorName** | **string** | Name of the Data Security Auditor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSecurityAuditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddDataSecurityAuditor200Response**](AddDataSecurityAuditor200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataSecurityAuditor

> AddDataSecurityAuditor200Response UpdateDataSecurityAuditor(ctx, dataSecurityAuditorName).UpdateRequest(updateRequest).Execute()

Update an existing Data Security Auditor by name

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
    dataSecurityAuditorName := "dataSecurityAuditorName_example" // string | Name of the Data Security Auditor
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Data Security Auditor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataSecurityAuditorApi.UpdateDataSecurityAuditor(context.Background(), dataSecurityAuditorName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSecurityAuditorApi.UpdateDataSecurityAuditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDataSecurityAuditor`: AddDataSecurityAuditor200Response
    fmt.Fprintf(os.Stdout, "Response from `DataSecurityAuditorApi.UpdateDataSecurityAuditor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSecurityAuditorName** | **string** | Name of the Data Security Auditor | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataSecurityAuditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Data Security Auditor | 

### Return type

[**AddDataSecurityAuditor200Response**](AddDataSecurityAuditor200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

