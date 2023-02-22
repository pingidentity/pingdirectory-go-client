# AddLengthBasedPasswordValidatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidatorName** | **string** | Name of the new Password Validator | 
**Schemas** | [**[]EnumlengthBasedPasswordValidatorSchemaUrn**](EnumlengthBasedPasswordValidatorSchemaUrn.md) |  | 
**MaxPasswordLength** | Pointer to **int32** | Specifies the maximum number of characters that can be included in a proposed password. | [optional] 
**MinPasswordLength** | Pointer to **int32** | Specifies the minimum number of characters that must be included in a proposed password. | [optional] 
**Description** | Pointer to **string** | A description for this Password Validator | [optional] 
**Enabled** | **bool** | Indicates whether the password validator is enabled for use. | 
**ValidatorRequirementDescription** | Pointer to **string** | Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator. | [optional] 
**ValidatorFailureMessage** | Pointer to **string** | Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator. | [optional] 

## Methods

### NewAddLengthBasedPasswordValidatorRequest

`func NewAddLengthBasedPasswordValidatorRequest(validatorName string, schemas []EnumlengthBasedPasswordValidatorSchemaUrn, enabled bool, ) *AddLengthBasedPasswordValidatorRequest`

NewAddLengthBasedPasswordValidatorRequest instantiates a new AddLengthBasedPasswordValidatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLengthBasedPasswordValidatorRequestWithDefaults

`func NewAddLengthBasedPasswordValidatorRequestWithDefaults() *AddLengthBasedPasswordValidatorRequest`

NewAddLengthBasedPasswordValidatorRequestWithDefaults instantiates a new AddLengthBasedPasswordValidatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidatorName

`func (o *AddLengthBasedPasswordValidatorRequest) GetValidatorName() string`

GetValidatorName returns the ValidatorName field if non-nil, zero value otherwise.

### GetValidatorNameOk

`func (o *AddLengthBasedPasswordValidatorRequest) GetValidatorNameOk() (*string, bool)`

GetValidatorNameOk returns a tuple with the ValidatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorName

`func (o *AddLengthBasedPasswordValidatorRequest) SetValidatorName(v string)`

SetValidatorName sets ValidatorName field to given value.


### GetSchemas

`func (o *AddLengthBasedPasswordValidatorRequest) GetSchemas() []EnumlengthBasedPasswordValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLengthBasedPasswordValidatorRequest) GetSchemasOk() (*[]EnumlengthBasedPasswordValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLengthBasedPasswordValidatorRequest) SetSchemas(v []EnumlengthBasedPasswordValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMaxPasswordLength

`func (o *AddLengthBasedPasswordValidatorRequest) GetMaxPasswordLength() int32`

GetMaxPasswordLength returns the MaxPasswordLength field if non-nil, zero value otherwise.

### GetMaxPasswordLengthOk

`func (o *AddLengthBasedPasswordValidatorRequest) GetMaxPasswordLengthOk() (*int32, bool)`

GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordLength

`func (o *AddLengthBasedPasswordValidatorRequest) SetMaxPasswordLength(v int32)`

SetMaxPasswordLength sets MaxPasswordLength field to given value.

### HasMaxPasswordLength

`func (o *AddLengthBasedPasswordValidatorRequest) HasMaxPasswordLength() bool`

HasMaxPasswordLength returns a boolean if a field has been set.

### GetMinPasswordLength

`func (o *AddLengthBasedPasswordValidatorRequest) GetMinPasswordLength() int32`

GetMinPasswordLength returns the MinPasswordLength field if non-nil, zero value otherwise.

### GetMinPasswordLengthOk

`func (o *AddLengthBasedPasswordValidatorRequest) GetMinPasswordLengthOk() (*int32, bool)`

GetMinPasswordLengthOk returns a tuple with the MinPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPasswordLength

`func (o *AddLengthBasedPasswordValidatorRequest) SetMinPasswordLength(v int32)`

SetMinPasswordLength sets MinPasswordLength field to given value.

### HasMinPasswordLength

`func (o *AddLengthBasedPasswordValidatorRequest) HasMinPasswordLength() bool`

HasMinPasswordLength returns a boolean if a field has been set.

### GetDescription

`func (o *AddLengthBasedPasswordValidatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLengthBasedPasswordValidatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLengthBasedPasswordValidatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLengthBasedPasswordValidatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddLengthBasedPasswordValidatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddLengthBasedPasswordValidatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddLengthBasedPasswordValidatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetValidatorRequirementDescription

`func (o *AddLengthBasedPasswordValidatorRequest) GetValidatorRequirementDescription() string`

GetValidatorRequirementDescription returns the ValidatorRequirementDescription field if non-nil, zero value otherwise.

### GetValidatorRequirementDescriptionOk

`func (o *AddLengthBasedPasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool)`

GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorRequirementDescription

`func (o *AddLengthBasedPasswordValidatorRequest) SetValidatorRequirementDescription(v string)`

SetValidatorRequirementDescription sets ValidatorRequirementDescription field to given value.

### HasValidatorRequirementDescription

`func (o *AddLengthBasedPasswordValidatorRequest) HasValidatorRequirementDescription() bool`

HasValidatorRequirementDescription returns a boolean if a field has been set.

### GetValidatorFailureMessage

`func (o *AddLengthBasedPasswordValidatorRequest) GetValidatorFailureMessage() string`

GetValidatorFailureMessage returns the ValidatorFailureMessage field if non-nil, zero value otherwise.

### GetValidatorFailureMessageOk

`func (o *AddLengthBasedPasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool)`

GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorFailureMessage

`func (o *AddLengthBasedPasswordValidatorRequest) SetValidatorFailureMessage(v string)`

SetValidatorFailureMessage sets ValidatorFailureMessage field to given value.

### HasValidatorFailureMessage

`func (o *AddLengthBasedPasswordValidatorRequest) HasValidatorFailureMessage() bool`

HasValidatorFailureMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

