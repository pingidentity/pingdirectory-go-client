# \TrustedCertificateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTrustedCertificate**](TrustedCertificateAPI.md#AddTrustedCertificate) | **Post** /trusted-certificates | Add a new Trusted Certificate to the config
[**DeleteTrustedCertificate**](TrustedCertificateAPI.md#DeleteTrustedCertificate) | **Delete** /trusted-certificates/{trusted-certificate-name} | Delete a Trusted Certificate
[**GetTrustedCertificate**](TrustedCertificateAPI.md#GetTrustedCertificate) | **Get** /trusted-certificates/{trusted-certificate-name} | Returns a single Trusted Certificate
[**ListTrustedCertificates**](TrustedCertificateAPI.md#ListTrustedCertificates) | **Get** /trusted-certificates | Returns a list of all Trusted Certificate objects
[**UpdateTrustedCertificate**](TrustedCertificateAPI.md#UpdateTrustedCertificate) | **Patch** /trusted-certificates/{trusted-certificate-name} | Update an existing Trusted Certificate by name



## AddTrustedCertificate

> TrustedCertificateResponse AddTrustedCertificate(ctx).AddTrustedCertificateRequest(addTrustedCertificateRequest).Execute()

Add a new Trusted Certificate to the config

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
    addTrustedCertificateRequest := *openapiclient.NewAddTrustedCertificateRequest("Certificate_example", "CertificateName_example") // AddTrustedCertificateRequest | Create a new Trusted Certificate in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedCertificateAPI.AddTrustedCertificate(context.Background()).AddTrustedCertificateRequest(addTrustedCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateAPI.AddTrustedCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTrustedCertificate`: TrustedCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TrustedCertificateAPI.AddTrustedCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTrustedCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addTrustedCertificateRequest** | [**AddTrustedCertificateRequest**](AddTrustedCertificateRequest.md) | Create a new Trusted Certificate in the config | 

### Return type

[**TrustedCertificateResponse**](TrustedCertificateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrustedCertificate

> DeleteTrustedCertificate(ctx, trustedCertificateName).Execute()

Delete a Trusted Certificate

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
    trustedCertificateName := "trustedCertificateName_example" // string | Name of the Trusted Certificate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TrustedCertificateAPI.DeleteTrustedCertificate(context.Background(), trustedCertificateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateAPI.DeleteTrustedCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedCertificateName** | **string** | Name of the Trusted Certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrustedCertificateRequest struct via the builder pattern


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


## GetTrustedCertificate

> TrustedCertificateResponse GetTrustedCertificate(ctx, trustedCertificateName).Execute()

Returns a single Trusted Certificate

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
    trustedCertificateName := "trustedCertificateName_example" // string | Name of the Trusted Certificate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedCertificateAPI.GetTrustedCertificate(context.Background(), trustedCertificateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateAPI.GetTrustedCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrustedCertificate`: TrustedCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TrustedCertificateAPI.GetTrustedCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedCertificateName** | **string** | Name of the Trusted Certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustedCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedCertificateResponse**](TrustedCertificateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrustedCertificates

> TrustedCertificateListResponse ListTrustedCertificates(ctx).Filter(filter).Execute()

Returns a list of all Trusted Certificate objects

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
    resp, r, err := apiClient.TrustedCertificateAPI.ListTrustedCertificates(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateAPI.ListTrustedCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTrustedCertificates`: TrustedCertificateListResponse
    fmt.Fprintf(os.Stdout, "Response from `TrustedCertificateAPI.ListTrustedCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrustedCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**TrustedCertificateListResponse**](TrustedCertificateListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrustedCertificate

> TrustedCertificateResponse UpdateTrustedCertificate(ctx, trustedCertificateName).UpdateRequest(updateRequest).Execute()

Update an existing Trusted Certificate by name

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
    trustedCertificateName := "trustedCertificateName_example" // string | Name of the Trusted Certificate
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Trusted Certificate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedCertificateAPI.UpdateTrustedCertificate(context.Background(), trustedCertificateName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateAPI.UpdateTrustedCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrustedCertificate`: TrustedCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TrustedCertificateAPI.UpdateTrustedCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedCertificateName** | **string** | Name of the Trusted Certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrustedCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Trusted Certificate | 

### Return type

[**TrustedCertificateResponse**](TrustedCertificateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

