# AddSimilarityBasedPasswordValidatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidatorName** | **string** | Name of the new Password Validator | 
**Schemas** | [**[]EnumsimilarityBasedPasswordValidatorSchemaUrn**](EnumsimilarityBasedPasswordValidatorSchemaUrn.md) |  | 
**MinPasswordDifference** | **int32** | Specifies the minimum difference of new and old password. | 
**Description** | Pointer to **string** | A description for this Password Validator | [optional] 
**Enabled** | **bool** | Indicates whether the password validator is enabled for use. | 
**ValidatorRequirementDescription** | Pointer to **string** | Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator. | [optional] 
**ValidatorFailureMessage** | Pointer to **string** | Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator. | [optional] 

## Methods

### NewAddSimilarityBasedPasswordValidatorRequest

`func NewAddSimilarityBasedPasswordValidatorRequest(validatorName string, schemas []EnumsimilarityBasedPasswordValidatorSchemaUrn, minPasswordDifference int32, enabled bool, ) *AddSimilarityBasedPasswordValidatorRequest`

NewAddSimilarityBasedPasswordValidatorRequest instantiates a new AddSimilarityBasedPasswordValidatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSimilarityBasedPasswordValidatorRequestWithDefaults

`func NewAddSimilarityBasedPasswordValidatorRequestWithDefaults() *AddSimilarityBasedPasswordValidatorRequest`

NewAddSimilarityBasedPasswordValidatorRequestWithDefaults instantiates a new AddSimilarityBasedPasswordValidatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidatorName

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorName() string`

GetValidatorName returns the ValidatorName field if non-nil, zero value otherwise.

### GetValidatorNameOk

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorNameOk() (*string, bool)`

GetValidatorNameOk returns a tuple with the ValidatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorName

`func (o *AddSimilarityBasedPasswordValidatorRequest) SetValidatorName(v string)`

SetValidatorName sets ValidatorName field to given value.


### GetSchemas

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetSchemas() []EnumsimilarityBasedPasswordValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetSchemasOk() (*[]EnumsimilarityBasedPasswordValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSimilarityBasedPasswordValidatorRequest) SetSchemas(v []EnumsimilarityBasedPasswordValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetMinPasswordDifference

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetMinPasswordDifference() int32`

GetMinPasswordDifference returns the MinPasswordDifference field if non-nil, zero value otherwise.

### GetMinPasswordDifferenceOk

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetMinPasswordDifferenceOk() (*int32, bool)`

GetMinPasswordDifferenceOk returns a tuple with the MinPasswordDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPasswordDifference

`func (o *AddSimilarityBasedPasswordValidatorRequest) SetMinPasswordDifference(v int32)`

SetMinPasswordDifference sets MinPasswordDifference field to given value.


### GetDescription

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSimilarityBasedPasswordValidatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSimilarityBasedPasswordValidatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSimilarityBasedPasswordValidatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetValidatorRequirementDescription

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorRequirementDescription() string`

GetValidatorRequirementDescription returns the ValidatorRequirementDescription field if non-nil, zero value otherwise.

### GetValidatorRequirementDescriptionOk

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool)`

GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorRequirementDescription

`func (o *AddSimilarityBasedPasswordValidatorRequest) SetValidatorRequirementDescription(v string)`

SetValidatorRequirementDescription sets ValidatorRequirementDescription field to given value.

### HasValidatorRequirementDescription

`func (o *AddSimilarityBasedPasswordValidatorRequest) HasValidatorRequirementDescription() bool`

HasValidatorRequirementDescription returns a boolean if a field has been set.

### GetValidatorFailureMessage

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorFailureMessage() string`

GetValidatorFailureMessage returns the ValidatorFailureMessage field if non-nil, zero value otherwise.

### GetValidatorFailureMessageOk

`func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool)`

GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorFailureMessage

`func (o *AddSimilarityBasedPasswordValidatorRequest) SetValidatorFailureMessage(v string)`

SetValidatorFailureMessage sets ValidatorFailureMessage field to given value.

### HasValidatorFailureMessage

`func (o *AddSimilarityBasedPasswordValidatorRequest) HasValidatorFailureMessage() bool`

HasValidatorFailureMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

