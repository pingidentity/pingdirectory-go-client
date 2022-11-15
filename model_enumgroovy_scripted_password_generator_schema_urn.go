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

// EnumgroovyScriptedPasswordGeneratorSchemaUrn the model 'EnumgroovyScriptedPasswordGeneratorSchemaUrn'
type EnumgroovyScriptedPasswordGeneratorSchemaUrn string

// List of Enumgroovy-scripted-password-generatorSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSWORD_GENERATORGROOVY_SCRIPTED EnumgroovyScriptedPasswordGeneratorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:password-generator:groovy-scripted"
)

// All allowed values of EnumgroovyScriptedPasswordGeneratorSchemaUrn enum
var AllowedEnumgroovyScriptedPasswordGeneratorSchemaUrnEnumValues = []EnumgroovyScriptedPasswordGeneratorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:password-generator:groovy-scripted",
}

func (v *EnumgroovyScriptedPasswordGeneratorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgroovyScriptedPasswordGeneratorSchemaUrn(value)
	for _, existing := range AllowedEnumgroovyScriptedPasswordGeneratorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgroovyScriptedPasswordGeneratorSchemaUrn", value)
}

// NewEnumgroovyScriptedPasswordGeneratorSchemaUrnFromValue returns a pointer to a valid EnumgroovyScriptedPasswordGeneratorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgroovyScriptedPasswordGeneratorSchemaUrnFromValue(v string) (*EnumgroovyScriptedPasswordGeneratorSchemaUrn, error) {
	ev := EnumgroovyScriptedPasswordGeneratorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgroovyScriptedPasswordGeneratorSchemaUrn: valid values are %v", v, AllowedEnumgroovyScriptedPasswordGeneratorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgroovyScriptedPasswordGeneratorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgroovyScriptedPasswordGeneratorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgroovy-scripted-password-generatorSchemaUrn value
func (v EnumgroovyScriptedPasswordGeneratorSchemaUrn) Ptr() *EnumgroovyScriptedPasswordGeneratorSchemaUrn {
	return &v
}

type NullableEnumgroovyScriptedPasswordGeneratorSchemaUrn struct {
	value *EnumgroovyScriptedPasswordGeneratorSchemaUrn
	isSet bool
}

func (v NullableEnumgroovyScriptedPasswordGeneratorSchemaUrn) Get() *EnumgroovyScriptedPasswordGeneratorSchemaUrn {
	return v.value
}

func (v *NullableEnumgroovyScriptedPasswordGeneratorSchemaUrn) Set(val *EnumgroovyScriptedPasswordGeneratorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgroovyScriptedPasswordGeneratorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgroovyScriptedPasswordGeneratorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgroovyScriptedPasswordGeneratorSchemaUrn(val *EnumgroovyScriptedPasswordGeneratorSchemaUrn) *NullableEnumgroovyScriptedPasswordGeneratorSchemaUrn {
	return &NullableEnumgroovyScriptedPasswordGeneratorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgroovyScriptedPasswordGeneratorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgroovyScriptedPasswordGeneratorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

