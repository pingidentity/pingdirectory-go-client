# \NotificationManagerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNotificationManager**](NotificationManagerAPI.md#AddNotificationManager) | **Post** /notification-managers | Add a new Notification Manager to the config
[**DeleteNotificationManager**](NotificationManagerAPI.md#DeleteNotificationManager) | **Delete** /notification-managers/{notification-manager-name} | Delete a Notification Manager
[**GetNotificationManager**](NotificationManagerAPI.md#GetNotificationManager) | **Get** /notification-managers/{notification-manager-name} | Returns a single Notification Manager
[**ListNotificationManagers**](NotificationManagerAPI.md#ListNotificationManagers) | **Get** /notification-managers | Returns a list of all Notification Manager objects
[**UpdateNotificationManager**](NotificationManagerAPI.md#UpdateNotificationManager) | **Patch** /notification-managers/{notification-manager-name} | Update an existing Notification Manager by name



## AddNotificationManager

> ThirdPartyNotificationManagerResponse AddNotificationManager(ctx).AddThirdPartyNotificationManagerRequest(addThirdPartyNotificationManagerRequest).Execute()

Add a new Notification Manager to the config

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
    addThirdPartyNotificationManagerRequest := *openapiclient.NewAddThirdPartyNotificationManagerRequest([]openapiclient.EnumthirdPartyNotificationManagerSchemaUrn{openapiclient.Enumthird-party-notification-managerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:notification-manager:third-party")}, "ExtensionClass_example", false, "SubscriptionBaseDN_example", "ManagerName_example") // AddThirdPartyNotificationManagerRequest | Create a new Notification Manager in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationManagerAPI.AddNotificationManager(context.Background()).AddThirdPartyNotificationManagerRequest(addThirdPartyNotificationManagerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationManagerAPI.AddNotificationManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddNotificationManager`: ThirdPartyNotificationManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationManagerAPI.AddNotificationManager`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddNotificationManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addThirdPartyNotificationManagerRequest** | [**AddThirdPartyNotificationManagerRequest**](AddThirdPartyNotificationManagerRequest.md) | Create a new Notification Manager in the config | 

### Return type

[**ThirdPartyNotificationManagerResponse**](ThirdPartyNotificationManagerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationManager

> DeleteNotificationManager(ctx, notificationManagerName).Execute()

Delete a Notification Manager

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
    notificationManagerName := "notificationManagerName_example" // string | Name of the Notification Manager

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationManagerAPI.DeleteNotificationManager(context.Background(), notificationManagerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationManagerAPI.DeleteNotificationManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationManagerName** | **string** | Name of the Notification Manager | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationManagerRequest struct via the builder pattern


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


## GetNotificationManager

> ThirdPartyNotificationManagerResponse GetNotificationManager(ctx, notificationManagerName).Execute()

Returns a single Notification Manager

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
    notificationManagerName := "notificationManagerName_example" // string | Name of the Notification Manager

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationManagerAPI.GetNotificationManager(context.Background(), notificationManagerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationManagerAPI.GetNotificationManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationManager`: ThirdPartyNotificationManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationManagerAPI.GetNotificationManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationManagerName** | **string** | Name of the Notification Manager | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThirdPartyNotificationManagerResponse**](ThirdPartyNotificationManagerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationManagers

> NotificationManagerListResponse ListNotificationManagers(ctx).Filter(filter).Execute()

Returns a list of all Notification Manager objects

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
    resp, r, err := apiClient.NotificationManagerAPI.ListNotificationManagers(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationManagerAPI.ListNotificationManagers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationManagers`: NotificationManagerListResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationManagerAPI.ListNotificationManagers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**NotificationManagerListResponse**](NotificationManagerListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationManager

> ThirdPartyNotificationManagerResponse UpdateNotificationManager(ctx, notificationManagerName).UpdateRequest(updateRequest).Execute()

Update an existing Notification Manager by name

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
    notificationManagerName := "notificationManagerName_example" // string | Name of the Notification Manager
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Notification Manager

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationManagerAPI.UpdateNotificationManager(context.Background(), notificationManagerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationManagerAPI.UpdateNotificationManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationManager`: ThirdPartyNotificationManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `NotificationManagerAPI.UpdateNotificationManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationManagerName** | **string** | Name of the Notification Manager | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Notification Manager | 

### Return type

[**ThirdPartyNotificationManagerResponse**](ThirdPartyNotificationManagerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

