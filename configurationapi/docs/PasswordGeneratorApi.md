# \PasswordGeneratorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPasswordGenerator**](PasswordGeneratorAPI.md#AddPasswordGenerator) | **Post** /password-generators | Add a new Password Generator to the config
[**DeletePasswordGenerator**](PasswordGeneratorAPI.md#DeletePasswordGenerator) | **Delete** /password-generators/{password-generator-name} | Delete a Password Generator
[**GetPasswordGenerator**](PasswordGeneratorAPI.md#GetPasswordGenerator) | **Get** /password-generators/{password-generator-name} | Returns a single Password Generator
[**ListPasswordGenerators**](PasswordGeneratorAPI.md#ListPasswordGenerators) | **Get** /password-generators | Returns a list of all Password Generator objects
[**UpdatePasswordGenerator**](PasswordGeneratorAPI.md#UpdatePasswordGenerator) | **Patch** /password-generators/{password-generator-name} | Update an existing Password Generator by name



## AddPasswordGenerator

> AddPasswordGenerator200Response AddPasswordGenerator(ctx).AddPasswordGeneratorRequest(addPasswordGeneratorRequest).Execute()

Add a new Password Generator to the config

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
    addPasswordGeneratorRequest := openapiclient.add_password_generator_request{AddGroovyScriptedPasswordGeneratorRequest: openapiclient.NewAddGroovyScriptedPasswordGeneratorRequest([]openapiclient.EnumgroovyScriptedPasswordGeneratorSchemaUrn{openapiclient.Enumgroovy-scripted-password-generatorSchemaUrn("urn:pingidentity:schemas:configuration:2.0:password-generator:groovy-scripted")}, "ScriptClass_example", false, "GeneratorName_example")} // AddPasswordGeneratorRequest | Create a new Password Generator in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordGeneratorAPI.AddPasswordGenerator(context.Background()).AddPasswordGeneratorRequest(addPasswordGeneratorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordGeneratorAPI.AddPasswordGenerator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPasswordGenerator`: AddPasswordGenerator200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordGeneratorAPI.AddPasswordGenerator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPasswordGeneratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPasswordGeneratorRequest** | [**AddPasswordGeneratorRequest**](AddPasswordGeneratorRequest.md) | Create a new Password Generator in the config | 

### Return type

[**AddPasswordGenerator200Response**](AddPasswordGenerator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePasswordGenerator

> DeletePasswordGenerator(ctx, passwordGeneratorName).Execute()

Delete a Password Generator

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
    passwordGeneratorName := "passwordGeneratorName_example" // string | Name of the Password Generator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PasswordGeneratorAPI.DeletePasswordGenerator(context.Background(), passwordGeneratorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordGeneratorAPI.DeletePasswordGenerator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordGeneratorName** | **string** | Name of the Password Generator | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePasswordGeneratorRequest struct via the builder pattern


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


## GetPasswordGenerator

> AddPasswordGenerator200Response GetPasswordGenerator(ctx, passwordGeneratorName).Execute()

Returns a single Password Generator

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
    passwordGeneratorName := "passwordGeneratorName_example" // string | Name of the Password Generator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordGeneratorAPI.GetPasswordGenerator(context.Background(), passwordGeneratorName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordGeneratorAPI.GetPasswordGenerator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordGenerator`: AddPasswordGenerator200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordGeneratorAPI.GetPasswordGenerator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordGeneratorName** | **string** | Name of the Password Generator | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordGeneratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddPasswordGenerator200Response**](AddPasswordGenerator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPasswordGenerators

> PasswordGeneratorListResponse ListPasswordGenerators(ctx).Filter(filter).Execute()

Returns a list of all Password Generator objects

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
    resp, r, err := apiClient.PasswordGeneratorAPI.ListPasswordGenerators(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordGeneratorAPI.ListPasswordGenerators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPasswordGenerators`: PasswordGeneratorListResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordGeneratorAPI.ListPasswordGenerators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPasswordGeneratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**PasswordGeneratorListResponse**](PasswordGeneratorListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordGenerator

> AddPasswordGenerator200Response UpdatePasswordGenerator(ctx, passwordGeneratorName).UpdateRequest(updateRequest).Execute()

Update an existing Password Generator by name

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
    passwordGeneratorName := "passwordGeneratorName_example" // string | Name of the Password Generator
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Password Generator

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordGeneratorAPI.UpdatePasswordGenerator(context.Background(), passwordGeneratorName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordGeneratorAPI.UpdatePasswordGenerator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordGenerator`: AddPasswordGenerator200Response
    fmt.Fprintf(os.Stdout, "Response from `PasswordGeneratorAPI.UpdatePasswordGenerator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordGeneratorName** | **string** | Name of the Password Generator | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordGeneratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Password Generator | 

### Return type

[**AddPasswordGenerator200Response**](AddPasswordGenerator200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

