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

// EnumgroovyScriptedErrorLogPublisherSchemaUrn the model 'EnumgroovyScriptedErrorLogPublisherSchemaUrn'
type EnumgroovyScriptedErrorLogPublisherSchemaUrn string

// List of Enumgroovy-scripted-error-log-publisherSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERGROOVY_SCRIPTED_ERROR EnumgroovyScriptedErrorLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:groovy-scripted-error"
)

// All allowed values of EnumgroovyScriptedErrorLogPublisherSchemaUrn enum
var AllowedEnumgroovyScriptedErrorLogPublisherSchemaUrnEnumValues = []EnumgroovyScriptedErrorLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:groovy-scripted-error",
}

func (v *EnumgroovyScriptedErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgroovyScriptedErrorLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumgroovyScriptedErrorLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgroovyScriptedErrorLogPublisherSchemaUrn", value)
}

// NewEnumgroovyScriptedErrorLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumgroovyScriptedErrorLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgroovyScriptedErrorLogPublisherSchemaUrnFromValue(v string) (*EnumgroovyScriptedErrorLogPublisherSchemaUrn, error) {
	ev := EnumgroovyScriptedErrorLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgroovyScriptedErrorLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumgroovyScriptedErrorLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgroovyScriptedErrorLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgroovyScriptedErrorLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgroovy-scripted-error-log-publisherSchemaUrn value
func (v EnumgroovyScriptedErrorLogPublisherSchemaUrn) Ptr() *EnumgroovyScriptedErrorLogPublisherSchemaUrn {
	return &v
}

type NullableEnumgroovyScriptedErrorLogPublisherSchemaUrn struct {
	value *EnumgroovyScriptedErrorLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumgroovyScriptedErrorLogPublisherSchemaUrn) Get() *EnumgroovyScriptedErrorLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumgroovyScriptedErrorLogPublisherSchemaUrn) Set(val *EnumgroovyScriptedErrorLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgroovyScriptedErrorLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgroovyScriptedErrorLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgroovyScriptedErrorLogPublisherSchemaUrn(val *EnumgroovyScriptedErrorLogPublisherSchemaUrn) *NullableEnumgroovyScriptedErrorLogPublisherSchemaUrn {
	return &NullableEnumgroovyScriptedErrorLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgroovyScriptedErrorLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgroovyScriptedErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

