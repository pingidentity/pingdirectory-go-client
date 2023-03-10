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

// EnumsensitiveAttributeAllowInCompareProp Indicates whether clients will be allowed to target sensitive attributes with compare requests.
type EnumsensitiveAttributeAllowInCompareProp string

// List of Enumsensitive-attribute-allowInCompareProp
const (
	ENUMSENSITIVEATTRIBUTEALLOWINCOMPAREPROP_ALLOW       EnumsensitiveAttributeAllowInCompareProp = "allow"
	ENUMSENSITIVEATTRIBUTEALLOWINCOMPAREPROP_REJECT      EnumsensitiveAttributeAllowInCompareProp = "reject"
	ENUMSENSITIVEATTRIBUTEALLOWINCOMPAREPROP_SECURE_ONLY EnumsensitiveAttributeAllowInCompareProp = "secure-only"
)

// All allowed values of EnumsensitiveAttributeAllowInCompareProp enum
var AllowedEnumsensitiveAttributeAllowInComparePropEnumValues = []EnumsensitiveAttributeAllowInCompareProp{
	"allow",
	"reject",
	"secure-only",
}

func (v *EnumsensitiveAttributeAllowInCompareProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsensitiveAttributeAllowInCompareProp(value)
	for _, existing := range AllowedEnumsensitiveAttributeAllowInComparePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsensitiveAttributeAllowInCompareProp", value)
}

// NewEnumsensitiveAttributeAllowInComparePropFromValue returns a pointer to a valid EnumsensitiveAttributeAllowInCompareProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsensitiveAttributeAllowInComparePropFromValue(v string) (*EnumsensitiveAttributeAllowInCompareProp, error) {
	ev := EnumsensitiveAttributeAllowInCompareProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsensitiveAttributeAllowInCompareProp: valid values are %v", v, AllowedEnumsensitiveAttributeAllowInComparePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsensitiveAttributeAllowInCompareProp) IsValid() bool {
	for _, existing := range AllowedEnumsensitiveAttributeAllowInComparePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsensitive-attribute-allowInCompareProp value
func (v EnumsensitiveAttributeAllowInCompareProp) Ptr() *EnumsensitiveAttributeAllowInCompareProp {
	return &v
}

type NullableEnumsensitiveAttributeAllowInCompareProp struct {
	value *EnumsensitiveAttributeAllowInCompareProp
	isSet bool
}

func (v NullableEnumsensitiveAttributeAllowInCompareProp) Get() *EnumsensitiveAttributeAllowInCompareProp {
	return v.value
}

func (v *NullableEnumsensitiveAttributeAllowInCompareProp) Set(val *EnumsensitiveAttributeAllowInCompareProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsensitiveAttributeAllowInCompareProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsensitiveAttributeAllowInCompareProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsensitiveAttributeAllowInCompareProp(val *EnumsensitiveAttributeAllowInCompareProp) *NullableEnumsensitiveAttributeAllowInCompareProp {
	return &NullableEnumsensitiveAttributeAllowInCompareProp{value: val, isSet: true}
}

func (v NullableEnumsensitiveAttributeAllowInCompareProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsensitiveAttributeAllowInCompareProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
