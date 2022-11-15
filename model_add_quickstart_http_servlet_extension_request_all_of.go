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

// AddQuickstartHttpServletExtensionRequestAllOf struct for AddQuickstartHttpServletExtensionRequestAllOf
type AddQuickstartHttpServletExtensionRequestAllOf struct {
	// Name of the new HTTP Servlet Extension
	ExtensionName string `json:"extensionName"`
}

// NewAddQuickstartHttpServletExtensionRequestAllOf instantiates a new AddQuickstartHttpServletExtensionRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddQuickstartHttpServletExtensionRequestAllOf(extensionName string) *AddQuickstartHttpServletExtensionRequestAllOf {
	this := AddQuickstartHttpServletExtensionRequestAllOf{}
	this.ExtensionName = extensionName
	return &this
}

// NewAddQuickstartHttpServletExtensionRequestAllOfWithDefaults instantiates a new AddQuickstartHttpServletExtensionRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddQuickstartHttpServletExtensionRequestAllOfWithDefaults() *AddQuickstartHttpServletExtensionRequestAllOf {
	this := AddQuickstartHttpServletExtensionRequestAllOf{}
	return &this
}

// GetExtensionName returns the ExtensionName field value
func (o *AddQuickstartHttpServletExtensionRequestAllOf) GetExtensionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionName
}

// GetExtensionNameOk returns a tuple with the ExtensionName field value
// and a boolean to check if the value has been set.
func (o *AddQuickstartHttpServletExtensionRequestAllOf) GetExtensionNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionName, true
}

// SetExtensionName sets field value
func (o *AddQuickstartHttpServletExtensionRequestAllOf) SetExtensionName(v string) {
	o.ExtensionName = v
}

func (o AddQuickstartHttpServletExtensionRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["extensionName"] = o.ExtensionName
	}
	return json.Marshal(toSerialize)
}

type NullableAddQuickstartHttpServletExtensionRequestAllOf struct {
	value *AddQuickstartHttpServletExtensionRequestAllOf
	isSet bool
}

func (v NullableAddQuickstartHttpServletExtensionRequestAllOf) Get() *AddQuickstartHttpServletExtensionRequestAllOf {
	return v.value
}

func (v *NullableAddQuickstartHttpServletExtensionRequestAllOf) Set(val *AddQuickstartHttpServletExtensionRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddQuickstartHttpServletExtensionRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddQuickstartHttpServletExtensionRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddQuickstartHttpServletExtensionRequestAllOf(val *AddQuickstartHttpServletExtensionRequestAllOf) *NullableAddQuickstartHttpServletExtensionRequestAllOf {
	return &NullableAddQuickstartHttpServletExtensionRequestAllOf{value: val, isSet: true}
}

func (v NullableAddQuickstartHttpServletExtensionRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddQuickstartHttpServletExtensionRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


