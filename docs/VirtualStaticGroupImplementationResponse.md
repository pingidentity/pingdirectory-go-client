# VirtualStaticGroupImplementationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumvirtualStaticGroupImplementationSchemaUrn**](EnumvirtualStaticGroupImplementationSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Group Implementation | [optional] 
**Enabled** | **bool** | Indicates whether the Group Implementation is enabled. | 

## Methods

### NewVirtualStaticGroupImplementationResponse

`func NewVirtualStaticGroupImplementationResponse(schemas []EnumvirtualStaticGroupImplementationSchemaUrn, enabled bool, ) *VirtualStaticGroupImplementationResponse`

NewVirtualStaticGroupImplementationResponse instantiates a new VirtualStaticGroupImplementationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualStaticGroupImplementationResponseWithDefaults

`func NewVirtualStaticGroupImplementationResponseWithDefaults() *VirtualStaticGroupImplementationResponse`

NewVirtualStaticGroupImplementationResponseWithDefaults instantiates a new VirtualStaticGroupImplementationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *VirtualStaticGroupImplementationResponse) GetSchemas() []EnumvirtualStaticGroupImplementationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *VirtualStaticGroupImplementationResponse) GetSchemasOk() (*[]EnumvirtualStaticGroupImplementationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *VirtualStaticGroupImplementationResponse) SetSchemas(v []EnumvirtualStaticGroupImplementationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *VirtualStaticGroupImplementationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualStaticGroupImplementationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualStaticGroupImplementationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualStaticGroupImplementationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *VirtualStaticGroupImplementationResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VirtualStaticGroupImplementationResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VirtualStaticGroupImplementationResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


