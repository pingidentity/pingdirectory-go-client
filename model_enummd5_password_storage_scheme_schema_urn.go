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

// Enummd5PasswordStorageSchemeSchemaUrn the model 'Enummd5PasswordStorageSchemeSchemaUrn'
type Enummd5PasswordStorageSchemeSchemaUrn string

// List of Enummd5-password-storage-schemeSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_STORAGE_SCHEMEMD5 Enummd5PasswordStorageSchemeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:md5"
)

// All allowed values of Enummd5PasswordStorageSchemeSchemaUrn enum
var AllowedEnummd5PasswordStorageSchemeSchemaUrnEnumValues = []Enummd5PasswordStorageSchemeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:md5",
}

func (v *Enummd5PasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Enummd5PasswordStorageSchemeSchemaUrn(value)
	for _, existing := range AllowedEnummd5PasswordStorageSchemeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Enummd5PasswordStorageSchemeSchemaUrn", value)
}

// NewEnummd5PasswordStorageSchemeSchemaUrnFromValue returns a pointer to a valid Enummd5PasswordStorageSchemeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnummd5PasswordStorageSchemeSchemaUrnFromValue(v string) (*Enummd5PasswordStorageSchemeSchemaUrn, error) {
	ev := Enummd5PasswordStorageSchemeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Enummd5PasswordStorageSchemeSchemaUrn: valid values are %v", v, AllowedEnummd5PasswordStorageSchemeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Enummd5PasswordStorageSchemeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnummd5PasswordStorageSchemeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enummd5-password-storage-schemeSchemaUrn value
func (v Enummd5PasswordStorageSchemeSchemaUrn) Ptr() *Enummd5PasswordStorageSchemeSchemaUrn {
	return &v
}

type NullableEnummd5PasswordStorageSchemeSchemaUrn struct {
	value *Enummd5PasswordStorageSchemeSchemaUrn
	isSet bool
}

func (v NullableEnummd5PasswordStorageSchemeSchemaUrn) Get() *Enummd5PasswordStorageSchemeSchemaUrn {
	return v.value
}

func (v *NullableEnummd5PasswordStorageSchemeSchemaUrn) Set(val *Enummd5PasswordStorageSchemeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnummd5PasswordStorageSchemeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnummd5PasswordStorageSchemeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnummd5PasswordStorageSchemeSchemaUrn(val *Enummd5PasswordStorageSchemeSchemaUrn) *NullableEnummd5PasswordStorageSchemeSchemaUrn {
	return &NullableEnummd5PasswordStorageSchemeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnummd5PasswordStorageSchemeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnummd5PasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

