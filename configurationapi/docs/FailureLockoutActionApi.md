# \FailureLockoutActionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFailureLockoutAction**](FailureLockoutActionAPI.md#AddFailureLockoutAction) | **Post** /failure-lockout-actions | Add a new Failure Lockout Action to the config
[**DeleteFailureLockoutAction**](FailureLockoutActionAPI.md#DeleteFailureLockoutAction) | **Delete** /failure-lockout-actions/{failure-lockout-action-name} | Delete a Failure Lockout Action
[**GetFailureLockoutAction**](FailureLockoutActionAPI.md#GetFailureLockoutAction) | **Get** /failure-lockout-actions/{failure-lockout-action-name} | Returns a single Failure Lockout Action
[**ListFailureLockoutActions**](FailureLockoutActionAPI.md#ListFailureLockoutActions) | **Get** /failure-lockout-actions | Returns a list of all Failure Lockout Action objects
[**UpdateFailureLockoutAction**](FailureLockoutActionAPI.md#UpdateFailureLockoutAction) | **Patch** /failure-lockout-actions/{failure-lockout-action-name} | Update an existing Failure Lockout Action by name



## AddFailureLockoutAction

> AddFailureLockoutAction200Response AddFailureLockoutAction(ctx).AddFailureLockoutActionRequest(addFailureLockoutActionRequest).Execute()

Add a new Failure Lockout Action to the config

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
    addFailureLockoutActionRequest := openapiclient.add_failure_lockout_action_request{AddDelayBindResponseFailureLockoutActionRequest: openapiclient.NewAddDelayBindResponseFailureLockoutActionRequest([]openapiclient.EnumdelayBindResponseFailureLockoutActionSchemaUrn{openapiclient.Enumdelay-bind-response-failure-lockout-actionSchemaUrn("urn:pingidentity:schemas:configuration:2.0:failure-lockout-action:delay-bind-response")}, "Delay_example", "ActionName_example")} // AddFailureLockoutActionRequest | Create a new Failure Lockout Action in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FailureLockoutActionAPI.AddFailureLockoutAction(context.Background()).AddFailureLockoutActionRequest(addFailureLockoutActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FailureLockoutActionAPI.AddFailureLockoutAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFailureLockoutAction`: AddFailureLockoutAction200Response
    fmt.Fprintf(os.Stdout, "Response from `FailureLockoutActionAPI.AddFailureLockoutAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddFailureLockoutActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addFailureLockoutActionRequest** | [**AddFailureLockoutActionRequest**](AddFailureLockoutActionRequest.md) | Create a new Failure Lockout Action in the config | 

### Return type

[**AddFailureLockoutAction200Response**](AddFailureLockoutAction200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFailureLockoutAction

> DeleteFailureLockoutAction(ctx, failureLockoutActionName).Execute()

Delete a Failure Lockout Action

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
    failureLockoutActionName := "failureLockoutActionName_example" // string | Name of the Failure Lockout Action

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FailureLockoutActionAPI.DeleteFailureLockoutAction(context.Background(), failureLockoutActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FailureLockoutActionAPI.DeleteFailureLockoutAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**failureLockoutActionName** | **string** | Name of the Failure Lockout Action | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFailureLockoutActionRequest struct via the builder pattern


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


## GetFailureLockoutAction

> AddFailureLockoutAction200Response GetFailureLockoutAction(ctx, failureLockoutActionName).Execute()

Returns a single Failure Lockout Action

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
    failureLockoutActionName := "failureLockoutActionName_example" // string | Name of the Failure Lockout Action

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FailureLockoutActionAPI.GetFailureLockoutAction(context.Background(), failureLockoutActionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FailureLockoutActionAPI.GetFailureLockoutAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFailureLockoutAction`: AddFailureLockoutAction200Response
    fmt.Fprintf(os.Stdout, "Response from `FailureLockoutActionAPI.GetFailureLockoutAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**failureLockoutActionName** | **string** | Name of the Failure Lockout Action | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFailureLockoutActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddFailureLockoutAction200Response**](AddFailureLockoutAction200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFailureLockoutActions

> FailureLockoutActionListResponse ListFailureLockoutActions(ctx).Filter(filter).Execute()

Returns a list of all Failure Lockout Action objects

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
    resp, r, err := apiClient.FailureLockoutActionAPI.ListFailureLockoutActions(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FailureLockoutActionAPI.ListFailureLockoutActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFailureLockoutActions`: FailureLockoutActionListResponse
    fmt.Fprintf(os.Stdout, "Response from `FailureLockoutActionAPI.ListFailureLockoutActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFailureLockoutActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**FailureLockoutActionListResponse**](FailureLockoutActionListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFailureLockoutAction

> AddFailureLockoutAction200Response UpdateFailureLockoutAction(ctx, failureLockoutActionName).UpdateRequest(updateRequest).Execute()

Update an existing Failure Lockout Action by name

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
    failureLockoutActionName := "failureLockoutActionName_example" // string | Name of the Failure Lockout Action
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Failure Lockout Action

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FailureLockoutActionAPI.UpdateFailureLockoutAction(context.Background(), failureLockoutActionName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FailureLockoutActionAPI.UpdateFailureLockoutAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFailureLockoutAction`: AddFailureLockoutAction200Response
    fmt.Fprintf(os.Stdout, "Response from `FailureLockoutActionAPI.UpdateFailureLockoutAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**failureLockoutActionName** | **string** | Name of the Failure Lockout Action | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFailureLockoutActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Failure Lockout Action | 

### Return type

[**AddFailureLockoutAction200Response**](AddFailureLockoutAction200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

