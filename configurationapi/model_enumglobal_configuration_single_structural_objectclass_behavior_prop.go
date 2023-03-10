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

// EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp Specifies how the Directory Server should handle operations for an entry does not contain a structural object class, or for an entry that contains multiple structural classes.
type EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp string

// List of Enumglobal-configuration-singleStructuralObjectclassBehaviorProp
const (
	ENUMGLOBALCONFIGURATIONSINGLESTRUCTURALOBJECTCLASSBEHAVIORPROP_ACCEPT EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp = "accept"
	ENUMGLOBALCONFIGURATIONSINGLESTRUCTURALOBJECTCLASSBEHAVIORPROP_REJECT EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp = "reject"
	ENUMGLOBALCONFIGURATIONSINGLESTRUCTURALOBJECTCLASSBEHAVIORPROP_WARN   EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp = "warn"
)

// All allowed values of EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp enum
var AllowedEnumglobalConfigurationSingleStructuralObjectclassBehaviorPropEnumValues = []EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp{
	"accept",
	"reject",
	"warn",
}

func (v *EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp(value)
	for _, existing := range AllowedEnumglobalConfigurationSingleStructuralObjectclassBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp", value)
}

// NewEnumglobalConfigurationSingleStructuralObjectclassBehaviorPropFromValue returns a pointer to a valid EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumglobalConfigurationSingleStructuralObjectclassBehaviorPropFromValue(v string) (*EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp, error) {
	ev := EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp: valid values are %v", v, AllowedEnumglobalConfigurationSingleStructuralObjectclassBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumglobalConfigurationSingleStructuralObjectclassBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumglobal-configuration-singleStructuralObjectclassBehaviorProp value
func (v EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp) Ptr() *EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp {
	return &v
}

type NullableEnumglobalConfigurationSingleStructuralObjectclassBehaviorProp struct {
	value *EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp
	isSet bool
}

func (v NullableEnumglobalConfigurationSingleStructuralObjectclassBehaviorProp) Get() *EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp {
	return v.value
}

func (v *NullableEnumglobalConfigurationSingleStructuralObjectclassBehaviorProp) Set(val *EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumglobalConfigurationSingleStructuralObjectclassBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumglobalConfigurationSingleStructuralObjectclassBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumglobalConfigurationSingleStructuralObjectclassBehaviorProp(val *EnumglobalConfigurationSingleStructuralObjectclassBehaviorProp) *NullableEnumglobalConfigurationSingleStructuralObjectclassBehaviorProp {
	return &NullableEnumglobalConfigurationSingleStructuralObjectclassBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumglobalConfigurationSingleStructuralObjectclassBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumglobalConfigurationSingleStructuralObjectclassBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
