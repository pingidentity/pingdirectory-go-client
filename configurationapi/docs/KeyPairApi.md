# \KeyPairAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKeyPair**](KeyPairAPI.md#AddKeyPair) | **Post** /key-pairs | Add a new Key Pair to the config
[**DeleteKeyPair**](KeyPairAPI.md#DeleteKeyPair) | **Delete** /key-pairs/{key-pair-name} | Delete a Key Pair
[**GetKeyPair**](KeyPairAPI.md#GetKeyPair) | **Get** /key-pairs/{key-pair-name} | Returns a single Key Pair
[**ListKeyPairs**](KeyPairAPI.md#ListKeyPairs) | **Get** /key-pairs | Returns a list of all Key Pair objects
[**UpdateKeyPair**](KeyPairAPI.md#UpdateKeyPair) | **Patch** /key-pairs/{key-pair-name} | Update an existing Key Pair by name



## AddKeyPair

> KeyPairResponse AddKeyPair(ctx).AddKeyPairRequest(addKeyPairRequest).Execute()

Add a new Key Pair to the config

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
    addKeyPairRequest := *openapiclient.NewAddKeyPairRequest("PairName_example") // AddKeyPairRequest | Create a new Key Pair in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairAPI.AddKeyPair(context.Background()).AddKeyPairRequest(addKeyPairRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairAPI.AddKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddKeyPair`: KeyPairResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyPairAPI.AddKeyPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addKeyPairRequest** | [**AddKeyPairRequest**](AddKeyPairRequest.md) | Create a new Key Pair in the config | 

### Return type

[**KeyPairResponse**](KeyPairResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeyPair

> DeleteKeyPair(ctx, keyPairName).Execute()

Delete a Key Pair

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
    keyPairName := "keyPairName_example" // string | Name of the Key Pair

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyPairAPI.DeleteKeyPair(context.Background(), keyPairName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairAPI.DeleteKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyPairName** | **string** | Name of the Key Pair | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyPairRequest struct via the builder pattern


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


## GetKeyPair

> KeyPairResponse GetKeyPair(ctx, keyPairName).Execute()

Returns a single Key Pair

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
    keyPairName := "keyPairName_example" // string | Name of the Key Pair

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairAPI.GetKeyPair(context.Background(), keyPairName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairAPI.GetKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyPair`: KeyPairResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyPairAPI.GetKeyPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyPairName** | **string** | Name of the Key Pair | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeyPairResponse**](KeyPairResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeyPairs

> KeyPairListResponse ListKeyPairs(ctx).Filter(filter).Execute()

Returns a list of all Key Pair objects

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
    resp, r, err := apiClient.KeyPairAPI.ListKeyPairs(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairAPI.ListKeyPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKeyPairs`: KeyPairListResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyPairAPI.ListKeyPairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKeyPairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**KeyPairListResponse**](KeyPairListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeyPair

> KeyPairResponse UpdateKeyPair(ctx, keyPairName).UpdateRequest(updateRequest).Execute()

Update an existing Key Pair by name

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
    keyPairName := "keyPairName_example" // string | Name of the Key Pair
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Key Pair

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyPairAPI.UpdateKeyPair(context.Background(), keyPairName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyPairAPI.UpdateKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKeyPair`: KeyPairResponse
    fmt.Fprintf(os.Stdout, "Response from `KeyPairAPI.UpdateKeyPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyPairName** | **string** | Name of the Key Pair | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Key Pair | 

### Return type

[**KeyPairResponse**](KeyPairResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

