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

// AddThirdPartyPasswordGeneratorRequest struct for AddThirdPartyPasswordGeneratorRequest
type AddThirdPartyPasswordGeneratorRequest struct {
	// Name of the new Password Generator
	GeneratorName string `json:"generatorName"`
	Schemas []EnumthirdPartyPasswordGeneratorSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Password Generator.
	ExtensionClass string `json:"extensionClass"`
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Password Generator
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Generator is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddThirdPartyPasswordGeneratorRequest instantiates a new AddThirdPartyPasswordGeneratorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyPasswordGeneratorRequest(generatorName string, schemas []EnumthirdPartyPasswordGeneratorSchemaUrn, extensionClass string, enabled bool) *AddThirdPartyPasswordGeneratorRequest {
	this := AddThirdPartyPasswordGeneratorRequest{}
	this.GeneratorName = generatorName
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewAddThirdPartyPasswordGeneratorRequestWithDefaults instantiates a new AddThirdPartyPasswordGeneratorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyPasswordGeneratorRequestWithDefaults() *AddThirdPartyPasswordGeneratorRequest {
	this := AddThirdPartyPasswordGeneratorRequest{}
	return &this
}

// GetGeneratorName returns the GeneratorName field value
func (o *AddThirdPartyPasswordGeneratorRequest) GetGeneratorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GeneratorName
}

// GetGeneratorNameOk returns a tuple with the GeneratorName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordGeneratorRequest) GetGeneratorNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GeneratorName, true
}

// SetGeneratorName sets field value
func (o *AddThirdPartyPasswordGeneratorRequest) SetGeneratorName(v string) {
	o.GeneratorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyPasswordGeneratorRequest) GetSchemas() []EnumthirdPartyPasswordGeneratorSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyPasswordGeneratorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordGeneratorRequest) GetSchemasOk() ([]EnumthirdPartyPasswordGeneratorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyPasswordGeneratorRequest) SetSchemas(v []EnumthirdPartyPasswordGeneratorSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyPasswordGeneratorRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordGeneratorRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyPasswordGeneratorRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyPasswordGeneratorRequest) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordGeneratorRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyPasswordGeneratorRequest) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyPasswordGeneratorRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyPasswordGeneratorRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordGeneratorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyPasswordGeneratorRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyPasswordGeneratorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyPasswordGeneratorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyPasswordGeneratorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyPasswordGeneratorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddThirdPartyPasswordGeneratorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["generatorName"] = o.GeneratorName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["extensionClass"] = o.ExtensionClass
	}
	if !isNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddThirdPartyPasswordGeneratorRequest struct {
	value *AddThirdPartyPasswordGeneratorRequest
	isSet bool
}

func (v NullableAddThirdPartyPasswordGeneratorRequest) Get() *AddThirdPartyPasswordGeneratorRequest {
	return v.value
}

func (v *NullableAddThirdPartyPasswordGeneratorRequest) Set(val *AddThirdPartyPasswordGeneratorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyPasswordGeneratorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyPasswordGeneratorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyPasswordGeneratorRequest(val *AddThirdPartyPasswordGeneratorRequest) *NullableAddThirdPartyPasswordGeneratorRequest {
	return &NullableAddThirdPartyPasswordGeneratorRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyPasswordGeneratorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyPasswordGeneratorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


