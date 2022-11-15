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

// AddScimSchemaRequestAllOf struct for AddScimSchemaRequestAllOf
type AddScimSchemaRequestAllOf struct {
	// Name of the new SCIM Schema
	SchemaName string `json:"schemaName"`
}

// NewAddScimSchemaRequestAllOf instantiates a new AddScimSchemaRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddScimSchemaRequestAllOf(schemaName string) *AddScimSchemaRequestAllOf {
	this := AddScimSchemaRequestAllOf{}
	this.SchemaName = schemaName
	return &this
}

// NewAddScimSchemaRequestAllOfWithDefaults instantiates a new AddScimSchemaRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddScimSchemaRequestAllOfWithDefaults() *AddScimSchemaRequestAllOf {
	this := AddScimSchemaRequestAllOf{}
	return &this
}

// GetSchemaName returns the SchemaName field value
func (o *AddScimSchemaRequestAllOf) GetSchemaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value
// and a boolean to check if the value has been set.
func (o *AddScimSchemaRequestAllOf) GetSchemaNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SchemaName, true
}

// SetSchemaName sets field value
func (o *AddScimSchemaRequestAllOf) SetSchemaName(v string) {
	o.SchemaName = v
}

func (o AddScimSchemaRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemaName"] = o.SchemaName
	}
	return json.Marshal(toSerialize)
}

type NullableAddScimSchemaRequestAllOf struct {
	value *AddScimSchemaRequestAllOf
	isSet bool
}

func (v NullableAddScimSchemaRequestAllOf) Get() *AddScimSchemaRequestAllOf {
	return v.value
}

func (v *NullableAddScimSchemaRequestAllOf) Set(val *AddScimSchemaRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddScimSchemaRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddScimSchemaRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddScimSchemaRequestAllOf(val *AddScimSchemaRequestAllOf) *NullableAddScimSchemaRequestAllOf {
	return &NullableAddScimSchemaRequestAllOf{value: val, isSet: true}
}

func (v NullableAddScimSchemaRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddScimSchemaRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


