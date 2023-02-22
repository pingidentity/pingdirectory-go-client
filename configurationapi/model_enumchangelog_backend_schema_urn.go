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

// EnumchangelogBackendSchemaUrn the model 'EnumchangelogBackendSchemaUrn'
type EnumchangelogBackendSchemaUrn string

// List of Enumchangelog-backendSchemaUrn
const (
	ENUMCHANGELOGBACKENDSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0BACKENDCHANGELOG EnumchangelogBackendSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:backend:changelog"
)

// All allowed values of EnumchangelogBackendSchemaUrn enum
var AllowedEnumchangelogBackendSchemaUrnEnumValues = []EnumchangelogBackendSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:backend:changelog",
}

func (v *EnumchangelogBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumchangelogBackendSchemaUrn(value)
	for _, existing := range AllowedEnumchangelogBackendSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumchangelogBackendSchemaUrn", value)
}

// NewEnumchangelogBackendSchemaUrnFromValue returns a pointer to a valid EnumchangelogBackendSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumchangelogBackendSchemaUrnFromValue(v string) (*EnumchangelogBackendSchemaUrn, error) {
	ev := EnumchangelogBackendSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumchangelogBackendSchemaUrn: valid values are %v", v, AllowedEnumchangelogBackendSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumchangelogBackendSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumchangelogBackendSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumchangelog-backendSchemaUrn value
func (v EnumchangelogBackendSchemaUrn) Ptr() *EnumchangelogBackendSchemaUrn {
	return &v
}

type NullableEnumchangelogBackendSchemaUrn struct {
	value *EnumchangelogBackendSchemaUrn
	isSet bool
}

func (v NullableEnumchangelogBackendSchemaUrn) Get() *EnumchangelogBackendSchemaUrn {
	return v.value
}

func (v *NullableEnumchangelogBackendSchemaUrn) Set(val *EnumchangelogBackendSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumchangelogBackendSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumchangelogBackendSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumchangelogBackendSchemaUrn(val *EnumchangelogBackendSchemaUrn) *NullableEnumchangelogBackendSchemaUrn {
	return &NullableEnumchangelogBackendSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumchangelogBackendSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumchangelogBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}