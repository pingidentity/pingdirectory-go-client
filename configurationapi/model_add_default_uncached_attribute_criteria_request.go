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

// AddDefaultUncachedAttributeCriteriaRequest struct for AddDefaultUncachedAttributeCriteriaRequest
type AddDefaultUncachedAttributeCriteriaRequest struct {
	// Name of the new Uncached Attribute Criteria
	CriteriaName string                                          `json:"criteriaName"`
	Schemas      []EnumdefaultUncachedAttributeCriteriaSchemaUrn `json:"schemas"`
	// A description for this Uncached Attribute Criteria
	Description *string `json:"description,omitempty"`
	// Indicates whether this Uncached Attribute Criteria is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewAddDefaultUncachedAttributeCriteriaRequest instantiates a new AddDefaultUncachedAttributeCriteriaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDefaultUncachedAttributeCriteriaRequest(criteriaName string, schemas []EnumdefaultUncachedAttributeCriteriaSchemaUrn, enabled bool) *AddDefaultUncachedAttributeCriteriaRequest {
	this := AddDefaultUncachedAttributeCriteriaRequest{}
	this.CriteriaName = criteriaName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddDefaultUncachedAttributeCriteriaRequestWithDefaults instantiates a new AddDefaultUncachedAttributeCriteriaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDefaultUncachedAttributeCriteriaRequestWithDefaults() *AddDefaultUncachedAttributeCriteriaRequest {
	this := AddDefaultUncachedAttributeCriteriaRequest{}
	return &this
}

// GetCriteriaName returns the CriteriaName field value
func (o *AddDefaultUncachedAttributeCriteriaRequest) GetCriteriaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CriteriaName
}

// GetCriteriaNameOk returns a tuple with the CriteriaName field value
// and a boolean to check if the value has been set.
func (o *AddDefaultUncachedAttributeCriteriaRequest) GetCriteriaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CriteriaName, true
}

// SetCriteriaName sets field value
func (o *AddDefaultUncachedAttributeCriteriaRequest) SetCriteriaName(v string) {
	o.CriteriaName = v
}

// GetSchemas returns the Schemas field value
func (o *AddDefaultUncachedAttributeCriteriaRequest) GetSchemas() []EnumdefaultUncachedAttributeCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumdefaultUncachedAttributeCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddDefaultUncachedAttributeCriteriaRequest) GetSchemasOk() ([]EnumdefaultUncachedAttributeCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddDefaultUncachedAttributeCriteriaRequest) SetSchemas(v []EnumdefaultUncachedAttributeCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddDefaultUncachedAttributeCriteriaRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDefaultUncachedAttributeCriteriaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddDefaultUncachedAttributeCriteriaRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddDefaultUncachedAttributeCriteriaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddDefaultUncachedAttributeCriteriaRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddDefaultUncachedAttributeCriteriaRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddDefaultUncachedAttributeCriteriaRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddDefaultUncachedAttributeCriteriaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["criteriaName"] = o.CriteriaName
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

type NullableAddDefaultUncachedAttributeCriteriaRequest struct {
	value *AddDefaultUncachedAttributeCriteriaRequest
	isSet bool
}

func (v NullableAddDefaultUncachedAttributeCriteriaRequest) Get() *AddDefaultUncachedAttributeCriteriaRequest {
	return v.value
}

func (v *NullableAddDefaultUncachedAttributeCriteriaRequest) Set(val *AddDefaultUncachedAttributeCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDefaultUncachedAttributeCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDefaultUncachedAttributeCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDefaultUncachedAttributeCriteriaRequest(val *AddDefaultUncachedAttributeCriteriaRequest) *NullableAddDefaultUncachedAttributeCriteriaRequest {
	return &NullableAddDefaultUncachedAttributeCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddDefaultUncachedAttributeCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDefaultUncachedAttributeCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}