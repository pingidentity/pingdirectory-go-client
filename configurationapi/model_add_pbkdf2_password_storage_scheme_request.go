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

// checks if the AddPbkdf2PasswordStorageSchemeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPbkdf2PasswordStorageSchemeRequest{}

// AddPbkdf2PasswordStorageSchemeRequest struct for AddPbkdf2PasswordStorageSchemeRequest
type AddPbkdf2PasswordStorageSchemeRequest struct {
	// Name of the new Password Storage Scheme
	SchemeName      string                                        `json:"schemeName"`
	Schemas         []Enumpbkdf2PasswordStorageSchemeSchemaUrn    `json:"schemas"`
	DigestAlgorithm *EnumpasswordStorageSchemeDigestAlgorithmProp `json:"digestAlgorithm,omitempty"`
	// Specifies the number of iterations to use when encoding passwords. The value must be greater than or equal to 1000.
	IterationCount *int32 `json:"iterationCount,omitempty"`
	// Specifies the number of bytes to use for the generated salt. The value must be greater than or equal to 8.
	SaltLengthBytes *int32 `json:"saltLengthBytes,omitempty"`
	// Specifies the number of bytes to use for the derived key. The value must be greater than or equal to 8.
	DerivedKeyLengthBytes *int32 `json:"derivedKeyLengthBytes,omitempty"`
	// Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords.
	MaxPasswordLength *int32 `json:"maxPasswordLength,omitempty"`
	// A description for this Password Storage Scheme
	Description *string `json:"description,omitempty"`
	// Indicates whether the Password Storage Scheme is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddPbkdf2PasswordStorageSchemeRequest instantiates a new AddPbkdf2PasswordStorageSchemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPbkdf2PasswordStorageSchemeRequest(schemeName string, schemas []Enumpbkdf2PasswordStorageSchemeSchemaUrn, enabled bool) *AddPbkdf2PasswordStorageSchemeRequest {
	this := AddPbkdf2PasswordStorageSchemeRequest{}
	this.SchemeName = schemeName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddPbkdf2PasswordStorageSchemeRequestWithDefaults instantiates a new AddPbkdf2PasswordStorageSchemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPbkdf2PasswordStorageSchemeRequestWithDefaults() *AddPbkdf2PasswordStorageSchemeRequest {
	this := AddPbkdf2PasswordStorageSchemeRequest{}
	return &this
}

// GetSchemeName returns the SchemeName field value
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSchemeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemeName
}

// GetSchemeNameOk returns a tuple with the SchemeName field value
// and a boolean to check if the value has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemeName, true
}

// SetSchemeName sets field value
func (o *AddPbkdf2PasswordStorageSchemeRequest) SetSchemeName(v string) {
	o.SchemeName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSchemas() []Enumpbkdf2PasswordStorageSchemeSchemaUrn {
	if o == nil {
		var ret []Enumpbkdf2PasswordStorageSchemeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSchemasOk() ([]Enumpbkdf2PasswordStorageSchemeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPbkdf2PasswordStorageSchemeRequest) SetSchemas(v []Enumpbkdf2PasswordStorageSchemeSchemaUrn) {
	o.Schemas = v
}

// GetDigestAlgorithm returns the DigestAlgorithm field value if set, zero value otherwise.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDigestAlgorithm() EnumpasswordStorageSchemeDigestAlgorithmProp {
	if o == nil || IsNil(o.DigestAlgorithm) {
		var ret EnumpasswordStorageSchemeDigestAlgorithmProp
		return ret
	}
	return *o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDigestAlgorithmOk() (*EnumpasswordStorageSchemeDigestAlgorithmProp, bool) {
	if o == nil || IsNil(o.DigestAlgorithm) {
		return nil, false
	}
	return o.DigestAlgorithm, true
}

// HasDigestAlgorithm returns a boolean if a field has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) HasDigestAlgorithm() bool {
	if o != nil && !IsNil(o.DigestAlgorithm) {
		return true
	}

	return false
}

// SetDigestAlgorithm gets a reference to the given EnumpasswordStorageSchemeDigestAlgorithmProp and assigns it to the DigestAlgorithm field.
func (o *AddPbkdf2PasswordStorageSchemeRequest) SetDigestAlgorithm(v EnumpasswordStorageSchemeDigestAlgorithmProp) {
	o.DigestAlgorithm = &v
}

// GetIterationCount returns the IterationCount field value if set, zero value otherwise.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetIterationCount() int32 {
	if o == nil || IsNil(o.IterationCount) {
		var ret int32
		return ret
	}
	return *o.IterationCount
}

// GetIterationCountOk returns a tuple with the IterationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetIterationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IterationCount) {
		return nil, false
	}
	return o.IterationCount, true
}

// HasIterationCount returns a boolean if a field has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) HasIterationCount() bool {
	if o != nil && !IsNil(o.IterationCount) {
		return true
	}

	return false
}

// SetIterationCount gets a reference to the given int32 and assigns it to the IterationCount field.
func (o *AddPbkdf2PasswordStorageSchemeRequest) SetIterationCount(v int32) {
	o.IterationCount = &v
}

// GetSaltLengthBytes returns the SaltLengthBytes field value if set, zero value otherwise.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSaltLengthBytes() int32 {
	if o == nil || IsNil(o.SaltLengthBytes) {
		var ret int32
		return ret
	}
	return *o.SaltLengthBytes
}

// GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetSaltLengthBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.SaltLengthBytes) {
		return nil, false
	}
	return o.SaltLengthBytes, true
}

// HasSaltLengthBytes returns a boolean if a field has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) HasSaltLengthBytes() bool {
	if o != nil && !IsNil(o.SaltLengthBytes) {
		return true
	}

	return false
}

// SetSaltLengthBytes gets a reference to the given int32 and assigns it to the SaltLengthBytes field.
func (o *AddPbkdf2PasswordStorageSchemeRequest) SetSaltLengthBytes(v int32) {
	o.SaltLengthBytes = &v
}

// GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field value if set, zero value otherwise.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDerivedKeyLengthBytes() int32 {
	if o == nil || IsNil(o.DerivedKeyLengthBytes) {
		var ret int32
		return ret
	}
	return *o.DerivedKeyLengthBytes
}

// GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDerivedKeyLengthBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.DerivedKeyLengthBytes) {
		return nil, false
	}
	return o.DerivedKeyLengthBytes, true
}

// HasDerivedKeyLengthBytes returns a boolean if a field has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) HasDerivedKeyLengthBytes() bool {
	if o != nil && !IsNil(o.DerivedKeyLengthBytes) {
		return true
	}

	return false
}

// SetDerivedKeyLengthBytes gets a reference to the given int32 and assigns it to the DerivedKeyLengthBytes field.
func (o *AddPbkdf2PasswordStorageSchemeRequest) SetDerivedKeyLengthBytes(v int32) {
	o.DerivedKeyLengthBytes = &v
}

// GetMaxPasswordLength returns the MaxPasswordLength field value if set, zero value otherwise.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetMaxPasswordLength() int32 {
	if o == nil || IsNil(o.MaxPasswordLength) {
		var ret int32
		return ret
	}
	return *o.MaxPasswordLength
}

// GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetMaxPasswordLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPasswordLength) {
		return nil, false
	}
	return o.MaxPasswordLength, true
}

// HasMaxPasswordLength returns a boolean if a field has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) HasMaxPasswordLength() bool {
	if o != nil && !IsNil(o.MaxPasswordLength) {
		return true
	}

	return false
}

// SetMaxPasswordLength gets a reference to the given int32 and assigns it to the MaxPasswordLength field.
func (o *AddPbkdf2PasswordStorageSchemeRequest) SetMaxPasswordLength(v int32) {
	o.MaxPasswordLength = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPbkdf2PasswordStorageSchemeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddPbkdf2PasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddPbkdf2PasswordStorageSchemeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddPbkdf2PasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPbkdf2PasswordStorageSchemeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemeName"] = o.SchemeName
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.DigestAlgorithm) {
		toSerialize["digestAlgorithm"] = o.DigestAlgorithm
	}
	if !IsNil(o.IterationCount) {
		toSerialize["iterationCount"] = o.IterationCount
	}
	if !IsNil(o.SaltLengthBytes) {
		toSerialize["saltLengthBytes"] = o.SaltLengthBytes
	}
	if !IsNil(o.DerivedKeyLengthBytes) {
		toSerialize["derivedKeyLengthBytes"] = o.DerivedKeyLengthBytes
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

type NullableAddPbkdf2PasswordStorageSchemeRequest struct {
	value *AddPbkdf2PasswordStorageSchemeRequest
	isSet bool
}

func (v NullableAddPbkdf2PasswordStorageSchemeRequest) Get() *AddPbkdf2PasswordStorageSchemeRequest {
	return v.value
}

func (v *NullableAddPbkdf2PasswordStorageSchemeRequest) Set(val *AddPbkdf2PasswordStorageSchemeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPbkdf2PasswordStorageSchemeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPbkdf2PasswordStorageSchemeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPbkdf2PasswordStorageSchemeRequest(val *AddPbkdf2PasswordStorageSchemeRequest) *NullableAddPbkdf2PasswordStorageSchemeRequest {
	return &NullableAddPbkdf2PasswordStorageSchemeRequest{value: val, isSet: true}
}

func (v NullableAddPbkdf2PasswordStorageSchemeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPbkdf2PasswordStorageSchemeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
