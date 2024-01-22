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

// EnumexternalServerLdapConnectionSecurityProp The mechanism to use to secure communication with the directory server.
type EnumexternalServerLdapConnectionSecurityProp string

// List of Enumexternal-server-ldap-connectionSecurityProp
const (
	ENUMEXTERNALSERVERLDAPCONNECTIONSECURITYPROP_NONE     EnumexternalServerLdapConnectionSecurityProp = "none"
	ENUMEXTERNALSERVERLDAPCONNECTIONSECURITYPROP_SSL      EnumexternalServerLdapConnectionSecurityProp = "ssl"
	ENUMEXTERNALSERVERLDAPCONNECTIONSECURITYPROP_STARTTLS EnumexternalServerLdapConnectionSecurityProp = "starttls"
)

// All allowed values of EnumexternalServerLdapConnectionSecurityProp enum
var AllowedEnumexternalServerLdapConnectionSecurityPropEnumValues = []EnumexternalServerLdapConnectionSecurityProp{
	"none",
	"ssl",
	"starttls",
}

func (v *EnumexternalServerLdapConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerLdapConnectionSecurityProp(value)
	for _, existing := range AllowedEnumexternalServerLdapConnectionSecurityPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerLdapConnectionSecurityProp", value)
}

// NewEnumexternalServerLdapConnectionSecurityPropFromValue returns a pointer to a valid EnumexternalServerLdapConnectionSecurityProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerLdapConnectionSecurityPropFromValue(v string) (*EnumexternalServerLdapConnectionSecurityProp, error) {
	ev := EnumexternalServerLdapConnectionSecurityProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerLdapConnectionSecurityProp: valid values are %v", v, AllowedEnumexternalServerLdapConnectionSecurityPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerLdapConnectionSecurityProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerLdapConnectionSecurityPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-ldap-connectionSecurityProp value
func (v EnumexternalServerLdapConnectionSecurityProp) Ptr() *EnumexternalServerLdapConnectionSecurityProp {
	return &v
}

type NullableEnumexternalServerLdapConnectionSecurityProp struct {
	value *EnumexternalServerLdapConnectionSecurityProp
	isSet bool
}

func (v NullableEnumexternalServerLdapConnectionSecurityProp) Get() *EnumexternalServerLdapConnectionSecurityProp {
	return v.value
}

func (v *NullableEnumexternalServerLdapConnectionSecurityProp) Set(val *EnumexternalServerLdapConnectionSecurityProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerLdapConnectionSecurityProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerLdapConnectionSecurityProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerLdapConnectionSecurityProp(val *EnumexternalServerLdapConnectionSecurityProp) *NullableEnumexternalServerLdapConnectionSecurityProp {
	return &NullableEnumexternalServerLdapConnectionSecurityProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerLdapConnectionSecurityProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerLdapConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}