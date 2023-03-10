/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"fmt"
)

// EnumpasswordPolicyPasswordRetirementBehaviorProp Specifies the conditions under which the server may retire a user's current password in the course of setting a new password for that user (whether via a modify operation or a password modify extended operation).
type EnumpasswordPolicyPasswordRetirementBehaviorProp string

// List of Enumpassword-policy-passwordRetirementBehaviorProp
const (
	ENUMPASSWORDPOLICYPASSWORDRETIREMENTBEHAVIORPROP_SELF_CHANGE          EnumpasswordPolicyPasswordRetirementBehaviorProp = "retire-on-self-change"
	ENUMPASSWORDPOLICYPASSWORDRETIREMENTBEHAVIORPROP_ADMINISTRATIVE_RESET EnumpasswordPolicyPasswordRetirementBehaviorProp = "retire-on-administrative-reset"
	ENUMPASSWORDPOLICYPASSWORDRETIREMENTBEHAVIORPROP_REQUEST_WITH_CONTROL EnumpasswordPolicyPasswordRetirementBehaviorProp = "retire-on-request-with-control"
)

// All allowed values of EnumpasswordPolicyPasswordRetirementBehaviorProp enum
var AllowedEnumpasswordPolicyPasswordRetirementBehaviorPropEnumValues = []EnumpasswordPolicyPasswordRetirementBehaviorProp{
	"retire-on-self-change",
	"retire-on-administrative-reset",
	"retire-on-request-with-control",
}

func (v *EnumpasswordPolicyPasswordRetirementBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpasswordPolicyPasswordRetirementBehaviorProp(value)
	for _, existing := range AllowedEnumpasswordPolicyPasswordRetirementBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpasswordPolicyPasswordRetirementBehaviorProp", value)
}

// NewEnumpasswordPolicyPasswordRetirementBehaviorPropFromValue returns a pointer to a valid EnumpasswordPolicyPasswordRetirementBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpasswordPolicyPasswordRetirementBehaviorPropFromValue(v string) (*EnumpasswordPolicyPasswordRetirementBehaviorProp, error) {
	ev := EnumpasswordPolicyPasswordRetirementBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpasswordPolicyPasswordRetirementBehaviorProp: valid values are %v", v, AllowedEnumpasswordPolicyPasswordRetirementBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpasswordPolicyPasswordRetirementBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumpasswordPolicyPasswordRetirementBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumpassword-policy-passwordRetirementBehaviorProp value
func (v EnumpasswordPolicyPasswordRetirementBehaviorProp) Ptr() *EnumpasswordPolicyPasswordRetirementBehaviorProp {
	return &v
}

type NullableEnumpasswordPolicyPasswordRetirementBehaviorProp struct {
	value *EnumpasswordPolicyPasswordRetirementBehaviorProp
	isSet bool
}

func (v NullableEnumpasswordPolicyPasswordRetirementBehaviorProp) Get() *EnumpasswordPolicyPasswordRetirementBehaviorProp {
	return v.value
}

func (v *NullableEnumpasswordPolicyPasswordRetirementBehaviorProp) Set(val *EnumpasswordPolicyPasswordRetirementBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpasswordPolicyPasswordRetirementBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpasswordPolicyPasswordRetirementBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpasswordPolicyPasswordRetirementBehaviorProp(val *EnumpasswordPolicyPasswordRetirementBehaviorProp) *NullableEnumpasswordPolicyPasswordRetirementBehaviorProp {
	return &NullableEnumpasswordPolicyPasswordRetirementBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumpasswordPolicyPasswordRetirementBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpasswordPolicyPasswordRetirementBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
