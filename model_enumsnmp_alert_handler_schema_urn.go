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

// EnumsnmpAlertHandlerSchemaUrn the model 'EnumsnmpAlertHandlerSchemaUrn'
type EnumsnmpAlertHandlerSchemaUrn string

// List of Enumsnmp-alert-handlerSchemaUrn
const (
	ENUMSNMPALERTHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ALERT_HANDLERSNMP EnumsnmpAlertHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:alert-handler:snmp"
)

// All allowed values of EnumsnmpAlertHandlerSchemaUrn enum
var AllowedEnumsnmpAlertHandlerSchemaUrnEnumValues = []EnumsnmpAlertHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:alert-handler:snmp",
}

func (v *EnumsnmpAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsnmpAlertHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumsnmpAlertHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsnmpAlertHandlerSchemaUrn", value)
}

// NewEnumsnmpAlertHandlerSchemaUrnFromValue returns a pointer to a valid EnumsnmpAlertHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsnmpAlertHandlerSchemaUrnFromValue(v string) (*EnumsnmpAlertHandlerSchemaUrn, error) {
	ev := EnumsnmpAlertHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsnmpAlertHandlerSchemaUrn: valid values are %v", v, AllowedEnumsnmpAlertHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsnmpAlertHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsnmpAlertHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsnmp-alert-handlerSchemaUrn value
func (v EnumsnmpAlertHandlerSchemaUrn) Ptr() *EnumsnmpAlertHandlerSchemaUrn {
	return &v
}

type NullableEnumsnmpAlertHandlerSchemaUrn struct {
	value *EnumsnmpAlertHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumsnmpAlertHandlerSchemaUrn) Get() *EnumsnmpAlertHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumsnmpAlertHandlerSchemaUrn) Set(val *EnumsnmpAlertHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsnmpAlertHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsnmpAlertHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsnmpAlertHandlerSchemaUrn(val *EnumsnmpAlertHandlerSchemaUrn) *NullableEnumsnmpAlertHandlerSchemaUrn {
	return &NullableEnumsnmpAlertHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsnmpAlertHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsnmpAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
