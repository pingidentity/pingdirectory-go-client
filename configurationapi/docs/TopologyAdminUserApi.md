# \TopologyAdminUserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTopologyAdminUser**](TopologyAdminUserAPI.md#AddTopologyAdminUser) | **Post** /topology-admin-users | Add a new Topology Admin User to the config
[**DeleteTopologyAdminUser**](TopologyAdminUserAPI.md#DeleteTopologyAdminUser) | **Delete** /topology-admin-users/{topology-admin-user-name} | Delete a Topology Admin User
[**GetTopologyAdminUser**](TopologyAdminUserAPI.md#GetTopologyAdminUser) | **Get** /topology-admin-users/{topology-admin-user-name} | Returns a single Topology Admin User
[**ListTopologyAdminUsers**](TopologyAdminUserAPI.md#ListTopologyAdminUsers) | **Get** /topology-admin-users | Returns a list of all Topology Admin User objects
[**UpdateTopologyAdminUser**](TopologyAdminUserAPI.md#UpdateTopologyAdminUser) | **Patch** /topology-admin-users/{topology-admin-user-name} | Update an existing Topology Admin User by name



## AddTopologyAdminUser

> TopologyAdminUserResponse AddTopologyAdminUser(ctx).AddTopologyAdminUserRequest(addTopologyAdminUserRequest).Execute()

Add a new Topology Admin User to the config

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
    addTopologyAdminUserRequest := *openapiclient.NewAddTopologyAdminUserRequest("UserName_example") // AddTopologyAdminUserRequest | Create a new Topology Admin User in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyAdminUserAPI.AddTopologyAdminUser(context.Background()).AddTopologyAdminUserRequest(addTopologyAdminUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAdminUserAPI.AddTopologyAdminUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTopologyAdminUser`: TopologyAdminUserResponse
    fmt.Fprintf(os.Stdout, "Response from `TopologyAdminUserAPI.AddTopologyAdminUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTopologyAdminUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addTopologyAdminUserRequest** | [**AddTopologyAdminUserRequest**](AddTopologyAdminUserRequest.md) | Create a new Topology Admin User in the config | 

### Return type

[**TopologyAdminUserResponse**](TopologyAdminUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTopologyAdminUser

> DeleteTopologyAdminUser(ctx, topologyAdminUserName).Execute()

Delete a Topology Admin User

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
    topologyAdminUserName := "topologyAdminUserName_example" // string | Name of the Topology Admin User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TopologyAdminUserAPI.DeleteTopologyAdminUser(context.Background(), topologyAdminUserName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAdminUserAPI.DeleteTopologyAdminUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topologyAdminUserName** | **string** | Name of the Topology Admin User | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTopologyAdminUserRequest struct via the builder pattern


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


## GetTopologyAdminUser

> TopologyAdminUserResponse GetTopologyAdminUser(ctx, topologyAdminUserName).Execute()

Returns a single Topology Admin User

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
    topologyAdminUserName := "topologyAdminUserName_example" // string | Name of the Topology Admin User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyAdminUserAPI.GetTopologyAdminUser(context.Background(), topologyAdminUserName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAdminUserAPI.GetTopologyAdminUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopologyAdminUser`: TopologyAdminUserResponse
    fmt.Fprintf(os.Stdout, "Response from `TopologyAdminUserAPI.GetTopologyAdminUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topologyAdminUserName** | **string** | Name of the Topology Admin User | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyAdminUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TopologyAdminUserResponse**](TopologyAdminUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTopologyAdminUsers

> TopologyAdminUserListResponse ListTopologyAdminUsers(ctx).Filter(filter).Execute()

Returns a list of all Topology Admin User objects

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
    resp, r, err := apiClient.TopologyAdminUserAPI.ListTopologyAdminUsers(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAdminUserAPI.ListTopologyAdminUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTopologyAdminUsers`: TopologyAdminUserListResponse
    fmt.Fprintf(os.Stdout, "Response from `TopologyAdminUserAPI.ListTopologyAdminUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTopologyAdminUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**TopologyAdminUserListResponse**](TopologyAdminUserListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTopologyAdminUser

> TopologyAdminUserResponse UpdateTopologyAdminUser(ctx, topologyAdminUserName).UpdateRequest(updateRequest).Execute()

Update an existing Topology Admin User by name

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
    topologyAdminUserName := "topologyAdminUserName_example" // string | Name of the Topology Admin User
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Topology Admin User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyAdminUserAPI.UpdateTopologyAdminUser(context.Background(), topologyAdminUserName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAdminUserAPI.UpdateTopologyAdminUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTopologyAdminUser`: TopologyAdminUserResponse
    fmt.Fprintf(os.Stdout, "Response from `TopologyAdminUserAPI.UpdateTopologyAdminUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topologyAdminUserName** | **string** | Name of the Topology Admin User | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTopologyAdminUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Topology Admin User | 

### Return type

[**TopologyAdminUserResponse**](TopologyAdminUserResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

