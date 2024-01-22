# \RootDseBackendAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRootDseBackend**](RootDseBackendAPI.md#GetRootDseBackend) | **Get** /root-dse-backend | Returns a single Root DSE Backend
[**UpdateRootDseBackend**](RootDseBackendAPI.md#UpdateRootDseBackend) | **Patch** /root-dse-backend | Update an existing Root DSE Backend by name



## GetRootDseBackend

> RootDseBackendResponse GetRootDseBackend(ctx).Execute()

Returns a single Root DSE Backend

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RootDseBackendAPI.GetRootDseBackend(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootDseBackendAPI.GetRootDseBackend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRootDseBackend`: RootDseBackendResponse
    fmt.Fprintf(os.Stdout, "Response from `RootDseBackendAPI.GetRootDseBackend`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootDseBackendRequest struct via the builder pattern


### Return type

[**RootDseBackendResponse**](RootDseBackendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRootDseBackend

> RootDseBackendResponse UpdateRootDseBackend(ctx).UpdateRequest(updateRequest).Execute()

Update an existing Root DSE Backend by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Root DSE Backend

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RootDseBackendAPI.UpdateRootDseBackend(context.Background()).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RootDseBackendAPI.UpdateRootDseBackend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRootDseBackend`: RootDseBackendResponse
    fmt.Fprintf(os.Stdout, "Response from `RootDseBackendAPI.UpdateRootDseBackend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRootDseBackendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Root DSE Backend | 

### Return type

[**RootDseBackendResponse**](RootDseBackendResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

