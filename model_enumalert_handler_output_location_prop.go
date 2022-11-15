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

// EnumalertHandlerOutputLocationProp The location to which alert messages will be written.
type EnumalertHandlerOutputLocationProp string

// List of Enumalert-handler-outputLocationProp
const (
	SERVER_OUT_FILE EnumalertHandlerOutputLocationProp = "server-out-file"
	STANDARD_OUTPUT EnumalertHandlerOutputLocationProp = "standard-output"
	STANDARD_ERROR EnumalertHandlerOutputLocationProp = "standard-error"
)

// All allowed values of EnumalertHandlerOutputLocationProp enum
var AllowedEnumalertHandlerOutputLocationPropEnumValues = []EnumalertHandlerOutputLocationProp{
	"server-out-file",
	"standard-output",
	"standard-error",
}

func (v *EnumalertHandlerOutputLocationProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumalertHandlerOutputLocationProp(value)
	for _, existing := range AllowedEnumalertHandlerOutputLocationPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumalertHandlerOutputLocationProp", value)
}

// NewEnumalertHandlerOutputLocationPropFromValue returns a pointer to a valid EnumalertHandlerOutputLocationProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumalertHandlerOutputLocationPropFromValue(v string) (*EnumalertHandlerOutputLocationProp, error) {
	ev := EnumalertHandlerOutputLocationProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumalertHandlerOutputLocationProp: valid values are %v", v, AllowedEnumalertHandlerOutputLocationPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumalertHandlerOutputLocationProp) IsValid() bool {
	for _, existing := range AllowedEnumalertHandlerOutputLocationPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumalert-handler-outputLocationProp value
func (v EnumalertHandlerOutputLocationProp) Ptr() *EnumalertHandlerOutputLocationProp {
	return &v
}

type NullableEnumalertHandlerOutputLocationProp struct {
	value *EnumalertHandlerOutputLocationProp
	isSet bool
}

func (v NullableEnumalertHandlerOutputLocationProp) Get() *EnumalertHandlerOutputLocationProp {
	return v.value
}

func (v *NullableEnumalertHandlerOutputLocationProp) Set(val *EnumalertHandlerOutputLocationProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumalertHandlerOutputLocationProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumalertHandlerOutputLocationProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumalertHandlerOutputLocationProp(val *EnumalertHandlerOutputLocationProp) *NullableEnumalertHandlerOutputLocationProp {
	return &NullableEnumalertHandlerOutputLocationProp{value: val, isSet: true}
}

func (v NullableEnumalertHandlerOutputLocationProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumalertHandlerOutputLocationProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

