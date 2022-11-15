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

// Base64PasswordStorageSchemeResponse struct for Base64PasswordStorageSchemeResponse
type Base64PasswordStorageSchemeResponse struct {
	Schemas []Enumbase64PasswordStorageSchemeSchemaUrn `json:"schemas"`
	// Indicates whether the Base64 Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
}

// NewBase64PasswordStorageSchemeResponse instantiates a new Base64PasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBase64PasswordStorageSchemeResponse(schemas []Enumbase64PasswordStorageSchemeSchemaUrn, enabled bool) *Base64PasswordStorageSchemeResponse {
	this := Base64PasswordStorageSchemeResponse{}
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewBase64PasswordStorageSchemeResponseWithDefaults instantiates a new Base64PasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBase64PasswordStorageSchemeResponseWithDefaults() *Base64PasswordStorageSchemeResponse {
	this := Base64PasswordStorageSchemeResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *Base64PasswordStorageSchemeResponse) GetSchemas() []Enumbase64PasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []Enumbase64PasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *Base64PasswordStorageSchemeResponse) GetSchemasOk() ([]Enumbase64PasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *Base64PasswordStorageSchemeResponse) SetSchemas(v []Enumbase64PasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetEnabled returns the Enabled field value
func (o *Base64PasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Base64PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Base64PasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Base64PasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Base64PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Base64PasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Base64PasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

func (o Base64PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableBase64PasswordStorageSchemeResponse struct {
	value *Base64PasswordStorageSchemeResponse
	isSet bool
}

func (v NullableBase64PasswordStorageSchemeResponse) Get() *Base64PasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableBase64PasswordStorageSchemeResponse) Set(val *Base64PasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBase64PasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBase64PasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBase64PasswordStorageSchemeResponse(val *Base64PasswordStorageSchemeResponse) *NullableBase64PasswordStorageSchemeResponse {
	return &NullableBase64PasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableBase64PasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBase64PasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


