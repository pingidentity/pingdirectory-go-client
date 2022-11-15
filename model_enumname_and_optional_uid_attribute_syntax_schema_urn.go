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

// EnumnameAndOptionalUidAttributeSyntaxSchemaUrn the model 'EnumnameAndOptionalUidAttributeSyntaxSchemaUrn'
type EnumnameAndOptionalUidAttributeSyntaxSchemaUrn string

// List of Enumname-and-optional-uid-attribute-syntaxSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0ATTRIBUTE_SYNTAXNAME_AND_OPTIONAL_UID EnumnameAndOptionalUidAttributeSyntaxSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:attribute-syntax:name-and-optional-uid"
)

// All allowed values of EnumnameAndOptionalUidAttributeSyntaxSchemaUrn enum
var AllowedEnumnameAndOptionalUidAttributeSyntaxSchemaUrnEnumValues = []EnumnameAndOptionalUidAttributeSyntaxSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:attribute-syntax:name-and-optional-uid",
}

func (v *EnumnameAndOptionalUidAttributeSyntaxSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumnameAndOptionalUidAttributeSyntaxSchemaUrn(value)
	for _, existing := range AllowedEnumnameAndOptionalUidAttributeSyntaxSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumnameAndOptionalUidAttributeSyntaxSchemaUrn", value)
}

// NewEnumnameAndOptionalUidAttributeSyntaxSchemaUrnFromValue returns a pointer to a valid EnumnameAndOptionalUidAttributeSyntaxSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumnameAndOptionalUidAttributeSyntaxSchemaUrnFromValue(v string) (*EnumnameAndOptionalUidAttributeSyntaxSchemaUrn, error) {
	ev := EnumnameAndOptionalUidAttributeSyntaxSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumnameAndOptionalUidAttributeSyntaxSchemaUrn: valid values are %v", v, AllowedEnumnameAndOptionalUidAttributeSyntaxSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumnameAndOptionalUidAttributeSyntaxSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumnameAndOptionalUidAttributeSyntaxSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumname-and-optional-uid-attribute-syntaxSchemaUrn value
func (v EnumnameAndOptionalUidAttributeSyntaxSchemaUrn) Ptr() *EnumnameAndOptionalUidAttributeSyntaxSchemaUrn {
	return &v
}

type NullableEnumnameAndOptionalUidAttributeSyntaxSchemaUrn struct {
	value *EnumnameAndOptionalUidAttributeSyntaxSchemaUrn
	isSet bool
}

func (v NullableEnumnameAndOptionalUidAttributeSyntaxSchemaUrn) Get() *EnumnameAndOptionalUidAttributeSyntaxSchemaUrn {
	return v.value
}

func (v *NullableEnumnameAndOptionalUidAttributeSyntaxSchemaUrn) Set(val *EnumnameAndOptionalUidAttributeSyntaxSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumnameAndOptionalUidAttributeSyntaxSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumnameAndOptionalUidAttributeSyntaxSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumnameAndOptionalUidAttributeSyntaxSchemaUrn(val *EnumnameAndOptionalUidAttributeSyntaxSchemaUrn) *NullableEnumnameAndOptionalUidAttributeSyntaxSchemaUrn {
	return &NullableEnumnameAndOptionalUidAttributeSyntaxSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumnameAndOptionalUidAttributeSyntaxSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumnameAndOptionalUidAttributeSyntaxSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

