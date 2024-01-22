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

// checks if the AddBlindTrustManagerProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddBlindTrustManagerProviderRequest{}

// AddBlindTrustManagerProviderRequest struct for AddBlindTrustManagerProviderRequest
type AddBlindTrustManagerProviderRequest struct {
	Schemas []EnumblindTrustManagerProviderSchemaUrn `json:"schemas"`
	// Indicate whether the Trust Manager Provider is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether certificates issued by an authority included in the JVM's set of default issuers should be automatically trusted, even if they would not otherwise be trusted by this provider.
	IncludeJVMDefaultIssuers *bool `json:"includeJVMDefaultIssuers,omitempty"`
	// Name of the new Trust Manager Provider
	ProviderName string `json:"providerName"`
}

// NewAddBlindTrustManagerProviderRequest instantiates a new AddBlindTrustManagerProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddBlindTrustManagerProviderRequest(schemas []EnumblindTrustManagerProviderSchemaUrn, enabled bool, providerName string) *AddBlindTrustManagerProviderRequest {
	this := AddBlindTrustManagerProviderRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.ProviderName = providerName
	return &this
}

// NewAddBlindTrustManagerProviderRequestWithDefaults instantiates a new AddBlindTrustManagerProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddBlindTrustManagerProviderRequestWithDefaults() *AddBlindTrustManagerProviderRequest {
	this := AddBlindTrustManagerProviderRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddBlindTrustManagerProviderRequest) GetSchemas() []EnumblindTrustManagerProviderSchemaUrn {
	if o == nil {
		var ret []EnumblindTrustManagerProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddBlindTrustManagerProviderRequest) GetSchemasOk() ([]EnumblindTrustManagerProviderSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddBlindTrustManagerProviderRequest) SetSchemas(v []EnumblindTrustManagerProviderSchemaUrn) {
	o.Schemas = v
}

// GetEnabled returns the Enabled field value
func (o *AddBlindTrustManagerProviderRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddBlindTrustManagerProviderRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddBlindTrustManagerProviderRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIncludeJVMDefaultIssuers returns the IncludeJVMDefaultIssuers field value if set, zero value otherwise.
func (o *AddBlindTrustManagerProviderRequest) GetIncludeJVMDefaultIssuers() bool {
	if o == nil || IsNil(o.IncludeJVMDefaultIssuers) {
		var ret bool
		return ret
	}
	return *o.IncludeJVMDefaultIssuers
}

// GetIncludeJVMDefaultIssuersOk returns a tuple with the IncludeJVMDefaultIssuers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddBlindTrustManagerProviderRequest) GetIncludeJVMDefaultIssuersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeJVMDefaultIssuers) {
		return nil, false
	}
	return o.IncludeJVMDefaultIssuers, true
}

// HasIncludeJVMDefaultIssuers returns a boolean if a field has been set.
func (o *AddBlindTrustManagerProviderRequest) HasIncludeJVMDefaultIssuers() bool {
	if o != nil && !IsNil(o.IncludeJVMDefaultIssuers) {
		return true
	}

	return false
}

// SetIncludeJVMDefaultIssuers gets a reference to the given bool and assigns it to the IncludeJVMDefaultIssuers field.
func (o *AddBlindTrustManagerProviderRequest) SetIncludeJVMDefaultIssuers(v bool) {
	o.IncludeJVMDefaultIssuers = &v
}

// GetProviderName returns the ProviderName field value
func (o *AddBlindTrustManagerProviderRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *AddBlindTrustManagerProviderRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *AddBlindTrustManagerProviderRequest) SetProviderName(v string) {
	o.ProviderName = v
}

func (o AddBlindTrustManagerProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddBlindTrustManagerProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.IncludeJVMDefaultIssuers) {
		toSerialize["includeJVMDefaultIssuers"] = o.IncludeJVMDefaultIssuers
	}
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableAddBlindTrustManagerProviderRequest struct {
	value *AddBlindTrustManagerProviderRequest
	isSet bool
}

func (v NullableAddBlindTrustManagerProviderRequest) Get() *AddBlindTrustManagerProviderRequest {
	return v.value
}

func (v *NullableAddBlindTrustManagerProviderRequest) Set(val *AddBlindTrustManagerProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddBlindTrustManagerProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddBlindTrustManagerProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddBlindTrustManagerProviderRequest(val *AddBlindTrustManagerProviderRequest) *NullableAddBlindTrustManagerProviderRequest {
	return &NullableAddBlindTrustManagerProviderRequest{value: val, isSet: true}
}

func (v NullableAddBlindTrustManagerProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddBlindTrustManagerProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
