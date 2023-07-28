# VirtualAttributeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **float64** |  | [optional] 
**Resources** | Pointer to [**[]GetVirtualAttribute200Response**](GetVirtualAttribute200Response.md) |  | [optional] 

## Methods

### NewVirtualAttributeListResponse

`func NewVirtualAttributeListResponse() *VirtualAttributeListResponse`

NewVirtualAttributeListResponse instantiates a new VirtualAttributeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAttributeListResponseWithDefaults

`func NewVirtualAttributeListResponseWithDefaults() *VirtualAttributeListResponse`

NewVirtualAttributeListResponseWithDefaults instantiates a new VirtualAttributeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *VirtualAttributeListResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VirtualAttributeListResponse) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VirtualAttributeListResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *VirtualAttributeListResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetTotalResults

`func (o *VirtualAttributeListResponse) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *VirtualAttributeListResponse) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *VirtualAttributeListResponse) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *VirtualAttributeListResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResources

`func (o *VirtualAttributeListResponse) GetResources() []GetVirtualAttribute200Response`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VirtualAttributeListResponse) GetResourcesOk() (*[]GetVirtualAttribute200Response, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VirtualAttributeListResponse) SetResources(v []GetVirtualAttribute200Response)`

SetResources sets Resources field to given value.

### HasResources

`func (o *VirtualAttributeListResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


