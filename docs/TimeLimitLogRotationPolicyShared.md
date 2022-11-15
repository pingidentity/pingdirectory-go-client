# TimeLimitLogRotationPolicyShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumtimeLimitLogRotationPolicySchemaUrn**](EnumtimeLimitLogRotationPolicySchemaUrn.md) |  | 
**RotationInterval** | **string** | Specifies the time interval between rotations. | 
**Description** | Pointer to **string** | A description for this Log Rotation Policy | [optional] 

## Methods

### NewTimeLimitLogRotationPolicyShared

`func NewTimeLimitLogRotationPolicyShared(schemas []EnumtimeLimitLogRotationPolicySchemaUrn, rotationInterval string, ) *TimeLimitLogRotationPolicyShared`

NewTimeLimitLogRotationPolicyShared instantiates a new TimeLimitLogRotationPolicyShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeLimitLogRotationPolicySharedWithDefaults

`func NewTimeLimitLogRotationPolicySharedWithDefaults() *TimeLimitLogRotationPolicyShared`

NewTimeLimitLogRotationPolicySharedWithDefaults instantiates a new TimeLimitLogRotationPolicyShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *TimeLimitLogRotationPolicyShared) GetSchemas() []EnumtimeLimitLogRotationPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TimeLimitLogRotationPolicyShared) GetSchemasOk() (*[]EnumtimeLimitLogRotationPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TimeLimitLogRotationPolicyShared) SetSchemas(v []EnumtimeLimitLogRotationPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRotationInterval

`func (o *TimeLimitLogRotationPolicyShared) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *TimeLimitLogRotationPolicyShared) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *TimeLimitLogRotationPolicyShared) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.


### GetDescription

`func (o *TimeLimitLogRotationPolicyShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TimeLimitLogRotationPolicyShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TimeLimitLogRotationPolicyShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TimeLimitLogRotationPolicyShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


