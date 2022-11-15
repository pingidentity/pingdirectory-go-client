# TimeLimitLogRetentionPolicyShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumtimeLimitLogRetentionPolicySchemaUrn**](EnumtimeLimitLogRetentionPolicySchemaUrn.md) |  | 
**RetainDuration** | **string** | Specifies the desired minimum length of time that each log file should be retained. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 

## Methods

### NewTimeLimitLogRetentionPolicyShared

`func NewTimeLimitLogRetentionPolicyShared(schemas []EnumtimeLimitLogRetentionPolicySchemaUrn, retainDuration string, ) *TimeLimitLogRetentionPolicyShared`

NewTimeLimitLogRetentionPolicyShared instantiates a new TimeLimitLogRetentionPolicyShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeLimitLogRetentionPolicySharedWithDefaults

`func NewTimeLimitLogRetentionPolicySharedWithDefaults() *TimeLimitLogRetentionPolicyShared`

NewTimeLimitLogRetentionPolicySharedWithDefaults instantiates a new TimeLimitLogRetentionPolicyShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *TimeLimitLogRetentionPolicyShared) GetSchemas() []EnumtimeLimitLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TimeLimitLogRetentionPolicyShared) GetSchemasOk() (*[]EnumtimeLimitLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TimeLimitLogRetentionPolicyShared) SetSchemas(v []EnumtimeLimitLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRetainDuration

`func (o *TimeLimitLogRetentionPolicyShared) GetRetainDuration() string`

GetRetainDuration returns the RetainDuration field if non-nil, zero value otherwise.

### GetRetainDurationOk

`func (o *TimeLimitLogRetentionPolicyShared) GetRetainDurationOk() (*string, bool)`

GetRetainDurationOk returns a tuple with the RetainDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainDuration

`func (o *TimeLimitLogRetentionPolicyShared) SetRetainDuration(v string)`

SetRetainDuration sets RetainDuration field to given value.


### GetDescription

`func (o *TimeLimitLogRetentionPolicyShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TimeLimitLogRetentionPolicyShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TimeLimitLogRetentionPolicyShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TimeLimitLogRetentionPolicyShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


