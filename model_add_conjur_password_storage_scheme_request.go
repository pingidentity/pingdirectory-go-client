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

// AddConjurPasswordStorageSchemeRequest struct for AddConjurPasswordStorageSchemeRequest
type AddConjurPasswordStorageSchemeRequest struct {
	// Name of the new Password Storage Scheme
	SchemeName string                                     `json:"schemeName"`
	Schemas    []EnumconjurPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// An external server definition with information needed to connect and authenticate to the Conjur instance containing user passwords.
	ConjurExternalServer string `json:"conjurExternalServer"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddConjurPasswordStorageSchemeRequest instantiates a new AddConjurPasswordStorageSchemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddConjurPasswordStorageSchemeRequest(schemeName string, schemas []EnumconjurPasswordStorageSchemeSchemaUrn, conjurExternalServer string, enabled bool) *AddConjurPasswordStorageSchemeRequest {
	this := AddConjurPasswordStorageSchemeRequest{}
	this.SchemeName = schemeName
	this.Schemas = schemas
	this.ConjurExternalServer = conjurExternalServer
	this.Enabled = enabled
	return &this
}

// NewAddConjurPasswordStorageSchemeRequestWithDefaults instantiates a new AddConjurPasswordStorageSchemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddConjurPasswordStorageSchemeRequestWithDefaults() *AddConjurPasswordStorageSchemeRequest {
	this := AddConjurPasswordStorageSchemeRequest{}
	return &this
}

// GetSchemeName returns the SchemeName field value
func (o *AddConjurPasswordStorageSchemeRequest) GetSchemeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemeName
}

// GetSchemeNameOk returns a tuple with the SchemeName field value
// and a boolean to check if the value has been set.
func (o *AddConjurPasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemeName, true
}

// SetSchemeName sets field value
func (o *AddConjurPasswordStorageSchemeRequest) SetSchemeName(v string) {
	o.SchemeName = v
}

// GetSchemas returns the Schemas field value
func (o *AddConjurPasswordStorageSchemeRequest) GetSchemas() []EnumconjurPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumconjurPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddConjurPasswordStorageSchemeRequest) GetSchemasOk() ([]EnumconjurPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddConjurPasswordStorageSchemeRequest) SetSchemas(v []EnumconjurPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetConjurExternalServer returns the ConjurExternalServer field value
func (o *AddConjurPasswordStorageSchemeRequest) GetConjurExternalServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConjurExternalServer
}

// GetConjurExternalServerOk returns a tuple with the ConjurExternalServer field value
// and a boolean to check if the value has been set.
func (o *AddConjurPasswordStorageSchemeRequest) GetConjurExternalServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConjurExternalServer, true
}

// SetConjurExternalServer sets field value
func (o *AddConjurPasswordStorageSchemeRequest) SetConjurExternalServer(v string) {
	o.ConjurExternalServer = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddConjurPasswordStorageSchemeRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddConjurPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddConjurPasswordStorageSchemeRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddConjurPasswordStorageSchemeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddConjurPasswordStorageSchemeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddConjurPasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddConjurPasswordStorageSchemeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddConjurPasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemeName"] = o.SchemeName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["conjurExternalServer"] = o.ConjurExternalServer
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddConjurPasswordStorageSchemeRequest struct {
	value *AddConjurPasswordStorageSchemeRequest
	isSet bool
}

func (v NullableAddConjurPasswordStorageSchemeRequest) Get() *AddConjurPasswordStorageSchemeRequest {
	return v.value
}

func (v *NullableAddConjurPasswordStorageSchemeRequest) Set(val *AddConjurPasswordStorageSchemeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddConjurPasswordStorageSchemeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddConjurPasswordStorageSchemeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddConjurPasswordStorageSchemeRequest(val *AddConjurPasswordStorageSchemeRequest) *NullableAddConjurPasswordStorageSchemeRequest {
	return &NullableAddConjurPasswordStorageSchemeRequest{value: val, isSet: true}
}

func (v NullableAddConjurPasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddConjurPasswordStorageSchemeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
