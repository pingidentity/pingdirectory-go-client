# \VirtualAttributeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVirtualAttribute**](VirtualAttributeAPI.md#AddVirtualAttribute) | **Post** /virtual-attributes | Add a new Virtual Attribute to the config
[**DeleteVirtualAttribute**](VirtualAttributeAPI.md#DeleteVirtualAttribute) | **Delete** /virtual-attributes/{virtual-attribute-name} | Delete a Virtual Attribute
[**GetVirtualAttribute**](VirtualAttributeAPI.md#GetVirtualAttribute) | **Get** /virtual-attributes/{virtual-attribute-name} | Returns a single Virtual Attribute
[**ListVirtualAttributes**](VirtualAttributeAPI.md#ListVirtualAttributes) | **Get** /virtual-attributes | Returns a list of all Virtual Attribute objects
[**UpdateVirtualAttribute**](VirtualAttributeAPI.md#UpdateVirtualAttribute) | **Patch** /virtual-attributes/{virtual-attribute-name} | Update an existing Virtual Attribute by name



## AddVirtualAttribute

> AddVirtualAttribute200Response AddVirtualAttribute(ctx).AddVirtualAttributeRequest(addVirtualAttributeRequest).Execute()

Add a new Virtual Attribute to the config

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
    addVirtualAttributeRequest := openapiclient.add_virtual_attribute_request{AddConstructedVirtualAttributeRequest: openapiclient.NewAddConstructedVirtualAttributeRequest([]openapiclient.EnumconstructedVirtualAttributeSchemaUrn{openapiclient.Enumconstructed-virtual-attributeSchemaUrn("urn:pingidentity:schemas:configuration:2.0:virtual-attribute:constructed")}, []string{"ValuePattern_example"}, false, "AttributeType_example", "Name_example")} // AddVirtualAttributeRequest | Create a new Virtual Attribute in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualAttributeAPI.AddVirtualAttribute(context.Background()).AddVirtualAttributeRequest(addVirtualAttributeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualAttributeAPI.AddVirtualAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVirtualAttribute`: AddVirtualAttribute200Response
    fmt.Fprintf(os.Stdout, "Response from `VirtualAttributeAPI.AddVirtualAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVirtualAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addVirtualAttributeRequest** | [**AddVirtualAttributeRequest**](AddVirtualAttributeRequest.md) | Create a new Virtual Attribute in the config | 

### Return type

[**AddVirtualAttribute200Response**](AddVirtualAttribute200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualAttribute

> DeleteVirtualAttribute(ctx, virtualAttributeName).Execute()

Delete a Virtual Attribute

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
    virtualAttributeName := "virtualAttributeName_example" // string | Name of the Virtual Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VirtualAttributeAPI.DeleteVirtualAttribute(context.Background(), virtualAttributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualAttributeAPI.DeleteVirtualAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualAttributeName** | **string** | Name of the Virtual Attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualAttributeRequest struct via the builder pattern


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


## GetVirtualAttribute

> GetVirtualAttribute200Response GetVirtualAttribute(ctx, virtualAttributeName).Execute()

Returns a single Virtual Attribute

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
    virtualAttributeName := "virtualAttributeName_example" // string | Name of the Virtual Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualAttributeAPI.GetVirtualAttribute(context.Background(), virtualAttributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualAttributeAPI.GetVirtualAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualAttribute`: GetVirtualAttribute200Response
    fmt.Fprintf(os.Stdout, "Response from `VirtualAttributeAPI.GetVirtualAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualAttributeName** | **string** | Name of the Virtual Attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVirtualAttribute200Response**](GetVirtualAttribute200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualAttributes

> VirtualAttributeListResponse ListVirtualAttributes(ctx).Filter(filter).Execute()

Returns a list of all Virtual Attribute objects

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
    resp, r, err := apiClient.VirtualAttributeAPI.ListVirtualAttributes(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualAttributeAPI.ListVirtualAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualAttributes`: VirtualAttributeListResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualAttributeAPI.ListVirtualAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**VirtualAttributeListResponse**](VirtualAttributeListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualAttribute

> GetVirtualAttribute200Response UpdateVirtualAttribute(ctx, virtualAttributeName).UpdateRequest(updateRequest).Execute()

Update an existing Virtual Attribute by name

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
    virtualAttributeName := "virtualAttributeName_example" // string | Name of the Virtual Attribute
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Virtual Attribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualAttributeAPI.UpdateVirtualAttribute(context.Background(), virtualAttributeName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualAttributeAPI.UpdateVirtualAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualAttribute`: GetVirtualAttribute200Response
    fmt.Fprintf(os.Stdout, "Response from `VirtualAttributeAPI.UpdateVirtualAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualAttributeName** | **string** | Name of the Virtual Attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Virtual Attribute | 

### Return type

[**GetVirtualAttribute200Response**](GetVirtualAttribute200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

