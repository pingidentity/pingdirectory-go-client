# \ConsentDefinitionLocalizationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConsentDefinitionLocalization**](ConsentDefinitionLocalizationAPI.md#AddConsentDefinitionLocalization) | **Post** /consent-definitions/{consent-definition-name}/consent-definition-localizations | Add a new Consent Definition Localization to the config
[**DeleteConsentDefinitionLocalization**](ConsentDefinitionLocalizationAPI.md#DeleteConsentDefinitionLocalization) | **Delete** /consent-definitions/{consent-definition-name}/consent-definition-localizations/{consent-definition-localization-name} | Delete a Consent Definition Localization
[**GetConsentDefinitionLocalization**](ConsentDefinitionLocalizationAPI.md#GetConsentDefinitionLocalization) | **Get** /consent-definitions/{consent-definition-name}/consent-definition-localizations/{consent-definition-localization-name} | Returns a single Consent Definition Localization
[**ListConsentDefinitionLocalizations**](ConsentDefinitionLocalizationAPI.md#ListConsentDefinitionLocalizations) | **Get** /consent-definitions/{consent-definition-name}/consent-definition-localizations | Returns a list of all Consent Definition Localization objects
[**UpdateConsentDefinitionLocalization**](ConsentDefinitionLocalizationAPI.md#UpdateConsentDefinitionLocalization) | **Patch** /consent-definitions/{consent-definition-name}/consent-definition-localizations/{consent-definition-localization-name} | Update an existing Consent Definition Localization by name



## AddConsentDefinitionLocalization

> ConsentDefinitionLocalizationResponse AddConsentDefinitionLocalization(ctx, consentDefinitionName).AddConsentDefinitionLocalizationRequest(addConsentDefinitionLocalizationRequest).Execute()

Add a new Consent Definition Localization to the config

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
    consentDefinitionName := "consentDefinitionName_example" // string | Name of the Consent Definition
    addConsentDefinitionLocalizationRequest := *openapiclient.NewAddConsentDefinitionLocalizationRequest("Locale_example", "Version_example", "DataText_example", "PurposeText_example", "LocalizationName_example") // AddConsentDefinitionLocalizationRequest | Create a new Consent Definition Localization in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsentDefinitionLocalizationAPI.AddConsentDefinitionLocalization(context.Background(), consentDefinitionName).AddConsentDefinitionLocalizationRequest(addConsentDefinitionLocalizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentDefinitionLocalizationAPI.AddConsentDefinitionLocalization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConsentDefinitionLocalization`: ConsentDefinitionLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConsentDefinitionLocalizationAPI.AddConsentDefinitionLocalization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentDefinitionName** | **string** | Name of the Consent Definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddConsentDefinitionLocalizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addConsentDefinitionLocalizationRequest** | [**AddConsentDefinitionLocalizationRequest**](AddConsentDefinitionLocalizationRequest.md) | Create a new Consent Definition Localization in the config | 

### Return type

[**ConsentDefinitionLocalizationResponse**](ConsentDefinitionLocalizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConsentDefinitionLocalization

> DeleteConsentDefinitionLocalization(ctx, consentDefinitionLocalizationName, consentDefinitionName).Execute()

Delete a Consent Definition Localization

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
    consentDefinitionLocalizationName := "consentDefinitionLocalizationName_example" // string | Name of the Consent Definition Localization
    consentDefinitionName := "consentDefinitionName_example" // string | Name of the Consent Definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConsentDefinitionLocalizationAPI.DeleteConsentDefinitionLocalization(context.Background(), consentDefinitionLocalizationName, consentDefinitionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentDefinitionLocalizationAPI.DeleteConsentDefinitionLocalization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentDefinitionLocalizationName** | **string** | Name of the Consent Definition Localization | 
**consentDefinitionName** | **string** | Name of the Consent Definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsentDefinitionLocalizationRequest struct via the builder pattern


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


## GetConsentDefinitionLocalization

> ConsentDefinitionLocalizationResponse GetConsentDefinitionLocalization(ctx, consentDefinitionLocalizationName, consentDefinitionName).Execute()

Returns a single Consent Definition Localization

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
    consentDefinitionLocalizationName := "consentDefinitionLocalizationName_example" // string | Name of the Consent Definition Localization
    consentDefinitionName := "consentDefinitionName_example" // string | Name of the Consent Definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsentDefinitionLocalizationAPI.GetConsentDefinitionLocalization(context.Background(), consentDefinitionLocalizationName, consentDefinitionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentDefinitionLocalizationAPI.GetConsentDefinitionLocalization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentDefinitionLocalization`: ConsentDefinitionLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConsentDefinitionLocalizationAPI.GetConsentDefinitionLocalization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentDefinitionLocalizationName** | **string** | Name of the Consent Definition Localization | 
**consentDefinitionName** | **string** | Name of the Consent Definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentDefinitionLocalizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConsentDefinitionLocalizationResponse**](ConsentDefinitionLocalizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConsentDefinitionLocalizations

> ConsentDefinitionLocalizationListResponse ListConsentDefinitionLocalizations(ctx, consentDefinitionName).Filter(filter).Execute()

Returns a list of all Consent Definition Localization objects

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
    consentDefinitionName := "consentDefinitionName_example" // string | Name of the Consent Definition
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsentDefinitionLocalizationAPI.ListConsentDefinitionLocalizations(context.Background(), consentDefinitionName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentDefinitionLocalizationAPI.ListConsentDefinitionLocalizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConsentDefinitionLocalizations`: ConsentDefinitionLocalizationListResponse
    fmt.Fprintf(os.Stdout, "Response from `ConsentDefinitionLocalizationAPI.ListConsentDefinitionLocalizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentDefinitionName** | **string** | Name of the Consent Definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConsentDefinitionLocalizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | SCIM filter | 

### Return type

[**ConsentDefinitionLocalizationListResponse**](ConsentDefinitionLocalizationListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConsentDefinitionLocalization

> ConsentDefinitionLocalizationResponse UpdateConsentDefinitionLocalization(ctx, consentDefinitionLocalizationName, consentDefinitionName).UpdateRequest(updateRequest).Execute()

Update an existing Consent Definition Localization by name

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
    consentDefinitionLocalizationName := "consentDefinitionLocalizationName_example" // string | Name of the Consent Definition Localization
    consentDefinitionName := "consentDefinitionName_example" // string | Name of the Consent Definition
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Consent Definition Localization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsentDefinitionLocalizationAPI.UpdateConsentDefinitionLocalization(context.Background(), consentDefinitionLocalizationName, consentDefinitionName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentDefinitionLocalizationAPI.UpdateConsentDefinitionLocalization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConsentDefinitionLocalization`: ConsentDefinitionLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConsentDefinitionLocalizationAPI.UpdateConsentDefinitionLocalization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentDefinitionLocalizationName** | **string** | Name of the Consent Definition Localization | 
**consentDefinitionName** | **string** | Name of the Consent Definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConsentDefinitionLocalizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Consent Definition Localization | 

### Return type

[**ConsentDefinitionLocalizationResponse**](ConsentDefinitionLocalizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

