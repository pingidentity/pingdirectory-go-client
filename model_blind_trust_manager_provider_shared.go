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

// BlindTrustManagerProviderShared struct for BlindTrustManagerProviderShared
type BlindTrustManagerProviderShared struct {
	Schemas []EnumblindTrustManagerProviderSchemaUrn `json:"schemas"`
	// Indicate whether the Trust Manager Provider is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether certificates issued by an authority included in the JVM's set of default issuers should be automatically trusted, even if they would not otherwise be trusted by this provider.
	IncludeJVMDefaultIssuers *bool `json:"includeJVMDefaultIssuers,omitempty"`
}

// NewBlindTrustManagerProviderShared instantiates a new BlindTrustManagerProviderShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlindTrustManagerProviderShared(schemas []EnumblindTrustManagerProviderSchemaUrn, enabled bool) *BlindTrustManagerProviderShared {
	this := BlindTrustManagerProviderShared{}
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewBlindTrustManagerProviderSharedWithDefaults instantiates a new BlindTrustManagerProviderShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlindTrustManagerProviderSharedWithDefaults() *BlindTrustManagerProviderShared {
	this := BlindTrustManagerProviderShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *BlindTrustManagerProviderShared) GetSchemas() []EnumblindTrustManagerProviderSchemaUrn {
	if o == nil {
		var ret []EnumblindTrustManagerProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *BlindTrustManagerProviderShared) GetSchemasOk() ([]EnumblindTrustManagerProviderSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *BlindTrustManagerProviderShared) SetSchemas(v []EnumblindTrustManagerProviderSchemaUrn) {
	o.Schemas = v
}

// GetEnabled returns the Enabled field value
func (o *BlindTrustManagerProviderShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *BlindTrustManagerProviderShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *BlindTrustManagerProviderShared) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIncludeJVMDefaultIssuers returns the IncludeJVMDefaultIssuers field value if set, zero value otherwise.
func (o *BlindTrustManagerProviderShared) GetIncludeJVMDefaultIssuers() bool {
	if o == nil || isNil(o.IncludeJVMDefaultIssuers) {
		var ret bool
		return ret
	}
	return *o.IncludeJVMDefaultIssuers
}

// GetIncludeJVMDefaultIssuersOk returns a tuple with the IncludeJVMDefaultIssuers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlindTrustManagerProviderShared) GetIncludeJVMDefaultIssuersOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeJVMDefaultIssuers) {
    return nil, false
	}
	return o.IncludeJVMDefaultIssuers, true
}

// HasIncludeJVMDefaultIssuers returns a boolean if a field has been set.
func (o *BlindTrustManagerProviderShared) HasIncludeJVMDefaultIssuers() bool {
	if o != nil && !isNil(o.IncludeJVMDefaultIssuers) {
		return true
	}

	return false
}

// SetIncludeJVMDefaultIssuers gets a reference to the given bool and assigns it to the IncludeJVMDefaultIssuers field.
func (o *BlindTrustManagerProviderShared) SetIncludeJVMDefaultIssuers(v bool) {
	o.IncludeJVMDefaultIssuers = &v
}

func (o BlindTrustManagerProviderShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.IncludeJVMDefaultIssuers) {
		toSerialize["includeJVMDefaultIssuers"] = o.IncludeJVMDefaultIssuers
	}
	return json.Marshal(toSerialize)
}

type NullableBlindTrustManagerProviderShared struct {
	value *BlindTrustManagerProviderShared
	isSet bool
}

func (v NullableBlindTrustManagerProviderShared) Get() *BlindTrustManagerProviderShared {
	return v.value
}

func (v *NullableBlindTrustManagerProviderShared) Set(val *BlindTrustManagerProviderShared) {
	v.value = val
	v.isSet = true
}

func (v NullableBlindTrustManagerProviderShared) IsSet() bool {
	return v.isSet
}

func (v *NullableBlindTrustManagerProviderShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlindTrustManagerProviderShared(val *BlindTrustManagerProviderShared) *NullableBlindTrustManagerProviderShared {
	return &NullableBlindTrustManagerProviderShared{value: val, isSet: true}
}

func (v NullableBlindTrustManagerProviderShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlindTrustManagerProviderShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


