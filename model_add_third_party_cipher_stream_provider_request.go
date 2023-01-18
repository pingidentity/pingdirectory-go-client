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

// AddThirdPartyCipherStreamProviderRequest struct for AddThirdPartyCipherStreamProviderRequest
type AddThirdPartyCipherStreamProviderRequest struct {
	// Name of the new Cipher Stream Provider
	ProviderName string `json:"providerName"`
	Schemas []EnumthirdPartyCipherStreamProviderSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Cipher Stream Provider.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Cipher Stream Provider. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Cipher Stream Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
}

// NewAddThirdPartyCipherStreamProviderRequest instantiates a new AddThirdPartyCipherStreamProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyCipherStreamProviderRequest(providerName string, schemas []EnumthirdPartyCipherStreamProviderSchemaUrn, extensionClass string, enabled bool) *AddThirdPartyCipherStreamProviderRequest {
	this := AddThirdPartyCipherStreamProviderRequest{}
	this.ProviderName = providerName
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewAddThirdPartyCipherStreamProviderRequestWithDefaults instantiates a new AddThirdPartyCipherStreamProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyCipherStreamProviderRequestWithDefaults() *AddThirdPartyCipherStreamProviderRequest {
	this := AddThirdPartyCipherStreamProviderRequest{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *AddThirdPartyCipherStreamProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyCipherStreamProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddThirdPartyCipherStreamProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyCipherStreamProviderRequest) GetSchemas() []EnumthirdPartyCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyCipherStreamProviderRequest) GetSchemasOk() ([]EnumthirdPartyCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyCipherStreamProviderRequest) SetSchemas(v []EnumthirdPartyCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyCipherStreamProviderRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyCipherStreamProviderRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyCipherStreamProviderRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyCipherStreamProviderRequest) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyCipherStreamProviderRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyCipherStreamProviderRequest) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyCipherStreamProviderRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyCipherStreamProviderRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyCipherStreamProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyCipherStreamProviderRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyCipherStreamProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyCipherStreamProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyCipherStreamProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyCipherStreamProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddThirdPartyCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["providerName"] = o.ProviderName
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

type NullableAddThirdPartyCipherStreamProviderRequest struct {
	value *AddThirdPartyCipherStreamProviderRequest
	isSet bool
}

func (v NullableAddThirdPartyCipherStreamProviderRequest) Get() *AddThirdPartyCipherStreamProviderRequest {
	return v.value
}

func (v *NullableAddThirdPartyCipherStreamProviderRequest) Set(val *AddThirdPartyCipherStreamProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyCipherStreamProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyCipherStreamProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyCipherStreamProviderRequest(val *AddThirdPartyCipherStreamProviderRequest) *NullableAddThirdPartyCipherStreamProviderRequest {
	return &NullableAddThirdPartyCipherStreamProviderRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyCipherStreamProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

