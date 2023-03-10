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

// EnumpluginScopeProp The scope to use for the search.
type EnumpluginScopeProp string

// List of Enumplugin-scopeProp
const (
	ENUMPLUGINSCOPEPROP_BASE                EnumpluginScopeProp = "base"
	ENUMPLUGINSCOPEPROP_ONE                 EnumpluginScopeProp = "one"
	ENUMPLUGINSCOPEPROP_SUB                 EnumpluginScopeProp = "sub"
	ENUMPLUGINSCOPEPROP_SUBORDINATE_SUBTREE EnumpluginScopeProp = "subordinate-subtree"
)

// All allowed values of EnumpluginScopeProp enum
var AllowedEnumpluginScopePropEnumValues = []EnumpluginScopeProp{
	"base",
	"one",
	"sub",
	"subordinate-subtree",
}

func (v *EnumpluginScopeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginScopeProp(value)
	for _, existing := range AllowedEnumpluginScopePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginScopeProp", value)
}

// NewEnumpluginScopePropFromValue returns a pointer to a valid EnumpluginScopeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginScopePropFromValue(v string) (*EnumpluginScopeProp, error) {
	ev := EnumpluginScopeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginScopeProp: valid values are %v", v, AllowedEnumpluginScopePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginScopeProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginScopePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-scopeProp value
func (v EnumpluginScopeProp) Ptr() *EnumpluginScopeProp {
	return &v
}

type NullableEnumpluginScopeProp struct {
	value *EnumpluginScopeProp
	isSet bool
}

func (v NullableEnumpluginScopeProp) Get() *EnumpluginScopeProp {
	return v.value
}

func (v *NullableEnumpluginScopeProp) Set(val *EnumpluginScopeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginScopeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginScopeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginScopeProp(val *EnumpluginScopeProp) *NullableEnumpluginScopeProp {
	return &NullableEnumpluginScopeProp{value: val, isSet: true}
}

func (v NullableEnumpluginScopeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginScopeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
