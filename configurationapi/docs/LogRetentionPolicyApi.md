# \LogRetentionPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogRetentionPolicy**](LogRetentionPolicyAPI.md#AddLogRetentionPolicy) | **Post** /log-retention-policies | Add a new Log Retention Policy to the config
[**DeleteLogRetentionPolicy**](LogRetentionPolicyAPI.md#DeleteLogRetentionPolicy) | **Delete** /log-retention-policies/{log-retention-policy-name} | Delete a Log Retention Policy
[**GetLogRetentionPolicy**](LogRetentionPolicyAPI.md#GetLogRetentionPolicy) | **Get** /log-retention-policies/{log-retention-policy-name} | Returns a single Log Retention Policy
[**ListLogRetentionPolicies**](LogRetentionPolicyAPI.md#ListLogRetentionPolicies) | **Get** /log-retention-policies | Returns a list of all Log Retention Policy objects
[**UpdateLogRetentionPolicy**](LogRetentionPolicyAPI.md#UpdateLogRetentionPolicy) | **Patch** /log-retention-policies/{log-retention-policy-name} | Update an existing Log Retention Policy by name



## AddLogRetentionPolicy

> AddLogRetentionPolicy200Response AddLogRetentionPolicy(ctx).AddLogRetentionPolicyRequest(addLogRetentionPolicyRequest).Execute()

Add a new Log Retention Policy to the config

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
    addLogRetentionPolicyRequest := openapiclient.add_log_retention_policy_request{AddFileCountLogRetentionPolicyRequest: openapiclient.NewAddFileCountLogRetentionPolicyRequest([]openapiclient.EnumfileCountLogRetentionPolicySchemaUrn{openapiclient.Enumfile-count-log-retention-policySchemaUrn("urn:pingidentity:schemas:configuration:2.0:log-retention-policy:file-count")}, int64(123), "PolicyName_example")} // AddLogRetentionPolicyRequest | Create a new Log Retention Policy in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogRetentionPolicyAPI.AddLogRetentionPolicy(context.Background()).AddLogRetentionPolicyRequest(addLogRetentionPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogRetentionPolicyAPI.AddLogRetentionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLogRetentionPolicy`: AddLogRetentionPolicy200Response
    fmt.Fprintf(os.Stdout, "Response from `LogRetentionPolicyAPI.AddLogRetentionPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLogRetentionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLogRetentionPolicyRequest** | [**AddLogRetentionPolicyRequest**](AddLogRetentionPolicyRequest.md) | Create a new Log Retention Policy in the config | 

### Return type

[**AddLogRetentionPolicy200Response**](AddLogRetentionPolicy200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogRetentionPolicy

> DeleteLogRetentionPolicy(ctx, logRetentionPolicyName).Execute()

Delete a Log Retention Policy

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
    logRetentionPolicyName := "logRetentionPolicyName_example" // string | Name of the Log Retention Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LogRetentionPolicyAPI.DeleteLogRetentionPolicy(context.Background(), logRetentionPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogRetentionPolicyAPI.DeleteLogRetentionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logRetentionPolicyName** | **string** | Name of the Log Retention Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogRetentionPolicyRequest struct via the builder pattern


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


## GetLogRetentionPolicy

> AddLogRetentionPolicy200Response GetLogRetentionPolicy(ctx, logRetentionPolicyName).Execute()

Returns a single Log Retention Policy

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
    logRetentionPolicyName := "logRetentionPolicyName_example" // string | Name of the Log Retention Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogRetentionPolicyAPI.GetLogRetentionPolicy(context.Background(), logRetentionPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogRetentionPolicyAPI.GetLogRetentionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogRetentionPolicy`: AddLogRetentionPolicy200Response
    fmt.Fprintf(os.Stdout, "Response from `LogRetentionPolicyAPI.GetLogRetentionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logRetentionPolicyName** | **string** | Name of the Log Retention Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogRetentionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddLogRetentionPolicy200Response**](AddLogRetentionPolicy200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogRetentionPolicies

> LogRetentionPolicyListResponse ListLogRetentionPolicies(ctx).Filter(filter).Execute()

Returns a list of all Log Retention Policy objects

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
    resp, r, err := apiClient.LogRetentionPolicyAPI.ListLogRetentionPolicies(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogRetentionPolicyAPI.ListLogRetentionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogRetentionPolicies`: LogRetentionPolicyListResponse
    fmt.Fprintf(os.Stdout, "Response from `LogRetentionPolicyAPI.ListLogRetentionPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogRetentionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**LogRetentionPolicyListResponse**](LogRetentionPolicyListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogRetentionPolicy

> AddLogRetentionPolicy200Response UpdateLogRetentionPolicy(ctx, logRetentionPolicyName).UpdateRequest(updateRequest).Execute()

Update an existing Log Retention Policy by name

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
    logRetentionPolicyName := "logRetentionPolicyName_example" // string | Name of the Log Retention Policy
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Log Retention Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogRetentionPolicyAPI.UpdateLogRetentionPolicy(context.Background(), logRetentionPolicyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogRetentionPolicyAPI.UpdateLogRetentionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogRetentionPolicy`: AddLogRetentionPolicy200Response
    fmt.Fprintf(os.Stdout, "Response from `LogRetentionPolicyAPI.UpdateLogRetentionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logRetentionPolicyName** | **string** | Name of the Log Retention Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogRetentionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Log Retention Policy | 

### Return type

[**AddLogRetentionPolicy200Response**](AddLogRetentionPolicy200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

