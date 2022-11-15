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

// ThirdPartyEnhancedPasswordStorageSchemeShared struct for ThirdPartyEnhancedPasswordStorageSchemeShared
type ThirdPartyEnhancedPasswordStorageSchemeShared struct {
	Schemas []EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Enhanced Password Storage Scheme.
	ExtensionClass string `json:"extensionClass"`
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewThirdPartyEnhancedPasswordStorageSchemeShared instantiates a new ThirdPartyEnhancedPasswordStorageSchemeShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyEnhancedPasswordStorageSchemeShared(schemas []EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn, extensionClass string, enabled bool) *ThirdPartyEnhancedPasswordStorageSchemeShared {
	this := ThirdPartyEnhancedPasswordStorageSchemeShared{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyEnhancedPasswordStorageSchemeSharedWithDefaults instantiates a new ThirdPartyEnhancedPasswordStorageSchemeShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyEnhancedPasswordStorageSchemeSharedWithDefaults() *ThirdPartyEnhancedPasswordStorageSchemeShared {
	this := ThirdPartyEnhancedPasswordStorageSchemeShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) GetSchemas() []EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) GetSchemasOk() ([]EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) SetSchemas(v []EnumthirdPartyEnhancedPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyEnhancedPasswordStorageSchemeShared) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ThirdPartyEnhancedPasswordStorageSchemeShared) MarshalJSON() ([]byte, error) {
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

type NullableThirdPartyEnhancedPasswordStorageSchemeShared struct {
	value *ThirdPartyEnhancedPasswordStorageSchemeShared
	isSet bool
}

func (v NullableThirdPartyEnhancedPasswordStorageSchemeShared) Get() *ThirdPartyEnhancedPasswordStorageSchemeShared {
	return v.value
}

func (v *NullableThirdPartyEnhancedPasswordStorageSchemeShared) Set(val *ThirdPartyEnhancedPasswordStorageSchemeShared) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyEnhancedPasswordStorageSchemeShared) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyEnhancedPasswordStorageSchemeShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyEnhancedPasswordStorageSchemeShared(val *ThirdPartyEnhancedPasswordStorageSchemeShared) *NullableThirdPartyEnhancedPasswordStorageSchemeShared {
	return &NullableThirdPartyEnhancedPasswordStorageSchemeShared{value: val, isSet: true}
}

func (v NullableThirdPartyEnhancedPasswordStorageSchemeShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyEnhancedPasswordStorageSchemeShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


