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

// EnumhttpServletExtensionAllowedAuthenticationTypeProp The types of authentication that may be used to authenticate to the file servlet.
type EnumhttpServletExtensionAllowedAuthenticationTypeProp string

// List of Enumhttp-servlet-extension-allowedAuthenticationTypeProp
const (
	ENUMHTTPSERVLETEXTENSIONALLOWEDAUTHENTICATIONTYPEPROP_BASIC EnumhttpServletExtensionAllowedAuthenticationTypeProp = "basic"
	ENUMHTTPSERVLETEXTENSIONALLOWEDAUTHENTICATIONTYPEPROP_ACCESS_TOKEN EnumhttpServletExtensionAllowedAuthenticationTypeProp = "access-token"
	ENUMHTTPSERVLETEXTENSIONALLOWEDAUTHENTICATIONTYPEPROP_ID_TOKEN EnumhttpServletExtensionAllowedAuthenticationTypeProp = "id-token"
)

// All allowed values of EnumhttpServletExtensionAllowedAuthenticationTypeProp enum
var AllowedEnumhttpServletExtensionAllowedAuthenticationTypePropEnumValues = []EnumhttpServletExtensionAllowedAuthenticationTypeProp{
	"basic",
	"access-token",
	"id-token",
}

func (v *EnumhttpServletExtensionAllowedAuthenticationTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumhttpServletExtensionAllowedAuthenticationTypeProp(value)
	for _, existing := range AllowedEnumhttpServletExtensionAllowedAuthenticationTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumhttpServletExtensionAllowedAuthenticationTypeProp", value)
}

// NewEnumhttpServletExtensionAllowedAuthenticationTypePropFromValue returns a pointer to a valid EnumhttpServletExtensionAllowedAuthenticationTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumhttpServletExtensionAllowedAuthenticationTypePropFromValue(v string) (*EnumhttpServletExtensionAllowedAuthenticationTypeProp, error) {
	ev := EnumhttpServletExtensionAllowedAuthenticationTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumhttpServletExtensionAllowedAuthenticationTypeProp: valid values are %v", v, AllowedEnumhttpServletExtensionAllowedAuthenticationTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumhttpServletExtensionAllowedAuthenticationTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumhttpServletExtensionAllowedAuthenticationTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumhttp-servlet-extension-allowedAuthenticationTypeProp value
func (v EnumhttpServletExtensionAllowedAuthenticationTypeProp) Ptr() *EnumhttpServletExtensionAllowedAuthenticationTypeProp {
	return &v
}

type NullableEnumhttpServletExtensionAllowedAuthenticationTypeProp struct {
	value *EnumhttpServletExtensionAllowedAuthenticationTypeProp
	isSet bool
}

func (v NullableEnumhttpServletExtensionAllowedAuthenticationTypeProp) Get() *EnumhttpServletExtensionAllowedAuthenticationTypeProp {
	return v.value
}

func (v *NullableEnumhttpServletExtensionAllowedAuthenticationTypeProp) Set(val *EnumhttpServletExtensionAllowedAuthenticationTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumhttpServletExtensionAllowedAuthenticationTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumhttpServletExtensionAllowedAuthenticationTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumhttpServletExtensionAllowedAuthenticationTypeProp(val *EnumhttpServletExtensionAllowedAuthenticationTypeProp) *NullableEnumhttpServletExtensionAllowedAuthenticationTypeProp {
	return &NullableEnumhttpServletExtensionAllowedAuthenticationTypeProp{value: val, isSet: true}
}

func (v NullableEnumhttpServletExtensionAllowedAuthenticationTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumhttpServletExtensionAllowedAuthenticationTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

