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

// EnumldapSdkDebugLoggerDebugLevelProp The minimum debug level that should be used for messages to be logged.
type EnumldapSdkDebugLoggerDebugLevelProp string

// List of Enumldap-sdk-debug-logger-debugLevelProp
const (
	ENUMLDAPSDKDEBUGLOGGERDEBUGLEVELPROP_SEVERE EnumldapSdkDebugLoggerDebugLevelProp = "severe"
	ENUMLDAPSDKDEBUGLOGGERDEBUGLEVELPROP_WARNING EnumldapSdkDebugLoggerDebugLevelProp = "warning"
	ENUMLDAPSDKDEBUGLOGGERDEBUGLEVELPROP_INFO EnumldapSdkDebugLoggerDebugLevelProp = "info"
	ENUMLDAPSDKDEBUGLOGGERDEBUGLEVELPROP_CONFIG EnumldapSdkDebugLoggerDebugLevelProp = "config"
	ENUMLDAPSDKDEBUGLOGGERDEBUGLEVELPROP_FINE EnumldapSdkDebugLoggerDebugLevelProp = "fine"
	ENUMLDAPSDKDEBUGLOGGERDEBUGLEVELPROP_FINER EnumldapSdkDebugLoggerDebugLevelProp = "finer"
	ENUMLDAPSDKDEBUGLOGGERDEBUGLEVELPROP_FINEST EnumldapSdkDebugLoggerDebugLevelProp = "finest"
)

// All allowed values of EnumldapSdkDebugLoggerDebugLevelProp enum
var AllowedEnumldapSdkDebugLoggerDebugLevelPropEnumValues = []EnumldapSdkDebugLoggerDebugLevelProp{
	"severe",
	"warning",
	"info",
	"config",
	"fine",
	"finer",
	"finest",
}

func (v *EnumldapSdkDebugLoggerDebugLevelProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumldapSdkDebugLoggerDebugLevelProp(value)
	for _, existing := range AllowedEnumldapSdkDebugLoggerDebugLevelPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumldapSdkDebugLoggerDebugLevelProp", value)
}

// NewEnumldapSdkDebugLoggerDebugLevelPropFromValue returns a pointer to a valid EnumldapSdkDebugLoggerDebugLevelProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumldapSdkDebugLoggerDebugLevelPropFromValue(v string) (*EnumldapSdkDebugLoggerDebugLevelProp, error) {
	ev := EnumldapSdkDebugLoggerDebugLevelProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumldapSdkDebugLoggerDebugLevelProp: valid values are %v", v, AllowedEnumldapSdkDebugLoggerDebugLevelPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumldapSdkDebugLoggerDebugLevelProp) IsValid() bool {
	for _, existing := range AllowedEnumldapSdkDebugLoggerDebugLevelPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumldap-sdk-debug-logger-debugLevelProp value
func (v EnumldapSdkDebugLoggerDebugLevelProp) Ptr() *EnumldapSdkDebugLoggerDebugLevelProp {
	return &v
}

type NullableEnumldapSdkDebugLoggerDebugLevelProp struct {
	value *EnumldapSdkDebugLoggerDebugLevelProp
	isSet bool
}

func (v NullableEnumldapSdkDebugLoggerDebugLevelProp) Get() *EnumldapSdkDebugLoggerDebugLevelProp {
	return v.value
}

func (v *NullableEnumldapSdkDebugLoggerDebugLevelProp) Set(val *EnumldapSdkDebugLoggerDebugLevelProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumldapSdkDebugLoggerDebugLevelProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumldapSdkDebugLoggerDebugLevelProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumldapSdkDebugLoggerDebugLevelProp(val *EnumldapSdkDebugLoggerDebugLevelProp) *NullableEnumldapSdkDebugLoggerDebugLevelProp {
	return &NullableEnumldapSdkDebugLoggerDebugLevelProp{value: val, isSet: true}
}

func (v NullableEnumldapSdkDebugLoggerDebugLevelProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumldapSdkDebugLoggerDebugLevelProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
