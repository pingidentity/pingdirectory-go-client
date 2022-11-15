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

// EnumrecurringTaskTimestampFormatProp The format to use for the timestamp represented by the \"${timestamp}\" token in the filename pattern.
type EnumrecurringTaskTimestampFormatProp string

// List of Enumrecurring-task-timestampFormatProp
const (
	ENUMRECURRINGTASKTIMESTAMPFORMATPROP_GENERALIZED_TIME_UTC_WITH_MILLISECONDS EnumrecurringTaskTimestampFormatProp = "generalized-time-utc-with-milliseconds"
	ENUMRECURRINGTASKTIMESTAMPFORMATPROP_GENERALIZED_TIME_UTC_WITH_SECONDS EnumrecurringTaskTimestampFormatProp = "generalized-time-utc-with-seconds"
	ENUMRECURRINGTASKTIMESTAMPFORMATPROP_GENERALIZED_TIME_UTC_WITH_MINUTES EnumrecurringTaskTimestampFormatProp = "generalized-time-utc-with-minutes"
	ENUMRECURRINGTASKTIMESTAMPFORMATPROP_LOCAL_TIME_WITH_MILLISECONDS EnumrecurringTaskTimestampFormatProp = "local-time-with-milliseconds"
	ENUMRECURRINGTASKTIMESTAMPFORMATPROP_LOCAL_TIME_WITH_SECONDS EnumrecurringTaskTimestampFormatProp = "local-time-with-seconds"
	ENUMRECURRINGTASKTIMESTAMPFORMATPROP_LOCAL_TIME_WITH_MINUTES EnumrecurringTaskTimestampFormatProp = "local-time-with-minutes"
	ENUMRECURRINGTASKTIMESTAMPFORMATPROP_LOCAL_DATE EnumrecurringTaskTimestampFormatProp = "local-date"
)

// All allowed values of EnumrecurringTaskTimestampFormatProp enum
var AllowedEnumrecurringTaskTimestampFormatPropEnumValues = []EnumrecurringTaskTimestampFormatProp{
	"generalized-time-utc-with-milliseconds",
	"generalized-time-utc-with-seconds",
	"generalized-time-utc-with-minutes",
	"local-time-with-milliseconds",
	"local-time-with-seconds",
	"local-time-with-minutes",
	"local-date",
}

func (v *EnumrecurringTaskTimestampFormatProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumrecurringTaskTimestampFormatProp(value)
	for _, existing := range AllowedEnumrecurringTaskTimestampFormatPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumrecurringTaskTimestampFormatProp", value)
}

// NewEnumrecurringTaskTimestampFormatPropFromValue returns a pointer to a valid EnumrecurringTaskTimestampFormatProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumrecurringTaskTimestampFormatPropFromValue(v string) (*EnumrecurringTaskTimestampFormatProp, error) {
	ev := EnumrecurringTaskTimestampFormatProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumrecurringTaskTimestampFormatProp: valid values are %v", v, AllowedEnumrecurringTaskTimestampFormatPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumrecurringTaskTimestampFormatProp) IsValid() bool {
	for _, existing := range AllowedEnumrecurringTaskTimestampFormatPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumrecurring-task-timestampFormatProp value
func (v EnumrecurringTaskTimestampFormatProp) Ptr() *EnumrecurringTaskTimestampFormatProp {
	return &v
}

type NullableEnumrecurringTaskTimestampFormatProp struct {
	value *EnumrecurringTaskTimestampFormatProp
	isSet bool
}

func (v NullableEnumrecurringTaskTimestampFormatProp) Get() *EnumrecurringTaskTimestampFormatProp {
	return v.value
}

func (v *NullableEnumrecurringTaskTimestampFormatProp) Set(val *EnumrecurringTaskTimestampFormatProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumrecurringTaskTimestampFormatProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumrecurringTaskTimestampFormatProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumrecurringTaskTimestampFormatProp(val *EnumrecurringTaskTimestampFormatProp) *NullableEnumrecurringTaskTimestampFormatProp {
	return &NullableEnumrecurringTaskTimestampFormatProp{value: val, isSet: true}
}

func (v NullableEnumrecurringTaskTimestampFormatProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumrecurringTaskTimestampFormatProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

