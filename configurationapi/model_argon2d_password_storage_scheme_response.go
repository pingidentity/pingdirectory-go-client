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

// checks if the Argon2dPasswordStorageSchemeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Argon2dPasswordStorageSchemeResponse{}

// Argon2dPasswordStorageSchemeResponse struct for Argon2dPasswordStorageSchemeResponse
type Argon2dPasswordStorageSchemeResponse struct {
	Schemas []Enumargon2dPasswordStorageSchemeSchemaUrn `json:"schemas"`
	// The number of rounds of cryptographic processing required in the course of encoding each password.
	IterationCount int64 `json:"iterationCount"`
	// The number of concurrent threads that will be used in the course of encoding each password.
	ParallelismFactor int64 `json:"parallelismFactor"`
	// The number of kilobytes of memory that must be used in the course of encoding each password.
	MemoryUsageKb int64 `json:"memoryUsageKb"`
	// The number of bytes to use for the generated salt.
	SaltLengthBytes int64 `json:"saltLengthBytes"`
	// The number of bytes to use for the derived key. The value must be greater than or equal to 8 and less than or equal to 512.
	DerivedKeyLengthBytes int64 `json:"derivedKeyLengthBytes"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Password Storage Scheme
	Id string `json:"id"`
}

// NewArgon2dPasswordStorageSchemeResponse instantiates a new Argon2dPasswordStorageSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArgon2dPasswordStorageSchemeResponse(schemas []Enumargon2dPasswordStorageSchemeSchemaUrn, iterationCount int64, parallelismFactor int64, memoryUsageKb int64, saltLengthBytes int64, derivedKeyLengthBytes int64, enabled bool, id string) *Argon2dPasswordStorageSchemeResponse {
	this := Argon2dPasswordStorageSchemeResponse{}
	this.Schemas = schemas
	this.IterationCount = iterationCount
	this.ParallelismFactor = parallelismFactor
	this.MemoryUsageKb = memoryUsageKb
	this.SaltLengthBytes = saltLengthBytes
	this.DerivedKeyLengthBytes = derivedKeyLengthBytes
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewArgon2dPasswordStorageSchemeResponseWithDefaults instantiates a new Argon2dPasswordStorageSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArgon2dPasswordStorageSchemeResponseWithDefaults() *Argon2dPasswordStorageSchemeResponse {
	this := Argon2dPasswordStorageSchemeResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *Argon2dPasswordStorageSchemeResponse) GetSchemas() []Enumargon2dPasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []Enumargon2dPasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *Argon2dPasswordStorageSchemeResponse) GetSchemasOk() ([]Enumargon2dPasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *Argon2dPasswordStorageSchemeResponse) SetSchemas(v []Enumargon2dPasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetIterationCount returns the IterationCount field value
func (o *Argon2dPasswordStorageSchemeResponse) GetIterationCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IterationCount
}

// GetIterationCountOk returns a tuple with the IterationCount field value
// and a boolean to check if the value has been set.
func (o *Argon2dPasswordStorageSchemeResponse) GetIterationCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IterationCount, true
}

// SetIterationCount sets field value
func (o *Argon2dPasswordStorageSchemeResponse) SetIterationCount(v int64) {
	o.IterationCount = v
}

// GetParallelismFactor returns the ParallelismFactor field value
func (o *Argon2dPasswordStorageSchemeResponse) GetParallelismFactor() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ParallelismFactor
}

// GetParallelismFactorOk returns a tuple with the ParallelismFactor field value
// and a boolean to check if the value has been set.
func (o *Argon2dPasswordStorageSchemeResponse) GetParallelismFactorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParallelismFactor, true
}

// SetParallelismFactor sets field value
func (o *Argon2dPasswordStorageSchemeResponse) SetParallelismFactor(v int64) {
	o.ParallelismFactor = v
}

// GetMemoryUsageKb returns the MemoryUsageKb field value
func (o *Argon2dPasswordStorageSchemeResponse) GetMemoryUsageKb() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MemoryUsageKb
}

// GetMemoryUsageKbOk returns a tuple with the MemoryUsageKb field value
// and a boolean to check if the value has been set.
func (o *Argon2dPasswordStorageSchemeResponse) GetMemoryUsageKbOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryUsageKb, true
}

// SetMemoryUsageKb sets field value
func (o *Argon2dPasswordStorageSchemeResponse) SetMemoryUsageKb(v int64) {
	o.MemoryUsageKb = v
}

// GetSaltLengthBytes returns the SaltLengthBytes field value
func (o *Argon2dPasswordStorageSchemeResponse) GetSaltLengthBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SaltLengthBytes
}

// GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field value
// and a boolean to check if the value has been set.
func (o *Argon2dPasswordStorageSchemeResponse) GetSaltLengthBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SaltLengthBytes, true
}

// SetSaltLengthBytes sets field value
func (o *Argon2dPasswordStorageSchemeResponse) SetSaltLengthBytes(v int64) {
	o.SaltLengthBytes = v
}

// GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field value
func (o *Argon2dPasswordStorageSchemeResponse) GetDerivedKeyLengthBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DerivedKeyLengthBytes
}

// GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field value
// and a boolean to check if the value has been set.
func (o *Argon2dPasswordStorageSchemeResponse) GetDerivedKeyLengthBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DerivedKeyLengthBytes, true
}

// SetDerivedKeyLengthBytes sets field value
func (o *Argon2dPasswordStorageSchemeResponse) SetDerivedKeyLengthBytes(v int64) {
	o.DerivedKeyLengthBytes = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Argon2dPasswordStorageSchemeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Argon2dPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Argon2dPasswordStorageSchemeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Argon2dPasswordStorageSchemeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *Argon2dPasswordStorageSchemeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Argon2dPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Argon2dPasswordStorageSchemeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Argon2dPasswordStorageSchemeResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Argon2dPasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Argon2dPasswordStorageSchemeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *Argon2dPasswordStorageSchemeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *Argon2dPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Argon2dPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *Argon2dPasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *Argon2dPasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *Argon2dPasswordStorageSchemeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Argon2dPasswordStorageSchemeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Argon2dPasswordStorageSchemeResponse) SetId(v string) {
	o.Id = v
}

func (o Argon2dPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Argon2dPasswordStorageSchemeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableArgon2dPasswordStorageSchemeResponse struct {
	value *Argon2dPasswordStorageSchemeResponse
	isSet bool
}

func (v NullableArgon2dPasswordStorageSchemeResponse) Get() *Argon2dPasswordStorageSchemeResponse {
	return v.value
}

func (v *NullableArgon2dPasswordStorageSchemeResponse) Set(val *Argon2dPasswordStorageSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableArgon2dPasswordStorageSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableArgon2dPasswordStorageSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArgon2dPasswordStorageSchemeResponse(val *Argon2dPasswordStorageSchemeResponse) *NullableArgon2dPasswordStorageSchemeResponse {
	return &NullableArgon2dPasswordStorageSchemeResponse{value: val, isSet: true}
}

func (v NullableArgon2dPasswordStorageSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArgon2dPasswordStorageSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
