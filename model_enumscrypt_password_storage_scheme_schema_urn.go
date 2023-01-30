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

// EnumscryptPasswordStorageSchemeSchemaUrn the model 'EnumscryptPasswordStorageSchemeSchemaUrn'
type EnumscryptPasswordStorageSchemeSchemaUrn string

// List of Enumscrypt-password-storage-schemeSchemaUrn
const (
	ENUMSCRYPTPASSWORDSTORAGESCHEMESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_STORAGE_SCHEMESCRYPT EnumscryptPasswordStorageSchemeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:scrypt"
)

// All allowed values of EnumscryptPasswordStorageSchemeSchemaUrn enum
var AllowedEnumscryptPasswordStorageSchemeSchemaUrnEnumValues = []EnumscryptPasswordStorageSchemeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:scrypt",
}

func (v *EnumscryptPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumscryptPasswordStorageSchemeSchemaUrn(value)
	for _, existing := range AllowedEnumscryptPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumscryptPasswordStorageSchemeSchemaUrn", value)
}

// NewEnumscryptPasswordStorageSchemeSchemaUrnFromValue returns a pointer to a valid EnumscryptPasswordStorageSchemeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumscryptPasswordStorageSchemeSchemaUrnFromValue(v string) (*EnumscryptPasswordStorageSchemeSchemaUrn, error) {
	ev := EnumscryptPasswordStorageSchemeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumscryptPasswordStorageSchemeSchemaUrn: valid values are %v", v, AllowedEnumscryptPasswordStorageSchemeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumscryptPasswordStorageSchemeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumscryptPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumscrypt-password-storage-schemeSchemaUrn value
func (v EnumscryptPasswordStorageSchemeSchemaUrn) Ptr() *EnumscryptPasswordStorageSchemeSchemaUrn {
	return &v
}

type NullableEnumscryptPasswordStorageSchemeSchemaUrn struct {
	value *EnumscryptPasswordStorageSchemeSchemaUrn
	isSet bool
}

func (v NullableEnumscryptPasswordStorageSchemeSchemaUrn) Get() *EnumscryptPasswordStorageSchemeSchemaUrn {
	return v.value
}

func (v *NullableEnumscryptPasswordStorageSchemeSchemaUrn) Set(val *EnumscryptPasswordStorageSchemeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumscryptPasswordStorageSchemeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumscryptPasswordStorageSchemeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumscryptPasswordStorageSchemeSchemaUrn(val *EnumscryptPasswordStorageSchemeSchemaUrn) *NullableEnumscryptPasswordStorageSchemeSchemaUrn {
	return &NullableEnumscryptPasswordStorageSchemeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumscryptPasswordStorageSchemeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumscryptPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
