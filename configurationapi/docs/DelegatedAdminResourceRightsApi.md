# \DelegatedAdminResourceRightsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDelegatedAdminResourceRights**](DelegatedAdminResourceRightsAPI.md#AddDelegatedAdminResourceRights) | **Post** /delegated-admin-rights/{delegated-admin-rights-name}/delegated-admin-resource-rights | Add a new Delegated Admin Resource Rights to the config
[**DeleteDelegatedAdminResourceRights**](DelegatedAdminResourceRightsAPI.md#DeleteDelegatedAdminResourceRights) | **Delete** /delegated-admin-rights/{delegated-admin-rights-name}/delegated-admin-resource-rights/{delegated-admin-resource-rights-name} | Delete a Delegated Admin Resource Rights
[**GetDelegatedAdminResourceRights**](DelegatedAdminResourceRightsAPI.md#GetDelegatedAdminResourceRights) | **Get** /delegated-admin-rights/{delegated-admin-rights-name}/delegated-admin-resource-rights/{delegated-admin-resource-rights-name} | Returns a single Delegated Admin Resource Rights
[**ListDelegatedAdminResourceRights**](DelegatedAdminResourceRightsAPI.md#ListDelegatedAdminResourceRights) | **Get** /delegated-admin-rights/{delegated-admin-rights-name}/delegated-admin-resource-rights | Returns a list of all Delegated Admin Resource Rights objects
[**UpdateDelegatedAdminResourceRights**](DelegatedAdminResourceRightsAPI.md#UpdateDelegatedAdminResourceRights) | **Patch** /delegated-admin-rights/{delegated-admin-rights-name}/delegated-admin-resource-rights/{delegated-admin-resource-rights-name} | Update an existing Delegated Admin Resource Rights by name



## AddDelegatedAdminResourceRights

> DelegatedAdminResourceRightsResponse AddDelegatedAdminResourceRights(ctx, delegatedAdminRightsName).AddDelegatedAdminResourceRightsRequest(addDelegatedAdminResourceRightsRequest).Execute()

Add a new Delegated Admin Resource Rights to the config

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
    addDelegatedAdminResourceRightsRequest := *openapiclient.NewAddDelegatedAdminResourceRightsRequest(false, "RestResourceType_example") // AddDelegatedAdminResourceRightsRequest | Create a new Delegated Admin Resource Rights in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminResourceRightsAPI.AddDelegatedAdminResourceRights(context.Background(), delegatedAdminRightsName).AddDelegatedAdminResourceRightsRequest(addDelegatedAdminResourceRightsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminResourceRightsAPI.AddDelegatedAdminResourceRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDelegatedAdminResourceRights`: DelegatedAdminResourceRightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminResourceRightsAPI.AddDelegatedAdminResourceRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDelegatedAdminResourceRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addDelegatedAdminResourceRightsRequest** | [**AddDelegatedAdminResourceRightsRequest**](AddDelegatedAdminResourceRightsRequest.md) | Create a new Delegated Admin Resource Rights in the config | 

### Return type

[**DelegatedAdminResourceRightsResponse**](DelegatedAdminResourceRightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDelegatedAdminResourceRights

> DeleteDelegatedAdminResourceRights(ctx, delegatedAdminResourceRightsName, delegatedAdminRightsName).Execute()

Delete a Delegated Admin Resource Rights

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
    delegatedAdminResourceRightsName := "delegatedAdminResourceRightsName_example" // string | Name of the Delegated Admin Resource Rights
    delegatedAdminRightsName := "delegatedAdminRightsName_example" // string | Name of the Delegated Admin Rights

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DelegatedAdminResourceRightsAPI.DeleteDelegatedAdminResourceRights(context.Background(), delegatedAdminResourceRightsName, delegatedAdminRightsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminResourceRightsAPI.DeleteDelegatedAdminResourceRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminResourceRightsName** | **string** | Name of the Delegated Admin Resource Rights | 
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDelegatedAdminResourceRightsRequest struct via the builder pattern


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


## GetDelegatedAdminResourceRights

> DelegatedAdminResourceRightsResponse GetDelegatedAdminResourceRights(ctx, delegatedAdminResourceRightsName, delegatedAdminRightsName).Execute()

Returns a single Delegated Admin Resource Rights

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
    delegatedAdminResourceRightsName := "delegatedAdminResourceRightsName_example" // string | Name of the Delegated Admin Resource Rights
    delegatedAdminRightsName := "delegatedAdminRightsName_example" // string | Name of the Delegated Admin Rights

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminResourceRightsAPI.GetDelegatedAdminResourceRights(context.Background(), delegatedAdminResourceRightsName, delegatedAdminRightsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminResourceRightsAPI.GetDelegatedAdminResourceRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDelegatedAdminResourceRights`: DelegatedAdminResourceRightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminResourceRightsAPI.GetDelegatedAdminResourceRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminResourceRightsName** | **string** | Name of the Delegated Admin Resource Rights | 
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDelegatedAdminResourceRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DelegatedAdminResourceRightsResponse**](DelegatedAdminResourceRightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDelegatedAdminResourceRights

> DelegatedAdminResourceRightsListResponse ListDelegatedAdminResourceRights(ctx, delegatedAdminRightsName).Filter(filter).Execute()

Returns a list of all Delegated Admin Resource Rights objects

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
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminResourceRightsAPI.ListDelegatedAdminResourceRights(context.Background(), delegatedAdminRightsName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminResourceRightsAPI.ListDelegatedAdminResourceRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDelegatedAdminResourceRights`: DelegatedAdminResourceRightsListResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminResourceRightsAPI.ListDelegatedAdminResourceRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDelegatedAdminResourceRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**DelegatedAdminResourceRightsListResponse**](DelegatedAdminResourceRightsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDelegatedAdminResourceRights

> DelegatedAdminResourceRightsResponse UpdateDelegatedAdminResourceRights(ctx, delegatedAdminResourceRightsName, delegatedAdminRightsName).UpdateRequest(updateRequest).Execute()

Update an existing Delegated Admin Resource Rights by name

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
    delegatedAdminResourceRightsName := "delegatedAdminResourceRightsName_example" // string | Name of the Delegated Admin Resource Rights
    delegatedAdminRightsName := "delegatedAdminRightsName_example" // string | Name of the Delegated Admin Rights
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Delegated Admin Resource Rights

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelegatedAdminResourceRightsAPI.UpdateDelegatedAdminResourceRights(context.Background(), delegatedAdminResourceRightsName, delegatedAdminRightsName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdminResourceRightsAPI.UpdateDelegatedAdminResourceRights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDelegatedAdminResourceRights`: DelegatedAdminResourceRightsResponse
    fmt.Fprintf(os.Stdout, "Response from `DelegatedAdminResourceRightsAPI.UpdateDelegatedAdminResourceRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**delegatedAdminResourceRightsName** | **string** | Name of the Delegated Admin Resource Rights | 
**delegatedAdminRightsName** | **string** | Name of the Delegated Admin Rights | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDelegatedAdminResourceRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Delegated Admin Resource Rights | 

### Return type

[**DelegatedAdminResourceRightsResponse**](DelegatedAdminResourceRightsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

