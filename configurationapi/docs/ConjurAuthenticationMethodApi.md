# \ConjurAuthenticationMethodAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConjurAuthenticationMethod**](ConjurAuthenticationMethodAPI.md#AddConjurAuthenticationMethod) | **Post** /conjur-authentication-methods | Add a new Conjur Authentication Method to the config
[**DeleteConjurAuthenticationMethod**](ConjurAuthenticationMethodAPI.md#DeleteConjurAuthenticationMethod) | **Delete** /conjur-authentication-methods/{conjur-authentication-method-name} | Delete a Conjur Authentication Method
[**GetConjurAuthenticationMethod**](ConjurAuthenticationMethodAPI.md#GetConjurAuthenticationMethod) | **Get** /conjur-authentication-methods/{conjur-authentication-method-name} | Returns a single Conjur Authentication Method
[**ListConjurAuthenticationMethods**](ConjurAuthenticationMethodAPI.md#ListConjurAuthenticationMethods) | **Get** /conjur-authentication-methods | Returns a list of all Conjur Authentication Method objects
[**UpdateConjurAuthenticationMethod**](ConjurAuthenticationMethodAPI.md#UpdateConjurAuthenticationMethod) | **Patch** /conjur-authentication-methods/{conjur-authentication-method-name} | Update an existing Conjur Authentication Method by name



## AddConjurAuthenticationMethod

> ApiKeyConjurAuthenticationMethodResponse AddConjurAuthenticationMethod(ctx).AddApiKeyConjurAuthenticationMethodRequest(addApiKeyConjurAuthenticationMethodRequest).Execute()

Add a new Conjur Authentication Method to the config

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
    addApiKeyConjurAuthenticationMethodRequest := *openapiclient.NewAddApiKeyConjurAuthenticationMethodRequest([]openapiclient.EnumapiKeyConjurAuthenticationMethodSchemaUrn{openapiclient.Enumapi-key-conjur-authentication-methodSchemaUrn("urn:pingidentity:schemas:configuration:2.0:conjur-authentication-method:api-key")}, "Username_example", "MethodName_example") // AddApiKeyConjurAuthenticationMethodRequest | Create a new Conjur Authentication Method in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConjurAuthenticationMethodAPI.AddConjurAuthenticationMethod(context.Background()).AddApiKeyConjurAuthenticationMethodRequest(addApiKeyConjurAuthenticationMethodRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConjurAuthenticationMethodAPI.AddConjurAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConjurAuthenticationMethod`: ApiKeyConjurAuthenticationMethodResponse
    fmt.Fprintf(os.Stdout, "Response from `ConjurAuthenticationMethodAPI.AddConjurAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddConjurAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addApiKeyConjurAuthenticationMethodRequest** | [**AddApiKeyConjurAuthenticationMethodRequest**](AddApiKeyConjurAuthenticationMethodRequest.md) | Create a new Conjur Authentication Method in the config | 

### Return type

[**ApiKeyConjurAuthenticationMethodResponse**](ApiKeyConjurAuthenticationMethodResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConjurAuthenticationMethod

> DeleteConjurAuthenticationMethod(ctx, conjurAuthenticationMethodName).Execute()

Delete a Conjur Authentication Method

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
    conjurAuthenticationMethodName := "conjurAuthenticationMethodName_example" // string | Name of the Conjur Authentication Method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConjurAuthenticationMethodAPI.DeleteConjurAuthenticationMethod(context.Background(), conjurAuthenticationMethodName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConjurAuthenticationMethodAPI.DeleteConjurAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conjurAuthenticationMethodName** | **string** | Name of the Conjur Authentication Method | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConjurAuthenticationMethodRequest struct via the builder pattern


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


## GetConjurAuthenticationMethod

> ApiKeyConjurAuthenticationMethodResponse GetConjurAuthenticationMethod(ctx, conjurAuthenticationMethodName).Execute()

Returns a single Conjur Authentication Method

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
    conjurAuthenticationMethodName := "conjurAuthenticationMethodName_example" // string | Name of the Conjur Authentication Method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConjurAuthenticationMethodAPI.GetConjurAuthenticationMethod(context.Background(), conjurAuthenticationMethodName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConjurAuthenticationMethodAPI.GetConjurAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConjurAuthenticationMethod`: ApiKeyConjurAuthenticationMethodResponse
    fmt.Fprintf(os.Stdout, "Response from `ConjurAuthenticationMethodAPI.GetConjurAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conjurAuthenticationMethodName** | **string** | Name of the Conjur Authentication Method | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConjurAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKeyConjurAuthenticationMethodResponse**](ApiKeyConjurAuthenticationMethodResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConjurAuthenticationMethods

> ConjurAuthenticationMethodListResponse ListConjurAuthenticationMethods(ctx).Filter(filter).Execute()

Returns a list of all Conjur Authentication Method objects

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
    resp, r, err := apiClient.ConjurAuthenticationMethodAPI.ListConjurAuthenticationMethods(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConjurAuthenticationMethodAPI.ListConjurAuthenticationMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConjurAuthenticationMethods`: ConjurAuthenticationMethodListResponse
    fmt.Fprintf(os.Stdout, "Response from `ConjurAuthenticationMethodAPI.ListConjurAuthenticationMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConjurAuthenticationMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ConjurAuthenticationMethodListResponse**](ConjurAuthenticationMethodListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConjurAuthenticationMethod

> ApiKeyConjurAuthenticationMethodResponse UpdateConjurAuthenticationMethod(ctx, conjurAuthenticationMethodName).UpdateRequest(updateRequest).Execute()

Update an existing Conjur Authentication Method by name

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
    conjurAuthenticationMethodName := "conjurAuthenticationMethodName_example" // string | Name of the Conjur Authentication Method
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Conjur Authentication Method

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConjurAuthenticationMethodAPI.UpdateConjurAuthenticationMethod(context.Background(), conjurAuthenticationMethodName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConjurAuthenticationMethodAPI.UpdateConjurAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConjurAuthenticationMethod`: ApiKeyConjurAuthenticationMethodResponse
    fmt.Fprintf(os.Stdout, "Response from `ConjurAuthenticationMethodAPI.UpdateConjurAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conjurAuthenticationMethodName** | **string** | Name of the Conjur Authentication Method | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConjurAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Conjur Authentication Method | 

### Return type

[**ApiKeyConjurAuthenticationMethodResponse**](ApiKeyConjurAuthenticationMethodResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

