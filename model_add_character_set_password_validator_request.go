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

// AddCharacterSetPasswordValidatorRequest struct for AddCharacterSetPasswordValidatorRequest
type AddCharacterSetPasswordValidatorRequest struct {
	// Name of the new Password Validator
	ValidatorName string                                       `json:"validatorName"`
	Schemas       []EnumcharacterSetPasswordValidatorSchemaUrn `json:"schemas"`
	// Specifies a character set containing characters that a password may contain and a value indicating the minimum number of characters required from that set.
	CharacterSet []string `json:"characterSet"`
	// Indicates whether this password validator allows passwords to contain characters outside of any of the user-defined character sets.
	AllowUnclassifiedCharacters bool `json:"allowUnclassifiedCharacters"`
	// Specifies the minimum number of character sets that must be represented in a proposed password.
	MinimumRequiredCharacterSets *int32 `json:"minimumRequiredCharacterSets,omitempty"`
	// A description for this Password Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether the password validator is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator.
	ValidatorRequirementDescription *string `json:"validatorRequirementDescription,omitempty"`
	// Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator.
	ValidatorFailureMessage *string `json:"validatorFailureMessage,omitempty"`
}

// NewAddCharacterSetPasswordValidatorRequest instantiates a new AddCharacterSetPasswordValidatorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCharacterSetPasswordValidatorRequest(validatorName string, schemas []EnumcharacterSetPasswordValidatorSchemaUrn, characterSet []string, allowUnclassifiedCharacters bool, enabled bool) *AddCharacterSetPasswordValidatorRequest {
	this := AddCharacterSetPasswordValidatorRequest{}
	this.ValidatorName = validatorName
	this.Schemas = schemas
	this.CharacterSet = characterSet
	this.AllowUnclassifiedCharacters = allowUnclassifiedCharacters
	this.Enabled = enabled
	return &this
}

// NewAddCharacterSetPasswordValidatorRequestWithDefaults instantiates a new AddCharacterSetPasswordValidatorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCharacterSetPasswordValidatorRequestWithDefaults() *AddCharacterSetPasswordValidatorRequest {
	this := AddCharacterSetPasswordValidatorRequest{}
	return &this
}

// GetValidatorName returns the ValidatorName field value
func (o *AddCharacterSetPasswordValidatorRequest) GetValidatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorName
}

// GetValidatorNameOk returns a tuple with the ValidatorName field value
// and a boolean to check if the value has been set.
func (o *AddCharacterSetPasswordValidatorRequest) GetValidatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorName, true
}

// SetValidatorName sets field value
func (o *AddCharacterSetPasswordValidatorRequest) SetValidatorName(v string) {
	o.ValidatorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddCharacterSetPasswordValidatorRequest) GetSchemas() []EnumcharacterSetPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumcharacterSetPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddCharacterSetPasswordValidatorRequest) GetSchemasOk() ([]EnumcharacterSetPasswordValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddCharacterSetPasswordValidatorRequest) SetSchemas(v []EnumcharacterSetPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetCharacterSet returns the CharacterSet field value
func (o *AddCharacterSetPasswordValidatorRequest) GetCharacterSet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CharacterSet
}

// GetCharacterSetOk returns a tuple with the CharacterSet field value
// and a boolean to check if the value has been set.
func (o *AddCharacterSetPasswordValidatorRequest) GetCharacterSetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CharacterSet, true
}

// SetCharacterSet sets field value
func (o *AddCharacterSetPasswordValidatorRequest) SetCharacterSet(v []string) {
	o.CharacterSet = v
}

// GetAllowUnclassifiedCharacters returns the AllowUnclassifiedCharacters field value
func (o *AddCharacterSetPasswordValidatorRequest) GetAllowUnclassifiedCharacters() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowUnclassifiedCharacters
}

// GetAllowUnclassifiedCharactersOk returns a tuple with the AllowUnclassifiedCharacters field value
// and a boolean to check if the value has been set.
func (o *AddCharacterSetPasswordValidatorRequest) GetAllowUnclassifiedCharactersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowUnclassifiedCharacters, true
}

// SetAllowUnclassifiedCharacters sets field value
func (o *AddCharacterSetPasswordValidatorRequest) SetAllowUnclassifiedCharacters(v bool) {
	o.AllowUnclassifiedCharacters = v
}

// GetMinimumRequiredCharacterSets returns the MinimumRequiredCharacterSets field value if set, zero value otherwise.
func (o *AddCharacterSetPasswordValidatorRequest) GetMinimumRequiredCharacterSets() int32 {
	if o == nil || isNil(o.MinimumRequiredCharacterSets) {
		var ret int32
		return ret
	}
	return *o.MinimumRequiredCharacterSets
}

// GetMinimumRequiredCharacterSetsOk returns a tuple with the MinimumRequiredCharacterSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCharacterSetPasswordValidatorRequest) GetMinimumRequiredCharacterSetsOk() (*int32, bool) {
	if o == nil || isNil(o.MinimumRequiredCharacterSets) {
		return nil, false
	}
	return o.MinimumRequiredCharacterSets, true
}

// HasMinimumRequiredCharacterSets returns a boolean if a field has been set.
func (o *AddCharacterSetPasswordValidatorRequest) HasMinimumRequiredCharacterSets() bool {
	if o != nil && !isNil(o.MinimumRequiredCharacterSets) {
		return true
	}

	return false
}

// SetMinimumRequiredCharacterSets gets a reference to the given int32 and assigns it to the MinimumRequiredCharacterSets field.
func (o *AddCharacterSetPasswordValidatorRequest) SetMinimumRequiredCharacterSets(v int32) {
	o.MinimumRequiredCharacterSets = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddCharacterSetPasswordValidatorRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCharacterSetPasswordValidatorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddCharacterSetPasswordValidatorRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddCharacterSetPasswordValidatorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddCharacterSetPasswordValidatorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddCharacterSetPasswordValidatorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddCharacterSetPasswordValidatorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *AddCharacterSetPasswordValidatorRequest) GetValidatorRequirementDescription() string {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCharacterSetPasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
		return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *AddCharacterSetPasswordValidatorRequest) HasValidatorRequirementDescription() bool {
	if o != nil && !isNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *AddCharacterSetPasswordValidatorRequest) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *AddCharacterSetPasswordValidatorRequest) GetValidatorFailureMessage() string {
	if o == nil || isNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCharacterSetPasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorFailureMessage) {
		return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *AddCharacterSetPasswordValidatorRequest) HasValidatorFailureMessage() bool {
	if o != nil && !isNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *AddCharacterSetPasswordValidatorRequest) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

func (o AddCharacterSetPasswordValidatorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["validatorName"] = o.ValidatorName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["characterSet"] = o.CharacterSet
	}
	if true {
		toSerialize["allowUnclassifiedCharacters"] = o.AllowUnclassifiedCharacters
	}
	if !isNil(o.MinimumRequiredCharacterSets) {
		toSerialize["minimumRequiredCharacterSets"] = o.MinimumRequiredCharacterSets
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

type NullableAddCharacterSetPasswordValidatorRequest struct {
	value *AddCharacterSetPasswordValidatorRequest
	isSet bool
}

func (v NullableAddCharacterSetPasswordValidatorRequest) Get() *AddCharacterSetPasswordValidatorRequest {
	return v.value
}

func (v *NullableAddCharacterSetPasswordValidatorRequest) Set(val *AddCharacterSetPasswordValidatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCharacterSetPasswordValidatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCharacterSetPasswordValidatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCharacterSetPasswordValidatorRequest(val *AddCharacterSetPasswordValidatorRequest) *NullableAddCharacterSetPasswordValidatorRequest {
	return &NullableAddCharacterSetPasswordValidatorRequest{value: val, isSet: true}
}

func (v NullableAddCharacterSetPasswordValidatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCharacterSetPasswordValidatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
