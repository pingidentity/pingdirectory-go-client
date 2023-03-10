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

// EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn the model 'EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn'
type EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn string

// List of Enumgroovy-scripted-uncached-attribute-criteriaSchemaUrn
const (
	ENUMGROOVYSCRIPTEDUNCACHEDATTRIBUTECRITERIASCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0UNCACHED_ATTRIBUTE_CRITERIAGROOVY_SCRIPTED EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:uncached-attribute-criteria:groovy-scripted"
)

// All allowed values of EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn enum
var AllowedEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrnEnumValues = []EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:uncached-attribute-criteria:groovy-scripted",
}

func (v *EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn(value)
	for _, existing := range AllowedEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn", value)
}

// NewEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrnFromValue returns a pointer to a valid EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrnFromValue(v string) (*EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn, error) {
	ev := EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn: valid values are %v", v, AllowedEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgroovy-scripted-uncached-attribute-criteriaSchemaUrn value
func (v EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) Ptr() *EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn {
	return &v
}

type NullableEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn struct {
	value *EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn
	isSet bool
}

func (v NullableEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) Get() *EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn {
	return v.value
}

func (v *NullableEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) Set(val *EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn(val *EnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) *NullableEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn {
	return &NullableEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgroovyScriptedUncachedAttributeCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
