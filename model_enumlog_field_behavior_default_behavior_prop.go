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

// EnumlogFieldBehaviorDefaultBehaviorProp The default behavior that the server should exhibit for fields for which no explicit behavior is defined. If no default behavior is defined, the server will fall back to using the default behavior configured for the syntax used for each log field.
type EnumlogFieldBehaviorDefaultBehaviorProp string

// List of Enumlog-field-behavior-defaultBehaviorProp
const (
	ENUMLOGFIELDBEHAVIORDEFAULTBEHAVIORPROP_PRESERVE                  EnumlogFieldBehaviorDefaultBehaviorProp = "preserve"
	ENUMLOGFIELDBEHAVIORDEFAULTBEHAVIORPROP_OMIT                      EnumlogFieldBehaviorDefaultBehaviorProp = "omit"
	ENUMLOGFIELDBEHAVIORDEFAULTBEHAVIORPROP_REDACT_ENTIRE_VALUE       EnumlogFieldBehaviorDefaultBehaviorProp = "redact-entire-value"
	ENUMLOGFIELDBEHAVIORDEFAULTBEHAVIORPROP_REDACT_VALUE_COMPONENTS   EnumlogFieldBehaviorDefaultBehaviorProp = "redact-value-components"
	ENUMLOGFIELDBEHAVIORDEFAULTBEHAVIORPROP_TOKENIZE_ENTIRE_VALUE     EnumlogFieldBehaviorDefaultBehaviorProp = "tokenize-entire-value"
	ENUMLOGFIELDBEHAVIORDEFAULTBEHAVIORPROP_TOKENIZE_VALUE_COMPONENTS EnumlogFieldBehaviorDefaultBehaviorProp = "tokenize-value-components"
)

// All allowed values of EnumlogFieldBehaviorDefaultBehaviorProp enum
var AllowedEnumlogFieldBehaviorDefaultBehaviorPropEnumValues = []EnumlogFieldBehaviorDefaultBehaviorProp{
	"preserve",
	"omit",
	"redact-entire-value",
	"redact-value-components",
	"tokenize-entire-value",
	"tokenize-value-components",
}

func (v *EnumlogFieldBehaviorDefaultBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogFieldBehaviorDefaultBehaviorProp(value)
	for _, existing := range AllowedEnumlogFieldBehaviorDefaultBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogFieldBehaviorDefaultBehaviorProp", value)
}

// NewEnumlogFieldBehaviorDefaultBehaviorPropFromValue returns a pointer to a valid EnumlogFieldBehaviorDefaultBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogFieldBehaviorDefaultBehaviorPropFromValue(v string) (*EnumlogFieldBehaviorDefaultBehaviorProp, error) {
	ev := EnumlogFieldBehaviorDefaultBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogFieldBehaviorDefaultBehaviorProp: valid values are %v", v, AllowedEnumlogFieldBehaviorDefaultBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogFieldBehaviorDefaultBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumlogFieldBehaviorDefaultBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-field-behavior-defaultBehaviorProp value
func (v EnumlogFieldBehaviorDefaultBehaviorProp) Ptr() *EnumlogFieldBehaviorDefaultBehaviorProp {
	return &v
}

type NullableEnumlogFieldBehaviorDefaultBehaviorProp struct {
	value *EnumlogFieldBehaviorDefaultBehaviorProp
	isSet bool
}

func (v NullableEnumlogFieldBehaviorDefaultBehaviorProp) Get() *EnumlogFieldBehaviorDefaultBehaviorProp {
	return v.value
}

func (v *NullableEnumlogFieldBehaviorDefaultBehaviorProp) Set(val *EnumlogFieldBehaviorDefaultBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogFieldBehaviorDefaultBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogFieldBehaviorDefaultBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogFieldBehaviorDefaultBehaviorProp(val *EnumlogFieldBehaviorDefaultBehaviorProp) *NullableEnumlogFieldBehaviorDefaultBehaviorProp {
	return &NullableEnumlogFieldBehaviorDefaultBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumlogFieldBehaviorDefaultBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogFieldBehaviorDefaultBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
