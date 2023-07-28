# \ConsentDefinitionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConsentDefinition**](ConsentDefinitionApi.md#AddConsentDefinition) | **Post** /consent-definitions | Add a new Consent Definition to the config
[**DeleteConsentDefinition**](ConsentDefinitionApi.md#DeleteConsentDefinition) | **Delete** /consent-definitions/{consent-definition-name} | Delete a Consent Definition
[**GetConsentDefinition**](ConsentDefinitionApi.md#GetConsentDefinition) | **Get** /consent-definitions/{consent-definition-name} | Returns a single Consent Definition
[**ListConsentDefinitions**](ConsentDefinitionApi.md#ListConsentDefinitions) | **Get** /consent-definitions | Returns a list of all Consent Definition objects
[**UpdateConsentDefinition**](ConsentDefinitionApi.md#UpdateConsentDefinition) | **Patch** /consent-definitions/{consent-definition-name} | Update an existing Consent Definition by name



## AddConsentDefinition

> ConsentDefinitionResponse AddConsentDefinition(ctx).AddConsentDefinitionRequest(addConsentDefinitionRequest).Execute()

Add a new Consent Definition to the config

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
    addConsentDefinitionRequest := *openapiclient.NewAddConsentDefinitionRequest("DefinitionName_example", "UniqueID_example") // AddConsentDefinitionRequest | Create a new Consent Definition in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsentDefinitionApi.AddConsentDefinition(context.Background()).AddConsentDefinitionRequest(addConsentDefinitionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentDefinitionApi.AddConsentDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConsentDefinition`: ConsentDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConsentDefinitionApi.AddConsentDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddConsentDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addConsentDefinitionRequest** | [**AddConsentDefinitionRequest**](AddConsentDefinitionRequest.md) | Create a new Consent Definition in the config | 

### Return type

[**ConsentDefinitionResponse**](ConsentDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConsentDefinition

> DeleteConsentDefinition(ctx, consentDefinitionName).Execute()

Delete a Consent Definition

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConsentDefinitionApi.DeleteConsentDefinition(context.Background(), consentDefinitionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentDefinitionApi.DeleteConsentDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentDefinitionName** | **string** | Name of the Consent Definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConsentDefinitionRequest struct via the builder pattern


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


## GetConsentDefinition

> ConsentDefinitionResponse GetConsentDefinition(ctx, consentDefinitionName).Execute()

Returns a single Consent Definition

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsentDefinitionApi.GetConsentDefinition(context.Background(), consentDefinitionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentDefinitionApi.GetConsentDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsentDefinition`: ConsentDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConsentDefinitionApi.GetConsentDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentDefinitionName** | **string** | Name of the Consent Definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsentDefinitionResponse**](ConsentDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConsentDefinitions

> ConsentDefinitionListResponse ListConsentDefinitions(ctx).Filter(filter).Execute()

Returns a list of all Consent Definition objects

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
    resp, r, err := apiClient.ConsentDefinitionApi.ListConsentDefinitions(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentDefinitionApi.ListConsentDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConsentDefinitions`: ConsentDefinitionListResponse
    fmt.Fprintf(os.Stdout, "Response from `ConsentDefinitionApi.ListConsentDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConsentDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ConsentDefinitionListResponse**](ConsentDefinitionListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConsentDefinition

> ConsentDefinitionResponse UpdateConsentDefinition(ctx, consentDefinitionName).UpdateRequest(updateRequest).Execute()

Update an existing Consent Definition by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Consent Definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConsentDefinitionApi.UpdateConsentDefinition(context.Background(), consentDefinitionName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConsentDefinitionApi.UpdateConsentDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConsentDefinition`: ConsentDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConsentDefinitionApi.UpdateConsentDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consentDefinitionName** | **string** | Name of the Consent Definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConsentDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Consent Definition | 

### Return type

[**ConsentDefinitionResponse**](ConsentDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

