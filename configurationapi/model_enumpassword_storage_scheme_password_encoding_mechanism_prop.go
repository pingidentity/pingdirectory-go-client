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

// EnumpasswordStorageSchemePasswordEncodingMechanismProp Specifies the mechanism that should be used to encode clear-text passwords for use with this scheme.
type EnumpasswordStorageSchemePasswordEncodingMechanismProp string

// List of Enumpassword-storage-scheme-passwordEncodingMechanismProp
const (
	ENUMPASSWORDSTORAGESCHEMEPASSWORDENCODINGMECHANISMPROP_CRYPT     EnumpasswordStorageSchemePasswordEncodingMechanismProp = "crypt"
	ENUMPASSWORDSTORAGESCHEMEPASSWORDENCODINGMECHANISMPROP_MD5       EnumpasswordStorageSchemePasswordEncodingMechanismProp = "md5"
	ENUMPASSWORDSTORAGESCHEMEPASSWORDENCODINGMECHANISMPROP_SHA_2_256 EnumpasswordStorageSchemePasswordEncodingMechanismProp = "sha-2-256"
	ENUMPASSWORDSTORAGESCHEMEPASSWORDENCODINGMECHANISMPROP_SHA_2_512 EnumpasswordStorageSchemePasswordEncodingMechanismProp = "sha-2-512"
)

// All allowed values of EnumpasswordStorageSchemePasswordEncodingMechanismProp enum
var AllowedEnumpasswordStorageSchemePasswordEncodingMechanismPropEnumValues = []EnumpasswordStorageSchemePasswordEncodingMechanismProp{
	"crypt",
	"md5",
	"sha-2-256",
	"sha-2-512",
}

func (v *EnumpasswordStorageSchemePasswordEncodingMechanismProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpasswordStorageSchemePasswordEncodingMechanismProp(value)
	for _, existing := range AllowedEnumpasswordStorageSchemePasswordEncodingMechanismPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpasswordStorageSchemePasswordEncodingMechanismProp", value)
}

// NewEnumpasswordStorageSchemePasswordEncodingMechanismPropFromValue returns a pointer to a valid EnumpasswordStorageSchemePasswordEncodingMechanismProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpasswordStorageSchemePasswordEncodingMechanismPropFromValue(v string) (*EnumpasswordStorageSchemePasswordEncodingMechanismProp, error) {
	ev := EnumpasswordStorageSchemePasswordEncodingMechanismProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpasswordStorageSchemePasswordEncodingMechanismProp: valid values are %v", v, AllowedEnumpasswordStorageSchemePasswordEncodingMechanismPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpasswordStorageSchemePasswordEncodingMechanismProp) IsValid() bool {
	for _, existing := range AllowedEnumpasswordStorageSchemePasswordEncodingMechanismPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumpassword-storage-scheme-passwordEncodingMechanismProp value
func (v EnumpasswordStorageSchemePasswordEncodingMechanismProp) Ptr() *EnumpasswordStorageSchemePasswordEncodingMechanismProp {
	return &v
}

type NullableEnumpasswordStorageSchemePasswordEncodingMechanismProp struct {
	value *EnumpasswordStorageSchemePasswordEncodingMechanismProp
	isSet bool
}

func (v NullableEnumpasswordStorageSchemePasswordEncodingMechanismProp) Get() *EnumpasswordStorageSchemePasswordEncodingMechanismProp {
	return v.value
}

func (v *NullableEnumpasswordStorageSchemePasswordEncodingMechanismProp) Set(val *EnumpasswordStorageSchemePasswordEncodingMechanismProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpasswordStorageSchemePasswordEncodingMechanismProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpasswordStorageSchemePasswordEncodingMechanismProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpasswordStorageSchemePasswordEncodingMechanismProp(val *EnumpasswordStorageSchemePasswordEncodingMechanismProp) *NullableEnumpasswordStorageSchemePasswordEncodingMechanismProp {
	return &NullableEnumpasswordStorageSchemePasswordEncodingMechanismProp{value: val, isSet: true}
}

func (v NullableEnumpasswordStorageSchemePasswordEncodingMechanismProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpasswordStorageSchemePasswordEncodingMechanismProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
