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

// checks if the AddArgon2PasswordStorageSchemeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddArgon2PasswordStorageSchemeRequest{}

// AddArgon2PasswordStorageSchemeRequest struct for AddArgon2PasswordStorageSchemeRequest
type AddArgon2PasswordStorageSchemeRequest struct {
	// Name of the new Password Storage Scheme
	SchemeName string                                     `json:"schemeName"`
	Schemas    []Enumargon2PasswordStorageSchemeSchemaUrn `json:"schemas"`
	// The number of rounds of cryptographic processing required in the course of encoding each password.
	IterationCount int32 `json:"iterationCount"`
	// The number of concurrent threads that will be used in the course of encoding each password.
	ParallelismFactor int32 `json:"parallelismFactor"`
	// The number of kilobytes of memory that must be used in the course of encoding each password.
	MemoryUsageKb int32 `json:"memoryUsageKb"`
	// The number of bytes to use for the generated salt.
	SaltLengthBytes int32 `json:"saltLengthBytes"`
	// The number of bytes to use for the derived key. The value must be greater than or equal to 8 and less than or equal to 512.
	DerivedKeyLengthBytes int32 `json:"derivedKeyLengthBytes"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddArgon2PasswordStorageSchemeRequest instantiates a new AddArgon2PasswordStorageSchemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddArgon2PasswordStorageSchemeRequest(schemeName string, schemas []Enumargon2PasswordStorageSchemeSchemaUrn, iterationCount int32, parallelismFactor int32, memoryUsageKb int32, saltLengthBytes int32, derivedKeyLengthBytes int32, enabled bool) *AddArgon2PasswordStorageSchemeRequest {
	this := AddArgon2PasswordStorageSchemeRequest{}
	this.SchemeName = schemeName
	this.Schemas = schemas
	this.IterationCount = iterationCount
	this.ParallelismFactor = parallelismFactor
	this.MemoryUsageKb = memoryUsageKb
	this.SaltLengthBytes = saltLengthBytes
	this.DerivedKeyLengthBytes = derivedKeyLengthBytes
	this.Enabled = enabled
	return &this
}

// NewAddArgon2PasswordStorageSchemeRequestWithDefaults instantiates a new AddArgon2PasswordStorageSchemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddArgon2PasswordStorageSchemeRequestWithDefaults() *AddArgon2PasswordStorageSchemeRequest {
	this := AddArgon2PasswordStorageSchemeRequest{}
	return &this
}

// GetSchemeName returns the SchemeName field value
func (o *AddArgon2PasswordStorageSchemeRequest) GetSchemeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemeName
}

// GetSchemeNameOk returns a tuple with the SchemeName field value
// and a boolean to check if the value has been set.
func (o *AddArgon2PasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemeName, true
}

// SetSchemeName sets field value
func (o *AddArgon2PasswordStorageSchemeRequest) SetSchemeName(v string) {
	o.SchemeName = v
}

// GetSchemas returns the Schemas field value
func (o *AddArgon2PasswordStorageSchemeRequest) GetSchemas() []Enumargon2PasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []Enumargon2PasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddArgon2PasswordStorageSchemeRequest) GetSchemasOk() ([]Enumargon2PasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddArgon2PasswordStorageSchemeRequest) SetSchemas(v []Enumargon2PasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetIterationCount returns the IterationCount field value
func (o *AddArgon2PasswordStorageSchemeRequest) GetIterationCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IterationCount
}

// GetIterationCountOk returns a tuple with the IterationCount field value
// and a boolean to check if the value has been set.
func (o *AddArgon2PasswordStorageSchemeRequest) GetIterationCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IterationCount, true
}

// SetIterationCount sets field value
func (o *AddArgon2PasswordStorageSchemeRequest) SetIterationCount(v int32) {
	o.IterationCount = v
}

// GetParallelismFactor returns the ParallelismFactor field value
func (o *AddArgon2PasswordStorageSchemeRequest) GetParallelismFactor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParallelismFactor
}

// GetParallelismFactorOk returns a tuple with the ParallelismFactor field value
// and a boolean to check if the value has been set.
func (o *AddArgon2PasswordStorageSchemeRequest) GetParallelismFactorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParallelismFactor, true
}

// SetParallelismFactor sets field value
func (o *AddArgon2PasswordStorageSchemeRequest) SetParallelismFactor(v int32) {
	o.ParallelismFactor = v
}

// GetMemoryUsageKb returns the MemoryUsageKb field value
func (o *AddArgon2PasswordStorageSchemeRequest) GetMemoryUsageKb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemoryUsageKb
}

// GetMemoryUsageKbOk returns a tuple with the MemoryUsageKb field value
// and a boolean to check if the value has been set.
func (o *AddArgon2PasswordStorageSchemeRequest) GetMemoryUsageKbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryUsageKb, true
}

// SetMemoryUsageKb sets field value
func (o *AddArgon2PasswordStorageSchemeRequest) SetMemoryUsageKb(v int32) {
	o.MemoryUsageKb = v
}

// GetSaltLengthBytes returns the SaltLengthBytes field value
func (o *AddArgon2PasswordStorageSchemeRequest) GetSaltLengthBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SaltLengthBytes
}

// GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field value
// and a boolean to check if the value has been set.
func (o *AddArgon2PasswordStorageSchemeRequest) GetSaltLengthBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SaltLengthBytes, true
}

// SetSaltLengthBytes sets field value
func (o *AddArgon2PasswordStorageSchemeRequest) SetSaltLengthBytes(v int32) {
	o.SaltLengthBytes = v
}

// GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field value
func (o *AddArgon2PasswordStorageSchemeRequest) GetDerivedKeyLengthBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DerivedKeyLengthBytes
}

// GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field value
// and a boolean to check if the value has been set.
func (o *AddArgon2PasswordStorageSchemeRequest) GetDerivedKeyLengthBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DerivedKeyLengthBytes, true
}

// SetDerivedKeyLengthBytes sets field value
func (o *AddArgon2PasswordStorageSchemeRequest) SetDerivedKeyLengthBytes(v int32) {
	o.DerivedKeyLengthBytes = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddArgon2PasswordStorageSchemeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddArgon2PasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddArgon2PasswordStorageSchemeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddArgon2PasswordStorageSchemeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddArgon2PasswordStorageSchemeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddArgon2PasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddArgon2PasswordStorageSchemeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddArgon2PasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddArgon2PasswordStorageSchemeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemeName"] = o.SchemeName
	toSerialize["schemas"] = o.Schemas
	toSerialize["iterationCount"] = o.IterationCount
	toSerialize["parallelismFactor"] = o.ParallelismFactor
	toSerialize["memoryUsageKb"] = o.MemoryUsageKb
	toSerialize["saltLengthBytes"] = o.SaltLengthBytes
	toSerialize["derivedKeyLengthBytes"] = o.DerivedKeyLengthBytes
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAddArgon2PasswordStorageSchemeRequest struct {
	value *AddArgon2PasswordStorageSchemeRequest
	isSet bool
}

func (v NullableAddArgon2PasswordStorageSchemeRequest) Get() *AddArgon2PasswordStorageSchemeRequest {
	return v.value
}

func (v *NullableAddArgon2PasswordStorageSchemeRequest) Set(val *AddArgon2PasswordStorageSchemeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddArgon2PasswordStorageSchemeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddArgon2PasswordStorageSchemeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddArgon2PasswordStorageSchemeRequest(val *AddArgon2PasswordStorageSchemeRequest) *NullableAddArgon2PasswordStorageSchemeRequest {
	return &NullableAddArgon2PasswordStorageSchemeRequest{value: val, isSet: true}
}

func (v NullableAddArgon2PasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddArgon2PasswordStorageSchemeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
