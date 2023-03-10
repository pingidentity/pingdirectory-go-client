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

// EnumsensitiveAttributeAllowInAddProp Indicates whether clients will be allowed to include sensitive attributes in add requests.
type EnumsensitiveAttributeAllowInAddProp string

// List of Enumsensitive-attribute-allowInAddProp
const (
	ENUMSENSITIVEATTRIBUTEALLOWINADDPROP_ALLOW       EnumsensitiveAttributeAllowInAddProp = "allow"
	ENUMSENSITIVEATTRIBUTEALLOWINADDPROP_REJECT      EnumsensitiveAttributeAllowInAddProp = "reject"
	ENUMSENSITIVEATTRIBUTEALLOWINADDPROP_SECURE_ONLY EnumsensitiveAttributeAllowInAddProp = "secure-only"
)

// All allowed values of EnumsensitiveAttributeAllowInAddProp enum
var AllowedEnumsensitiveAttributeAllowInAddPropEnumValues = []EnumsensitiveAttributeAllowInAddProp{
	"allow",
	"reject",
	"secure-only",
}

func (v *EnumsensitiveAttributeAllowInAddProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsensitiveAttributeAllowInAddProp(value)
	for _, existing := range AllowedEnumsensitiveAttributeAllowInAddPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsensitiveAttributeAllowInAddProp", value)
}

// NewEnumsensitiveAttributeAllowInAddPropFromValue returns a pointer to a valid EnumsensitiveAttributeAllowInAddProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsensitiveAttributeAllowInAddPropFromValue(v string) (*EnumsensitiveAttributeAllowInAddProp, error) {
	ev := EnumsensitiveAttributeAllowInAddProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsensitiveAttributeAllowInAddProp: valid values are %v", v, AllowedEnumsensitiveAttributeAllowInAddPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsensitiveAttributeAllowInAddProp) IsValid() bool {
	for _, existing := range AllowedEnumsensitiveAttributeAllowInAddPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsensitive-attribute-allowInAddProp value
func (v EnumsensitiveAttributeAllowInAddProp) Ptr() *EnumsensitiveAttributeAllowInAddProp {
	return &v
}

type NullableEnumsensitiveAttributeAllowInAddProp struct {
	value *EnumsensitiveAttributeAllowInAddProp
	isSet bool
}

func (v NullableEnumsensitiveAttributeAllowInAddProp) Get() *EnumsensitiveAttributeAllowInAddProp {
	return v.value
}

func (v *NullableEnumsensitiveAttributeAllowInAddProp) Set(val *EnumsensitiveAttributeAllowInAddProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsensitiveAttributeAllowInAddProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsensitiveAttributeAllowInAddProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsensitiveAttributeAllowInAddProp(val *EnumsensitiveAttributeAllowInAddProp) *NullableEnumsensitiveAttributeAllowInAddProp {
	return &NullableEnumsensitiveAttributeAllowInAddProp{value: val, isSet: true}
}

func (v NullableEnumsensitiveAttributeAllowInAddProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsensitiveAttributeAllowInAddProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
