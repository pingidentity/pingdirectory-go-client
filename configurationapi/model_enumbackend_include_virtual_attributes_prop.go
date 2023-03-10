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

// EnumbackendIncludeVirtualAttributesProp Specifies the changelog entry elements (if any) in which virtual attributes should be included.
type EnumbackendIncludeVirtualAttributesProp string

// List of Enumbackend-includeVirtualAttributesProp
const (
	ENUMBACKENDINCLUDEVIRTUALATTRIBUTESPROP_ADD_ATTRIBUTES           EnumbackendIncludeVirtualAttributesProp = "add-attributes"
	ENUMBACKENDINCLUDEVIRTUALATTRIBUTESPROP_DELETED_ENTRY_ATTRIBUTES EnumbackendIncludeVirtualAttributesProp = "deleted-entry-attributes"
	ENUMBACKENDINCLUDEVIRTUALATTRIBUTESPROP_BEFORE_AND_AFTER_VALUES  EnumbackendIncludeVirtualAttributesProp = "before-and-after-values"
	ENUMBACKENDINCLUDEVIRTUALATTRIBUTESPROP_KEY_ATTRIBUTE_VALUES     EnumbackendIncludeVirtualAttributesProp = "key-attribute-values"
)

// All allowed values of EnumbackendIncludeVirtualAttributesProp enum
var AllowedEnumbackendIncludeVirtualAttributesPropEnumValues = []EnumbackendIncludeVirtualAttributesProp{
	"add-attributes",
	"deleted-entry-attributes",
	"before-and-after-values",
	"key-attribute-values",
}

func (v *EnumbackendIncludeVirtualAttributesProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumbackendIncludeVirtualAttributesProp(value)
	for _, existing := range AllowedEnumbackendIncludeVirtualAttributesPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumbackendIncludeVirtualAttributesProp", value)
}

// NewEnumbackendIncludeVirtualAttributesPropFromValue returns a pointer to a valid EnumbackendIncludeVirtualAttributesProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumbackendIncludeVirtualAttributesPropFromValue(v string) (*EnumbackendIncludeVirtualAttributesProp, error) {
	ev := EnumbackendIncludeVirtualAttributesProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumbackendIncludeVirtualAttributesProp: valid values are %v", v, AllowedEnumbackendIncludeVirtualAttributesPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumbackendIncludeVirtualAttributesProp) IsValid() bool {
	for _, existing := range AllowedEnumbackendIncludeVirtualAttributesPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumbackend-includeVirtualAttributesProp value
func (v EnumbackendIncludeVirtualAttributesProp) Ptr() *EnumbackendIncludeVirtualAttributesProp {
	return &v
}

type NullableEnumbackendIncludeVirtualAttributesProp struct {
	value *EnumbackendIncludeVirtualAttributesProp
	isSet bool
}

func (v NullableEnumbackendIncludeVirtualAttributesProp) Get() *EnumbackendIncludeVirtualAttributesProp {
	return v.value
}

func (v *NullableEnumbackendIncludeVirtualAttributesProp) Set(val *EnumbackendIncludeVirtualAttributesProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumbackendIncludeVirtualAttributesProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumbackendIncludeVirtualAttributesProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumbackendIncludeVirtualAttributesProp(val *EnumbackendIncludeVirtualAttributesProp) *NullableEnumbackendIncludeVirtualAttributesProp {
	return &NullableEnumbackendIncludeVirtualAttributesProp{value: val, isSet: true}
}

func (v NullableEnumbackendIncludeVirtualAttributesProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumbackendIncludeVirtualAttributesProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
