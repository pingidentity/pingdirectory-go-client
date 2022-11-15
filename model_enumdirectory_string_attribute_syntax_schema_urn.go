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

// EnumdirectoryStringAttributeSyntaxSchemaUrn the model 'EnumdirectoryStringAttributeSyntaxSchemaUrn'
type EnumdirectoryStringAttributeSyntaxSchemaUrn string

// List of Enumdirectory-string-attribute-syntaxSchemaUrn
const (
	ENUMDIRECTORYSTRINGATTRIBUTESYNTAXSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ATTRIBUTE_SYNTAXDIRECTORY_STRING EnumdirectoryStringAttributeSyntaxSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:attribute-syntax:directory-string"
)

// All allowed values of EnumdirectoryStringAttributeSyntaxSchemaUrn enum
var AllowedEnumdirectoryStringAttributeSyntaxSchemaUrnEnumValues = []EnumdirectoryStringAttributeSyntaxSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:attribute-syntax:directory-string",
}

func (v *EnumdirectoryStringAttributeSyntaxSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdirectoryStringAttributeSyntaxSchemaUrn(value)
	for _, existing := range AllowedEnumdirectoryStringAttributeSyntaxSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdirectoryStringAttributeSyntaxSchemaUrn", value)
}

// NewEnumdirectoryStringAttributeSyntaxSchemaUrnFromValue returns a pointer to a valid EnumdirectoryStringAttributeSyntaxSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdirectoryStringAttributeSyntaxSchemaUrnFromValue(v string) (*EnumdirectoryStringAttributeSyntaxSchemaUrn, error) {
	ev := EnumdirectoryStringAttributeSyntaxSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdirectoryStringAttributeSyntaxSchemaUrn: valid values are %v", v, AllowedEnumdirectoryStringAttributeSyntaxSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdirectoryStringAttributeSyntaxSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdirectoryStringAttributeSyntaxSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdirectory-string-attribute-syntaxSchemaUrn value
func (v EnumdirectoryStringAttributeSyntaxSchemaUrn) Ptr() *EnumdirectoryStringAttributeSyntaxSchemaUrn {
	return &v
}

type NullableEnumdirectoryStringAttributeSyntaxSchemaUrn struct {
	value *EnumdirectoryStringAttributeSyntaxSchemaUrn
	isSet bool
}

func (v NullableEnumdirectoryStringAttributeSyntaxSchemaUrn) Get() *EnumdirectoryStringAttributeSyntaxSchemaUrn {
	return v.value
}

func (v *NullableEnumdirectoryStringAttributeSyntaxSchemaUrn) Set(val *EnumdirectoryStringAttributeSyntaxSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdirectoryStringAttributeSyntaxSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdirectoryStringAttributeSyntaxSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdirectoryStringAttributeSyntaxSchemaUrn(val *EnumdirectoryStringAttributeSyntaxSchemaUrn) *NullableEnumdirectoryStringAttributeSyntaxSchemaUrn {
	return &NullableEnumdirectoryStringAttributeSyntaxSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdirectoryStringAttributeSyntaxSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdirectoryStringAttributeSyntaxSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

