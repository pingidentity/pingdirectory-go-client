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

// EnumexternalServerPingIdentityDsConnectionSecurityProp The mechanism to use to secure communication with the directory server.
type EnumexternalServerPingIdentityDsConnectionSecurityProp string

// List of Enumexternal-server-ping-identity-ds-connectionSecurityProp
const (
	ENUMEXTERNALSERVERPINGIDENTITYDSCONNECTIONSECURITYPROP_NONE     EnumexternalServerPingIdentityDsConnectionSecurityProp = "none"
	ENUMEXTERNALSERVERPINGIDENTITYDSCONNECTIONSECURITYPROP_SSL      EnumexternalServerPingIdentityDsConnectionSecurityProp = "ssl"
	ENUMEXTERNALSERVERPINGIDENTITYDSCONNECTIONSECURITYPROP_STARTTLS EnumexternalServerPingIdentityDsConnectionSecurityProp = "starttls"
)

// All allowed values of EnumexternalServerPingIdentityDsConnectionSecurityProp enum
var AllowedEnumexternalServerPingIdentityDsConnectionSecurityPropEnumValues = []EnumexternalServerPingIdentityDsConnectionSecurityProp{
	"none",
	"ssl",
	"starttls",
}

func (v *EnumexternalServerPingIdentityDsConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerPingIdentityDsConnectionSecurityProp(value)
	for _, existing := range AllowedEnumexternalServerPingIdentityDsConnectionSecurityPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerPingIdentityDsConnectionSecurityProp", value)
}

// NewEnumexternalServerPingIdentityDsConnectionSecurityPropFromValue returns a pointer to a valid EnumexternalServerPingIdentityDsConnectionSecurityProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerPingIdentityDsConnectionSecurityPropFromValue(v string) (*EnumexternalServerPingIdentityDsConnectionSecurityProp, error) {
	ev := EnumexternalServerPingIdentityDsConnectionSecurityProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerPingIdentityDsConnectionSecurityProp: valid values are %v", v, AllowedEnumexternalServerPingIdentityDsConnectionSecurityPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerPingIdentityDsConnectionSecurityProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerPingIdentityDsConnectionSecurityPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-ping-identity-ds-connectionSecurityProp value
func (v EnumexternalServerPingIdentityDsConnectionSecurityProp) Ptr() *EnumexternalServerPingIdentityDsConnectionSecurityProp {
	return &v
}

type NullableEnumexternalServerPingIdentityDsConnectionSecurityProp struct {
	value *EnumexternalServerPingIdentityDsConnectionSecurityProp
	isSet bool
}

func (v NullableEnumexternalServerPingIdentityDsConnectionSecurityProp) Get() *EnumexternalServerPingIdentityDsConnectionSecurityProp {
	return v.value
}

func (v *NullableEnumexternalServerPingIdentityDsConnectionSecurityProp) Set(val *EnumexternalServerPingIdentityDsConnectionSecurityProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerPingIdentityDsConnectionSecurityProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerPingIdentityDsConnectionSecurityProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerPingIdentityDsConnectionSecurityProp(val *EnumexternalServerPingIdentityDsConnectionSecurityProp) *NullableEnumexternalServerPingIdentityDsConnectionSecurityProp {
	return &NullableEnumexternalServerPingIdentityDsConnectionSecurityProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerPingIdentityDsConnectionSecurityProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerPingIdentityDsConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}