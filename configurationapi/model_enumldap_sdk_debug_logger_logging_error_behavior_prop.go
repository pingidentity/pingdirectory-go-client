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

// EnumldapSdkDebugLoggerLoggingErrorBehaviorProp Specifies the behavior that the server should exhibit if an error occurs during logging processing.
type EnumldapSdkDebugLoggerLoggingErrorBehaviorProp string

// List of Enumldap-sdk-debug-logger-loggingErrorBehaviorProp
const (
	ENUMLDAPSDKDEBUGLOGGERLOGGINGERRORBEHAVIORPROP_STANDARD_ERROR EnumldapSdkDebugLoggerLoggingErrorBehaviorProp = "standard-error"
	ENUMLDAPSDKDEBUGLOGGERLOGGINGERRORBEHAVIORPROP_LOCKDOWN_MODE  EnumldapSdkDebugLoggerLoggingErrorBehaviorProp = "lockdown-mode"
)

// All allowed values of EnumldapSdkDebugLoggerLoggingErrorBehaviorProp enum
var AllowedEnumldapSdkDebugLoggerLoggingErrorBehaviorPropEnumValues = []EnumldapSdkDebugLoggerLoggingErrorBehaviorProp{
	"standard-error",
	"lockdown-mode",
}

func (v *EnumldapSdkDebugLoggerLoggingErrorBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumldapSdkDebugLoggerLoggingErrorBehaviorProp(value)
	for _, existing := range AllowedEnumldapSdkDebugLoggerLoggingErrorBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumldapSdkDebugLoggerLoggingErrorBehaviorProp", value)
}

// NewEnumldapSdkDebugLoggerLoggingErrorBehaviorPropFromValue returns a pointer to a valid EnumldapSdkDebugLoggerLoggingErrorBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumldapSdkDebugLoggerLoggingErrorBehaviorPropFromValue(v string) (*EnumldapSdkDebugLoggerLoggingErrorBehaviorProp, error) {
	ev := EnumldapSdkDebugLoggerLoggingErrorBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumldapSdkDebugLoggerLoggingErrorBehaviorProp: valid values are %v", v, AllowedEnumldapSdkDebugLoggerLoggingErrorBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumldapSdkDebugLoggerLoggingErrorBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumldapSdkDebugLoggerLoggingErrorBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumldap-sdk-debug-logger-loggingErrorBehaviorProp value
func (v EnumldapSdkDebugLoggerLoggingErrorBehaviorProp) Ptr() *EnumldapSdkDebugLoggerLoggingErrorBehaviorProp {
	return &v
}

type NullableEnumldapSdkDebugLoggerLoggingErrorBehaviorProp struct {
	value *EnumldapSdkDebugLoggerLoggingErrorBehaviorProp
	isSet bool
}

func (v NullableEnumldapSdkDebugLoggerLoggingErrorBehaviorProp) Get() *EnumldapSdkDebugLoggerLoggingErrorBehaviorProp {
	return v.value
}

func (v *NullableEnumldapSdkDebugLoggerLoggingErrorBehaviorProp) Set(val *EnumldapSdkDebugLoggerLoggingErrorBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumldapSdkDebugLoggerLoggingErrorBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumldapSdkDebugLoggerLoggingErrorBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumldapSdkDebugLoggerLoggingErrorBehaviorProp(val *EnumldapSdkDebugLoggerLoggingErrorBehaviorProp) *NullableEnumldapSdkDebugLoggerLoggingErrorBehaviorProp {
	return &NullableEnumldapSdkDebugLoggerLoggingErrorBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumldapSdkDebugLoggerLoggingErrorBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumldapSdkDebugLoggerLoggingErrorBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
