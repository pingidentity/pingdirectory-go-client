# \ChangeSubscriptionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddChangeSubscription**](ChangeSubscriptionAPI.md#AddChangeSubscription) | **Post** /change-subscriptions | Add a new Change Subscription to the config
[**DeleteChangeSubscription**](ChangeSubscriptionAPI.md#DeleteChangeSubscription) | **Delete** /change-subscriptions/{change-subscription-name} | Delete a Change Subscription
[**GetChangeSubscription**](ChangeSubscriptionAPI.md#GetChangeSubscription) | **Get** /change-subscriptions/{change-subscription-name} | Returns a single Change Subscription
[**ListChangeSubscriptions**](ChangeSubscriptionAPI.md#ListChangeSubscriptions) | **Get** /change-subscriptions | Returns a list of all Change Subscription objects
[**UpdateChangeSubscription**](ChangeSubscriptionAPI.md#UpdateChangeSubscription) | **Patch** /change-subscriptions/{change-subscription-name} | Update an existing Change Subscription by name



## AddChangeSubscription

> ChangeSubscriptionResponse AddChangeSubscription(ctx).AddChangeSubscriptionRequest(addChangeSubscriptionRequest).Execute()

Add a new Change Subscription to the config

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
    addChangeSubscriptionRequest := *openapiclient.NewAddChangeSubscriptionRequest("SubscriptionName_example") // AddChangeSubscriptionRequest | Create a new Change Subscription in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChangeSubscriptionAPI.AddChangeSubscription(context.Background()).AddChangeSubscriptionRequest(addChangeSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChangeSubscriptionAPI.AddChangeSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddChangeSubscription`: ChangeSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `ChangeSubscriptionAPI.AddChangeSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddChangeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addChangeSubscriptionRequest** | [**AddChangeSubscriptionRequest**](AddChangeSubscriptionRequest.md) | Create a new Change Subscription in the config | 

### Return type

[**ChangeSubscriptionResponse**](ChangeSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChangeSubscription

> DeleteChangeSubscription(ctx, changeSubscriptionName).Execute()

Delete a Change Subscription

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
    changeSubscriptionName := "changeSubscriptionName_example" // string | Name of the Change Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ChangeSubscriptionAPI.DeleteChangeSubscription(context.Background(), changeSubscriptionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChangeSubscriptionAPI.DeleteChangeSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeSubscriptionName** | **string** | Name of the Change Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChangeSubscriptionRequest struct via the builder pattern


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


## GetChangeSubscription

> ChangeSubscriptionResponse GetChangeSubscription(ctx, changeSubscriptionName).Execute()

Returns a single Change Subscription

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
    changeSubscriptionName := "changeSubscriptionName_example" // string | Name of the Change Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChangeSubscriptionAPI.GetChangeSubscription(context.Background(), changeSubscriptionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChangeSubscriptionAPI.GetChangeSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChangeSubscription`: ChangeSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `ChangeSubscriptionAPI.GetChangeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeSubscriptionName** | **string** | Name of the Change Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChangeSubscriptionResponse**](ChangeSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChangeSubscriptions

> ChangeSubscriptionListResponse ListChangeSubscriptions(ctx).Filter(filter).Execute()

Returns a list of all Change Subscription objects

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
    resp, r, err := apiClient.ChangeSubscriptionAPI.ListChangeSubscriptions(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChangeSubscriptionAPI.ListChangeSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChangeSubscriptions`: ChangeSubscriptionListResponse
    fmt.Fprintf(os.Stdout, "Response from `ChangeSubscriptionAPI.ListChangeSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListChangeSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ChangeSubscriptionListResponse**](ChangeSubscriptionListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChangeSubscription

> ChangeSubscriptionResponse UpdateChangeSubscription(ctx, changeSubscriptionName).UpdateRequest(updateRequest).Execute()

Update an existing Change Subscription by name

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
    changeSubscriptionName := "changeSubscriptionName_example" // string | Name of the Change Subscription
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Change Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChangeSubscriptionAPI.UpdateChangeSubscription(context.Background(), changeSubscriptionName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChangeSubscriptionAPI.UpdateChangeSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChangeSubscription`: ChangeSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `ChangeSubscriptionAPI.UpdateChangeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changeSubscriptionName** | **string** | Name of the Change Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChangeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Change Subscription | 

### Return type

[**ChangeSubscriptionResponse**](ChangeSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

