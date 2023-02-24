# \ClientConnectionPolicyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddClientConnectionPolicy**](ClientConnectionPolicyApi.md#AddClientConnectionPolicy) | **Post** /client-connection-policies | Add a new Client Connection Policy to the config
[**DeleteClientConnectionPolicy**](ClientConnectionPolicyApi.md#DeleteClientConnectionPolicy) | **Delete** /client-connection-policies/{client-connection-policy-name} | Delete a Client Connection Policy
[**GetClientConnectionPolicy**](ClientConnectionPolicyApi.md#GetClientConnectionPolicy) | **Get** /client-connection-policies/{client-connection-policy-name} | Returns a single Client Connection Policy
[**UpdateClientConnectionPolicy**](ClientConnectionPolicyApi.md#UpdateClientConnectionPolicy) | **Patch** /client-connection-policies/{client-connection-policy-name} | Update an existing Client Connection Policy by name



## AddClientConnectionPolicy

> ClientConnectionPolicyResponse AddClientConnectionPolicy(ctx).AddClientConnectionPolicyRequest(addClientConnectionPolicyRequest).Execute()

Add a new Client Connection Policy to the config

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
    addClientConnectionPolicyRequest := *openapiclient.NewAddClientConnectionPolicyRequest("PolicyName_example", "PolicyID_example", false, int32(123)) // AddClientConnectionPolicyRequest | Create a new Client Connection Policy in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientConnectionPolicyApi.AddClientConnectionPolicy(context.Background()).AddClientConnectionPolicyRequest(addClientConnectionPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientConnectionPolicyApi.AddClientConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddClientConnectionPolicy`: ClientConnectionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientConnectionPolicyApi.AddClientConnectionPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddClientConnectionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addClientConnectionPolicyRequest** | [**AddClientConnectionPolicyRequest**](AddClientConnectionPolicyRequest.md) | Create a new Client Connection Policy in the config | 

### Return type

[**ClientConnectionPolicyResponse**](ClientConnectionPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientConnectionPolicy

> DeleteClientConnectionPolicy(ctx, clientConnectionPolicyName).Execute()

Delete a Client Connection Policy

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
    clientConnectionPolicyName := "clientConnectionPolicyName_example" // string | Name of the Client Connection Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientConnectionPolicyApi.DeleteClientConnectionPolicy(context.Background(), clientConnectionPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientConnectionPolicyApi.DeleteClientConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientConnectionPolicyName** | **string** | Name of the Client Connection Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientConnectionPolicyRequest struct via the builder pattern


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


## GetClientConnectionPolicy

> ClientConnectionPolicyResponse GetClientConnectionPolicy(ctx, clientConnectionPolicyName).Execute()

Returns a single Client Connection Policy

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
    clientConnectionPolicyName := "clientConnectionPolicyName_example" // string | Name of the Client Connection Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientConnectionPolicyApi.GetClientConnectionPolicy(context.Background(), clientConnectionPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientConnectionPolicyApi.GetClientConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientConnectionPolicy`: ClientConnectionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientConnectionPolicyApi.GetClientConnectionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientConnectionPolicyName** | **string** | Name of the Client Connection Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientConnectionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientConnectionPolicyResponse**](ClientConnectionPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientConnectionPolicy

> ClientConnectionPolicyResponse UpdateClientConnectionPolicy(ctx, clientConnectionPolicyName).UpdateRequest(updateRequest).Execute()

Update an existing Client Connection Policy by name

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
    clientConnectionPolicyName := "clientConnectionPolicyName_example" // string | Name of the Client Connection Policy
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Client Connection Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientConnectionPolicyApi.UpdateClientConnectionPolicy(context.Background(), clientConnectionPolicyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientConnectionPolicyApi.UpdateClientConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientConnectionPolicy`: ClientConnectionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientConnectionPolicyApi.UpdateClientConnectionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientConnectionPolicyName** | **string** | Name of the Client Connection Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientConnectionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Client Connection Policy | 

### Return type

[**ClientConnectionPolicyResponse**](ClientConnectionPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

