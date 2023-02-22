# AddTimeLimitLogRotationPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | **string** | Name of the new Log Rotation Policy | 
**Schemas** | [**[]EnumtimeLimitLogRotationPolicySchemaUrn**](EnumtimeLimitLogRotationPolicySchemaUrn.md) |  | 
**RotationInterval** | **string** | Specifies the time interval between rotations. | 
**Description** | Pointer to **string** | A description for this Log Rotation Policy | [optional] 

## Methods

### NewAddTimeLimitLogRotationPolicyRequest

`func NewAddTimeLimitLogRotationPolicyRequest(policyName string, schemas []EnumtimeLimitLogRotationPolicySchemaUrn, rotationInterval string, ) *AddTimeLimitLogRotationPolicyRequest`

NewAddTimeLimitLogRotationPolicyRequest instantiates a new AddTimeLimitLogRotationPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTimeLimitLogRotationPolicyRequestWithDefaults

`func NewAddTimeLimitLogRotationPolicyRequestWithDefaults() *AddTimeLimitLogRotationPolicyRequest`

NewAddTimeLimitLogRotationPolicyRequestWithDefaults instantiates a new AddTimeLimitLogRotationPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *AddTimeLimitLogRotationPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddTimeLimitLogRotationPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddTimeLimitLogRotationPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetSchemas

`func (o *AddTimeLimitLogRotationPolicyRequest) GetSchemas() []EnumtimeLimitLogRotationPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTimeLimitLogRotationPolicyRequest) GetSchemasOk() (*[]EnumtimeLimitLogRotationPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTimeLimitLogRotationPolicyRequest) SetSchemas(v []EnumtimeLimitLogRotationPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRotationInterval

`func (o *AddTimeLimitLogRotationPolicyRequest) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *AddTimeLimitLogRotationPolicyRequest) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *AddTimeLimitLogRotationPolicyRequest) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.


### GetDescription

`func (o *AddTimeLimitLogRotationPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTimeLimitLogRotationPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTimeLimitLogRotationPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTimeLimitLogRotationPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


