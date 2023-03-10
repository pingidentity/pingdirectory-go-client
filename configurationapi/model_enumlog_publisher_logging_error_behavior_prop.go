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

// EnumlogPublisherLoggingErrorBehaviorProp Specifies the behavior that the server should exhibit if an error occurs during logging processing.
type EnumlogPublisherLoggingErrorBehaviorProp string

// List of Enumlog-publisher-loggingErrorBehaviorProp
const (
	ENUMLOGPUBLISHERLOGGINGERRORBEHAVIORPROP_STANDARD_ERROR EnumlogPublisherLoggingErrorBehaviorProp = "standard-error"
	ENUMLOGPUBLISHERLOGGINGERRORBEHAVIORPROP_LOCKDOWN_MODE  EnumlogPublisherLoggingErrorBehaviorProp = "lockdown-mode"
)

// All allowed values of EnumlogPublisherLoggingErrorBehaviorProp enum
var AllowedEnumlogPublisherLoggingErrorBehaviorPropEnumValues = []EnumlogPublisherLoggingErrorBehaviorProp{
	"standard-error",
	"lockdown-mode",
}

func (v *EnumlogPublisherLoggingErrorBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherLoggingErrorBehaviorProp(value)
	for _, existing := range AllowedEnumlogPublisherLoggingErrorBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherLoggingErrorBehaviorProp", value)
}

// NewEnumlogPublisherLoggingErrorBehaviorPropFromValue returns a pointer to a valid EnumlogPublisherLoggingErrorBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherLoggingErrorBehaviorPropFromValue(v string) (*EnumlogPublisherLoggingErrorBehaviorProp, error) {
	ev := EnumlogPublisherLoggingErrorBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherLoggingErrorBehaviorProp: valid values are %v", v, AllowedEnumlogPublisherLoggingErrorBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherLoggingErrorBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherLoggingErrorBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-loggingErrorBehaviorProp value
func (v EnumlogPublisherLoggingErrorBehaviorProp) Ptr() *EnumlogPublisherLoggingErrorBehaviorProp {
	return &v
}

type NullableEnumlogPublisherLoggingErrorBehaviorProp struct {
	value *EnumlogPublisherLoggingErrorBehaviorProp
	isSet bool
}

func (v NullableEnumlogPublisherLoggingErrorBehaviorProp) Get() *EnumlogPublisherLoggingErrorBehaviorProp {
	return v.value
}

func (v *NullableEnumlogPublisherLoggingErrorBehaviorProp) Set(val *EnumlogPublisherLoggingErrorBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherLoggingErrorBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherLoggingErrorBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherLoggingErrorBehaviorProp(val *EnumlogPublisherLoggingErrorBehaviorProp) *NullableEnumlogPublisherLoggingErrorBehaviorProp {
	return &NullableEnumlogPublisherLoggingErrorBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherLoggingErrorBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherLoggingErrorBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
