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

// AddCryptPasswordStorageSchemeRequest struct for AddCryptPasswordStorageSchemeRequest
type AddCryptPasswordStorageSchemeRequest struct {
	// Name of the new Password Storage Scheme
	SchemeName string `json:"schemeName"`
	Schemas []EnumcryptPasswordStorageSchemeSchemaUrn `json:"schemas"`
	PasswordEncodingMechanism *EnumpasswordStorageSchemePasswordEncodingMechanismProp `json:"passwordEncodingMechanism,omitempty"`
	// Specifies the number of digest rounds to use for the SHA-2 encodings. This will not be used for the legacy or MD5-based encodings.
	NumDigestRounds *int32 `json:"numDigestRounds,omitempty"`
	// Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords.
	MaxPasswordLength *int32 `json:"maxPasswordLength,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddCryptPasswordStorageSchemeRequest instantiates a new AddCryptPasswordStorageSchemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCryptPasswordStorageSchemeRequest(schemeName string, schemas []EnumcryptPasswordStorageSchemeSchemaUrn, enabled bool) *AddCryptPasswordStorageSchemeRequest {
	this := AddCryptPasswordStorageSchemeRequest{}
	this.SchemeName = schemeName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddCryptPasswordStorageSchemeRequestWithDefaults instantiates a new AddCryptPasswordStorageSchemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCryptPasswordStorageSchemeRequestWithDefaults() *AddCryptPasswordStorageSchemeRequest {
	this := AddCryptPasswordStorageSchemeRequest{}
	return &this
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
	if o == nil || isNil(o.PasswordEncodingMechanism) {
		var ret EnumpasswordStorageSchemePasswordEncodingMechanismProp
		return ret
	}
	return *o.PasswordEncodingMechanism
}

// GetPasswordEncodingMechanismOk returns a tuple with the PasswordEncodingMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCryptPasswordStorageSchemeRequest) GetPasswordEncodingMechanismOk() (*EnumpasswordStorageSchemePasswordEncodingMechanismProp, bool) {
	if o == nil || isNil(o.PasswordEncodingMechanism) {
    return nil, false
	}
	return o.PasswordEncodingMechanism, true
}

// HasPasswordEncodingMechanism returns a boolean if a field has been set.
func (o *AddCryptPasswordStorageSchemeRequest) HasPasswordEncodingMechanism() bool {
	if o != nil && !isNil(o.PasswordEncodingMechanism) {
		return true
	}

	return false
}

// SetPasswordEncodingMechanism gets a reference to the given EnumpasswordStorageSchemePasswordEncodingMechanismProp and assigns it to the PasswordEncodingMechanism field.
func (o *AddCryptPasswordStorageSchemeRequest) SetPasswordEncodingMechanism(v EnumpasswordStorageSchemePasswordEncodingMechanismProp) {
	o.PasswordEncodingMechanism = &v
}

// GetNumDigestRounds returns the NumDigestRounds field value if set, zero value otherwise.
func (o *AddCryptPasswordStorageSchemeRequest) GetNumDigestRounds() int32 {
	if o == nil || isNil(o.NumDigestRounds) {
		var ret int32
		return ret
	}
	return *o.NumDigestRounds
}

// GetNumDigestRoundsOk returns a tuple with the NumDigestRounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCryptPasswordStorageSchemeRequest) GetNumDigestRoundsOk() (*int32, bool) {
	if o == nil || isNil(o.NumDigestRounds) {
    return nil, false
	}
	return o.NumDigestRounds, true
}

// HasNumDigestRounds returns a boolean if a field has been set.
func (o *AddCryptPasswordStorageSchemeRequest) HasNumDigestRounds() bool {
	if o != nil && !isNil(o.NumDigestRounds) {
		return true
	}

	return false
}

// SetNumDigestRounds gets a reference to the given int32 and assigns it to the NumDigestRounds field.
func (o *AddCryptPasswordStorageSchemeRequest) SetNumDigestRounds(v int32) {
	o.NumDigestRounds = &v
}

// GetMaxPasswordLength returns the MaxPasswordLength field value if set, zero value otherwise.
func (o *AddCryptPasswordStorageSchemeRequest) GetMaxPasswordLength() int32 {
	if o == nil || isNil(o.MaxPasswordLength) {
		var ret int32
		return ret
	}
	return *o.MaxPasswordLength
}

// GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCryptPasswordStorageSchemeRequest) GetMaxPasswordLengthOk() (*int32, bool) {
	if o == nil || isNil(o.MaxPasswordLength) {
    return nil, false
	}
	return o.MaxPasswordLength, true
}

// HasMaxPasswordLength returns a boolean if a field has been set.
func (o *AddCryptPasswordStorageSchemeRequest) HasMaxPasswordLength() bool {
	if o != nil && !isNil(o.MaxPasswordLength) {
		return true
	}

	return false
}

// SetMaxPasswordLength gets a reference to the given int32 and assigns it to the MaxPasswordLength field.
func (o *AddCryptPasswordStorageSchemeRequest) SetMaxPasswordLength(v int32) {
	o.MaxPasswordLength = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddCryptPasswordStorageSchemeRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCryptPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddCryptPasswordStorageSchemeRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
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

func (o AddCryptPasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemeName"] = o.SchemeName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.PasswordEncodingMechanism) {
		toSerialize["passwordEncodingMechanism"] = o.PasswordEncodingMechanism
	}
	if !isNil(o.NumDigestRounds) {
		toSerialize["numDigestRounds"] = o.NumDigestRounds
	}
	if !isNil(o.MaxPasswordLength) {
		toSerialize["maxPasswordLength"] = o.MaxPasswordLength
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
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

