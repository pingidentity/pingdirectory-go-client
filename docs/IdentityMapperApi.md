# \IdentityMapperApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIdentityMapper**](IdentityMapperApi.md#AddIdentityMapper) | **Post** /identity-mappers | Add a new Identity Mapper to the config
[**DeleteIdentityMapper**](IdentityMapperApi.md#DeleteIdentityMapper) | **Delete** /identity-mappers/{identity-mapper-name} | Delete a Identity Mapper
[**GetIdentityMapper**](IdentityMapperApi.md#GetIdentityMapper) | **Get** /identity-mappers/{identity-mapper-name} | Returns a single Identity Mapper
[**UpdateIdentityMapper**](IdentityMapperApi.md#UpdateIdentityMapper) | **Patch** /identity-mappers/{identity-mapper-name} | Update an existing Identity Mapper by name



## AddIdentityMapper

> AddIdentityMapper200Response AddIdentityMapper(ctx).AddIdentityMapperRequest(addIdentityMapperRequest).Execute()

Add a new Identity Mapper to the config

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    addIdentityMapperRequest := openapiclient.add_identity_mapper_request{AddAggregateIdentityMapperRequest: openapiclient.NewAddAggregateIdentityMapperRequest("MapperName_example", []openapiclient.EnumaggregateIdentityMapperSchemaUrn{openapiclient.Enumaggregate-identity-mapperSchemaUrn("urn:pingidentity:schemas:configuration:2.0:identity-mapper:aggregate")}, false)} // AddIdentityMapperRequest | Create a new Identity Mapper in the config

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityMapperApi.AddIdentityMapper(context.Background()).AddIdentityMapperRequest(addIdentityMapperRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityMapperApi.AddIdentityMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddIdentityMapper`: AddIdentityMapper200Response
    fmt.Fprintf(os.Stdout, "Response from `IdentityMapperApi.AddIdentityMapper`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIdentityMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addIdentityMapperRequest** | [**AddIdentityMapperRequest**](AddIdentityMapperRequest.md) | Create a new Identity Mapper in the config | 

### Return type

[**AddIdentityMapper200Response**](AddIdentityMapper200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityMapper

> DeleteIdentityMapper(ctx, identityMapperName).Execute()

Delete a Identity Mapper

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityMapperName := "identityMapperName_example" // string | Name of the Identity Mapper to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityMapperApi.DeleteIdentityMapper(context.Background(), identityMapperName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityMapperApi.DeleteIdentityMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityMapperName** | **string** | Name of the Identity Mapper to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityMapperRequest struct via the builder pattern


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


## GetIdentityMapper

> AddIdentityMapper200Response GetIdentityMapper(ctx, identityMapperName).Execute()

Returns a single Identity Mapper

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityMapperName := "identityMapperName_example" // string | Name of the Identity Mapper to be read

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityMapperApi.GetIdentityMapper(context.Background(), identityMapperName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityMapperApi.GetIdentityMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityMapper`: AddIdentityMapper200Response
    fmt.Fprintf(os.Stdout, "Response from `IdentityMapperApi.GetIdentityMapper`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityMapperName** | **string** | Name of the Identity Mapper to be read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddIdentityMapper200Response**](AddIdentityMapper200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityMapper

> AddIdentityMapper200Response UpdateIdentityMapper(ctx, identityMapperName).UpdateRequest(updateRequest).Execute()

Update an existing Identity Mapper by name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityMapperName := "identityMapperName_example" // string | Name of the Identity Mapper to be updated
    updateRequest := *openapiclient.NewUpdateRequest([]openapiclient.Operation{*openapiclient.NewOperation(openapiclient.EnumOperation("add"), "Path_example")}) // UpdateRequest | Update an existing Identity Mapper

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityMapperApi.UpdateIdentityMapper(context.Background(), identityMapperName).UpdateRequest(updateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityMapperApi.UpdateIdentityMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentityMapper`: AddIdentityMapper200Response
    fmt.Fprintf(os.Stdout, "Response from `IdentityMapperApi.UpdateIdentityMapper`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityMapperName** | **string** | Name of the Identity Mapper to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequest** | [**UpdateRequest**](UpdateRequest.md) | Update an existing Identity Mapper | 

### Return type

[**AddIdentityMapper200Response**](AddIdentityMapper200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

