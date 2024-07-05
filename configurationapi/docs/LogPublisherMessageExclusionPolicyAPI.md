# \LogPublisherMessageExclusionPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLogPublisherMessageExclusionPolicy**](LogPublisherMessageExclusionPolicyAPI.md#AddLogPublisherMessageExclusionPolicy) | **Post** /log-publisher-message-exclusion-policies | Add a new Log Publisher Message Exclusion Policy to the config
[**DeleteLogPublisherMessageExclusionPolicy**](LogPublisherMessageExclusionPolicyAPI.md#DeleteLogPublisherMessageExclusionPolicy) | **Delete** /log-publisher-message-exclusion-policies/{log-publisher-message-exclusion-policy-name} | Delete a Log Publisher Message Exclusion Policy
[**GetLogPublisherMessageExclusionPolicy**](LogPublisherMessageExclusionPolicyAPI.md#GetLogPublisherMessageExclusionPolicy) | **Get** /log-publisher-message-exclusion-policies/{log-publisher-message-exclusion-policy-name} | Returns a single Log Publisher Message Exclusion Policy
[**ListLogPublisherMessageExclusionPolicies**](LogPublisherMessageExclusionPolicyAPI.md#ListLogPublisherMessageExclusionPolicies) | **Get** /log-publisher-message-exclusion-policies | Returns a list of all Log Publisher Message Exclusion Policy objects
[**UpdateLogPublisherMessageExclusionPolicy**](LogPublisherMessageExclusionPolicyAPI.md#UpdateLogPublisherMessageExclusionPolicy) | **Patch** /log-publisher-message-exclusion-policies/{log-publisher-message-exclusion-policy-name} | Update an existing Log Publisher Message Exclusion Policy by name



## AddLogPublisherMessageExclusionPolicy

> ErrorLogPublisherMessageExclusionPolicyResponse AddLogPublisherMessageExclusionPolicy(ctx).AddErrorLogPublisherMessageExclusionPolicyRequest(addErrorLogPublisherMessageExclusionPolicyRequest).Execute()

Add a new Log Publisher Message Exclusion Policy to the config

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
    addErrorLogPublisherMessageExclusionPolicyRequest := *openapiclient.NewAddErrorLogPublisherMessageExclusionPolicyRequest([]openapiclient.EnumerrorLogPublisherMessageExclusionPolicySchemaUrn{openapiclient.Enumerror-log-publisher-message-exclusion-policySchemaUrn("urn:pingidentity:schemas:configuration:2.0:log-publisher-message-exclusion-policy:error")}, openapiclient.Enumlog-publisher-message-exclusion-policy-logMessageCategoryProp("CORE"), openapiclient.Enumlog-publisher-message-exclusion-policy-logMessageSeverityProp("FATAL_ERROR"), "LogMessageRegex_example", false, "PolicyName_example") // AddErrorLogPublisherMessageExclusionPolicyRequest | Create a new Log Publisher Message Exclusion Policy in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogPublisherMessageExclusionPolicyAPI.AddLogPublisherMessageExclusionPolicy(context.Background()).AddErrorLogPublisherMessageExclusionPolicyRequest(addErrorLogPublisherMessageExclusionPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogPublisherMessageExclusionPolicyAPI.AddLogPublisherMessageExclusionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLogPublisherMessageExclusionPolicy`: ErrorLogPublisherMessageExclusionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `LogPublisherMessageExclusionPolicyAPI.AddLogPublisherMessageExclusionPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLogPublisherMessageExclusionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addErrorLogPublisherMessageExclusionPolicyRequest** | [**AddErrorLogPublisherMessageExclusionPolicyRequest**](AddErrorLogPublisherMessageExclusionPolicyRequest.md) | Create a new Log Publisher Message Exclusion Policy in the config | 

### Return type

[**ErrorLogPublisherMessageExclusionPolicyResponse**](ErrorLogPublisherMessageExclusionPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogPublisherMessageExclusionPolicy

> DeleteLogPublisherMessageExclusionPolicy(ctx, logPublisherMessageExclusionPolicyName).Execute()

Delete a Log Publisher Message Exclusion Policy

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
    logPublisherMessageExclusionPolicyName := "logPublisherMessageExclusionPolicyName_example" // string | Name of the Log Publisher Message Exclusion Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LogPublisherMessageExclusionPolicyAPI.DeleteLogPublisherMessageExclusionPolicy(context.Background(), logPublisherMessageExclusionPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogPublisherMessageExclusionPolicyAPI.DeleteLogPublisherMessageExclusionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logPublisherMessageExclusionPolicyName** | **string** | Name of the Log Publisher Message Exclusion Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogPublisherMessageExclusionPolicyRequest struct via the builder pattern


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


## GetLogPublisherMessageExclusionPolicy

> ErrorLogPublisherMessageExclusionPolicyResponse GetLogPublisherMessageExclusionPolicy(ctx, logPublisherMessageExclusionPolicyName).Execute()

Returns a single Log Publisher Message Exclusion Policy

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
    logPublisherMessageExclusionPolicyName := "logPublisherMessageExclusionPolicyName_example" // string | Name of the Log Publisher Message Exclusion Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogPublisherMessageExclusionPolicyAPI.GetLogPublisherMessageExclusionPolicy(context.Background(), logPublisherMessageExclusionPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogPublisherMessageExclusionPolicyAPI.GetLogPublisherMessageExclusionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogPublisherMessageExclusionPolicy`: ErrorLogPublisherMessageExclusionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `LogPublisherMessageExclusionPolicyAPI.GetLogPublisherMessageExclusionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logPublisherMessageExclusionPolicyName** | **string** | Name of the Log Publisher Message Exclusion Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogPublisherMessageExclusionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ErrorLogPublisherMessageExclusionPolicyResponse**](ErrorLogPublisherMessageExclusionPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogPublisherMessageExclusionPolicies

> LogPublisherMessageExclusionPolicyListResponse ListLogPublisherMessageExclusionPolicies(ctx).Filter(filter).Execute()

Returns a list of all Log Publisher Message Exclusion Policy objects

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
    resp, r, err := apiClient.LogPublisherMessageExclusionPolicyAPI.ListLogPublisherMessageExclusionPolicies(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogPublisherMessageExclusionPolicyAPI.ListLogPublisherMessageExclusionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogPublisherMessageExclusionPolicies`: LogPublisherMessageExclusionPolicyListResponse
    fmt.Fprintf(os.Stdout, "Response from `LogPublisherMessageExclusionPolicyAPI.ListLogPublisherMessageExclusionPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogPublisherMessageExclusionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**LogPublisherMessageExclusionPolicyListResponse**](LogPublisherMessageExclusionPolicyListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogPublisherMessageExclusionPolicy

> ErrorLogPublisherMessageExclusionPolicyResponse UpdateLogPublisherMessageExclusionPolicy(ctx, logPublisherMessageExclusionPolicyName).UpdateRequest(updateRequest).Execute()

Update an existing Log Publisher Message Exclusion Policy by name

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
    logPublisherMessageExclusionPolicyName := "logPublisherMessageExclusionPolicyName_example" // string | Name of the Log Publisher Message Exclusion Policy
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Log Publisher Message Exclusion Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogPublisherMessageExclusionPolicyAPI.UpdateLogPublisherMessageExclusionPolicy(context.Background(), logPublisherMessageExclusionPolicyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogPublisherMessageExclusionPolicyAPI.UpdateLogPublisherMessageExclusionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogPublisherMessageExclusionPolicy`: ErrorLogPublisherMessageExclusionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `LogPublisherMessageExclusionPolicyAPI.UpdateLogPublisherMessageExclusionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logPublisherMessageExclusionPolicyName** | **string** | Name of the Log Publisher Message Exclusion Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogPublisherMessageExclusionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Log Publisher Message Exclusion Policy | 

### Return type

[**ErrorLogPublisherMessageExclusionPolicyResponse**](ErrorLogPublisherMessageExclusionPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

