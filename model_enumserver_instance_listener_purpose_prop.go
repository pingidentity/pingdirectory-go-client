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

// EnumserverInstanceListenerPurposeProp Identifies the purpose of this Server Instance Listener.
type EnumserverInstanceListenerPurposeProp string

// List of Enumserver-instance-listener-purposeProp
const (
	ENUMSERVERINSTANCELISTENERPURPOSEPROP_MIRRORED_CONFIG EnumserverInstanceListenerPurposeProp = "mirrored-config"
)

// All allowed values of EnumserverInstanceListenerPurposeProp enum
var AllowedEnumserverInstanceListenerPurposePropEnumValues = []EnumserverInstanceListenerPurposeProp{
	"mirrored-config",
}

func (v *EnumserverInstanceListenerPurposeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumserverInstanceListenerPurposeProp(value)
	for _, existing := range AllowedEnumserverInstanceListenerPurposePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumserverInstanceListenerPurposeProp", value)
}

// NewEnumserverInstanceListenerPurposePropFromValue returns a pointer to a valid EnumserverInstanceListenerPurposeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumserverInstanceListenerPurposePropFromValue(v string) (*EnumserverInstanceListenerPurposeProp, error) {
	ev := EnumserverInstanceListenerPurposeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumserverInstanceListenerPurposeProp: valid values are %v", v, AllowedEnumserverInstanceListenerPurposePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumserverInstanceListenerPurposeProp) IsValid() bool {
	for _, existing := range AllowedEnumserverInstanceListenerPurposePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumserver-instance-listener-purposeProp value
func (v EnumserverInstanceListenerPurposeProp) Ptr() *EnumserverInstanceListenerPurposeProp {
	return &v
}

type NullableEnumserverInstanceListenerPurposeProp struct {
	value *EnumserverInstanceListenerPurposeProp
	isSet bool
}

func (v NullableEnumserverInstanceListenerPurposeProp) Get() *EnumserverInstanceListenerPurposeProp {
	return v.value
}

func (v *NullableEnumserverInstanceListenerPurposeProp) Set(val *EnumserverInstanceListenerPurposeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumserverInstanceListenerPurposeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumserverInstanceListenerPurposeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumserverInstanceListenerPurposeProp(val *EnumserverInstanceListenerPurposeProp) *NullableEnumserverInstanceListenerPurposeProp {
	return &NullableEnumserverInstanceListenerPurposeProp{value: val, isSet: true}
}

func (v NullableEnumserverInstanceListenerPurposeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumserverInstanceListenerPurposeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
