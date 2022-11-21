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

// DefaultUncachedEntryCriteriaResponse struct for DefaultUncachedEntryCriteriaResponse
type DefaultUncachedEntryCriteriaResponse struct {
	// Name of the Uncached Entry Criteria
	Id string `json:"id"`
	Schemas []EnumdefaultUncachedEntryCriteriaSchemaUrn `json:"schemas"`
	// A description for this Uncached Entry Criteria
	Description *string `json:"description,omitempty"`
	// Indicates whether this Uncached Entry Criteria is enabled for use in the server.
	Enabled bool `json:"enabled"`
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewDefaultUncachedEntryCriteriaResponse instantiates a new DefaultUncachedEntryCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultUncachedEntryCriteriaResponse(id string, schemas []EnumdefaultUncachedEntryCriteriaSchemaUrn, enabled bool) *DefaultUncachedEntryCriteriaResponse {
	this := DefaultUncachedEntryCriteriaResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewDefaultUncachedEntryCriteriaResponseWithDefaults instantiates a new DefaultUncachedEntryCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultUncachedEntryCriteriaResponseWithDefaults() *DefaultUncachedEntryCriteriaResponse {
	this := DefaultUncachedEntryCriteriaResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DefaultUncachedEntryCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DefaultUncachedEntryCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DefaultUncachedEntryCriteriaResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *DefaultUncachedEntryCriteriaResponse) GetSchemas() []EnumdefaultUncachedEntryCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumdefaultUncachedEntryCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *DefaultUncachedEntryCriteriaResponse) GetSchemasOk() ([]EnumdefaultUncachedEntryCriteriaSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *DefaultUncachedEntryCriteriaResponse) SetSchemas(v []EnumdefaultUncachedEntryCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DefaultUncachedEntryCriteriaResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultUncachedEntryCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DefaultUncachedEntryCriteriaResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DefaultUncachedEntryCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *DefaultUncachedEntryCriteriaResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DefaultUncachedEntryCriteriaResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DefaultUncachedEntryCriteriaResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DefaultUncachedEntryCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultUncachedEntryCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DefaultUncachedEntryCriteriaResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DefaultUncachedEntryCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o DefaultUncachedEntryCriteriaResponse) MarshalJSON() ([]byte, error) {
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

type NullableDefaultUncachedEntryCriteriaResponse struct {
	value *DefaultUncachedEntryCriteriaResponse
	isSet bool
}

func (v NullableDefaultUncachedEntryCriteriaResponse) Get() *DefaultUncachedEntryCriteriaResponse {
	return v.value
}

func (v *NullableDefaultUncachedEntryCriteriaResponse) Set(val *DefaultUncachedEntryCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultUncachedEntryCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultUncachedEntryCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultUncachedEntryCriteriaResponse(val *DefaultUncachedEntryCriteriaResponse) *NullableDefaultUncachedEntryCriteriaResponse {
	return &NullableDefaultUncachedEntryCriteriaResponse{value: val, isSet: true}
}

func (v NullableDefaultUncachedEntryCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultUncachedEntryCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


