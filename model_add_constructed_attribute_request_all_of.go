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

// AddConstructedAttributeRequestAllOf struct for AddConstructedAttributeRequestAllOf
type AddConstructedAttributeRequestAllOf struct {
	// Name of the new Constructed Attribute
	AttributeName string `json:"attributeName"`
}

// NewAddConstructedAttributeRequestAllOf instantiates a new AddConstructedAttributeRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddConstructedAttributeRequestAllOf(attributeName string) *AddConstructedAttributeRequestAllOf {
	this := AddConstructedAttributeRequestAllOf{}
	this.AttributeName = attributeName
	return &this
}

// NewAddConstructedAttributeRequestAllOfWithDefaults instantiates a new AddConstructedAttributeRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddConstructedAttributeRequestAllOfWithDefaults() *AddConstructedAttributeRequestAllOf {
	this := AddConstructedAttributeRequestAllOf{}
	return &this
}

// GetAttributeName returns the AttributeName field value
func (o *AddConstructedAttributeRequestAllOf) GetAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value
// and a boolean to check if the value has been set.
func (o *AddConstructedAttributeRequestAllOf) GetAttributeNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AttributeName, true
}

// SetAttributeName sets field value
func (o *AddConstructedAttributeRequestAllOf) SetAttributeName(v string) {
	o.AttributeName = v
}

func (o AddConstructedAttributeRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributeName"] = o.AttributeName
	}
	return json.Marshal(toSerialize)
}

type NullableAddConstructedAttributeRequestAllOf struct {
	value *AddConstructedAttributeRequestAllOf
	isSet bool
}

func (v NullableAddConstructedAttributeRequestAllOf) Get() *AddConstructedAttributeRequestAllOf {
	return v.value
}

func (v *NullableAddConstructedAttributeRequestAllOf) Set(val *AddConstructedAttributeRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddConstructedAttributeRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddConstructedAttributeRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddConstructedAttributeRequestAllOf(val *AddConstructedAttributeRequestAllOf) *NullableAddConstructedAttributeRequestAllOf {
	return &NullableAddConstructedAttributeRequestAllOf{value: val, isSet: true}
}

func (v NullableAddConstructedAttributeRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddConstructedAttributeRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


