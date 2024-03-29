# AddUtf8PasswordValidatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumutf8PasswordValidatorSchemaUrn**](Enumutf8PasswordValidatorSchemaUrn.md) |  | 
**AllowNonAsciiCharacters** | Pointer to **bool** | Indicates whether passwords will be allowed to include characters from outside the ASCII character set. | [optional] 
**AllowUnknownCharacters** | Pointer to **bool** | Indicates whether passwords will be allowed to include characters that are not recognized by the JVM&#39;s Unicode support. | [optional] 
**AllowedCharacterType** | Pointer to [**[]EnumpasswordValidatorAllowedCharacterTypeProp**](EnumpasswordValidatorAllowedCharacterTypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Password Validator | [optional] 
**Enabled** | **bool** | Indicates whether the password validator is enabled for use. | 
**ValidatorRequirementDescription** | Pointer to **string** | Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator. | [optional] 
**ValidatorFailureMessage** | Pointer to **string** | Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator. | [optional] 
**ValidatorName** | **string** | Name of the new Password Validator | 

## Methods

### NewAddUtf8PasswordValidatorRequest

`func NewAddUtf8PasswordValidatorRequest(schemas []Enumutf8PasswordValidatorSchemaUrn, enabled bool, validatorName string, ) *AddUtf8PasswordValidatorRequest`

NewAddUtf8PasswordValidatorRequest instantiates a new AddUtf8PasswordValidatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUtf8PasswordValidatorRequestWithDefaults

`func NewAddUtf8PasswordValidatorRequestWithDefaults() *AddUtf8PasswordValidatorRequest`

NewAddUtf8PasswordValidatorRequestWithDefaults instantiates a new AddUtf8PasswordValidatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddUtf8PasswordValidatorRequest) GetSchemas() []Enumutf8PasswordValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddUtf8PasswordValidatorRequest) GetSchemasOk() (*[]Enumutf8PasswordValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddUtf8PasswordValidatorRequest) SetSchemas(v []Enumutf8PasswordValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowNonAsciiCharacters

`func (o *AddUtf8PasswordValidatorRequest) GetAllowNonAsciiCharacters() bool`

GetAllowNonAsciiCharacters returns the AllowNonAsciiCharacters field if non-nil, zero value otherwise.

### GetAllowNonAsciiCharactersOk

`func (o *AddUtf8PasswordValidatorRequest) GetAllowNonAsciiCharactersOk() (*bool, bool)`

GetAllowNonAsciiCharactersOk returns a tuple with the AllowNonAsciiCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNonAsciiCharacters

`func (o *AddUtf8PasswordValidatorRequest) SetAllowNonAsciiCharacters(v bool)`

SetAllowNonAsciiCharacters sets AllowNonAsciiCharacters field to given value.

### HasAllowNonAsciiCharacters

`func (o *AddUtf8PasswordValidatorRequest) HasAllowNonAsciiCharacters() bool`

HasAllowNonAsciiCharacters returns a boolean if a field has been set.

### GetAllowUnknownCharacters

`func (o *AddUtf8PasswordValidatorRequest) GetAllowUnknownCharacters() bool`

GetAllowUnknownCharacters returns the AllowUnknownCharacters field if non-nil, zero value otherwise.

### GetAllowUnknownCharactersOk

`func (o *AddUtf8PasswordValidatorRequest) GetAllowUnknownCharactersOk() (*bool, bool)`

GetAllowUnknownCharactersOk returns a tuple with the AllowUnknownCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnknownCharacters

`func (o *AddUtf8PasswordValidatorRequest) SetAllowUnknownCharacters(v bool)`

SetAllowUnknownCharacters sets AllowUnknownCharacters field to given value.

### HasAllowUnknownCharacters

`func (o *AddUtf8PasswordValidatorRequest) HasAllowUnknownCharacters() bool`

HasAllowUnknownCharacters returns a boolean if a field has been set.

### GetAllowedCharacterType

`func (o *AddUtf8PasswordValidatorRequest) GetAllowedCharacterType() []EnumpasswordValidatorAllowedCharacterTypeProp`

GetAllowedCharacterType returns the AllowedCharacterType field if non-nil, zero value otherwise.

### GetAllowedCharacterTypeOk

`func (o *AddUtf8PasswordValidatorRequest) GetAllowedCharacterTypeOk() (*[]EnumpasswordValidatorAllowedCharacterTypeProp, bool)`

GetAllowedCharacterTypeOk returns a tuple with the AllowedCharacterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCharacterType

`func (o *AddUtf8PasswordValidatorRequest) SetAllowedCharacterType(v []EnumpasswordValidatorAllowedCharacterTypeProp)`

SetAllowedCharacterType sets AllowedCharacterType field to given value.

### HasAllowedCharacterType

`func (o *AddUtf8PasswordValidatorRequest) HasAllowedCharacterType() bool`

HasAllowedCharacterType returns a boolean if a field has been set.

### GetDescription

`func (o *AddUtf8PasswordValidatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUtf8PasswordValidatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUtf8PasswordValidatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUtf8PasswordValidatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddUtf8PasswordValidatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddUtf8PasswordValidatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddUtf8PasswordValidatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetValidatorRequirementDescription

`func (o *AddUtf8PasswordValidatorRequest) GetValidatorRequirementDescription() string`

GetValidatorRequirementDescription returns the ValidatorRequirementDescription field if non-nil, zero value otherwise.

### GetValidatorRequirementDescriptionOk

`func (o *AddUtf8PasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool)`

GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorRequirementDescription

`func (o *AddUtf8PasswordValidatorRequest) SetValidatorRequirementDescription(v string)`

SetValidatorRequirementDescription sets ValidatorRequirementDescription field to given value.

### HasValidatorRequirementDescription

`func (o *AddUtf8PasswordValidatorRequest) HasValidatorRequirementDescription() bool`

HasValidatorRequirementDescription returns a boolean if a field has been set.

### GetValidatorFailureMessage

`func (o *AddUtf8PasswordValidatorRequest) GetValidatorFailureMessage() string`

GetValidatorFailureMessage returns the ValidatorFailureMessage field if non-nil, zero value otherwise.

### GetValidatorFailureMessageOk

`func (o *AddUtf8PasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool)`

GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorFailureMessage

`func (o *AddUtf8PasswordValidatorRequest) SetValidatorFailureMessage(v string)`

SetValidatorFailureMessage sets ValidatorFailureMessage field to given value.

### HasValidatorFailureMessage

`func (o *AddUtf8PasswordValidatorRequest) HasValidatorFailureMessage() bool`

HasValidatorFailureMessage returns a boolean if a field has been set.

### GetValidatorName

`func (o *AddUtf8PasswordValidatorRequest) GetValidatorName() string`

GetValidatorName returns the ValidatorName field if non-nil, zero value otherwise.

### GetValidatorNameOk

`func (o *AddUtf8PasswordValidatorRequest) GetValidatorNameOk() (*string, bool)`

GetValidatorNameOk returns a tuple with the ValidatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorName

`func (o *AddUtf8PasswordValidatorRequest) SetValidatorName(v string)`

SetValidatorName sets ValidatorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


