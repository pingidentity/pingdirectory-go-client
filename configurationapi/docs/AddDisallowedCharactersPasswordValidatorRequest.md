# AddDisallowedCharactersPasswordValidatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdisallowedCharactersPasswordValidatorSchemaUrn**](EnumdisallowedCharactersPasswordValidatorSchemaUrn.md) |  | 
**DisallowedCharacters** | Pointer to **string** | A set of characters that will not be allowed anywhere in a password. | [optional] 
**DisallowedLeadingCharacters** | Pointer to **string** | A set of characters that will not be allowed as the first character of the password. | [optional] 
**DisallowedTrailingCharacters** | Pointer to **string** | A set of characters that will not be allowed as the last character of the password. | [optional] 
**Description** | Pointer to **string** | A description for this Password Validator | [optional] 
**Enabled** | **bool** | Indicates whether the password validator is enabled for use. | 
**ValidatorRequirementDescription** | Pointer to **string** | Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator. | [optional] 
**ValidatorFailureMessage** | Pointer to **string** | Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator. | [optional] 
**ValidatorName** | **string** | Name of the new Password Validator | 

## Methods

### NewAddDisallowedCharactersPasswordValidatorRequest

`func NewAddDisallowedCharactersPasswordValidatorRequest(schemas []EnumdisallowedCharactersPasswordValidatorSchemaUrn, enabled bool, validatorName string, ) *AddDisallowedCharactersPasswordValidatorRequest`

NewAddDisallowedCharactersPasswordValidatorRequest instantiates a new AddDisallowedCharactersPasswordValidatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDisallowedCharactersPasswordValidatorRequestWithDefaults

`func NewAddDisallowedCharactersPasswordValidatorRequestWithDefaults() *AddDisallowedCharactersPasswordValidatorRequest`

NewAddDisallowedCharactersPasswordValidatorRequestWithDefaults instantiates a new AddDisallowedCharactersPasswordValidatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetSchemas() []EnumdisallowedCharactersPasswordValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetSchemasOk() (*[]EnumdisallowedCharactersPasswordValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDisallowedCharactersPasswordValidatorRequest) SetSchemas(v []EnumdisallowedCharactersPasswordValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDisallowedCharacters

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetDisallowedCharacters() string`

GetDisallowedCharacters returns the DisallowedCharacters field if non-nil, zero value otherwise.

### GetDisallowedCharactersOk

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetDisallowedCharactersOk() (*string, bool)`

GetDisallowedCharactersOk returns a tuple with the DisallowedCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowedCharacters

`func (o *AddDisallowedCharactersPasswordValidatorRequest) SetDisallowedCharacters(v string)`

SetDisallowedCharacters sets DisallowedCharacters field to given value.

### HasDisallowedCharacters

`func (o *AddDisallowedCharactersPasswordValidatorRequest) HasDisallowedCharacters() bool`

HasDisallowedCharacters returns a boolean if a field has been set.

### GetDisallowedLeadingCharacters

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetDisallowedLeadingCharacters() string`

GetDisallowedLeadingCharacters returns the DisallowedLeadingCharacters field if non-nil, zero value otherwise.

### GetDisallowedLeadingCharactersOk

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetDisallowedLeadingCharactersOk() (*string, bool)`

GetDisallowedLeadingCharactersOk returns a tuple with the DisallowedLeadingCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowedLeadingCharacters

`func (o *AddDisallowedCharactersPasswordValidatorRequest) SetDisallowedLeadingCharacters(v string)`

SetDisallowedLeadingCharacters sets DisallowedLeadingCharacters field to given value.

### HasDisallowedLeadingCharacters

`func (o *AddDisallowedCharactersPasswordValidatorRequest) HasDisallowedLeadingCharacters() bool`

HasDisallowedLeadingCharacters returns a boolean if a field has been set.

### GetDisallowedTrailingCharacters

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetDisallowedTrailingCharacters() string`

GetDisallowedTrailingCharacters returns the DisallowedTrailingCharacters field if non-nil, zero value otherwise.

### GetDisallowedTrailingCharactersOk

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetDisallowedTrailingCharactersOk() (*string, bool)`

GetDisallowedTrailingCharactersOk returns a tuple with the DisallowedTrailingCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowedTrailingCharacters

`func (o *AddDisallowedCharactersPasswordValidatorRequest) SetDisallowedTrailingCharacters(v string)`

SetDisallowedTrailingCharacters sets DisallowedTrailingCharacters field to given value.

### HasDisallowedTrailingCharacters

`func (o *AddDisallowedCharactersPasswordValidatorRequest) HasDisallowedTrailingCharacters() bool`

HasDisallowedTrailingCharacters returns a boolean if a field has been set.

### GetDescription

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDisallowedCharactersPasswordValidatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDisallowedCharactersPasswordValidatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddDisallowedCharactersPasswordValidatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetValidatorRequirementDescription

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetValidatorRequirementDescription() string`

GetValidatorRequirementDescription returns the ValidatorRequirementDescription field if non-nil, zero value otherwise.

### GetValidatorRequirementDescriptionOk

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool)`

GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorRequirementDescription

`func (o *AddDisallowedCharactersPasswordValidatorRequest) SetValidatorRequirementDescription(v string)`

SetValidatorRequirementDescription sets ValidatorRequirementDescription field to given value.

### HasValidatorRequirementDescription

`func (o *AddDisallowedCharactersPasswordValidatorRequest) HasValidatorRequirementDescription() bool`

HasValidatorRequirementDescription returns a boolean if a field has been set.

### GetValidatorFailureMessage

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetValidatorFailureMessage() string`

GetValidatorFailureMessage returns the ValidatorFailureMessage field if non-nil, zero value otherwise.

### GetValidatorFailureMessageOk

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool)`

GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorFailureMessage

`func (o *AddDisallowedCharactersPasswordValidatorRequest) SetValidatorFailureMessage(v string)`

SetValidatorFailureMessage sets ValidatorFailureMessage field to given value.

### HasValidatorFailureMessage

`func (o *AddDisallowedCharactersPasswordValidatorRequest) HasValidatorFailureMessage() bool`

HasValidatorFailureMessage returns a boolean if a field has been set.

### GetValidatorName

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetValidatorName() string`

GetValidatorName returns the ValidatorName field if non-nil, zero value otherwise.

### GetValidatorNameOk

`func (o *AddDisallowedCharactersPasswordValidatorRequest) GetValidatorNameOk() (*string, bool)`

GetValidatorNameOk returns a tuple with the ValidatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorName

`func (o *AddDisallowedCharactersPasswordValidatorRequest) SetValidatorName(v string)`

SetValidatorName sets ValidatorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


