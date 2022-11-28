/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DictionaryPasswordValidatorResponse struct for DictionaryPasswordValidatorResponse
type DictionaryPasswordValidatorResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Password Validator
	Id string `json:"id"`
	Schemas []EnumdictionaryPasswordValidatorSchemaUrn `json:"schemas"`
	// Specifies the path to the file containing a list of words that cannot be used as passwords.
	DictionaryFile string `json:"dictionaryFile"`
	// Indicates whether this password validator is to treat password characters in a case-sensitive manner.
	CaseSensitiveValidation bool `json:"caseSensitiveValidation"`
	// Indicates whether this password validator is to test the reversed value of the provided password as well as the order in which it was given.
	TestReversedPassword bool `json:"testReversedPassword"`
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

// NewDictionaryPasswordValidatorResponse instantiates a new DictionaryPasswordValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDictionaryPasswordValidatorResponse(id string, schemas []EnumdictionaryPasswordValidatorSchemaUrn, dictionaryFile string, caseSensitiveValidation bool, testReversedPassword bool, enabled bool) *DictionaryPasswordValidatorResponse {
	this := DictionaryPasswordValidatorResponse{}
	this.Id = id
	this.Schemas = schemas
	this.DictionaryFile = dictionaryFile
	this.CaseSensitiveValidation = caseSensitiveValidation
	this.TestReversedPassword = testReversedPassword
	this.Enabled = enabled
	return &this
}

// NewDictionaryPasswordValidatorResponseWithDefaults instantiates a new DictionaryPasswordValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDictionaryPasswordValidatorResponseWithDefaults() *DictionaryPasswordValidatorResponse {
	this := DictionaryPasswordValidatorResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DictionaryPasswordValidatorResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DictionaryPasswordValidatorResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DictionaryPasswordValidatorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *DictionaryPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *DictionaryPasswordValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *DictionaryPasswordValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *DictionaryPasswordValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DictionaryPasswordValidatorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *DictionaryPasswordValidatorResponse) GetSchemas() []EnumdictionaryPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumdictionaryPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetSchemasOk() ([]EnumdictionaryPasswordValidatorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *DictionaryPasswordValidatorResponse) SetSchemas(v []EnumdictionaryPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetDictionaryFile returns the DictionaryFile field value
func (o *DictionaryPasswordValidatorResponse) GetDictionaryFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DictionaryFile
}

// GetDictionaryFileOk returns a tuple with the DictionaryFile field value
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetDictionaryFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DictionaryFile, true
}

// SetDictionaryFile sets field value
func (o *DictionaryPasswordValidatorResponse) SetDictionaryFile(v string) {
	o.DictionaryFile = v
}

// GetCaseSensitiveValidation returns the CaseSensitiveValidation field value
func (o *DictionaryPasswordValidatorResponse) GetCaseSensitiveValidation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CaseSensitiveValidation
}

// GetCaseSensitiveValidationOk returns a tuple with the CaseSensitiveValidation field value
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetCaseSensitiveValidationOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CaseSensitiveValidation, true
}

// SetCaseSensitiveValidation sets field value
func (o *DictionaryPasswordValidatorResponse) SetCaseSensitiveValidation(v bool) {
	o.CaseSensitiveValidation = v
}

// GetTestReversedPassword returns the TestReversedPassword field value
func (o *DictionaryPasswordValidatorResponse) GetTestReversedPassword() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TestReversedPassword
}

// GetTestReversedPasswordOk returns a tuple with the TestReversedPassword field value
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetTestReversedPasswordOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TestReversedPassword, true
}

// SetTestReversedPassword sets field value
func (o *DictionaryPasswordValidatorResponse) SetTestReversedPassword(v bool) {
	o.TestReversedPassword = v
}

// GetIgnoreLeadingNonAlphabeticCharacters returns the IgnoreLeadingNonAlphabeticCharacters field value if set, zero value otherwise.
func (o *DictionaryPasswordValidatorResponse) GetIgnoreLeadingNonAlphabeticCharacters() bool {
	if o == nil || isNil(o.IgnoreLeadingNonAlphabeticCharacters) {
		var ret bool
		return ret
	}
	return *o.IgnoreLeadingNonAlphabeticCharacters
}

// GetIgnoreLeadingNonAlphabeticCharactersOk returns a tuple with the IgnoreLeadingNonAlphabeticCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetIgnoreLeadingNonAlphabeticCharactersOk() (*bool, bool) {
	if o == nil || isNil(o.IgnoreLeadingNonAlphabeticCharacters) {
    return nil, false
	}
	return o.IgnoreLeadingNonAlphabeticCharacters, true
}

// HasIgnoreLeadingNonAlphabeticCharacters returns a boolean if a field has been set.
func (o *DictionaryPasswordValidatorResponse) HasIgnoreLeadingNonAlphabeticCharacters() bool {
	if o != nil && !isNil(o.IgnoreLeadingNonAlphabeticCharacters) {
		return true
	}

	return false
}

// SetIgnoreLeadingNonAlphabeticCharacters gets a reference to the given bool and assigns it to the IgnoreLeadingNonAlphabeticCharacters field.
func (o *DictionaryPasswordValidatorResponse) SetIgnoreLeadingNonAlphabeticCharacters(v bool) {
	o.IgnoreLeadingNonAlphabeticCharacters = &v
}

// GetIgnoreTrailingNonAlphabeticCharacters returns the IgnoreTrailingNonAlphabeticCharacters field value if set, zero value otherwise.
func (o *DictionaryPasswordValidatorResponse) GetIgnoreTrailingNonAlphabeticCharacters() bool {
	if o == nil || isNil(o.IgnoreTrailingNonAlphabeticCharacters) {
		var ret bool
		return ret
	}
	return *o.IgnoreTrailingNonAlphabeticCharacters
}

// GetIgnoreTrailingNonAlphabeticCharactersOk returns a tuple with the IgnoreTrailingNonAlphabeticCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetIgnoreTrailingNonAlphabeticCharactersOk() (*bool, bool) {
	if o == nil || isNil(o.IgnoreTrailingNonAlphabeticCharacters) {
    return nil, false
	}
	return o.IgnoreTrailingNonAlphabeticCharacters, true
}

// HasIgnoreTrailingNonAlphabeticCharacters returns a boolean if a field has been set.
func (o *DictionaryPasswordValidatorResponse) HasIgnoreTrailingNonAlphabeticCharacters() bool {
	if o != nil && !isNil(o.IgnoreTrailingNonAlphabeticCharacters) {
		return true
	}

	return false
}

// SetIgnoreTrailingNonAlphabeticCharacters gets a reference to the given bool and assigns it to the IgnoreTrailingNonAlphabeticCharacters field.
func (o *DictionaryPasswordValidatorResponse) SetIgnoreTrailingNonAlphabeticCharacters(v bool) {
	o.IgnoreTrailingNonAlphabeticCharacters = &v
}

// GetStripDiacriticalMarks returns the StripDiacriticalMarks field value if set, zero value otherwise.
func (o *DictionaryPasswordValidatorResponse) GetStripDiacriticalMarks() bool {
	if o == nil || isNil(o.StripDiacriticalMarks) {
		var ret bool
		return ret
	}
	return *o.StripDiacriticalMarks
}

// GetStripDiacriticalMarksOk returns a tuple with the StripDiacriticalMarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetStripDiacriticalMarksOk() (*bool, bool) {
	if o == nil || isNil(o.StripDiacriticalMarks) {
    return nil, false
	}
	return o.StripDiacriticalMarks, true
}

// HasStripDiacriticalMarks returns a boolean if a field has been set.
func (o *DictionaryPasswordValidatorResponse) HasStripDiacriticalMarks() bool {
	if o != nil && !isNil(o.StripDiacriticalMarks) {
		return true
	}

	return false
}

// SetStripDiacriticalMarks gets a reference to the given bool and assigns it to the StripDiacriticalMarks field.
func (o *DictionaryPasswordValidatorResponse) SetStripDiacriticalMarks(v bool) {
	o.StripDiacriticalMarks = &v
}

// GetAlternativePasswordCharacterMapping returns the AlternativePasswordCharacterMapping field value if set, zero value otherwise.
func (o *DictionaryPasswordValidatorResponse) GetAlternativePasswordCharacterMapping() []string {
	if o == nil || isNil(o.AlternativePasswordCharacterMapping) {
		var ret []string
		return ret
	}
	return o.AlternativePasswordCharacterMapping
}

// GetAlternativePasswordCharacterMappingOk returns a tuple with the AlternativePasswordCharacterMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetAlternativePasswordCharacterMappingOk() ([]string, bool) {
	if o == nil || isNil(o.AlternativePasswordCharacterMapping) {
    return nil, false
	}
	return o.AlternativePasswordCharacterMapping, true
}

// HasAlternativePasswordCharacterMapping returns a boolean if a field has been set.
func (o *DictionaryPasswordValidatorResponse) HasAlternativePasswordCharacterMapping() bool {
	if o != nil && !isNil(o.AlternativePasswordCharacterMapping) {
		return true
	}

	return false
}

// SetAlternativePasswordCharacterMapping gets a reference to the given []string and assigns it to the AlternativePasswordCharacterMapping field.
func (o *DictionaryPasswordValidatorResponse) SetAlternativePasswordCharacterMapping(v []string) {
	o.AlternativePasswordCharacterMapping = v
}

// GetMaximumAllowedPercentOfPassword returns the MaximumAllowedPercentOfPassword field value if set, zero value otherwise.
func (o *DictionaryPasswordValidatorResponse) GetMaximumAllowedPercentOfPassword() int32 {
	if o == nil || isNil(o.MaximumAllowedPercentOfPassword) {
		var ret int32
		return ret
	}
	return *o.MaximumAllowedPercentOfPassword
}

// GetMaximumAllowedPercentOfPasswordOk returns a tuple with the MaximumAllowedPercentOfPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetMaximumAllowedPercentOfPasswordOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumAllowedPercentOfPassword) {
    return nil, false
	}
	return o.MaximumAllowedPercentOfPassword, true
}

// HasMaximumAllowedPercentOfPassword returns a boolean if a field has been set.
func (o *DictionaryPasswordValidatorResponse) HasMaximumAllowedPercentOfPassword() bool {
	if o != nil && !isNil(o.MaximumAllowedPercentOfPassword) {
		return true
	}

	return false
}

// SetMaximumAllowedPercentOfPassword gets a reference to the given int32 and assigns it to the MaximumAllowedPercentOfPassword field.
func (o *DictionaryPasswordValidatorResponse) SetMaximumAllowedPercentOfPassword(v int32) {
	o.MaximumAllowedPercentOfPassword = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DictionaryPasswordValidatorResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DictionaryPasswordValidatorResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DictionaryPasswordValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *DictionaryPasswordValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DictionaryPasswordValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *DictionaryPasswordValidatorResponse) GetValidatorRequirementDescription() string {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
    return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *DictionaryPasswordValidatorResponse) HasValidatorRequirementDescription() bool {
	if o != nil && !isNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *DictionaryPasswordValidatorResponse) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *DictionaryPasswordValidatorResponse) GetValidatorFailureMessage() string {
	if o == nil || isNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryPasswordValidatorResponse) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorFailureMessage) {
    return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *DictionaryPasswordValidatorResponse) HasValidatorFailureMessage() bool {
	if o != nil && !isNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *DictionaryPasswordValidatorResponse) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

func (o DictionaryPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["dictionaryFile"] = o.DictionaryFile
	}
	if true {
		toSerialize["caseSensitiveValidation"] = o.CaseSensitiveValidation
	}
	if true {
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

type NullableDictionaryPasswordValidatorResponse struct {
	value *DictionaryPasswordValidatorResponse
	isSet bool
}

func (v NullableDictionaryPasswordValidatorResponse) Get() *DictionaryPasswordValidatorResponse {
	return v.value
}

func (v *NullableDictionaryPasswordValidatorResponse) Set(val *DictionaryPasswordValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDictionaryPasswordValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDictionaryPasswordValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDictionaryPasswordValidatorResponse(val *DictionaryPasswordValidatorResponse) *NullableDictionaryPasswordValidatorResponse {
	return &NullableDictionaryPasswordValidatorResponse{value: val, isSet: true}
}

func (v NullableDictionaryPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDictionaryPasswordValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


