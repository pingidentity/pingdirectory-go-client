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

// SubjectDnToUserAttributeCertificateMapperResponse struct for SubjectDnToUserAttributeCertificateMapperResponse
type SubjectDnToUserAttributeCertificateMapperResponse struct {
	// Name of the Certificate Mapper
	Id string `json:"id"`
	Schemas []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn `json:"schemas"`
	// Specifies the name or OID of the attribute whose value should exactly match the certificate subject DN.
	SubjectAttribute string `json:"subjectAttribute"`
	// Specifies the base DNs that should be used when performing searches to map the client certificate to a user entry.
	UserBaseDN []string `json:"userBaseDN,omitempty"`
	// A description for this Certificate Mapper
	Description *string `json:"description,omitempty"`
	// Indicates whether the Certificate Mapper is enabled.
	Enabled bool `json:"enabled"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewSubjectDnToUserAttributeCertificateMapperResponse instantiates a new SubjectDnToUserAttributeCertificateMapperResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectDnToUserAttributeCertificateMapperResponse(id string, schemas []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn, subjectAttribute string, enabled bool) *SubjectDnToUserAttributeCertificateMapperResponse {
	this := SubjectDnToUserAttributeCertificateMapperResponse{}
	this.Id = id
	this.Schemas = schemas
	this.SubjectAttribute = subjectAttribute
	this.Enabled = enabled
	return &this
}

// NewSubjectDnToUserAttributeCertificateMapperResponseWithDefaults instantiates a new SubjectDnToUserAttributeCertificateMapperResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectDnToUserAttributeCertificateMapperResponseWithDefaults() *SubjectDnToUserAttributeCertificateMapperResponse {
	this := SubjectDnToUserAttributeCertificateMapperResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetSchemas() []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn {
	if o == nil {
		var ret []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetSchemasOk() ([]EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetSchemas(v []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) {
	o.Schemas = v
}

// GetSubjectAttribute returns the SubjectAttribute field value
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetSubjectAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectAttribute
}

// GetSubjectAttributeOk returns a tuple with the SubjectAttribute field value
// and a boolean to check if the value has been set.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetSubjectAttributeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubjectAttribute, true
}

// SetSubjectAttribute sets field value
func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetSubjectAttribute(v string) {
	o.SubjectAttribute = v
}

// GetUserBaseDN returns the UserBaseDN field value if set, zero value otherwise.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetUserBaseDN() []string {
	if o == nil || isNil(o.UserBaseDN) {
		var ret []string
		return ret
	}
	return o.UserBaseDN
}

// GetUserBaseDNOk returns a tuple with the UserBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetUserBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.UserBaseDN) {
    return nil, false
	}
	return o.UserBaseDN, true
}

// HasUserBaseDN returns a boolean if a field has been set.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) HasUserBaseDN() bool {
	if o != nil && !isNil(o.UserBaseDN) {
		return true
	}

	return false
}

// SetUserBaseDN gets a reference to the given []string and assigns it to the UserBaseDN field.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetUserBaseDN(v []string) {
	o.UserBaseDN = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SubjectDnToUserAttributeCertificateMapperResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o SubjectDnToUserAttributeCertificateMapperResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["subjectAttribute"] = o.SubjectAttribute
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
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableSubjectDnToUserAttributeCertificateMapperResponse struct {
	value *SubjectDnToUserAttributeCertificateMapperResponse
	isSet bool
}

func (v NullableSubjectDnToUserAttributeCertificateMapperResponse) Get() *SubjectDnToUserAttributeCertificateMapperResponse {
	return v.value
}

func (v *NullableSubjectDnToUserAttributeCertificateMapperResponse) Set(val *SubjectDnToUserAttributeCertificateMapperResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectDnToUserAttributeCertificateMapperResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectDnToUserAttributeCertificateMapperResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectDnToUserAttributeCertificateMapperResponse(val *SubjectDnToUserAttributeCertificateMapperResponse) *NullableSubjectDnToUserAttributeCertificateMapperResponse {
	return &NullableSubjectDnToUserAttributeCertificateMapperResponse{value: val, isSet: true}
}

func (v NullableSubjectDnToUserAttributeCertificateMapperResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectDnToUserAttributeCertificateMapperResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


