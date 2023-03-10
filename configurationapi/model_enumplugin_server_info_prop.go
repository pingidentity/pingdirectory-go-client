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

// EnumpluginServerInfoProp Specifies whether statistics related to resource utilization such as JVM memory and CPU/Network/Disk utilization.
type EnumpluginServerInfoProp string

// List of Enumplugin-serverInfoProp
const (
	ENUMPLUGINSERVERINFOPROP_NONE     EnumpluginServerInfoProp = "none"
	ENUMPLUGINSERVERINFOPROP_BASIC    EnumpluginServerInfoProp = "basic"
	ENUMPLUGINSERVERINFOPROP_EXTENDED EnumpluginServerInfoProp = "extended"
)

// All allowed values of EnumpluginServerInfoProp enum
var AllowedEnumpluginServerInfoPropEnumValues = []EnumpluginServerInfoProp{
	"none",
	"basic",
	"extended",
}

func (v *EnumpluginServerInfoProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginServerInfoProp(value)
	for _, existing := range AllowedEnumpluginServerInfoPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginServerInfoProp", value)
}

// NewEnumpluginServerInfoPropFromValue returns a pointer to a valid EnumpluginServerInfoProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginServerInfoPropFromValue(v string) (*EnumpluginServerInfoProp, error) {
	ev := EnumpluginServerInfoProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginServerInfoProp: valid values are %v", v, AllowedEnumpluginServerInfoPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginServerInfoProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginServerInfoPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-serverInfoProp value
func (v EnumpluginServerInfoProp) Ptr() *EnumpluginServerInfoProp {
	return &v
}

type NullableEnumpluginServerInfoProp struct {
	value *EnumpluginServerInfoProp
	isSet bool
}

func (v NullableEnumpluginServerInfoProp) Get() *EnumpluginServerInfoProp {
	return v.value
}

func (v *NullableEnumpluginServerInfoProp) Set(val *EnumpluginServerInfoProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginServerInfoProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginServerInfoProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginServerInfoProp(val *EnumpluginServerInfoProp) *NullableEnumpluginServerInfoProp {
	return &NullableEnumpluginServerInfoProp{value: val, isSet: true}
}

func (v NullableEnumpluginServerInfoProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginServerInfoProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
