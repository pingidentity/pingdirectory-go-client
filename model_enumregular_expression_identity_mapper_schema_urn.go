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

// EnumregularExpressionIdentityMapperSchemaUrn the model 'EnumregularExpressionIdentityMapperSchemaUrn'
type EnumregularExpressionIdentityMapperSchemaUrn string

// List of Enumregular-expression-identity-mapperSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0IDENTITY_MAPPERREGULAR_EXPRESSION EnumregularExpressionIdentityMapperSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:identity-mapper:regular-expression"
)

// All allowed values of EnumregularExpressionIdentityMapperSchemaUrn enum
var AllowedEnumregularExpressionIdentityMapperSchemaUrnEnumValues = []EnumregularExpressionIdentityMapperSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:identity-mapper:regular-expression",
}

func (v *EnumregularExpressionIdentityMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumregularExpressionIdentityMapperSchemaUrn(value)
	for _, existing := range AllowedEnumregularExpressionIdentityMapperSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumregularExpressionIdentityMapperSchemaUrn", value)
}

// NewEnumregularExpressionIdentityMapperSchemaUrnFromValue returns a pointer to a valid EnumregularExpressionIdentityMapperSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumregularExpressionIdentityMapperSchemaUrnFromValue(v string) (*EnumregularExpressionIdentityMapperSchemaUrn, error) {
	ev := EnumregularExpressionIdentityMapperSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumregularExpressionIdentityMapperSchemaUrn: valid values are %v", v, AllowedEnumregularExpressionIdentityMapperSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumregularExpressionIdentityMapperSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumregularExpressionIdentityMapperSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumregular-expression-identity-mapperSchemaUrn value
func (v EnumregularExpressionIdentityMapperSchemaUrn) Ptr() *EnumregularExpressionIdentityMapperSchemaUrn {
	return &v
}

type NullableEnumregularExpressionIdentityMapperSchemaUrn struct {
	value *EnumregularExpressionIdentityMapperSchemaUrn
	isSet bool
}

func (v NullableEnumregularExpressionIdentityMapperSchemaUrn) Get() *EnumregularExpressionIdentityMapperSchemaUrn {
	return v.value
}

func (v *NullableEnumregularExpressionIdentityMapperSchemaUrn) Set(val *EnumregularExpressionIdentityMapperSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumregularExpressionIdentityMapperSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumregularExpressionIdentityMapperSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumregularExpressionIdentityMapperSchemaUrn(val *EnumregularExpressionIdentityMapperSchemaUrn) *NullableEnumregularExpressionIdentityMapperSchemaUrn {
	return &NullableEnumregularExpressionIdentityMapperSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumregularExpressionIdentityMapperSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumregularExpressionIdentityMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

