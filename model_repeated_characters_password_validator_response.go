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

// RepeatedCharactersPasswordValidatorResponse struct for RepeatedCharactersPasswordValidatorResponse
type RepeatedCharactersPasswordValidatorResponse struct {
	// Name of the Password Validator
	Id string `json:"id"`
	Schemas []EnumrepeatedCharactersPasswordValidatorSchemaUrn `json:"schemas"`
	// Specifies the maximum number of times that any character can appear consecutively in a password value.
	MaxConsecutiveLength int32 `json:"maxConsecutiveLength"`
	// Indicates whether this password validator should treat password characters in a case-sensitive manner.
	CaseSensitiveValidation bool `json:"caseSensitiveValidation"`
	CharacterSet []string `json:"characterSet,omitempty"`
	// A description for this Password Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether the password validator is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator.
	ValidatorRequirementDescription *string `json:"validatorRequirementDescription,omitempty"`
	// Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator.
	ValidatorFailureMessage *string `json:"validatorFailureMessage,omitempty"`
}

// NewRepeatedCharactersPasswordValidatorResponse instantiates a new RepeatedCharactersPasswordValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepeatedCharactersPasswordValidatorResponse(id string, schemas []EnumrepeatedCharactersPasswordValidatorSchemaUrn, maxConsecutiveLength int32, caseSensitiveValidation bool, enabled bool) *RepeatedCharactersPasswordValidatorResponse {
	this := RepeatedCharactersPasswordValidatorResponse{}
	this.Id = id
	this.Schemas = schemas
	this.MaxConsecutiveLength = maxConsecutiveLength
	this.CaseSensitiveValidation = caseSensitiveValidation
	this.Enabled = enabled
	return &this
}

// NewRepeatedCharactersPasswordValidatorResponseWithDefaults instantiates a new RepeatedCharactersPasswordValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepeatedCharactersPasswordValidatorResponseWithDefaults() *RepeatedCharactersPasswordValidatorResponse {
	this := RepeatedCharactersPasswordValidatorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *RepeatedCharactersPasswordValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RepeatedCharactersPasswordValidatorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *RepeatedCharactersPasswordValidatorResponse) GetSchemas() []EnumrepeatedCharactersPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumrepeatedCharactersPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) GetSchemasOk() ([]EnumrepeatedCharactersPasswordValidatorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *RepeatedCharactersPasswordValidatorResponse) SetSchemas(v []EnumrepeatedCharactersPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetMaxConsecutiveLength returns the MaxConsecutiveLength field value
func (o *RepeatedCharactersPasswordValidatorResponse) GetMaxConsecutiveLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxConsecutiveLength
}

// GetMaxConsecutiveLengthOk returns a tuple with the MaxConsecutiveLength field value
// and a boolean to check if the value has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) GetMaxConsecutiveLengthOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MaxConsecutiveLength, true
}

// SetMaxConsecutiveLength sets field value
func (o *RepeatedCharactersPasswordValidatorResponse) SetMaxConsecutiveLength(v int32) {
	o.MaxConsecutiveLength = v
}

// GetCaseSensitiveValidation returns the CaseSensitiveValidation field value
func (o *RepeatedCharactersPasswordValidatorResponse) GetCaseSensitiveValidation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CaseSensitiveValidation
}

// GetCaseSensitiveValidationOk returns a tuple with the CaseSensitiveValidation field value
// and a boolean to check if the value has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) GetCaseSensitiveValidationOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CaseSensitiveValidation, true
}

// SetCaseSensitiveValidation sets field value
func (o *RepeatedCharactersPasswordValidatorResponse) SetCaseSensitiveValidation(v bool) {
	o.CaseSensitiveValidation = v
}

// GetCharacterSet returns the CharacterSet field value if set, zero value otherwise.
func (o *RepeatedCharactersPasswordValidatorResponse) GetCharacterSet() []string {
	if o == nil || isNil(o.CharacterSet) {
		var ret []string
		return ret
	}
	return o.CharacterSet
}

// GetCharacterSetOk returns a tuple with the CharacterSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) GetCharacterSetOk() ([]string, bool) {
	if o == nil || isNil(o.CharacterSet) {
    return nil, false
	}
	return o.CharacterSet, true
}

// HasCharacterSet returns a boolean if a field has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) HasCharacterSet() bool {
	if o != nil && !isNil(o.CharacterSet) {
		return true
	}

	return false
}

// SetCharacterSet gets a reference to the given []string and assigns it to the CharacterSet field.
func (o *RepeatedCharactersPasswordValidatorResponse) SetCharacterSet(v []string) {
	o.CharacterSet = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RepeatedCharactersPasswordValidatorResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RepeatedCharactersPasswordValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *RepeatedCharactersPasswordValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *RepeatedCharactersPasswordValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *RepeatedCharactersPasswordValidatorResponse) GetValidatorRequirementDescription() string {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
    return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) HasValidatorRequirementDescription() bool {
	if o != nil && !isNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *RepeatedCharactersPasswordValidatorResponse) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *RepeatedCharactersPasswordValidatorResponse) GetValidatorFailureMessage() string {
	if o == nil || isNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorFailureMessage) {
    return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *RepeatedCharactersPasswordValidatorResponse) HasValidatorFailureMessage() bool {
	if o != nil && !isNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *RepeatedCharactersPasswordValidatorResponse) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

func (o RepeatedCharactersPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["maxConsecutiveLength"] = o.MaxConsecutiveLength
	}
	if true {
		toSerialize["caseSensitiveValidation"] = o.CaseSensitiveValidation
	}
	if !isNil(o.CharacterSet) {
		toSerialize["characterSet"] = o.CharacterSet
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

type NullableRepeatedCharactersPasswordValidatorResponse struct {
	value *RepeatedCharactersPasswordValidatorResponse
	isSet bool
}

func (v NullableRepeatedCharactersPasswordValidatorResponse) Get() *RepeatedCharactersPasswordValidatorResponse {
	return v.value
}

func (v *NullableRepeatedCharactersPasswordValidatorResponse) Set(val *RepeatedCharactersPasswordValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRepeatedCharactersPasswordValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRepeatedCharactersPasswordValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepeatedCharactersPasswordValidatorResponse(val *RepeatedCharactersPasswordValidatorResponse) *NullableRepeatedCharactersPasswordValidatorResponse {
	return &NullableRepeatedCharactersPasswordValidatorResponse{value: val, isSet: true}
}

func (v NullableRepeatedCharactersPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepeatedCharactersPasswordValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


