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

// EnumpasswordStorageSchemeDigestAlgorithmProp Specifies the digest algorithm that will be used when encoding passwords.
type EnumpasswordStorageSchemeDigestAlgorithmProp string

// List of Enumpassword-storage-scheme-digestAlgorithmProp
const (
	ENUMPASSWORDSTORAGESCHEMEDIGESTALGORITHMPROP__1 EnumpasswordStorageSchemeDigestAlgorithmProp = "sha-1"
	ENUMPASSWORDSTORAGESCHEMEDIGESTALGORITHMPROP__256 EnumpasswordStorageSchemeDigestAlgorithmProp = "sha-256"
	ENUMPASSWORDSTORAGESCHEMEDIGESTALGORITHMPROP__384 EnumpasswordStorageSchemeDigestAlgorithmProp = "sha-384"
	ENUMPASSWORDSTORAGESCHEMEDIGESTALGORITHMPROP__512 EnumpasswordStorageSchemeDigestAlgorithmProp = "sha-512"
)

// All allowed values of EnumpasswordStorageSchemeDigestAlgorithmProp enum
var AllowedEnumpasswordStorageSchemeDigestAlgorithmPropEnumValues = []EnumpasswordStorageSchemeDigestAlgorithmProp{
	"sha-1",
	"sha-256",
	"sha-384",
	"sha-512",
}

func (v *EnumpasswordStorageSchemeDigestAlgorithmProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpasswordStorageSchemeDigestAlgorithmProp(value)
	for _, existing := range AllowedEnumpasswordStorageSchemeDigestAlgorithmPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpasswordStorageSchemeDigestAlgorithmProp", value)
}

// NewEnumpasswordStorageSchemeDigestAlgorithmPropFromValue returns a pointer to a valid EnumpasswordStorageSchemeDigestAlgorithmProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpasswordStorageSchemeDigestAlgorithmPropFromValue(v string) (*EnumpasswordStorageSchemeDigestAlgorithmProp, error) {
	ev := EnumpasswordStorageSchemeDigestAlgorithmProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpasswordStorageSchemeDigestAlgorithmProp: valid values are %v", v, AllowedEnumpasswordStorageSchemeDigestAlgorithmPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpasswordStorageSchemeDigestAlgorithmProp) IsValid() bool {
	for _, existing := range AllowedEnumpasswordStorageSchemeDigestAlgorithmPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumpassword-storage-scheme-digestAlgorithmProp value
func (v EnumpasswordStorageSchemeDigestAlgorithmProp) Ptr() *EnumpasswordStorageSchemeDigestAlgorithmProp {
	return &v
}

type NullableEnumpasswordStorageSchemeDigestAlgorithmProp struct {
	value *EnumpasswordStorageSchemeDigestAlgorithmProp
	isSet bool
}

func (v NullableEnumpasswordStorageSchemeDigestAlgorithmProp) Get() *EnumpasswordStorageSchemeDigestAlgorithmProp {
	return v.value
}

func (v *NullableEnumpasswordStorageSchemeDigestAlgorithmProp) Set(val *EnumpasswordStorageSchemeDigestAlgorithmProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpasswordStorageSchemeDigestAlgorithmProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpasswordStorageSchemeDigestAlgorithmProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpasswordStorageSchemeDigestAlgorithmProp(val *EnumpasswordStorageSchemeDigestAlgorithmProp) *NullableEnumpasswordStorageSchemeDigestAlgorithmProp {
	return &NullableEnumpasswordStorageSchemeDigestAlgorithmProp{value: val, isSet: true}
}

func (v NullableEnumpasswordStorageSchemeDigestAlgorithmProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpasswordStorageSchemeDigestAlgorithmProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
