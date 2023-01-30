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

// EnumgroovyScriptedAlertHandlerSchemaUrn the model 'EnumgroovyScriptedAlertHandlerSchemaUrn'
type EnumgroovyScriptedAlertHandlerSchemaUrn string

// List of Enumgroovy-scripted-alert-handlerSchemaUrn
const (
	ENUMGROOVYSCRIPTEDALERTHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ALERT_HANDLERGROOVY_SCRIPTED EnumgroovyScriptedAlertHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:alert-handler:groovy-scripted"
)

// All allowed values of EnumgroovyScriptedAlertHandlerSchemaUrn enum
var AllowedEnumgroovyScriptedAlertHandlerSchemaUrnEnumValues = []EnumgroovyScriptedAlertHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:alert-handler:groovy-scripted",
}

func (v *EnumgroovyScriptedAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgroovyScriptedAlertHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumgroovyScriptedAlertHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgroovyScriptedAlertHandlerSchemaUrn", value)
}

// NewEnumgroovyScriptedAlertHandlerSchemaUrnFromValue returns a pointer to a valid EnumgroovyScriptedAlertHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgroovyScriptedAlertHandlerSchemaUrnFromValue(v string) (*EnumgroovyScriptedAlertHandlerSchemaUrn, error) {
	ev := EnumgroovyScriptedAlertHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgroovyScriptedAlertHandlerSchemaUrn: valid values are %v", v, AllowedEnumgroovyScriptedAlertHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgroovyScriptedAlertHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgroovyScriptedAlertHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgroovy-scripted-alert-handlerSchemaUrn value
func (v EnumgroovyScriptedAlertHandlerSchemaUrn) Ptr() *EnumgroovyScriptedAlertHandlerSchemaUrn {
	return &v
}

type NullableEnumgroovyScriptedAlertHandlerSchemaUrn struct {
	value *EnumgroovyScriptedAlertHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumgroovyScriptedAlertHandlerSchemaUrn) Get() *EnumgroovyScriptedAlertHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumgroovyScriptedAlertHandlerSchemaUrn) Set(val *EnumgroovyScriptedAlertHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgroovyScriptedAlertHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgroovyScriptedAlertHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgroovyScriptedAlertHandlerSchemaUrn(val *EnumgroovyScriptedAlertHandlerSchemaUrn) *NullableEnumgroovyScriptedAlertHandlerSchemaUrn {
	return &NullableEnumgroovyScriptedAlertHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgroovyScriptedAlertHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgroovyScriptedAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
