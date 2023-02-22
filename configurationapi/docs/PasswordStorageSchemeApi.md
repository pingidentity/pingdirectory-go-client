# \PasswordStorageSchemeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPasswordStorageScheme**](PasswordStorageSchemeApi.md#AddPasswordStorageScheme) | **Post** /password-storage-schemes | Add a new Password Storage Scheme to the config
[**DeletePasswordStorageScheme**](PasswordStorageSchemeApi.md#DeletePasswordStorageScheme) | **Delete** /password-storage-schemes/{password-storage-scheme-name} | Delete a Password Storage Scheme
[**GetPasswordStorageScheme**](PasswordStorageSchemeApi.md#GetPasswordStorageScheme) | **Get** /password-storage-schemes/{password-storage-scheme-name} | Returns a single Password Storage Scheme
[**UpdatePasswordStorageScheme**](PasswordStorageSchemeApi.md#UpdatePasswordStorageScheme) | **Patch** /password-storage-schemes/{password-storage-scheme-name} | Update an existing Password Storage Scheme by name



## AddPasswordStorageScheme

> AddPasswordStorageScheme200Response AddPasswordStorageScheme(ctx).AddPasswordStorageSchemeRequest(addPasswordStorageSchemeRequest).Execute()

Add a new Password Storage Scheme to the config

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
    addPasswordStorageSchemeRequest := openapiclient.add_password_storage_scheme_request{AddAes256PasswordStorageSchemeRequest: openapiclient.NewAddAes256PasswordStorageSchemeRequest("SchemeName_example", []openapiclient.Enumaes256PasswordStorageSchemeSchemaUrn{openapiclient.Enumaes-256-password-storage-schemeSchemaUrn("urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:aes-256")}, false)} // AddPasswordStorageSchemeRequest | Create a new Password Storage Scheme in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordStorageSchemeApi.AddPasswordStorageScheme(context.Background()).AddPasswordStorageSchemeRequest(addPasswordStorageSchemeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordStorageSchemeApi.AddPasswordStorageScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPasswordStorageScheme`: AddPasswordStorageScheme200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordStorageSchemeApi.AddPasswordStorageScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPasswordStorageSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPasswordStorageSchemeRequest** | [**AddPasswordStorageSchemeRequest**](AddPasswordStorageSchemeRequest.md) | Create a new Password Storage Scheme in the config | 

### Return type

[**AddPasswordStorageScheme200Response**](AddPasswordStorageScheme200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePasswordStorageScheme

> DeletePasswordStorageScheme(ctx, passwordStorageSchemeName).Execute()

Delete a Password Storage Scheme

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
    passwordStorageSchemeName := "passwordStorageSchemeName_example" // string | Name of the Password Storage Scheme

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordStorageSchemeApi.DeletePasswordStorageScheme(context.Background(), passwordStorageSchemeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordStorageSchemeApi.DeletePasswordStorageScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordStorageSchemeName** | **string** | Name of the Password Storage Scheme | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePasswordStorageSchemeRequest struct via the builder pattern


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


## GetPasswordStorageScheme

> GetPasswordStorageScheme200Response GetPasswordStorageScheme(ctx, passwordStorageSchemeName).Execute()

Returns a single Password Storage Scheme

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
    passwordStorageSchemeName := "passwordStorageSchemeName_example" // string | Name of the Password Storage Scheme

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordStorageSchemeApi.GetPasswordStorageScheme(context.Background(), passwordStorageSchemeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordStorageSchemeApi.GetPasswordStorageScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordStorageScheme`: GetPasswordStorageScheme200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordStorageSchemeApi.GetPasswordStorageScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordStorageSchemeName** | **string** | Name of the Password Storage Scheme | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordStorageSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPasswordStorageScheme200Response**](GetPasswordStorageScheme200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordStorageScheme

> GetPasswordStorageScheme200Response UpdatePasswordStorageScheme(ctx, passwordStorageSchemeName).UpdateRequest(updateRequest).Execute()

Update an existing Password Storage Scheme by name

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
    passwordStorageSchemeName := "passwordStorageSchemeName_example" // string | Name of the Password Storage Scheme
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Password Storage Scheme

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordStorageSchemeApi.UpdatePasswordStorageScheme(context.Background(), passwordStorageSchemeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordStorageSchemeApi.UpdatePasswordStorageScheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordStorageScheme`: GetPasswordStorageScheme200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordStorageSchemeApi.UpdatePasswordStorageScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordStorageSchemeName** | **string** | Name of the Password Storage Scheme | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordStorageSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Password Storage Scheme | 

### Return type

[**GetPasswordStorageScheme200Response**](GetPasswordStorageScheme200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

