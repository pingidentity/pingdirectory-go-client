# \LdapSdkDebugLoggerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLdapSdkDebugLogger**](LdapSdkDebugLoggerAPI.md#GetLdapSdkDebugLogger) | **Get** /ldap-sdk-debug-logger | Returns a single LDAP SDK Debug Logger
[**UpdateLdapSdkDebugLogger**](LdapSdkDebugLoggerAPI.md#UpdateLdapSdkDebugLogger) | **Patch** /ldap-sdk-debug-logger | Update an existing LDAP SDK Debug Logger by name



## GetLdapSdkDebugLogger

> LdapSdkDebugLoggerResponse GetLdapSdkDebugLogger(ctx).Execute()

Returns a single LDAP SDK Debug Logger

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LdapSdkDebugLoggerAPI.GetLdapSdkDebugLogger(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapSdkDebugLoggerAPI.GetLdapSdkDebugLogger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLdapSdkDebugLogger`: LdapSdkDebugLoggerResponse
    fmt.Fprintf(os.Stdout, "Response from `LdapSdkDebugLoggerAPI.GetLdapSdkDebugLogger`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapSdkDebugLoggerRequest struct via the builder pattern


### Return type

[**LdapSdkDebugLoggerResponse**](LdapSdkDebugLoggerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLdapSdkDebugLogger

> LdapSdkDebugLoggerResponse UpdateLdapSdkDebugLogger(ctx).UpdateRequest(updateRequest).Execute()

Update an existing LDAP SDK Debug Logger by name

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
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing LDAP SDK Debug Logger

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LdapSdkDebugLoggerAPI.UpdateLdapSdkDebugLogger(context.Background()).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LdapSdkDebugLoggerAPI.UpdateLdapSdkDebugLogger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLdapSdkDebugLogger`: LdapSdkDebugLoggerResponse
    fmt.Fprintf(os.Stdout, "Response from `LdapSdkDebugLoggerAPI.UpdateLdapSdkDebugLogger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapSdkDebugLoggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing LDAP SDK Debug Logger | 

### Return type

[**LdapSdkDebugLoggerResponse**](LdapSdkDebugLoggerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

