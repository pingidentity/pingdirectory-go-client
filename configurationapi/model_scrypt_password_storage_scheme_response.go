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

// checks if the ScryptPasswordStorageSchemeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScryptPasswordStorageSchemeResponse{}

// ScryptPasswordStorageSchemeResponse struct for ScryptPasswordStorageSchemeResponse
type ScryptPasswordStorageSchemeResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Password Storage Scheme
	Id      string                                     `json:"id"`
	Schemas []EnumscryptPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// Specifies the exponent that should be used for the CPU/memory cost factor. The cost factor must be a power of two, so the value of this property represents the power to which two is raised. The CPU/memory cost factor specifies the number of iterations required for encoding the password, and also affects the amount of memory required during processing. A higher cost factor requires more processing and more memory to generate a password, which makes attacks against the password more expensive.
	ScryptCpuMemoryCostFactorExponent *int32 `json:"scryptCpuMemoryCostFactorExponent,omitempty"`
	// Specifies the block size for the digest that will be used in the course of encoding passwords. Increasing the block size while keeping the CPU/memory cost factor constant will increase the amount of memory required to encode a password, but it also increases the ratio of sequential memory access to random memory access (and sequential memory access is generally faster than random memory access).
	ScryptBlockSize *int32 `json:"scryptBlockSize,omitempty"`
	// Specifies the number of times that scrypt has to perform the entire encoding process to produce the final result.
	ScryptParallelizationParameter *int32 `json:"scryptParallelizationParameter,omitempty"`
	// Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords.
	MaxPasswordLength *int32 `json:"maxPasswordLength,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewScryptPasswordStorageSchemeResponse instantiates a new ScryptPasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScryptPasswordStorageSchemeResponse(id string, schemas []EnumscryptPasswordStorageSchemeSchemaUrn, enabled bool) *ScryptPasswordStorageSchemeResponse {
	this := ScryptPasswordStorageSchemeResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewScryptPasswordStorageSchemeResponseWithDefaults instantiates a new ScryptPasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScryptPasswordStorageSchemeResponseWithDefaults() *ScryptPasswordStorageSchemeResponse {
	this := ScryptPasswordStorageSchemeResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScryptPasswordStorageSchemeResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScryptPasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScryptPasswordStorageSchemeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ScryptPasswordStorageSchemeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ScryptPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScryptPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ScryptPasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ScryptPasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ScryptPasswordStorageSchemeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScryptPasswordStorageSchemeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScryptPasswordStorageSchemeResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ScryptPasswordStorageSchemeResponse) GetSchemas() []EnumscryptPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []EnumscryptPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ScryptPasswordStorageSchemeResponse) GetSchemasOk() ([]EnumscryptPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ScryptPasswordStorageSchemeResponse) SetSchemas(v []EnumscryptPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetScryptCpuMemoryCostFactorExponent returns the ScryptCpuMemoryCostFactorExponent field value if set, zero value otherwise.
func (o *ScryptPasswordStorageSchemeResponse) GetScryptCpuMemoryCostFactorExponent() int32 {
	if o == nil || IsNil(o.ScryptCpuMemoryCostFactorExponent) {
		var ret int32
		return ret
	}
	return *o.ScryptCpuMemoryCostFactorExponent
}

// GetScryptCpuMemoryCostFactorExponentOk returns a tuple with the ScryptCpuMemoryCostFactorExponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScryptPasswordStorageSchemeResponse) GetScryptCpuMemoryCostFactorExponentOk() (*int32, bool) {
	if o == nil || IsNil(o.ScryptCpuMemoryCostFactorExponent) {
		return nil, false
	}
	return o.ScryptCpuMemoryCostFactorExponent, true
}

// HasScryptCpuMemoryCostFactorExponent returns a boolean if a field has been set.
func (o *ScryptPasswordStorageSchemeResponse) HasScryptCpuMemoryCostFactorExponent() bool {
	if o != nil && !IsNil(o.ScryptCpuMemoryCostFactorExponent) {
		return true
	}

	return false
}

// SetScryptCpuMemoryCostFactorExponent gets a reference to the given int32 and assigns it to the ScryptCpuMemoryCostFactorExponent field.
func (o *ScryptPasswordStorageSchemeResponse) SetScryptCpuMemoryCostFactorExponent(v int32) {
	o.ScryptCpuMemoryCostFactorExponent = &v
}

// GetScryptBlockSize returns the ScryptBlockSize field value if set, zero value otherwise.
func (o *ScryptPasswordStorageSchemeResponse) GetScryptBlockSize() int32 {
	if o == nil || IsNil(o.ScryptBlockSize) {
		var ret int32
		return ret
	}
	return *o.ScryptBlockSize
}

// GetScryptBlockSizeOk returns a tuple with the ScryptBlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScryptPasswordStorageSchemeResponse) GetScryptBlockSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ScryptBlockSize) {
		return nil, false
	}
	return o.ScryptBlockSize, true
}

// HasScryptBlockSize returns a boolean if a field has been set.
func (o *ScryptPasswordStorageSchemeResponse) HasScryptBlockSize() bool {
	if o != nil && !IsNil(o.ScryptBlockSize) {
		return true
	}

	return false
}

// SetScryptBlockSize gets a reference to the given int32 and assigns it to the ScryptBlockSize field.
func (o *ScryptPasswordStorageSchemeResponse) SetScryptBlockSize(v int32) {
	o.ScryptBlockSize = &v
}

// GetScryptParallelizationParameter returns the ScryptParallelizationParameter field value if set, zero value otherwise.
func (o *ScryptPasswordStorageSchemeResponse) GetScryptParallelizationParameter() int32 {
	if o == nil || IsNil(o.ScryptParallelizationParameter) {
		var ret int32
		return ret
	}
	return *o.ScryptParallelizationParameter
}

// GetScryptParallelizationParameterOk returns a tuple with the ScryptParallelizationParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScryptPasswordStorageSchemeResponse) GetScryptParallelizationParameterOk() (*int32, bool) {
	if o == nil || IsNil(o.ScryptParallelizationParameter) {
		return nil, false
	}
	return o.ScryptParallelizationParameter, true
}

// HasScryptParallelizationParameter returns a boolean if a field has been set.
func (o *ScryptPasswordStorageSchemeResponse) HasScryptParallelizationParameter() bool {
	if o != nil && !IsNil(o.ScryptParallelizationParameter) {
		return true
	}

	return false
}

// SetScryptParallelizationParameter gets a reference to the given int32 and assigns it to the ScryptParallelizationParameter field.
func (o *ScryptPasswordStorageSchemeResponse) SetScryptParallelizationParameter(v int32) {
	o.ScryptParallelizationParameter = &v
}

// GetMaxPasswordLength returns the MaxPasswordLength field value if set, zero value otherwise.
func (o *ScryptPasswordStorageSchemeResponse) GetMaxPasswordLength() int32 {
	if o == nil || IsNil(o.MaxPasswordLength) {
		var ret int32
		return ret
	}
	return *o.MaxPasswordLength
}

// GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScryptPasswordStorageSchemeResponse) GetMaxPasswordLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPasswordLength) {
		return nil, false
	}
	return o.MaxPasswordLength, true
}

// HasMaxPasswordLength returns a boolean if a field has been set.
func (o *ScryptPasswordStorageSchemeResponse) HasMaxPasswordLength() bool {
	if o != nil && !IsNil(o.MaxPasswordLength) {
		return true
	}

	return false
}

// SetMaxPasswordLength gets a reference to the given int32 and assigns it to the MaxPasswordLength field.
func (o *ScryptPasswordStorageSchemeResponse) SetMaxPasswordLength(v int32) {
	o.MaxPasswordLength = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScryptPasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScryptPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScryptPasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScryptPasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ScryptPasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ScryptPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ScryptPasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ScryptPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScryptPasswordStorageSchemeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
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
	return toSerialize, nil
}

type NullableScryptPasswordStorageSchemeResponse struct {
	value *ScryptPasswordStorageSchemeResponse
	isSet bool
}

func (v NullableScryptPasswordStorageSchemeResponse) Get() *ScryptPasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableScryptPasswordStorageSchemeResponse) Set(val *ScryptPasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScryptPasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScryptPasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScryptPasswordStorageSchemeResponse(val *ScryptPasswordStorageSchemeResponse) *NullableScryptPasswordStorageSchemeResponse {
	return &NullableScryptPasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableScryptPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScryptPasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
