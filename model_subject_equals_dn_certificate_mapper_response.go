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

// SubjectEqualsDnCertificateMapperResponse struct for SubjectEqualsDnCertificateMapperResponse
type SubjectEqualsDnCertificateMapperResponse struct {
	// Name of the Certificate Mapper
	Id string `json:"id"`
	Schemas []EnumsubjectEqualsDnCertificateMapperSchemaUrn `json:"schemas"`
	// A description for this Certificate Mapper
	Description *string `json:"description,omitempty"`
	// Indicates whether the Certificate Mapper is enabled.
	Enabled bool `json:"enabled"`
}

// NewSubjectEqualsDnCertificateMapperResponse instantiates a new SubjectEqualsDnCertificateMapperResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectEqualsDnCertificateMapperResponse(id string, schemas []EnumsubjectEqualsDnCertificateMapperSchemaUrn, enabled bool) *SubjectEqualsDnCertificateMapperResponse {
	this := SubjectEqualsDnCertificateMapperResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewSubjectEqualsDnCertificateMapperResponseWithDefaults instantiates a new SubjectEqualsDnCertificateMapperResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectEqualsDnCertificateMapperResponseWithDefaults() *SubjectEqualsDnCertificateMapperResponse {
	this := SubjectEqualsDnCertificateMapperResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SubjectEqualsDnCertificateMapperResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubjectEqualsDnCertificateMapperResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubjectEqualsDnCertificateMapperResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *SubjectEqualsDnCertificateMapperResponse) GetSchemas() []EnumsubjectEqualsDnCertificateMapperSchemaUrn {
	if o == nil {
		var ret []EnumsubjectEqualsDnCertificateMapperSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SubjectEqualsDnCertificateMapperResponse) GetSchemasOk() ([]EnumsubjectEqualsDnCertificateMapperSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SubjectEqualsDnCertificateMapperResponse) SetSchemas(v []EnumsubjectEqualsDnCertificateMapperSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubjectEqualsDnCertificateMapperResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectEqualsDnCertificateMapperResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubjectEqualsDnCertificateMapperResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SubjectEqualsDnCertificateMapperResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SubjectEqualsDnCertificateMapperResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SubjectEqualsDnCertificateMapperResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SubjectEqualsDnCertificateMapperResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o SubjectEqualsDnCertificateMapperResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSubjectEqualsDnCertificateMapperResponse struct {
	value *SubjectEqualsDnCertificateMapperResponse
	isSet bool
}

func (v NullableSubjectEqualsDnCertificateMapperResponse) Get() *SubjectEqualsDnCertificateMapperResponse {
	return v.value
}

func (v *NullableSubjectEqualsDnCertificateMapperResponse) Set(val *SubjectEqualsDnCertificateMapperResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectEqualsDnCertificateMapperResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectEqualsDnCertificateMapperResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectEqualsDnCertificateMapperResponse(val *SubjectEqualsDnCertificateMapperResponse) *NullableSubjectEqualsDnCertificateMapperResponse {
	return &NullableSubjectEqualsDnCertificateMapperResponse{value: val, isSet: true}
}

func (v NullableSubjectEqualsDnCertificateMapperResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectEqualsDnCertificateMapperResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


