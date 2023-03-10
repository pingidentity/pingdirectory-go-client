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

// EnumjsonAttributeConstraintsSchemaUrn the model 'EnumjsonAttributeConstraintsSchemaUrn'
type EnumjsonAttributeConstraintsSchemaUrn string

// List of Enumjson-attribute-constraintsSchemaUrn
const (
	ENUMJSONATTRIBUTECONSTRAINTSSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0JSON_ATTRIBUTE_CONSTRAINTS EnumjsonAttributeConstraintsSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:json-attribute-constraints"
)

// All allowed values of EnumjsonAttributeConstraintsSchemaUrn enum
var AllowedEnumjsonAttributeConstraintsSchemaUrnEnumValues = []EnumjsonAttributeConstraintsSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:json-attribute-constraints",
}

func (v *EnumjsonAttributeConstraintsSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumjsonAttributeConstraintsSchemaUrn(value)
	for _, existing := range AllowedEnumjsonAttributeConstraintsSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumjsonAttributeConstraintsSchemaUrn", value)
}

// NewEnumjsonAttributeConstraintsSchemaUrnFromValue returns a pointer to a valid EnumjsonAttributeConstraintsSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumjsonAttributeConstraintsSchemaUrnFromValue(v string) (*EnumjsonAttributeConstraintsSchemaUrn, error) {
	ev := EnumjsonAttributeConstraintsSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumjsonAttributeConstraintsSchemaUrn: valid values are %v", v, AllowedEnumjsonAttributeConstraintsSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumjsonAttributeConstraintsSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumjsonAttributeConstraintsSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumjson-attribute-constraintsSchemaUrn value
func (v EnumjsonAttributeConstraintsSchemaUrn) Ptr() *EnumjsonAttributeConstraintsSchemaUrn {
	return &v
}

type NullableEnumjsonAttributeConstraintsSchemaUrn struct {
	value *EnumjsonAttributeConstraintsSchemaUrn
	isSet bool
}

func (v NullableEnumjsonAttributeConstraintsSchemaUrn) Get() *EnumjsonAttributeConstraintsSchemaUrn {
	return v.value
}

func (v *NullableEnumjsonAttributeConstraintsSchemaUrn) Set(val *EnumjsonAttributeConstraintsSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumjsonAttributeConstraintsSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumjsonAttributeConstraintsSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumjsonAttributeConstraintsSchemaUrn(val *EnumjsonAttributeConstraintsSchemaUrn) *NullableEnumjsonAttributeConstraintsSchemaUrn {
	return &NullableEnumjsonAttributeConstraintsSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumjsonAttributeConstraintsSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumjsonAttributeConstraintsSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
