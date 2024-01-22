# \DelegatedAdminRightsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDelegatedAdminRights**](DelegatedAdminRightsAPI.md#AddDelegatedAdminRights) | **Post** /delegated-admin-rights | Add a new Delegated Admin Rights to the config
[**DeleteDelegatedAdminRights**](DelegatedAdminRightsAPI.md#DeleteDelegatedAdminRights) | **Delete** /delegated-admin-rights/{delegated-admin-rights-name} | Delete a Delegated Admin Rights
[**GetDelegatedAdminRights**](DelegatedAdminRightsAPI.md#GetDelegatedAdminRights) | **Get** /delegated-admin-rights/{delegated-admin-rights-name} | Returns a single Delegated Admin Rights
[**ListDelegatedAdminRights**](DelegatedAdminRightsAPI.md#ListDelegatedAdminRights) | **Get** /delegated-admin-rights | Returns a list of all Delegated Admin Rights objects
[**UpdateDelegatedAdminRights**](DelegatedAdminRightsAPI.md#UpdateDelegatedAdminRights) | **Patch** /delegated-admin-rights/{delegated-admin-rights-name} | Update an existing Delegated Admin Rights by name



## AddDelegatedAdminRights

> DelegatedAdminRightsResponse AddDelegatedAdminRights(ctx).AddDelegatedAdminRightsRequest(addDelegatedAdminRightsRequest).Execute()

Add a new Delegated Admin Rights to the config

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
    addDelegatedAdminRightsRequest := *openapiclient.NewAddDelegatedAdminRightsRequest(false, "RightsName_example") // AddDelegatedAdminRightsRequest | Create a new Delegated Admin Rights in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminRightsAPI.AddDelegatedAdminRights(context.Background()).AddDelegatedAdminRightsRequest(addDelegatedAdminRightsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminRightsAPI.AddDelegatedAdminRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDelegatedAdminRights`: DelegatedAdminRightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminRightsAPI.AddDelegatedAdminRights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDelegatedAdminRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addDelegatedAdminRightsRequest** | [**AddDelegatedAdminRightsRequest**](AddDelegatedAdminRightsRequest.md) | Create a new Delegated Admin Rights in the config | 

### Return type

[**DelegatedAdminRightsResponse**](DelegatedAdminRightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDelegatedAdminRights

> DeleteDelegatedAdminRights(ctx, delegatedAdminRightsName).Execute()

Delete a Delegated Admin Rights

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
    delegatedAdminRightsName := "delegatedAdminRightsName_example" // string | Name of the Delegated Admin Rights

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DelegatedAdminRightsAPI.DeleteDelegatedAdminRights(context.Background(), delegatedAdminRightsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminRightsAPI.DeleteDelegatedAdminRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDelegatedAdminRightsRequest struct via the builder pattern


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


## GetDelegatedAdminRights

> DelegatedAdminRightsResponse GetDelegatedAdminRights(ctx, delegatedAdminRightsName).Execute()

Returns a single Delegated Admin Rights

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
    delegatedAdminRightsName := "delegatedAdminRightsName_example" // string | Name of the Delegated Admin Rights

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminRightsAPI.GetDelegatedAdminRights(context.Background(), delegatedAdminRightsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminRightsAPI.GetDelegatedAdminRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDelegatedAdminRights`: DelegatedAdminRightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminRightsAPI.GetDelegatedAdminRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDelegatedAdminRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DelegatedAdminRightsResponse**](DelegatedAdminRightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDelegatedAdminRights

> DelegatedAdminRightsListResponse ListDelegatedAdminRights(ctx).Filter(filter).Execute()

Returns a list of all Delegated Admin Rights objects

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
    resp, r, err := apiClient.DelegatedAdminRightsAPI.ListDelegatedAdminRights(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminRightsAPI.ListDelegatedAdminRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDelegatedAdminRights`: DelegatedAdminRightsListResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminRightsAPI.ListDelegatedAdminRights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDelegatedAdminRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**DelegatedAdminRightsListResponse**](DelegatedAdminRightsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDelegatedAdminRights

> DelegatedAdminRightsResponse UpdateDelegatedAdminRights(ctx, delegatedAdminRightsName).UpdateRequest(updateRequest).Execute()

Update an existing Delegated Admin Rights by name

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
    delegatedAdminRightsName := "delegatedAdminRightsName_example" // string | Name of the Delegated Admin Rights
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Delegated Admin Rights

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminRightsAPI.UpdateDelegatedAdminRights(context.Background(), delegatedAdminRightsName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminRightsAPI.UpdateDelegatedAdminRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDelegatedAdminRights`: DelegatedAdminRightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminRightsAPI.UpdateDelegatedAdminRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDelegatedAdminRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Delegated Admin Rights | 

### Return type

[**DelegatedAdminRightsResponse**](DelegatedAdminRightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

