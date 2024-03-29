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

// EnumexternalServerNokiaDsConnectionSecurityProp The mechanism to use to secure communication with the directory server.
type EnumexternalServerNokiaDsConnectionSecurityProp string

// List of Enumexternal-server-nokia-ds-connectionSecurityProp
const (
	ENUMEXTERNALSERVERNOKIADSCONNECTIONSECURITYPROP_NONE     EnumexternalServerNokiaDsConnectionSecurityProp = "none"
	ENUMEXTERNALSERVERNOKIADSCONNECTIONSECURITYPROP_SSL      EnumexternalServerNokiaDsConnectionSecurityProp = "ssl"
	ENUMEXTERNALSERVERNOKIADSCONNECTIONSECURITYPROP_STARTTLS EnumexternalServerNokiaDsConnectionSecurityProp = "starttls"
)

// All allowed values of EnumexternalServerNokiaDsConnectionSecurityProp enum
var AllowedEnumexternalServerNokiaDsConnectionSecurityPropEnumValues = []EnumexternalServerNokiaDsConnectionSecurityProp{
	"none",
	"ssl",
	"starttls",
}

func (v *EnumexternalServerNokiaDsConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerNokiaDsConnectionSecurityProp(value)
	for _, existing := range AllowedEnumexternalServerNokiaDsConnectionSecurityPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerNokiaDsConnectionSecurityProp", value)
}

// NewEnumexternalServerNokiaDsConnectionSecurityPropFromValue returns a pointer to a valid EnumexternalServerNokiaDsConnectionSecurityProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerNokiaDsConnectionSecurityPropFromValue(v string) (*EnumexternalServerNokiaDsConnectionSecurityProp, error) {
	ev := EnumexternalServerNokiaDsConnectionSecurityProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerNokiaDsConnectionSecurityProp: valid values are %v", v, AllowedEnumexternalServerNokiaDsConnectionSecurityPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerNokiaDsConnectionSecurityProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerNokiaDsConnectionSecurityPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-nokia-ds-connectionSecurityProp value
func (v EnumexternalServerNokiaDsConnectionSecurityProp) Ptr() *EnumexternalServerNokiaDsConnectionSecurityProp {
	return &v
}

type NullableEnumexternalServerNokiaDsConnectionSecurityProp struct {
	value *EnumexternalServerNokiaDsConnectionSecurityProp
	isSet bool
}

func (v NullableEnumexternalServerNokiaDsConnectionSecurityProp) Get() *EnumexternalServerNokiaDsConnectionSecurityProp {
	return v.value
}

func (v *NullableEnumexternalServerNokiaDsConnectionSecurityProp) Set(val *EnumexternalServerNokiaDsConnectionSecurityProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerNokiaDsConnectionSecurityProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerNokiaDsConnectionSecurityProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerNokiaDsConnectionSecurityProp(val *EnumexternalServerNokiaDsConnectionSecurityProp) *NullableEnumexternalServerNokiaDsConnectionSecurityProp {
	return &NullableEnumexternalServerNokiaDsConnectionSecurityProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerNokiaDsConnectionSecurityProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerNokiaDsConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
