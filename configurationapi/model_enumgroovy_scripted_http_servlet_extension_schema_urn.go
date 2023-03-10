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

// EnumgroovyScriptedHttpServletExtensionSchemaUrn the model 'EnumgroovyScriptedHttpServletExtensionSchemaUrn'
type EnumgroovyScriptedHttpServletExtensionSchemaUrn string

// List of Enumgroovy-scripted-http-servlet-extensionSchemaUrn
const (
	ENUMGROOVYSCRIPTEDHTTPSERVLETEXTENSIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0HTTP_SERVLET_EXTENSIONGROOVY_SCRIPTED EnumgroovyScriptedHttpServletExtensionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:groovy-scripted"
)

// All allowed values of EnumgroovyScriptedHttpServletExtensionSchemaUrn enum
var AllowedEnumgroovyScriptedHttpServletExtensionSchemaUrnEnumValues = []EnumgroovyScriptedHttpServletExtensionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:groovy-scripted",
}

func (v *EnumgroovyScriptedHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgroovyScriptedHttpServletExtensionSchemaUrn(value)
	for _, existing := range AllowedEnumgroovyScriptedHttpServletExtensionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgroovyScriptedHttpServletExtensionSchemaUrn", value)
}

// NewEnumgroovyScriptedHttpServletExtensionSchemaUrnFromValue returns a pointer to a valid EnumgroovyScriptedHttpServletExtensionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgroovyScriptedHttpServletExtensionSchemaUrnFromValue(v string) (*EnumgroovyScriptedHttpServletExtensionSchemaUrn, error) {
	ev := EnumgroovyScriptedHttpServletExtensionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgroovyScriptedHttpServletExtensionSchemaUrn: valid values are %v", v, AllowedEnumgroovyScriptedHttpServletExtensionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgroovyScriptedHttpServletExtensionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgroovyScriptedHttpServletExtensionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgroovy-scripted-http-servlet-extensionSchemaUrn value
func (v EnumgroovyScriptedHttpServletExtensionSchemaUrn) Ptr() *EnumgroovyScriptedHttpServletExtensionSchemaUrn {
	return &v
}

type NullableEnumgroovyScriptedHttpServletExtensionSchemaUrn struct {
	value *EnumgroovyScriptedHttpServletExtensionSchemaUrn
	isSet bool
}

func (v NullableEnumgroovyScriptedHttpServletExtensionSchemaUrn) Get() *EnumgroovyScriptedHttpServletExtensionSchemaUrn {
	return v.value
}

func (v *NullableEnumgroovyScriptedHttpServletExtensionSchemaUrn) Set(val *EnumgroovyScriptedHttpServletExtensionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgroovyScriptedHttpServletExtensionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgroovyScriptedHttpServletExtensionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgroovyScriptedHttpServletExtensionSchemaUrn(val *EnumgroovyScriptedHttpServletExtensionSchemaUrn) *NullableEnumgroovyScriptedHttpServletExtensionSchemaUrn {
	return &NullableEnumgroovyScriptedHttpServletExtensionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgroovyScriptedHttpServletExtensionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgroovyScriptedHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
