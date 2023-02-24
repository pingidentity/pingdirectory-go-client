/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// AddDictionaryPasswordValidatorRequest struct for AddDictionaryPasswordValidatorRequest
type AddDictionaryPasswordValidatorRequest struct {
	// Name of the new Password Validator
	ValidatorName string                                     `json:"validatorName"`
	Schemas       []EnumdictionaryPasswordValidatorSchemaUrn `json:"schemas"`
	// Specifies the path to the file containing a list of words that cannot be used as passwords.
	DictionaryFile *string `json:"dictionaryFile,omitempty"`
	// Indicates whether this password validator is to treat password characters in a case-sensitive manner.
	CaseSensitiveValidation *bool `json:"caseSensitiveValidation,omitempty"`
	// Indicates whether this password validator is to test the reversed value of the provided password as well as the order in which it was given.
	TestReversedPassword *bool `json:"testReversedPassword,omitempty"`
	// Indicates whether to ignore any digits, symbols, or other non-alphabetic characters that may appear at the beginning of a proposed password.
	IgnoreLeadingNonAlphabeticCharacters *bool `json:"ignoreLeadingNonAlphabeticCharacters,omitempty"`
	// Indicates whether to ignore any digits, symbols, or other non-alphabetic characters that may appear at the end of a proposed password.
	IgnoreTrailingNonAlphabeticCharacters *bool `json:"ignoreTrailingNonAlphabeticCharacters,omitempty"`
	// Indicates whether to strip characters of any diacritical marks (like accents, cedillas, circumflexes, diaereses, tildes, and umlauts) they may contain. Any characters with a diacritical mark would be replaced with a base version
	StripDiacriticalMarks *bool `json:"stripDiacriticalMarks,omitempty"`
	// Provides a set of character substitutions that can be applied to the proposed password when checking to see if it is in the provided dictionary. Each mapping should consist of a single character followed by a colon and a list of the alternative characters that may be used in place of that character.
	AlternativePasswordCharacterMapping []string `json:"alternativePasswordCharacterMapping,omitempty"`
	// The maximum allowed percent of a proposed password that any single dictionary word is allowed to comprise. A value of 100 indicates that a proposed password will only be rejected if the dictionary contains the entire proposed password (after any configured transformations have been applied).
	MaximumAllowedPercentOfPassword *int32 `json:"maximumAllowedPercentOfPassword,omitempty"`
	// A description for this Password Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether the password validator is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator.
	ValidatorRequirementDescription *string `json:"validatorRequirementDescription,omitempty"`
	// Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator.
	ValidatorFailureMessage *string `json:"validatorFailureMessage,omitempty"`
}

// NewAddDictionaryPasswordValidatorRequest instantiates a new AddDictionaryPasswordValidatorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDictionaryPasswordValidatorRequest(validatorName string, schemas []EnumdictionaryPasswordValidatorSchemaUrn, enabled bool) *AddDictionaryPasswordValidatorRequest {
	this := AddDictionaryPasswordValidatorRequest{}
	this.ValidatorName = validatorName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddDictionaryPasswordValidatorRequestWithDefaults instantiates a new AddDictionaryPasswordValidatorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDictionaryPasswordValidatorRequestWithDefaults() *AddDictionaryPasswordValidatorRequest {
	this := AddDictionaryPasswordValidatorRequest{}
	return &this
}

// GetValidatorName returns the ValidatorName field value
func (o *AddDictionaryPasswordValidatorRequest) GetValidatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorName
}

// GetValidatorNameOk returns a tuple with the ValidatorName field value
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetValidatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorName, true
}

// SetValidatorName sets field value
func (o *AddDictionaryPasswordValidatorRequest) SetValidatorName(v string) {
	o.ValidatorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddDictionaryPasswordValidatorRequest) GetSchemas() []EnumdictionaryPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumdictionaryPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetSchemasOk() ([]EnumdictionaryPasswordValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddDictionaryPasswordValidatorRequest) SetSchemas(v []EnumdictionaryPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetDictionaryFile returns the DictionaryFile field value if set, zero value otherwise.
func (o *AddDictionaryPasswordValidatorRequest) GetDictionaryFile() string {
	if o == nil || isNil(o.DictionaryFile) {
		var ret string
		return ret
	}
	return *o.DictionaryFile
}

// GetDictionaryFileOk returns a tuple with the DictionaryFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetDictionaryFileOk() (*string, bool) {
	if o == nil || isNil(o.DictionaryFile) {
		return nil, false
	}
	return o.DictionaryFile, true
}

// HasDictionaryFile returns a boolean if a field has been set.
func (o *AddDictionaryPasswordValidatorRequest) HasDictionaryFile() bool {
	if o != nil && !isNil(o.DictionaryFile) {
		return true
	}

	return false
}

// SetDictionaryFile gets a reference to the given string and assigns it to the DictionaryFile field.
func (o *AddDictionaryPasswordValidatorRequest) SetDictionaryFile(v string) {
	o.DictionaryFile = &v
}

// GetCaseSensitiveValidation returns the CaseSensitiveValidation field value if set, zero value otherwise.
func (o *AddDictionaryPasswordValidatorRequest) GetCaseSensitiveValidation() bool {
	if o == nil || isNil(o.CaseSensitiveValidation) {
		var ret bool
		return ret
	}
	return *o.CaseSensitiveValidation
}

// GetCaseSensitiveValidationOk returns a tuple with the CaseSensitiveValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetCaseSensitiveValidationOk() (*bool, bool) {
	if o == nil || isNil(o.CaseSensitiveValidation) {
		return nil, false
	}
	return o.CaseSensitiveValidation, true
}

// HasCaseSensitiveValidation returns a boolean if a field has been set.
func (o *AddDictionaryPasswordValidatorRequest) HasCaseSensitiveValidation() bool {
	if o != nil && !isNil(o.CaseSensitiveValidation) {
		return true
	}

	return false
}

// SetCaseSensitiveValidation gets a reference to the given bool and assigns it to the CaseSensitiveValidation field.
func (o *AddDictionaryPasswordValidatorRequest) SetCaseSensitiveValidation(v bool) {
	o.CaseSensitiveValidation = &v
}

// GetTestReversedPassword returns the TestReversedPassword field value if set, zero value otherwise.
func (o *AddDictionaryPasswordValidatorRequest) GetTestReversedPassword() bool {
	if o == nil || isNil(o.TestReversedPassword) {
		var ret bool
		return ret
	}
	return *o.TestReversedPassword
}

// GetTestReversedPasswordOk returns a tuple with the TestReversedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetTestReversedPasswordOk() (*bool, bool) {
	if o == nil || isNil(o.TestReversedPassword) {
		return nil, false
	}
	return o.TestReversedPassword, true
}

// HasTestReversedPassword returns a boolean if a field has been set.
func (o *AddDictionaryPasswordValidatorRequest) HasTestReversedPassword() bool {
	if o != nil && !isNil(o.TestReversedPassword) {
		return true
	}

	return false
}

// SetTestReversedPassword gets a reference to the given bool and assigns it to the TestReversedPassword field.
func (o *AddDictionaryPasswordValidatorRequest) SetTestReversedPassword(v bool) {
	o.TestReversedPassword = &v
}

// GetIgnoreLeadingNonAlphabeticCharacters returns the IgnoreLeadingNonAlphabeticCharacters field value if set, zero value otherwise.
func (o *AddDictionaryPasswordValidatorRequest) GetIgnoreLeadingNonAlphabeticCharacters() bool {
	if o == nil || isNil(o.IgnoreLeadingNonAlphabeticCharacters) {
		var ret bool
		return ret
	}
	return *o.IgnoreLeadingNonAlphabeticCharacters
}

// GetIgnoreLeadingNonAlphabeticCharactersOk returns a tuple with the IgnoreLeadingNonAlphabeticCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetIgnoreLeadingNonAlphabeticCharactersOk() (*bool, bool) {
	if o == nil || isNil(o.IgnoreLeadingNonAlphabeticCharacters) {
		return nil, false
	}
	return o.IgnoreLeadingNonAlphabeticCharacters, true
}

// HasIgnoreLeadingNonAlphabeticCharacters returns a boolean if a field has been set.
func (o *AddDictionaryPasswordValidatorRequest) HasIgnoreLeadingNonAlphabeticCharacters() bool {
	if o != nil && !isNil(o.IgnoreLeadingNonAlphabeticCharacters) {
		return true
	}

	return false
}

// SetIgnoreLeadingNonAlphabeticCharacters gets a reference to the given bool and assigns it to the IgnoreLeadingNonAlphabeticCharacters field.
func (o *AddDictionaryPasswordValidatorRequest) SetIgnoreLeadingNonAlphabeticCharacters(v bool) {
	o.IgnoreLeadingNonAlphabeticCharacters = &v
}

// GetIgnoreTrailingNonAlphabeticCharacters returns the IgnoreTrailingNonAlphabeticCharacters field value if set, zero value otherwise.
func (o *AddDictionaryPasswordValidatorRequest) GetIgnoreTrailingNonAlphabeticCharacters() bool {
	if o == nil || isNil(o.IgnoreTrailingNonAlphabeticCharacters) {
		var ret bool
		return ret
	}
	return *o.IgnoreTrailingNonAlphabeticCharacters
}

// GetIgnoreTrailingNonAlphabeticCharactersOk returns a tuple with the IgnoreTrailingNonAlphabeticCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetIgnoreTrailingNonAlphabeticCharactersOk() (*bool, bool) {
	if o == nil || isNil(o.IgnoreTrailingNonAlphabeticCharacters) {
		return nil, false
	}
	return o.IgnoreTrailingNonAlphabeticCharacters, true
}

// HasIgnoreTrailingNonAlphabeticCharacters returns a boolean if a field has been set.
func (o *AddDictionaryPasswordValidatorRequest) HasIgnoreTrailingNonAlphabeticCharacters() bool {
	if o != nil && !isNil(o.IgnoreTrailingNonAlphabeticCharacters) {
		return true
	}

	return false
}

// SetIgnoreTrailingNonAlphabeticCharacters gets a reference to the given bool and assigns it to the IgnoreTrailingNonAlphabeticCharacters field.
func (o *AddDictionaryPasswordValidatorRequest) SetIgnoreTrailingNonAlphabeticCharacters(v bool) {
	o.IgnoreTrailingNonAlphabeticCharacters = &v
}

// GetStripDiacriticalMarks returns the StripDiacriticalMarks field value if set, zero value otherwise.
func (o *AddDictionaryPasswordValidatorRequest) GetStripDiacriticalMarks() bool {
	if o == nil || isNil(o.StripDiacriticalMarks) {
		var ret bool
		return ret
	}
	return *o.StripDiacriticalMarks
}

// GetStripDiacriticalMarksOk returns a tuple with the StripDiacriticalMarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetStripDiacriticalMarksOk() (*bool, bool) {
	if o == nil || isNil(o.StripDiacriticalMarks) {
		return nil, false
	}
	return o.StripDiacriticalMarks, true
}

// HasStripDiacriticalMarks returns a boolean if a field has been set.
func (o *AddDictionaryPasswordValidatorRequest) HasStripDiacriticalMarks() bool {
	if o != nil && !isNil(o.StripDiacriticalMarks) {
		return true
	}

	return false
}

// SetStripDiacriticalMarks gets a reference to the given bool and assigns it to the StripDiacriticalMarks field.
func (o *AddDictionaryPasswordValidatorRequest) SetStripDiacriticalMarks(v bool) {
	o.StripDiacriticalMarks = &v
}

// GetAlternativePasswordCharacterMapping returns the AlternativePasswordCharacterMapping field value if set, zero value otherwise.
func (o *AddDictionaryPasswordValidatorRequest) GetAlternativePasswordCharacterMapping() []string {
	if o == nil || isNil(o.AlternativePasswordCharacterMapping) {
		var ret []string
		return ret
	}
	return o.AlternativePasswordCharacterMapping
}

// GetAlternativePasswordCharacterMappingOk returns a tuple with the AlternativePasswordCharacterMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetAlternativePasswordCharacterMappingOk() ([]string, bool) {
	if o == nil || isNil(o.AlternativePasswordCharacterMapping) {
		return nil, false
	}
	return o.AlternativePasswordCharacterMapping, true
}

// HasAlternativePasswordCharacterMapping returns a boolean if a field has been set.
func (o *AddDictionaryPasswordValidatorRequest) HasAlternativePasswordCharacterMapping() bool {
	if o != nil && !isNil(o.AlternativePasswordCharacterMapping) {
		return true
	}

	return false
}

// SetAlternativePasswordCharacterMapping gets a reference to the given []string and assigns it to the AlternativePasswordCharacterMapping field.
func (o *AddDictionaryPasswordValidatorRequest) SetAlternativePasswordCharacterMapping(v []string) {
	o.AlternativePasswordCharacterMapping = v
}

// GetMaximumAllowedPercentOfPassword returns the MaximumAllowedPercentOfPassword field value if set, zero value otherwise.
func (o *AddDictionaryPasswordValidatorRequest) GetMaximumAllowedPercentOfPassword() int32 {
	if o == nil || isNil(o.MaximumAllowedPercentOfPassword) {
		var ret int32
		return ret
	}
	return *o.MaximumAllowedPercentOfPassword
}

// GetMaximumAllowedPercentOfPasswordOk returns a tuple with the MaximumAllowedPercentOfPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetMaximumAllowedPercentOfPasswordOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumAllowedPercentOfPassword) {
		return nil, false
	}
	return o.MaximumAllowedPercentOfPassword, true
}

// HasMaximumAllowedPercentOfPassword returns a boolean if a field has been set.
func (o *AddDictionaryPasswordValidatorRequest) HasMaximumAllowedPercentOfPassword() bool {
	if o != nil && !isNil(o.MaximumAllowedPercentOfPassword) {
		return true
	}

	return false
}

// SetMaximumAllowedPercentOfPassword gets a reference to the given int32 and assigns it to the MaximumAllowedPercentOfPassword field.
func (o *AddDictionaryPasswordValidatorRequest) SetMaximumAllowedPercentOfPassword(v int32) {
	o.MaximumAllowedPercentOfPassword = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddDictionaryPasswordValidatorRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddDictionaryPasswordValidatorRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddDictionaryPasswordValidatorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddDictionaryPasswordValidatorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddDictionaryPasswordValidatorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *AddDictionaryPasswordValidatorRequest) GetValidatorRequirementDescription() string {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
		return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *AddDictionaryPasswordValidatorRequest) HasValidatorRequirementDescription() bool {
	if o != nil && !isNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *AddDictionaryPasswordValidatorRequest) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *AddDictionaryPasswordValidatorRequest) GetValidatorFailureMessage() string {
	if o == nil || isNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDictionaryPasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorFailureMessage) {
		return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *AddDictionaryPasswordValidatorRequest) HasValidatorFailureMessage() bool {
	if o != nil && !isNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *AddDictionaryPasswordValidatorRequest) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

func (o AddDictionaryPasswordValidatorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["validatorName"] = o.ValidatorName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.DictionaryFile) {
		toSerialize["dictionaryFile"] = o.DictionaryFile
	}
	if !isNil(o.CaseSensitiveValidation) {
		toSerialize["caseSensitiveValidation"] = o.CaseSensitiveValidation
	}
	if !isNil(o.TestReversedPassword) {
		toSerialize["testReversedPassword"] = o.TestReversedPassword
	}
	if !isNil(o.IgnoreLeadingNonAlphabeticCharacters) {
		toSerialize["ignoreLeadingNonAlphabeticCharacters"] = o.IgnoreLeadingNonAlphabeticCharacters
	}
	if !isNil(o.IgnoreTrailingNonAlphabeticCharacters) {
		toSerialize["ignoreTrailingNonAlphabeticCharacters"] = o.IgnoreTrailingNonAlphabeticCharacters
	}
	if !isNil(o.StripDiacriticalMarks) {
		toSerialize["stripDiacriticalMarks"] = o.StripDiacriticalMarks
	}
	if !isNil(o.AlternativePasswordCharacterMapping) {
		toSerialize["alternativePasswordCharacterMapping"] = o.AlternativePasswordCharacterMapping
	}
	if !isNil(o.MaximumAllowedPercentOfPassword) {
		toSerialize["maximumAllowedPercentOfPassword"] = o.MaximumAllowedPercentOfPassword
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.ValidatorRequirementDescription) {
		toSerialize["validatorRequirementDescription"] = o.ValidatorRequirementDescription
	}
	if !isNil(o.ValidatorFailureMessage) {
		toSerialize["validatorFailureMessage"] = o.ValidatorFailureMessage
	}
	return json.Marshal(toSerialize)
}

type NullableAddDictionaryPasswordValidatorRequest struct {
	value *AddDictionaryPasswordValidatorRequest
	isSet bool
}

func (v NullableAddDictionaryPasswordValidatorRequest) Get() *AddDictionaryPasswordValidatorRequest {
	return v.value
}

func (v *NullableAddDictionaryPasswordValidatorRequest) Set(val *AddDictionaryPasswordValidatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDictionaryPasswordValidatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDictionaryPasswordValidatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDictionaryPasswordValidatorRequest(val *AddDictionaryPasswordValidatorRequest) *NullableAddDictionaryPasswordValidatorRequest {
	return &NullableAddDictionaryPasswordValidatorRequest{value: val, isSet: true}
}

func (v NullableAddDictionaryPasswordValidatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDictionaryPasswordValidatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
