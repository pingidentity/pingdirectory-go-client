# \LogFieldSyntaxAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogFieldSyntax**](LogFieldSyntaxAPI.md#GetLogFieldSyntax) | **Get** /log-field-syntaxes/{log-field-syntax-name} | Returns a single Log Field Syntax
[**ListLogFieldSyntaxes**](LogFieldSyntaxAPI.md#ListLogFieldSyntaxes) | **Get** /log-field-syntaxes | Returns a list of all Log Field Syntax objects
[**UpdateLogFieldSyntax**](LogFieldSyntaxAPI.md#UpdateLogFieldSyntax) | **Patch** /log-field-syntaxes/{log-field-syntax-name} | Update an existing Log Field Syntax by name



## GetLogFieldSyntax

> GetLogFieldSyntax200Response GetLogFieldSyntax(ctx, logFieldSyntaxName).Execute()

Returns a single Log Field Syntax

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
    logFieldSyntaxName := "logFieldSyntaxName_example" // string | Name of the Log Field Syntax

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldSyntaxAPI.GetLogFieldSyntax(context.Background(), logFieldSyntaxName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldSyntaxAPI.GetLogFieldSyntax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogFieldSyntax`: GetLogFieldSyntax200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFieldSyntaxAPI.GetLogFieldSyntax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logFieldSyntaxName** | **string** | Name of the Log Field Syntax | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogFieldSyntaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLogFieldSyntax200Response**](GetLogFieldSyntax200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogFieldSyntaxes

> LogFieldSyntaxListResponse ListLogFieldSyntaxes(ctx).Filter(filter).Execute()

Returns a list of all Log Field Syntax objects

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
    resp, r, err := apiClient.LogFieldSyntaxAPI.ListLogFieldSyntaxes(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldSyntaxAPI.ListLogFieldSyntaxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogFieldSyntaxes`: LogFieldSyntaxListResponse
    fmt.Fprintf(os.Stdout, "Response from `LogFieldSyntaxAPI.ListLogFieldSyntaxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogFieldSyntaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**LogFieldSyntaxListResponse**](LogFieldSyntaxListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogFieldSyntax

> GetLogFieldSyntax200Response UpdateLogFieldSyntax(ctx, logFieldSyntaxName).UpdateRequest(updateRequest).Execute()

Update an existing Log Field Syntax by name

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
    logFieldSyntaxName := "logFieldSyntaxName_example" // string | Name of the Log Field Syntax
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Log Field Syntax

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogFieldSyntaxAPI.UpdateLogFieldSyntax(context.Background(), logFieldSyntaxName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogFieldSyntaxAPI.UpdateLogFieldSyntax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogFieldSyntax`: GetLogFieldSyntax200Response
    fmt.Fprintf(os.Stdout, "Response from `LogFieldSyntaxAPI.UpdateLogFieldSyntax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**logFieldSyntaxName** | **string** | Name of the Log Field Syntax | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogFieldSyntaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Log Field Syntax | 

### Return type

[**GetLogFieldSyntax200Response**](GetLogFieldSyntax200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

