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

// EnumpingOneHttpExternalServerSchemaUrn the model 'EnumpingOneHttpExternalServerSchemaUrn'
type EnumpingOneHttpExternalServerSchemaUrn string

// List of Enumping-one-http-external-serverSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERPING_ONE_HTTP EnumpingOneHttpExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:ping-one-http"
)

// All allowed values of EnumpingOneHttpExternalServerSchemaUrn enum
var AllowedEnumpingOneHttpExternalServerSchemaUrnEnumValues = []EnumpingOneHttpExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:ping-one-http",
}

func (v *EnumpingOneHttpExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpingOneHttpExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumpingOneHttpExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpingOneHttpExternalServerSchemaUrn", value)
}

// NewEnumpingOneHttpExternalServerSchemaUrnFromValue returns a pointer to a valid EnumpingOneHttpExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpingOneHttpExternalServerSchemaUrnFromValue(v string) (*EnumpingOneHttpExternalServerSchemaUrn, error) {
	ev := EnumpingOneHttpExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpingOneHttpExternalServerSchemaUrn: valid values are %v", v, AllowedEnumpingOneHttpExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpingOneHttpExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumpingOneHttpExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumping-one-http-external-serverSchemaUrn value
func (v EnumpingOneHttpExternalServerSchemaUrn) Ptr() *EnumpingOneHttpExternalServerSchemaUrn {
	return &v
}

type NullableEnumpingOneHttpExternalServerSchemaUrn struct {
	value *EnumpingOneHttpExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumpingOneHttpExternalServerSchemaUrn) Get() *EnumpingOneHttpExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumpingOneHttpExternalServerSchemaUrn) Set(val *EnumpingOneHttpExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpingOneHttpExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpingOneHttpExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpingOneHttpExternalServerSchemaUrn(val *EnumpingOneHttpExternalServerSchemaUrn) *NullableEnumpingOneHttpExternalServerSchemaUrn {
	return &NullableEnumpingOneHttpExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumpingOneHttpExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpingOneHttpExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

