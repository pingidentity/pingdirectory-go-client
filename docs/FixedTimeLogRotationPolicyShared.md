# FixedTimeLogRotationPolicyShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfixedTimeLogRotationPolicySchemaUrn**](EnumfixedTimeLogRotationPolicySchemaUrn.md) |  | 
**TimeOfDay** | **[]string** |  | 
**Description** | Pointer to **string** | A description for this Log Rotation Policy | [optional] 

## Methods

### NewFixedTimeLogRotationPolicyShared

`func NewFixedTimeLogRotationPolicyShared(schemas []EnumfixedTimeLogRotationPolicySchemaUrn, timeOfDay []string, ) *FixedTimeLogRotationPolicyShared`

NewFixedTimeLogRotationPolicyShared instantiates a new FixedTimeLogRotationPolicyShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedTimeLogRotationPolicySharedWithDefaults

`func NewFixedTimeLogRotationPolicySharedWithDefaults() *FixedTimeLogRotationPolicyShared`

NewFixedTimeLogRotationPolicySharedWithDefaults instantiates a new FixedTimeLogRotationPolicyShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FixedTimeLogRotationPolicyShared) GetSchemas() []EnumfixedTimeLogRotationPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FixedTimeLogRotationPolicyShared) GetSchemasOk() (*[]EnumfixedTimeLogRotationPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FixedTimeLogRotationPolicyShared) SetSchemas(v []EnumfixedTimeLogRotationPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTimeOfDay

`func (o *FixedTimeLogRotationPolicyShared) GetTimeOfDay() []string`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *FixedTimeLogRotationPolicyShared) GetTimeOfDayOk() (*[]string, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *FixedTimeLogRotationPolicyShared) SetTimeOfDay(v []string)`

SetTimeOfDay sets TimeOfDay field to given value.


### GetDescription

`func (o *FixedTimeLogRotationPolicyShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FixedTimeLogRotationPolicyShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FixedTimeLogRotationPolicyShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FixedTimeLogRotationPolicyShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


