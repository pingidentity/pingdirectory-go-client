# \SensitiveAttributeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSensitiveAttribute**](SensitiveAttributeAPI.md#AddSensitiveAttribute) | **Post** /sensitive-attributes | Add a new Sensitive Attribute to the config
[**DeleteSensitiveAttribute**](SensitiveAttributeAPI.md#DeleteSensitiveAttribute) | **Delete** /sensitive-attributes/{sensitive-attribute-name} | Delete a Sensitive Attribute
[**GetSensitiveAttribute**](SensitiveAttributeAPI.md#GetSensitiveAttribute) | **Get** /sensitive-attributes/{sensitive-attribute-name} | Returns a single Sensitive Attribute
[**ListSensitiveAttributes**](SensitiveAttributeAPI.md#ListSensitiveAttributes) | **Get** /sensitive-attributes | Returns a list of all Sensitive Attribute objects
[**UpdateSensitiveAttribute**](SensitiveAttributeAPI.md#UpdateSensitiveAttribute) | **Patch** /sensitive-attributes/{sensitive-attribute-name} | Update an existing Sensitive Attribute by name



## AddSensitiveAttribute

> SensitiveAttributeResponse AddSensitiveAttribute(ctx).AddSensitiveAttributeRequest(addSensitiveAttributeRequest).Execute()

Add a new Sensitive Attribute to the config

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
    addSensitiveAttributeRequest := *openapiclient.NewAddSensitiveAttributeRequest([]string{"AttributeType_example"}, "AttributeName_example") // AddSensitiveAttributeRequest | Create a new Sensitive Attribute in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SensitiveAttributeAPI.AddSensitiveAttribute(context.Background()).AddSensitiveAttributeRequest(addSensitiveAttributeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensitiveAttributeAPI.AddSensitiveAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSensitiveAttribute`: SensitiveAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `SensitiveAttributeAPI.AddSensitiveAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSensitiveAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addSensitiveAttributeRequest** | [**AddSensitiveAttributeRequest**](AddSensitiveAttributeRequest.md) | Create a new Sensitive Attribute in the config | 

### Return type

[**SensitiveAttributeResponse**](SensitiveAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSensitiveAttribute

> DeleteSensitiveAttribute(ctx, sensitiveAttributeName).Execute()

Delete a Sensitive Attribute

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
    sensitiveAttributeName := "sensitiveAttributeName_example" // string | Name of the Sensitive Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SensitiveAttributeAPI.DeleteSensitiveAttribute(context.Background(), sensitiveAttributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensitiveAttributeAPI.DeleteSensitiveAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sensitiveAttributeName** | **string** | Name of the Sensitive Attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSensitiveAttributeRequest struct via the builder pattern


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


## GetSensitiveAttribute

> SensitiveAttributeResponse GetSensitiveAttribute(ctx, sensitiveAttributeName).Execute()

Returns a single Sensitive Attribute

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
    sensitiveAttributeName := "sensitiveAttributeName_example" // string | Name of the Sensitive Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SensitiveAttributeAPI.GetSensitiveAttribute(context.Background(), sensitiveAttributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensitiveAttributeAPI.GetSensitiveAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSensitiveAttribute`: SensitiveAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `SensitiveAttributeAPI.GetSensitiveAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sensitiveAttributeName** | **string** | Name of the Sensitive Attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSensitiveAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SensitiveAttributeResponse**](SensitiveAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSensitiveAttributes

> SensitiveAttributeListResponse ListSensitiveAttributes(ctx).Filter(filter).Execute()

Returns a list of all Sensitive Attribute objects

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
    resp, r, err := apiClient.SensitiveAttributeAPI.ListSensitiveAttributes(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensitiveAttributeAPI.ListSensitiveAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSensitiveAttributes`: SensitiveAttributeListResponse
    fmt.Fprintf(os.Stdout, "Response from `SensitiveAttributeAPI.ListSensitiveAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSensitiveAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**SensitiveAttributeListResponse**](SensitiveAttributeListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSensitiveAttribute

> SensitiveAttributeResponse UpdateSensitiveAttribute(ctx, sensitiveAttributeName).UpdateRequest(updateRequest).Execute()

Update an existing Sensitive Attribute by name

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
    sensitiveAttributeName := "sensitiveAttributeName_example" // string | Name of the Sensitive Attribute
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Sensitive Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SensitiveAttributeAPI.UpdateSensitiveAttribute(context.Background(), sensitiveAttributeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensitiveAttributeAPI.UpdateSensitiveAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSensitiveAttribute`: SensitiveAttributeResponse
    fmt.Fprintf(os.Stdout, "Response from `SensitiveAttributeAPI.UpdateSensitiveAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sensitiveAttributeName** | **string** | Name of the Sensitive Attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSensitiveAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Sensitive Attribute | 

### Return type

[**SensitiveAttributeResponse**](SensitiveAttributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

