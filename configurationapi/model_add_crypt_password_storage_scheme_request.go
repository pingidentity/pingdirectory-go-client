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

// checks if the AddCryptPasswordStorageSchemeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCryptPasswordStorageSchemeRequest{}

// AddCryptPasswordStorageSchemeRequest struct for AddCryptPasswordStorageSchemeRequest
type AddCryptPasswordStorageSchemeRequest struct {
	Schemas                   []EnumcryptPasswordStorageSchemeSchemaUrn               `json:"schemas"`
	PasswordEncodingMechanism *EnumpasswordStorageSchemePasswordEncodingMechanismProp `json:"passwordEncodingMechanism,omitempty"`
	// Specifies the number of digest rounds to use for the SHA-2 encodings. This will not be used for the legacy or MD5-based encodings.
	NumDigestRounds *int64 `json:"numDigestRounds,omitempty"`
	// Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords.
	MaxPasswordLength *int64 `json:"maxPasswordLength,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
	// Name of the new Password Storage Scheme
	SchemeName string `json:"schemeName"`
}

// NewAddCryptPasswordStorageSchemeRequest instantiates a new AddCryptPasswordStorageSchemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCryptPasswordStorageSchemeRequest(schemas []EnumcryptPasswordStorageSchemeSchemaUrn, enabled bool, schemeName string) *AddCryptPasswordStorageSchemeRequest {
	this := AddCryptPasswordStorageSchemeRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.SchemeName = schemeName
	return &this
}

// NewAddCryptPasswordStorageSchemeRequestWithDefaults instantiates a new AddCryptPasswordStorageSchemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCryptPasswordStorageSchemeRequestWithDefaults() *AddCryptPasswordStorageSchemeRequest {
	this := AddCryptPasswordStorageSchemeRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddCryptPasswordStorageSchemeRequest) GetSchemas() []EnumcryptPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumcryptPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddCryptPasswordStorageSchemeRequest) GetSchemasOk() ([]EnumcryptPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddCryptPasswordStorageSchemeRequest) SetSchemas(v []EnumcryptPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetPasswordEncodingMechanism returns the PasswordEncodingMechanism field value if set, zero value otherwise.
func (o *AddCryptPasswordStorageSchemeRequest) GetPasswordEncodingMechanism() EnumpasswordStorageSchemePasswordEncodingMechanismProp {
	if o == nil || IsNil(o.PasswordEncodingMechanism) {
		var ret EnumpasswordStorageSchemePasswordEncodingMechanismProp
		return ret
	}
	return *o.PasswordEncodingMechanism
}

// GetPasswordEncodingMechanismOk returns a tuple with the PasswordEncodingMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCryptPasswordStorageSchemeRequest) GetPasswordEncodingMechanismOk() (*EnumpasswordStorageSchemePasswordEncodingMechanismProp, bool) {
	if o == nil || IsNil(o.PasswordEncodingMechanism) {
		return nil, false
	}
	return o.PasswordEncodingMechanism, true
}

// HasPasswordEncodingMechanism returns a boolean if a field has been set.
func (o *AddCryptPasswordStorageSchemeRequest) HasPasswordEncodingMechanism() bool {
	if o != nil && !IsNil(o.PasswordEncodingMechanism) {
		return true
	}

	return false
}

// SetPasswordEncodingMechanism gets a reference to the given EnumpasswordStorageSchemePasswordEncodingMechanismProp and assigns it to the PasswordEncodingMechanism field.
func (o *AddCryptPasswordStorageSchemeRequest) SetPasswordEncodingMechanism(v EnumpasswordStorageSchemePasswordEncodingMechanismProp) {
	o.PasswordEncodingMechanism = &v
}

// GetNumDigestRounds returns the NumDigestRounds field value if set, zero value otherwise.
func (o *AddCryptPasswordStorageSchemeRequest) GetNumDigestRounds() int64 {
	if o == nil || IsNil(o.NumDigestRounds) {
		var ret int64
		return ret
	}
	return *o.NumDigestRounds
}

// GetNumDigestRoundsOk returns a tuple with the NumDigestRounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCryptPasswordStorageSchemeRequest) GetNumDigestRoundsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumDigestRounds) {
		return nil, false
	}
	return o.NumDigestRounds, true
}

// HasNumDigestRounds returns a boolean if a field has been set.
func (o *AddCryptPasswordStorageSchemeRequest) HasNumDigestRounds() bool {
	if o != nil && !IsNil(o.NumDigestRounds) {
		return true
	}

	return false
}

// SetNumDigestRounds gets a reference to the given int64 and assigns it to the NumDigestRounds field.
func (o *AddCryptPasswordStorageSchemeRequest) SetNumDigestRounds(v int64) {
	o.NumDigestRounds = &v
}

// GetMaxPasswordLength returns the MaxPasswordLength field value if set, zero value otherwise.
func (o *AddCryptPasswordStorageSchemeRequest) GetMaxPasswordLength() int64 {
	if o == nil || IsNil(o.MaxPasswordLength) {
		var ret int64
		return ret
	}
	return *o.MaxPasswordLength
}

// GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCryptPasswordStorageSchemeRequest) GetMaxPasswordLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxPasswordLength) {
		return nil, false
	}
	return o.MaxPasswordLength, true
}

// HasMaxPasswordLength returns a boolean if a field has been set.
func (o *AddCryptPasswordStorageSchemeRequest) HasMaxPasswordLength() bool {
	if o != nil && !IsNil(o.MaxPasswordLength) {
		return true
	}

	return false
}

// SetMaxPasswordLength gets a reference to the given int64 and assigns it to the MaxPasswordLength field.
func (o *AddCryptPasswordStorageSchemeRequest) SetMaxPasswordLength(v int64) {
	o.MaxPasswordLength = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddCryptPasswordStorageSchemeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCryptPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddCryptPasswordStorageSchemeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddCryptPasswordStorageSchemeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddCryptPasswordStorageSchemeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddCryptPasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddCryptPasswordStorageSchemeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSchemeName returns the SchemeName field value
func (o *AddCryptPasswordStorageSchemeRequest) GetSchemeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemeName
}

// GetSchemeNameOk returns a tuple with the SchemeName field value
// and a boolean to check if the value has been set.
func (o *AddCryptPasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemeName, true
}

// SetSchemeName sets field value
func (o *AddCryptPasswordStorageSchemeRequest) SetSchemeName(v string) {
	o.SchemeName = v
}

func (o AddCryptPasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCryptPasswordStorageSchemeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.PasswordEncodingMechanism) {
		toSerialize["passwordEncodingMechanism"] = o.PasswordEncodingMechanism
	}
	if !IsNil(o.NumDigestRounds) {
		toSerialize["numDigestRounds"] = o.NumDigestRounds
	}
	if !IsNil(o.MaxPasswordLength) {
		toSerialize["maxPasswordLength"] = o.MaxPasswordLength
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["schemeName"] = o.SchemeName
	return toSerialize, nil
}

type NullableAddCryptPasswordStorageSchemeRequest struct {
	value *AddCryptPasswordStorageSchemeRequest
	isSet bool
}

func (v NullableAddCryptPasswordStorageSchemeRequest) Get() *AddCryptPasswordStorageSchemeRequest {
	return v.value
}

func (v *NullableAddCryptPasswordStorageSchemeRequest) Set(val *AddCryptPasswordStorageSchemeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCryptPasswordStorageSchemeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCryptPasswordStorageSchemeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCryptPasswordStorageSchemeRequest(val *AddCryptPasswordStorageSchemeRequest) *NullableAddCryptPasswordStorageSchemeRequest {
	return &NullableAddCryptPasswordStorageSchemeRequest{value: val, isSet: true}
}

func (v NullableAddCryptPasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCryptPasswordStorageSchemeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
