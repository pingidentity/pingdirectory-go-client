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

// EnumvirtualAttributeConflictBehaviorProp Specifies the behavior that the server is to exhibit for entries that already contain one or more real values for the associated attribute.
type EnumvirtualAttributeConflictBehaviorProp string

// List of Enumvirtual-attribute-conflictBehaviorProp
const (
	ENUMVIRTUALATTRIBUTECONFLICTBEHAVIORPROP_REAL_OVERRIDES_VIRTUAL EnumvirtualAttributeConflictBehaviorProp = "real-overrides-virtual"
	ENUMVIRTUALATTRIBUTECONFLICTBEHAVIORPROP_VIRTUAL_OVERRIDES_REAL EnumvirtualAttributeConflictBehaviorProp = "virtual-overrides-real"
	ENUMVIRTUALATTRIBUTECONFLICTBEHAVIORPROP_MERGE_REAL_AND_VIRTUAL EnumvirtualAttributeConflictBehaviorProp = "merge-real-and-virtual"
)

// All allowed values of EnumvirtualAttributeConflictBehaviorProp enum
var AllowedEnumvirtualAttributeConflictBehaviorPropEnumValues = []EnumvirtualAttributeConflictBehaviorProp{
	"real-overrides-virtual",
	"virtual-overrides-real",
	"merge-real-and-virtual",
}

func (v *EnumvirtualAttributeConflictBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumvirtualAttributeConflictBehaviorProp(value)
	for _, existing := range AllowedEnumvirtualAttributeConflictBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumvirtualAttributeConflictBehaviorProp", value)
}

// NewEnumvirtualAttributeConflictBehaviorPropFromValue returns a pointer to a valid EnumvirtualAttributeConflictBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumvirtualAttributeConflictBehaviorPropFromValue(v string) (*EnumvirtualAttributeConflictBehaviorProp, error) {
	ev := EnumvirtualAttributeConflictBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumvirtualAttributeConflictBehaviorProp: valid values are %v", v, AllowedEnumvirtualAttributeConflictBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumvirtualAttributeConflictBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumvirtualAttributeConflictBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumvirtual-attribute-conflictBehaviorProp value
func (v EnumvirtualAttributeConflictBehaviorProp) Ptr() *EnumvirtualAttributeConflictBehaviorProp {
	return &v
}

type NullableEnumvirtualAttributeConflictBehaviorProp struct {
	value *EnumvirtualAttributeConflictBehaviorProp
	isSet bool
}

func (v NullableEnumvirtualAttributeConflictBehaviorProp) Get() *EnumvirtualAttributeConflictBehaviorProp {
	return v.value
}

func (v *NullableEnumvirtualAttributeConflictBehaviorProp) Set(val *EnumvirtualAttributeConflictBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumvirtualAttributeConflictBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumvirtualAttributeConflictBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumvirtualAttributeConflictBehaviorProp(val *EnumvirtualAttributeConflictBehaviorProp) *NullableEnumvirtualAttributeConflictBehaviorProp {
	return &NullableEnumvirtualAttributeConflictBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumvirtualAttributeConflictBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumvirtualAttributeConflictBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}