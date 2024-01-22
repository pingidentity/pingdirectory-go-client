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

// checks if the AddFingerprintCertificateMapperRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddFingerprintCertificateMapperRequest{}

// AddFingerprintCertificateMapperRequest struct for AddFingerprintCertificateMapperRequest
type AddFingerprintCertificateMapperRequest struct {
	Schemas []EnumfingerprintCertificateMapperSchemaUrn `json:"schemas"`
	// Specifies the attribute in which to look for the fingerprint.
	FingerprintAttribute *string                                       `json:"fingerprintAttribute,omitempty"`
	FingerprintAlgorithm EnumcertificateMapperFingerprintAlgorithmProp `json:"fingerprintAlgorithm"`
	// Specifies the set of base DNs below which to search for users.
	UserBaseDN []string `json:"userBaseDN,omitempty"`
	// A description for this Certificate Mapper
	Description *string `json:"description,omitempty"`
	// Indicates whether the Certificate Mapper is enabled.
	Enabled bool `json:"enabled"`
	// Name of the new Certificate Mapper
	MapperName string `json:"mapperName"`
}

// NewAddFingerprintCertificateMapperRequest instantiates a new AddFingerprintCertificateMapperRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFingerprintCertificateMapperRequest(schemas []EnumfingerprintCertificateMapperSchemaUrn, fingerprintAlgorithm EnumcertificateMapperFingerprintAlgorithmProp, enabled bool, mapperName string) *AddFingerprintCertificateMapperRequest {
	this := AddFingerprintCertificateMapperRequest{}
	this.Schemas = schemas
	this.FingerprintAlgorithm = fingerprintAlgorithm
	this.Enabled = enabled
	this.MapperName = mapperName
	return &this
}

// NewAddFingerprintCertificateMapperRequestWithDefaults instantiates a new AddFingerprintCertificateMapperRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFingerprintCertificateMapperRequestWithDefaults() *AddFingerprintCertificateMapperRequest {
	this := AddFingerprintCertificateMapperRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddFingerprintCertificateMapperRequest) GetSchemas() []EnumfingerprintCertificateMapperSchemaUrn {
	if o == nil {
		var ret []EnumfingerprintCertificateMapperSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddFingerprintCertificateMapperRequest) GetSchemasOk() ([]EnumfingerprintCertificateMapperSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddFingerprintCertificateMapperRequest) SetSchemas(v []EnumfingerprintCertificateMapperSchemaUrn) {
	o.Schemas = v
}

// GetFingerprintAttribute returns the FingerprintAttribute field value if set, zero value otherwise.
func (o *AddFingerprintCertificateMapperRequest) GetFingerprintAttribute() string {
	if o == nil || IsNil(o.FingerprintAttribute) {
		var ret string
		return ret
	}
	return *o.FingerprintAttribute
}

// GetFingerprintAttributeOk returns a tuple with the FingerprintAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFingerprintCertificateMapperRequest) GetFingerprintAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.FingerprintAttribute) {
		return nil, false
	}
	return o.FingerprintAttribute, true
}

// HasFingerprintAttribute returns a boolean if a field has been set.
func (o *AddFingerprintCertificateMapperRequest) HasFingerprintAttribute() bool {
	if o != nil && !IsNil(o.FingerprintAttribute) {
		return true
	}

	return false
}

// SetFingerprintAttribute gets a reference to the given string and assigns it to the FingerprintAttribute field.
func (o *AddFingerprintCertificateMapperRequest) SetFingerprintAttribute(v string) {
	o.FingerprintAttribute = &v
}

// GetFingerprintAlgorithm returns the FingerprintAlgorithm field value
func (o *AddFingerprintCertificateMapperRequest) GetFingerprintAlgorithm() EnumcertificateMapperFingerprintAlgorithmProp {
	if o == nil {
		var ret EnumcertificateMapperFingerprintAlgorithmProp
		return ret
	}

	return o.FingerprintAlgorithm
}

// GetFingerprintAlgorithmOk returns a tuple with the FingerprintAlgorithm field value
// and a boolean to check if the value has been set.
func (o *AddFingerprintCertificateMapperRequest) GetFingerprintAlgorithmOk() (*EnumcertificateMapperFingerprintAlgorithmProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FingerprintAlgorithm, true
}

// SetFingerprintAlgorithm sets field value
func (o *AddFingerprintCertificateMapperRequest) SetFingerprintAlgorithm(v EnumcertificateMapperFingerprintAlgorithmProp) {
	o.FingerprintAlgorithm = v
}

// GetUserBaseDN returns the UserBaseDN field value if set, zero value otherwise.
func (o *AddFingerprintCertificateMapperRequest) GetUserBaseDN() []string {
	if o == nil || IsNil(o.UserBaseDN) {
		var ret []string
		return ret
	}
	return o.UserBaseDN
}

// GetUserBaseDNOk returns a tuple with the UserBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFingerprintCertificateMapperRequest) GetUserBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.UserBaseDN) {
		return nil, false
	}
	return o.UserBaseDN, true
}

// HasUserBaseDN returns a boolean if a field has been set.
func (o *AddFingerprintCertificateMapperRequest) HasUserBaseDN() bool {
	if o != nil && !IsNil(o.UserBaseDN) {
		return true
	}

	return false
}

// SetUserBaseDN gets a reference to the given []string and assigns it to the UserBaseDN field.
func (o *AddFingerprintCertificateMapperRequest) SetUserBaseDN(v []string) {
	o.UserBaseDN = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddFingerprintCertificateMapperRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddFingerprintCertificateMapperRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddFingerprintCertificateMapperRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddFingerprintCertificateMapperRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddFingerprintCertificateMapperRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddFingerprintCertificateMapperRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddFingerprintCertificateMapperRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMapperName returns the MapperName field value
func (o *AddFingerprintCertificateMapperRequest) GetMapperName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MapperName
}

// GetMapperNameOk returns a tuple with the MapperName field value
// and a boolean to check if the value has been set.
func (o *AddFingerprintCertificateMapperRequest) GetMapperNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapperName, true
}

// SetMapperName sets field value
func (o *AddFingerprintCertificateMapperRequest) SetMapperName(v string) {
	o.MapperName = v
}

func (o AddFingerprintCertificateMapperRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddFingerprintCertificateMapperRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.FingerprintAttribute) {
		toSerialize["fingerprintAttribute"] = o.FingerprintAttribute
	}
	toSerialize["fingerprintAlgorithm"] = o.FingerprintAlgorithm
	if !IsNil(o.UserBaseDN) {
		toSerialize["userBaseDN"] = o.UserBaseDN
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["mapperName"] = o.MapperName
	return toSerialize, nil
}

type NullableAddFingerprintCertificateMapperRequest struct {
	value *AddFingerprintCertificateMapperRequest
	isSet bool
}

func (v NullableAddFingerprintCertificateMapperRequest) Get() *AddFingerprintCertificateMapperRequest {
	return v.value
}

func (v *NullableAddFingerprintCertificateMapperRequest) Set(val *AddFingerprintCertificateMapperRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFingerprintCertificateMapperRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFingerprintCertificateMapperRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFingerprintCertificateMapperRequest(val *AddFingerprintCertificateMapperRequest) *NullableAddFingerprintCertificateMapperRequest {
	return &NullableAddFingerprintCertificateMapperRequest{value: val, isSet: true}
}

func (v NullableAddFingerprintCertificateMapperRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFingerprintCertificateMapperRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
