# AddTimeLimitLogRetentionPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | **string** | Name of the new Log Retention Policy | 
**Schemas** | [**[]EnumtimeLimitLogRetentionPolicySchemaUrn**](EnumtimeLimitLogRetentionPolicySchemaUrn.md) |  | 
**RetainDuration** | **string** | Specifies the desired minimum length of time that each log file should be retained. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 

## Methods

### NewAddTimeLimitLogRetentionPolicyRequest

`func NewAddTimeLimitLogRetentionPolicyRequest(policyName string, schemas []EnumtimeLimitLogRetentionPolicySchemaUrn, retainDuration string, ) *AddTimeLimitLogRetentionPolicyRequest`

NewAddTimeLimitLogRetentionPolicyRequest instantiates a new AddTimeLimitLogRetentionPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTimeLimitLogRetentionPolicyRequestWithDefaults

`func NewAddTimeLimitLogRetentionPolicyRequestWithDefaults() *AddTimeLimitLogRetentionPolicyRequest`

NewAddTimeLimitLogRetentionPolicyRequestWithDefaults instantiates a new AddTimeLimitLogRetentionPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *AddTimeLimitLogRetentionPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddTimeLimitLogRetentionPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddTimeLimitLogRetentionPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetSchemas

`func (o *AddTimeLimitLogRetentionPolicyRequest) GetSchemas() []EnumtimeLimitLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTimeLimitLogRetentionPolicyRequest) GetSchemasOk() (*[]EnumtimeLimitLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTimeLimitLogRetentionPolicyRequest) SetSchemas(v []EnumtimeLimitLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRetainDuration

`func (o *AddTimeLimitLogRetentionPolicyRequest) GetRetainDuration() string`

GetRetainDuration returns the RetainDuration field if non-nil, zero value otherwise.

### GetRetainDurationOk

`func (o *AddTimeLimitLogRetentionPolicyRequest) GetRetainDurationOk() (*string, bool)`

GetRetainDurationOk returns a tuple with the RetainDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainDuration

`func (o *AddTimeLimitLogRetentionPolicyRequest) SetRetainDuration(v string)`

SetRetainDuration sets RetainDuration field to given value.


### GetDescription

`func (o *AddTimeLimitLogRetentionPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTimeLimitLogRetentionPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTimeLimitLogRetentionPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTimeLimitLogRetentionPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


