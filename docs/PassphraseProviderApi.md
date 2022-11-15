# \PassphraseProviderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPassphraseProvider**](PassphraseProviderApi.md#AddPassphraseProvider) | **Post** /passphrase-providers | Add a new Passphrase Provider to the config
[**DeletePassphraseProvider**](PassphraseProviderApi.md#DeletePassphraseProvider) | **Delete** /passphrase-providers/{passphrase-provider-name} | Delete a Passphrase Provider
[**GetPassphraseProvider**](PassphraseProviderApi.md#GetPassphraseProvider) | **Get** /passphrase-providers/{passphrase-provider-name} | Returns a single Passphrase Provider
[**UpdatePassphraseProvider**](PassphraseProviderApi.md#UpdatePassphraseProvider) | **Patch** /passphrase-providers/{passphrase-provider-name} | Update an existing Passphrase Provider by name



## AddPassphraseProvider

> AddPassphraseProvider200Response AddPassphraseProvider(ctx).AddPassphraseProviderRequest(addPassphraseProviderRequest).Execute()

Add a new Passphrase Provider to the config

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
    addPassphraseProviderRequest := openapiclient.add_passphrase_provider_request{AddAmazonSecretsManagerPassphraseProviderRequest: openapiclient.NewAddAmazonSecretsManagerPassphraseProviderRequest("ProviderName_example", []openapiclient.EnumamazonSecretsManagerPassphraseProviderSchemaUrn{openapiclient.Enumamazon-secrets-manager-passphrase-providerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:passphrase-provider:amazon-secrets-manager")}, "AwsExternalServer_example", "SecretID_example", "SecretFieldName_example", false)} // AddPassphraseProviderRequest | Create a new Passphrase Provider in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PassphraseProviderApi.AddPassphraseProvider(context.Background()).AddPassphraseProviderRequest(addPassphraseProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassphraseProviderApi.AddPassphraseProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPassphraseProvider`: AddPassphraseProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `PassphraseProviderApi.AddPassphraseProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPassphraseProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPassphraseProviderRequest** | [**AddPassphraseProviderRequest**](AddPassphraseProviderRequest.md) | Create a new Passphrase Provider in the config | 

### Return type

[**AddPassphraseProvider200Response**](AddPassphraseProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePassphraseProvider

> DeletePassphraseProvider(ctx, passphraseProviderName).Execute()

Delete a Passphrase Provider

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
    passphraseProviderName := "passphraseProviderName_example" // string | Name of the Passphrase Provider to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PassphraseProviderApi.DeletePassphraseProvider(context.Background(), passphraseProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassphraseProviderApi.DeletePassphraseProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passphraseProviderName** | **string** | Name of the Passphrase Provider to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePassphraseProviderRequest struct via the builder pattern


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


## GetPassphraseProvider

> AddPassphraseProvider200Response GetPassphraseProvider(ctx, passphraseProviderName).Execute()

Returns a single Passphrase Provider

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
    passphraseProviderName := "passphraseProviderName_example" // string | Name of the Passphrase Provider to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PassphraseProviderApi.GetPassphraseProvider(context.Background(), passphraseProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassphraseProviderApi.GetPassphraseProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPassphraseProvider`: AddPassphraseProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `PassphraseProviderApi.GetPassphraseProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passphraseProviderName** | **string** | Name of the Passphrase Provider to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPassphraseProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddPassphraseProvider200Response**](AddPassphraseProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePassphraseProvider

> AddPassphraseProvider200Response UpdatePassphraseProvider(ctx, passphraseProviderName).UpdateRequest(updateRequest).Execute()

Update an existing Passphrase Provider by name

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
    passphraseProviderName := "passphraseProviderName_example" // string | Name of the Passphrase Provider to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Passphrase Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PassphraseProviderApi.UpdatePassphraseProvider(context.Background(), passphraseProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassphraseProviderApi.UpdatePassphraseProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePassphraseProvider`: AddPassphraseProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `PassphraseProviderApi.UpdatePassphraseProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passphraseProviderName** | **string** | Name of the Passphrase Provider to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePassphraseProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Passphrase Provider | 

### Return type

[**AddPassphraseProvider200Response**](AddPassphraseProvider200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

