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

// checks if the AddScryptPasswordStorageSchemeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddScryptPasswordStorageSchemeRequest{}

// AddScryptPasswordStorageSchemeRequest struct for AddScryptPasswordStorageSchemeRequest
type AddScryptPasswordStorageSchemeRequest struct {
	Schemas []EnumscryptPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// Specifies the exponent that should be used for the CPU/memory cost factor. The cost factor must be a power of two, so the value of this property represents the power to which two is raised. The CPU/memory cost factor specifies the number of iterations required for encoding the password, and also affects the amount of memory required during processing. A higher cost factor requires more processing and more memory to generate a password, which makes attacks against the password more expensive.
	ScryptCpuMemoryCostFactorExponent *int64 `json:"scryptCpuMemoryCostFactorExponent,omitempty"`
	// Specifies the block size for the digest that will be used in the course of encoding passwords. Increasing the block size while keeping the CPU/memory cost factor constant will increase the amount of memory required to encode a password, but it also increases the ratio of sequential memory access to random memory access (and sequential memory access is generally faster than random memory access).
	ScryptBlockSize *int64 `json:"scryptBlockSize,omitempty"`
	// Specifies the number of times that scrypt has to perform the entire encoding process to produce the final result.
	ScryptParallelizationParameter *int64 `json:"scryptParallelizationParameter,omitempty"`
	// Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords.
	MaxPasswordLength *int64 `json:"maxPasswordLength,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
	// Name of the new Password Storage Scheme
	SchemeName string `json:"schemeName"`
}

// NewAddScryptPasswordStorageSchemeRequest instantiates a new AddScryptPasswordStorageSchemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddScryptPasswordStorageSchemeRequest(schemas []EnumscryptPasswordStorageSchemeSchemaUrn, enabled bool, schemeName string) *AddScryptPasswordStorageSchemeRequest {
	this := AddScryptPasswordStorageSchemeRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.SchemeName = schemeName
	return &this
}

// NewAddScryptPasswordStorageSchemeRequestWithDefaults instantiates a new AddScryptPasswordStorageSchemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddScryptPasswordStorageSchemeRequestWithDefaults() *AddScryptPasswordStorageSchemeRequest {
	this := AddScryptPasswordStorageSchemeRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddScryptPasswordStorageSchemeRequest) GetSchemas() []EnumscryptPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumscryptPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddScryptPasswordStorageSchemeRequest) GetSchemasOk() ([]EnumscryptPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddScryptPasswordStorageSchemeRequest) SetSchemas(v []EnumscryptPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetScryptCpuMemoryCostFactorExponent returns the ScryptCpuMemoryCostFactorExponent field value if set, zero value otherwise.
func (o *AddScryptPasswordStorageSchemeRequest) GetScryptCpuMemoryCostFactorExponent() int64 {
	if o == nil || IsNil(o.ScryptCpuMemoryCostFactorExponent) {
		var ret int64
		return ret
	}
	return *o.ScryptCpuMemoryCostFactorExponent
}

// GetScryptCpuMemoryCostFactorExponentOk returns a tuple with the ScryptCpuMemoryCostFactorExponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScryptPasswordStorageSchemeRequest) GetScryptCpuMemoryCostFactorExponentOk() (*int64, bool) {
	if o == nil || IsNil(o.ScryptCpuMemoryCostFactorExponent) {
		return nil, false
	}
	return o.ScryptCpuMemoryCostFactorExponent, true
}

// HasScryptCpuMemoryCostFactorExponent returns a boolean if a field has been set.
func (o *AddScryptPasswordStorageSchemeRequest) HasScryptCpuMemoryCostFactorExponent() bool {
	if o != nil && !IsNil(o.ScryptCpuMemoryCostFactorExponent) {
		return true
	}

	return false
}

// SetScryptCpuMemoryCostFactorExponent gets a reference to the given int64 and assigns it to the ScryptCpuMemoryCostFactorExponent field.
func (o *AddScryptPasswordStorageSchemeRequest) SetScryptCpuMemoryCostFactorExponent(v int64) {
	o.ScryptCpuMemoryCostFactorExponent = &v
}

// GetScryptBlockSize returns the ScryptBlockSize field value if set, zero value otherwise.
func (o *AddScryptPasswordStorageSchemeRequest) GetScryptBlockSize() int64 {
	if o == nil || IsNil(o.ScryptBlockSize) {
		var ret int64
		return ret
	}
	return *o.ScryptBlockSize
}

// GetScryptBlockSizeOk returns a tuple with the ScryptBlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScryptPasswordStorageSchemeRequest) GetScryptBlockSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.ScryptBlockSize) {
		return nil, false
	}
	return o.ScryptBlockSize, true
}

// HasScryptBlockSize returns a boolean if a field has been set.
func (o *AddScryptPasswordStorageSchemeRequest) HasScryptBlockSize() bool {
	if o != nil && !IsNil(o.ScryptBlockSize) {
		return true
	}

	return false
}

// SetScryptBlockSize gets a reference to the given int64 and assigns it to the ScryptBlockSize field.
func (o *AddScryptPasswordStorageSchemeRequest) SetScryptBlockSize(v int64) {
	o.ScryptBlockSize = &v
}

// GetScryptParallelizationParameter returns the ScryptParallelizationParameter field value if set, zero value otherwise.
func (o *AddScryptPasswordStorageSchemeRequest) GetScryptParallelizationParameter() int64 {
	if o == nil || IsNil(o.ScryptParallelizationParameter) {
		var ret int64
		return ret
	}
	return *o.ScryptParallelizationParameter
}

// GetScryptParallelizationParameterOk returns a tuple with the ScryptParallelizationParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScryptPasswordStorageSchemeRequest) GetScryptParallelizationParameterOk() (*int64, bool) {
	if o == nil || IsNil(o.ScryptParallelizationParameter) {
		return nil, false
	}
	return o.ScryptParallelizationParameter, true
}

// HasScryptParallelizationParameter returns a boolean if a field has been set.
func (o *AddScryptPasswordStorageSchemeRequest) HasScryptParallelizationParameter() bool {
	if o != nil && !IsNil(o.ScryptParallelizationParameter) {
		return true
	}

	return false
}

// SetScryptParallelizationParameter gets a reference to the given int64 and assigns it to the ScryptParallelizationParameter field.
func (o *AddScryptPasswordStorageSchemeRequest) SetScryptParallelizationParameter(v int64) {
	o.ScryptParallelizationParameter = &v
}

// GetMaxPasswordLength returns the MaxPasswordLength field value if set, zero value otherwise.
func (o *AddScryptPasswordStorageSchemeRequest) GetMaxPasswordLength() int64 {
	if o == nil || IsNil(o.MaxPasswordLength) {
		var ret int64
		return ret
	}
	return *o.MaxPasswordLength
}

// GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScryptPasswordStorageSchemeRequest) GetMaxPasswordLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxPasswordLength) {
		return nil, false
	}
	return o.MaxPasswordLength, true
}

// HasMaxPasswordLength returns a boolean if a field has been set.
func (o *AddScryptPasswordStorageSchemeRequest) HasMaxPasswordLength() bool {
	if o != nil && !IsNil(o.MaxPasswordLength) {
		return true
	}

	return false
}

// SetMaxPasswordLength gets a reference to the given int64 and assigns it to the MaxPasswordLength field.
func (o *AddScryptPasswordStorageSchemeRequest) SetMaxPasswordLength(v int64) {
	o.MaxPasswordLength = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddScryptPasswordStorageSchemeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScryptPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddScryptPasswordStorageSchemeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddScryptPasswordStorageSchemeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddScryptPasswordStorageSchemeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddScryptPasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddScryptPasswordStorageSchemeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSchemeName returns the SchemeName field value
func (o *AddScryptPasswordStorageSchemeRequest) GetSchemeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemeName
}

// GetSchemeNameOk returns a tuple with the SchemeName field value
// and a boolean to check if the value has been set.
func (o *AddScryptPasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemeName, true
}

// SetSchemeName sets field value
func (o *AddScryptPasswordStorageSchemeRequest) SetSchemeName(v string) {
	o.SchemeName = v
}

func (o AddScryptPasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddScryptPasswordStorageSchemeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.ScryptCpuMemoryCostFactorExponent) {
		toSerialize["scryptCpuMemoryCostFactorExponent"] = o.ScryptCpuMemoryCostFactorExponent
	}
	if !IsNil(o.ScryptBlockSize) {
		toSerialize["scryptBlockSize"] = o.ScryptBlockSize
	}
	if !IsNil(o.ScryptParallelizationParameter) {
		toSerialize["scryptParallelizationParameter"] = o.ScryptParallelizationParameter
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

type NullableAddScryptPasswordStorageSchemeRequest struct {
	value *AddScryptPasswordStorageSchemeRequest
	isSet bool
}

func (v NullableAddScryptPasswordStorageSchemeRequest) Get() *AddScryptPasswordStorageSchemeRequest {
	return v.value
}

func (v *NullableAddScryptPasswordStorageSchemeRequest) Set(val *AddScryptPasswordStorageSchemeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddScryptPasswordStorageSchemeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddScryptPasswordStorageSchemeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddScryptPasswordStorageSchemeRequest(val *AddScryptPasswordStorageSchemeRequest) *NullableAddScryptPasswordStorageSchemeRequest {
	return &NullableAddScryptPasswordStorageSchemeRequest{value: val, isSet: true}
}

func (v NullableAddScryptPasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddScryptPasswordStorageSchemeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
