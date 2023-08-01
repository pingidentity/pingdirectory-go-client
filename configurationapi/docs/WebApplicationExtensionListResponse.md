# WebApplicationExtensionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]GetWebApplicationExtension200Response**](GetWebApplicationExtension200Response.md) |  | [optional] 

## Methods

### NewWebApplicationExtensionListResponse

`func NewWebApplicationExtensionListResponse() *WebApplicationExtensionListResponse`

NewWebApplicationExtensionListResponse instantiates a new WebApplicationExtensionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebApplicationExtensionListResponseWithDefaults

`func NewWebApplicationExtensionListResponseWithDefaults() *WebApplicationExtensionListResponse`

NewWebApplicationExtensionListResponseWithDefaults instantiates a new WebApplicationExtensionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *WebApplicationExtensionListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *WebApplicationExtensionListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *WebApplicationExtensionListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *WebApplicationExtensionListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *WebApplicationExtensionListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *WebApplicationExtensionListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *WebApplicationExtensionListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *WebApplicationExtensionListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *WebApplicationExtensionListResponse) GetResources() []GetWebApplicationExtension200Response`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *WebApplicationExtensionListResponse) GetResourcesOk() (*[]GetWebApplicationExtension200Response, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *WebApplicationExtensionListResponse) SetResources(v []GetWebApplicationExtension200Response)`

SetResources sets Resources field to given value.

### HasResources

`func (o *WebApplicationExtensionListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


