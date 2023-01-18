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

// EnumgroovyScriptedAccessLogPublisherSchemaUrn the model 'EnumgroovyScriptedAccessLogPublisherSchemaUrn'
type EnumgroovyScriptedAccessLogPublisherSchemaUrn string

// List of Enumgroovy-scripted-access-log-publisherSchemaUrn
const (
	ENUMGROOVYSCRIPTEDACCESSLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERGROOVY_SCRIPTED_ACCESS EnumgroovyScriptedAccessLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:groovy-scripted-access"
)

// All allowed values of EnumgroovyScriptedAccessLogPublisherSchemaUrn enum
var AllowedEnumgroovyScriptedAccessLogPublisherSchemaUrnEnumValues = []EnumgroovyScriptedAccessLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:groovy-scripted-access",
}

func (v *EnumgroovyScriptedAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgroovyScriptedAccessLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumgroovyScriptedAccessLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgroovyScriptedAccessLogPublisherSchemaUrn", value)
}

// NewEnumgroovyScriptedAccessLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumgroovyScriptedAccessLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgroovyScriptedAccessLogPublisherSchemaUrnFromValue(v string) (*EnumgroovyScriptedAccessLogPublisherSchemaUrn, error) {
	ev := EnumgroovyScriptedAccessLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgroovyScriptedAccessLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumgroovyScriptedAccessLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgroovyScriptedAccessLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgroovyScriptedAccessLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgroovy-scripted-access-log-publisherSchemaUrn value
func (v EnumgroovyScriptedAccessLogPublisherSchemaUrn) Ptr() *EnumgroovyScriptedAccessLogPublisherSchemaUrn {
	return &v
}

type NullableEnumgroovyScriptedAccessLogPublisherSchemaUrn struct {
	value *EnumgroovyScriptedAccessLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumgroovyScriptedAccessLogPublisherSchemaUrn) Get() *EnumgroovyScriptedAccessLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumgroovyScriptedAccessLogPublisherSchemaUrn) Set(val *EnumgroovyScriptedAccessLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgroovyScriptedAccessLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgroovyScriptedAccessLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgroovyScriptedAccessLogPublisherSchemaUrn(val *EnumgroovyScriptedAccessLogPublisherSchemaUrn) *NullableEnumgroovyScriptedAccessLogPublisherSchemaUrn {
	return &NullableEnumgroovyScriptedAccessLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgroovyScriptedAccessLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgroovyScriptedAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
