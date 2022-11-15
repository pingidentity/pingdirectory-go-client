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

// EnummetricsBackendSchemaUrn the model 'EnummetricsBackendSchemaUrn'
type EnummetricsBackendSchemaUrn string

// List of Enummetrics-backendSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0BACKENDMETRICS EnummetricsBackendSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:backend:metrics"
)

// All allowed values of EnummetricsBackendSchemaUrn enum
var AllowedEnummetricsBackendSchemaUrnEnumValues = []EnummetricsBackendSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:backend:metrics",
}

func (v *EnummetricsBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnummetricsBackendSchemaUrn(value)
	for _, existing := range AllowedEnummetricsBackendSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnummetricsBackendSchemaUrn", value)
}

// NewEnummetricsBackendSchemaUrnFromValue returns a pointer to a valid EnummetricsBackendSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnummetricsBackendSchemaUrnFromValue(v string) (*EnummetricsBackendSchemaUrn, error) {
	ev := EnummetricsBackendSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnummetricsBackendSchemaUrn: valid values are %v", v, AllowedEnummetricsBackendSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnummetricsBackendSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnummetricsBackendSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enummetrics-backendSchemaUrn value
func (v EnummetricsBackendSchemaUrn) Ptr() *EnummetricsBackendSchemaUrn {
	return &v
}

type NullableEnummetricsBackendSchemaUrn struct {
	value *EnummetricsBackendSchemaUrn
	isSet bool
}

func (v NullableEnummetricsBackendSchemaUrn) Get() *EnummetricsBackendSchemaUrn {
	return v.value
}

func (v *NullableEnummetricsBackendSchemaUrn) Set(val *EnummetricsBackendSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnummetricsBackendSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnummetricsBackendSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnummetricsBackendSchemaUrn(val *EnummetricsBackendSchemaUrn) *NullableEnummetricsBackendSchemaUrn {
	return &NullableEnummetricsBackendSchemaUrn{value: val, isSet: true}
}

func (v NullableEnummetricsBackendSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnummetricsBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

