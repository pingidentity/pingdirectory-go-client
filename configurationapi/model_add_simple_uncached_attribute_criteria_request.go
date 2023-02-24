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

// checks if the AddSimpleUncachedAttributeCriteriaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSimpleUncachedAttributeCriteriaRequest{}

// AddSimpleUncachedAttributeCriteriaRequest struct for AddSimpleUncachedAttributeCriteriaRequest
type AddSimpleUncachedAttributeCriteriaRequest struct {
	// Name of the new Uncached Attribute Criteria
	CriteriaName string                                         `json:"criteriaName"`
	Schemas      []EnumsimpleUncachedAttributeCriteriaSchemaUrn `json:"schemas"`
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

// NewAddSimpleUncachedAttributeCriteriaRequest instantiates a new AddSimpleUncachedAttributeCriteriaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSimpleUncachedAttributeCriteriaRequest(criteriaName string, schemas []EnumsimpleUncachedAttributeCriteriaSchemaUrn, attributeType []string, enabled bool) *AddSimpleUncachedAttributeCriteriaRequest {
	this := AddSimpleUncachedAttributeCriteriaRequest{}
	this.CriteriaName = criteriaName
	this.Schemas = schemas
	this.AttributeType = attributeType
	this.Enabled = enabled
	return &this
}

// NewAddSimpleUncachedAttributeCriteriaRequestWithDefaults instantiates a new AddSimpleUncachedAttributeCriteriaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSimpleUncachedAttributeCriteriaRequestWithDefaults() *AddSimpleUncachedAttributeCriteriaRequest {
	this := AddSimpleUncachedAttributeCriteriaRequest{}
	return &this
}

// GetCriteriaName returns the CriteriaName field value
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetCriteriaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CriteriaName
}

// GetCriteriaNameOk returns a tuple with the CriteriaName field value
// and a boolean to check if the value has been set.
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetCriteriaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CriteriaName, true
}

// SetCriteriaName sets field value
func (o *AddSimpleUncachedAttributeCriteriaRequest) SetCriteriaName(v string) {
	o.CriteriaName = v
}

// GetSchemas returns the Schemas field value
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetSchemas() []EnumsimpleUncachedAttributeCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumsimpleUncachedAttributeCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetSchemasOk() ([]EnumsimpleUncachedAttributeCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSimpleUncachedAttributeCriteriaRequest) SetSchemas(v []EnumsimpleUncachedAttributeCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetAttributeType returns the AttributeType field value
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetAttributeType() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetAttributeTypeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeType, true
}

// SetAttributeType sets field value
func (o *AddSimpleUncachedAttributeCriteriaRequest) SetAttributeType(v []string) {
	o.AttributeType = v
}

// GetMinValueCount returns the MinValueCount field value if set, zero value otherwise.
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetMinValueCount() int32 {
	if o == nil || IsNil(o.MinValueCount) {
		var ret int32
		return ret
	}
	return *o.MinValueCount
}

// GetMinValueCountOk returns a tuple with the MinValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetMinValueCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MinValueCount) {
		return nil, false
	}
	return o.MinValueCount, true
}

// HasMinValueCount returns a boolean if a field has been set.
func (o *AddSimpleUncachedAttributeCriteriaRequest) HasMinValueCount() bool {
	if o != nil && !IsNil(o.MinValueCount) {
		return true
	}

	return false
}

// SetMinValueCount gets a reference to the given int32 and assigns it to the MinValueCount field.
func (o *AddSimpleUncachedAttributeCriteriaRequest) SetMinValueCount(v int32) {
	o.MinValueCount = &v
}

// GetMinTotalValueSize returns the MinTotalValueSize field value if set, zero value otherwise.
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetMinTotalValueSize() string {
	if o == nil || IsNil(o.MinTotalValueSize) {
		var ret string
		return ret
	}
	return *o.MinTotalValueSize
}

// GetMinTotalValueSizeOk returns a tuple with the MinTotalValueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetMinTotalValueSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MinTotalValueSize) {
		return nil, false
	}
	return o.MinTotalValueSize, true
}

// HasMinTotalValueSize returns a boolean if a field has been set.
func (o *AddSimpleUncachedAttributeCriteriaRequest) HasMinTotalValueSize() bool {
	if o != nil && !IsNil(o.MinTotalValueSize) {
		return true
	}

	return false
}

// SetMinTotalValueSize gets a reference to the given string and assigns it to the MinTotalValueSize field.
func (o *AddSimpleUncachedAttributeCriteriaRequest) SetMinTotalValueSize(v string) {
	o.MinTotalValueSize = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSimpleUncachedAttributeCriteriaRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSimpleUncachedAttributeCriteriaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddSimpleUncachedAttributeCriteriaRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddSimpleUncachedAttributeCriteriaRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddSimpleUncachedAttributeCriteriaRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSimpleUncachedAttributeCriteriaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["criteriaName"] = o.CriteriaName
	toSerialize["schemas"] = o.Schemas
	toSerialize["attributeType"] = o.AttributeType
	if !IsNil(o.MinValueCount) {
		toSerialize["minValueCount"] = o.MinValueCount
	}
	if !IsNil(o.MinTotalValueSize) {
		toSerialize["minTotalValueSize"] = o.MinTotalValueSize
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAddSimpleUncachedAttributeCriteriaRequest struct {
	value *AddSimpleUncachedAttributeCriteriaRequest
	isSet bool
}

func (v NullableAddSimpleUncachedAttributeCriteriaRequest) Get() *AddSimpleUncachedAttributeCriteriaRequest {
	return v.value
}

func (v *NullableAddSimpleUncachedAttributeCriteriaRequest) Set(val *AddSimpleUncachedAttributeCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSimpleUncachedAttributeCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSimpleUncachedAttributeCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSimpleUncachedAttributeCriteriaRequest(val *AddSimpleUncachedAttributeCriteriaRequest) *NullableAddSimpleUncachedAttributeCriteriaRequest {
	return &NullableAddSimpleUncachedAttributeCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddSimpleUncachedAttributeCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSimpleUncachedAttributeCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
