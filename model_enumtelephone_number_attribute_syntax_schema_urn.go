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

// EnumtelephoneNumberAttributeSyntaxSchemaUrn the model 'EnumtelephoneNumberAttributeSyntaxSchemaUrn'
type EnumtelephoneNumberAttributeSyntaxSchemaUrn string

// List of Enumtelephone-number-attribute-syntaxSchemaUrn
const (
	ENUMTELEPHONENUMBERATTRIBUTESYNTAXSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ATTRIBUTE_SYNTAXTELEPHONE_NUMBER EnumtelephoneNumberAttributeSyntaxSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:attribute-syntax:telephone-number"
)

// All allowed values of EnumtelephoneNumberAttributeSyntaxSchemaUrn enum
var AllowedEnumtelephoneNumberAttributeSyntaxSchemaUrnEnumValues = []EnumtelephoneNumberAttributeSyntaxSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:attribute-syntax:telephone-number",
}

func (v *EnumtelephoneNumberAttributeSyntaxSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumtelephoneNumberAttributeSyntaxSchemaUrn(value)
	for _, existing := range AllowedEnumtelephoneNumberAttributeSyntaxSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumtelephoneNumberAttributeSyntaxSchemaUrn", value)
}

// NewEnumtelephoneNumberAttributeSyntaxSchemaUrnFromValue returns a pointer to a valid EnumtelephoneNumberAttributeSyntaxSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumtelephoneNumberAttributeSyntaxSchemaUrnFromValue(v string) (*EnumtelephoneNumberAttributeSyntaxSchemaUrn, error) {
	ev := EnumtelephoneNumberAttributeSyntaxSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumtelephoneNumberAttributeSyntaxSchemaUrn: valid values are %v", v, AllowedEnumtelephoneNumberAttributeSyntaxSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumtelephoneNumberAttributeSyntaxSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumtelephoneNumberAttributeSyntaxSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumtelephone-number-attribute-syntaxSchemaUrn value
func (v EnumtelephoneNumberAttributeSyntaxSchemaUrn) Ptr() *EnumtelephoneNumberAttributeSyntaxSchemaUrn {
	return &v
}

type NullableEnumtelephoneNumberAttributeSyntaxSchemaUrn struct {
	value *EnumtelephoneNumberAttributeSyntaxSchemaUrn
	isSet bool
}

func (v NullableEnumtelephoneNumberAttributeSyntaxSchemaUrn) Get() *EnumtelephoneNumberAttributeSyntaxSchemaUrn {
	return v.value
}

func (v *NullableEnumtelephoneNumberAttributeSyntaxSchemaUrn) Set(val *EnumtelephoneNumberAttributeSyntaxSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumtelephoneNumberAttributeSyntaxSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumtelephoneNumberAttributeSyntaxSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumtelephoneNumberAttributeSyntaxSchemaUrn(val *EnumtelephoneNumberAttributeSyntaxSchemaUrn) *NullableEnumtelephoneNumberAttributeSyntaxSchemaUrn {
	return &NullableEnumtelephoneNumberAttributeSyntaxSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumtelephoneNumberAttributeSyntaxSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumtelephoneNumberAttributeSyntaxSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

