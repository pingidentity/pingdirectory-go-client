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

// EnumvirtualAttributeJoinScopeProp The scope for searches used to identify joined entries.
type EnumvirtualAttributeJoinScopeProp string

// List of Enumvirtual-attribute-joinScopeProp
const (
	ENUMVIRTUALATTRIBUTEJOINSCOPEPROP_BASE_OBJECT         EnumvirtualAttributeJoinScopeProp = "base-object"
	ENUMVIRTUALATTRIBUTEJOINSCOPEPROP_SINGLE_LEVEL        EnumvirtualAttributeJoinScopeProp = "single-level"
	ENUMVIRTUALATTRIBUTEJOINSCOPEPROP_WHOLE_SUBTREE       EnumvirtualAttributeJoinScopeProp = "whole-subtree"
	ENUMVIRTUALATTRIBUTEJOINSCOPEPROP_SUBORDINATE_SUBTREE EnumvirtualAttributeJoinScopeProp = "subordinate-subtree"
)

// All allowed values of EnumvirtualAttributeJoinScopeProp enum
var AllowedEnumvirtualAttributeJoinScopePropEnumValues = []EnumvirtualAttributeJoinScopeProp{
	"base-object",
	"single-level",
	"whole-subtree",
	"subordinate-subtree",
}

func (v *EnumvirtualAttributeJoinScopeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumvirtualAttributeJoinScopeProp(value)
	for _, existing := range AllowedEnumvirtualAttributeJoinScopePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumvirtualAttributeJoinScopeProp", value)
}

// NewEnumvirtualAttributeJoinScopePropFromValue returns a pointer to a valid EnumvirtualAttributeJoinScopeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumvirtualAttributeJoinScopePropFromValue(v string) (*EnumvirtualAttributeJoinScopeProp, error) {
	ev := EnumvirtualAttributeJoinScopeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumvirtualAttributeJoinScopeProp: valid values are %v", v, AllowedEnumvirtualAttributeJoinScopePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumvirtualAttributeJoinScopeProp) IsValid() bool {
	for _, existing := range AllowedEnumvirtualAttributeJoinScopePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumvirtual-attribute-joinScopeProp value
func (v EnumvirtualAttributeJoinScopeProp) Ptr() *EnumvirtualAttributeJoinScopeProp {
	return &v
}

type NullableEnumvirtualAttributeJoinScopeProp struct {
	value *EnumvirtualAttributeJoinScopeProp
	isSet bool
}

func (v NullableEnumvirtualAttributeJoinScopeProp) Get() *EnumvirtualAttributeJoinScopeProp {
	return v.value
}

func (v *NullableEnumvirtualAttributeJoinScopeProp) Set(val *EnumvirtualAttributeJoinScopeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumvirtualAttributeJoinScopeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumvirtualAttributeJoinScopeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumvirtualAttributeJoinScopeProp(val *EnumvirtualAttributeJoinScopeProp) *NullableEnumvirtualAttributeJoinScopeProp {
	return &NullableEnumvirtualAttributeJoinScopeProp{value: val, isSet: true}
}

func (v NullableEnumvirtualAttributeJoinScopeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumvirtualAttributeJoinScopeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
