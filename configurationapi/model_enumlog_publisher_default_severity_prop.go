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

// EnumlogPublisherDefaultSeverityProp Specifies the default severity levels for the logger.
type EnumlogPublisherDefaultSeverityProp string

// List of Enumlog-publisher-defaultSeverityProp
const (
	ENUMLOGPUBLISHERDEFAULTSEVERITYPROP_ALL            EnumlogPublisherDefaultSeverityProp = "all"
	ENUMLOGPUBLISHERDEFAULTSEVERITYPROP_NONE           EnumlogPublisherDefaultSeverityProp = "none"
	ENUMLOGPUBLISHERDEFAULTSEVERITYPROP_FATAL_ERROR    EnumlogPublisherDefaultSeverityProp = "fatal-error"
	ENUMLOGPUBLISHERDEFAULTSEVERITYPROP_INFO           EnumlogPublisherDefaultSeverityProp = "info"
	ENUMLOGPUBLISHERDEFAULTSEVERITYPROP_MILD_ERROR     EnumlogPublisherDefaultSeverityProp = "mild-error"
	ENUMLOGPUBLISHERDEFAULTSEVERITYPROP_MILD_WARNING   EnumlogPublisherDefaultSeverityProp = "mild-warning"
	ENUMLOGPUBLISHERDEFAULTSEVERITYPROP_NOTICE         EnumlogPublisherDefaultSeverityProp = "notice"
	ENUMLOGPUBLISHERDEFAULTSEVERITYPROP_SEVERE_ERROR   EnumlogPublisherDefaultSeverityProp = "severe-error"
	ENUMLOGPUBLISHERDEFAULTSEVERITYPROP_SEVERE_WARNING EnumlogPublisherDefaultSeverityProp = "severe-warning"
	ENUMLOGPUBLISHERDEFAULTSEVERITYPROP_DEBUG          EnumlogPublisherDefaultSeverityProp = "debug"
)

// All allowed values of EnumlogPublisherDefaultSeverityProp enum
var AllowedEnumlogPublisherDefaultSeverityPropEnumValues = []EnumlogPublisherDefaultSeverityProp{
	"all",
	"none",
	"fatal-error",
	"info",
	"mild-error",
	"mild-warning",
	"notice",
	"severe-error",
	"severe-warning",
	"debug",
}

func (v *EnumlogPublisherDefaultSeverityProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherDefaultSeverityProp(value)
	for _, existing := range AllowedEnumlogPublisherDefaultSeverityPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherDefaultSeverityProp", value)
}

// NewEnumlogPublisherDefaultSeverityPropFromValue returns a pointer to a valid EnumlogPublisherDefaultSeverityProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherDefaultSeverityPropFromValue(v string) (*EnumlogPublisherDefaultSeverityProp, error) {
	ev := EnumlogPublisherDefaultSeverityProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherDefaultSeverityProp: valid values are %v", v, AllowedEnumlogPublisherDefaultSeverityPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherDefaultSeverityProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherDefaultSeverityPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-defaultSeverityProp value
func (v EnumlogPublisherDefaultSeverityProp) Ptr() *EnumlogPublisherDefaultSeverityProp {
	return &v
}

type NullableEnumlogPublisherDefaultSeverityProp struct {
	value *EnumlogPublisherDefaultSeverityProp
	isSet bool
}

func (v NullableEnumlogPublisherDefaultSeverityProp) Get() *EnumlogPublisherDefaultSeverityProp {
	return v.value
}

func (v *NullableEnumlogPublisherDefaultSeverityProp) Set(val *EnumlogPublisherDefaultSeverityProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherDefaultSeverityProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherDefaultSeverityProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherDefaultSeverityProp(val *EnumlogPublisherDefaultSeverityProp) *NullableEnumlogPublisherDefaultSeverityProp {
	return &NullableEnumlogPublisherDefaultSeverityProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherDefaultSeverityProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherDefaultSeverityProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
