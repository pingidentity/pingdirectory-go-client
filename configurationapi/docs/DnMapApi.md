# \DnMapApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDnMap**](DnMapApi.md#AddDnMap) | **Post** /dn-maps | Add a new DN Map to the config
[**DeleteDnMap**](DnMapApi.md#DeleteDnMap) | **Delete** /dn-maps/{dn-map-name} | Delete a DN Map
[**GetDnMap**](DnMapApi.md#GetDnMap) | **Get** /dn-maps/{dn-map-name} | Returns a single DN Map
[**ListDnMaps**](DnMapApi.md#ListDnMaps) | **Get** /dn-maps | Returns a list of all DN Map objects
[**UpdateDnMap**](DnMapApi.md#UpdateDnMap) | **Patch** /dn-maps/{dn-map-name} | Update an existing DN Map by name



## AddDnMap

> DnMapResponse AddDnMap(ctx).AddDnMapRequest(addDnMapRequest).Execute()

Add a new DN Map to the config

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
    addDnMapRequest := *openapiclient.NewAddDnMapRequest("MapName_example", "FromDNPattern_example", "ToDNPattern_example") // AddDnMapRequest | Create a new DN Map in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnMapApi.AddDnMap(context.Background()).AddDnMapRequest(addDnMapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnMapApi.AddDnMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDnMap`: DnMapResponse
    fmt.Fprintf(os.Stdout, "Response from `DnMapApi.AddDnMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDnMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addDnMapRequest** | [**AddDnMapRequest**](AddDnMapRequest.md) | Create a new DN Map in the config | 

### Return type

[**DnMapResponse**](DnMapResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDnMap

> DeleteDnMap(ctx, dnMapName).Execute()

Delete a DN Map

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
    dnMapName := "dnMapName_example" // string | Name of the DN Map

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DnMapApi.DeleteDnMap(context.Background(), dnMapName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnMapApi.DeleteDnMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnMapName** | **string** | Name of the DN Map | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDnMapRequest struct via the builder pattern


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


## GetDnMap

> DnMapResponse GetDnMap(ctx, dnMapName).Execute()

Returns a single DN Map

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
    dnMapName := "dnMapName_example" // string | Name of the DN Map

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnMapApi.GetDnMap(context.Background(), dnMapName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnMapApi.GetDnMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDnMap`: DnMapResponse
    fmt.Fprintf(os.Stdout, "Response from `DnMapApi.GetDnMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnMapName** | **string** | Name of the DN Map | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DnMapResponse**](DnMapResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDnMaps

> DnMapListResponse ListDnMaps(ctx).Filter(filter).Execute()

Returns a list of all DN Map objects

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
    resp, r, err := apiClient.DnMapApi.ListDnMaps(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnMapApi.ListDnMaps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDnMaps`: DnMapListResponse
    fmt.Fprintf(os.Stdout, "Response from `DnMapApi.ListDnMaps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDnMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**DnMapListResponse**](DnMapListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDnMap

> DnMapResponse UpdateDnMap(ctx, dnMapName).UpdateRequest(updateRequest).Execute()

Update an existing DN Map by name

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
    dnMapName := "dnMapName_example" // string | Name of the DN Map
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing DN Map

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnMapApi.UpdateDnMap(context.Background(), dnMapName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnMapApi.UpdateDnMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDnMap`: DnMapResponse
    fmt.Fprintf(os.Stdout, "Response from `DnMapApi.UpdateDnMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnMapName** | **string** | Name of the DN Map | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDnMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing DN Map | 

### Return type

[**DnMapResponse**](DnMapResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

