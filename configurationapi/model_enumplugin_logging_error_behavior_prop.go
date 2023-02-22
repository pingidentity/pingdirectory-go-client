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

// EnumpluginLoggingErrorBehaviorProp Specifies the behavior that the server should exhibit if an error occurs during logging processing.
type EnumpluginLoggingErrorBehaviorProp string

// List of Enumplugin-loggingErrorBehaviorProp
const (
	ENUMPLUGINLOGGINGERRORBEHAVIORPROP_STANDARD_ERROR EnumpluginLoggingErrorBehaviorProp = "standard-error"
	ENUMPLUGINLOGGINGERRORBEHAVIORPROP_LOCKDOWN_MODE  EnumpluginLoggingErrorBehaviorProp = "lockdown-mode"
)

// All allowed values of EnumpluginLoggingErrorBehaviorProp enum
var AllowedEnumpluginLoggingErrorBehaviorPropEnumValues = []EnumpluginLoggingErrorBehaviorProp{
	"standard-error",
	"lockdown-mode",
}

func (v *EnumpluginLoggingErrorBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginLoggingErrorBehaviorProp(value)
	for _, existing := range AllowedEnumpluginLoggingErrorBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginLoggingErrorBehaviorProp", value)
}

// NewEnumpluginLoggingErrorBehaviorPropFromValue returns a pointer to a valid EnumpluginLoggingErrorBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginLoggingErrorBehaviorPropFromValue(v string) (*EnumpluginLoggingErrorBehaviorProp, error) {
	ev := EnumpluginLoggingErrorBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginLoggingErrorBehaviorProp: valid values are %v", v, AllowedEnumpluginLoggingErrorBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginLoggingErrorBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginLoggingErrorBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-loggingErrorBehaviorProp value
func (v EnumpluginLoggingErrorBehaviorProp) Ptr() *EnumpluginLoggingErrorBehaviorProp {
	return &v
}

type NullableEnumpluginLoggingErrorBehaviorProp struct {
	value *EnumpluginLoggingErrorBehaviorProp
	isSet bool
}

func (v NullableEnumpluginLoggingErrorBehaviorProp) Get() *EnumpluginLoggingErrorBehaviorProp {
	return v.value
}

func (v *NullableEnumpluginLoggingErrorBehaviorProp) Set(val *EnumpluginLoggingErrorBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginLoggingErrorBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginLoggingErrorBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginLoggingErrorBehaviorProp(val *EnumpluginLoggingErrorBehaviorProp) *NullableEnumpluginLoggingErrorBehaviorProp {
	return &NullableEnumpluginLoggingErrorBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumpluginLoggingErrorBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginLoggingErrorBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}