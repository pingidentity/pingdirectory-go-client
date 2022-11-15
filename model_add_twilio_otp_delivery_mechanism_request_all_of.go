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

// AddTwilioOtpDeliveryMechanismRequestAllOf struct for AddTwilioOtpDeliveryMechanismRequestAllOf
type AddTwilioOtpDeliveryMechanismRequestAllOf struct {
	// Name of the new OTP Delivery Mechanism
	MechanismName string `json:"mechanismName"`
}

// NewAddTwilioOtpDeliveryMechanismRequestAllOf instantiates a new AddTwilioOtpDeliveryMechanismRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTwilioOtpDeliveryMechanismRequestAllOf(mechanismName string) *AddTwilioOtpDeliveryMechanismRequestAllOf {
	this := AddTwilioOtpDeliveryMechanismRequestAllOf{}
	this.MechanismName = mechanismName
	return &this
}

// NewAddTwilioOtpDeliveryMechanismRequestAllOfWithDefaults instantiates a new AddTwilioOtpDeliveryMechanismRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTwilioOtpDeliveryMechanismRequestAllOfWithDefaults() *AddTwilioOtpDeliveryMechanismRequestAllOf {
	this := AddTwilioOtpDeliveryMechanismRequestAllOf{}
	return &this
}

// GetMechanismName returns the MechanismName field value
func (o *AddTwilioOtpDeliveryMechanismRequestAllOf) GetMechanismName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MechanismName
}

// GetMechanismNameOk returns a tuple with the MechanismName field value
// and a boolean to check if the value has been set.
func (o *AddTwilioOtpDeliveryMechanismRequestAllOf) GetMechanismNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MechanismName, true
}

// SetMechanismName sets field value
func (o *AddTwilioOtpDeliveryMechanismRequestAllOf) SetMechanismName(v string) {
	o.MechanismName = v
}

func (o AddTwilioOtpDeliveryMechanismRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mechanismName"] = o.MechanismName
	}
	return json.Marshal(toSerialize)
}

type NullableAddTwilioOtpDeliveryMechanismRequestAllOf struct {
	value *AddTwilioOtpDeliveryMechanismRequestAllOf
	isSet bool
}

func (v NullableAddTwilioOtpDeliveryMechanismRequestAllOf) Get() *AddTwilioOtpDeliveryMechanismRequestAllOf {
	return v.value
}

func (v *NullableAddTwilioOtpDeliveryMechanismRequestAllOf) Set(val *AddTwilioOtpDeliveryMechanismRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTwilioOtpDeliveryMechanismRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTwilioOtpDeliveryMechanismRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTwilioOtpDeliveryMechanismRequestAllOf(val *AddTwilioOtpDeliveryMechanismRequestAllOf) *NullableAddTwilioOtpDeliveryMechanismRequestAllOf {
	return &NullableAddTwilioOtpDeliveryMechanismRequestAllOf{value: val, isSet: true}
}

func (v NullableAddTwilioOtpDeliveryMechanismRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTwilioOtpDeliveryMechanismRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


