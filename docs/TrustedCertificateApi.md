# \TrustedCertificateApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTrustedCertificate**](TrustedCertificateApi.md#AddTrustedCertificate) | **Post** /trusted-certificates | Add a new Trusted Certificate to the config
[**DeleteTrustedCertificate**](TrustedCertificateApi.md#DeleteTrustedCertificate) | **Delete** /trusted-certificates/{trusted-certificate-name} | Delete a Trusted Certificate
[**GetTrustedCertificate**](TrustedCertificateApi.md#GetTrustedCertificate) | **Get** /trusted-certificates/{trusted-certificate-name} | Returns a single Trusted Certificate
[**UpdateTrustedCertificate**](TrustedCertificateApi.md#UpdateTrustedCertificate) | **Patch** /trusted-certificates/{trusted-certificate-name} | Update an existing Trusted Certificate by name



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
    openapiclient "./openapi"
)

func main() {
    addTrustedCertificateRequest := *openapiclient.NewAddTrustedCertificateRequest("CertificateName_example", "Certificate_example") // AddTrustedCertificateRequest | Create a new Trusted Certificate in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedCertificateApi.AddTrustedCertificate(context.Background()).AddTrustedCertificateRequest(addTrustedCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateApi.AddTrustedCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTrustedCertificate`: TrustedCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TrustedCertificateApi.AddTrustedCertificate`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    trustedCertificateName := "trustedCertificateName_example" // string | Name of the Trusted Certificate to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedCertificateApi.DeleteTrustedCertificate(context.Background(), trustedCertificateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateApi.DeleteTrustedCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedCertificateName** | **string** | Name of the Trusted Certificate to be deleted | 

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
    openapiclient "./openapi"
)

func main() {
    trustedCertificateName := "trustedCertificateName_example" // string | Name of the Trusted Certificate to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedCertificateApi.GetTrustedCertificate(context.Background(), trustedCertificateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateApi.GetTrustedCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrustedCertificate`: TrustedCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TrustedCertificateApi.GetTrustedCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedCertificateName** | **string** | Name of the Trusted Certificate to be read | 

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
    openapiclient "./openapi"
)

func main() {
    trustedCertificateName := "trustedCertificateName_example" // string | Name of the Trusted Certificate to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Trusted Certificate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedCertificateApi.UpdateTrustedCertificate(context.Background(), trustedCertificateName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedCertificateApi.UpdateTrustedCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrustedCertificate`: TrustedCertificateResponse
    fmt.Fprintf(os.Stdout, "Response from `TrustedCertificateApi.UpdateTrustedCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trustedCertificateName** | **string** | Name of the Trusted Certificate to be updated | 

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

