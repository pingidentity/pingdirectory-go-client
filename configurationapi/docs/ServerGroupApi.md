# \ServerGroupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServerGroup**](ServerGroupApi.md#AddServerGroup) | **Post** /server-groups | Add a new Server Group to the config
[**DeleteServerGroup**](ServerGroupApi.md#DeleteServerGroup) | **Delete** /server-groups/{server-group-name} | Delete a Server Group
[**GetServerGroup**](ServerGroupApi.md#GetServerGroup) | **Get** /server-groups/{server-group-name} | Returns a single Server Group
[**ListServerGroups**](ServerGroupApi.md#ListServerGroups) | **Get** /server-groups | Returns a list of all Server Group objects
[**UpdateServerGroup**](ServerGroupApi.md#UpdateServerGroup) | **Patch** /server-groups/{server-group-name} | Update an existing Server Group by name



## AddServerGroup

> ServerGroupResponse AddServerGroup(ctx).AddServerGroupRequest(addServerGroupRequest).Execute()

Add a new Server Group to the config

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
    addServerGroupRequest := *openapiclient.NewAddServerGroupRequest("GroupName_example") // AddServerGroupRequest | Create a new Server Group in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.AddServerGroup(context.Background()).AddServerGroupRequest(addServerGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.AddServerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddServerGroup`: ServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.AddServerGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addServerGroupRequest** | [**AddServerGroupRequest**](AddServerGroupRequest.md) | Create a new Server Group in the config | 

### Return type

[**ServerGroupResponse**](ServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServerGroup

> DeleteServerGroup(ctx, serverGroupName).Execute()

Delete a Server Group

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
    serverGroupName := "serverGroupName_example" // string | Name of the Server Group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServerGroupApi.DeleteServerGroup(context.Background(), serverGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.DeleteServerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverGroupName** | **string** | Name of the Server Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerGroupRequest struct via the builder pattern


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


## GetServerGroup

> ServerGroupResponse GetServerGroup(ctx, serverGroupName).Execute()

Returns a single Server Group

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
    serverGroupName := "serverGroupName_example" // string | Name of the Server Group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.GetServerGroup(context.Background(), serverGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.GetServerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerGroup`: ServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.GetServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverGroupName** | **string** | Name of the Server Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerGroupResponse**](ServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerGroups

> ServerGroupListResponse ListServerGroups(ctx).Filter(filter).Execute()

Returns a list of all Server Group objects

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
    resp, r, err := apiClient.ServerGroupApi.ListServerGroups(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.ListServerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServerGroups`: ServerGroupListResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.ListServerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ServerGroupListResponse**](ServerGroupListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServerGroup

> ServerGroupResponse UpdateServerGroup(ctx, serverGroupName).UpdateRequest(updateRequest).Execute()

Update an existing Server Group by name

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
    serverGroupName := "serverGroupName_example" // string | Name of the Server Group
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Server Group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerGroupApi.UpdateServerGroup(context.Background(), serverGroupName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerGroupApi.UpdateServerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServerGroup`: ServerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ServerGroupApi.UpdateServerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverGroupName** | **string** | Name of the Server Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Server Group | 

### Return type

[**ServerGroupResponse**](ServerGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

