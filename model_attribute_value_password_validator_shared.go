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

// AttributeValuePasswordValidatorShared struct for AttributeValuePasswordValidatorShared
type AttributeValuePasswordValidatorShared struct {
	Schemas []EnumattributeValuePasswordValidatorSchemaUrn `json:"schemas"`
	MatchAttribute []string `json:"matchAttribute,omitempty"`
	// Indicates whether to reject any proposed password that is a substring of a value in one of the match attributes in the target user's entry.
	TestPasswordSubstringOfAttributeValue *bool `json:"testPasswordSubstringOfAttributeValue,omitempty"`
	// Indicates whether to reject any proposed password in which a value in one of the match attributes in the target user's entry is a substring of that password.
	TestAttributeValueSubstringOfPassword *bool `json:"testAttributeValueSubstringOfPassword,omitempty"`
	// The minimum length that an attribute value must have for it to be considered when rejecting passwords that contain the value of another attribute as a substring.
	MinimumAttributeValueLengthForSubstringMatches *int32 `json:"minimumAttributeValueLengthForSubstringMatches,omitempty"`
	// Indicates whether to perform matching against the reversed value of the provided password in addition to the order in which it was given.
	TestReversedPassword bool `json:"testReversedPassword"`
	// A description for this Password Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether the password validator is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator.
	ValidatorRequirementDescription *string `json:"validatorRequirementDescription,omitempty"`
	// Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator.
	ValidatorFailureMessage *string `json:"validatorFailureMessage,omitempty"`
}

// NewAttributeValuePasswordValidatorShared instantiates a new AttributeValuePasswordValidatorShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeValuePasswordValidatorShared(schemas []EnumattributeValuePasswordValidatorSchemaUrn, testReversedPassword bool, enabled bool) *AttributeValuePasswordValidatorShared {
	this := AttributeValuePasswordValidatorShared{}
	this.Schemas = schemas
	this.TestReversedPassword = testReversedPassword
	this.Enabled = enabled
	return &this
}

// NewAttributeValuePasswordValidatorSharedWithDefaults instantiates a new AttributeValuePasswordValidatorShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeValuePasswordValidatorSharedWithDefaults() *AttributeValuePasswordValidatorShared {
	this := AttributeValuePasswordValidatorShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AttributeValuePasswordValidatorShared) GetSchemas() []EnumattributeValuePasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumattributeValuePasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AttributeValuePasswordValidatorShared) GetSchemasOk() ([]EnumattributeValuePasswordValidatorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AttributeValuePasswordValidatorShared) SetSchemas(v []EnumattributeValuePasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetMatchAttribute returns the MatchAttribute field value if set, zero value otherwise.
func (o *AttributeValuePasswordValidatorShared) GetMatchAttribute() []string {
	if o == nil || isNil(o.MatchAttribute) {
		var ret []string
		return ret
	}
	return o.MatchAttribute
}

// GetMatchAttributeOk returns a tuple with the MatchAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeValuePasswordValidatorShared) GetMatchAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.MatchAttribute) {
    return nil, false
	}
	return o.MatchAttribute, true
}

// HasMatchAttribute returns a boolean if a field has been set.
func (o *AttributeValuePasswordValidatorShared) HasMatchAttribute() bool {
	if o != nil && !isNil(o.MatchAttribute) {
		return true
	}

	return false
}

// SetMatchAttribute gets a reference to the given []string and assigns it to the MatchAttribute field.
func (o *AttributeValuePasswordValidatorShared) SetMatchAttribute(v []string) {
	o.MatchAttribute = v
}

// GetTestPasswordSubstringOfAttributeValue returns the TestPasswordSubstringOfAttributeValue field value if set, zero value otherwise.
func (o *AttributeValuePasswordValidatorShared) GetTestPasswordSubstringOfAttributeValue() bool {
	if o == nil || isNil(o.TestPasswordSubstringOfAttributeValue) {
		var ret bool
		return ret
	}
	return *o.TestPasswordSubstringOfAttributeValue
}

// GetTestPasswordSubstringOfAttributeValueOk returns a tuple with the TestPasswordSubstringOfAttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeValuePasswordValidatorShared) GetTestPasswordSubstringOfAttributeValueOk() (*bool, bool) {
	if o == nil || isNil(o.TestPasswordSubstringOfAttributeValue) {
    return nil, false
	}
	return o.TestPasswordSubstringOfAttributeValue, true
}

// HasTestPasswordSubstringOfAttributeValue returns a boolean if a field has been set.
func (o *AttributeValuePasswordValidatorShared) HasTestPasswordSubstringOfAttributeValue() bool {
	if o != nil && !isNil(o.TestPasswordSubstringOfAttributeValue) {
		return true
	}

	return false
}

// SetTestPasswordSubstringOfAttributeValue gets a reference to the given bool and assigns it to the TestPasswordSubstringOfAttributeValue field.
func (o *AttributeValuePasswordValidatorShared) SetTestPasswordSubstringOfAttributeValue(v bool) {
	o.TestPasswordSubstringOfAttributeValue = &v
}

// GetTestAttributeValueSubstringOfPassword returns the TestAttributeValueSubstringOfPassword field value if set, zero value otherwise.
func (o *AttributeValuePasswordValidatorShared) GetTestAttributeValueSubstringOfPassword() bool {
	if o == nil || isNil(o.TestAttributeValueSubstringOfPassword) {
		var ret bool
		return ret
	}
	return *o.TestAttributeValueSubstringOfPassword
}

// GetTestAttributeValueSubstringOfPasswordOk returns a tuple with the TestAttributeValueSubstringOfPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeValuePasswordValidatorShared) GetTestAttributeValueSubstringOfPasswordOk() (*bool, bool) {
	if o == nil || isNil(o.TestAttributeValueSubstringOfPassword) {
    return nil, false
	}
	return o.TestAttributeValueSubstringOfPassword, true
}

// HasTestAttributeValueSubstringOfPassword returns a boolean if a field has been set.
func (o *AttributeValuePasswordValidatorShared) HasTestAttributeValueSubstringOfPassword() bool {
	if o != nil && !isNil(o.TestAttributeValueSubstringOfPassword) {
		return true
	}

	return false
}

// SetTestAttributeValueSubstringOfPassword gets a reference to the given bool and assigns it to the TestAttributeValueSubstringOfPassword field.
func (o *AttributeValuePasswordValidatorShared) SetTestAttributeValueSubstringOfPassword(v bool) {
	o.TestAttributeValueSubstringOfPassword = &v
}

// GetMinimumAttributeValueLengthForSubstringMatches returns the MinimumAttributeValueLengthForSubstringMatches field value if set, zero value otherwise.
func (o *AttributeValuePasswordValidatorShared) GetMinimumAttributeValueLengthForSubstringMatches() int32 {
	if o == nil || isNil(o.MinimumAttributeValueLengthForSubstringMatches) {
		var ret int32
		return ret
	}
	return *o.MinimumAttributeValueLengthForSubstringMatches
}

// GetMinimumAttributeValueLengthForSubstringMatchesOk returns a tuple with the MinimumAttributeValueLengthForSubstringMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeValuePasswordValidatorShared) GetMinimumAttributeValueLengthForSubstringMatchesOk() (*int32, bool) {
	if o == nil || isNil(o.MinimumAttributeValueLengthForSubstringMatches) {
    return nil, false
	}
	return o.MinimumAttributeValueLengthForSubstringMatches, true
}

// HasMinimumAttributeValueLengthForSubstringMatches returns a boolean if a field has been set.
func (o *AttributeValuePasswordValidatorShared) HasMinimumAttributeValueLengthForSubstringMatches() bool {
	if o != nil && !isNil(o.MinimumAttributeValueLengthForSubstringMatches) {
		return true
	}

	return false
}

// SetMinimumAttributeValueLengthForSubstringMatches gets a reference to the given int32 and assigns it to the MinimumAttributeValueLengthForSubstringMatches field.
func (o *AttributeValuePasswordValidatorShared) SetMinimumAttributeValueLengthForSubstringMatches(v int32) {
	o.MinimumAttributeValueLengthForSubstringMatches = &v
}

// GetTestReversedPassword returns the TestReversedPassword field value
func (o *AttributeValuePasswordValidatorShared) GetTestReversedPassword() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TestReversedPassword
}

// GetTestReversedPasswordOk returns a tuple with the TestReversedPassword field value
// and a boolean to check if the value has been set.
func (o *AttributeValuePasswordValidatorShared) GetTestReversedPasswordOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TestReversedPassword, true
}

// SetTestReversedPassword sets field value
func (o *AttributeValuePasswordValidatorShared) SetTestReversedPassword(v bool) {
	o.TestReversedPassword = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AttributeValuePasswordValidatorShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeValuePasswordValidatorShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AttributeValuePasswordValidatorShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AttributeValuePasswordValidatorShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AttributeValuePasswordValidatorShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AttributeValuePasswordValidatorShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AttributeValuePasswordValidatorShared) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *AttributeValuePasswordValidatorShared) GetValidatorRequirementDescription() string {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeValuePasswordValidatorShared) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorRequirementDescription) {
    return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *AttributeValuePasswordValidatorShared) HasValidatorRequirementDescription() bool {
	if o != nil && !isNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *AttributeValuePasswordValidatorShared) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *AttributeValuePasswordValidatorShared) GetValidatorFailureMessage() string {
	if o == nil || isNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeValuePasswordValidatorShared) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || isNil(o.ValidatorFailureMessage) {
    return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *AttributeValuePasswordValidatorShared) HasValidatorFailureMessage() bool {
	if o != nil && !isNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *AttributeValuePasswordValidatorShared) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

func (o AttributeValuePasswordValidatorShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.MatchAttribute) {
		toSerialize["matchAttribute"] = o.MatchAttribute
	}
	if !isNil(o.TestPasswordSubstringOfAttributeValue) {
		toSerialize["testPasswordSubstringOfAttributeValue"] = o.TestPasswordSubstringOfAttributeValue
	}
	if !isNil(o.TestAttributeValueSubstringOfPassword) {
		toSerialize["testAttributeValueSubstringOfPassword"] = o.TestAttributeValueSubstringOfPassword
	}
	if !isNil(o.MinimumAttributeValueLengthForSubstringMatches) {
		toSerialize["minimumAttributeValueLengthForSubstringMatches"] = o.MinimumAttributeValueLengthForSubstringMatches
	}
	if true {
		toSerialize["testReversedPassword"] = o.TestReversedPassword
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

type NullableAttributeValuePasswordValidatorShared struct {
	value *AttributeValuePasswordValidatorShared
	isSet bool
}

func (v NullableAttributeValuePasswordValidatorShared) Get() *AttributeValuePasswordValidatorShared {
	return v.value
}

func (v *NullableAttributeValuePasswordValidatorShared) Set(val *AttributeValuePasswordValidatorShared) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeValuePasswordValidatorShared) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeValuePasswordValidatorShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeValuePasswordValidatorShared(val *AttributeValuePasswordValidatorShared) *NullableAttributeValuePasswordValidatorShared {
	return &NullableAttributeValuePasswordValidatorShared{value: val, isSet: true}
}

func (v NullableAttributeValuePasswordValidatorShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeValuePasswordValidatorShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


