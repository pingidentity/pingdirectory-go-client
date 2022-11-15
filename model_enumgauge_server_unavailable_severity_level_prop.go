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

// EnumgaugeServerUnavailableSeverityLevelProp Specifies the alarm severity level at or above which the server is considered unavailable.
type EnumgaugeServerUnavailableSeverityLevelProp string

// List of Enumgauge-serverUnavailableSeverityLevelProp
const (
	ENUMGAUGESERVERUNAVAILABLESEVERITYLEVELPROP_CRITICAL EnumgaugeServerUnavailableSeverityLevelProp = "critical"
	ENUMGAUGESERVERUNAVAILABLESEVERITYLEVELPROP_MAJOR EnumgaugeServerUnavailableSeverityLevelProp = "major"
	ENUMGAUGESERVERUNAVAILABLESEVERITYLEVELPROP_MINOR EnumgaugeServerUnavailableSeverityLevelProp = "minor"
	ENUMGAUGESERVERUNAVAILABLESEVERITYLEVELPROP_WARNING EnumgaugeServerUnavailableSeverityLevelProp = "warning"
	ENUMGAUGESERVERUNAVAILABLESEVERITYLEVELPROP_NONE EnumgaugeServerUnavailableSeverityLevelProp = "none"
)

// All allowed values of EnumgaugeServerUnavailableSeverityLevelProp enum
var AllowedEnumgaugeServerUnavailableSeverityLevelPropEnumValues = []EnumgaugeServerUnavailableSeverityLevelProp{
	"critical",
	"major",
	"minor",
	"warning",
	"none",
}

func (v *EnumgaugeServerUnavailableSeverityLevelProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgaugeServerUnavailableSeverityLevelProp(value)
	for _, existing := range AllowedEnumgaugeServerUnavailableSeverityLevelPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgaugeServerUnavailableSeverityLevelProp", value)
}

// NewEnumgaugeServerUnavailableSeverityLevelPropFromValue returns a pointer to a valid EnumgaugeServerUnavailableSeverityLevelProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgaugeServerUnavailableSeverityLevelPropFromValue(v string) (*EnumgaugeServerUnavailableSeverityLevelProp, error) {
	ev := EnumgaugeServerUnavailableSeverityLevelProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgaugeServerUnavailableSeverityLevelProp: valid values are %v", v, AllowedEnumgaugeServerUnavailableSeverityLevelPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgaugeServerUnavailableSeverityLevelProp) IsValid() bool {
	for _, existing := range AllowedEnumgaugeServerUnavailableSeverityLevelPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgauge-serverUnavailableSeverityLevelProp value
func (v EnumgaugeServerUnavailableSeverityLevelProp) Ptr() *EnumgaugeServerUnavailableSeverityLevelProp {
	return &v
}

type NullableEnumgaugeServerUnavailableSeverityLevelProp struct {
	value *EnumgaugeServerUnavailableSeverityLevelProp
	isSet bool
}

func (v NullableEnumgaugeServerUnavailableSeverityLevelProp) Get() *EnumgaugeServerUnavailableSeverityLevelProp {
	return v.value
}

func (v *NullableEnumgaugeServerUnavailableSeverityLevelProp) Set(val *EnumgaugeServerUnavailableSeverityLevelProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgaugeServerUnavailableSeverityLevelProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgaugeServerUnavailableSeverityLevelProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgaugeServerUnavailableSeverityLevelProp(val *EnumgaugeServerUnavailableSeverityLevelProp) *NullableEnumgaugeServerUnavailableSeverityLevelProp {
	return &NullableEnumgaugeServerUnavailableSeverityLevelProp{value: val, isSet: true}
}

func (v NullableEnumgaugeServerUnavailableSeverityLevelProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgaugeServerUnavailableSeverityLevelProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

