# AddLogRotationPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | **string** | Name of the new Log Rotation Policy | 
**Schemas** | [**[]EnumsizeLimitLogRotationPolicySchemaUrn**](EnumsizeLimitLogRotationPolicySchemaUrn.md) |  | 
**RotationInterval** | **string** | Specifies the time interval between rotations. | 
**Description** | Pointer to **string** | A description for this Log Rotation Policy | [optional] 
**TimeOfDay** | **[]string** | Specifies the time of day at which log rotation should occur. | 
**FileSizeLimit** | **string** | Specifies the maximum size that a log file can reach before it is rotated. | 

## Methods

### NewAddLogRotationPolicyRequest

`func NewAddLogRotationPolicyRequest(policyName string, schemas []EnumsizeLimitLogRotationPolicySchemaUrn, rotationInterval string, timeOfDay []string, fileSizeLimit string, ) *AddLogRotationPolicyRequest`

NewAddLogRotationPolicyRequest instantiates a new AddLogRotationPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLogRotationPolicyRequestWithDefaults

`func NewAddLogRotationPolicyRequestWithDefaults() *AddLogRotationPolicyRequest`

NewAddLogRotationPolicyRequestWithDefaults instantiates a new AddLogRotationPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *AddLogRotationPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddLogRotationPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddLogRotationPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetSchemas

`func (o *AddLogRotationPolicyRequest) GetSchemas() []EnumsizeLimitLogRotationPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLogRotationPolicyRequest) GetSchemasOk() (*[]EnumsizeLimitLogRotationPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLogRotationPolicyRequest) SetSchemas(v []EnumsizeLimitLogRotationPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRotationInterval

`func (o *AddLogRotationPolicyRequest) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *AddLogRotationPolicyRequest) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *AddLogRotationPolicyRequest) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.


### GetDescription

`func (o *AddLogRotationPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLogRotationPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLogRotationPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLogRotationPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimeOfDay

`func (o *AddLogRotationPolicyRequest) GetTimeOfDay() []string`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *AddLogRotationPolicyRequest) GetTimeOfDayOk() (*[]string, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *AddLogRotationPolicyRequest) SetTimeOfDay(v []string)`

SetTimeOfDay sets TimeOfDay field to given value.


### GetFileSizeLimit

`func (o *AddLogRotationPolicyRequest) GetFileSizeLimit() string`

GetFileSizeLimit returns the FileSizeLimit field if non-nil, zero value otherwise.

### GetFileSizeLimitOk

`func (o *AddLogRotationPolicyRequest) GetFileSizeLimitOk() (*string, bool)`

GetFileSizeLimitOk returns a tuple with the FileSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeLimit

`func (o *AddLogRotationPolicyRequest) SetFileSizeLimit(v string)`

SetFileSizeLimit sets FileSizeLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


