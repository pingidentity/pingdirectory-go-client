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

// AddFileBasedCipherStreamProviderRequest struct for AddFileBasedCipherStreamProviderRequest
type AddFileBasedCipherStreamProviderRequest struct {
	// Name of the new Cipher Stream Provider
	ProviderName string `json:"providerName"`
	Schemas []EnumfileBasedCipherStreamProviderSchemaUrn `json:"schemas"`
	// The path to the file containing the password to use when generating ciphers.
	PasswordFile string `json:"passwordFile"`
	// Indicates whether the server should wait for the password file to become available if it does not exist.
	WaitForPasswordFile *bool `json:"waitForPasswordFile,omitempty"`
	// A description for this Cipher Stream Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Cipher Stream Provider is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
}

// NewAddFileBasedCipherStreamProviderRequest instantiates a new AddFileBasedCipherStreamProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFileBasedCipherStreamProviderRequest(providerName string, schemas []EnumfileBasedCipherStreamProviderSchemaUrn, passwordFile string, enabled bool) *AddFileBasedCipherStreamProviderRequest {
	this := AddFileBasedCipherStreamProviderRequest{}
	this.ProviderName = providerName
	this.Schemas = schemas
	this.PasswordFile = passwordFile
	this.Enabled = enabled
	return &this
}

// NewAddFileBasedCipherStreamProviderRequestWithDefaults instantiates a new AddFileBasedCipherStreamProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFileBasedCipherStreamProviderRequestWithDefaults() *AddFileBasedCipherStreamProviderRequest {
	this := AddFileBasedCipherStreamProviderRequest{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *AddFileBasedCipherStreamProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddFileBasedCipherStreamProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddFileBasedCipherStreamProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetSchemas returns the Schemas field value
func (o *AddFileBasedCipherStreamProviderRequest) GetSchemas() []EnumfileBasedCipherStreamProviderSchemaUrn {
	if o == nil {
		var ret []EnumfileBasedCipherStreamProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddFileBasedCipherStreamProviderRequest) GetSchemasOk() ([]EnumfileBasedCipherStreamProviderSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddFileBasedCipherStreamProviderRequest) SetSchemas(v []EnumfileBasedCipherStreamProviderSchemaUrn) {
	o.Schemas = v
}

// GetPasswordFile returns the PasswordFile field value
func (o *AddFileBasedCipherStreamProviderRequest) GetPasswordFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordFile
}

// GetPasswordFileOk returns a tuple with the PasswordFile field value
// and a boolean to check if the value has been set.
func (o *AddFileBasedCipherStreamProviderRequest) GetPasswordFileOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PasswordFile, true
}

// SetPasswordFile sets field value
func (o *AddFileBasedCipherStreamProviderRequest) SetPasswordFile(v string) {
	o.PasswordFile = v
}

// GetWaitForPasswordFile returns the WaitForPasswordFile field value if set, zero value otherwise.
func (o *AddFileBasedCipherStreamProviderRequest) GetWaitForPasswordFile() bool {
	if o == nil || isNil(o.WaitForPasswordFile) {
		var ret bool
		return ret
	}
	return *o.WaitForPasswordFile
}

// GetWaitForPasswordFileOk returns a tuple with the WaitForPasswordFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileBasedCipherStreamProviderRequest) GetWaitForPasswordFileOk() (*bool, bool) {
	if o == nil || isNil(o.WaitForPasswordFile) {
    return nil, false
	}
	return o.WaitForPasswordFile, true
}

// HasWaitForPasswordFile returns a boolean if a field has been set.
func (o *AddFileBasedCipherStreamProviderRequest) HasWaitForPasswordFile() bool {
	if o != nil && !isNil(o.WaitForPasswordFile) {
		return true
	}

	return false
}

// SetWaitForPasswordFile gets a reference to the given bool and assigns it to the WaitForPasswordFile field.
func (o *AddFileBasedCipherStreamProviderRequest) SetWaitForPasswordFile(v bool) {
	o.WaitForPasswordFile = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddFileBasedCipherStreamProviderRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileBasedCipherStreamProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddFileBasedCipherStreamProviderRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddFileBasedCipherStreamProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddFileBasedCipherStreamProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddFileBasedCipherStreamProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddFileBasedCipherStreamProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddFileBasedCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["providerName"] = o.ProviderName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["passwordFile"] = o.PasswordFile
	}
	if !isNil(o.WaitForPasswordFile) {
		toSerialize["waitForPasswordFile"] = o.WaitForPasswordFile
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddFileBasedCipherStreamProviderRequest struct {
	value *AddFileBasedCipherStreamProviderRequest
	isSet bool
}

func (v NullableAddFileBasedCipherStreamProviderRequest) Get() *AddFileBasedCipherStreamProviderRequest {
	return v.value
}

func (v *NullableAddFileBasedCipherStreamProviderRequest) Set(val *AddFileBasedCipherStreamProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFileBasedCipherStreamProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFileBasedCipherStreamProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFileBasedCipherStreamProviderRequest(val *AddFileBasedCipherStreamProviderRequest) *NullableAddFileBasedCipherStreamProviderRequest {
	return &NullableAddFileBasedCipherStreamProviderRequest{value: val, isSet: true}
}

func (v NullableAddFileBasedCipherStreamProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFileBasedCipherStreamProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

