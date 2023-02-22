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

// EnumlogPublisherOutputLocationProp Specifies the output stream to which JSON-formatted error log messages should be written.
type EnumlogPublisherOutputLocationProp string

// List of Enumlog-publisher-outputLocationProp
const (
	ENUMLOGPUBLISHEROUTPUTLOCATIONPROP_OUTPUT EnumlogPublisherOutputLocationProp = "standard-output"
	ENUMLOGPUBLISHEROUTPUTLOCATIONPROP_ERROR  EnumlogPublisherOutputLocationProp = "standard-error"
)

// All allowed values of EnumlogPublisherOutputLocationProp enum
var AllowedEnumlogPublisherOutputLocationPropEnumValues = []EnumlogPublisherOutputLocationProp{
	"standard-output",
	"standard-error",
}

func (v *EnumlogPublisherOutputLocationProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherOutputLocationProp(value)
	for _, existing := range AllowedEnumlogPublisherOutputLocationPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherOutputLocationProp", value)
}

// NewEnumlogPublisherOutputLocationPropFromValue returns a pointer to a valid EnumlogPublisherOutputLocationProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherOutputLocationPropFromValue(v string) (*EnumlogPublisherOutputLocationProp, error) {
	ev := EnumlogPublisherOutputLocationProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherOutputLocationProp: valid values are %v", v, AllowedEnumlogPublisherOutputLocationPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherOutputLocationProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherOutputLocationPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-outputLocationProp value
func (v EnumlogPublisherOutputLocationProp) Ptr() *EnumlogPublisherOutputLocationProp {
	return &v
}

type NullableEnumlogPublisherOutputLocationProp struct {
	value *EnumlogPublisherOutputLocationProp
	isSet bool
}

func (v NullableEnumlogPublisherOutputLocationProp) Get() *EnumlogPublisherOutputLocationProp {
	return v.value
}

func (v *NullableEnumlogPublisherOutputLocationProp) Set(val *EnumlogPublisherOutputLocationProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherOutputLocationProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherOutputLocationProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherOutputLocationProp(val *EnumlogPublisherOutputLocationProp) *NullableEnumlogPublisherOutputLocationProp {
	return &NullableEnumlogPublisherOutputLocationProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherOutputLocationProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherOutputLocationProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}