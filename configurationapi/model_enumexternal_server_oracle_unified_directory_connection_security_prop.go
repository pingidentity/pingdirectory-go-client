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

// EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp The mechanism to use to secure communication with the directory server.
type EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp string

// List of Enumexternal-server-oracle-unified-directory-connectionSecurityProp
const (
	ENUMEXTERNALSERVERORACLEUNIFIEDDIRECTORYCONNECTIONSECURITYPROP_NONE     EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp = "none"
	ENUMEXTERNALSERVERORACLEUNIFIEDDIRECTORYCONNECTIONSECURITYPROP_SSL      EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp = "ssl"
	ENUMEXTERNALSERVERORACLEUNIFIEDDIRECTORYCONNECTIONSECURITYPROP_STARTTLS EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp = "starttls"
)

// All allowed values of EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp enum
var AllowedEnumexternalServerOracleUnifiedDirectoryConnectionSecurityPropEnumValues = []EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp{
	"none",
	"ssl",
	"starttls",
}

func (v *EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp(value)
	for _, existing := range AllowedEnumexternalServerOracleUnifiedDirectoryConnectionSecurityPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp", value)
}

// NewEnumexternalServerOracleUnifiedDirectoryConnectionSecurityPropFromValue returns a pointer to a valid EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerOracleUnifiedDirectoryConnectionSecurityPropFromValue(v string) (*EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp, error) {
	ev := EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp: valid values are %v", v, AllowedEnumexternalServerOracleUnifiedDirectoryConnectionSecurityPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerOracleUnifiedDirectoryConnectionSecurityPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-oracle-unified-directory-connectionSecurityProp value
func (v EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp) Ptr() *EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp {
	return &v
}

type NullableEnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp struct {
	value *EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp
	isSet bool
}

func (v NullableEnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp) Get() *EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp {
	return v.value
}

func (v *NullableEnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp) Set(val *EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp(val *EnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp) *NullableEnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp {
	return &NullableEnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerOracleUnifiedDirectoryConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
