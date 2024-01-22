# \AzureAuthenticationMethodAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAzureAuthenticationMethod**](AzureAuthenticationMethodAPI.md#AddAzureAuthenticationMethod) | **Post** /azure-authentication-methods | Add a new Azure Authentication Method to the config
[**DeleteAzureAuthenticationMethod**](AzureAuthenticationMethodAPI.md#DeleteAzureAuthenticationMethod) | **Delete** /azure-authentication-methods/{azure-authentication-method-name} | Delete a Azure Authentication Method
[**GetAzureAuthenticationMethod**](AzureAuthenticationMethodAPI.md#GetAzureAuthenticationMethod) | **Get** /azure-authentication-methods/{azure-authentication-method-name} | Returns a single Azure Authentication Method
[**ListAzureAuthenticationMethods**](AzureAuthenticationMethodAPI.md#ListAzureAuthenticationMethods) | **Get** /azure-authentication-methods | Returns a list of all Azure Authentication Method objects
[**UpdateAzureAuthenticationMethod**](AzureAuthenticationMethodAPI.md#UpdateAzureAuthenticationMethod) | **Patch** /azure-authentication-methods/{azure-authentication-method-name} | Update an existing Azure Authentication Method by name



## AddAzureAuthenticationMethod

> AddAzureAuthenticationMethod200Response AddAzureAuthenticationMethod(ctx).AddAzureAuthenticationMethodRequest(addAzureAuthenticationMethodRequest).Execute()

Add a new Azure Authentication Method to the config

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
    addAzureAuthenticationMethodRequest := openapiclient.add_azure_authentication_method_request{AddClientSecretAzureAuthenticationMethodRequest: openapiclient.NewAddClientSecretAzureAuthenticationMethodRequest([]openapiclient.EnumclientSecretAzureAuthenticationMethodSchemaUrn{openapiclient.Enumclient-secret-azure-authentication-methodSchemaUrn("urn:pingidentity:schemas:configuration:2.0:azure-authentication-method:client-secret")}, "TenantID_example", "ClientID_example", "ClientSecret_example", "MethodName_example")} // AddAzureAuthenticationMethodRequest | Create a new Azure Authentication Method in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureAuthenticationMethodAPI.AddAzureAuthenticationMethod(context.Background()).AddAzureAuthenticationMethodRequest(addAzureAuthenticationMethodRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureAuthenticationMethodAPI.AddAzureAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAzureAuthenticationMethod`: AddAzureAuthenticationMethod200Response
    fmt.Fprintf(os.Stdout, "Response from `AzureAuthenticationMethodAPI.AddAzureAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAzureAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addAzureAuthenticationMethodRequest** | [**AddAzureAuthenticationMethodRequest**](AddAzureAuthenticationMethodRequest.md) | Create a new Azure Authentication Method in the config | 

### Return type

[**AddAzureAuthenticationMethod200Response**](AddAzureAuthenticationMethod200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAzureAuthenticationMethod

> DeleteAzureAuthenticationMethod(ctx, azureAuthenticationMethodName).Execute()

Delete a Azure Authentication Method

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
    azureAuthenticationMethodName := "azureAuthenticationMethodName_example" // string | Name of the Azure Authentication Method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AzureAuthenticationMethodAPI.DeleteAzureAuthenticationMethod(context.Background(), azureAuthenticationMethodName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureAuthenticationMethodAPI.DeleteAzureAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAuthenticationMethodName** | **string** | Name of the Azure Authentication Method | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureAuthenticationMethodRequest struct via the builder pattern


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


## GetAzureAuthenticationMethod

> AddAzureAuthenticationMethod200Response GetAzureAuthenticationMethod(ctx, azureAuthenticationMethodName).Execute()

Returns a single Azure Authentication Method

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
    azureAuthenticationMethodName := "azureAuthenticationMethodName_example" // string | Name of the Azure Authentication Method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureAuthenticationMethodAPI.GetAzureAuthenticationMethod(context.Background(), azureAuthenticationMethodName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureAuthenticationMethodAPI.GetAzureAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAzureAuthenticationMethod`: AddAzureAuthenticationMethod200Response
    fmt.Fprintf(os.Stdout, "Response from `AzureAuthenticationMethodAPI.GetAzureAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAuthenticationMethodName** | **string** | Name of the Azure Authentication Method | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddAzureAuthenticationMethod200Response**](AddAzureAuthenticationMethod200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAzureAuthenticationMethods

> AzureAuthenticationMethodListResponse ListAzureAuthenticationMethods(ctx).Filter(filter).Execute()

Returns a list of all Azure Authentication Method objects

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
    resp, r, err := apiClient.AzureAuthenticationMethodAPI.ListAzureAuthenticationMethods(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureAuthenticationMethodAPI.ListAzureAuthenticationMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAzureAuthenticationMethods`: AzureAuthenticationMethodListResponse
    fmt.Fprintf(os.Stdout, "Response from `AzureAuthenticationMethodAPI.ListAzureAuthenticationMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAzureAuthenticationMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**AzureAuthenticationMethodListResponse**](AzureAuthenticationMethodListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAzureAuthenticationMethod

> AddAzureAuthenticationMethod200Response UpdateAzureAuthenticationMethod(ctx, azureAuthenticationMethodName).UpdateRequest(updateRequest).Execute()

Update an existing Azure Authentication Method by name

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
    azureAuthenticationMethodName := "azureAuthenticationMethodName_example" // string | Name of the Azure Authentication Method
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Azure Authentication Method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureAuthenticationMethodAPI.UpdateAzureAuthenticationMethod(context.Background(), azureAuthenticationMethodName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureAuthenticationMethodAPI.UpdateAzureAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAzureAuthenticationMethod`: AddAzureAuthenticationMethod200Response
    fmt.Fprintf(os.Stdout, "Response from `AzureAuthenticationMethodAPI.UpdateAzureAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**azureAuthenticationMethodName** | **string** | Name of the Azure Authentication Method | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAzureAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Azure Authentication Method | 

### Return type

[**AddAzureAuthenticationMethod200Response**](AddAzureAuthenticationMethod200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

