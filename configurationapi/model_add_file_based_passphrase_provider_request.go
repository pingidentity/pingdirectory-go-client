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

// checks if the AddFileBasedPassphraseProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddFileBasedPassphraseProviderRequest{}

// AddFileBasedPassphraseProviderRequest struct for AddFileBasedPassphraseProviderRequest
type AddFileBasedPassphraseProviderRequest struct {
	Schemas []EnumfileBasedPassphraseProviderSchemaUrn `json:"schemas"`
	// The path to the file containing the passphrase.
	PasswordFile string `json:"passwordFile"`
	// The maximum length of time that the passphrase provider may cache the passphrase that has been read from the target file. A value of zero seconds indicates that the provider should always attempt to read the passphrase from the file.
	MaxCacheDuration *string `json:"maxCacheDuration,omitempty"`
	// A description for this Passphrase Provider
	Description *string `json:"description,omitempty"`
	// Indicates whether this Passphrase Provider is enabled for use in the server.
	Enabled bool `json:"enabled"`
	// Name of the new Passphrase Provider
	ProviderName string `json:"providerName"`
}

// NewAddFileBasedPassphraseProviderRequest instantiates a new AddFileBasedPassphraseProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFileBasedPassphraseProviderRequest(schemas []EnumfileBasedPassphraseProviderSchemaUrn, passwordFile string, enabled bool, providerName string) *AddFileBasedPassphraseProviderRequest {
	this := AddFileBasedPassphraseProviderRequest{}
	this.Schemas = schemas
	this.PasswordFile = passwordFile
	this.Enabled = enabled
	this.ProviderName = providerName
	return &this
}

// NewAddFileBasedPassphraseProviderRequestWithDefaults instantiates a new AddFileBasedPassphraseProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFileBasedPassphraseProviderRequestWithDefaults() *AddFileBasedPassphraseProviderRequest {
	this := AddFileBasedPassphraseProviderRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddFileBasedPassphraseProviderRequest) GetSchemas() []EnumfileBasedPassphraseProviderSchemaUrn {
	if o == nil {
		var ret []EnumfileBasedPassphraseProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddFileBasedPassphraseProviderRequest) GetSchemasOk() ([]EnumfileBasedPassphraseProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddFileBasedPassphraseProviderRequest) SetSchemas(v []EnumfileBasedPassphraseProviderSchemaUrn) {
	o.Schemas = v
}

// GetPasswordFile returns the PasswordFile field value
func (o *AddFileBasedPassphraseProviderRequest) GetPasswordFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordFile
}

// GetPasswordFileOk returns a tuple with the PasswordFile field value
// and a boolean to check if the value has been set.
func (o *AddFileBasedPassphraseProviderRequest) GetPasswordFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordFile, true
}

// SetPasswordFile sets field value
func (o *AddFileBasedPassphraseProviderRequest) SetPasswordFile(v string) {
	o.PasswordFile = v
}

// GetMaxCacheDuration returns the MaxCacheDuration field value if set, zero value otherwise.
func (o *AddFileBasedPassphraseProviderRequest) GetMaxCacheDuration() string {
	if o == nil || IsNil(o.MaxCacheDuration) {
		var ret string
		return ret
	}
	return *o.MaxCacheDuration
}

// GetMaxCacheDurationOk returns a tuple with the MaxCacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileBasedPassphraseProviderRequest) GetMaxCacheDurationOk() (*string, bool) {
	if o == nil || IsNil(o.MaxCacheDuration) {
		return nil, false
	}
	return o.MaxCacheDuration, true
}

// HasMaxCacheDuration returns a boolean if a field has been set.
func (o *AddFileBasedPassphraseProviderRequest) HasMaxCacheDuration() bool {
	if o != nil && !IsNil(o.MaxCacheDuration) {
		return true
	}

	return false
}

// SetMaxCacheDuration gets a reference to the given string and assigns it to the MaxCacheDuration field.
func (o *AddFileBasedPassphraseProviderRequest) SetMaxCacheDuration(v string) {
	o.MaxCacheDuration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddFileBasedPassphraseProviderRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFileBasedPassphraseProviderRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddFileBasedPassphraseProviderRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddFileBasedPassphraseProviderRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddFileBasedPassphraseProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddFileBasedPassphraseProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddFileBasedPassphraseProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProviderName returns the ProviderName field value
func (o *AddFileBasedPassphraseProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddFileBasedPassphraseProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddFileBasedPassphraseProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

func (o AddFileBasedPassphraseProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddFileBasedPassphraseProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["passwordFile"] = o.PasswordFile
	if !IsNil(o.MaxCacheDuration) {
		toSerialize["maxCacheDuration"] = o.MaxCacheDuration
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableAddFileBasedPassphraseProviderRequest struct {
	value *AddFileBasedPassphraseProviderRequest
	isSet bool
}

func (v NullableAddFileBasedPassphraseProviderRequest) Get() *AddFileBasedPassphraseProviderRequest {
	return v.value
}

func (v *NullableAddFileBasedPassphraseProviderRequest) Set(val *AddFileBasedPassphraseProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFileBasedPassphraseProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFileBasedPassphraseProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFileBasedPassphraseProviderRequest(val *AddFileBasedPassphraseProviderRequest) *NullableAddFileBasedPassphraseProviderRequest {
	return &NullableAddFileBasedPassphraseProviderRequest{value: val, isSet: true}
}

func (v NullableAddFileBasedPassphraseProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFileBasedPassphraseProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
