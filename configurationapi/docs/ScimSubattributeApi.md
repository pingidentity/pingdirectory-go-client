# \ScimSubattributeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScimSubattribute**](ScimSubattributeAPI.md#AddScimSubattribute) | **Post** /scim-schemas/{scim-schema-name}/scim-attributes/{scim-attribute-name}/scim-subattributes | Add a new SCIM Subattribute to the config
[**DeleteScimSubattribute**](ScimSubattributeAPI.md#DeleteScimSubattribute) | **Delete** /scim-schemas/{scim-schema-name}/scim-attributes/{scim-attribute-name}/scim-subattributes/{scim-subattribute-name} | Delete a SCIM Subattribute
[**GetScimSubattribute**](ScimSubattributeAPI.md#GetScimSubattribute) | **Get** /scim-schemas/{scim-schema-name}/scim-attributes/{scim-attribute-name}/scim-subattributes/{scim-subattribute-name} | Returns a single SCIM Subattribute
[**ListScimSubattributes**](ScimSubattributeAPI.md#ListScimSubattributes) | **Get** /scim-schemas/{scim-schema-name}/scim-attributes/{scim-attribute-name}/scim-subattributes | Returns a list of all SCIM Subattribute objects
[**UpdateScimSubattribute**](ScimSubattributeAPI.md#UpdateScimSubattribute) | **Patch** /scim-schemas/{scim-schema-name}/scim-attributes/{scim-attribute-name}/scim-subattributes/{scim-subattribute-name} | Update an existing SCIM Subattribute by name



## AddScimSubattribute

> ScimSubattributeResponse AddScimSubattribute(ctx, scimAttributeName, scimSchemaName).AddScimSubattributeRequest(addScimSubattributeRequest).Execute()

Add a new SCIM Subattribute to the config

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
    scimAttributeName := "scimAttributeName_example" // string | Name of the SCIM Attribute
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema
    addScimSubattributeRequest := *openapiclient.NewAddScimSubattributeRequest("SubattributeName_example") // AddScimSubattributeRequest | Create a new SCIM Subattribute in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimSubattributeAPI.AddScimSubattribute(context.Background(), scimAttributeName, scimSchemaName).AddScimSubattributeRequest(addScimSubattributeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimSubattributeAPI.AddScimSubattribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddScimSubattribute`: ScimSubattributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimSubattributeAPI.AddScimSubattribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimAttributeName** | **string** | Name of the SCIM Attribute | 
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddScimSubattributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addScimSubattributeRequest** | [**AddScimSubattributeRequest**](AddScimSubattributeRequest.md) | Create a new SCIM Subattribute in the config | 

### Return type

[**ScimSubattributeResponse**](ScimSubattributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScimSubattribute

> DeleteScimSubattribute(ctx, scimSubattributeName, scimAttributeName, scimSchemaName).Execute()

Delete a SCIM Subattribute

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
    scimSubattributeName := "scimSubattributeName_example" // string | Name of the SCIM Subattribute
    scimAttributeName := "scimAttributeName_example" // string | Name of the SCIM Attribute
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScimSubattributeAPI.DeleteScimSubattribute(context.Background(), scimSubattributeName, scimAttributeName, scimSchemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimSubattributeAPI.DeleteScimSubattribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimSubattributeName** | **string** | Name of the SCIM Subattribute | 
**scimAttributeName** | **string** | Name of the SCIM Attribute | 
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScimSubattributeRequest struct via the builder pattern


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


## GetScimSubattribute

> ScimSubattributeResponse GetScimSubattribute(ctx, scimSubattributeName, scimAttributeName, scimSchemaName).Execute()

Returns a single SCIM Subattribute

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
    scimSubattributeName := "scimSubattributeName_example" // string | Name of the SCIM Subattribute
    scimAttributeName := "scimAttributeName_example" // string | Name of the SCIM Attribute
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimSubattributeAPI.GetScimSubattribute(context.Background(), scimSubattributeName, scimAttributeName, scimSchemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimSubattributeAPI.GetScimSubattribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScimSubattribute`: ScimSubattributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimSubattributeAPI.GetScimSubattribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimSubattributeName** | **string** | Name of the SCIM Subattribute | 
**scimAttributeName** | **string** | Name of the SCIM Attribute | 
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScimSubattributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ScimSubattributeResponse**](ScimSubattributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScimSubattributes

> ScimSubattributeListResponse ListScimSubattributes(ctx, scimAttributeName, scimSchemaName).Filter(filter).Execute()

Returns a list of all SCIM Subattribute objects

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
    scimAttributeName := "scimAttributeName_example" // string | Name of the SCIM Attribute
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema
    filter := "filter_example" // string | SCIM filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimSubattributeAPI.ListScimSubattributes(context.Background(), scimAttributeName, scimSchemaName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimSubattributeAPI.ListScimSubattributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScimSubattributes`: ScimSubattributeListResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimSubattributeAPI.ListScimSubattributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimAttributeName** | **string** | Name of the SCIM Attribute | 
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiListScimSubattributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | SCIM filter | 

### Return type

[**ScimSubattributeListResponse**](ScimSubattributeListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScimSubattribute

> ScimSubattributeResponse UpdateScimSubattribute(ctx, scimSubattributeName, scimAttributeName, scimSchemaName).UpdateRequest(updateRequest).Execute()

Update an existing SCIM Subattribute by name

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
    scimSubattributeName := "scimSubattributeName_example" // string | Name of the SCIM Subattribute
    scimAttributeName := "scimAttributeName_example" // string | Name of the SCIM Attribute
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing SCIM Subattribute

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimSubattributeAPI.UpdateScimSubattribute(context.Background(), scimSubattributeName, scimAttributeName, scimSchemaName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimSubattributeAPI.UpdateScimSubattribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScimSubattribute`: ScimSubattributeResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimSubattributeAPI.UpdateScimSubattribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimSubattributeName** | **string** | Name of the SCIM Subattribute | 
**scimAttributeName** | **string** | Name of the SCIM Attribute | 
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScimSubattributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing SCIM Subattribute | 

### Return type

[**ScimSubattributeResponse**](ScimSubattributeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

