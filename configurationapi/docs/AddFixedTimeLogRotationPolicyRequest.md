# AddFixedTimeLogRotationPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfixedTimeLogRotationPolicySchemaUrn**](EnumfixedTimeLogRotationPolicySchemaUrn.md) |  | 
**TimeOfDay** | **[]string** | Specifies the time of day at which log rotation should occur. | 
**Description** | Pointer to **string** | A description for this Log Rotation Policy | [optional] 
**PolicyName** | **string** | Name of the new Log Rotation Policy | 

## Methods

### NewAddFixedTimeLogRotationPolicyRequest

`func NewAddFixedTimeLogRotationPolicyRequest(schemas []EnumfixedTimeLogRotationPolicySchemaUrn, timeOfDay []string, policyName string, ) *AddFixedTimeLogRotationPolicyRequest`

NewAddFixedTimeLogRotationPolicyRequest instantiates a new AddFixedTimeLogRotationPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFixedTimeLogRotationPolicyRequestWithDefaults

`func NewAddFixedTimeLogRotationPolicyRequestWithDefaults() *AddFixedTimeLogRotationPolicyRequest`

NewAddFixedTimeLogRotationPolicyRequestWithDefaults instantiates a new AddFixedTimeLogRotationPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddFixedTimeLogRotationPolicyRequest) GetSchemas() []EnumfixedTimeLogRotationPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddFixedTimeLogRotationPolicyRequest) GetSchemasOk() (*[]EnumfixedTimeLogRotationPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddFixedTimeLogRotationPolicyRequest) SetSchemas(v []EnumfixedTimeLogRotationPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTimeOfDay

`func (o *AddFixedTimeLogRotationPolicyRequest) GetTimeOfDay() []string`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *AddFixedTimeLogRotationPolicyRequest) GetTimeOfDayOk() (*[]string, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *AddFixedTimeLogRotationPolicyRequest) SetTimeOfDay(v []string)`

SetTimeOfDay sets TimeOfDay field to given value.


### GetDescription

`func (o *AddFixedTimeLogRotationPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddFixedTimeLogRotationPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddFixedTimeLogRotationPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddFixedTimeLogRotationPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyName

`func (o *AddFixedTimeLogRotationPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddFixedTimeLogRotationPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddFixedTimeLogRotationPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


