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

// EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp Specifies how the Directory Server should handle operations whenever an attribute value violates the associated attribute syntax.
type EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp string

// List of Enumglobal-configuration-invalidAttributeSyntaxBehaviorProp
const (
	ENUMGLOBALCONFIGURATIONINVALIDATTRIBUTESYNTAXBEHAVIORPROP_ACCEPT EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp = "accept"
	ENUMGLOBALCONFIGURATIONINVALIDATTRIBUTESYNTAXBEHAVIORPROP_REJECT EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp = "reject"
	ENUMGLOBALCONFIGURATIONINVALIDATTRIBUTESYNTAXBEHAVIORPROP_WARN   EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp = "warn"
)

// All allowed values of EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp enum
var AllowedEnumglobalConfigurationInvalidAttributeSyntaxBehaviorPropEnumValues = []EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp{
	"accept",
	"reject",
	"warn",
}

func (v *EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp(value)
	for _, existing := range AllowedEnumglobalConfigurationInvalidAttributeSyntaxBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp", value)
}

// NewEnumglobalConfigurationInvalidAttributeSyntaxBehaviorPropFromValue returns a pointer to a valid EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumglobalConfigurationInvalidAttributeSyntaxBehaviorPropFromValue(v string) (*EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp, error) {
	ev := EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp: valid values are %v", v, AllowedEnumglobalConfigurationInvalidAttributeSyntaxBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumglobalConfigurationInvalidAttributeSyntaxBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumglobal-configuration-invalidAttributeSyntaxBehaviorProp value
func (v EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp) Ptr() *EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp {
	return &v
}

type NullableEnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp struct {
	value *EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp
	isSet bool
}

func (v NullableEnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp) Get() *EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp {
	return v.value
}

func (v *NullableEnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp) Set(val *EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp(val *EnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp) *NullableEnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp {
	return &NullableEnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumglobalConfigurationInvalidAttributeSyntaxBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}