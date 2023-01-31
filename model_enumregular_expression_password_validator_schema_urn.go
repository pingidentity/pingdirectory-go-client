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

// EnumregularExpressionPasswordValidatorSchemaUrn the model 'EnumregularExpressionPasswordValidatorSchemaUrn'
type EnumregularExpressionPasswordValidatorSchemaUrn string

// List of Enumregular-expression-password-validatorSchemaUrn
const (
	ENUMREGULAREXPRESSIONPASSWORDVALIDATORSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_VALIDATORREGULAR_EXPRESSION EnumregularExpressionPasswordValidatorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-validator:regular-expression"
)

// All allowed values of EnumregularExpressionPasswordValidatorSchemaUrn enum
var AllowedEnumregularExpressionPasswordValidatorSchemaUrnEnumValues = []EnumregularExpressionPasswordValidatorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-validator:regular-expression",
}

func (v *EnumregularExpressionPasswordValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumregularExpressionPasswordValidatorSchemaUrn(value)
	for _, existing := range AllowedEnumregularExpressionPasswordValidatorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumregularExpressionPasswordValidatorSchemaUrn", value)
}

// NewEnumregularExpressionPasswordValidatorSchemaUrnFromValue returns a pointer to a valid EnumregularExpressionPasswordValidatorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumregularExpressionPasswordValidatorSchemaUrnFromValue(v string) (*EnumregularExpressionPasswordValidatorSchemaUrn, error) {
	ev := EnumregularExpressionPasswordValidatorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumregularExpressionPasswordValidatorSchemaUrn: valid values are %v", v, AllowedEnumregularExpressionPasswordValidatorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumregularExpressionPasswordValidatorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumregularExpressionPasswordValidatorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumregular-expression-password-validatorSchemaUrn value
func (v EnumregularExpressionPasswordValidatorSchemaUrn) Ptr() *EnumregularExpressionPasswordValidatorSchemaUrn {
	return &v
}

type NullableEnumregularExpressionPasswordValidatorSchemaUrn struct {
	value *EnumregularExpressionPasswordValidatorSchemaUrn
	isSet bool
}

func (v NullableEnumregularExpressionPasswordValidatorSchemaUrn) Get() *EnumregularExpressionPasswordValidatorSchemaUrn {
	return v.value
}

func (v *NullableEnumregularExpressionPasswordValidatorSchemaUrn) Set(val *EnumregularExpressionPasswordValidatorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumregularExpressionPasswordValidatorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumregularExpressionPasswordValidatorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumregularExpressionPasswordValidatorSchemaUrn(val *EnumregularExpressionPasswordValidatorSchemaUrn) *NullableEnumregularExpressionPasswordValidatorSchemaUrn {
	return &NullableEnumregularExpressionPasswordValidatorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumregularExpressionPasswordValidatorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumregularExpressionPasswordValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
