# \ClientConnectionPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddClientConnectionPolicy**](ClientConnectionPolicyAPI.md#AddClientConnectionPolicy) | **Post** /client-connection-policies | Add a new Client Connection Policy to the config
[**DeleteClientConnectionPolicy**](ClientConnectionPolicyAPI.md#DeleteClientConnectionPolicy) | **Delete** /client-connection-policies/{client-connection-policy-name} | Delete a Client Connection Policy
[**GetClientConnectionPolicy**](ClientConnectionPolicyAPI.md#GetClientConnectionPolicy) | **Get** /client-connection-policies/{client-connection-policy-name} | Returns a single Client Connection Policy
[**ListClientConnectionPolicies**](ClientConnectionPolicyAPI.md#ListClientConnectionPolicies) | **Get** /client-connection-policies | Returns a list of all Client Connection Policy objects
[**UpdateClientConnectionPolicy**](ClientConnectionPolicyAPI.md#UpdateClientConnectionPolicy) | **Patch** /client-connection-policies/{client-connection-policy-name} | Update an existing Client Connection Policy by name



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
    addClientConnectionPolicyRequest := *openapiclient.NewAddClientConnectionPolicyRequest("PolicyID_example", false, int64(123), "PolicyName_example") // AddClientConnectionPolicyRequest | Create a new Client Connection Policy in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientConnectionPolicyAPI.AddClientConnectionPolicy(context.Background()).AddClientConnectionPolicyRequest(addClientConnectionPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientConnectionPolicyAPI.AddClientConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddClientConnectionPolicy`: ClientConnectionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientConnectionPolicyAPI.AddClientConnectionPolicy`: %v\n", resp)
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
    r, err := apiClient.ClientConnectionPolicyAPI.DeleteClientConnectionPolicy(context.Background(), clientConnectionPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientConnectionPolicyAPI.DeleteClientConnectionPolicy``: %v\n", err)
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
    resp, r, err := apiClient.ClientConnectionPolicyAPI.GetClientConnectionPolicy(context.Background(), clientConnectionPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientConnectionPolicyAPI.GetClientConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientConnectionPolicy`: ClientConnectionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientConnectionPolicyAPI.GetClientConnectionPolicy`: %v\n", resp)
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


## ListClientConnectionPolicies

> ClientConnectionPolicyListResponse ListClientConnectionPolicies(ctx).Filter(filter).Execute()

Returns a list of all Client Connection Policy objects

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
    resp, r, err := apiClient.ClientConnectionPolicyAPI.ListClientConnectionPolicies(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientConnectionPolicyAPI.ListClientConnectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClientConnectionPolicies`: ClientConnectionPolicyListResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientConnectionPolicyAPI.ListClientConnectionPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClientConnectionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ClientConnectionPolicyListResponse**](ClientConnectionPolicyListResponse.md)

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
    resp, r, err := apiClient.ClientConnectionPolicyAPI.UpdateClientConnectionPolicy(context.Background(), clientConnectionPolicyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientConnectionPolicyAPI.UpdateClientConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientConnectionPolicy`: ClientConnectionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientConnectionPolicyAPI.UpdateClientConnectionPolicy`: %v\n", resp)
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

