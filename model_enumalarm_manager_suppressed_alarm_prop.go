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

// EnumalarmManagerSuppressedAlarmProp Specifies the names of the alarm alert types that should be suppressed. If the condition that triggers an alarm in this list occurs, then the alarm will not be raised and no alerts will be generated. Only a subset of alarms can be suppressed in this way. Alarms triggered by a gauge can be disabled by disabling the gauge.
type EnumalarmManagerSuppressedAlarmProp string

// List of Enumalarm-manager-suppressedAlarmProp
const (
	NO_ENABLED_ALERT_HANDLERS EnumalarmManagerSuppressedAlarmProp = "no-enabled-alert-handlers"
	PDP_UNAVAILABLE EnumalarmManagerSuppressedAlarmProp = "pdp-unavailable"
)

// All allowed values of EnumalarmManagerSuppressedAlarmProp enum
var AllowedEnumalarmManagerSuppressedAlarmPropEnumValues = []EnumalarmManagerSuppressedAlarmProp{
	"no-enabled-alert-handlers",
	"pdp-unavailable",
}

func (v *EnumalarmManagerSuppressedAlarmProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumalarmManagerSuppressedAlarmProp(value)
	for _, existing := range AllowedEnumalarmManagerSuppressedAlarmPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumalarmManagerSuppressedAlarmProp", value)
}

// NewEnumalarmManagerSuppressedAlarmPropFromValue returns a pointer to a valid EnumalarmManagerSuppressedAlarmProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumalarmManagerSuppressedAlarmPropFromValue(v string) (*EnumalarmManagerSuppressedAlarmProp, error) {
	ev := EnumalarmManagerSuppressedAlarmProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumalarmManagerSuppressedAlarmProp: valid values are %v", v, AllowedEnumalarmManagerSuppressedAlarmPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumalarmManagerSuppressedAlarmProp) IsValid() bool {
	for _, existing := range AllowedEnumalarmManagerSuppressedAlarmPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumalarm-manager-suppressedAlarmProp value
func (v EnumalarmManagerSuppressedAlarmProp) Ptr() *EnumalarmManagerSuppressedAlarmProp {
	return &v
}

type NullableEnumalarmManagerSuppressedAlarmProp struct {
	value *EnumalarmManagerSuppressedAlarmProp
	isSet bool
}

func (v NullableEnumalarmManagerSuppressedAlarmProp) Get() *EnumalarmManagerSuppressedAlarmProp {
	return v.value
}

func (v *NullableEnumalarmManagerSuppressedAlarmProp) Set(val *EnumalarmManagerSuppressedAlarmProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumalarmManagerSuppressedAlarmProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumalarmManagerSuppressedAlarmProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumalarmManagerSuppressedAlarmProp(val *EnumalarmManagerSuppressedAlarmProp) *NullableEnumalarmManagerSuppressedAlarmProp {
	return &NullableEnumalarmManagerSuppressedAlarmProp{value: val, isSet: true}
}

func (v NullableEnumalarmManagerSuppressedAlarmProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumalarmManagerSuppressedAlarmProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

