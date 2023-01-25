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

// SubstringMatchingRuleResponse struct for SubstringMatchingRuleResponse
type SubstringMatchingRuleResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumsubstringMatchingRuleSchemaUrn `json:"schemas"`
	// Name of the Matching Rule
	Id string `json:"id"`
	// Indicates whether the Matching Rule is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewSubstringMatchingRuleResponse instantiates a new SubstringMatchingRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubstringMatchingRuleResponse(schemas []EnumsubstringMatchingRuleSchemaUrn, id string, enabled bool) *SubstringMatchingRuleResponse {
	this := SubstringMatchingRuleResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewSubstringMatchingRuleResponseWithDefaults instantiates a new SubstringMatchingRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubstringMatchingRuleResponseWithDefaults() *SubstringMatchingRuleResponse {
	this := SubstringMatchingRuleResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubstringMatchingRuleResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstringMatchingRuleResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubstringMatchingRuleResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SubstringMatchingRuleResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SubstringMatchingRuleResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstringMatchingRuleResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SubstringMatchingRuleResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SubstringMatchingRuleResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *SubstringMatchingRuleResponse) GetSchemas() []EnumsubstringMatchingRuleSchemaUrn {
	if o == nil {
		var ret []EnumsubstringMatchingRuleSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SubstringMatchingRuleResponse) GetSchemasOk() ([]EnumsubstringMatchingRuleSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SubstringMatchingRuleResponse) SetSchemas(v []EnumsubstringMatchingRuleSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *SubstringMatchingRuleResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubstringMatchingRuleResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubstringMatchingRuleResponse) SetId(v string) {
	o.Id = v
}

// GetEnabled returns the Enabled field value
func (o *SubstringMatchingRuleResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SubstringMatchingRuleResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SubstringMatchingRuleResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o SubstringMatchingRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSubstringMatchingRuleResponse struct {
	value *SubstringMatchingRuleResponse
	isSet bool
}

func (v NullableSubstringMatchingRuleResponse) Get() *SubstringMatchingRuleResponse {
	return v.value
}

func (v *NullableSubstringMatchingRuleResponse) Set(val *SubstringMatchingRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubstringMatchingRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubstringMatchingRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubstringMatchingRuleResponse(val *SubstringMatchingRuleResponse) *NullableSubstringMatchingRuleResponse {
	return &NullableSubstringMatchingRuleResponse{value: val, isSet: true}
}

func (v NullableSubstringMatchingRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubstringMatchingRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


