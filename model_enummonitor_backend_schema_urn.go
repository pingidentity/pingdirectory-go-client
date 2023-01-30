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

// EnummonitorBackendSchemaUrn the model 'EnummonitorBackendSchemaUrn'
type EnummonitorBackendSchemaUrn string

// List of Enummonitor-backendSchemaUrn
const (
	ENUMMONITORBACKENDSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0BACKENDMONITOR EnummonitorBackendSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:backend:monitor"
)

// All allowed values of EnummonitorBackendSchemaUrn enum
var AllowedEnummonitorBackendSchemaUrnEnumValues = []EnummonitorBackendSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:backend:monitor",
}

func (v *EnummonitorBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnummonitorBackendSchemaUrn(value)
	for _, existing := range AllowedEnummonitorBackendSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnummonitorBackendSchemaUrn", value)
}

// NewEnummonitorBackendSchemaUrnFromValue returns a pointer to a valid EnummonitorBackendSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnummonitorBackendSchemaUrnFromValue(v string) (*EnummonitorBackendSchemaUrn, error) {
	ev := EnummonitorBackendSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnummonitorBackendSchemaUrn: valid values are %v", v, AllowedEnummonitorBackendSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnummonitorBackendSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnummonitorBackendSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enummonitor-backendSchemaUrn value
func (v EnummonitorBackendSchemaUrn) Ptr() *EnummonitorBackendSchemaUrn {
	return &v
}

type NullableEnummonitorBackendSchemaUrn struct {
	value *EnummonitorBackendSchemaUrn
	isSet bool
}

func (v NullableEnummonitorBackendSchemaUrn) Get() *EnummonitorBackendSchemaUrn {
	return v.value
}

func (v *NullableEnummonitorBackendSchemaUrn) Set(val *EnummonitorBackendSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnummonitorBackendSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnummonitorBackendSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnummonitorBackendSchemaUrn(val *EnummonitorBackendSchemaUrn) *NullableEnummonitorBackendSchemaUrn {
	return &NullableEnummonitorBackendSchemaUrn{value: val, isSet: true}
}

func (v NullableEnummonitorBackendSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnummonitorBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
