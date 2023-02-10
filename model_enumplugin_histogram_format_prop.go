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

// EnumpluginHistogramFormatProp The format of the data in the processing time histogram.
type EnumpluginHistogramFormatProp string

// List of Enumplugin-histogramFormatProp
const (
	ENUMPLUGINHISTOGRAMFORMATPROP_COUNT                EnumpluginHistogramFormatProp = "count"
	ENUMPLUGINHISTOGRAMFORMATPROP_AGGREGATE_PERCENTAGE EnumpluginHistogramFormatProp = "aggregate-percentage"
)

// All allowed values of EnumpluginHistogramFormatProp enum
var AllowedEnumpluginHistogramFormatPropEnumValues = []EnumpluginHistogramFormatProp{
	"count",
	"aggregate-percentage",
}

func (v *EnumpluginHistogramFormatProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginHistogramFormatProp(value)
	for _, existing := range AllowedEnumpluginHistogramFormatPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginHistogramFormatProp", value)
}

// NewEnumpluginHistogramFormatPropFromValue returns a pointer to a valid EnumpluginHistogramFormatProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginHistogramFormatPropFromValue(v string) (*EnumpluginHistogramFormatProp, error) {
	ev := EnumpluginHistogramFormatProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginHistogramFormatProp: valid values are %v", v, AllowedEnumpluginHistogramFormatPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginHistogramFormatProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginHistogramFormatPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-histogramFormatProp value
func (v EnumpluginHistogramFormatProp) Ptr() *EnumpluginHistogramFormatProp {
	return &v
}

type NullableEnumpluginHistogramFormatProp struct {
	value *EnumpluginHistogramFormatProp
	isSet bool
}

func (v NullableEnumpluginHistogramFormatProp) Get() *EnumpluginHistogramFormatProp {
	return v.value
}

func (v *NullableEnumpluginHistogramFormatProp) Set(val *EnumpluginHistogramFormatProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginHistogramFormatProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginHistogramFormatProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginHistogramFormatProp(val *EnumpluginHistogramFormatProp) *NullableEnumpluginHistogramFormatProp {
	return &NullableEnumpluginHistogramFormatProp{value: val, isSet: true}
}

func (v NullableEnumpluginHistogramFormatProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginHistogramFormatProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
