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

// Enumrc4PasswordStorageSchemeSchemaUrn the model 'Enumrc4PasswordStorageSchemeSchemaUrn'
type Enumrc4PasswordStorageSchemeSchemaUrn string

// List of Enumrc4-password-storage-schemeSchemaUrn
const (
	ENUMRC4PASSWORDSTORAGESCHEMESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_STORAGE_SCHEMERC4 Enumrc4PasswordStorageSchemeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:rc4"
)

// All allowed values of Enumrc4PasswordStorageSchemeSchemaUrn enum
var AllowedEnumrc4PasswordStorageSchemeSchemaUrnEnumValues = []Enumrc4PasswordStorageSchemeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:rc4",
}

func (v *Enumrc4PasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Enumrc4PasswordStorageSchemeSchemaUrn(value)
	for _, existing := range AllowedEnumrc4PasswordStorageSchemeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Enumrc4PasswordStorageSchemeSchemaUrn", value)
}

// NewEnumrc4PasswordStorageSchemeSchemaUrnFromValue returns a pointer to a valid Enumrc4PasswordStorageSchemeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumrc4PasswordStorageSchemeSchemaUrnFromValue(v string) (*Enumrc4PasswordStorageSchemeSchemaUrn, error) {
	ev := Enumrc4PasswordStorageSchemeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Enumrc4PasswordStorageSchemeSchemaUrn: valid values are %v", v, AllowedEnumrc4PasswordStorageSchemeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Enumrc4PasswordStorageSchemeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumrc4PasswordStorageSchemeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumrc4-password-storage-schemeSchemaUrn value
func (v Enumrc4PasswordStorageSchemeSchemaUrn) Ptr() *Enumrc4PasswordStorageSchemeSchemaUrn {
	return &v
}

type NullableEnumrc4PasswordStorageSchemeSchemaUrn struct {
	value *Enumrc4PasswordStorageSchemeSchemaUrn
	isSet bool
}

func (v NullableEnumrc4PasswordStorageSchemeSchemaUrn) Get() *Enumrc4PasswordStorageSchemeSchemaUrn {
	return v.value
}

func (v *NullableEnumrc4PasswordStorageSchemeSchemaUrn) Set(val *Enumrc4PasswordStorageSchemeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumrc4PasswordStorageSchemeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumrc4PasswordStorageSchemeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumrc4PasswordStorageSchemeSchemaUrn(val *Enumrc4PasswordStorageSchemeSchemaUrn) *NullableEnumrc4PasswordStorageSchemeSchemaUrn {
	return &NullableEnumrc4PasswordStorageSchemeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumrc4PasswordStorageSchemeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumrc4PasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
