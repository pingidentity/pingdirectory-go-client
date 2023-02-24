# \HttpServletCrossOriginPolicyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHttpServletCrossOriginPolicy**](HttpServletCrossOriginPolicyApi.md#AddHttpServletCrossOriginPolicy) | **Post** /http-servlet-cross-origin-policies | Add a new HTTP Servlet Cross Origin Policy to the config
[**DeleteHttpServletCrossOriginPolicy**](HttpServletCrossOriginPolicyApi.md#DeleteHttpServletCrossOriginPolicy) | **Delete** /http-servlet-cross-origin-policies/{http-servlet-cross-origin-policy-name} | Delete a HTTP Servlet Cross Origin Policy
[**GetHttpServletCrossOriginPolicy**](HttpServletCrossOriginPolicyApi.md#GetHttpServletCrossOriginPolicy) | **Get** /http-servlet-cross-origin-policies/{http-servlet-cross-origin-policy-name} | Returns a single HTTP Servlet Cross Origin Policy
[**UpdateHttpServletCrossOriginPolicy**](HttpServletCrossOriginPolicyApi.md#UpdateHttpServletCrossOriginPolicy) | **Patch** /http-servlet-cross-origin-policies/{http-servlet-cross-origin-policy-name} | Update an existing HTTP Servlet Cross Origin Policy by name



## AddHttpServletCrossOriginPolicy

> HttpServletCrossOriginPolicyResponse AddHttpServletCrossOriginPolicy(ctx).AddHttpServletCrossOriginPolicyRequest(addHttpServletCrossOriginPolicyRequest).Execute()

Add a new HTTP Servlet Cross Origin Policy to the config

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
    addHttpServletCrossOriginPolicyRequest := *openapiclient.NewAddHttpServletCrossOriginPolicyRequest("PolicyName_example") // AddHttpServletCrossOriginPolicyRequest | Create a new HTTP Servlet Cross Origin Policy in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpServletCrossOriginPolicyApi.AddHttpServletCrossOriginPolicy(context.Background()).AddHttpServletCrossOriginPolicyRequest(addHttpServletCrossOriginPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServletCrossOriginPolicyApi.AddHttpServletCrossOriginPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddHttpServletCrossOriginPolicy`: HttpServletCrossOriginPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `HttpServletCrossOriginPolicyApi.AddHttpServletCrossOriginPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHttpServletCrossOriginPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addHttpServletCrossOriginPolicyRequest** | [**AddHttpServletCrossOriginPolicyRequest**](AddHttpServletCrossOriginPolicyRequest.md) | Create a new HTTP Servlet Cross Origin Policy in the config | 

### Return type

[**HttpServletCrossOriginPolicyResponse**](HttpServletCrossOriginPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHttpServletCrossOriginPolicy

> DeleteHttpServletCrossOriginPolicy(ctx, httpServletCrossOriginPolicyName).Execute()

Delete a HTTP Servlet Cross Origin Policy

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
    httpServletCrossOriginPolicyName := "httpServletCrossOriginPolicyName_example" // string | Name of the HTTP Servlet Cross Origin Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HttpServletCrossOriginPolicyApi.DeleteHttpServletCrossOriginPolicy(context.Background(), httpServletCrossOriginPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServletCrossOriginPolicyApi.DeleteHttpServletCrossOriginPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletCrossOriginPolicyName** | **string** | Name of the HTTP Servlet Cross Origin Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHttpServletCrossOriginPolicyRequest struct via the builder pattern


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


## GetHttpServletCrossOriginPolicy

> HttpServletCrossOriginPolicyResponse GetHttpServletCrossOriginPolicy(ctx, httpServletCrossOriginPolicyName).Execute()

Returns a single HTTP Servlet Cross Origin Policy

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
    httpServletCrossOriginPolicyName := "httpServletCrossOriginPolicyName_example" // string | Name of the HTTP Servlet Cross Origin Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpServletCrossOriginPolicyApi.GetHttpServletCrossOriginPolicy(context.Background(), httpServletCrossOriginPolicyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServletCrossOriginPolicyApi.GetHttpServletCrossOriginPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHttpServletCrossOriginPolicy`: HttpServletCrossOriginPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `HttpServletCrossOriginPolicyApi.GetHttpServletCrossOriginPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletCrossOriginPolicyName** | **string** | Name of the HTTP Servlet Cross Origin Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHttpServletCrossOriginPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HttpServletCrossOriginPolicyResponse**](HttpServletCrossOriginPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHttpServletCrossOriginPolicy

> HttpServletCrossOriginPolicyResponse UpdateHttpServletCrossOriginPolicy(ctx, httpServletCrossOriginPolicyName).UpdateRequest(updateRequest).Execute()

Update an existing HTTP Servlet Cross Origin Policy by name

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
    httpServletCrossOriginPolicyName := "httpServletCrossOriginPolicyName_example" // string | Name of the HTTP Servlet Cross Origin Policy
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing HTTP Servlet Cross Origin Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpServletCrossOriginPolicyApi.UpdateHttpServletCrossOriginPolicy(context.Background(), httpServletCrossOriginPolicyName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServletCrossOriginPolicyApi.UpdateHttpServletCrossOriginPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHttpServletCrossOriginPolicy`: HttpServletCrossOriginPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `HttpServletCrossOriginPolicyApi.UpdateHttpServletCrossOriginPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpServletCrossOriginPolicyName** | **string** | Name of the HTTP Servlet Cross Origin Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHttpServletCrossOriginPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing HTTP Servlet Cross Origin Policy | 

### Return type

[**HttpServletCrossOriginPolicyResponse**](HttpServletCrossOriginPolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

