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

// BcryptPasswordStorageSchemeShared struct for BcryptPasswordStorageSchemeShared
type BcryptPasswordStorageSchemeShared struct {
	Schemas []EnumbcryptPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// Specifies the cost factor to use when encoding passwords with Bcrypt. A higher cost factor requires more processing to generate a password, which makes attacks against the password more expensive.
	BcryptCostFactor *int32 `json:"bcryptCostFactor,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewBcryptPasswordStorageSchemeShared instantiates a new BcryptPasswordStorageSchemeShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBcryptPasswordStorageSchemeShared(schemas []EnumbcryptPasswordStorageSchemeSchemaUrn, enabled bool) *BcryptPasswordStorageSchemeShared {
	this := BcryptPasswordStorageSchemeShared{}
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewBcryptPasswordStorageSchemeSharedWithDefaults instantiates a new BcryptPasswordStorageSchemeShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBcryptPasswordStorageSchemeSharedWithDefaults() *BcryptPasswordStorageSchemeShared {
	this := BcryptPasswordStorageSchemeShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *BcryptPasswordStorageSchemeShared) GetSchemas() []EnumbcryptPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumbcryptPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *BcryptPasswordStorageSchemeShared) GetSchemasOk() ([]EnumbcryptPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *BcryptPasswordStorageSchemeShared) SetSchemas(v []EnumbcryptPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetBcryptCostFactor returns the BcryptCostFactor field value if set, zero value otherwise.
func (o *BcryptPasswordStorageSchemeShared) GetBcryptCostFactor() int32 {
	if o == nil || isNil(o.BcryptCostFactor) {
		var ret int32
		return ret
	}
	return *o.BcryptCostFactor
}

// GetBcryptCostFactorOk returns a tuple with the BcryptCostFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BcryptPasswordStorageSchemeShared) GetBcryptCostFactorOk() (*int32, bool) {
	if o == nil || isNil(o.BcryptCostFactor) {
    return nil, false
	}
	return o.BcryptCostFactor, true
}

// HasBcryptCostFactor returns a boolean if a field has been set.
func (o *BcryptPasswordStorageSchemeShared) HasBcryptCostFactor() bool {
	if o != nil && !isNil(o.BcryptCostFactor) {
		return true
	}

	return false
}

// SetBcryptCostFactor gets a reference to the given int32 and assigns it to the BcryptCostFactor field.
func (o *BcryptPasswordStorageSchemeShared) SetBcryptCostFactor(v int32) {
	o.BcryptCostFactor = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BcryptPasswordStorageSchemeShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BcryptPasswordStorageSchemeShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BcryptPasswordStorageSchemeShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BcryptPasswordStorageSchemeShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *BcryptPasswordStorageSchemeShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *BcryptPasswordStorageSchemeShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *BcryptPasswordStorageSchemeShared) SetEnabled(v bool) {
	o.Enabled = v
}

func (o BcryptPasswordStorageSchemeShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableBcryptPasswordStorageSchemeShared struct {
	value *BcryptPasswordStorageSchemeShared
	isSet bool
}

func (v NullableBcryptPasswordStorageSchemeShared) Get() *BcryptPasswordStorageSchemeShared {
	return v.value
}

func (v *NullableBcryptPasswordStorageSchemeShared) Set(val *BcryptPasswordStorageSchemeShared) {
	v.value = val
	v.isSet = true
}

func (v NullableBcryptPasswordStorageSchemeShared) IsSet() bool {
	return v.isSet
}

func (v *NullableBcryptPasswordStorageSchemeShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBcryptPasswordStorageSchemeShared(val *BcryptPasswordStorageSchemeShared) *NullableBcryptPasswordStorageSchemeShared {
	return &NullableBcryptPasswordStorageSchemeShared{value: val, isSet: true}
}

func (v NullableBcryptPasswordStorageSchemeShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBcryptPasswordStorageSchemeShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


