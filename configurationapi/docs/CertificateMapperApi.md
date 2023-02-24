# \CertificateMapperApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCertificateMapper**](CertificateMapperApi.md#AddCertificateMapper) | **Post** /certificate-mappers | Add a new Certificate Mapper to the config
[**DeleteCertificateMapper**](CertificateMapperApi.md#DeleteCertificateMapper) | **Delete** /certificate-mappers/{certificate-mapper-name} | Delete a Certificate Mapper
[**GetCertificateMapper**](CertificateMapperApi.md#GetCertificateMapper) | **Get** /certificate-mappers/{certificate-mapper-name} | Returns a single Certificate Mapper
[**UpdateCertificateMapper**](CertificateMapperApi.md#UpdateCertificateMapper) | **Patch** /certificate-mappers/{certificate-mapper-name} | Update an existing Certificate Mapper by name



## AddCertificateMapper

> AddCertificateMapper200Response AddCertificateMapper(ctx).AddCertificateMapperRequest(addCertificateMapperRequest).Execute()

Add a new Certificate Mapper to the config

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
    addCertificateMapperRequest := openapiclient.add_certificate_mapper_request{AddFingerprintCertificateMapperRequest: openapiclient.NewAddFingerprintCertificateMapperRequest("MapperName_example", []openapiclient.EnumfingerprintCertificateMapperSchemaUrn{openapiclient.Enumfingerprint-certificate-mapperSchemaUrn("urn:pingidentity:schemas:configuration:2.0:certificate-mapper:fingerprint")}, openapiclient.Enumcertificate-mapper-fingerprintAlgorithmProp("md5"), false)} // AddCertificateMapperRequest | Create a new Certificate Mapper in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateMapperApi.AddCertificateMapper(context.Background()).AddCertificateMapperRequest(addCertificateMapperRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateMapperApi.AddCertificateMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCertificateMapper`: AddCertificateMapper200Response
    fmt.Fprintf(os.Stdout, "Response from `CertificateMapperApi.AddCertificateMapper`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCertificateMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addCertificateMapperRequest** | [**AddCertificateMapperRequest**](AddCertificateMapperRequest.md) | Create a new Certificate Mapper in the config | 

### Return type

[**AddCertificateMapper200Response**](AddCertificateMapper200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificateMapper

> DeleteCertificateMapper(ctx, certificateMapperName).Execute()

Delete a Certificate Mapper

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
    certificateMapperName := "certificateMapperName_example" // string | Name of the Certificate Mapper

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateMapperApi.DeleteCertificateMapper(context.Background(), certificateMapperName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateMapperApi.DeleteCertificateMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateMapperName** | **string** | Name of the Certificate Mapper | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateMapperRequest struct via the builder pattern


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


## GetCertificateMapper

> AddCertificateMapper200Response GetCertificateMapper(ctx, certificateMapperName).Execute()

Returns a single Certificate Mapper

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
    certificateMapperName := "certificateMapperName_example" // string | Name of the Certificate Mapper

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateMapperApi.GetCertificateMapper(context.Background(), certificateMapperName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateMapperApi.GetCertificateMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateMapper`: AddCertificateMapper200Response
    fmt.Fprintf(os.Stdout, "Response from `CertificateMapperApi.GetCertificateMapper`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateMapperName** | **string** | Name of the Certificate Mapper | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddCertificateMapper200Response**](AddCertificateMapper200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateMapper

> AddCertificateMapper200Response UpdateCertificateMapper(ctx, certificateMapperName).UpdateRequest(updateRequest).Execute()

Update an existing Certificate Mapper by name

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
    certificateMapperName := "certificateMapperName_example" // string | Name of the Certificate Mapper
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Certificate Mapper

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateMapperApi.UpdateCertificateMapper(context.Background(), certificateMapperName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateMapperApi.UpdateCertificateMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificateMapper`: AddCertificateMapper200Response
    fmt.Fprintf(os.Stdout, "Response from `CertificateMapperApi.UpdateCertificateMapper`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateMapperName** | **string** | Name of the Certificate Mapper | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Certificate Mapper | 

### Return type

[**AddCertificateMapper200Response**](AddCertificateMapper200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

