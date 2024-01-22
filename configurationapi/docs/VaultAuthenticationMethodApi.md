# \VaultAuthenticationMethodAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVaultAuthenticationMethod**](VaultAuthenticationMethodAPI.md#AddVaultAuthenticationMethod) | **Post** /vault-authentication-methods | Add a new Vault Authentication Method to the config
[**DeleteVaultAuthenticationMethod**](VaultAuthenticationMethodAPI.md#DeleteVaultAuthenticationMethod) | **Delete** /vault-authentication-methods/{vault-authentication-method-name} | Delete a Vault Authentication Method
[**GetVaultAuthenticationMethod**](VaultAuthenticationMethodAPI.md#GetVaultAuthenticationMethod) | **Get** /vault-authentication-methods/{vault-authentication-method-name} | Returns a single Vault Authentication Method
[**ListVaultAuthenticationMethods**](VaultAuthenticationMethodAPI.md#ListVaultAuthenticationMethods) | **Get** /vault-authentication-methods | Returns a list of all Vault Authentication Method objects
[**UpdateVaultAuthenticationMethod**](VaultAuthenticationMethodAPI.md#UpdateVaultAuthenticationMethod) | **Patch** /vault-authentication-methods/{vault-authentication-method-name} | Update an existing Vault Authentication Method by name



## AddVaultAuthenticationMethod

> AddVaultAuthenticationMethod200Response AddVaultAuthenticationMethod(ctx).AddVaultAuthenticationMethodRequest(addVaultAuthenticationMethodRequest).Execute()

Add a new Vault Authentication Method to the config

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
    addVaultAuthenticationMethodRequest := openapiclient.add_vault_authentication_method_request{AddAppRoleVaultAuthenticationMethodRequest: openapiclient.NewAddAppRoleVaultAuthenticationMethodRequest([]openapiclient.EnumappRoleVaultAuthenticationMethodSchemaUrn{openapiclient.Enumapp-role-vault-authentication-methodSchemaUrn("urn:pingidentity:schemas:configuration:2.0:vault-authentication-method:app-role")}, "VaultRoleID_example", "VaultSecretID_example", "MethodName_example")} // AddVaultAuthenticationMethodRequest | Create a new Vault Authentication Method in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultAuthenticationMethodAPI.AddVaultAuthenticationMethod(context.Background()).AddVaultAuthenticationMethodRequest(addVaultAuthenticationMethodRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultAuthenticationMethodAPI.AddVaultAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVaultAuthenticationMethod`: AddVaultAuthenticationMethod200Response
    fmt.Fprintf(os.Stdout, "Response from `VaultAuthenticationMethodAPI.AddVaultAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVaultAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addVaultAuthenticationMethodRequest** | [**AddVaultAuthenticationMethodRequest**](AddVaultAuthenticationMethodRequest.md) | Create a new Vault Authentication Method in the config | 

### Return type

[**AddVaultAuthenticationMethod200Response**](AddVaultAuthenticationMethod200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVaultAuthenticationMethod

> DeleteVaultAuthenticationMethod(ctx, vaultAuthenticationMethodName).Execute()

Delete a Vault Authentication Method

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
    vaultAuthenticationMethodName := "vaultAuthenticationMethodName_example" // string | Name of the Vault Authentication Method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VaultAuthenticationMethodAPI.DeleteVaultAuthenticationMethod(context.Background(), vaultAuthenticationMethodName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultAuthenticationMethodAPI.DeleteVaultAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultAuthenticationMethodName** | **string** | Name of the Vault Authentication Method | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultAuthenticationMethodRequest struct via the builder pattern


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


## GetVaultAuthenticationMethod

> AddVaultAuthenticationMethod200Response GetVaultAuthenticationMethod(ctx, vaultAuthenticationMethodName).Execute()

Returns a single Vault Authentication Method

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
    vaultAuthenticationMethodName := "vaultAuthenticationMethodName_example" // string | Name of the Vault Authentication Method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultAuthenticationMethodAPI.GetVaultAuthenticationMethod(context.Background(), vaultAuthenticationMethodName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultAuthenticationMethodAPI.GetVaultAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultAuthenticationMethod`: AddVaultAuthenticationMethod200Response
    fmt.Fprintf(os.Stdout, "Response from `VaultAuthenticationMethodAPI.GetVaultAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultAuthenticationMethodName** | **string** | Name of the Vault Authentication Method | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddVaultAuthenticationMethod200Response**](AddVaultAuthenticationMethod200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVaultAuthenticationMethods

> VaultAuthenticationMethodListResponse ListVaultAuthenticationMethods(ctx).Filter(filter).Execute()

Returns a list of all Vault Authentication Method objects

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
    resp, r, err := apiClient.VaultAuthenticationMethodAPI.ListVaultAuthenticationMethods(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultAuthenticationMethodAPI.ListVaultAuthenticationMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultAuthenticationMethods`: VaultAuthenticationMethodListResponse
    fmt.Fprintf(os.Stdout, "Response from `VaultAuthenticationMethodAPI.ListVaultAuthenticationMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVaultAuthenticationMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**VaultAuthenticationMethodListResponse**](VaultAuthenticationMethodListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVaultAuthenticationMethod

> AddVaultAuthenticationMethod200Response UpdateVaultAuthenticationMethod(ctx, vaultAuthenticationMethodName).UpdateRequest(updateRequest).Execute()

Update an existing Vault Authentication Method by name

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
    vaultAuthenticationMethodName := "vaultAuthenticationMethodName_example" // string | Name of the Vault Authentication Method
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Vault Authentication Method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VaultAuthenticationMethodAPI.UpdateVaultAuthenticationMethod(context.Background(), vaultAuthenticationMethodName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultAuthenticationMethodAPI.UpdateVaultAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVaultAuthenticationMethod`: AddVaultAuthenticationMethod200Response
    fmt.Fprintf(os.Stdout, "Response from `VaultAuthenticationMethodAPI.UpdateVaultAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vaultAuthenticationMethodName** | **string** | Name of the Vault Authentication Method | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Vault Authentication Method | 

### Return type

[**AddVaultAuthenticationMethod200Response**](AddVaultAuthenticationMethod200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

