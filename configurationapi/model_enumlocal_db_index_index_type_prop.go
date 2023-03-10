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

// EnumlocalDbIndexIndexTypeProp Specifies the type(s) of indexing that should be performed for the associated attribute.
type EnumlocalDbIndexIndexTypeProp string

// List of Enumlocal-db-index-indexTypeProp
const (
	ENUMLOCALDBINDEXINDEXTYPEPROP_EQUALITY    EnumlocalDbIndexIndexTypeProp = "equality"
	ENUMLOCALDBINDEXINDEXTYPEPROP_ORDERING    EnumlocalDbIndexIndexTypeProp = "ordering"
	ENUMLOCALDBINDEXINDEXTYPEPROP_PRESENCE    EnumlocalDbIndexIndexTypeProp = "presence"
	ENUMLOCALDBINDEXINDEXTYPEPROP_SUBSTRING   EnumlocalDbIndexIndexTypeProp = "substring"
	ENUMLOCALDBINDEXINDEXTYPEPROP_APPROXIMATE EnumlocalDbIndexIndexTypeProp = "approximate"
)

// All allowed values of EnumlocalDbIndexIndexTypeProp enum
var AllowedEnumlocalDbIndexIndexTypePropEnumValues = []EnumlocalDbIndexIndexTypeProp{
	"equality",
	"ordering",
	"presence",
	"substring",
	"approximate",
}

func (v *EnumlocalDbIndexIndexTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlocalDbIndexIndexTypeProp(value)
	for _, existing := range AllowedEnumlocalDbIndexIndexTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlocalDbIndexIndexTypeProp", value)
}

// NewEnumlocalDbIndexIndexTypePropFromValue returns a pointer to a valid EnumlocalDbIndexIndexTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlocalDbIndexIndexTypePropFromValue(v string) (*EnumlocalDbIndexIndexTypeProp, error) {
	ev := EnumlocalDbIndexIndexTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlocalDbIndexIndexTypeProp: valid values are %v", v, AllowedEnumlocalDbIndexIndexTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlocalDbIndexIndexTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumlocalDbIndexIndexTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlocal-db-index-indexTypeProp value
func (v EnumlocalDbIndexIndexTypeProp) Ptr() *EnumlocalDbIndexIndexTypeProp {
	return &v
}

type NullableEnumlocalDbIndexIndexTypeProp struct {
	value *EnumlocalDbIndexIndexTypeProp
	isSet bool
}

func (v NullableEnumlocalDbIndexIndexTypeProp) Get() *EnumlocalDbIndexIndexTypeProp {
	return v.value
}

func (v *NullableEnumlocalDbIndexIndexTypeProp) Set(val *EnumlocalDbIndexIndexTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlocalDbIndexIndexTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlocalDbIndexIndexTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlocalDbIndexIndexTypeProp(val *EnumlocalDbIndexIndexTypeProp) *NullableEnumlocalDbIndexIndexTypeProp {
	return &NullableEnumlocalDbIndexIndexTypeProp{value: val, isSet: true}
}

func (v NullableEnumlocalDbIndexIndexTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlocalDbIndexIndexTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
