# \RootDnUserApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRootDnUser**](RootDnUserApi.md#AddRootDnUser) | **Post** /root-dn/root-dn-users | Add a new Root DN User to the config
[**DeleteRootDnUser**](RootDnUserApi.md#DeleteRootDnUser) | **Delete** /root-dn/root-dn-users/{root-dn-user-name} | Delete a Root DN User
[**GetRootDnUser**](RootDnUserApi.md#GetRootDnUser) | **Get** /root-dn/root-dn-users/{root-dn-user-name} | Returns a single Root DN User
[**ListRootDnUsers**](RootDnUserApi.md#ListRootDnUsers) | **Get** /root-dn/root-dn-users | Returns a list of all Root DN User objects
[**UpdateRootDnUser**](RootDnUserApi.md#UpdateRootDnUser) | **Patch** /root-dn/root-dn-users/{root-dn-user-name} | Update an existing Root DN User by name



## AddRootDnUser

> RootDnUserResponse AddRootDnUser(ctx).AddRootDnUserRequest(addRootDnUserRequest).Execute()

Add a new Root DN User to the config

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
    addRootDnUserRequest := *openapiclient.NewAddRootDnUserRequest("UserName_example") // AddRootDnUserRequest | Create a new Root DN User in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RootDnUserApi.AddRootDnUser(context.Background()).AddRootDnUserRequest(addRootDnUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootDnUserApi.AddRootDnUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRootDnUser`: RootDnUserResponse
    fmt.Fprintf(os.Stdout, "Response from `RootDnUserApi.AddRootDnUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRootDnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addRootDnUserRequest** | [**AddRootDnUserRequest**](AddRootDnUserRequest.md) | Create a new Root DN User in the config | 

### Return type

[**RootDnUserResponse**](RootDnUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRootDnUser

> DeleteRootDnUser(ctx, rootDnUserName).Execute()

Delete a Root DN User

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
    rootDnUserName := "rootDnUserName_example" // string | Name of the Root DN User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RootDnUserApi.DeleteRootDnUser(context.Background(), rootDnUserName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootDnUserApi.DeleteRootDnUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rootDnUserName** | **string** | Name of the Root DN User | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRootDnUserRequest struct via the builder pattern


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


## GetRootDnUser

> RootDnUserResponse GetRootDnUser(ctx, rootDnUserName).Execute()

Returns a single Root DN User

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
    rootDnUserName := "rootDnUserName_example" // string | Name of the Root DN User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RootDnUserApi.GetRootDnUser(context.Background(), rootDnUserName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootDnUserApi.GetRootDnUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRootDnUser`: RootDnUserResponse
    fmt.Fprintf(os.Stdout, "Response from `RootDnUserApi.GetRootDnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rootDnUserName** | **string** | Name of the Root DN User | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootDnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RootDnUserResponse**](RootDnUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRootDnUsers

> RootDnUserListResponse ListRootDnUsers(ctx).Filter(filter).Execute()

Returns a list of all Root DN User objects

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
    resp, r, err := apiClient.RootDnUserApi.ListRootDnUsers(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootDnUserApi.ListRootDnUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRootDnUsers`: RootDnUserListResponse
    fmt.Fprintf(os.Stdout, "Response from `RootDnUserApi.ListRootDnUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRootDnUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**RootDnUserListResponse**](RootDnUserListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRootDnUser

> RootDnUserResponse UpdateRootDnUser(ctx, rootDnUserName).UpdateRequest(updateRequest).Execute()

Update an existing Root DN User by name

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
    rootDnUserName := "rootDnUserName_example" // string | Name of the Root DN User
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Root DN User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RootDnUserApi.UpdateRootDnUser(context.Background(), rootDnUserName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootDnUserApi.UpdateRootDnUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRootDnUser`: RootDnUserResponse
    fmt.Fprintf(os.Stdout, "Response from `RootDnUserApi.UpdateRootDnUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rootDnUserName** | **string** | Name of the Root DN User | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRootDnUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Root DN User | 

### Return type

[**RootDnUserResponse**](RootDnUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

