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

// checks if the AddRandomPasswordGeneratorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddRandomPasswordGeneratorRequest{}

// AddRandomPasswordGeneratorRequest struct for AddRandomPasswordGeneratorRequest
type AddRandomPasswordGeneratorRequest struct {
	Schemas []EnumrandomPasswordGeneratorSchemaUrn `json:"schemas"`
	// Specifies one or more named character sets.
	PasswordCharacterSet []string `json:"passwordCharacterSet"`
	// Specifies the format to use for the generated password.
	PasswordFormat string `json:"passwordFormat"`
	// A description for this Password Generator
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Generator is enabled for use.
	Enabled bool `json:"enabled"`
	// Name of the new Password Generator
	GeneratorName string `json:"generatorName"`
}

// NewAddRandomPasswordGeneratorRequest instantiates a new AddRandomPasswordGeneratorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddRandomPasswordGeneratorRequest(schemas []EnumrandomPasswordGeneratorSchemaUrn, passwordCharacterSet []string, passwordFormat string, enabled bool, generatorName string) *AddRandomPasswordGeneratorRequest {
	this := AddRandomPasswordGeneratorRequest{}
	this.Schemas = schemas
	this.PasswordCharacterSet = passwordCharacterSet
	this.PasswordFormat = passwordFormat
	this.Enabled = enabled
	this.GeneratorName = generatorName
	return &this
}

// NewAddRandomPasswordGeneratorRequestWithDefaults instantiates a new AddRandomPasswordGeneratorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddRandomPasswordGeneratorRequestWithDefaults() *AddRandomPasswordGeneratorRequest {
	this := AddRandomPasswordGeneratorRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddRandomPasswordGeneratorRequest) GetSchemas() []EnumrandomPasswordGeneratorSchemaUrn {
	if o == nil {
		var ret []EnumrandomPasswordGeneratorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddRandomPasswordGeneratorRequest) GetSchemasOk() ([]EnumrandomPasswordGeneratorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddRandomPasswordGeneratorRequest) SetSchemas(v []EnumrandomPasswordGeneratorSchemaUrn) {
	o.Schemas = v
}

// GetPasswordCharacterSet returns the PasswordCharacterSet field value
func (o *AddRandomPasswordGeneratorRequest) GetPasswordCharacterSet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PasswordCharacterSet
}

// GetPasswordCharacterSetOk returns a tuple with the PasswordCharacterSet field value
// and a boolean to check if the value has been set.
func (o *AddRandomPasswordGeneratorRequest) GetPasswordCharacterSetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordCharacterSet, true
}

// SetPasswordCharacterSet sets field value
func (o *AddRandomPasswordGeneratorRequest) SetPasswordCharacterSet(v []string) {
	o.PasswordCharacterSet = v
}

// GetPasswordFormat returns the PasswordFormat field value
func (o *AddRandomPasswordGeneratorRequest) GetPasswordFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordFormat
}

// GetPasswordFormatOk returns a tuple with the PasswordFormat field value
// and a boolean to check if the value has been set.
func (o *AddRandomPasswordGeneratorRequest) GetPasswordFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordFormat, true
}

// SetPasswordFormat sets field value
func (o *AddRandomPasswordGeneratorRequest) SetPasswordFormat(v string) {
	o.PasswordFormat = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddRandomPasswordGeneratorRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRandomPasswordGeneratorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddRandomPasswordGeneratorRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddRandomPasswordGeneratorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddRandomPasswordGeneratorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddRandomPasswordGeneratorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddRandomPasswordGeneratorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetGeneratorName returns the GeneratorName field value
func (o *AddRandomPasswordGeneratorRequest) GetGeneratorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GeneratorName
}

// GetGeneratorNameOk returns a tuple with the GeneratorName field value
// and a boolean to check if the value has been set.
func (o *AddRandomPasswordGeneratorRequest) GetGeneratorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneratorName, true
}

// SetGeneratorName sets field value
func (o *AddRandomPasswordGeneratorRequest) SetGeneratorName(v string) {
	o.GeneratorName = v
}

func (o AddRandomPasswordGeneratorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddRandomPasswordGeneratorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["passwordCharacterSet"] = o.PasswordCharacterSet
	toSerialize["passwordFormat"] = o.PasswordFormat
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["generatorName"] = o.GeneratorName
	return toSerialize, nil
}

type NullableAddRandomPasswordGeneratorRequest struct {
	value *AddRandomPasswordGeneratorRequest
	isSet bool
}

func (v NullableAddRandomPasswordGeneratorRequest) Get() *AddRandomPasswordGeneratorRequest {
	return v.value
}

func (v *NullableAddRandomPasswordGeneratorRequest) Set(val *AddRandomPasswordGeneratorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddRandomPasswordGeneratorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddRandomPasswordGeneratorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddRandomPasswordGeneratorRequest(val *AddRandomPasswordGeneratorRequest) *NullableAddRandomPasswordGeneratorRequest {
	return &NullableAddRandomPasswordGeneratorRequest{value: val, isSet: true}
}

func (v NullableAddRandomPasswordGeneratorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddRandomPasswordGeneratorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
