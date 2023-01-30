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

// EnumtaskBackendSchemaUrn the model 'EnumtaskBackendSchemaUrn'
type EnumtaskBackendSchemaUrn string

// List of Enumtask-backendSchemaUrn
const (
	ENUMTASKBACKENDSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0BACKENDTASK EnumtaskBackendSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:backend:task"
)

// All allowed values of EnumtaskBackendSchemaUrn enum
var AllowedEnumtaskBackendSchemaUrnEnumValues = []EnumtaskBackendSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:backend:task",
}

func (v *EnumtaskBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumtaskBackendSchemaUrn(value)
	for _, existing := range AllowedEnumtaskBackendSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumtaskBackendSchemaUrn", value)
}

// NewEnumtaskBackendSchemaUrnFromValue returns a pointer to a valid EnumtaskBackendSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumtaskBackendSchemaUrnFromValue(v string) (*EnumtaskBackendSchemaUrn, error) {
	ev := EnumtaskBackendSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumtaskBackendSchemaUrn: valid values are %v", v, AllowedEnumtaskBackendSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumtaskBackendSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumtaskBackendSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumtask-backendSchemaUrn value
func (v EnumtaskBackendSchemaUrn) Ptr() *EnumtaskBackendSchemaUrn {
	return &v
}

type NullableEnumtaskBackendSchemaUrn struct {
	value *EnumtaskBackendSchemaUrn
	isSet bool
}

func (v NullableEnumtaskBackendSchemaUrn) Get() *EnumtaskBackendSchemaUrn {
	return v.value
}

func (v *NullableEnumtaskBackendSchemaUrn) Set(val *EnumtaskBackendSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumtaskBackendSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumtaskBackendSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumtaskBackendSchemaUrn(val *EnumtaskBackendSchemaUrn) *NullableEnumtaskBackendSchemaUrn {
	return &NullableEnumtaskBackendSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumtaskBackendSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumtaskBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
