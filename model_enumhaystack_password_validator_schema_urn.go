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

// EnumhaystackPasswordValidatorSchemaUrn the model 'EnumhaystackPasswordValidatorSchemaUrn'
type EnumhaystackPasswordValidatorSchemaUrn string

// List of Enumhaystack-password-validatorSchemaUrn
const (
	ENUMHAYSTACKPASSWORDVALIDATORSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_VALIDATORHAYSTACK EnumhaystackPasswordValidatorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-validator:haystack"
)

// All allowed values of EnumhaystackPasswordValidatorSchemaUrn enum
var AllowedEnumhaystackPasswordValidatorSchemaUrnEnumValues = []EnumhaystackPasswordValidatorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-validator:haystack",
}

func (v *EnumhaystackPasswordValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumhaystackPasswordValidatorSchemaUrn(value)
	for _, existing := range AllowedEnumhaystackPasswordValidatorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumhaystackPasswordValidatorSchemaUrn", value)
}

// NewEnumhaystackPasswordValidatorSchemaUrnFromValue returns a pointer to a valid EnumhaystackPasswordValidatorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumhaystackPasswordValidatorSchemaUrnFromValue(v string) (*EnumhaystackPasswordValidatorSchemaUrn, error) {
	ev := EnumhaystackPasswordValidatorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumhaystackPasswordValidatorSchemaUrn: valid values are %v", v, AllowedEnumhaystackPasswordValidatorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumhaystackPasswordValidatorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumhaystackPasswordValidatorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumhaystack-password-validatorSchemaUrn value
func (v EnumhaystackPasswordValidatorSchemaUrn) Ptr() *EnumhaystackPasswordValidatorSchemaUrn {
	return &v
}

type NullableEnumhaystackPasswordValidatorSchemaUrn struct {
	value *EnumhaystackPasswordValidatorSchemaUrn
	isSet bool
}

func (v NullableEnumhaystackPasswordValidatorSchemaUrn) Get() *EnumhaystackPasswordValidatorSchemaUrn {
	return v.value
}

func (v *NullableEnumhaystackPasswordValidatorSchemaUrn) Set(val *EnumhaystackPasswordValidatorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumhaystackPasswordValidatorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumhaystackPasswordValidatorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumhaystackPasswordValidatorSchemaUrn(val *EnumhaystackPasswordValidatorSchemaUrn) *NullableEnumhaystackPasswordValidatorSchemaUrn {
	return &NullableEnumhaystackPasswordValidatorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumhaystackPasswordValidatorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumhaystackPasswordValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
