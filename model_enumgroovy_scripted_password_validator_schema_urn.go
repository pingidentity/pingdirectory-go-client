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

// EnumgroovyScriptedPasswordValidatorSchemaUrn the model 'EnumgroovyScriptedPasswordValidatorSchemaUrn'
type EnumgroovyScriptedPasswordValidatorSchemaUrn string

// List of Enumgroovy-scripted-password-validatorSchemaUrn
const (
	ENUMGROOVYSCRIPTEDPASSWORDVALIDATORSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_VALIDATORGROOVY_SCRIPTED EnumgroovyScriptedPasswordValidatorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-validator:groovy-scripted"
)

// All allowed values of EnumgroovyScriptedPasswordValidatorSchemaUrn enum
var AllowedEnumgroovyScriptedPasswordValidatorSchemaUrnEnumValues = []EnumgroovyScriptedPasswordValidatorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-validator:groovy-scripted",
}

func (v *EnumgroovyScriptedPasswordValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgroovyScriptedPasswordValidatorSchemaUrn(value)
	for _, existing := range AllowedEnumgroovyScriptedPasswordValidatorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgroovyScriptedPasswordValidatorSchemaUrn", value)
}

// NewEnumgroovyScriptedPasswordValidatorSchemaUrnFromValue returns a pointer to a valid EnumgroovyScriptedPasswordValidatorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgroovyScriptedPasswordValidatorSchemaUrnFromValue(v string) (*EnumgroovyScriptedPasswordValidatorSchemaUrn, error) {
	ev := EnumgroovyScriptedPasswordValidatorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgroovyScriptedPasswordValidatorSchemaUrn: valid values are %v", v, AllowedEnumgroovyScriptedPasswordValidatorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgroovyScriptedPasswordValidatorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgroovyScriptedPasswordValidatorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgroovy-scripted-password-validatorSchemaUrn value
func (v EnumgroovyScriptedPasswordValidatorSchemaUrn) Ptr() *EnumgroovyScriptedPasswordValidatorSchemaUrn {
	return &v
}

type NullableEnumgroovyScriptedPasswordValidatorSchemaUrn struct {
	value *EnumgroovyScriptedPasswordValidatorSchemaUrn
	isSet bool
}

func (v NullableEnumgroovyScriptedPasswordValidatorSchemaUrn) Get() *EnumgroovyScriptedPasswordValidatorSchemaUrn {
	return v.value
}

func (v *NullableEnumgroovyScriptedPasswordValidatorSchemaUrn) Set(val *EnumgroovyScriptedPasswordValidatorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgroovyScriptedPasswordValidatorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgroovyScriptedPasswordValidatorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgroovyScriptedPasswordValidatorSchemaUrn(val *EnumgroovyScriptedPasswordValidatorSchemaUrn) *NullableEnumgroovyScriptedPasswordValidatorSchemaUrn {
	return &NullableEnumgroovyScriptedPasswordValidatorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgroovyScriptedPasswordValidatorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgroovyScriptedPasswordValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

