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

// ThirdPartyKeyManagerProviderShared struct for ThirdPartyKeyManagerProviderShared
type ThirdPartyKeyManagerProviderShared struct {
	Schemas []EnumthirdPartyKeyManagerProviderSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Key Manager Provider.
	ExtensionClass string `json:"extensionClass"`
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Key Manager Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether the Key Manager Provider is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewThirdPartyKeyManagerProviderShared instantiates a new ThirdPartyKeyManagerProviderShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyKeyManagerProviderShared(schemas []EnumthirdPartyKeyManagerProviderSchemaUrn, extensionClass string, enabled bool) *ThirdPartyKeyManagerProviderShared {
	this := ThirdPartyKeyManagerProviderShared{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyKeyManagerProviderSharedWithDefaults instantiates a new ThirdPartyKeyManagerProviderShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyKeyManagerProviderSharedWithDefaults() *ThirdPartyKeyManagerProviderShared {
	this := ThirdPartyKeyManagerProviderShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyKeyManagerProviderShared) GetSchemas() []EnumthirdPartyKeyManagerProviderSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyKeyManagerProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyKeyManagerProviderShared) GetSchemasOk() ([]EnumthirdPartyKeyManagerProviderSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyKeyManagerProviderShared) SetSchemas(v []EnumthirdPartyKeyManagerProviderSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyKeyManagerProviderShared) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyKeyManagerProviderShared) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyKeyManagerProviderShared) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyKeyManagerProviderShared) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyKeyManagerProviderShared) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyKeyManagerProviderShared) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyKeyManagerProviderShared) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyKeyManagerProviderShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyKeyManagerProviderShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyKeyManagerProviderShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyKeyManagerProviderShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyKeyManagerProviderShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyKeyManagerProviderShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyKeyManagerProviderShared) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ThirdPartyKeyManagerProviderShared) MarshalJSON() ([]byte, error) {
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

type NullableThirdPartyKeyManagerProviderShared struct {
	value *ThirdPartyKeyManagerProviderShared
	isSet bool
}

func (v NullableThirdPartyKeyManagerProviderShared) Get() *ThirdPartyKeyManagerProviderShared {
	return v.value
}

func (v *NullableThirdPartyKeyManagerProviderShared) Set(val *ThirdPartyKeyManagerProviderShared) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyKeyManagerProviderShared) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyKeyManagerProviderShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyKeyManagerProviderShared(val *ThirdPartyKeyManagerProviderShared) *NullableThirdPartyKeyManagerProviderShared {
	return &NullableThirdPartyKeyManagerProviderShared{value: val, isSet: true}
}

func (v NullableThirdPartyKeyManagerProviderShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyKeyManagerProviderShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


