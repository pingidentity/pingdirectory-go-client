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

// AggregateRequestCriteriaShared struct for AggregateRequestCriteriaShared
type AggregateRequestCriteriaShared struct {
	Schemas []EnumaggregateRequestCriteriaSchemaUrn `json:"schemas"`
	AllIncludedRequestCriteria []string `json:"allIncludedRequestCriteria,omitempty"`
	AnyIncludedRequestCriteria []string `json:"anyIncludedRequestCriteria,omitempty"`
	NotAllIncludedRequestCriteria []string `json:"notAllIncludedRequestCriteria,omitempty"`
	NoneIncludedRequestCriteria []string `json:"noneIncludedRequestCriteria,omitempty"`
	// A description for this Request Criteria
	Description *string `json:"description,omitempty"`
}

// NewAggregateRequestCriteriaShared instantiates a new AggregateRequestCriteriaShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregateRequestCriteriaShared(schemas []EnumaggregateRequestCriteriaSchemaUrn) *AggregateRequestCriteriaShared {
	this := AggregateRequestCriteriaShared{}
	this.Schemas = schemas
	return &this
}

// NewAggregateRequestCriteriaSharedWithDefaults instantiates a new AggregateRequestCriteriaShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregateRequestCriteriaSharedWithDefaults() *AggregateRequestCriteriaShared {
	this := AggregateRequestCriteriaShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AggregateRequestCriteriaShared) GetSchemas() []EnumaggregateRequestCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumaggregateRequestCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AggregateRequestCriteriaShared) GetSchemasOk() ([]EnumaggregateRequestCriteriaSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AggregateRequestCriteriaShared) SetSchemas(v []EnumaggregateRequestCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetAllIncludedRequestCriteria returns the AllIncludedRequestCriteria field value if set, zero value otherwise.
func (o *AggregateRequestCriteriaShared) GetAllIncludedRequestCriteria() []string {
	if o == nil || isNil(o.AllIncludedRequestCriteria) {
		var ret []string
		return ret
	}
	return o.AllIncludedRequestCriteria
}

// GetAllIncludedRequestCriteriaOk returns a tuple with the AllIncludedRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateRequestCriteriaShared) GetAllIncludedRequestCriteriaOk() ([]string, bool) {
	if o == nil || isNil(o.AllIncludedRequestCriteria) {
    return nil, false
	}
	return o.AllIncludedRequestCriteria, true
}

// HasAllIncludedRequestCriteria returns a boolean if a field has been set.
func (o *AggregateRequestCriteriaShared) HasAllIncludedRequestCriteria() bool {
	if o != nil && !isNil(o.AllIncludedRequestCriteria) {
		return true
	}

	return false
}

// SetAllIncludedRequestCriteria gets a reference to the given []string and assigns it to the AllIncludedRequestCriteria field.
func (o *AggregateRequestCriteriaShared) SetAllIncludedRequestCriteria(v []string) {
	o.AllIncludedRequestCriteria = v
}

// GetAnyIncludedRequestCriteria returns the AnyIncludedRequestCriteria field value if set, zero value otherwise.
func (o *AggregateRequestCriteriaShared) GetAnyIncludedRequestCriteria() []string {
	if o == nil || isNil(o.AnyIncludedRequestCriteria) {
		var ret []string
		return ret
	}
	return o.AnyIncludedRequestCriteria
}

// GetAnyIncludedRequestCriteriaOk returns a tuple with the AnyIncludedRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateRequestCriteriaShared) GetAnyIncludedRequestCriteriaOk() ([]string, bool) {
	if o == nil || isNil(o.AnyIncludedRequestCriteria) {
    return nil, false
	}
	return o.AnyIncludedRequestCriteria, true
}

// HasAnyIncludedRequestCriteria returns a boolean if a field has been set.
func (o *AggregateRequestCriteriaShared) HasAnyIncludedRequestCriteria() bool {
	if o != nil && !isNil(o.AnyIncludedRequestCriteria) {
		return true
	}

	return false
}

// SetAnyIncludedRequestCriteria gets a reference to the given []string and assigns it to the AnyIncludedRequestCriteria field.
func (o *AggregateRequestCriteriaShared) SetAnyIncludedRequestCriteria(v []string) {
	o.AnyIncludedRequestCriteria = v
}

// GetNotAllIncludedRequestCriteria returns the NotAllIncludedRequestCriteria field value if set, zero value otherwise.
func (o *AggregateRequestCriteriaShared) GetNotAllIncludedRequestCriteria() []string {
	if o == nil || isNil(o.NotAllIncludedRequestCriteria) {
		var ret []string
		return ret
	}
	return o.NotAllIncludedRequestCriteria
}

// GetNotAllIncludedRequestCriteriaOk returns a tuple with the NotAllIncludedRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateRequestCriteriaShared) GetNotAllIncludedRequestCriteriaOk() ([]string, bool) {
	if o == nil || isNil(o.NotAllIncludedRequestCriteria) {
    return nil, false
	}
	return o.NotAllIncludedRequestCriteria, true
}

// HasNotAllIncludedRequestCriteria returns a boolean if a field has been set.
func (o *AggregateRequestCriteriaShared) HasNotAllIncludedRequestCriteria() bool {
	if o != nil && !isNil(o.NotAllIncludedRequestCriteria) {
		return true
	}

	return false
}

// SetNotAllIncludedRequestCriteria gets a reference to the given []string and assigns it to the NotAllIncludedRequestCriteria field.
func (o *AggregateRequestCriteriaShared) SetNotAllIncludedRequestCriteria(v []string) {
	o.NotAllIncludedRequestCriteria = v
}

// GetNoneIncludedRequestCriteria returns the NoneIncludedRequestCriteria field value if set, zero value otherwise.
func (o *AggregateRequestCriteriaShared) GetNoneIncludedRequestCriteria() []string {
	if o == nil || isNil(o.NoneIncludedRequestCriteria) {
		var ret []string
		return ret
	}
	return o.NoneIncludedRequestCriteria
}

// GetNoneIncludedRequestCriteriaOk returns a tuple with the NoneIncludedRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateRequestCriteriaShared) GetNoneIncludedRequestCriteriaOk() ([]string, bool) {
	if o == nil || isNil(o.NoneIncludedRequestCriteria) {
    return nil, false
	}
	return o.NoneIncludedRequestCriteria, true
}

// HasNoneIncludedRequestCriteria returns a boolean if a field has been set.
func (o *AggregateRequestCriteriaShared) HasNoneIncludedRequestCriteria() bool {
	if o != nil && !isNil(o.NoneIncludedRequestCriteria) {
		return true
	}

	return false
}

// SetNoneIncludedRequestCriteria gets a reference to the given []string and assigns it to the NoneIncludedRequestCriteria field.
func (o *AggregateRequestCriteriaShared) SetNoneIncludedRequestCriteria(v []string) {
	o.NoneIncludedRequestCriteria = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AggregateRequestCriteriaShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateRequestCriteriaShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AggregateRequestCriteriaShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AggregateRequestCriteriaShared) SetDescription(v string) {
	o.Description = &v
}

func (o AggregateRequestCriteriaShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.AllIncludedRequestCriteria) {
		toSerialize["allIncludedRequestCriteria"] = o.AllIncludedRequestCriteria
	}
	if !isNil(o.AnyIncludedRequestCriteria) {
		toSerialize["anyIncludedRequestCriteria"] = o.AnyIncludedRequestCriteria
	}
	if !isNil(o.NotAllIncludedRequestCriteria) {
		toSerialize["notAllIncludedRequestCriteria"] = o.NotAllIncludedRequestCriteria
	}
	if !isNil(o.NoneIncludedRequestCriteria) {
		toSerialize["noneIncludedRequestCriteria"] = o.NoneIncludedRequestCriteria
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAggregateRequestCriteriaShared struct {
	value *AggregateRequestCriteriaShared
	isSet bool
}

func (v NullableAggregateRequestCriteriaShared) Get() *AggregateRequestCriteriaShared {
	return v.value
}

func (v *NullableAggregateRequestCriteriaShared) Set(val *AggregateRequestCriteriaShared) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregateRequestCriteriaShared) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregateRequestCriteriaShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregateRequestCriteriaShared(val *AggregateRequestCriteriaShared) *NullableAggregateRequestCriteriaShared {
	return &NullableAggregateRequestCriteriaShared{value: val, isSet: true}
}

func (v NullableAggregateRequestCriteriaShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregateRequestCriteriaShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


