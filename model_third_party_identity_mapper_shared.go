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

// ThirdPartyIdentityMapperShared struct for ThirdPartyIdentityMapperShared
type ThirdPartyIdentityMapperShared struct {
	Schemas []EnumthirdPartyIdentityMapperSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Identity Mapper.
	ExtensionClass string `json:"extensionClass"`
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Identity Mapper
	Description *string `json:"description,omitempty"`
	// Indicates whether the Identity Mapper is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewThirdPartyIdentityMapperShared instantiates a new ThirdPartyIdentityMapperShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyIdentityMapperShared(schemas []EnumthirdPartyIdentityMapperSchemaUrn, extensionClass string, enabled bool) *ThirdPartyIdentityMapperShared {
	this := ThirdPartyIdentityMapperShared{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyIdentityMapperSharedWithDefaults instantiates a new ThirdPartyIdentityMapperShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyIdentityMapperSharedWithDefaults() *ThirdPartyIdentityMapperShared {
	this := ThirdPartyIdentityMapperShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyIdentityMapperShared) GetSchemas() []EnumthirdPartyIdentityMapperSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyIdentityMapperSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperShared) GetSchemasOk() ([]EnumthirdPartyIdentityMapperSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyIdentityMapperShared) SetSchemas(v []EnumthirdPartyIdentityMapperSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyIdentityMapperShared) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperShared) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyIdentityMapperShared) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyIdentityMapperShared) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperShared) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyIdentityMapperShared) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyIdentityMapperShared) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyIdentityMapperShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyIdentityMapperShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyIdentityMapperShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyIdentityMapperShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyIdentityMapperShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyIdentityMapperShared) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ThirdPartyIdentityMapperShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableThirdPartyIdentityMapperShared struct {
	value *ThirdPartyIdentityMapperShared
	isSet bool
}

func (v NullableThirdPartyIdentityMapperShared) Get() *ThirdPartyIdentityMapperShared {
	return v.value
}

func (v *NullableThirdPartyIdentityMapperShared) Set(val *ThirdPartyIdentityMapperShared) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyIdentityMapperShared) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyIdentityMapperShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyIdentityMapperShared(val *ThirdPartyIdentityMapperShared) *NullableThirdPartyIdentityMapperShared {
	return &NullableThirdPartyIdentityMapperShared{value: val, isSet: true}
}

func (v NullableThirdPartyIdentityMapperShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyIdentityMapperShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


