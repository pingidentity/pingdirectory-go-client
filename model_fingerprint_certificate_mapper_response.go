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

// FingerprintCertificateMapperResponse struct for FingerprintCertificateMapperResponse
type FingerprintCertificateMapperResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Certificate Mapper
	Id string `json:"id"`
	Schemas []EnumfingerprintCertificateMapperSchemaUrn `json:"schemas"`
	// Specifies the attribute in which to look for the fingerprint.
	FingerprintAttribute string `json:"fingerprintAttribute"`
	FingerprintAlgorithm EnumcertificateMapperFingerprintAlgorithmProp `json:"fingerprintAlgorithm"`
	// Specifies the set of base DNs below which to search for users.
	UserBaseDN []string `json:"userBaseDN,omitempty"`
	// A description for this Certificate Mapper
	Description *string `json:"description,omitempty"`
	// Indicates whether the Certificate Mapper is enabled.
	Enabled bool `json:"enabled"`
}

// NewFingerprintCertificateMapperResponse instantiates a new FingerprintCertificateMapperResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFingerprintCertificateMapperResponse(id string, schemas []EnumfingerprintCertificateMapperSchemaUrn, fingerprintAttribute string, fingerprintAlgorithm EnumcertificateMapperFingerprintAlgorithmProp, enabled bool) *FingerprintCertificateMapperResponse {
	this := FingerprintCertificateMapperResponse{}
	this.Id = id
	this.Schemas = schemas
	this.FingerprintAttribute = fingerprintAttribute
	this.FingerprintAlgorithm = fingerprintAlgorithm
	this.Enabled = enabled
	return &this
}

// NewFingerprintCertificateMapperResponseWithDefaults instantiates a new FingerprintCertificateMapperResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFingerprintCertificateMapperResponseWithDefaults() *FingerprintCertificateMapperResponse {
	this := FingerprintCertificateMapperResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FingerprintCertificateMapperResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FingerprintCertificateMapperResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FingerprintCertificateMapperResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *FingerprintCertificateMapperResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *FingerprintCertificateMapperResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FingerprintCertificateMapperResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *FingerprintCertificateMapperResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *FingerprintCertificateMapperResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *FingerprintCertificateMapperResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FingerprintCertificateMapperResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FingerprintCertificateMapperResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *FingerprintCertificateMapperResponse) GetSchemas() []EnumfingerprintCertificateMapperSchemaUrn {
	if o == nil {
		var ret []EnumfingerprintCertificateMapperSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *FingerprintCertificateMapperResponse) GetSchemasOk() ([]EnumfingerprintCertificateMapperSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *FingerprintCertificateMapperResponse) SetSchemas(v []EnumfingerprintCertificateMapperSchemaUrn) {
	o.Schemas = v
}

// GetFingerprintAttribute returns the FingerprintAttribute field value
func (o *FingerprintCertificateMapperResponse) GetFingerprintAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FingerprintAttribute
}

// GetFingerprintAttributeOk returns a tuple with the FingerprintAttribute field value
// and a boolean to check if the value has been set.
func (o *FingerprintCertificateMapperResponse) GetFingerprintAttributeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FingerprintAttribute, true
}

// SetFingerprintAttribute sets field value
func (o *FingerprintCertificateMapperResponse) SetFingerprintAttribute(v string) {
	o.FingerprintAttribute = v
}

// GetFingerprintAlgorithm returns the FingerprintAlgorithm field value
func (o *FingerprintCertificateMapperResponse) GetFingerprintAlgorithm() EnumcertificateMapperFingerprintAlgorithmProp {
	if o == nil {
		var ret EnumcertificateMapperFingerprintAlgorithmProp
		return ret
	}

	return o.FingerprintAlgorithm
}

// GetFingerprintAlgorithmOk returns a tuple with the FingerprintAlgorithm field value
// and a boolean to check if the value has been set.
func (o *FingerprintCertificateMapperResponse) GetFingerprintAlgorithmOk() (*EnumcertificateMapperFingerprintAlgorithmProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FingerprintAlgorithm, true
}

// SetFingerprintAlgorithm sets field value
func (o *FingerprintCertificateMapperResponse) SetFingerprintAlgorithm(v EnumcertificateMapperFingerprintAlgorithmProp) {
	o.FingerprintAlgorithm = v
}

// GetUserBaseDN returns the UserBaseDN field value if set, zero value otherwise.
func (o *FingerprintCertificateMapperResponse) GetUserBaseDN() []string {
	if o == nil || isNil(o.UserBaseDN) {
		var ret []string
		return ret
	}
	return o.UserBaseDN
}

// GetUserBaseDNOk returns a tuple with the UserBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FingerprintCertificateMapperResponse) GetUserBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.UserBaseDN) {
    return nil, false
	}
	return o.UserBaseDN, true
}

// HasUserBaseDN returns a boolean if a field has been set.
func (o *FingerprintCertificateMapperResponse) HasUserBaseDN() bool {
	if o != nil && !isNil(o.UserBaseDN) {
		return true
	}

	return false
}

// SetUserBaseDN gets a reference to the given []string and assigns it to the UserBaseDN field.
func (o *FingerprintCertificateMapperResponse) SetUserBaseDN(v []string) {
	o.UserBaseDN = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FingerprintCertificateMapperResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FingerprintCertificateMapperResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FingerprintCertificateMapperResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FingerprintCertificateMapperResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *FingerprintCertificateMapperResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FingerprintCertificateMapperResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FingerprintCertificateMapperResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o FingerprintCertificateMapperResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["fingerprintAttribute"] = o.FingerprintAttribute
	}
	if true {
		toSerialize["fingerprintAlgorithm"] = o.FingerprintAlgorithm
	}
	if !isNil(o.UserBaseDN) {
		toSerialize["userBaseDN"] = o.UserBaseDN
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableFingerprintCertificateMapperResponse struct {
	value *FingerprintCertificateMapperResponse
	isSet bool
}

func (v NullableFingerprintCertificateMapperResponse) Get() *FingerprintCertificateMapperResponse {
	return v.value
}

func (v *NullableFingerprintCertificateMapperResponse) Set(val *FingerprintCertificateMapperResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFingerprintCertificateMapperResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFingerprintCertificateMapperResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFingerprintCertificateMapperResponse(val *FingerprintCertificateMapperResponse) *NullableFingerprintCertificateMapperResponse {
	return &NullableFingerprintCertificateMapperResponse{value: val, isSet: true}
}

func (v NullableFingerprintCertificateMapperResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFingerprintCertificateMapperResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


