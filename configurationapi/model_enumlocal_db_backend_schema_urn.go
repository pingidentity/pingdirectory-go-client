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

// EnumlocalDbBackendSchemaUrn the model 'EnumlocalDbBackendSchemaUrn'
type EnumlocalDbBackendSchemaUrn string

// List of Enumlocal-db-backendSchemaUrn
const (
	ENUMLOCALDBBACKENDSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0BACKENDLOCAL_DB EnumlocalDbBackendSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:backend:local-db"
)

// All allowed values of EnumlocalDbBackendSchemaUrn enum
var AllowedEnumlocalDbBackendSchemaUrnEnumValues = []EnumlocalDbBackendSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:backend:local-db",
}

func (v *EnumlocalDbBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlocalDbBackendSchemaUrn(value)
	for _, existing := range AllowedEnumlocalDbBackendSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlocalDbBackendSchemaUrn", value)
}

// NewEnumlocalDbBackendSchemaUrnFromValue returns a pointer to a valid EnumlocalDbBackendSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlocalDbBackendSchemaUrnFromValue(v string) (*EnumlocalDbBackendSchemaUrn, error) {
	ev := EnumlocalDbBackendSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlocalDbBackendSchemaUrn: valid values are %v", v, AllowedEnumlocalDbBackendSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlocalDbBackendSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumlocalDbBackendSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlocal-db-backendSchemaUrn value
func (v EnumlocalDbBackendSchemaUrn) Ptr() *EnumlocalDbBackendSchemaUrn {
	return &v
}

type NullableEnumlocalDbBackendSchemaUrn struct {
	value *EnumlocalDbBackendSchemaUrn
	isSet bool
}

func (v NullableEnumlocalDbBackendSchemaUrn) Get() *EnumlocalDbBackendSchemaUrn {
	return v.value
}

func (v *NullableEnumlocalDbBackendSchemaUrn) Set(val *EnumlocalDbBackendSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlocalDbBackendSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlocalDbBackendSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlocalDbBackendSchemaUrn(val *EnumlocalDbBackendSchemaUrn) *NullableEnumlocalDbBackendSchemaUrn {
	return &NullableEnumlocalDbBackendSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumlocalDbBackendSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlocalDbBackendSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
