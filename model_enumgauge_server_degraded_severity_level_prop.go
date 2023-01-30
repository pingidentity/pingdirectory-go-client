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

// EnumgaugeServerDegradedSeverityLevelProp Specifies the alarm severity level at or above which the server is considered degraded.
type EnumgaugeServerDegradedSeverityLevelProp string

// List of Enumgauge-serverDegradedSeverityLevelProp
const (
	ENUMGAUGESERVERDEGRADEDSEVERITYLEVELPROP_CRITICAL EnumgaugeServerDegradedSeverityLevelProp = "critical"
	ENUMGAUGESERVERDEGRADEDSEVERITYLEVELPROP_MAJOR    EnumgaugeServerDegradedSeverityLevelProp = "major"
	ENUMGAUGESERVERDEGRADEDSEVERITYLEVELPROP_MINOR    EnumgaugeServerDegradedSeverityLevelProp = "minor"
	ENUMGAUGESERVERDEGRADEDSEVERITYLEVELPROP_WARNING  EnumgaugeServerDegradedSeverityLevelProp = "warning"
	ENUMGAUGESERVERDEGRADEDSEVERITYLEVELPROP_NONE     EnumgaugeServerDegradedSeverityLevelProp = "none"
)

// All allowed values of EnumgaugeServerDegradedSeverityLevelProp enum
var AllowedEnumgaugeServerDegradedSeverityLevelPropEnumValues = []EnumgaugeServerDegradedSeverityLevelProp{
	"critical",
	"major",
	"minor",
	"warning",
	"none",
}

func (v *EnumgaugeServerDegradedSeverityLevelProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgaugeServerDegradedSeverityLevelProp(value)
	for _, existing := range AllowedEnumgaugeServerDegradedSeverityLevelPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgaugeServerDegradedSeverityLevelProp", value)
}

// NewEnumgaugeServerDegradedSeverityLevelPropFromValue returns a pointer to a valid EnumgaugeServerDegradedSeverityLevelProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgaugeServerDegradedSeverityLevelPropFromValue(v string) (*EnumgaugeServerDegradedSeverityLevelProp, error) {
	ev := EnumgaugeServerDegradedSeverityLevelProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgaugeServerDegradedSeverityLevelProp: valid values are %v", v, AllowedEnumgaugeServerDegradedSeverityLevelPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgaugeServerDegradedSeverityLevelProp) IsValid() bool {
	for _, existing := range AllowedEnumgaugeServerDegradedSeverityLevelPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgauge-serverDegradedSeverityLevelProp value
func (v EnumgaugeServerDegradedSeverityLevelProp) Ptr() *EnumgaugeServerDegradedSeverityLevelProp {
	return &v
}

type NullableEnumgaugeServerDegradedSeverityLevelProp struct {
	value *EnumgaugeServerDegradedSeverityLevelProp
	isSet bool
}

func (v NullableEnumgaugeServerDegradedSeverityLevelProp) Get() *EnumgaugeServerDegradedSeverityLevelProp {
	return v.value
}

func (v *NullableEnumgaugeServerDegradedSeverityLevelProp) Set(val *EnumgaugeServerDegradedSeverityLevelProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgaugeServerDegradedSeverityLevelProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgaugeServerDegradedSeverityLevelProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgaugeServerDegradedSeverityLevelProp(val *EnumgaugeServerDegradedSeverityLevelProp) *NullableEnumgaugeServerDegradedSeverityLevelProp {
	return &NullableEnumgaugeServerDegradedSeverityLevelProp{value: val, isSet: true}
}

func (v NullableEnumgaugeServerDegradedSeverityLevelProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgaugeServerDegradedSeverityLevelProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
