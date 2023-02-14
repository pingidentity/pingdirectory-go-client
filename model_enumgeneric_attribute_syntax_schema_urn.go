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

// EnumgenericAttributeSyntaxSchemaUrn the model 'EnumgenericAttributeSyntaxSchemaUrn'
type EnumgenericAttributeSyntaxSchemaUrn string

// List of Enumgeneric-attribute-syntaxSchemaUrn
const (
	ENUMGENERICATTRIBUTESYNTAXSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ATTRIBUTE_SYNTAXGENERIC EnumgenericAttributeSyntaxSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:attribute-syntax:generic"
)

// All allowed values of EnumgenericAttributeSyntaxSchemaUrn enum
var AllowedEnumgenericAttributeSyntaxSchemaUrnEnumValues = []EnumgenericAttributeSyntaxSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:attribute-syntax:generic",
}

func (v *EnumgenericAttributeSyntaxSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgenericAttributeSyntaxSchemaUrn(value)
	for _, existing := range AllowedEnumgenericAttributeSyntaxSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgenericAttributeSyntaxSchemaUrn", value)
}

// NewEnumgenericAttributeSyntaxSchemaUrnFromValue returns a pointer to a valid EnumgenericAttributeSyntaxSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgenericAttributeSyntaxSchemaUrnFromValue(v string) (*EnumgenericAttributeSyntaxSchemaUrn, error) {
	ev := EnumgenericAttributeSyntaxSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgenericAttributeSyntaxSchemaUrn: valid values are %v", v, AllowedEnumgenericAttributeSyntaxSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgenericAttributeSyntaxSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgenericAttributeSyntaxSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgeneric-attribute-syntaxSchemaUrn value
func (v EnumgenericAttributeSyntaxSchemaUrn) Ptr() *EnumgenericAttributeSyntaxSchemaUrn {
	return &v
}

type NullableEnumgenericAttributeSyntaxSchemaUrn struct {
	value *EnumgenericAttributeSyntaxSchemaUrn
	isSet bool
}

func (v NullableEnumgenericAttributeSyntaxSchemaUrn) Get() *EnumgenericAttributeSyntaxSchemaUrn {
	return v.value
}

func (v *NullableEnumgenericAttributeSyntaxSchemaUrn) Set(val *EnumgenericAttributeSyntaxSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgenericAttributeSyntaxSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgenericAttributeSyntaxSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgenericAttributeSyntaxSchemaUrn(val *EnumgenericAttributeSyntaxSchemaUrn) *NullableEnumgenericAttributeSyntaxSchemaUrn {
	return &NullableEnumgenericAttributeSyntaxSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgenericAttributeSyntaxSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgenericAttributeSyntaxSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
