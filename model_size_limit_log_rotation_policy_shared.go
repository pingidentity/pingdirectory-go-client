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

// SizeLimitLogRotationPolicyShared struct for SizeLimitLogRotationPolicyShared
type SizeLimitLogRotationPolicyShared struct {
	Schemas []EnumsizeLimitLogRotationPolicySchemaUrn `json:"schemas"`
	// Specifies the maximum size that a log file can reach before it is rotated.
	FileSizeLimit string `json:"fileSizeLimit"`
	// A description for this Log Rotation Policy
	Description *string `json:"description,omitempty"`
}

// NewSizeLimitLogRotationPolicyShared instantiates a new SizeLimitLogRotationPolicyShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSizeLimitLogRotationPolicyShared(schemas []EnumsizeLimitLogRotationPolicySchemaUrn, fileSizeLimit string) *SizeLimitLogRotationPolicyShared {
	this := SizeLimitLogRotationPolicyShared{}
	this.Schemas = schemas
	this.FileSizeLimit = fileSizeLimit
	return &this
}

// NewSizeLimitLogRotationPolicySharedWithDefaults instantiates a new SizeLimitLogRotationPolicyShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSizeLimitLogRotationPolicySharedWithDefaults() *SizeLimitLogRotationPolicyShared {
	this := SizeLimitLogRotationPolicyShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *SizeLimitLogRotationPolicyShared) GetSchemas() []EnumsizeLimitLogRotationPolicySchemaUrn {
	if o == nil {
		var ret []EnumsizeLimitLogRotationPolicySchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SizeLimitLogRotationPolicyShared) GetSchemasOk() ([]EnumsizeLimitLogRotationPolicySchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SizeLimitLogRotationPolicyShared) SetSchemas(v []EnumsizeLimitLogRotationPolicySchemaUrn) {
	o.Schemas = v
}

// GetFileSizeLimit returns the FileSizeLimit field value
func (o *SizeLimitLogRotationPolicyShared) GetFileSizeLimit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileSizeLimit
}

// GetFileSizeLimitOk returns a tuple with the FileSizeLimit field value
// and a boolean to check if the value has been set.
func (o *SizeLimitLogRotationPolicyShared) GetFileSizeLimitOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FileSizeLimit, true
}

// SetFileSizeLimit sets field value
func (o *SizeLimitLogRotationPolicyShared) SetFileSizeLimit(v string) {
	o.FileSizeLimit = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SizeLimitLogRotationPolicyShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SizeLimitLogRotationPolicyShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SizeLimitLogRotationPolicyShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SizeLimitLogRotationPolicyShared) SetDescription(v string) {
	o.Description = &v
}

func (o SizeLimitLogRotationPolicyShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["fileSizeLimit"] = o.FileSizeLimit
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableSizeLimitLogRotationPolicyShared struct {
	value *SizeLimitLogRotationPolicyShared
	isSet bool
}

func (v NullableSizeLimitLogRotationPolicyShared) Get() *SizeLimitLogRotationPolicyShared {
	return v.value
}

func (v *NullableSizeLimitLogRotationPolicyShared) Set(val *SizeLimitLogRotationPolicyShared) {
	v.value = val
	v.isSet = true
}

func (v NullableSizeLimitLogRotationPolicyShared) IsSet() bool {
	return v.isSet
}

func (v *NullableSizeLimitLogRotationPolicyShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSizeLimitLogRotationPolicyShared(val *SizeLimitLogRotationPolicyShared) *NullableSizeLimitLogRotationPolicyShared {
	return &NullableSizeLimitLogRotationPolicyShared{value: val, isSet: true}
}

func (v NullableSizeLimitLogRotationPolicyShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSizeLimitLogRotationPolicyShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


