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

// EnumexternalServerConnectionSecurityProp The mechanism to use to secure communication with the directory server.
type EnumexternalServerConnectionSecurityProp string

// List of Enumexternal-server-connectionSecurityProp
const (
	ENUMEXTERNALSERVERCONNECTIONSECURITYPROP_NONE     EnumexternalServerConnectionSecurityProp = "none"
	ENUMEXTERNALSERVERCONNECTIONSECURITYPROP_SSL      EnumexternalServerConnectionSecurityProp = "ssl"
	ENUMEXTERNALSERVERCONNECTIONSECURITYPROP_STARTTLS EnumexternalServerConnectionSecurityProp = "starttls"
)

// All allowed values of EnumexternalServerConnectionSecurityProp enum
var AllowedEnumexternalServerConnectionSecurityPropEnumValues = []EnumexternalServerConnectionSecurityProp{
	"none",
	"ssl",
	"starttls",
}

func (v *EnumexternalServerConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerConnectionSecurityProp(value)
	for _, existing := range AllowedEnumexternalServerConnectionSecurityPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerConnectionSecurityProp", value)
}

// NewEnumexternalServerConnectionSecurityPropFromValue returns a pointer to a valid EnumexternalServerConnectionSecurityProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerConnectionSecurityPropFromValue(v string) (*EnumexternalServerConnectionSecurityProp, error) {
	ev := EnumexternalServerConnectionSecurityProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerConnectionSecurityProp: valid values are %v", v, AllowedEnumexternalServerConnectionSecurityPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerConnectionSecurityProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerConnectionSecurityPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-connectionSecurityProp value
func (v EnumexternalServerConnectionSecurityProp) Ptr() *EnumexternalServerConnectionSecurityProp {
	return &v
}

type NullableEnumexternalServerConnectionSecurityProp struct {
	value *EnumexternalServerConnectionSecurityProp
	isSet bool
}

func (v NullableEnumexternalServerConnectionSecurityProp) Get() *EnumexternalServerConnectionSecurityProp {
	return v.value
}

func (v *NullableEnumexternalServerConnectionSecurityProp) Set(val *EnumexternalServerConnectionSecurityProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerConnectionSecurityProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerConnectionSecurityProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerConnectionSecurityProp(val *EnumexternalServerConnectionSecurityProp) *NullableEnumexternalServerConnectionSecurityProp {
	return &NullableEnumexternalServerConnectionSecurityProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerConnectionSecurityProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
