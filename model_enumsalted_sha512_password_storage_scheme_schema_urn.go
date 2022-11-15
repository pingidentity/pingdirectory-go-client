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

// EnumsaltedSha512PasswordStorageSchemeSchemaUrn the model 'EnumsaltedSha512PasswordStorageSchemeSchemaUrn'
type EnumsaltedSha512PasswordStorageSchemeSchemaUrn string

// List of Enumsalted-sha512-password-storage-schemeSchemaUrn
const (
	ENUMSALTEDSHA512PASSWORDSTORAGESCHEMESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_STORAGE_SCHEMESALTED_SHA512 EnumsaltedSha512PasswordStorageSchemeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:salted-sha512"
)

// All allowed values of EnumsaltedSha512PasswordStorageSchemeSchemaUrn enum
var AllowedEnumsaltedSha512PasswordStorageSchemeSchemaUrnEnumValues = []EnumsaltedSha512PasswordStorageSchemeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:salted-sha512",
}

func (v *EnumsaltedSha512PasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsaltedSha512PasswordStorageSchemeSchemaUrn(value)
	for _, existing := range AllowedEnumsaltedSha512PasswordStorageSchemeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsaltedSha512PasswordStorageSchemeSchemaUrn", value)
}

// NewEnumsaltedSha512PasswordStorageSchemeSchemaUrnFromValue returns a pointer to a valid EnumsaltedSha512PasswordStorageSchemeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsaltedSha512PasswordStorageSchemeSchemaUrnFromValue(v string) (*EnumsaltedSha512PasswordStorageSchemeSchemaUrn, error) {
	ev := EnumsaltedSha512PasswordStorageSchemeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsaltedSha512PasswordStorageSchemeSchemaUrn: valid values are %v", v, AllowedEnumsaltedSha512PasswordStorageSchemeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsaltedSha512PasswordStorageSchemeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsaltedSha512PasswordStorageSchemeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsalted-sha512-password-storage-schemeSchemaUrn value
func (v EnumsaltedSha512PasswordStorageSchemeSchemaUrn) Ptr() *EnumsaltedSha512PasswordStorageSchemeSchemaUrn {
	return &v
}

type NullableEnumsaltedSha512PasswordStorageSchemeSchemaUrn struct {
	value *EnumsaltedSha512PasswordStorageSchemeSchemaUrn
	isSet bool
}

func (v NullableEnumsaltedSha512PasswordStorageSchemeSchemaUrn) Get() *EnumsaltedSha512PasswordStorageSchemeSchemaUrn {
	return v.value
}

func (v *NullableEnumsaltedSha512PasswordStorageSchemeSchemaUrn) Set(val *EnumsaltedSha512PasswordStorageSchemeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsaltedSha512PasswordStorageSchemeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsaltedSha512PasswordStorageSchemeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsaltedSha512PasswordStorageSchemeSchemaUrn(val *EnumsaltedSha512PasswordStorageSchemeSchemaUrn) *NullableEnumsaltedSha512PasswordStorageSchemeSchemaUrn {
	return &NullableEnumsaltedSha512PasswordStorageSchemeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsaltedSha512PasswordStorageSchemeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsaltedSha512PasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

