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

// SimpleUncachedAttributeCriteriaResponse struct for SimpleUncachedAttributeCriteriaResponse
type SimpleUncachedAttributeCriteriaResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Uncached Attribute Criteria
	Id string `json:"id"`
	Schemas []EnumsimpleUncachedAttributeCriteriaSchemaUrn `json:"schemas"`
	// Specifies the attribute types for attributes that may be written to the uncached-id2entry database.
	AttributeType []string `json:"attributeType"`
	// Specifies the minimum number of values that an attribute must have before it will be written into the uncached-id2entry database.
	MinValueCount *int32 `json:"minValueCount,omitempty"`
	// Specifies the minimum total value size (i.e., the sum of the sizes of all values) that an attribute must have before it will be written into the uncached-id2entry database.
	MinTotalValueSize *string `json:"minTotalValueSize,omitempty"`
	// A description for this Uncached Attribute Criteria
	Description *string `json:"description,omitempty"`
	// Indicates whether this Uncached Attribute Criteria is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewSimpleUncachedAttributeCriteriaResponse instantiates a new SimpleUncachedAttributeCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleUncachedAttributeCriteriaResponse(id string, schemas []EnumsimpleUncachedAttributeCriteriaSchemaUrn, attributeType []string, enabled bool) *SimpleUncachedAttributeCriteriaResponse {
	this := SimpleUncachedAttributeCriteriaResponse{}
	this.Id = id
	this.Schemas = schemas
	this.AttributeType = attributeType
	this.Enabled = enabled
	return &this
}

// NewSimpleUncachedAttributeCriteriaResponseWithDefaults instantiates a new SimpleUncachedAttributeCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleUncachedAttributeCriteriaResponseWithDefaults() *SimpleUncachedAttributeCriteriaResponse {
	this := SimpleUncachedAttributeCriteriaResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SimpleUncachedAttributeCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SimpleUncachedAttributeCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SimpleUncachedAttributeCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SimpleUncachedAttributeCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *SimpleUncachedAttributeCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimpleUncachedAttributeCriteriaResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *SimpleUncachedAttributeCriteriaResponse) GetSchemas() []EnumsimpleUncachedAttributeCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumsimpleUncachedAttributeCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) GetSchemasOk() ([]EnumsimpleUncachedAttributeCriteriaSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SimpleUncachedAttributeCriteriaResponse) SetSchemas(v []EnumsimpleUncachedAttributeCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetAttributeType returns the AttributeType field value
func (o *SimpleUncachedAttributeCriteriaResponse) GetAttributeType() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) GetAttributeTypeOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AttributeType, true
}

// SetAttributeType sets field value
func (o *SimpleUncachedAttributeCriteriaResponse) SetAttributeType(v []string) {
	o.AttributeType = v
}

// GetMinValueCount returns the MinValueCount field value if set, zero value otherwise.
func (o *SimpleUncachedAttributeCriteriaResponse) GetMinValueCount() int32 {
	if o == nil || isNil(o.MinValueCount) {
		var ret int32
		return ret
	}
	return *o.MinValueCount
}

// GetMinValueCountOk returns a tuple with the MinValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) GetMinValueCountOk() (*int32, bool) {
	if o == nil || isNil(o.MinValueCount) {
    return nil, false
	}
	return o.MinValueCount, true
}

// HasMinValueCount returns a boolean if a field has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) HasMinValueCount() bool {
	if o != nil && !isNil(o.MinValueCount) {
		return true
	}

	return false
}

// SetMinValueCount gets a reference to the given int32 and assigns it to the MinValueCount field.
func (o *SimpleUncachedAttributeCriteriaResponse) SetMinValueCount(v int32) {
	o.MinValueCount = &v
}

// GetMinTotalValueSize returns the MinTotalValueSize field value if set, zero value otherwise.
func (o *SimpleUncachedAttributeCriteriaResponse) GetMinTotalValueSize() string {
	if o == nil || isNil(o.MinTotalValueSize) {
		var ret string
		return ret
	}
	return *o.MinTotalValueSize
}

// GetMinTotalValueSizeOk returns a tuple with the MinTotalValueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) GetMinTotalValueSizeOk() (*string, bool) {
	if o == nil || isNil(o.MinTotalValueSize) {
    return nil, false
	}
	return o.MinTotalValueSize, true
}

// HasMinTotalValueSize returns a boolean if a field has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) HasMinTotalValueSize() bool {
	if o != nil && !isNil(o.MinTotalValueSize) {
		return true
	}

	return false
}

// SetMinTotalValueSize gets a reference to the given string and assigns it to the MinTotalValueSize field.
func (o *SimpleUncachedAttributeCriteriaResponse) SetMinTotalValueSize(v string) {
	o.MinTotalValueSize = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SimpleUncachedAttributeCriteriaResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SimpleUncachedAttributeCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SimpleUncachedAttributeCriteriaResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SimpleUncachedAttributeCriteriaResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SimpleUncachedAttributeCriteriaResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o SimpleUncachedAttributeCriteriaResponse) MarshalJSON() ([]byte, error) {
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
		toSerialize["attributeType"] = o.AttributeType
	}
	if !isNil(o.MinValueCount) {
		toSerialize["minValueCount"] = o.MinValueCount
	}
	if !isNil(o.MinTotalValueSize) {
		toSerialize["minTotalValueSize"] = o.MinTotalValueSize
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleUncachedAttributeCriteriaResponse struct {
	value *SimpleUncachedAttributeCriteriaResponse
	isSet bool
}

func (v NullableSimpleUncachedAttributeCriteriaResponse) Get() *SimpleUncachedAttributeCriteriaResponse {
	return v.value
}

func (v *NullableSimpleUncachedAttributeCriteriaResponse) Set(val *SimpleUncachedAttributeCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleUncachedAttributeCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleUncachedAttributeCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleUncachedAttributeCriteriaResponse(val *SimpleUncachedAttributeCriteriaResponse) *NullableSimpleUncachedAttributeCriteriaResponse {
	return &NullableSimpleUncachedAttributeCriteriaResponse{value: val, isSet: true}
}

func (v NullableSimpleUncachedAttributeCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleUncachedAttributeCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

