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

// AddBcryptPasswordStorageSchemeRequest struct for AddBcryptPasswordStorageSchemeRequest
type AddBcryptPasswordStorageSchemeRequest struct {
	// Name of the new Password Storage Scheme
	SchemeName string `json:"schemeName"`
	Schemas []EnumbcryptPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// Specifies the cost factor to use when encoding passwords with Bcrypt. A higher cost factor requires more processing to generate a password, which makes attacks against the password more expensive.
	BcryptCostFactor *int32 `json:"bcryptCostFactor,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddBcryptPasswordStorageSchemeRequest instantiates a new AddBcryptPasswordStorageSchemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddBcryptPasswordStorageSchemeRequest(schemeName string, schemas []EnumbcryptPasswordStorageSchemeSchemaUrn, enabled bool) *AddBcryptPasswordStorageSchemeRequest {
	this := AddBcryptPasswordStorageSchemeRequest{}
	this.SchemeName = schemeName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddBcryptPasswordStorageSchemeRequestWithDefaults instantiates a new AddBcryptPasswordStorageSchemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddBcryptPasswordStorageSchemeRequestWithDefaults() *AddBcryptPasswordStorageSchemeRequest {
	this := AddBcryptPasswordStorageSchemeRequest{}
	return &this
}

// GetSchemeName returns the SchemeName field value
func (o *AddBcryptPasswordStorageSchemeRequest) GetSchemeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemeName
}

// GetSchemeNameOk returns a tuple with the SchemeName field value
// and a boolean to check if the value has been set.
func (o *AddBcryptPasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemeName, true
}

// SetSchemeName sets field value
func (o *AddBcryptPasswordStorageSchemeRequest) SetSchemeName(v string) {
	o.SchemeName = v
}

// GetSchemas returns the Schemas field value
func (o *AddBcryptPasswordStorageSchemeRequest) GetSchemas() []EnumbcryptPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumbcryptPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddBcryptPasswordStorageSchemeRequest) GetSchemasOk() ([]EnumbcryptPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddBcryptPasswordStorageSchemeRequest) SetSchemas(v []EnumbcryptPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetBcryptCostFactor returns the BcryptCostFactor field value if set, zero value otherwise.
func (o *AddBcryptPasswordStorageSchemeRequest) GetBcryptCostFactor() int32 {
	if o == nil || isNil(o.BcryptCostFactor) {
		var ret int32
		return ret
	}
	return *o.BcryptCostFactor
}

// GetBcryptCostFactorOk returns a tuple with the BcryptCostFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBcryptPasswordStorageSchemeRequest) GetBcryptCostFactorOk() (*int32, bool) {
	if o == nil || isNil(o.BcryptCostFactor) {
    return nil, false
	}
	return o.BcryptCostFactor, true
}

// HasBcryptCostFactor returns a boolean if a field has been set.
func (o *AddBcryptPasswordStorageSchemeRequest) HasBcryptCostFactor() bool {
	if o != nil && !isNil(o.BcryptCostFactor) {
		return true
	}

	return false
}

// SetBcryptCostFactor gets a reference to the given int32 and assigns it to the BcryptCostFactor field.
func (o *AddBcryptPasswordStorageSchemeRequest) SetBcryptCostFactor(v int32) {
	o.BcryptCostFactor = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddBcryptPasswordStorageSchemeRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBcryptPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddBcryptPasswordStorageSchemeRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddBcryptPasswordStorageSchemeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddBcryptPasswordStorageSchemeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddBcryptPasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddBcryptPasswordStorageSchemeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddBcryptPasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemeName"] = o.SchemeName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.BcryptCostFactor) {
		toSerialize["bcryptCostFactor"] = o.BcryptCostFactor
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddBcryptPasswordStorageSchemeRequest struct {
	value *AddBcryptPasswordStorageSchemeRequest
	isSet bool
}

func (v NullableAddBcryptPasswordStorageSchemeRequest) Get() *AddBcryptPasswordStorageSchemeRequest {
	return v.value
}

func (v *NullableAddBcryptPasswordStorageSchemeRequest) Set(val *AddBcryptPasswordStorageSchemeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddBcryptPasswordStorageSchemeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddBcryptPasswordStorageSchemeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddBcryptPasswordStorageSchemeRequest(val *AddBcryptPasswordStorageSchemeRequest) *NullableAddBcryptPasswordStorageSchemeRequest {
	return &NullableAddBcryptPasswordStorageSchemeRequest{value: val, isSet: true}
}

func (v NullableAddBcryptPasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddBcryptPasswordStorageSchemeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


