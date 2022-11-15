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

// AddTimeLimitLogRetentionPolicyRequestAllOf struct for AddTimeLimitLogRetentionPolicyRequestAllOf
type AddTimeLimitLogRetentionPolicyRequestAllOf struct {
	// Name of the new Log Retention Policy
	PolicyName string `json:"policyName"`
}

// NewAddTimeLimitLogRetentionPolicyRequestAllOf instantiates a new AddTimeLimitLogRetentionPolicyRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTimeLimitLogRetentionPolicyRequestAllOf(policyName string) *AddTimeLimitLogRetentionPolicyRequestAllOf {
	this := AddTimeLimitLogRetentionPolicyRequestAllOf{}
	this.PolicyName = policyName
	return &this
}

// NewAddTimeLimitLogRetentionPolicyRequestAllOfWithDefaults instantiates a new AddTimeLimitLogRetentionPolicyRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTimeLimitLogRetentionPolicyRequestAllOfWithDefaults() *AddTimeLimitLogRetentionPolicyRequestAllOf {
	this := AddTimeLimitLogRetentionPolicyRequestAllOf{}
	return &this
}

// GetPolicyName returns the PolicyName field value
func (o *AddTimeLimitLogRetentionPolicyRequestAllOf) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *AddTimeLimitLogRetentionPolicyRequestAllOf) GetPolicyNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value
func (o *AddTimeLimitLogRetentionPolicyRequestAllOf) SetPolicyName(v string) {
	o.PolicyName = v
}

func (o AddTimeLimitLogRetentionPolicyRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["policyName"] = o.PolicyName
	}
	return json.Marshal(toSerialize)
}

type NullableAddTimeLimitLogRetentionPolicyRequestAllOf struct {
	value *AddTimeLimitLogRetentionPolicyRequestAllOf
	isSet bool
}

func (v NullableAddTimeLimitLogRetentionPolicyRequestAllOf) Get() *AddTimeLimitLogRetentionPolicyRequestAllOf {
	return v.value
}

func (v *NullableAddTimeLimitLogRetentionPolicyRequestAllOf) Set(val *AddTimeLimitLogRetentionPolicyRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTimeLimitLogRetentionPolicyRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTimeLimitLogRetentionPolicyRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTimeLimitLogRetentionPolicyRequestAllOf(val *AddTimeLimitLogRetentionPolicyRequestAllOf) *NullableAddTimeLimitLogRetentionPolicyRequestAllOf {
	return &NullableAddTimeLimitLogRetentionPolicyRequestAllOf{value: val, isSet: true}
}

func (v NullableAddTimeLimitLogRetentionPolicyRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTimeLimitLogRetentionPolicyRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


