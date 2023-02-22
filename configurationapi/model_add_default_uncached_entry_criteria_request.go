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

// AddDefaultUncachedEntryCriteriaRequest struct for AddDefaultUncachedEntryCriteriaRequest
type AddDefaultUncachedEntryCriteriaRequest struct {
	// Name of the new Uncached Entry Criteria
	CriteriaName string                                      `json:"criteriaName"`
	Schemas      []EnumdefaultUncachedEntryCriteriaSchemaUrn `json:"schemas"`
	// A description for this Uncached Entry Criteria
	Description *string `json:"description,omitempty"`
	// Indicates whether this Uncached Entry Criteria is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewAddDefaultUncachedEntryCriteriaRequest instantiates a new AddDefaultUncachedEntryCriteriaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDefaultUncachedEntryCriteriaRequest(criteriaName string, schemas []EnumdefaultUncachedEntryCriteriaSchemaUrn, enabled bool) *AddDefaultUncachedEntryCriteriaRequest {
	this := AddDefaultUncachedEntryCriteriaRequest{}
	this.CriteriaName = criteriaName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddDefaultUncachedEntryCriteriaRequestWithDefaults instantiates a new AddDefaultUncachedEntryCriteriaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDefaultUncachedEntryCriteriaRequestWithDefaults() *AddDefaultUncachedEntryCriteriaRequest {
	this := AddDefaultUncachedEntryCriteriaRequest{}
	return &this
}

// GetCriteriaName returns the CriteriaName field value
func (o *AddDefaultUncachedEntryCriteriaRequest) GetCriteriaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CriteriaName
}

// GetCriteriaNameOk returns a tuple with the CriteriaName field value
// and a boolean to check if the value has been set.
func (o *AddDefaultUncachedEntryCriteriaRequest) GetCriteriaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CriteriaName, true
}

// SetCriteriaName sets field value
func (o *AddDefaultUncachedEntryCriteriaRequest) SetCriteriaName(v string) {
	o.CriteriaName = v
}

// GetSchemas returns the Schemas field value
func (o *AddDefaultUncachedEntryCriteriaRequest) GetSchemas() []EnumdefaultUncachedEntryCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumdefaultUncachedEntryCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddDefaultUncachedEntryCriteriaRequest) GetSchemasOk() ([]EnumdefaultUncachedEntryCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddDefaultUncachedEntryCriteriaRequest) SetSchemas(v []EnumdefaultUncachedEntryCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddDefaultUncachedEntryCriteriaRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDefaultUncachedEntryCriteriaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddDefaultUncachedEntryCriteriaRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddDefaultUncachedEntryCriteriaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddDefaultUncachedEntryCriteriaRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddDefaultUncachedEntryCriteriaRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddDefaultUncachedEntryCriteriaRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddDefaultUncachedEntryCriteriaRequest) MarshalJSON() ([]byte, error) {
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

type NullableAddDefaultUncachedEntryCriteriaRequest struct {
	value *AddDefaultUncachedEntryCriteriaRequest
	isSet bool
}

func (v NullableAddDefaultUncachedEntryCriteriaRequest) Get() *AddDefaultUncachedEntryCriteriaRequest {
	return v.value
}

func (v *NullableAddDefaultUncachedEntryCriteriaRequest) Set(val *AddDefaultUncachedEntryCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDefaultUncachedEntryCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDefaultUncachedEntryCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDefaultUncachedEntryCriteriaRequest(val *AddDefaultUncachedEntryCriteriaRequest) *NullableAddDefaultUncachedEntryCriteriaRequest {
	return &NullableAddDefaultUncachedEntryCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddDefaultUncachedEntryCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDefaultUncachedEntryCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}