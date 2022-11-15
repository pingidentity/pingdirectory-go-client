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

// EnumcharacterSetPasswordValidatorSchemaUrn the model 'EnumcharacterSetPasswordValidatorSchemaUrn'
type EnumcharacterSetPasswordValidatorSchemaUrn string

// List of Enumcharacter-set-password-validatorSchemaUrn
const (
	ENUMCHARACTERSETPASSWORDVALIDATORSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_VALIDATORCHARACTER_SET EnumcharacterSetPasswordValidatorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-validator:character-set"
)

// All allowed values of EnumcharacterSetPasswordValidatorSchemaUrn enum
var AllowedEnumcharacterSetPasswordValidatorSchemaUrnEnumValues = []EnumcharacterSetPasswordValidatorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-validator:character-set",
}

func (v *EnumcharacterSetPasswordValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcharacterSetPasswordValidatorSchemaUrn(value)
	for _, existing := range AllowedEnumcharacterSetPasswordValidatorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcharacterSetPasswordValidatorSchemaUrn", value)
}

// NewEnumcharacterSetPasswordValidatorSchemaUrnFromValue returns a pointer to a valid EnumcharacterSetPasswordValidatorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcharacterSetPasswordValidatorSchemaUrnFromValue(v string) (*EnumcharacterSetPasswordValidatorSchemaUrn, error) {
	ev := EnumcharacterSetPasswordValidatorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcharacterSetPasswordValidatorSchemaUrn: valid values are %v", v, AllowedEnumcharacterSetPasswordValidatorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcharacterSetPasswordValidatorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcharacterSetPasswordValidatorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumcharacter-set-password-validatorSchemaUrn value
func (v EnumcharacterSetPasswordValidatorSchemaUrn) Ptr() *EnumcharacterSetPasswordValidatorSchemaUrn {
	return &v
}

type NullableEnumcharacterSetPasswordValidatorSchemaUrn struct {
	value *EnumcharacterSetPasswordValidatorSchemaUrn
	isSet bool
}

func (v NullableEnumcharacterSetPasswordValidatorSchemaUrn) Get() *EnumcharacterSetPasswordValidatorSchemaUrn {
	return v.value
}

func (v *NullableEnumcharacterSetPasswordValidatorSchemaUrn) Set(val *EnumcharacterSetPasswordValidatorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcharacterSetPasswordValidatorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcharacterSetPasswordValidatorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcharacterSetPasswordValidatorSchemaUrn(val *EnumcharacterSetPasswordValidatorSchemaUrn) *NullableEnumcharacterSetPasswordValidatorSchemaUrn {
	return &NullableEnumcharacterSetPasswordValidatorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcharacterSetPasswordValidatorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcharacterSetPasswordValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

