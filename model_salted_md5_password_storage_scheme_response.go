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

// SaltedMd5PasswordStorageSchemeResponse struct for SaltedMd5PasswordStorageSchemeResponse
type SaltedMd5PasswordStorageSchemeResponse struct {
	Schemas []EnumsaltedMd5PasswordStorageSchemeSchemaUrn `json:"schemas"`
	// Indicates whether the Salted MD5 Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies the number of bytes to use for the generated salt.
	SaltLengthBytes *int32 `json:"saltLengthBytes,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
}

// NewSaltedMd5PasswordStorageSchemeResponse instantiates a new SaltedMd5PasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaltedMd5PasswordStorageSchemeResponse(schemas []EnumsaltedMd5PasswordStorageSchemeSchemaUrn, enabled bool) *SaltedMd5PasswordStorageSchemeResponse {
	this := SaltedMd5PasswordStorageSchemeResponse{}
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewSaltedMd5PasswordStorageSchemeResponseWithDefaults instantiates a new SaltedMd5PasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaltedMd5PasswordStorageSchemeResponseWithDefaults() *SaltedMd5PasswordStorageSchemeResponse {
	this := SaltedMd5PasswordStorageSchemeResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *SaltedMd5PasswordStorageSchemeResponse) GetSchemas() []EnumsaltedMd5PasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumsaltedMd5PasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SaltedMd5PasswordStorageSchemeResponse) GetSchemasOk() ([]EnumsaltedMd5PasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SaltedMd5PasswordStorageSchemeResponse) SetSchemas(v []EnumsaltedMd5PasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetEnabled returns the Enabled field value
func (o *SaltedMd5PasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SaltedMd5PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SaltedMd5PasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSaltLengthBytes returns the SaltLengthBytes field value if set, zero value otherwise.
func (o *SaltedMd5PasswordStorageSchemeResponse) GetSaltLengthBytes() int32 {
	if o == nil || isNil(o.SaltLengthBytes) {
		var ret int32
		return ret
	}
	return *o.SaltLengthBytes
}

// GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaltedMd5PasswordStorageSchemeResponse) GetSaltLengthBytesOk() (*int32, bool) {
	if o == nil || isNil(o.SaltLengthBytes) {
    return nil, false
	}
	return o.SaltLengthBytes, true
}

// HasSaltLengthBytes returns a boolean if a field has been set.
func (o *SaltedMd5PasswordStorageSchemeResponse) HasSaltLengthBytes() bool {
	if o != nil && !isNil(o.SaltLengthBytes) {
		return true
	}

	return false
}

// SetSaltLengthBytes gets a reference to the given int32 and assigns it to the SaltLengthBytes field.
func (o *SaltedMd5PasswordStorageSchemeResponse) SetSaltLengthBytes(v int32) {
	o.SaltLengthBytes = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SaltedMd5PasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaltedMd5PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SaltedMd5PasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SaltedMd5PasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

func (o SaltedMd5PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.SaltLengthBytes) {
		toSerialize["saltLengthBytes"] = o.SaltLengthBytes
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableSaltedMd5PasswordStorageSchemeResponse struct {
	value *SaltedMd5PasswordStorageSchemeResponse
	isSet bool
}

func (v NullableSaltedMd5PasswordStorageSchemeResponse) Get() *SaltedMd5PasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableSaltedMd5PasswordStorageSchemeResponse) Set(val *SaltedMd5PasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSaltedMd5PasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSaltedMd5PasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaltedMd5PasswordStorageSchemeResponse(val *SaltedMd5PasswordStorageSchemeResponse) *NullableSaltedMd5PasswordStorageSchemeResponse {
	return &NullableSaltedMd5PasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableSaltedMd5PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaltedMd5PasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


