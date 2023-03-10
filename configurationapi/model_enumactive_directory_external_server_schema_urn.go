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

// EnumactiveDirectoryExternalServerSchemaUrn the model 'EnumactiveDirectoryExternalServerSchemaUrn'
type EnumactiveDirectoryExternalServerSchemaUrn string

// List of Enumactive-directory-external-serverSchemaUrn
const (
	ENUMACTIVEDIRECTORYEXTERNALSERVERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERACTIVE_DIRECTORY EnumactiveDirectoryExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:active-directory"
)

// All allowed values of EnumactiveDirectoryExternalServerSchemaUrn enum
var AllowedEnumactiveDirectoryExternalServerSchemaUrnEnumValues = []EnumactiveDirectoryExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:active-directory",
}

func (v *EnumactiveDirectoryExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumactiveDirectoryExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumactiveDirectoryExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumactiveDirectoryExternalServerSchemaUrn", value)
}

// NewEnumactiveDirectoryExternalServerSchemaUrnFromValue returns a pointer to a valid EnumactiveDirectoryExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumactiveDirectoryExternalServerSchemaUrnFromValue(v string) (*EnumactiveDirectoryExternalServerSchemaUrn, error) {
	ev := EnumactiveDirectoryExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumactiveDirectoryExternalServerSchemaUrn: valid values are %v", v, AllowedEnumactiveDirectoryExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumactiveDirectoryExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumactiveDirectoryExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumactive-directory-external-serverSchemaUrn value
func (v EnumactiveDirectoryExternalServerSchemaUrn) Ptr() *EnumactiveDirectoryExternalServerSchemaUrn {
	return &v
}

type NullableEnumactiveDirectoryExternalServerSchemaUrn struct {
	value *EnumactiveDirectoryExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumactiveDirectoryExternalServerSchemaUrn) Get() *EnumactiveDirectoryExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumactiveDirectoryExternalServerSchemaUrn) Set(val *EnumactiveDirectoryExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumactiveDirectoryExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumactiveDirectoryExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumactiveDirectoryExternalServerSchemaUrn(val *EnumactiveDirectoryExternalServerSchemaUrn) *NullableEnumactiveDirectoryExternalServerSchemaUrn {
	return &NullableEnumactiveDirectoryExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumactiveDirectoryExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumactiveDirectoryExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
