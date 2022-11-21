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

// DefaultUncachedAttributeCriteriaResponse struct for DefaultUncachedAttributeCriteriaResponse
type DefaultUncachedAttributeCriteriaResponse struct {
	// Name of the Uncached Attribute Criteria
	Id string `json:"id"`
	Schemas []EnumdefaultUncachedAttributeCriteriaSchemaUrn `json:"schemas"`
	// A description for this Uncached Attribute Criteria
	Description *string `json:"description,omitempty"`
	// Indicates whether this Uncached Attribute Criteria is enabled for use in the server.
	Enabled bool `json:"enabled"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewDefaultUncachedAttributeCriteriaResponse instantiates a new DefaultUncachedAttributeCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultUncachedAttributeCriteriaResponse(id string, schemas []EnumdefaultUncachedAttributeCriteriaSchemaUrn, enabled bool) *DefaultUncachedAttributeCriteriaResponse {
	this := DefaultUncachedAttributeCriteriaResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewDefaultUncachedAttributeCriteriaResponseWithDefaults instantiates a new DefaultUncachedAttributeCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultUncachedAttributeCriteriaResponseWithDefaults() *DefaultUncachedAttributeCriteriaResponse {
	this := DefaultUncachedAttributeCriteriaResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DefaultUncachedAttributeCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DefaultUncachedAttributeCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DefaultUncachedAttributeCriteriaResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *DefaultUncachedAttributeCriteriaResponse) GetSchemas() []EnumdefaultUncachedAttributeCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumdefaultUncachedAttributeCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *DefaultUncachedAttributeCriteriaResponse) GetSchemasOk() ([]EnumdefaultUncachedAttributeCriteriaSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *DefaultUncachedAttributeCriteriaResponse) SetSchemas(v []EnumdefaultUncachedAttributeCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DefaultUncachedAttributeCriteriaResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultUncachedAttributeCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DefaultUncachedAttributeCriteriaResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DefaultUncachedAttributeCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *DefaultUncachedAttributeCriteriaResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DefaultUncachedAttributeCriteriaResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DefaultUncachedAttributeCriteriaResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DefaultUncachedAttributeCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultUncachedAttributeCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DefaultUncachedAttributeCriteriaResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DefaultUncachedAttributeCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o DefaultUncachedAttributeCriteriaResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultUncachedAttributeCriteriaResponse struct {
	value *DefaultUncachedAttributeCriteriaResponse
	isSet bool
}

func (v NullableDefaultUncachedAttributeCriteriaResponse) Get() *DefaultUncachedAttributeCriteriaResponse {
	return v.value
}

func (v *NullableDefaultUncachedAttributeCriteriaResponse) Set(val *DefaultUncachedAttributeCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultUncachedAttributeCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultUncachedAttributeCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultUncachedAttributeCriteriaResponse(val *DefaultUncachedAttributeCriteriaResponse) *NullableDefaultUncachedAttributeCriteriaResponse {
	return &NullableDefaultUncachedAttributeCriteriaResponse{value: val, isSet: true}
}

func (v NullableDefaultUncachedAttributeCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultUncachedAttributeCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


