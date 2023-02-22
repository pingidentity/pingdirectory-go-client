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

// EnumcryptPasswordStorageSchemeSchemaUrn the model 'EnumcryptPasswordStorageSchemeSchemaUrn'
type EnumcryptPasswordStorageSchemeSchemaUrn string

// List of Enumcrypt-password-storage-schemeSchemaUrn
const (
	ENUMCRYPTPASSWORDSTORAGESCHEMESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_STORAGE_SCHEMECRYPT EnumcryptPasswordStorageSchemeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:crypt"
)

// All allowed values of EnumcryptPasswordStorageSchemeSchemaUrn enum
var AllowedEnumcryptPasswordStorageSchemeSchemaUrnEnumValues = []EnumcryptPasswordStorageSchemeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:crypt",
}

func (v *EnumcryptPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcryptPasswordStorageSchemeSchemaUrn(value)
	for _, existing := range AllowedEnumcryptPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcryptPasswordStorageSchemeSchemaUrn", value)
}

// NewEnumcryptPasswordStorageSchemeSchemaUrnFromValue returns a pointer to a valid EnumcryptPasswordStorageSchemeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcryptPasswordStorageSchemeSchemaUrnFromValue(v string) (*EnumcryptPasswordStorageSchemeSchemaUrn, error) {
	ev := EnumcryptPasswordStorageSchemeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcryptPasswordStorageSchemeSchemaUrn: valid values are %v", v, AllowedEnumcryptPasswordStorageSchemeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcryptPasswordStorageSchemeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcryptPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumcrypt-password-storage-schemeSchemaUrn value
func (v EnumcryptPasswordStorageSchemeSchemaUrn) Ptr() *EnumcryptPasswordStorageSchemeSchemaUrn {
	return &v
}

type NullableEnumcryptPasswordStorageSchemeSchemaUrn struct {
	value *EnumcryptPasswordStorageSchemeSchemaUrn
	isSet bool
}

func (v NullableEnumcryptPasswordStorageSchemeSchemaUrn) Get() *EnumcryptPasswordStorageSchemeSchemaUrn {
	return v.value
}

func (v *NullableEnumcryptPasswordStorageSchemeSchemaUrn) Set(val *EnumcryptPasswordStorageSchemeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcryptPasswordStorageSchemeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcryptPasswordStorageSchemeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcryptPasswordStorageSchemeSchemaUrn(val *EnumcryptPasswordStorageSchemeSchemaUrn) *NullableEnumcryptPasswordStorageSchemeSchemaUrn {
	return &NullableEnumcryptPasswordStorageSchemeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcryptPasswordStorageSchemeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcryptPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}