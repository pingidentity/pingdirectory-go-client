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

// EnumlogPublisherPolicyMessageTypeProp Specifies the policy message types to be logged.
type EnumlogPublisherPolicyMessageTypeProp string

// List of Enumlog-publisher-policyMessageTypeProp
const (
	ENUMLOGPUBLISHERPOLICYMESSAGETYPEPROP_DECISION EnumlogPublisherPolicyMessageTypeProp = "decision"
	ENUMLOGPUBLISHERPOLICYMESSAGETYPEPROP_ADVICE   EnumlogPublisherPolicyMessageTypeProp = "advice"
)

// All allowed values of EnumlogPublisherPolicyMessageTypeProp enum
var AllowedEnumlogPublisherPolicyMessageTypePropEnumValues = []EnumlogPublisherPolicyMessageTypeProp{
	"decision",
	"advice",
}

func (v *EnumlogPublisherPolicyMessageTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherPolicyMessageTypeProp(value)
	for _, existing := range AllowedEnumlogPublisherPolicyMessageTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherPolicyMessageTypeProp", value)
}

// NewEnumlogPublisherPolicyMessageTypePropFromValue returns a pointer to a valid EnumlogPublisherPolicyMessageTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherPolicyMessageTypePropFromValue(v string) (*EnumlogPublisherPolicyMessageTypeProp, error) {
	ev := EnumlogPublisherPolicyMessageTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherPolicyMessageTypeProp: valid values are %v", v, AllowedEnumlogPublisherPolicyMessageTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherPolicyMessageTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherPolicyMessageTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-policyMessageTypeProp value
func (v EnumlogPublisherPolicyMessageTypeProp) Ptr() *EnumlogPublisherPolicyMessageTypeProp {
	return &v
}

type NullableEnumlogPublisherPolicyMessageTypeProp struct {
	value *EnumlogPublisherPolicyMessageTypeProp
	isSet bool
}

func (v NullableEnumlogPublisherPolicyMessageTypeProp) Get() *EnumlogPublisherPolicyMessageTypeProp {
	return v.value
}

func (v *NullableEnumlogPublisherPolicyMessageTypeProp) Set(val *EnumlogPublisherPolicyMessageTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherPolicyMessageTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherPolicyMessageTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherPolicyMessageTypeProp(val *EnumlogPublisherPolicyMessageTypeProp) *NullableEnumlogPublisherPolicyMessageTypeProp {
	return &NullableEnumlogPublisherPolicyMessageTypeProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherPolicyMessageTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherPolicyMessageTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
