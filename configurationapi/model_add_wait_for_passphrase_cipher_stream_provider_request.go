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

// checks if the AddWaitForPassphraseCipherStreamProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddWaitForPassphraseCipherStreamProviderRequest{}

// AddWaitForPassphraseCipherStreamProviderRequest struct for AddWaitForPassphraseCipherStreamProviderRequest
type AddWaitForPassphraseCipherStreamProviderRequest struct {
	Schemas []EnumwaitForPassphraseCipherStreamProviderSchemaUrn `json:"schemas"`
	// A description for this Cipher Stream Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
	// Name of the new Cipher Stream Provider
	ProviderName string `json:"providerName"`
}

// NewAddWaitForPassphraseCipherStreamProviderRequest instantiates a new AddWaitForPassphraseCipherStreamProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddWaitForPassphraseCipherStreamProviderRequest(schemas []EnumwaitForPassphraseCipherStreamProviderSchemaUrn, enabled bool, providerName string) *AddWaitForPassphraseCipherStreamProviderRequest {
	this := AddWaitForPassphraseCipherStreamProviderRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.ProviderName = providerName
	return &this
}

// NewAddWaitForPassphraseCipherStreamProviderRequestWithDefaults instantiates a new AddWaitForPassphraseCipherStreamProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddWaitForPassphraseCipherStreamProviderRequestWithDefaults() *AddWaitForPassphraseCipherStreamProviderRequest {
	this := AddWaitForPassphraseCipherStreamProviderRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetSchemas() []EnumwaitForPassphraseCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumwaitForPassphraseCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetSchemasOk() ([]EnumwaitForPassphraseCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddWaitForPassphraseCipherStreamProviderRequest) SetSchemas(v []EnumwaitForPassphraseCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddWaitForPassphraseCipherStreamProviderRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddWaitForPassphraseCipherStreamProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddWaitForPassphraseCipherStreamProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProviderName returns the ProviderName field value
func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddWaitForPassphraseCipherStreamProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddWaitForPassphraseCipherStreamProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

func (o AddWaitForPassphraseCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddWaitForPassphraseCipherStreamProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableAddWaitForPassphraseCipherStreamProviderRequest struct {
	value *AddWaitForPassphraseCipherStreamProviderRequest
	isSet bool
}

func (v NullableAddWaitForPassphraseCipherStreamProviderRequest) Get() *AddWaitForPassphraseCipherStreamProviderRequest {
	return v.value
}

func (v *NullableAddWaitForPassphraseCipherStreamProviderRequest) Set(val *AddWaitForPassphraseCipherStreamProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddWaitForPassphraseCipherStreamProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddWaitForPassphraseCipherStreamProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddWaitForPassphraseCipherStreamProviderRequest(val *AddWaitForPassphraseCipherStreamProviderRequest) *NullableAddWaitForPassphraseCipherStreamProviderRequest {
	return &NullableAddWaitForPassphraseCipherStreamProviderRequest{value: val, isSet: true}
}

func (v NullableAddWaitForPassphraseCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddWaitForPassphraseCipherStreamProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
