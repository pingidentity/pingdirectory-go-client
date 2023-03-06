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

// EnumcustomPasswordValidatorSchemaUrn the model 'EnumcustomPasswordValidatorSchemaUrn'
type EnumcustomPasswordValidatorSchemaUrn string

// List of Enumcustom-password-validatorSchemaUrn
const (
	ENUMCUSTOMPASSWORDVALIDATORSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_VALIDATORCUSTOM EnumcustomPasswordValidatorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-validator:custom"
)

// All allowed values of EnumcustomPasswordValidatorSchemaUrn enum
var AllowedEnumcustomPasswordValidatorSchemaUrnEnumValues = []EnumcustomPasswordValidatorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-validator:custom",
}

func (v *EnumcustomPasswordValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcustomPasswordValidatorSchemaUrn(value)
	for _, existing := range AllowedEnumcustomPasswordValidatorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcustomPasswordValidatorSchemaUrn", value)
}

// NewEnumcustomPasswordValidatorSchemaUrnFromValue returns a pointer to a valid EnumcustomPasswordValidatorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcustomPasswordValidatorSchemaUrnFromValue(v string) (*EnumcustomPasswordValidatorSchemaUrn, error) {
	ev := EnumcustomPasswordValidatorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcustomPasswordValidatorSchemaUrn: valid values are %v", v, AllowedEnumcustomPasswordValidatorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcustomPasswordValidatorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcustomPasswordValidatorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumcustom-password-validatorSchemaUrn value
func (v EnumcustomPasswordValidatorSchemaUrn) Ptr() *EnumcustomPasswordValidatorSchemaUrn {
	return &v
}

type NullableEnumcustomPasswordValidatorSchemaUrn struct {
	value *EnumcustomPasswordValidatorSchemaUrn
	isSet bool
}

func (v NullableEnumcustomPasswordValidatorSchemaUrn) Get() *EnumcustomPasswordValidatorSchemaUrn {
	return v.value
}

func (v *NullableEnumcustomPasswordValidatorSchemaUrn) Set(val *EnumcustomPasswordValidatorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcustomPasswordValidatorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcustomPasswordValidatorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcustomPasswordValidatorSchemaUrn(val *EnumcustomPasswordValidatorSchemaUrn) *NullableEnumcustomPasswordValidatorSchemaUrn {
	return &NullableEnumcustomPasswordValidatorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcustomPasswordValidatorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcustomPasswordValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
