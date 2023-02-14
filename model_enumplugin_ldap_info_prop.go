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

// EnumpluginLdapInfoProp Specifies the level of detail to include about the LDAP connection handlers.
type EnumpluginLdapInfoProp string

// List of Enumplugin-ldapInfoProp
const (
	ENUMPLUGINLDAPINFOPROP_NONE     EnumpluginLdapInfoProp = "none"
	ENUMPLUGINLDAPINFOPROP_BASIC    EnumpluginLdapInfoProp = "basic"
	ENUMPLUGINLDAPINFOPROP_EXTENDED EnumpluginLdapInfoProp = "extended"
)

// All allowed values of EnumpluginLdapInfoProp enum
var AllowedEnumpluginLdapInfoPropEnumValues = []EnumpluginLdapInfoProp{
	"none",
	"basic",
	"extended",
}

func (v *EnumpluginLdapInfoProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginLdapInfoProp(value)
	for _, existing := range AllowedEnumpluginLdapInfoPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginLdapInfoProp", value)
}

// NewEnumpluginLdapInfoPropFromValue returns a pointer to a valid EnumpluginLdapInfoProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginLdapInfoPropFromValue(v string) (*EnumpluginLdapInfoProp, error) {
	ev := EnumpluginLdapInfoProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginLdapInfoProp: valid values are %v", v, AllowedEnumpluginLdapInfoPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginLdapInfoProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginLdapInfoPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-ldapInfoProp value
func (v EnumpluginLdapInfoProp) Ptr() *EnumpluginLdapInfoProp {
	return &v
}

type NullableEnumpluginLdapInfoProp struct {
	value *EnumpluginLdapInfoProp
	isSet bool
}

func (v NullableEnumpluginLdapInfoProp) Get() *EnumpluginLdapInfoProp {
	return v.value
}

func (v *NullableEnumpluginLdapInfoProp) Set(val *EnumpluginLdapInfoProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginLdapInfoProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginLdapInfoProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginLdapInfoProp(val *EnumpluginLdapInfoProp) *NullableEnumpluginLdapInfoProp {
	return &NullableEnumpluginLdapInfoProp{value: val, isSet: true}
}

func (v NullableEnumpluginLdapInfoProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginLdapInfoProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
