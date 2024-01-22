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

// checks if the AddThirdPartyPasswordValidatorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddThirdPartyPasswordValidatorRequest{}

// AddThirdPartyPasswordValidatorRequest struct for AddThirdPartyPasswordValidatorRequest
type AddThirdPartyPasswordValidatorRequest struct {
	Schemas []EnumthirdPartyPasswordValidatorSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Password Validator.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Password Validator. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
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

// NewAddThirdPartyPasswordValidatorRequest instantiates a new AddThirdPartyPasswordValidatorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyPasswordValidatorRequest(schemas []EnumthirdPartyPasswordValidatorSchemaUrn, extensionClass string, enabled bool, validatorName string) *AddThirdPartyPasswordValidatorRequest {
	this := AddThirdPartyPasswordValidatorRequest{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	this.ValidatorName = validatorName
	return &this
}

// NewAddThirdPartyPasswordValidatorRequestWithDefaults instantiates a new AddThirdPartyPasswordValidatorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyPasswordValidatorRequestWithDefaults() *AddThirdPartyPasswordValidatorRequest {
	this := AddThirdPartyPasswordValidatorRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyPasswordValidatorRequest) GetSchemas() []EnumthirdPartyPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordValidatorRequest) GetSchemasOk() ([]EnumthirdPartyPasswordValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyPasswordValidatorRequest) SetSchemas(v []EnumthirdPartyPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyPasswordValidatorRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordValidatorRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyPasswordValidatorRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyPasswordValidatorRequest) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordValidatorRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyPasswordValidatorRequest) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyPasswordValidatorRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyPasswordValidatorRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordValidatorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyPasswordValidatorRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyPasswordValidatorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyPasswordValidatorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordValidatorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyPasswordValidatorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorRequirementDescription() string {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *AddThirdPartyPasswordValidatorRequest) HasValidatorRequirementDescription() bool {
	if o != nil && !IsNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *AddThirdPartyPasswordValidatorRequest) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorFailureMessage() string {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *AddThirdPartyPasswordValidatorRequest) HasValidatorFailureMessage() bool {
	if o != nil && !IsNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *AddThirdPartyPasswordValidatorRequest) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

// GetValidatorName returns the ValidatorName field value
func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorName
}

// GetValidatorNameOk returns a tuple with the ValidatorName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordValidatorRequest) GetValidatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorName, true
}

// SetValidatorName sets field value
func (o *AddThirdPartyPasswordValidatorRequest) SetValidatorName(v string) {
	o.ValidatorName = v
}

func (o AddThirdPartyPasswordValidatorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddThirdPartyPasswordValidatorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
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

type NullableAddThirdPartyPasswordValidatorRequest struct {
	value *AddThirdPartyPasswordValidatorRequest
	isSet bool
}

func (v NullableAddThirdPartyPasswordValidatorRequest) Get() *AddThirdPartyPasswordValidatorRequest {
	return v.value
}

func (v *NullableAddThirdPartyPasswordValidatorRequest) Set(val *AddThirdPartyPasswordValidatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyPasswordValidatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyPasswordValidatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyPasswordValidatorRequest(val *AddThirdPartyPasswordValidatorRequest) *NullableAddThirdPartyPasswordValidatorRequest {
	return &NullableAddThirdPartyPasswordValidatorRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyPasswordValidatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyPasswordValidatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
