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

// Enumargon2dPasswordStorageSchemeSchemaUrn the model 'Enumargon2dPasswordStorageSchemeSchemaUrn'
type Enumargon2dPasswordStorageSchemeSchemaUrn string

// List of Enumargon2d-password-storage-schemeSchemaUrn
const (
	ENUMARGON2DPASSWORDSTORAGESCHEMESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_STORAGE_SCHEMEARGON2D Enumargon2dPasswordStorageSchemeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:argon2d"
)

// All allowed values of Enumargon2dPasswordStorageSchemeSchemaUrn enum
var AllowedEnumargon2dPasswordStorageSchemeSchemaUrnEnumValues = []Enumargon2dPasswordStorageSchemeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:argon2d",
}

func (v *Enumargon2dPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Enumargon2dPasswordStorageSchemeSchemaUrn(value)
	for _, existing := range AllowedEnumargon2dPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Enumargon2dPasswordStorageSchemeSchemaUrn", value)
}

// NewEnumargon2dPasswordStorageSchemeSchemaUrnFromValue returns a pointer to a valid Enumargon2dPasswordStorageSchemeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumargon2dPasswordStorageSchemeSchemaUrnFromValue(v string) (*Enumargon2dPasswordStorageSchemeSchemaUrn, error) {
	ev := Enumargon2dPasswordStorageSchemeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Enumargon2dPasswordStorageSchemeSchemaUrn: valid values are %v", v, AllowedEnumargon2dPasswordStorageSchemeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Enumargon2dPasswordStorageSchemeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumargon2dPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumargon2d-password-storage-schemeSchemaUrn value
func (v Enumargon2dPasswordStorageSchemeSchemaUrn) Ptr() *Enumargon2dPasswordStorageSchemeSchemaUrn {
	return &v
}

type NullableEnumargon2dPasswordStorageSchemeSchemaUrn struct {
	value *Enumargon2dPasswordStorageSchemeSchemaUrn
	isSet bool
}

func (v NullableEnumargon2dPasswordStorageSchemeSchemaUrn) Get() *Enumargon2dPasswordStorageSchemeSchemaUrn {
	return v.value
}

func (v *NullableEnumargon2dPasswordStorageSchemeSchemaUrn) Set(val *Enumargon2dPasswordStorageSchemeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumargon2dPasswordStorageSchemeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumargon2dPasswordStorageSchemeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumargon2dPasswordStorageSchemeSchemaUrn(val *Enumargon2dPasswordStorageSchemeSchemaUrn) *NullableEnumargon2dPasswordStorageSchemeSchemaUrn {
	return &NullableEnumargon2dPasswordStorageSchemeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumargon2dPasswordStorageSchemeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumargon2dPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
