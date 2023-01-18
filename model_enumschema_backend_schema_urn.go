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

// EnumschemaBackendSchemaUrn the model 'EnumschemaBackendSchemaUrn'
type EnumschemaBackendSchemaUrn string

// List of Enumschema-backendSchemaUrn
const (
	ENUMSCHEMABACKENDSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0BACKENDSCHEMA EnumschemaBackendSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:backend:schema"
)

// All allowed values of EnumschemaBackendSchemaUrn enum
var AllowedEnumschemaBackendSchemaUrnEnumValues = []EnumschemaBackendSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:backend:schema",
}

func (v *EnumschemaBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumschemaBackendSchemaUrn(value)
	for _, existing := range AllowedEnumschemaBackendSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumschemaBackendSchemaUrn", value)
}

// NewEnumschemaBackendSchemaUrnFromValue returns a pointer to a valid EnumschemaBackendSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumschemaBackendSchemaUrnFromValue(v string) (*EnumschemaBackendSchemaUrn, error) {
	ev := EnumschemaBackendSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumschemaBackendSchemaUrn: valid values are %v", v, AllowedEnumschemaBackendSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumschemaBackendSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumschemaBackendSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumschema-backendSchemaUrn value
func (v EnumschemaBackendSchemaUrn) Ptr() *EnumschemaBackendSchemaUrn {
	return &v
}

type NullableEnumschemaBackendSchemaUrn struct {
	value *EnumschemaBackendSchemaUrn
	isSet bool
}

func (v NullableEnumschemaBackendSchemaUrn) Get() *EnumschemaBackendSchemaUrn {
	return v.value
}

func (v *NullableEnumschemaBackendSchemaUrn) Set(val *EnumschemaBackendSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumschemaBackendSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumschemaBackendSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumschemaBackendSchemaUrn(val *EnumschemaBackendSchemaUrn) *NullableEnumschemaBackendSchemaUrn {
	return &NullableEnumschemaBackendSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumschemaBackendSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumschemaBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
