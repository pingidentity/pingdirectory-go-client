# \PassphraseProviderAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPassphraseProvider**](PassphraseProviderAPI.md#AddPassphraseProvider) | **Post** /passphrase-providers | Add a new Passphrase Provider to the config
[**DeletePassphraseProvider**](PassphraseProviderAPI.md#DeletePassphraseProvider) | **Delete** /passphrase-providers/{passphrase-provider-name} | Delete a Passphrase Provider
[**GetPassphraseProvider**](PassphraseProviderAPI.md#GetPassphraseProvider) | **Get** /passphrase-providers/{passphrase-provider-name} | Returns a single Passphrase Provider
[**ListPassphraseProviders**](PassphraseProviderAPI.md#ListPassphraseProviders) | **Get** /passphrase-providers | Returns a list of all Passphrase Provider objects
[**UpdatePassphraseProvider**](PassphraseProviderAPI.md#UpdatePassphraseProvider) | **Patch** /passphrase-providers/{passphrase-provider-name} | Update an existing Passphrase Provider by name



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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    addPassphraseProviderRequest := openapiclient.add_passphrase_provider_request{AddAmazonSecretsManagerPassphraseProviderRequest: openapiclient.NewAddAmazonSecretsManagerPassphraseProviderRequest([]openapiclient.EnumamazonSecretsManagerPassphraseProviderSchemaUrn{openapiclient.Enumamazon-secrets-manager-passphrase-providerSchemaUrn("urn:pingidentity:schemas:configuration:2.0:passphrase-provider:amazon-secrets-manager")}, "AwsExternalServer_example", "SecretID_example", "SecretFieldName_example", false, "ProviderName_example")} // AddPassphraseProviderRequest | Create a new Passphrase Provider in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PassphraseProviderAPI.AddPassphraseProvider(context.Background()).AddPassphraseProviderRequest(addPassphraseProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassphraseProviderAPI.AddPassphraseProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPassphraseProvider`: AddPassphraseProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `PassphraseProviderAPI.AddPassphraseProvider`: %v\n", resp)
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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    passphraseProviderName := "passphraseProviderName_example" // string | Name of the Passphrase Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PassphraseProviderAPI.DeletePassphraseProvider(context.Background(), passphraseProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassphraseProviderAPI.DeletePassphraseProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passphraseProviderName** | **string** | Name of the Passphrase Provider | 

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    passphraseProviderName := "passphraseProviderName_example" // string | Name of the Passphrase Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PassphraseProviderAPI.GetPassphraseProvider(context.Background(), passphraseProviderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassphraseProviderAPI.GetPassphraseProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPassphraseProvider`: AddPassphraseProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `PassphraseProviderAPI.GetPassphraseProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passphraseProviderName** | **string** | Name of the Passphrase Provider | 

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


## ListPassphraseProviders

> PassphraseProviderListResponse ListPassphraseProviders(ctx).Filter(filter).Execute()

Returns a list of all Passphrase Provider objects

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
    resp, r, err := apiClient.PassphraseProviderAPI.ListPassphraseProviders(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassphraseProviderAPI.ListPassphraseProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPassphraseProviders`: PassphraseProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `PassphraseProviderAPI.ListPassphraseProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPassphraseProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**PassphraseProviderListResponse**](PassphraseProviderListResponse.md)

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
    openapiclient "github.com/pingidentity/pingdirectory-go-client"
)

func main() {
    passphraseProviderName := "passphraseProviderName_example" // string | Name of the Passphrase Provider
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Passphrase Provider

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PassphraseProviderAPI.UpdatePassphraseProvider(context.Background(), passphraseProviderName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PassphraseProviderAPI.UpdatePassphraseProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePassphraseProvider`: AddPassphraseProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `PassphraseProviderAPI.UpdatePassphraseProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passphraseProviderName** | **string** | Name of the Passphrase Provider | 

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

