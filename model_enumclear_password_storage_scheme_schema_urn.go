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

// EnumclearPasswordStorageSchemeSchemaUrn the model 'EnumclearPasswordStorageSchemeSchemaUrn'
type EnumclearPasswordStorageSchemeSchemaUrn string

// List of Enumclear-password-storage-schemeSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_STORAGE_SCHEMECLEAR EnumclearPasswordStorageSchemeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:clear"
)

// All allowed values of EnumclearPasswordStorageSchemeSchemaUrn enum
var AllowedEnumclearPasswordStorageSchemeSchemaUrnEnumValues = []EnumclearPasswordStorageSchemeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:clear",
}

func (v *EnumclearPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumclearPasswordStorageSchemeSchemaUrn(value)
	for _, existing := range AllowedEnumclearPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumclearPasswordStorageSchemeSchemaUrn", value)
}

// NewEnumclearPasswordStorageSchemeSchemaUrnFromValue returns a pointer to a valid EnumclearPasswordStorageSchemeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumclearPasswordStorageSchemeSchemaUrnFromValue(v string) (*EnumclearPasswordStorageSchemeSchemaUrn, error) {
	ev := EnumclearPasswordStorageSchemeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumclearPasswordStorageSchemeSchemaUrn: valid values are %v", v, AllowedEnumclearPasswordStorageSchemeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumclearPasswordStorageSchemeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumclearPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumclear-password-storage-schemeSchemaUrn value
func (v EnumclearPasswordStorageSchemeSchemaUrn) Ptr() *EnumclearPasswordStorageSchemeSchemaUrn {
	return &v
}

type NullableEnumclearPasswordStorageSchemeSchemaUrn struct {
	value *EnumclearPasswordStorageSchemeSchemaUrn
	isSet bool
}

func (v NullableEnumclearPasswordStorageSchemeSchemaUrn) Get() *EnumclearPasswordStorageSchemeSchemaUrn {
	return v.value
}

func (v *NullableEnumclearPasswordStorageSchemeSchemaUrn) Set(val *EnumclearPasswordStorageSchemeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumclearPasswordStorageSchemeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumclearPasswordStorageSchemeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumclearPasswordStorageSchemeSchemaUrn(val *EnumclearPasswordStorageSchemeSchemaUrn) *NullableEnumclearPasswordStorageSchemeSchemaUrn {
	return &NullableEnumclearPasswordStorageSchemeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumclearPasswordStorageSchemeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumclearPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

