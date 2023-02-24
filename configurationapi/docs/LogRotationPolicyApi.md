# \LogRotationPolicyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogRotationPolicy**](LogRotationPolicyApi.md#AddLogRotationPolicy) | **Post** /log-rotation-policies | Add a new Log Rotation Policy to the config
[**DeleteLogRotationPolicy**](LogRotationPolicyApi.md#DeleteLogRotationPolicy) | **Delete** /log-rotation-policies/{log-rotation-policy-name} | Delete a Log Rotation Policy
[**GetLogRotationPolicy**](LogRotationPolicyApi.md#GetLogRotationPolicy) | **Get** /log-rotation-policies/{log-rotation-policy-name} | Returns a single Log Rotation Policy
[**UpdateLogRotationPolicy**](LogRotationPolicyApi.md#UpdateLogRotationPolicy) | **Patch** /log-rotation-policies/{log-rotation-policy-name} | Update an existing Log Rotation Policy by name



## AddLogRotationPolicy

> AddLogRotationPolicy200Response AddLogRotationPolicy(ctx).AddLogRotationPolicyRequest(addLogRotationPolicyRequest).Execute()

Add a new Log Rotation Policy to the config

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
    addLogRotationPolicyRequest := openapiclient.add_log_rotation_policy_request{AddFixedTimeLogRotationPolicyRequest: openapiclient.NewAddFixedTimeLogRotationPolicyRequest("PolicyName_example", []openapiclient.EnumfixedTimeLogRotationPolicySchemaUrn{openapiclient.Enumfixed-time-log-rotation-policySchemaUrn("urn:pingidentity:schemas:configuration:2.0:log-rotation-policy:fixed-time")}, []string{"TimeOfDay_example"})} // AddLogRotationPolicyRequest | Create a new Log Rotation Policy in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogRotationPolicyApi.AddLogRotationPolicy(context.Background()).AddLogRotationPolicyRequest(addLogRotationPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogRotationPolicyApi.AddLogRotationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLogRotationPolicy`: AddLogRotationPolicy200Response
    fmt.Fprintf(os.Stdout, "Response from `LogRotationPolicyApi.AddLogRotationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLogRotationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLogRotationPolicyRequest** | [**AddLogRotationPolicyRequest**](AddLogRotationPolicyRequest.md) | Create a new Log Rotation Policy in the config | 

### Return type

[**AddLogRotationPolicy200Response**](AddLogRotationPolicy200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogRotationPolicy

> DeleteLogRotationPolicy(ctx, logRotationPolicyName).Execute()

Delete a Log Rotation Policy

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
    logRotationPolicyName := "logRotationPolicyName_example" // string | Name of the Log Rotation Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LogRotationPolicyApi.DeleteLogRotationPolicy(context.Background(), logRotationPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogRotationPolicyApi.DeleteLogRotationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logRotationPolicyName** | **string** | Name of the Log Rotation Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogRotationPolicyRequest struct via the builder pattern


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


## GetLogRotationPolicy

> AddLogRotationPolicy200Response GetLogRotationPolicy(ctx, logRotationPolicyName).Execute()

Returns a single Log Rotation Policy

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
    logRotationPolicyName := "logRotationPolicyName_example" // string | Name of the Log Rotation Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogRotationPolicyApi.GetLogRotationPolicy(context.Background(), logRotationPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogRotationPolicyApi.GetLogRotationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogRotationPolicy`: AddLogRotationPolicy200Response
    fmt.Fprintf(os.Stdout, "Response from `LogRotationPolicyApi.GetLogRotationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logRotationPolicyName** | **string** | Name of the Log Rotation Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogRotationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddLogRotationPolicy200Response**](AddLogRotationPolicy200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogRotationPolicy

> AddLogRotationPolicy200Response UpdateLogRotationPolicy(ctx, logRotationPolicyName).UpdateRequest(updateRequest).Execute()

Update an existing Log Rotation Policy by name

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
    logRotationPolicyName := "logRotationPolicyName_example" // string | Name of the Log Rotation Policy
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Log Rotation Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogRotationPolicyApi.UpdateLogRotationPolicy(context.Background(), logRotationPolicyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogRotationPolicyApi.UpdateLogRotationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogRotationPolicy`: AddLogRotationPolicy200Response
    fmt.Fprintf(os.Stdout, "Response from `LogRotationPolicyApi.UpdateLogRotationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logRotationPolicyName** | **string** | Name of the Log Rotation Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogRotationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Log Rotation Policy | 

### Return type

[**AddLogRotationPolicy200Response**](AddLogRotationPolicy200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

