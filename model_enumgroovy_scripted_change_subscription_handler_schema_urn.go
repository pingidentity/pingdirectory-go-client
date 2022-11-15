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

// EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn the model 'EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn'
type EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn string

// List of Enumgroovy-scripted-change-subscription-handlerSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0CHANGE_SUBSCRIPTION_HANDLERGROOVY_SCRIPTED EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:change-subscription-handler:groovy-scripted"
)

// All allowed values of EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn enum
var AllowedEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrnEnumValues = []EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:change-subscription-handler:groovy-scripted",
}

func (v *EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn", value)
}

// NewEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrnFromValue returns a pointer to a valid EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrnFromValue(v string) (*EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn, error) {
	ev := EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn: valid values are %v", v, AllowedEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgroovy-scripted-change-subscription-handlerSchemaUrn value
func (v EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) Ptr() *EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn {
	return &v
}

type NullableEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn struct {
	value *EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) Get() *EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) Set(val *EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn(val *EnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) *NullableEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn {
	return &NullableEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgroovyScriptedChangeSubscriptionHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

