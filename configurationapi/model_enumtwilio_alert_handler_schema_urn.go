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

// EnumtwilioAlertHandlerSchemaUrn the model 'EnumtwilioAlertHandlerSchemaUrn'
type EnumtwilioAlertHandlerSchemaUrn string

// List of Enumtwilio-alert-handlerSchemaUrn
const (
	ENUMTWILIOALERTHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ALERT_HANDLERTWILIO EnumtwilioAlertHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:alert-handler:twilio"
)

// All allowed values of EnumtwilioAlertHandlerSchemaUrn enum
var AllowedEnumtwilioAlertHandlerSchemaUrnEnumValues = []EnumtwilioAlertHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:alert-handler:twilio",
}

func (v *EnumtwilioAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumtwilioAlertHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumtwilioAlertHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumtwilioAlertHandlerSchemaUrn", value)
}

// NewEnumtwilioAlertHandlerSchemaUrnFromValue returns a pointer to a valid EnumtwilioAlertHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumtwilioAlertHandlerSchemaUrnFromValue(v string) (*EnumtwilioAlertHandlerSchemaUrn, error) {
	ev := EnumtwilioAlertHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumtwilioAlertHandlerSchemaUrn: valid values are %v", v, AllowedEnumtwilioAlertHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumtwilioAlertHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumtwilioAlertHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumtwilio-alert-handlerSchemaUrn value
func (v EnumtwilioAlertHandlerSchemaUrn) Ptr() *EnumtwilioAlertHandlerSchemaUrn {
	return &v
}

type NullableEnumtwilioAlertHandlerSchemaUrn struct {
	value *EnumtwilioAlertHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumtwilioAlertHandlerSchemaUrn) Get() *EnumtwilioAlertHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumtwilioAlertHandlerSchemaUrn) Set(val *EnumtwilioAlertHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumtwilioAlertHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumtwilioAlertHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumtwilioAlertHandlerSchemaUrn(val *EnumtwilioAlertHandlerSchemaUrn) *NullableEnumtwilioAlertHandlerSchemaUrn {
	return &NullableEnumtwilioAlertHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumtwilioAlertHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumtwilioAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}