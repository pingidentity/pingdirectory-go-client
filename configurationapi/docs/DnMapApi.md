# \DnMapAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDnMap**](DnMapAPI.md#AddDnMap) | **Post** /dn-maps | Add a new DN Map to the config
[**DeleteDnMap**](DnMapAPI.md#DeleteDnMap) | **Delete** /dn-maps/{dn-map-name} | Delete a DN Map
[**GetDnMap**](DnMapAPI.md#GetDnMap) | **Get** /dn-maps/{dn-map-name} | Returns a single DN Map
[**ListDnMaps**](DnMapAPI.md#ListDnMaps) | **Get** /dn-maps | Returns a list of all DN Map objects
[**UpdateDnMap**](DnMapAPI.md#UpdateDnMap) | **Patch** /dn-maps/{dn-map-name} | Update an existing DN Map by name



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
    addDnMapRequest := *openapiclient.NewAddDnMapRequest("FromDNPattern_example", "ToDNPattern_example", "MapName_example") // AddDnMapRequest | Create a new DN Map in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DnMapAPI.AddDnMap(context.Background()).AddDnMapRequest(addDnMapRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnMapAPI.AddDnMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDnMap`: DnMapResponse
    fmt.Fprintf(os.Stdout, "Response from `DnMapAPI.AddDnMap`: %v\n", resp)
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
    r, err := apiClient.DnMapAPI.DeleteDnMap(context.Background(), dnMapName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnMapAPI.DeleteDnMap``: %v\n", err)
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
    resp, r, err := apiClient.DnMapAPI.GetDnMap(context.Background(), dnMapName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnMapAPI.GetDnMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDnMap`: DnMapResponse
    fmt.Fprintf(os.Stdout, "Response from `DnMapAPI.GetDnMap`: %v\n", resp)
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
    resp, r, err := apiClient.DnMapAPI.ListDnMaps(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnMapAPI.ListDnMaps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDnMaps`: DnMapListResponse
    fmt.Fprintf(os.Stdout, "Response from `DnMapAPI.ListDnMaps`: %v\n", resp)
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
    resp, r, err := apiClient.DnMapAPI.UpdateDnMap(context.Background(), dnMapName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DnMapAPI.UpdateDnMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDnMap`: DnMapResponse
    fmt.Fprintf(os.Stdout, "Response from `DnMapAPI.UpdateDnMap`: %v\n", resp)
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

