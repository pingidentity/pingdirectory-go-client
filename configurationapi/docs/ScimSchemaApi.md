# \ScimSchemaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScimSchema**](ScimSchemaApi.md#AddScimSchema) | **Post** /scim-schemas | Add a new SCIM Schema to the config
[**DeleteScimSchema**](ScimSchemaApi.md#DeleteScimSchema) | **Delete** /scim-schemas/{scim-schema-name} | Delete a SCIM Schema
[**GetScimSchema**](ScimSchemaApi.md#GetScimSchema) | **Get** /scim-schemas/{scim-schema-name} | Returns a single SCIM Schema
[**ListScimSchemas**](ScimSchemaApi.md#ListScimSchemas) | **Get** /scim-schemas | Returns a list of all SCIM Schema objects
[**UpdateScimSchema**](ScimSchemaApi.md#UpdateScimSchema) | **Patch** /scim-schemas/{scim-schema-name} | Update an existing SCIM Schema by name



## AddScimSchema

> ScimSchemaResponse AddScimSchema(ctx).AddScimSchemaRequest(addScimSchemaRequest).Execute()

Add a new SCIM Schema to the config

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
    addScimSchemaRequest := *openapiclient.NewAddScimSchemaRequest("SchemaName_example", "SchemaURN_example") // AddScimSchemaRequest | Create a new SCIM Schema in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimSchemaApi.AddScimSchema(context.Background()).AddScimSchemaRequest(addScimSchemaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimSchemaApi.AddScimSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddScimSchema`: ScimSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimSchemaApi.AddScimSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddScimSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addScimSchemaRequest** | [**AddScimSchemaRequest**](AddScimSchemaRequest.md) | Create a new SCIM Schema in the config | 

### Return type

[**ScimSchemaResponse**](ScimSchemaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScimSchema

> DeleteScimSchema(ctx, scimSchemaName).Execute()

Delete a SCIM Schema

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
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScimSchemaApi.DeleteScimSchema(context.Background(), scimSchemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimSchemaApi.DeleteScimSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScimSchemaRequest struct via the builder pattern


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


## GetScimSchema

> ScimSchemaResponse GetScimSchema(ctx, scimSchemaName).Execute()

Returns a single SCIM Schema

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
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimSchemaApi.GetScimSchema(context.Background(), scimSchemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimSchemaApi.GetScimSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScimSchema`: ScimSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimSchemaApi.GetScimSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScimSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScimSchemaResponse**](ScimSchemaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScimSchemas

> ScimSchemaListResponse ListScimSchemas(ctx).Filter(filter).Execute()

Returns a list of all SCIM Schema objects

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
    resp, r, err := apiClient.ScimSchemaApi.ListScimSchemas(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimSchemaApi.ListScimSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScimSchemas`: ScimSchemaListResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimSchemaApi.ListScimSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListScimSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter | 

### Return type

[**ScimSchemaListResponse**](ScimSchemaListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScimSchema

> ScimSchemaResponse UpdateScimSchema(ctx, scimSchemaName).UpdateRequest(updateRequest).Execute()

Update an existing SCIM Schema by name

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
    scimSchemaName := "scimSchemaName_example" // string | Name of the SCIM Schema
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing SCIM Schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScimSchemaApi.UpdateScimSchema(context.Background(), scimSchemaName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScimSchemaApi.UpdateScimSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScimSchema`: ScimSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `ScimSchemaApi.UpdateScimSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scimSchemaName** | **string** | Name of the SCIM Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScimSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing SCIM Schema | 

### Return type

[**ScimSchemaResponse**](ScimSchemaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

