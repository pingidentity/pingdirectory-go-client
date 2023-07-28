# \AccountStatusNotificationHandlerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccountStatusNotificationHandler**](AccountStatusNotificationHandlerApi.md#AddAccountStatusNotificationHandler) | **Post** /account-status-notification-handlers | Add a new Account Status Notification Handler to the config
[**DeleteAccountStatusNotificationHandler**](AccountStatusNotificationHandlerApi.md#DeleteAccountStatusNotificationHandler) | **Delete** /account-status-notification-handlers/{account-status-notification-handler-name} | Delete a Account Status Notification Handler
[**GetAccountStatusNotificationHandler**](AccountStatusNotificationHandlerApi.md#GetAccountStatusNotificationHandler) | **Get** /account-status-notification-handlers/{account-status-notification-handler-name} | Returns a single Account Status Notification Handler
[**ListAccountStatusNotificationHandlers**](AccountStatusNotificationHandlerApi.md#ListAccountStatusNotificationHandlers) | **Get** /account-status-notification-handlers | Returns a list of all Account Status Notification Handler objects
[**UpdateAccountStatusNotificationHandler**](AccountStatusNotificationHandlerApi.md#UpdateAccountStatusNotificationHandler) | **Patch** /account-status-notification-handlers/{account-status-notification-handler-name} | Update an existing Account Status Notification Handler by name



## AddAccountStatusNotificationHandler

> AddAccountStatusNotificationHandler200Response AddAccountStatusNotificationHandler(ctx).AddAccountStatusNotificationHandlerRequest(addAccountStatusNotificationHandlerRequest).Execute()

Add a new Account Status Notification Handler to the config

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
    addAccountStatusNotificationHandlerRequest := openapiclient.add_account_status_notification_handler_request{AddAdminAlertAccountStatusNotificationHandlerRequest: openapiclient.NewAddAdminAlertAccountStatusNotificationHandlerRequest("HandlerName_example", []openapiclient.EnumadminAlertAccountStatusNotificationHandlerSchemaUrn{openapiclient.Enumadmin-alert-account-status-notification-handlerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:account-status-notification-handler:admin-alert")}, []openapiclient.EnumaccountStatusNotificationHandlerAccountStatusNotificationTypeProp{openapiclient.Enumaccount-status-notification-handler-accountStatusNotificationTypeProp("account-temporarily-locked")}, false)} // AddAccountStatusNotificationHandlerRequest | Create a new Account Status Notification Handler in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountStatusNotificationHandlerApi.AddAccountStatusNotificationHandler(context.Background()).AddAccountStatusNotificationHandlerRequest(addAccountStatusNotificationHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStatusNotificationHandlerApi.AddAccountStatusNotificationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccountStatusNotificationHandler`: AddAccountStatusNotificationHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountStatusNotificationHandlerApi.AddAccountStatusNotificationHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccountStatusNotificationHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addAccountStatusNotificationHandlerRequest** | [**AddAccountStatusNotificationHandlerRequest**](AddAccountStatusNotificationHandlerRequest.md) | Create a new Account Status Notification Handler in the config | 

### Return type

[**AddAccountStatusNotificationHandler200Response**](AddAccountStatusNotificationHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountStatusNotificationHandler

> DeleteAccountStatusNotificationHandler(ctx, accountStatusNotificationHandlerName).Execute()

Delete a Account Status Notification Handler

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
    accountStatusNotificationHandlerName := "accountStatusNotificationHandlerName_example" // string | Name of the Account Status Notification Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountStatusNotificationHandlerApi.DeleteAccountStatusNotificationHandler(context.Background(), accountStatusNotificationHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStatusNotificationHandlerApi.DeleteAccountStatusNotificationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountStatusNotificationHandlerName** | **string** | Name of the Account Status Notification Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountStatusNotificationHandlerRequest struct via the builder pattern


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


## GetAccountStatusNotificationHandler

> AddAccountStatusNotificationHandler200Response GetAccountStatusNotificationHandler(ctx, accountStatusNotificationHandlerName).Execute()

Returns a single Account Status Notification Handler

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
    accountStatusNotificationHandlerName := "accountStatusNotificationHandlerName_example" // string | Name of the Account Status Notification Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountStatusNotificationHandlerApi.GetAccountStatusNotificationHandler(context.Background(), accountStatusNotificationHandlerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStatusNotificationHandlerApi.GetAccountStatusNotificationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountStatusNotificationHandler`: AddAccountStatusNotificationHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountStatusNotificationHandlerApi.GetAccountStatusNotificationHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountStatusNotificationHandlerName** | **string** | Name of the Account Status Notification Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountStatusNotificationHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddAccountStatusNotificationHandler200Response**](AddAccountStatusNotificationHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountStatusNotificationHandlers

> AccountStatusNotificationHandlerListResponse ListAccountStatusNotificationHandlers(ctx).Filter(filter).Execute()

Returns a list of all Account Status Notification Handler objects

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
    resp, r, err := apiClient.AccountStatusNotificationHandlerApi.ListAccountStatusNotificationHandlers(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStatusNotificationHandlerApi.ListAccountStatusNotificationHandlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountStatusNotificationHandlers`: AccountStatusNotificationHandlerListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountStatusNotificationHandlerApi.ListAccountStatusNotificationHandlers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountStatusNotificationHandlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**AccountStatusNotificationHandlerListResponse**](AccountStatusNotificationHandlerListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountStatusNotificationHandler

> AddAccountStatusNotificationHandler200Response UpdateAccountStatusNotificationHandler(ctx, accountStatusNotificationHandlerName).UpdateRequest(updateRequest).Execute()

Update an existing Account Status Notification Handler by name

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
    accountStatusNotificationHandlerName := "accountStatusNotificationHandlerName_example" // string | Name of the Account Status Notification Handler
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Account Status Notification Handler

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountStatusNotificationHandlerApi.UpdateAccountStatusNotificationHandler(context.Background(), accountStatusNotificationHandlerName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountStatusNotificationHandlerApi.UpdateAccountStatusNotificationHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountStatusNotificationHandler`: AddAccountStatusNotificationHandler200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountStatusNotificationHandlerApi.UpdateAccountStatusNotificationHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountStatusNotificationHandlerName** | **string** | Name of the Account Status Notification Handler | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountStatusNotificationHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Account Status Notification Handler | 

### Return type

[**AddAccountStatusNotificationHandler200Response**](AddAccountStatusNotificationHandler200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

