# StaticGroupImplementationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumstaticGroupImplementationSchemaUrn**](EnumstaticGroupImplementationSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Group Implementation | [optional] 
**Enabled** | **bool** | Indicates whether the Group Implementation is enabled. | 

## Methods

### NewStaticGroupImplementationResponse

`func NewStaticGroupImplementationResponse(schemas []EnumstaticGroupImplementationSchemaUrn, enabled bool, ) *StaticGroupImplementationResponse`

NewStaticGroupImplementationResponse instantiates a new StaticGroupImplementationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticGroupImplementationResponseWithDefaults

`func NewStaticGroupImplementationResponseWithDefaults() *StaticGroupImplementationResponse`

NewStaticGroupImplementationResponseWithDefaults instantiates a new StaticGroupImplementationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *StaticGroupImplementationResponse) GetSchemas() []EnumstaticGroupImplementationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *StaticGroupImplementationResponse) GetSchemasOk() (*[]EnumstaticGroupImplementationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *StaticGroupImplementationResponse) SetSchemas(v []EnumstaticGroupImplementationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *StaticGroupImplementationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StaticGroupImplementationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StaticGroupImplementationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StaticGroupImplementationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *StaticGroupImplementationResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StaticGroupImplementationResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StaticGroupImplementationResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


