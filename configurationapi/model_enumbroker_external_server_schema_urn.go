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

// EnumbrokerExternalServerSchemaUrn the model 'EnumbrokerExternalServerSchemaUrn'
type EnumbrokerExternalServerSchemaUrn string

// List of Enumbroker-external-serverSchemaUrn
const (
	ENUMBROKEREXTERNALSERVERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERBROKER EnumbrokerExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:broker"
)

// All allowed values of EnumbrokerExternalServerSchemaUrn enum
var AllowedEnumbrokerExternalServerSchemaUrnEnumValues = []EnumbrokerExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:broker",
}

func (v *EnumbrokerExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumbrokerExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumbrokerExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumbrokerExternalServerSchemaUrn", value)
}

// NewEnumbrokerExternalServerSchemaUrnFromValue returns a pointer to a valid EnumbrokerExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumbrokerExternalServerSchemaUrnFromValue(v string) (*EnumbrokerExternalServerSchemaUrn, error) {
	ev := EnumbrokerExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumbrokerExternalServerSchemaUrn: valid values are %v", v, AllowedEnumbrokerExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumbrokerExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumbrokerExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumbroker-external-serverSchemaUrn value
func (v EnumbrokerExternalServerSchemaUrn) Ptr() *EnumbrokerExternalServerSchemaUrn {
	return &v
}

type NullableEnumbrokerExternalServerSchemaUrn struct {
	value *EnumbrokerExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumbrokerExternalServerSchemaUrn) Get() *EnumbrokerExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumbrokerExternalServerSchemaUrn) Set(val *EnumbrokerExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumbrokerExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumbrokerExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumbrokerExternalServerSchemaUrn(val *EnumbrokerExternalServerSchemaUrn) *NullableEnumbrokerExternalServerSchemaUrn {
	return &NullableEnumbrokerExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumbrokerExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumbrokerExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
