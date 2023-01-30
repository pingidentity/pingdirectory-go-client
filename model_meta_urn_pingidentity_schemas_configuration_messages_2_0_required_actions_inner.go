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

// MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner struct for MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner
type MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner struct {
	Property *string `json:"property,omitempty"`
	Type     string  `json:"type"`
	Synopsis string  `json:"synopsis"`
}

// NewMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner instantiates a new MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner(type_ string, synopsis string) *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner {
	this := MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner{}
	this.Type = type_
	this.Synopsis = synopsis
	return &this
}

// NewMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInnerWithDefaults instantiates a new MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInnerWithDefaults() *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner {
	this := MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner{}
	return &this
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) GetProperty() string {
	if o == nil || isNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) GetPropertyOk() (*string, bool) {
	if o == nil || isNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) HasProperty() bool {
	if o != nil && !isNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) SetProperty(v string) {
	o.Property = &v
}

// GetType returns the Type field value
func (o *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) SetType(v string) {
	o.Type = v
}

// GetSynopsis returns the Synopsis field value
func (o *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) GetSynopsis() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Synopsis
}

// GetSynopsisOk returns a tuple with the Synopsis field value
// and a boolean to check if the value has been set.
func (o *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) GetSynopsisOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Synopsis, true
}

// SetSynopsis sets field value
func (o *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) SetSynopsis(v string) {
	o.Synopsis = v
}

func (o MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["synopsis"] = o.Synopsis
	}
	return json.Marshal(toSerialize)
}

type NullableMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner struct {
	value *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner
	isSet bool
}

func (v NullableMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) Get() *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner {
	return v.value
}

func (v *NullableMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) Set(val *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner(val *MetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) *NullableMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner {
	return &NullableMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner{value: val, isSet: true}
}

func (v NullableMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaUrnPingidentitySchemasConfigurationMessages20RequiredActionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
