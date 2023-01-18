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

// EnumsnmpSubAgentAlertHandlerSchemaUrn the model 'EnumsnmpSubAgentAlertHandlerSchemaUrn'
type EnumsnmpSubAgentAlertHandlerSchemaUrn string

// List of Enumsnmp-sub-agent-alert-handlerSchemaUrn
const (
	ENUMSNMPSUBAGENTALERTHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ALERT_HANDLERSNMP_SUB_AGENT EnumsnmpSubAgentAlertHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:alert-handler:snmp-sub-agent"
)

// All allowed values of EnumsnmpSubAgentAlertHandlerSchemaUrn enum
var AllowedEnumsnmpSubAgentAlertHandlerSchemaUrnEnumValues = []EnumsnmpSubAgentAlertHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:alert-handler:snmp-sub-agent",
}

func (v *EnumsnmpSubAgentAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsnmpSubAgentAlertHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumsnmpSubAgentAlertHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsnmpSubAgentAlertHandlerSchemaUrn", value)
}

// NewEnumsnmpSubAgentAlertHandlerSchemaUrnFromValue returns a pointer to a valid EnumsnmpSubAgentAlertHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsnmpSubAgentAlertHandlerSchemaUrnFromValue(v string) (*EnumsnmpSubAgentAlertHandlerSchemaUrn, error) {
	ev := EnumsnmpSubAgentAlertHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsnmpSubAgentAlertHandlerSchemaUrn: valid values are %v", v, AllowedEnumsnmpSubAgentAlertHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsnmpSubAgentAlertHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsnmpSubAgentAlertHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsnmp-sub-agent-alert-handlerSchemaUrn value
func (v EnumsnmpSubAgentAlertHandlerSchemaUrn) Ptr() *EnumsnmpSubAgentAlertHandlerSchemaUrn {
	return &v
}

type NullableEnumsnmpSubAgentAlertHandlerSchemaUrn struct {
	value *EnumsnmpSubAgentAlertHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumsnmpSubAgentAlertHandlerSchemaUrn) Get() *EnumsnmpSubAgentAlertHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumsnmpSubAgentAlertHandlerSchemaUrn) Set(val *EnumsnmpSubAgentAlertHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsnmpSubAgentAlertHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsnmpSubAgentAlertHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsnmpSubAgentAlertHandlerSchemaUrn(val *EnumsnmpSubAgentAlertHandlerSchemaUrn) *NullableEnumsnmpSubAgentAlertHandlerSchemaUrn {
	return &NullableEnumsnmpSubAgentAlertHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsnmpSubAgentAlertHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsnmpSubAgentAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
