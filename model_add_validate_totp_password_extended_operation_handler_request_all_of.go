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

// AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf struct for AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf
type AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf struct {
	// Name of the new Extended Operation Handler
	HandlerName string `json:"handlerName"`
}

// NewAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf instantiates a new AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf(handlerName string) *AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf {
	this := AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf{}
	this.HandlerName = handlerName
	return &this
}

// NewAddValidateTotpPasswordExtendedOperationHandlerRequestAllOfWithDefaults instantiates a new AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddValidateTotpPasswordExtendedOperationHandlerRequestAllOfWithDefaults() *AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf {
	this := AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) GetHandlerNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) SetHandlerName(v string) {
	o.HandlerName = v
}

func (o AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handlerName"] = o.HandlerName
	}
	return json.Marshal(toSerialize)
}

type NullableAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf struct {
	value *AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf
	isSet bool
}

func (v NullableAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) Get() *AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf {
	return v.value
}

func (v *NullableAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) Set(val *AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf(val *AddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) *NullableAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf {
	return &NullableAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf{value: val, isSet: true}
}

func (v NullableAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddValidateTotpPasswordExtendedOperationHandlerRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


