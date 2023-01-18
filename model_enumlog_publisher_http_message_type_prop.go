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

// EnumlogPublisherHttpMessageTypeProp Specifies the HTTP message types which can be logged.
type EnumlogPublisherHttpMessageTypeProp string

// List of Enumlog-publisher-httpMessageTypeProp
const (
	ENUMLOGPUBLISHERHTTPMESSAGETYPEPROP_REQUEST EnumlogPublisherHttpMessageTypeProp = "request"
	ENUMLOGPUBLISHERHTTPMESSAGETYPEPROP_RESPONSE EnumlogPublisherHttpMessageTypeProp = "response"
)

// All allowed values of EnumlogPublisherHttpMessageTypeProp enum
var AllowedEnumlogPublisherHttpMessageTypePropEnumValues = []EnumlogPublisherHttpMessageTypeProp{
	"request",
	"response",
}

func (v *EnumlogPublisherHttpMessageTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherHttpMessageTypeProp(value)
	for _, existing := range AllowedEnumlogPublisherHttpMessageTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherHttpMessageTypeProp", value)
}

// NewEnumlogPublisherHttpMessageTypePropFromValue returns a pointer to a valid EnumlogPublisherHttpMessageTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherHttpMessageTypePropFromValue(v string) (*EnumlogPublisherHttpMessageTypeProp, error) {
	ev := EnumlogPublisherHttpMessageTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherHttpMessageTypeProp: valid values are %v", v, AllowedEnumlogPublisherHttpMessageTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherHttpMessageTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherHttpMessageTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-httpMessageTypeProp value
func (v EnumlogPublisherHttpMessageTypeProp) Ptr() *EnumlogPublisherHttpMessageTypeProp {
	return &v
}

type NullableEnumlogPublisherHttpMessageTypeProp struct {
	value *EnumlogPublisherHttpMessageTypeProp
	isSet bool
}

func (v NullableEnumlogPublisherHttpMessageTypeProp) Get() *EnumlogPublisherHttpMessageTypeProp {
	return v.value
}

func (v *NullableEnumlogPublisherHttpMessageTypeProp) Set(val *EnumlogPublisherHttpMessageTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherHttpMessageTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherHttpMessageTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherHttpMessageTypeProp(val *EnumlogPublisherHttpMessageTypeProp) *NullableEnumlogPublisherHttpMessageTypeProp {
	return &NullableEnumlogPublisherHttpMessageTypeProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherHttpMessageTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherHttpMessageTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
