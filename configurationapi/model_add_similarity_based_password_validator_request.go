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

// checks if the AddSimilarityBasedPasswordValidatorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSimilarityBasedPasswordValidatorRequest{}

// AddSimilarityBasedPasswordValidatorRequest struct for AddSimilarityBasedPasswordValidatorRequest
type AddSimilarityBasedPasswordValidatorRequest struct {
	Schemas []EnumsimilarityBasedPasswordValidatorSchemaUrn `json:"schemas"`
	// Specifies the minimum difference of new and old password.
	MinPasswordDifference int64 `json:"minPasswordDifference"`
	// A description for this Password Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether the password validator is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator.
	ValidatorRequirementDescription *string `json:"validatorRequirementDescription,omitempty"`
	// Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator.
	ValidatorFailureMessage *string `json:"validatorFailureMessage,omitempty"`
	// Name of the new Password Validator
	ValidatorName string `json:"validatorName"`
}

// NewAddSimilarityBasedPasswordValidatorRequest instantiates a new AddSimilarityBasedPasswordValidatorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSimilarityBasedPasswordValidatorRequest(schemas []EnumsimilarityBasedPasswordValidatorSchemaUrn, minPasswordDifference int64, enabled bool, validatorName string) *AddSimilarityBasedPasswordValidatorRequest {
	this := AddSimilarityBasedPasswordValidatorRequest{}
	this.Schemas = schemas
	this.MinPasswordDifference = minPasswordDifference
	this.Enabled = enabled
	this.ValidatorName = validatorName
	return &this
}

// NewAddSimilarityBasedPasswordValidatorRequestWithDefaults instantiates a new AddSimilarityBasedPasswordValidatorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSimilarityBasedPasswordValidatorRequestWithDefaults() *AddSimilarityBasedPasswordValidatorRequest {
	this := AddSimilarityBasedPasswordValidatorRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddSimilarityBasedPasswordValidatorRequest) GetSchemas() []EnumsimilarityBasedPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumsimilarityBasedPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSimilarityBasedPasswordValidatorRequest) GetSchemasOk() ([]EnumsimilarityBasedPasswordValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSimilarityBasedPasswordValidatorRequest) SetSchemas(v []EnumsimilarityBasedPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetMinPasswordDifference returns the MinPasswordDifference field value
func (o *AddSimilarityBasedPasswordValidatorRequest) GetMinPasswordDifference() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinPasswordDifference
}

// GetMinPasswordDifferenceOk returns a tuple with the MinPasswordDifference field value
// and a boolean to check if the value has been set.
func (o *AddSimilarityBasedPasswordValidatorRequest) GetMinPasswordDifferenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinPasswordDifference, true
}

// SetMinPasswordDifference sets field value
func (o *AddSimilarityBasedPasswordValidatorRequest) SetMinPasswordDifference(v int64) {
	o.MinPasswordDifference = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSimilarityBasedPasswordValidatorRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimilarityBasedPasswordValidatorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSimilarityBasedPasswordValidatorRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSimilarityBasedPasswordValidatorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddSimilarityBasedPasswordValidatorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddSimilarityBasedPasswordValidatorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddSimilarityBasedPasswordValidatorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorRequirementDescription() string {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *AddSimilarityBasedPasswordValidatorRequest) HasValidatorRequirementDescription() bool {
	if o != nil && !IsNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *AddSimilarityBasedPasswordValidatorRequest) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorFailureMessage() string {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *AddSimilarityBasedPasswordValidatorRequest) HasValidatorFailureMessage() bool {
	if o != nil && !IsNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *AddSimilarityBasedPasswordValidatorRequest) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

// GetValidatorName returns the ValidatorName field value
func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorName
}

// GetValidatorNameOk returns a tuple with the ValidatorName field value
// and a boolean to check if the value has been set.
func (o *AddSimilarityBasedPasswordValidatorRequest) GetValidatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorName, true
}

// SetValidatorName sets field value
func (o *AddSimilarityBasedPasswordValidatorRequest) SetValidatorName(v string) {
	o.ValidatorName = v
}

func (o AddSimilarityBasedPasswordValidatorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSimilarityBasedPasswordValidatorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["minPasswordDifference"] = o.MinPasswordDifference
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.ValidatorRequirementDescription) {
		toSerialize["validatorRequirementDescription"] = o.ValidatorRequirementDescription
	}
	if !IsNil(o.ValidatorFailureMessage) {
		toSerialize["validatorFailureMessage"] = o.ValidatorFailureMessage
	}
	toSerialize["validatorName"] = o.ValidatorName
	return toSerialize, nil
}

type NullableAddSimilarityBasedPasswordValidatorRequest struct {
	value *AddSimilarityBasedPasswordValidatorRequest
	isSet bool
}

func (v NullableAddSimilarityBasedPasswordValidatorRequest) Get() *AddSimilarityBasedPasswordValidatorRequest {
	return v.value
}

func (v *NullableAddSimilarityBasedPasswordValidatorRequest) Set(val *AddSimilarityBasedPasswordValidatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSimilarityBasedPasswordValidatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSimilarityBasedPasswordValidatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSimilarityBasedPasswordValidatorRequest(val *AddSimilarityBasedPasswordValidatorRequest) *NullableAddSimilarityBasedPasswordValidatorRequest {
	return &NullableAddSimilarityBasedPasswordValidatorRequest{value: val, isSet: true}
}

func (v NullableAddSimilarityBasedPasswordValidatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSimilarityBasedPasswordValidatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
