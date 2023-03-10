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

// EnumjmxAlertHandlerSchemaUrn the model 'EnumjmxAlertHandlerSchemaUrn'
type EnumjmxAlertHandlerSchemaUrn string

// List of Enumjmx-alert-handlerSchemaUrn
const (
	ENUMJMXALERTHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ALERT_HANDLERJMX EnumjmxAlertHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:alert-handler:jmx"
)

// All allowed values of EnumjmxAlertHandlerSchemaUrn enum
var AllowedEnumjmxAlertHandlerSchemaUrnEnumValues = []EnumjmxAlertHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:alert-handler:jmx",
}

func (v *EnumjmxAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumjmxAlertHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumjmxAlertHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumjmxAlertHandlerSchemaUrn", value)
}

// NewEnumjmxAlertHandlerSchemaUrnFromValue returns a pointer to a valid EnumjmxAlertHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumjmxAlertHandlerSchemaUrnFromValue(v string) (*EnumjmxAlertHandlerSchemaUrn, error) {
	ev := EnumjmxAlertHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumjmxAlertHandlerSchemaUrn: valid values are %v", v, AllowedEnumjmxAlertHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumjmxAlertHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumjmxAlertHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumjmx-alert-handlerSchemaUrn value
func (v EnumjmxAlertHandlerSchemaUrn) Ptr() *EnumjmxAlertHandlerSchemaUrn {
	return &v
}

type NullableEnumjmxAlertHandlerSchemaUrn struct {
	value *EnumjmxAlertHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumjmxAlertHandlerSchemaUrn) Get() *EnumjmxAlertHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumjmxAlertHandlerSchemaUrn) Set(val *EnumjmxAlertHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumjmxAlertHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumjmxAlertHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumjmxAlertHandlerSchemaUrn(val *EnumjmxAlertHandlerSchemaUrn) *NullableEnumjmxAlertHandlerSchemaUrn {
	return &NullableEnumjmxAlertHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumjmxAlertHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumjmxAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
