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

// EnumlogPublisherTimestampPrecisionProp Specifies the smallest time unit to be included in timestamps.
type EnumlogPublisherTimestampPrecisionProp string

// List of Enumlog-publisher-timestampPrecisionProp
const (
	ENUMLOGPUBLISHERTIMESTAMPPRECISIONPROP_SECONDS EnumlogPublisherTimestampPrecisionProp = "seconds"
	ENUMLOGPUBLISHERTIMESTAMPPRECISIONPROP_MILLISECONDS EnumlogPublisherTimestampPrecisionProp = "milliseconds"
)

// All allowed values of EnumlogPublisherTimestampPrecisionProp enum
var AllowedEnumlogPublisherTimestampPrecisionPropEnumValues = []EnumlogPublisherTimestampPrecisionProp{
	"seconds",
	"milliseconds",
}

func (v *EnumlogPublisherTimestampPrecisionProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherTimestampPrecisionProp(value)
	for _, existing := range AllowedEnumlogPublisherTimestampPrecisionPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherTimestampPrecisionProp", value)
}

// NewEnumlogPublisherTimestampPrecisionPropFromValue returns a pointer to a valid EnumlogPublisherTimestampPrecisionProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherTimestampPrecisionPropFromValue(v string) (*EnumlogPublisherTimestampPrecisionProp, error) {
	ev := EnumlogPublisherTimestampPrecisionProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherTimestampPrecisionProp: valid values are %v", v, AllowedEnumlogPublisherTimestampPrecisionPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherTimestampPrecisionProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherTimestampPrecisionPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-timestampPrecisionProp value
func (v EnumlogPublisherTimestampPrecisionProp) Ptr() *EnumlogPublisherTimestampPrecisionProp {
	return &v
}

type NullableEnumlogPublisherTimestampPrecisionProp struct {
	value *EnumlogPublisherTimestampPrecisionProp
	isSet bool
}

func (v NullableEnumlogPublisherTimestampPrecisionProp) Get() *EnumlogPublisherTimestampPrecisionProp {
	return v.value
}

func (v *NullableEnumlogPublisherTimestampPrecisionProp) Set(val *EnumlogPublisherTimestampPrecisionProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherTimestampPrecisionProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherTimestampPrecisionProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherTimestampPrecisionProp(val *EnumlogPublisherTimestampPrecisionProp) *NullableEnumlogPublisherTimestampPrecisionProp {
	return &NullableEnumlogPublisherTimestampPrecisionProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherTimestampPrecisionProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherTimestampPrecisionProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

