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

// EnumpluginServerAccessModeProp Specifies the manner in which external servers should be used for pass-through authentication attempts if multiple servers are defined.
type EnumpluginServerAccessModeProp string

// List of Enumplugin-serverAccessModeProp
const (
	ENUMPLUGINSERVERACCESSMODEPROP_ROUND_ROBIN             EnumpluginServerAccessModeProp = "round-robin"
	ENUMPLUGINSERVERACCESSMODEPROP_FAILOVER_ON_UNAVAILABLE EnumpluginServerAccessModeProp = "failover-on-unavailable"
	ENUMPLUGINSERVERACCESSMODEPROP_FAILOVER_ON_ANY_FAILURE EnumpluginServerAccessModeProp = "failover-on-any-failure"
)

// All allowed values of EnumpluginServerAccessModeProp enum
var AllowedEnumpluginServerAccessModePropEnumValues = []EnumpluginServerAccessModeProp{
	"round-robin",
	"failover-on-unavailable",
	"failover-on-any-failure",
}

func (v *EnumpluginServerAccessModeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginServerAccessModeProp(value)
	for _, existing := range AllowedEnumpluginServerAccessModePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginServerAccessModeProp", value)
}

// NewEnumpluginServerAccessModePropFromValue returns a pointer to a valid EnumpluginServerAccessModeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginServerAccessModePropFromValue(v string) (*EnumpluginServerAccessModeProp, error) {
	ev := EnumpluginServerAccessModeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginServerAccessModeProp: valid values are %v", v, AllowedEnumpluginServerAccessModePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginServerAccessModeProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginServerAccessModePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-serverAccessModeProp value
func (v EnumpluginServerAccessModeProp) Ptr() *EnumpluginServerAccessModeProp {
	return &v
}

type NullableEnumpluginServerAccessModeProp struct {
	value *EnumpluginServerAccessModeProp
	isSet bool
}

func (v NullableEnumpluginServerAccessModeProp) Get() *EnumpluginServerAccessModeProp {
	return v.value
}

func (v *NullableEnumpluginServerAccessModeProp) Set(val *EnumpluginServerAccessModeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginServerAccessModeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginServerAccessModeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginServerAccessModeProp(val *EnumpluginServerAccessModeProp) *NullableEnumpluginServerAccessModeProp {
	return &NullableEnumpluginServerAccessModeProp{value: val, isSet: true}
}

func (v NullableEnumpluginServerAccessModeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginServerAccessModeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
