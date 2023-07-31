# \AttributeSyntaxApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttributeSyntax**](AttributeSyntaxApi.md#GetAttributeSyntax) | **Get** /attribute-syntaxes/{attribute-syntax-name} | Returns a single Attribute Syntax
[**ListAttributeSyntaxes**](AttributeSyntaxApi.md#ListAttributeSyntaxes) | **Get** /attribute-syntaxes | Returns a list of all Attribute Syntax objects
[**UpdateAttributeSyntax**](AttributeSyntaxApi.md#UpdateAttributeSyntax) | **Patch** /attribute-syntaxes/{attribute-syntax-name} | Update an existing Attribute Syntax by name



## GetAttributeSyntax

> GetAttributeSyntax200Response GetAttributeSyntax(ctx, attributeSyntaxName).Execute()

Returns a single Attribute Syntax

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
    attributeSyntaxName := "attributeSyntaxName_example" // string | Name of the Attribute Syntax

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttributeSyntaxApi.GetAttributeSyntax(context.Background(), attributeSyntaxName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttributeSyntaxApi.GetAttributeSyntax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributeSyntax`: GetAttributeSyntax200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeSyntaxApi.GetAttributeSyntax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeSyntaxName** | **string** | Name of the Attribute Syntax | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeSyntaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAttributeSyntax200Response**](GetAttributeSyntax200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttributeSyntaxes

> AttributeSyntaxListResponse ListAttributeSyntaxes(ctx).Filter(filter).Execute()

Returns a list of all Attribute Syntax objects

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
    resp, r, err := apiClient.AttributeSyntaxApi.ListAttributeSyntaxes(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttributeSyntaxApi.ListAttributeSyntaxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAttributeSyntaxes`: AttributeSyntaxListResponse
    fmt.Fprintf(os.Stdout, "Response from `AttributeSyntaxApi.ListAttributeSyntaxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAttributeSyntaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**AttributeSyntaxListResponse**](AttributeSyntaxListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttributeSyntax

> GetAttributeSyntax200Response UpdateAttributeSyntax(ctx, attributeSyntaxName).UpdateRequest(updateRequest).Execute()

Update an existing Attribute Syntax by name

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
    attributeSyntaxName := "attributeSyntaxName_example" // string | Name of the Attribute Syntax
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Attribute Syntax

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttributeSyntaxApi.UpdateAttributeSyntax(context.Background(), attributeSyntaxName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttributeSyntaxApi.UpdateAttributeSyntax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAttributeSyntax`: GetAttributeSyntax200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeSyntaxApi.UpdateAttributeSyntax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeSyntaxName** | **string** | Name of the Attribute Syntax | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributeSyntaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Attribute Syntax | 

### Return type

[**GetAttributeSyntax200Response**](GetAttributeSyntax200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

