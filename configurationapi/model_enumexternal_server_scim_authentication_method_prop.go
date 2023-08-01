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

// EnumexternalServerScimAuthenticationMethodProp The mechanism to use to authenticate to the target server.
type EnumexternalServerScimAuthenticationMethodProp string

// List of Enumexternal-server-scim-authenticationMethodProp
const (
	ENUMEXTERNALSERVERSCIMAUTHENTICATIONMETHODPROP_NONE  EnumexternalServerScimAuthenticationMethodProp = "none"
	ENUMEXTERNALSERVERSCIMAUTHENTICATIONMETHODPROP_BASIC EnumexternalServerScimAuthenticationMethodProp = "basic"
	ENUMEXTERNALSERVERSCIMAUTHENTICATIONMETHODPROP_OAUTH EnumexternalServerScimAuthenticationMethodProp = "oauth"
)

// All allowed values of EnumexternalServerScimAuthenticationMethodProp enum
var AllowedEnumexternalServerScimAuthenticationMethodPropEnumValues = []EnumexternalServerScimAuthenticationMethodProp{
	"none",
	"basic",
	"oauth",
}

func (v *EnumexternalServerScimAuthenticationMethodProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerScimAuthenticationMethodProp(value)
	for _, existing := range AllowedEnumexternalServerScimAuthenticationMethodPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerScimAuthenticationMethodProp", value)
}

// NewEnumexternalServerScimAuthenticationMethodPropFromValue returns a pointer to a valid EnumexternalServerScimAuthenticationMethodProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerScimAuthenticationMethodPropFromValue(v string) (*EnumexternalServerScimAuthenticationMethodProp, error) {
	ev := EnumexternalServerScimAuthenticationMethodProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerScimAuthenticationMethodProp: valid values are %v", v, AllowedEnumexternalServerScimAuthenticationMethodPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerScimAuthenticationMethodProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerScimAuthenticationMethodPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-scim-authenticationMethodProp value
func (v EnumexternalServerScimAuthenticationMethodProp) Ptr() *EnumexternalServerScimAuthenticationMethodProp {
	return &v
}

type NullableEnumexternalServerScimAuthenticationMethodProp struct {
	value *EnumexternalServerScimAuthenticationMethodProp
	isSet bool
}

func (v NullableEnumexternalServerScimAuthenticationMethodProp) Get() *EnumexternalServerScimAuthenticationMethodProp {
	return v.value
}

func (v *NullableEnumexternalServerScimAuthenticationMethodProp) Set(val *EnumexternalServerScimAuthenticationMethodProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerScimAuthenticationMethodProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerScimAuthenticationMethodProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerScimAuthenticationMethodProp(val *EnumexternalServerScimAuthenticationMethodProp) *NullableEnumexternalServerScimAuthenticationMethodProp {
	return &NullableEnumexternalServerScimAuthenticationMethodProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerScimAuthenticationMethodProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerScimAuthenticationMethodProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
