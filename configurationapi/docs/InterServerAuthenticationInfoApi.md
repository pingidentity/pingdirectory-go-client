# \InterServerAuthenticationInfoAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInterServerAuthenticationInfo**](InterServerAuthenticationInfoAPI.md#GetInterServerAuthenticationInfo) | **Get** /server-instances/{server-instance-name}/server-instance-listeners/{server-instance-listener-name}/inter-server-authentication-info/{inter-server-authentication-info-name} | Returns a single Inter Server Authentication Info
[**ListInterServerAuthenticationInfo**](InterServerAuthenticationInfoAPI.md#ListInterServerAuthenticationInfo) | **Get** /server-instances/{server-instance-name}/server-instance-listeners/{server-instance-listener-name}/inter-server-authentication-info | Returns a list of all Inter Server Authentication Info objects
[**UpdateInterServerAuthenticationInfo**](InterServerAuthenticationInfoAPI.md#UpdateInterServerAuthenticationInfo) | **Patch** /server-instances/{server-instance-name}/server-instance-listeners/{server-instance-listener-name}/inter-server-authentication-info/{inter-server-authentication-info-name} | Update an existing Inter Server Authentication Info by name



## GetInterServerAuthenticationInfo

> GetInterServerAuthenticationInfo200Response GetInterServerAuthenticationInfo(ctx, interServerAuthenticationInfoName, serverInstanceListenerName, serverInstanceName).Execute()

Returns a single Inter Server Authentication Info

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
    interServerAuthenticationInfoName := "interServerAuthenticationInfoName_example" // string | Name of the Inter Server Authentication Info
    serverInstanceListenerName := "serverInstanceListenerName_example" // string | Name of the Server Instance Listener
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterServerAuthenticationInfoAPI.GetInterServerAuthenticationInfo(context.Background(), interServerAuthenticationInfoName, serverInstanceListenerName, serverInstanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterServerAuthenticationInfoAPI.GetInterServerAuthenticationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterServerAuthenticationInfo`: GetInterServerAuthenticationInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `InterServerAuthenticationInfoAPI.GetInterServerAuthenticationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interServerAuthenticationInfoName** | **string** | Name of the Inter Server Authentication Info | 
**serverInstanceListenerName** | **string** | Name of the Server Instance Listener | 
**serverInstanceName** | **string** | Name of the Server Instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterServerAuthenticationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetInterServerAuthenticationInfo200Response**](GetInterServerAuthenticationInfo200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInterServerAuthenticationInfo

> InterServerAuthenticationInfoListResponse ListInterServerAuthenticationInfo(ctx, serverInstanceListenerName, serverInstanceName).Filter(filter).Execute()

Returns a list of all Inter Server Authentication Info objects

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
    serverInstanceListenerName := "serverInstanceListenerName_example" // string | Name of the Server Instance Listener
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterServerAuthenticationInfoAPI.ListInterServerAuthenticationInfo(context.Background(), serverInstanceListenerName, serverInstanceName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterServerAuthenticationInfoAPI.ListInterServerAuthenticationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInterServerAuthenticationInfo`: InterServerAuthenticationInfoListResponse
    fmt.Fprintf(os.Stdout, "Response from `InterServerAuthenticationInfoAPI.ListInterServerAuthenticationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverInstanceListenerName** | **string** | Name of the Server Instance Listener | 
**serverInstanceName** | **string** | Name of the Server Instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInterServerAuthenticationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | SCIM filter | 

### Return type

[**InterServerAuthenticationInfoListResponse**](InterServerAuthenticationInfoListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInterServerAuthenticationInfo

> GetInterServerAuthenticationInfo200Response UpdateInterServerAuthenticationInfo(ctx, interServerAuthenticationInfoName, serverInstanceListenerName, serverInstanceName).UpdateRequest(updateRequest).Execute()

Update an existing Inter Server Authentication Info by name

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
    interServerAuthenticationInfoName := "interServerAuthenticationInfoName_example" // string | Name of the Inter Server Authentication Info
    serverInstanceListenerName := "serverInstanceListenerName_example" // string | Name of the Server Instance Listener
    serverInstanceName := "serverInstanceName_example" // string | Name of the Server Instance
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Inter Server Authentication Info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterServerAuthenticationInfoAPI.UpdateInterServerAuthenticationInfo(context.Background(), interServerAuthenticationInfoName, serverInstanceListenerName, serverInstanceName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterServerAuthenticationInfoAPI.UpdateInterServerAuthenticationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInterServerAuthenticationInfo`: GetInterServerAuthenticationInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `InterServerAuthenticationInfoAPI.UpdateInterServerAuthenticationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interServerAuthenticationInfoName** | **string** | Name of the Inter Server Authentication Info | 
**serverInstanceListenerName** | **string** | Name of the Server Instance Listener | 
**serverInstanceName** | **string** | Name of the Server Instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInterServerAuthenticationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Inter Server Authentication Info | 

### Return type

[**GetInterServerAuthenticationInfo200Response**](GetInterServerAuthenticationInfo200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

