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

// Enumsha1PasswordStorageSchemeSchemaUrn the model 'Enumsha1PasswordStorageSchemeSchemaUrn'
type Enumsha1PasswordStorageSchemeSchemaUrn string

// List of Enumsha1-password-storage-schemeSchemaUrn
const (
	ENUMSHA1PASSWORDSTORAGESCHEMESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_STORAGE_SCHEMESHA1 Enumsha1PasswordStorageSchemeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:sha1"
)

// All allowed values of Enumsha1PasswordStorageSchemeSchemaUrn enum
var AllowedEnumsha1PasswordStorageSchemeSchemaUrnEnumValues = []Enumsha1PasswordStorageSchemeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:sha1",
}

func (v *Enumsha1PasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Enumsha1PasswordStorageSchemeSchemaUrn(value)
	for _, existing := range AllowedEnumsha1PasswordStorageSchemeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Enumsha1PasswordStorageSchemeSchemaUrn", value)
}

// NewEnumsha1PasswordStorageSchemeSchemaUrnFromValue returns a pointer to a valid Enumsha1PasswordStorageSchemeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsha1PasswordStorageSchemeSchemaUrnFromValue(v string) (*Enumsha1PasswordStorageSchemeSchemaUrn, error) {
	ev := Enumsha1PasswordStorageSchemeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Enumsha1PasswordStorageSchemeSchemaUrn: valid values are %v", v, AllowedEnumsha1PasswordStorageSchemeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Enumsha1PasswordStorageSchemeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsha1PasswordStorageSchemeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsha1-password-storage-schemeSchemaUrn value
func (v Enumsha1PasswordStorageSchemeSchemaUrn) Ptr() *Enumsha1PasswordStorageSchemeSchemaUrn {
	return &v
}

type NullableEnumsha1PasswordStorageSchemeSchemaUrn struct {
	value *Enumsha1PasswordStorageSchemeSchemaUrn
	isSet bool
}

func (v NullableEnumsha1PasswordStorageSchemeSchemaUrn) Get() *Enumsha1PasswordStorageSchemeSchemaUrn {
	return v.value
}

func (v *NullableEnumsha1PasswordStorageSchemeSchemaUrn) Set(val *Enumsha1PasswordStorageSchemeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsha1PasswordStorageSchemeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsha1PasswordStorageSchemeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsha1PasswordStorageSchemeSchemaUrn(val *Enumsha1PasswordStorageSchemeSchemaUrn) *NullableEnumsha1PasswordStorageSchemeSchemaUrn {
	return &NullableEnumsha1PasswordStorageSchemeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsha1PasswordStorageSchemeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsha1PasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
