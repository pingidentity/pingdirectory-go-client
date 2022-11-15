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

// EnumaesPasswordStorageSchemeSchemaUrn the model 'EnumaesPasswordStorageSchemeSchemaUrn'
type EnumaesPasswordStorageSchemeSchemaUrn string

// List of Enumaes-password-storage-schemeSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_STORAGE_SCHEMEAES EnumaesPasswordStorageSchemeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:aes"
)

// All allowed values of EnumaesPasswordStorageSchemeSchemaUrn enum
var AllowedEnumaesPasswordStorageSchemeSchemaUrnEnumValues = []EnumaesPasswordStorageSchemeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-storage-scheme:aes",
}

func (v *EnumaesPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumaesPasswordStorageSchemeSchemaUrn(value)
	for _, existing := range AllowedEnumaesPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumaesPasswordStorageSchemeSchemaUrn", value)
}

// NewEnumaesPasswordStorageSchemeSchemaUrnFromValue returns a pointer to a valid EnumaesPasswordStorageSchemeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumaesPasswordStorageSchemeSchemaUrnFromValue(v string) (*EnumaesPasswordStorageSchemeSchemaUrn, error) {
	ev := EnumaesPasswordStorageSchemeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumaesPasswordStorageSchemeSchemaUrn: valid values are %v", v, AllowedEnumaesPasswordStorageSchemeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumaesPasswordStorageSchemeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumaesPasswordStorageSchemeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumaes-password-storage-schemeSchemaUrn value
func (v EnumaesPasswordStorageSchemeSchemaUrn) Ptr() *EnumaesPasswordStorageSchemeSchemaUrn {
	return &v
}

type NullableEnumaesPasswordStorageSchemeSchemaUrn struct {
	value *EnumaesPasswordStorageSchemeSchemaUrn
	isSet bool
}

func (v NullableEnumaesPasswordStorageSchemeSchemaUrn) Get() *EnumaesPasswordStorageSchemeSchemaUrn {
	return v.value
}

func (v *NullableEnumaesPasswordStorageSchemeSchemaUrn) Set(val *EnumaesPasswordStorageSchemeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumaesPasswordStorageSchemeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumaesPasswordStorageSchemeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumaesPasswordStorageSchemeSchemaUrn(val *EnumaesPasswordStorageSchemeSchemaUrn) *NullableEnumaesPasswordStorageSchemeSchemaUrn {
	return &NullableEnumaesPasswordStorageSchemeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumaesPasswordStorageSchemeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumaesPasswordStorageSchemeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

