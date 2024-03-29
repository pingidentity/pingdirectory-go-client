# HaystackPasswordValidatorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumhaystackPasswordValidatorSchemaUrn**](EnumhaystackPasswordValidatorSchemaUrn.md) |  | 
**AssumedPasswordGuessesPerSecond** | **string** | The number of password guesses per second that a potential attacker may be expected to make. | 
**MinimumAcceptableTimeToExhaustSearchSpace** | **string** | The minimum length of time (using the configured number of password guesses per second) required to exhaust the entire search space for a proposed password in order for that password to be considered acceptable. | 
**Description** | Pointer to **string** | A description for this Password Validator | [optional] 
**Enabled** | **bool** | Indicates whether the password validator is enabled for use. | 
**ValidatorRequirementDescription** | Pointer to **string** | Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator. | [optional] 
**ValidatorFailureMessage** | Pointer to **string** | Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Password Validator | 

## Methods

### NewHaystackPasswordValidatorResponse

`func NewHaystackPasswordValidatorResponse(schemas []EnumhaystackPasswordValidatorSchemaUrn, assumedPasswordGuessesPerSecond string, minimumAcceptableTimeToExhaustSearchSpace string, enabled bool, id string, ) *HaystackPasswordValidatorResponse`

NewHaystackPasswordValidatorResponse instantiates a new HaystackPasswordValidatorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHaystackPasswordValidatorResponseWithDefaults

`func NewHaystackPasswordValidatorResponseWithDefaults() *HaystackPasswordValidatorResponse`

NewHaystackPasswordValidatorResponseWithDefaults instantiates a new HaystackPasswordValidatorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *HaystackPasswordValidatorResponse) GetSchemas() []EnumhaystackPasswordValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *HaystackPasswordValidatorResponse) GetSchemasOk() (*[]EnumhaystackPasswordValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *HaystackPasswordValidatorResponse) SetSchemas(v []EnumhaystackPasswordValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAssumedPasswordGuessesPerSecond

`func (o *HaystackPasswordValidatorResponse) GetAssumedPasswordGuessesPerSecond() string`

GetAssumedPasswordGuessesPerSecond returns the AssumedPasswordGuessesPerSecond field if non-nil, zero value otherwise.

### GetAssumedPasswordGuessesPerSecondOk

`func (o *HaystackPasswordValidatorResponse) GetAssumedPasswordGuessesPerSecondOk() (*string, bool)`

GetAssumedPasswordGuessesPerSecondOk returns a tuple with the AssumedPasswordGuessesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumedPasswordGuessesPerSecond

`func (o *HaystackPasswordValidatorResponse) SetAssumedPasswordGuessesPerSecond(v string)`

SetAssumedPasswordGuessesPerSecond sets AssumedPasswordGuessesPerSecond field to given value.


### GetMinimumAcceptableTimeToExhaustSearchSpace

`func (o *HaystackPasswordValidatorResponse) GetMinimumAcceptableTimeToExhaustSearchSpace() string`

GetMinimumAcceptableTimeToExhaustSearchSpace returns the MinimumAcceptableTimeToExhaustSearchSpace field if non-nil, zero value otherwise.

### GetMinimumAcceptableTimeToExhaustSearchSpaceOk

`func (o *HaystackPasswordValidatorResponse) GetMinimumAcceptableTimeToExhaustSearchSpaceOk() (*string, bool)`

GetMinimumAcceptableTimeToExhaustSearchSpaceOk returns a tuple with the MinimumAcceptableTimeToExhaustSearchSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAcceptableTimeToExhaustSearchSpace

`func (o *HaystackPasswordValidatorResponse) SetMinimumAcceptableTimeToExhaustSearchSpace(v string)`

SetMinimumAcceptableTimeToExhaustSearchSpace sets MinimumAcceptableTimeToExhaustSearchSpace field to given value.


### GetDescription

`func (o *HaystackPasswordValidatorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HaystackPasswordValidatorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HaystackPasswordValidatorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HaystackPasswordValidatorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *HaystackPasswordValidatorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HaystackPasswordValidatorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HaystackPasswordValidatorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetValidatorRequirementDescription

`func (o *HaystackPasswordValidatorResponse) GetValidatorRequirementDescription() string`

GetValidatorRequirementDescription returns the ValidatorRequirementDescription field if non-nil, zero value otherwise.

### GetValidatorRequirementDescriptionOk

`func (o *HaystackPasswordValidatorResponse) GetValidatorRequirementDescriptionOk() (*string, bool)`

GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorRequirementDescription

`func (o *HaystackPasswordValidatorResponse) SetValidatorRequirementDescription(v string)`

SetValidatorRequirementDescription sets ValidatorRequirementDescription field to given value.

### HasValidatorRequirementDescription

`func (o *HaystackPasswordValidatorResponse) HasValidatorRequirementDescription() bool`

HasValidatorRequirementDescription returns a boolean if a field has been set.

### GetValidatorFailureMessage

`func (o *HaystackPasswordValidatorResponse) GetValidatorFailureMessage() string`

GetValidatorFailureMessage returns the ValidatorFailureMessage field if non-nil, zero value otherwise.

### GetValidatorFailureMessageOk

`func (o *HaystackPasswordValidatorResponse) GetValidatorFailureMessageOk() (*string, bool)`

GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorFailureMessage

`func (o *HaystackPasswordValidatorResponse) SetValidatorFailureMessage(v string)`

SetValidatorFailureMessage sets ValidatorFailureMessage field to given value.

### HasValidatorFailureMessage

`func (o *HaystackPasswordValidatorResponse) HasValidatorFailureMessage() bool`

HasValidatorFailureMessage returns a boolean if a field has been set.

### GetMeta

`func (o *HaystackPasswordValidatorResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HaystackPasswordValidatorResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HaystackPasswordValidatorResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HaystackPasswordValidatorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *HaystackPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *HaystackPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *HaystackPasswordValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *HaystackPasswordValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *HaystackPasswordValidatorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HaystackPasswordValidatorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HaystackPasswordValidatorResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


