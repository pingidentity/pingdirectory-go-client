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

// EnumpluginLdapChangelogInfoProp Specifies the level of detail to include for the LDAP changelog.
type EnumpluginLdapChangelogInfoProp string

// List of Enumplugin-ldapChangelogInfoProp
const (
	ENUMPLUGINLDAPCHANGELOGINFOPROP_NONE     EnumpluginLdapChangelogInfoProp = "none"
	ENUMPLUGINLDAPCHANGELOGINFOPROP_BASIC    EnumpluginLdapChangelogInfoProp = "basic"
	ENUMPLUGINLDAPCHANGELOGINFOPROP_EXTENDED EnumpluginLdapChangelogInfoProp = "extended"
	ENUMPLUGINLDAPCHANGELOGINFOPROP_VERBOSE  EnumpluginLdapChangelogInfoProp = "verbose"
)

// All allowed values of EnumpluginLdapChangelogInfoProp enum
var AllowedEnumpluginLdapChangelogInfoPropEnumValues = []EnumpluginLdapChangelogInfoProp{
	"none",
	"basic",
	"extended",
	"verbose",
}

func (v *EnumpluginLdapChangelogInfoProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginLdapChangelogInfoProp(value)
	for _, existing := range AllowedEnumpluginLdapChangelogInfoPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginLdapChangelogInfoProp", value)
}

// NewEnumpluginLdapChangelogInfoPropFromValue returns a pointer to a valid EnumpluginLdapChangelogInfoProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginLdapChangelogInfoPropFromValue(v string) (*EnumpluginLdapChangelogInfoProp, error) {
	ev := EnumpluginLdapChangelogInfoProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginLdapChangelogInfoProp: valid values are %v", v, AllowedEnumpluginLdapChangelogInfoPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginLdapChangelogInfoProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginLdapChangelogInfoPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-ldapChangelogInfoProp value
func (v EnumpluginLdapChangelogInfoProp) Ptr() *EnumpluginLdapChangelogInfoProp {
	return &v
}

type NullableEnumpluginLdapChangelogInfoProp struct {
	value *EnumpluginLdapChangelogInfoProp
	isSet bool
}

func (v NullableEnumpluginLdapChangelogInfoProp) Get() *EnumpluginLdapChangelogInfoProp {
	return v.value
}

func (v *NullableEnumpluginLdapChangelogInfoProp) Set(val *EnumpluginLdapChangelogInfoProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginLdapChangelogInfoProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginLdapChangelogInfoProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginLdapChangelogInfoProp(val *EnumpluginLdapChangelogInfoProp) *NullableEnumpluginLdapChangelogInfoProp {
	return &NullableEnumpluginLdapChangelogInfoProp{value: val, isSet: true}
}

func (v NullableEnumpluginLdapChangelogInfoProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginLdapChangelogInfoProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
