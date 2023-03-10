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

// EnumscimSubattributeTypeProp Specifies the data type for this sub-attribute.
type EnumscimSubattributeTypeProp string

// List of Enumscim-subattribute-typeProp
const (
	ENUMSCIMSUBATTRIBUTETYPEPROP_STRING    EnumscimSubattributeTypeProp = "string"
	ENUMSCIMSUBATTRIBUTETYPEPROP_BOOLEAN   EnumscimSubattributeTypeProp = "boolean"
	ENUMSCIMSUBATTRIBUTETYPEPROP_DATETIME  EnumscimSubattributeTypeProp = "datetime"
	ENUMSCIMSUBATTRIBUTETYPEPROP_DECIMAL   EnumscimSubattributeTypeProp = "decimal"
	ENUMSCIMSUBATTRIBUTETYPEPROP_INTEGER   EnumscimSubattributeTypeProp = "integer"
	ENUMSCIMSUBATTRIBUTETYPEPROP_BINARY    EnumscimSubattributeTypeProp = "binary"
	ENUMSCIMSUBATTRIBUTETYPEPROP_REFERENCE EnumscimSubattributeTypeProp = "reference"
)

// All allowed values of EnumscimSubattributeTypeProp enum
var AllowedEnumscimSubattributeTypePropEnumValues = []EnumscimSubattributeTypeProp{
	"string",
	"boolean",
	"datetime",
	"decimal",
	"integer",
	"binary",
	"reference",
}

func (v *EnumscimSubattributeTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumscimSubattributeTypeProp(value)
	for _, existing := range AllowedEnumscimSubattributeTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumscimSubattributeTypeProp", value)
}

// NewEnumscimSubattributeTypePropFromValue returns a pointer to a valid EnumscimSubattributeTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumscimSubattributeTypePropFromValue(v string) (*EnumscimSubattributeTypeProp, error) {
	ev := EnumscimSubattributeTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumscimSubattributeTypeProp: valid values are %v", v, AllowedEnumscimSubattributeTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumscimSubattributeTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumscimSubattributeTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumscim-subattribute-typeProp value
func (v EnumscimSubattributeTypeProp) Ptr() *EnumscimSubattributeTypeProp {
	return &v
}

type NullableEnumscimSubattributeTypeProp struct {
	value *EnumscimSubattributeTypeProp
	isSet bool
}

func (v NullableEnumscimSubattributeTypeProp) Get() *EnumscimSubattributeTypeProp {
	return v.value
}

func (v *NullableEnumscimSubattributeTypeProp) Set(val *EnumscimSubattributeTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumscimSubattributeTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumscimSubattributeTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumscimSubattributeTypeProp(val *EnumscimSubattributeTypeProp) *NullableEnumscimSubattributeTypeProp {
	return &NullableEnumscimSubattributeTypeProp{value: val, isSet: true}
}

func (v NullableEnumscimSubattributeTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumscimSubattributeTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
