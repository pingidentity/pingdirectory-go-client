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

// EnumrandomPasswordGeneratorSchemaUrn the model 'EnumrandomPasswordGeneratorSchemaUrn'
type EnumrandomPasswordGeneratorSchemaUrn string

// List of Enumrandom-password-generatorSchemaUrn
const (
	ENUMRANDOMPASSWORDGENERATORSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_GENERATORRANDOM EnumrandomPasswordGeneratorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-generator:random"
)

// All allowed values of EnumrandomPasswordGeneratorSchemaUrn enum
var AllowedEnumrandomPasswordGeneratorSchemaUrnEnumValues = []EnumrandomPasswordGeneratorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-generator:random",
}

func (v *EnumrandomPasswordGeneratorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumrandomPasswordGeneratorSchemaUrn(value)
	for _, existing := range AllowedEnumrandomPasswordGeneratorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumrandomPasswordGeneratorSchemaUrn", value)
}

// NewEnumrandomPasswordGeneratorSchemaUrnFromValue returns a pointer to a valid EnumrandomPasswordGeneratorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumrandomPasswordGeneratorSchemaUrnFromValue(v string) (*EnumrandomPasswordGeneratorSchemaUrn, error) {
	ev := EnumrandomPasswordGeneratorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumrandomPasswordGeneratorSchemaUrn: valid values are %v", v, AllowedEnumrandomPasswordGeneratorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumrandomPasswordGeneratorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumrandomPasswordGeneratorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumrandom-password-generatorSchemaUrn value
func (v EnumrandomPasswordGeneratorSchemaUrn) Ptr() *EnumrandomPasswordGeneratorSchemaUrn {
	return &v
}

type NullableEnumrandomPasswordGeneratorSchemaUrn struct {
	value *EnumrandomPasswordGeneratorSchemaUrn
	isSet bool
}

func (v NullableEnumrandomPasswordGeneratorSchemaUrn) Get() *EnumrandomPasswordGeneratorSchemaUrn {
	return v.value
}

func (v *NullableEnumrandomPasswordGeneratorSchemaUrn) Set(val *EnumrandomPasswordGeneratorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumrandomPasswordGeneratorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumrandomPasswordGeneratorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumrandomPasswordGeneratorSchemaUrn(val *EnumrandomPasswordGeneratorSchemaUrn) *NullableEnumrandomPasswordGeneratorSchemaUrn {
	return &NullableEnumrandomPasswordGeneratorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumrandomPasswordGeneratorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumrandomPasswordGeneratorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
