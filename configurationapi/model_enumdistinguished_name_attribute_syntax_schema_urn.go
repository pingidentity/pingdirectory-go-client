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

// EnumdistinguishedNameAttributeSyntaxSchemaUrn the model 'EnumdistinguishedNameAttributeSyntaxSchemaUrn'
type EnumdistinguishedNameAttributeSyntaxSchemaUrn string

// List of Enumdistinguished-name-attribute-syntaxSchemaUrn
const (
	ENUMDISTINGUISHEDNAMEATTRIBUTESYNTAXSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ATTRIBUTE_SYNTAXDISTINGUISHED_NAME EnumdistinguishedNameAttributeSyntaxSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:attribute-syntax:distinguished-name"
)

// All allowed values of EnumdistinguishedNameAttributeSyntaxSchemaUrn enum
var AllowedEnumdistinguishedNameAttributeSyntaxSchemaUrnEnumValues = []EnumdistinguishedNameAttributeSyntaxSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:attribute-syntax:distinguished-name",
}

func (v *EnumdistinguishedNameAttributeSyntaxSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdistinguishedNameAttributeSyntaxSchemaUrn(value)
	for _, existing := range AllowedEnumdistinguishedNameAttributeSyntaxSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdistinguishedNameAttributeSyntaxSchemaUrn", value)
}

// NewEnumdistinguishedNameAttributeSyntaxSchemaUrnFromValue returns a pointer to a valid EnumdistinguishedNameAttributeSyntaxSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdistinguishedNameAttributeSyntaxSchemaUrnFromValue(v string) (*EnumdistinguishedNameAttributeSyntaxSchemaUrn, error) {
	ev := EnumdistinguishedNameAttributeSyntaxSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdistinguishedNameAttributeSyntaxSchemaUrn: valid values are %v", v, AllowedEnumdistinguishedNameAttributeSyntaxSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdistinguishedNameAttributeSyntaxSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdistinguishedNameAttributeSyntaxSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdistinguished-name-attribute-syntaxSchemaUrn value
func (v EnumdistinguishedNameAttributeSyntaxSchemaUrn) Ptr() *EnumdistinguishedNameAttributeSyntaxSchemaUrn {
	return &v
}

type NullableEnumdistinguishedNameAttributeSyntaxSchemaUrn struct {
	value *EnumdistinguishedNameAttributeSyntaxSchemaUrn
	isSet bool
}

func (v NullableEnumdistinguishedNameAttributeSyntaxSchemaUrn) Get() *EnumdistinguishedNameAttributeSyntaxSchemaUrn {
	return v.value
}

func (v *NullableEnumdistinguishedNameAttributeSyntaxSchemaUrn) Set(val *EnumdistinguishedNameAttributeSyntaxSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdistinguishedNameAttributeSyntaxSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdistinguishedNameAttributeSyntaxSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdistinguishedNameAttributeSyntaxSchemaUrn(val *EnumdistinguishedNameAttributeSyntaxSchemaUrn) *NullableEnumdistinguishedNameAttributeSyntaxSchemaUrn {
	return &NullableEnumdistinguishedNameAttributeSyntaxSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdistinguishedNameAttributeSyntaxSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdistinguishedNameAttributeSyntaxSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
