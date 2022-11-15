/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// EnumlogFieldSyntaxDefaultBehaviorProp The default behavior that the server should exhibit when logging fields with this syntax. This may be overridden on a per-field basis.
type EnumlogFieldSyntaxDefaultBehaviorProp string

// List of Enumlog-field-syntax-defaultBehaviorProp
const (
	ENUMLOGFIELDSYNTAXDEFAULTBEHAVIORPROP_PRESERVE EnumlogFieldSyntaxDefaultBehaviorProp = "preserve"
	ENUMLOGFIELDSYNTAXDEFAULTBEHAVIORPROP_OMIT EnumlogFieldSyntaxDefaultBehaviorProp = "omit"
	ENUMLOGFIELDSYNTAXDEFAULTBEHAVIORPROP_REDACT_ENTIRE_VALUE EnumlogFieldSyntaxDefaultBehaviorProp = "redact-entire-value"
	ENUMLOGFIELDSYNTAXDEFAULTBEHAVIORPROP_REDACT_VALUE_COMPONENTS EnumlogFieldSyntaxDefaultBehaviorProp = "redact-value-components"
	ENUMLOGFIELDSYNTAXDEFAULTBEHAVIORPROP_TOKENIZE_ENTIRE_VALUE EnumlogFieldSyntaxDefaultBehaviorProp = "tokenize-entire-value"
	ENUMLOGFIELDSYNTAXDEFAULTBEHAVIORPROP_TOKENIZE_VALUE_COMPONENTS EnumlogFieldSyntaxDefaultBehaviorProp = "tokenize-value-components"
)

// All allowed values of EnumlogFieldSyntaxDefaultBehaviorProp enum
var AllowedEnumlogFieldSyntaxDefaultBehaviorPropEnumValues = []EnumlogFieldSyntaxDefaultBehaviorProp{
	"preserve",
	"omit",
	"redact-entire-value",
	"redact-value-components",
	"tokenize-entire-value",
	"tokenize-value-components",
}

func (v *EnumlogFieldSyntaxDefaultBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogFieldSyntaxDefaultBehaviorProp(value)
	for _, existing := range AllowedEnumlogFieldSyntaxDefaultBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogFieldSyntaxDefaultBehaviorProp", value)
}

// NewEnumlogFieldSyntaxDefaultBehaviorPropFromValue returns a pointer to a valid EnumlogFieldSyntaxDefaultBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogFieldSyntaxDefaultBehaviorPropFromValue(v string) (*EnumlogFieldSyntaxDefaultBehaviorProp, error) {
	ev := EnumlogFieldSyntaxDefaultBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogFieldSyntaxDefaultBehaviorProp: valid values are %v", v, AllowedEnumlogFieldSyntaxDefaultBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogFieldSyntaxDefaultBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumlogFieldSyntaxDefaultBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-field-syntax-defaultBehaviorProp value
func (v EnumlogFieldSyntaxDefaultBehaviorProp) Ptr() *EnumlogFieldSyntaxDefaultBehaviorProp {
	return &v
}

type NullableEnumlogFieldSyntaxDefaultBehaviorProp struct {
	value *EnumlogFieldSyntaxDefaultBehaviorProp
	isSet bool
}

func (v NullableEnumlogFieldSyntaxDefaultBehaviorProp) Get() *EnumlogFieldSyntaxDefaultBehaviorProp {
	return v.value
}

func (v *NullableEnumlogFieldSyntaxDefaultBehaviorProp) Set(val *EnumlogFieldSyntaxDefaultBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogFieldSyntaxDefaultBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogFieldSyntaxDefaultBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogFieldSyntaxDefaultBehaviorProp(val *EnumlogFieldSyntaxDefaultBehaviorProp) *NullableEnumlogFieldSyntaxDefaultBehaviorProp {
	return &NullableEnumlogFieldSyntaxDefaultBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumlogFieldSyntaxDefaultBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogFieldSyntaxDefaultBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

