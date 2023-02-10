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

// EnumscimAttributeMappingSchemaUrn the model 'EnumscimAttributeMappingSchemaUrn'
type EnumscimAttributeMappingSchemaUrn string

// List of Enumscim-attribute-mappingSchemaUrn
const (
	ENUMSCIMATTRIBUTEMAPPINGSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0SCIM_ATTRIBUTE_MAPPING EnumscimAttributeMappingSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:scim-attribute-mapping"
)

// All allowed values of EnumscimAttributeMappingSchemaUrn enum
var AllowedEnumscimAttributeMappingSchemaUrnEnumValues = []EnumscimAttributeMappingSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:scim-attribute-mapping",
}

func (v *EnumscimAttributeMappingSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumscimAttributeMappingSchemaUrn(value)
	for _, existing := range AllowedEnumscimAttributeMappingSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumscimAttributeMappingSchemaUrn", value)
}

// NewEnumscimAttributeMappingSchemaUrnFromValue returns a pointer to a valid EnumscimAttributeMappingSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumscimAttributeMappingSchemaUrnFromValue(v string) (*EnumscimAttributeMappingSchemaUrn, error) {
	ev := EnumscimAttributeMappingSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumscimAttributeMappingSchemaUrn: valid values are %v", v, AllowedEnumscimAttributeMappingSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumscimAttributeMappingSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumscimAttributeMappingSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumscim-attribute-mappingSchemaUrn value
func (v EnumscimAttributeMappingSchemaUrn) Ptr() *EnumscimAttributeMappingSchemaUrn {
	return &v
}

type NullableEnumscimAttributeMappingSchemaUrn struct {
	value *EnumscimAttributeMappingSchemaUrn
	isSet bool
}

func (v NullableEnumscimAttributeMappingSchemaUrn) Get() *EnumscimAttributeMappingSchemaUrn {
	return v.value
}

func (v *NullableEnumscimAttributeMappingSchemaUrn) Set(val *EnumscimAttributeMappingSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumscimAttributeMappingSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumscimAttributeMappingSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumscimAttributeMappingSchemaUrn(val *EnumscimAttributeMappingSchemaUrn) *NullableEnumscimAttributeMappingSchemaUrn {
	return &NullableEnumscimAttributeMappingSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumscimAttributeMappingSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumscimAttributeMappingSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
