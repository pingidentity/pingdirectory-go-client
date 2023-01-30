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

// EnumsmtpAlertHandlerSchemaUrn the model 'EnumsmtpAlertHandlerSchemaUrn'
type EnumsmtpAlertHandlerSchemaUrn string

// List of Enumsmtp-alert-handlerSchemaUrn
const (
	ENUMSMTPALERTHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ALERT_HANDLERSMTP EnumsmtpAlertHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:alert-handler:smtp"
)

// All allowed values of EnumsmtpAlertHandlerSchemaUrn enum
var AllowedEnumsmtpAlertHandlerSchemaUrnEnumValues = []EnumsmtpAlertHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:alert-handler:smtp",
}

func (v *EnumsmtpAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsmtpAlertHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumsmtpAlertHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsmtpAlertHandlerSchemaUrn", value)
}

// NewEnumsmtpAlertHandlerSchemaUrnFromValue returns a pointer to a valid EnumsmtpAlertHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsmtpAlertHandlerSchemaUrnFromValue(v string) (*EnumsmtpAlertHandlerSchemaUrn, error) {
	ev := EnumsmtpAlertHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsmtpAlertHandlerSchemaUrn: valid values are %v", v, AllowedEnumsmtpAlertHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsmtpAlertHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsmtpAlertHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsmtp-alert-handlerSchemaUrn value
func (v EnumsmtpAlertHandlerSchemaUrn) Ptr() *EnumsmtpAlertHandlerSchemaUrn {
	return &v
}

type NullableEnumsmtpAlertHandlerSchemaUrn struct {
	value *EnumsmtpAlertHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumsmtpAlertHandlerSchemaUrn) Get() *EnumsmtpAlertHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumsmtpAlertHandlerSchemaUrn) Set(val *EnumsmtpAlertHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsmtpAlertHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsmtpAlertHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsmtpAlertHandlerSchemaUrn(val *EnumsmtpAlertHandlerSchemaUrn) *NullableEnumsmtpAlertHandlerSchemaUrn {
	return &NullableEnumsmtpAlertHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsmtpAlertHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsmtpAlertHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
