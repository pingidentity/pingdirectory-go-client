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

// EnumpassThroughAuthenticationHandlerServerAccessModeProp Specifies the manner in which external servers should be used for pass-through authentication attempts if multiple servers are defined.
type EnumpassThroughAuthenticationHandlerServerAccessModeProp string

// List of Enumpass-through-authentication-handler-serverAccessModeProp
const (
	ENUMPASSTHROUGHAUTHENTICATIONHANDLERSERVERACCESSMODEPROP_ROUND_ROBIN EnumpassThroughAuthenticationHandlerServerAccessModeProp = "round-robin"
	ENUMPASSTHROUGHAUTHENTICATIONHANDLERSERVERACCESSMODEPROP_FAILOVER_ON_UNAVAILABLE EnumpassThroughAuthenticationHandlerServerAccessModeProp = "failover-on-unavailable"
	ENUMPASSTHROUGHAUTHENTICATIONHANDLERSERVERACCESSMODEPROP_FAILOVER_ON_ANY_FAILURE EnumpassThroughAuthenticationHandlerServerAccessModeProp = "failover-on-any-failure"
)

// All allowed values of EnumpassThroughAuthenticationHandlerServerAccessModeProp enum
var AllowedEnumpassThroughAuthenticationHandlerServerAccessModePropEnumValues = []EnumpassThroughAuthenticationHandlerServerAccessModeProp{
	"round-robin",
	"failover-on-unavailable",
	"failover-on-any-failure",
}

func (v *EnumpassThroughAuthenticationHandlerServerAccessModeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpassThroughAuthenticationHandlerServerAccessModeProp(value)
	for _, existing := range AllowedEnumpassThroughAuthenticationHandlerServerAccessModePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpassThroughAuthenticationHandlerServerAccessModeProp", value)
}

// NewEnumpassThroughAuthenticationHandlerServerAccessModePropFromValue returns a pointer to a valid EnumpassThroughAuthenticationHandlerServerAccessModeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpassThroughAuthenticationHandlerServerAccessModePropFromValue(v string) (*EnumpassThroughAuthenticationHandlerServerAccessModeProp, error) {
	ev := EnumpassThroughAuthenticationHandlerServerAccessModeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpassThroughAuthenticationHandlerServerAccessModeProp: valid values are %v", v, AllowedEnumpassThroughAuthenticationHandlerServerAccessModePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpassThroughAuthenticationHandlerServerAccessModeProp) IsValid() bool {
	for _, existing := range AllowedEnumpassThroughAuthenticationHandlerServerAccessModePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumpass-through-authentication-handler-serverAccessModeProp value
func (v EnumpassThroughAuthenticationHandlerServerAccessModeProp) Ptr() *EnumpassThroughAuthenticationHandlerServerAccessModeProp {
	return &v
}

type NullableEnumpassThroughAuthenticationHandlerServerAccessModeProp struct {
	value *EnumpassThroughAuthenticationHandlerServerAccessModeProp
	isSet bool
}

func (v NullableEnumpassThroughAuthenticationHandlerServerAccessModeProp) Get() *EnumpassThroughAuthenticationHandlerServerAccessModeProp {
	return v.value
}

func (v *NullableEnumpassThroughAuthenticationHandlerServerAccessModeProp) Set(val *EnumpassThroughAuthenticationHandlerServerAccessModeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpassThroughAuthenticationHandlerServerAccessModeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpassThroughAuthenticationHandlerServerAccessModeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpassThroughAuthenticationHandlerServerAccessModeProp(val *EnumpassThroughAuthenticationHandlerServerAccessModeProp) *NullableEnumpassThroughAuthenticationHandlerServerAccessModeProp {
	return &NullableEnumpassThroughAuthenticationHandlerServerAccessModeProp{value: val, isSet: true}
}

func (v NullableEnumpassThroughAuthenticationHandlerServerAccessModeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpassThroughAuthenticationHandlerServerAccessModeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

