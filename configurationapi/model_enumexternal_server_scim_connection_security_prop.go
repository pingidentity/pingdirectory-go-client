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

// EnumexternalServerScimConnectionSecurityProp The mechanism to use to secure communication with the SCIM service provider.
type EnumexternalServerScimConnectionSecurityProp string

// List of Enumexternal-server-scim-connectionSecurityProp
const (
	ENUMEXTERNALSERVERSCIMCONNECTIONSECURITYPROP_NONE EnumexternalServerScimConnectionSecurityProp = "none"
	ENUMEXTERNALSERVERSCIMCONNECTIONSECURITYPROP_SSL  EnumexternalServerScimConnectionSecurityProp = "ssl"
)

// All allowed values of EnumexternalServerScimConnectionSecurityProp enum
var AllowedEnumexternalServerScimConnectionSecurityPropEnumValues = []EnumexternalServerScimConnectionSecurityProp{
	"none",
	"ssl",
}

func (v *EnumexternalServerScimConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerScimConnectionSecurityProp(value)
	for _, existing := range AllowedEnumexternalServerScimConnectionSecurityPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerScimConnectionSecurityProp", value)
}

// NewEnumexternalServerScimConnectionSecurityPropFromValue returns a pointer to a valid EnumexternalServerScimConnectionSecurityProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerScimConnectionSecurityPropFromValue(v string) (*EnumexternalServerScimConnectionSecurityProp, error) {
	ev := EnumexternalServerScimConnectionSecurityProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerScimConnectionSecurityProp: valid values are %v", v, AllowedEnumexternalServerScimConnectionSecurityPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerScimConnectionSecurityProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerScimConnectionSecurityPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-scim-connectionSecurityProp value
func (v EnumexternalServerScimConnectionSecurityProp) Ptr() *EnumexternalServerScimConnectionSecurityProp {
	return &v
}

type NullableEnumexternalServerScimConnectionSecurityProp struct {
	value *EnumexternalServerScimConnectionSecurityProp
	isSet bool
}

func (v NullableEnumexternalServerScimConnectionSecurityProp) Get() *EnumexternalServerScimConnectionSecurityProp {
	return v.value
}

func (v *NullableEnumexternalServerScimConnectionSecurityProp) Set(val *EnumexternalServerScimConnectionSecurityProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerScimConnectionSecurityProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerScimConnectionSecurityProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerScimConnectionSecurityProp(val *EnumexternalServerScimConnectionSecurityProp) *NullableEnumexternalServerScimConnectionSecurityProp {
	return &NullableEnumexternalServerScimConnectionSecurityProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerScimConnectionSecurityProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerScimConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
