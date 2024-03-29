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

// checks if the AddEnvironmentVariablePassphraseProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddEnvironmentVariablePassphraseProviderRequest{}

// AddEnvironmentVariablePassphraseProviderRequest struct for AddEnvironmentVariablePassphraseProviderRequest
type AddEnvironmentVariablePassphraseProviderRequest struct {
	Schemas []EnumenvironmentVariablePassphraseProviderSchemaUrn `json:"schemas"`
	// The name of the environment variable that is expected to hold the passphrase.
	EnvironmentVariable string `json:"environmentVariable"`
	// A description for this Passphrase Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Passphrase Provider is enabled for use in the server.
	Enabled bool `json:"enabled"`
	// Name of the new Passphrase Provider
	ProviderName string `json:"providerName"`
}

// NewAddEnvironmentVariablePassphraseProviderRequest instantiates a new AddEnvironmentVariablePassphraseProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddEnvironmentVariablePassphraseProviderRequest(schemas []EnumenvironmentVariablePassphraseProviderSchemaUrn, environmentVariable string, enabled bool, providerName string) *AddEnvironmentVariablePassphraseProviderRequest {
	this := AddEnvironmentVariablePassphraseProviderRequest{}
	this.Schemas = schemas
	this.EnvironmentVariable = environmentVariable
	this.Enabled = enabled
	this.ProviderName = providerName
	return &this
}

// NewAddEnvironmentVariablePassphraseProviderRequestWithDefaults instantiates a new AddEnvironmentVariablePassphraseProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddEnvironmentVariablePassphraseProviderRequestWithDefaults() *AddEnvironmentVariablePassphraseProviderRequest {
	this := AddEnvironmentVariablePassphraseProviderRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddEnvironmentVariablePassphraseProviderRequest) GetSchemas() []EnumenvironmentVariablePassphraseProviderSchemaUrn {
	if o == nil {
		var ret []EnumenvironmentVariablePassphraseProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddEnvironmentVariablePassphraseProviderRequest) GetSchemasOk() ([]EnumenvironmentVariablePassphraseProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddEnvironmentVariablePassphraseProviderRequest) SetSchemas(v []EnumenvironmentVariablePassphraseProviderSchemaUrn) {
	o.Schemas = v
}

// GetEnvironmentVariable returns the EnvironmentVariable field value
func (o *AddEnvironmentVariablePassphraseProviderRequest) GetEnvironmentVariable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentVariable
}

// GetEnvironmentVariableOk returns a tuple with the EnvironmentVariable field value
// and a boolean to check if the value has been set.
func (o *AddEnvironmentVariablePassphraseProviderRequest) GetEnvironmentVariableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentVariable, true
}

// SetEnvironmentVariable sets field value
func (o *AddEnvironmentVariablePassphraseProviderRequest) SetEnvironmentVariable(v string) {
	o.EnvironmentVariable = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddEnvironmentVariablePassphraseProviderRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEnvironmentVariablePassphraseProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddEnvironmentVariablePassphraseProviderRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddEnvironmentVariablePassphraseProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddEnvironmentVariablePassphraseProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddEnvironmentVariablePassphraseProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddEnvironmentVariablePassphraseProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProviderName returns the ProviderName field value
func (o *AddEnvironmentVariablePassphraseProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddEnvironmentVariablePassphraseProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddEnvironmentVariablePassphraseProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

func (o AddEnvironmentVariablePassphraseProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddEnvironmentVariablePassphraseProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["environmentVariable"] = o.EnvironmentVariable
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableAddEnvironmentVariablePassphraseProviderRequest struct {
	value *AddEnvironmentVariablePassphraseProviderRequest
	isSet bool
}

func (v NullableAddEnvironmentVariablePassphraseProviderRequest) Get() *AddEnvironmentVariablePassphraseProviderRequest {
	return v.value
}

func (v *NullableAddEnvironmentVariablePassphraseProviderRequest) Set(val *AddEnvironmentVariablePassphraseProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddEnvironmentVariablePassphraseProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddEnvironmentVariablePassphraseProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddEnvironmentVariablePassphraseProviderRequest(val *AddEnvironmentVariablePassphraseProviderRequest) *NullableAddEnvironmentVariablePassphraseProviderRequest {
	return &NullableAddEnvironmentVariablePassphraseProviderRequest{value: val, isSet: true}
}

func (v NullableAddEnvironmentVariablePassphraseProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddEnvironmentVariablePassphraseProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
