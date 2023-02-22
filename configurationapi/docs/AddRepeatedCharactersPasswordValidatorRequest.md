# AddRepeatedCharactersPasswordValidatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidatorName** | **string** | Name of the new Password Validator | 
**Schemas** | [**[]EnumrepeatedCharactersPasswordValidatorSchemaUrn**](EnumrepeatedCharactersPasswordValidatorSchemaUrn.md) |  | 
**MaxConsecutiveLength** | **int32** | Specifies the maximum number of times that any character can appear consecutively in a password value. | 
**CaseSensitiveValidation** | **bool** | Indicates whether this password validator should treat password characters in a case-sensitive manner. | 
**CharacterSet** | Pointer to **[]string** | Specifies a set of characters that should be considered equivalent for the purpose of this password validator. This can be used, for example, to ensure that passwords contain no more than three consecutive digits. | [optional] 
**Description** | Pointer to **string** | A description for this Password Validator | [optional] 
**Enabled** | **bool** | Indicates whether the password validator is enabled for use. | 
**ValidatorRequirementDescription** | Pointer to **string** | Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator. | [optional] 
**ValidatorFailureMessage** | Pointer to **string** | Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator. | [optional] 

## Methods

### NewAddRepeatedCharactersPasswordValidatorRequest

`func NewAddRepeatedCharactersPasswordValidatorRequest(validatorName string, schemas []EnumrepeatedCharactersPasswordValidatorSchemaUrn, maxConsecutiveLength int32, caseSensitiveValidation bool, enabled bool, ) *AddRepeatedCharactersPasswordValidatorRequest`

NewAddRepeatedCharactersPasswordValidatorRequest instantiates a new AddRepeatedCharactersPasswordValidatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRepeatedCharactersPasswordValidatorRequestWithDefaults

`func NewAddRepeatedCharactersPasswordValidatorRequestWithDefaults() *AddRepeatedCharactersPasswordValidatorRequest`

NewAddRepeatedCharactersPasswordValidatorRequestWithDefaults instantiates a new AddRepeatedCharactersPasswordValidatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidatorName

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetValidatorName() string`

GetValidatorName returns the ValidatorName field if non-nil, zero value otherwise.

### GetValidatorNameOk

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetValidatorNameOk() (*string, bool)`

GetValidatorNameOk returns a tuple with the ValidatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorName

`func (o *AddRepeatedCharactersPasswordValidatorRequest) SetValidatorName(v string)`

SetValidatorName sets ValidatorName field to given value.


### GetSchemas

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetSchemas() []EnumrepeatedCharactersPasswordValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetSchemasOk() (*[]EnumrepeatedCharactersPasswordValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddRepeatedCharactersPasswordValidatorRequest) SetSchemas(v []EnumrepeatedCharactersPasswordValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMaxConsecutiveLength

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetMaxConsecutiveLength() int32`

GetMaxConsecutiveLength returns the MaxConsecutiveLength field if non-nil, zero value otherwise.

### GetMaxConsecutiveLengthOk

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetMaxConsecutiveLengthOk() (*int32, bool)`

GetMaxConsecutiveLengthOk returns a tuple with the MaxConsecutiveLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConsecutiveLength

`func (o *AddRepeatedCharactersPasswordValidatorRequest) SetMaxConsecutiveLength(v int32)`

SetMaxConsecutiveLength sets MaxConsecutiveLength field to given value.


### GetCaseSensitiveValidation

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetCaseSensitiveValidation() bool`

GetCaseSensitiveValidation returns the CaseSensitiveValidation field if non-nil, zero value otherwise.

### GetCaseSensitiveValidationOk

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetCaseSensitiveValidationOk() (*bool, bool)`

GetCaseSensitiveValidationOk returns a tuple with the CaseSensitiveValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitiveValidation

`func (o *AddRepeatedCharactersPasswordValidatorRequest) SetCaseSensitiveValidation(v bool)`

SetCaseSensitiveValidation sets CaseSensitiveValidation field to given value.


### GetCharacterSet

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetCharacterSet() []string`

GetCharacterSet returns the CharacterSet field if non-nil, zero value otherwise.

### GetCharacterSetOk

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetCharacterSetOk() (*[]string, bool)`

GetCharacterSetOk returns a tuple with the CharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterSet

`func (o *AddRepeatedCharactersPasswordValidatorRequest) SetCharacterSet(v []string)`

SetCharacterSet sets CharacterSet field to given value.

### HasCharacterSet

`func (o *AddRepeatedCharactersPasswordValidatorRequest) HasCharacterSet() bool`

HasCharacterSet returns a boolean if a field has been set.

### GetDescription

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddRepeatedCharactersPasswordValidatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddRepeatedCharactersPasswordValidatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddRepeatedCharactersPasswordValidatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetValidatorRequirementDescription

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetValidatorRequirementDescription() string`

GetValidatorRequirementDescription returns the ValidatorRequirementDescription field if non-nil, zero value otherwise.

### GetValidatorRequirementDescriptionOk

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool)`

GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorRequirementDescription

`func (o *AddRepeatedCharactersPasswordValidatorRequest) SetValidatorRequirementDescription(v string)`

SetValidatorRequirementDescription sets ValidatorRequirementDescription field to given value.

### HasValidatorRequirementDescription

`func (o *AddRepeatedCharactersPasswordValidatorRequest) HasValidatorRequirementDescription() bool`

HasValidatorRequirementDescription returns a boolean if a field has been set.

### GetValidatorFailureMessage

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetValidatorFailureMessage() string`

GetValidatorFailureMessage returns the ValidatorFailureMessage field if non-nil, zero value otherwise.

### GetValidatorFailureMessageOk

`func (o *AddRepeatedCharactersPasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool)`

GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorFailureMessage

`func (o *AddRepeatedCharactersPasswordValidatorRequest) SetValidatorFailureMessage(v string)`

SetValidatorFailureMessage sets ValidatorFailureMessage field to given value.

### HasValidatorFailureMessage

`func (o *AddRepeatedCharactersPasswordValidatorRequest) HasValidatorFailureMessage() bool`

HasValidatorFailureMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

