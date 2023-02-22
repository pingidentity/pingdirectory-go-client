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

// EnumldapSdkDebugLoggerTimestampPrecisionProp Specifies the smallest time unit to be included in timestamps.
type EnumldapSdkDebugLoggerTimestampPrecisionProp string

// List of Enumldap-sdk-debug-logger-timestampPrecisionProp
const (
	ENUMLDAPSDKDEBUGLOGGERTIMESTAMPPRECISIONPROP_SECONDS      EnumldapSdkDebugLoggerTimestampPrecisionProp = "seconds"
	ENUMLDAPSDKDEBUGLOGGERTIMESTAMPPRECISIONPROP_MILLISECONDS EnumldapSdkDebugLoggerTimestampPrecisionProp = "milliseconds"
)

// All allowed values of EnumldapSdkDebugLoggerTimestampPrecisionProp enum
var AllowedEnumldapSdkDebugLoggerTimestampPrecisionPropEnumValues = []EnumldapSdkDebugLoggerTimestampPrecisionProp{
	"seconds",
	"milliseconds",
}

func (v *EnumldapSdkDebugLoggerTimestampPrecisionProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumldapSdkDebugLoggerTimestampPrecisionProp(value)
	for _, existing := range AllowedEnumldapSdkDebugLoggerTimestampPrecisionPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumldapSdkDebugLoggerTimestampPrecisionProp", value)
}

// NewEnumldapSdkDebugLoggerTimestampPrecisionPropFromValue returns a pointer to a valid EnumldapSdkDebugLoggerTimestampPrecisionProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumldapSdkDebugLoggerTimestampPrecisionPropFromValue(v string) (*EnumldapSdkDebugLoggerTimestampPrecisionProp, error) {
	ev := EnumldapSdkDebugLoggerTimestampPrecisionProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumldapSdkDebugLoggerTimestampPrecisionProp: valid values are %v", v, AllowedEnumldapSdkDebugLoggerTimestampPrecisionPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumldapSdkDebugLoggerTimestampPrecisionProp) IsValid() bool {
	for _, existing := range AllowedEnumldapSdkDebugLoggerTimestampPrecisionPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumldap-sdk-debug-logger-timestampPrecisionProp value
func (v EnumldapSdkDebugLoggerTimestampPrecisionProp) Ptr() *EnumldapSdkDebugLoggerTimestampPrecisionProp {
	return &v
}

type NullableEnumldapSdkDebugLoggerTimestampPrecisionProp struct {
	value *EnumldapSdkDebugLoggerTimestampPrecisionProp
	isSet bool
}

func (v NullableEnumldapSdkDebugLoggerTimestampPrecisionProp) Get() *EnumldapSdkDebugLoggerTimestampPrecisionProp {
	return v.value
}

func (v *NullableEnumldapSdkDebugLoggerTimestampPrecisionProp) Set(val *EnumldapSdkDebugLoggerTimestampPrecisionProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumldapSdkDebugLoggerTimestampPrecisionProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumldapSdkDebugLoggerTimestampPrecisionProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumldapSdkDebugLoggerTimestampPrecisionProp(val *EnumldapSdkDebugLoggerTimestampPrecisionProp) *NullableEnumldapSdkDebugLoggerTimestampPrecisionProp {
	return &NullableEnumldapSdkDebugLoggerTimestampPrecisionProp{value: val, isSet: true}
}

func (v NullableEnumldapSdkDebugLoggerTimestampPrecisionProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumldapSdkDebugLoggerTimestampPrecisionProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}