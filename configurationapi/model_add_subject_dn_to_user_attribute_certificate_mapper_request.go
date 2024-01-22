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

// checks if the AddSubjectDnToUserAttributeCertificateMapperRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSubjectDnToUserAttributeCertificateMapperRequest{}

// AddSubjectDnToUserAttributeCertificateMapperRequest struct for AddSubjectDnToUserAttributeCertificateMapperRequest
type AddSubjectDnToUserAttributeCertificateMapperRequest struct {
	Schemas []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn `json:"schemas"`
	// Specifies the name or OID of the attribute whose value should exactly match the certificate subject DN.
	SubjectAttribute *string `json:"subjectAttribute,omitempty"`
	// Specifies the base DNs that should be used when performing searches to map the client certificate to a user entry.
	UserBaseDN []string `json:"userBaseDN,omitempty"`
	// A description for this Certificate Mapper
	Description *string `json:"description,omitempty"`
	// Indicates whether the Certificate Mapper is enabled.
	Enabled bool `json:"enabled"`
	// Name of the new Certificate Mapper
	MapperName string `json:"mapperName"`
}

// NewAddSubjectDnToUserAttributeCertificateMapperRequest instantiates a new AddSubjectDnToUserAttributeCertificateMapperRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSubjectDnToUserAttributeCertificateMapperRequest(schemas []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn, enabled bool, mapperName string) *AddSubjectDnToUserAttributeCertificateMapperRequest {
	this := AddSubjectDnToUserAttributeCertificateMapperRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.MapperName = mapperName
	return &this
}

// NewAddSubjectDnToUserAttributeCertificateMapperRequestWithDefaults instantiates a new AddSubjectDnToUserAttributeCertificateMapperRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSubjectDnToUserAttributeCertificateMapperRequestWithDefaults() *AddSubjectDnToUserAttributeCertificateMapperRequest {
	this := AddSubjectDnToUserAttributeCertificateMapperRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetSchemas() []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn {
	if o == nil {
		var ret []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetSchemasOk() ([]EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetSchemas(v []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) {
	o.Schemas = v
}

// GetSubjectAttribute returns the SubjectAttribute field value if set, zero value otherwise.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetSubjectAttribute() string {
	if o == nil || IsNil(o.SubjectAttribute) {
		var ret string
		return ret
	}
	return *o.SubjectAttribute
}

// GetSubjectAttributeOk returns a tuple with the SubjectAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetSubjectAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectAttribute) {
		return nil, false
	}
	return o.SubjectAttribute, true
}

// HasSubjectAttribute returns a boolean if a field has been set.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) HasSubjectAttribute() bool {
	if o != nil && !IsNil(o.SubjectAttribute) {
		return true
	}

	return false
}

// SetSubjectAttribute gets a reference to the given string and assigns it to the SubjectAttribute field.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetSubjectAttribute(v string) {
	o.SubjectAttribute = &v
}

// GetUserBaseDN returns the UserBaseDN field value if set, zero value otherwise.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetUserBaseDN() []string {
	if o == nil || IsNil(o.UserBaseDN) {
		var ret []string
		return ret
	}
	return o.UserBaseDN
}

// GetUserBaseDNOk returns a tuple with the UserBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetUserBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.UserBaseDN) {
		return nil, false
	}
	return o.UserBaseDN, true
}

// HasUserBaseDN returns a boolean if a field has been set.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) HasUserBaseDN() bool {
	if o != nil && !IsNil(o.UserBaseDN) {
		return true
	}

	return false
}

// SetUserBaseDN gets a reference to the given []string and assigns it to the UserBaseDN field.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetUserBaseDN(v []string) {
	o.UserBaseDN = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMapperName returns the MapperName field value
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetMapperName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MapperName
}

// GetMapperNameOk returns a tuple with the MapperName field value
// and a boolean to check if the value has been set.
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) GetMapperNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapperName, true
}

// SetMapperName sets field value
func (o *AddSubjectDnToUserAttributeCertificateMapperRequest) SetMapperName(v string) {
	o.MapperName = v
}

func (o AddSubjectDnToUserAttributeCertificateMapperRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSubjectDnToUserAttributeCertificateMapperRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.SubjectAttribute) {
		toSerialize["subjectAttribute"] = o.SubjectAttribute
	}
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

type NullableAddSubjectDnToUserAttributeCertificateMapperRequest struct {
	value *AddSubjectDnToUserAttributeCertificateMapperRequest
	isSet bool
}

func (v NullableAddSubjectDnToUserAttributeCertificateMapperRequest) Get() *AddSubjectDnToUserAttributeCertificateMapperRequest {
	return v.value
}

func (v *NullableAddSubjectDnToUserAttributeCertificateMapperRequest) Set(val *AddSubjectDnToUserAttributeCertificateMapperRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSubjectDnToUserAttributeCertificateMapperRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSubjectDnToUserAttributeCertificateMapperRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSubjectDnToUserAttributeCertificateMapperRequest(val *AddSubjectDnToUserAttributeCertificateMapperRequest) *NullableAddSubjectDnToUserAttributeCertificateMapperRequest {
	return &NullableAddSubjectDnToUserAttributeCertificateMapperRequest{value: val, isSet: true}
}

func (v NullableAddSubjectDnToUserAttributeCertificateMapperRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSubjectDnToUserAttributeCertificateMapperRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
