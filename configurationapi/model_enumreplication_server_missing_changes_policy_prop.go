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

// EnumreplicationServerMissingChangesPolicyProp Determines how the server responds when replication detects that some changes might have been missed. Each missing changes policy is a set of missing changes actions to take for a set of missing changes types. The value configured here acts as a default for all replication domains on this replication server.
type EnumreplicationServerMissingChangesPolicyProp string

// List of Enumreplication-server-missingChangesPolicyProp
const (
	ENUMREPLICATIONSERVERMISSINGCHANGESPOLICYPROP_MAXIMUM_INTEGRITY    EnumreplicationServerMissingChangesPolicyProp = "maximum-integrity"
	ENUMREPLICATIONSERVERMISSINGCHANGESPOLICYPROP_FAVOR_INTEGRITY      EnumreplicationServerMissingChangesPolicyProp = "favor-integrity"
	ENUMREPLICATIONSERVERMISSINGCHANGESPOLICYPROP_FAVOR_AVAILABILITY   EnumreplicationServerMissingChangesPolicyProp = "favor-availability"
	ENUMREPLICATIONSERVERMISSINGCHANGESPOLICYPROP_MAXIMUM_AVAILABILITY EnumreplicationServerMissingChangesPolicyProp = "maximum-availability"
)

// All allowed values of EnumreplicationServerMissingChangesPolicyProp enum
var AllowedEnumreplicationServerMissingChangesPolicyPropEnumValues = []EnumreplicationServerMissingChangesPolicyProp{
	"maximum-integrity",
	"favor-integrity",
	"favor-availability",
	"maximum-availability",
}

func (v *EnumreplicationServerMissingChangesPolicyProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumreplicationServerMissingChangesPolicyProp(value)
	for _, existing := range AllowedEnumreplicationServerMissingChangesPolicyPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumreplicationServerMissingChangesPolicyProp", value)
}

// NewEnumreplicationServerMissingChangesPolicyPropFromValue returns a pointer to a valid EnumreplicationServerMissingChangesPolicyProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumreplicationServerMissingChangesPolicyPropFromValue(v string) (*EnumreplicationServerMissingChangesPolicyProp, error) {
	ev := EnumreplicationServerMissingChangesPolicyProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumreplicationServerMissingChangesPolicyProp: valid values are %v", v, AllowedEnumreplicationServerMissingChangesPolicyPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumreplicationServerMissingChangesPolicyProp) IsValid() bool {
	for _, existing := range AllowedEnumreplicationServerMissingChangesPolicyPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumreplication-server-missingChangesPolicyProp value
func (v EnumreplicationServerMissingChangesPolicyProp) Ptr() *EnumreplicationServerMissingChangesPolicyProp {
	return &v
}

type NullableEnumreplicationServerMissingChangesPolicyProp struct {
	value *EnumreplicationServerMissingChangesPolicyProp
	isSet bool
}

func (v NullableEnumreplicationServerMissingChangesPolicyProp) Get() *EnumreplicationServerMissingChangesPolicyProp {
	return v.value
}

func (v *NullableEnumreplicationServerMissingChangesPolicyProp) Set(val *EnumreplicationServerMissingChangesPolicyProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumreplicationServerMissingChangesPolicyProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumreplicationServerMissingChangesPolicyProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumreplicationServerMissingChangesPolicyProp(val *EnumreplicationServerMissingChangesPolicyProp) *NullableEnumreplicationServerMissingChangesPolicyProp {
	return &NullableEnumreplicationServerMissingChangesPolicyProp{value: val, isSet: true}
}

func (v NullableEnumreplicationServerMissingChangesPolicyProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumreplicationServerMissingChangesPolicyProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
