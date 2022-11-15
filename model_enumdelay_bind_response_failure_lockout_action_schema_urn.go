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

// EnumdelayBindResponseFailureLockoutActionSchemaUrn the model 'EnumdelayBindResponseFailureLockoutActionSchemaUrn'
type EnumdelayBindResponseFailureLockoutActionSchemaUrn string

// List of Enumdelay-bind-response-failure-lockout-actionSchemaUrn
const (
	ENUMDELAYBINDRESPONSEFAILURELOCKOUTACTIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0FAILURE_LOCKOUT_ACTIONDELAY_BIND_RESPONSE EnumdelayBindResponseFailureLockoutActionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:failure-lockout-action:delay-bind-response"
)

// All allowed values of EnumdelayBindResponseFailureLockoutActionSchemaUrn enum
var AllowedEnumdelayBindResponseFailureLockoutActionSchemaUrnEnumValues = []EnumdelayBindResponseFailureLockoutActionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:failure-lockout-action:delay-bind-response",
}

func (v *EnumdelayBindResponseFailureLockoutActionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdelayBindResponseFailureLockoutActionSchemaUrn(value)
	for _, existing := range AllowedEnumdelayBindResponseFailureLockoutActionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdelayBindResponseFailureLockoutActionSchemaUrn", value)
}

// NewEnumdelayBindResponseFailureLockoutActionSchemaUrnFromValue returns a pointer to a valid EnumdelayBindResponseFailureLockoutActionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdelayBindResponseFailureLockoutActionSchemaUrnFromValue(v string) (*EnumdelayBindResponseFailureLockoutActionSchemaUrn, error) {
	ev := EnumdelayBindResponseFailureLockoutActionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdelayBindResponseFailureLockoutActionSchemaUrn: valid values are %v", v, AllowedEnumdelayBindResponseFailureLockoutActionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdelayBindResponseFailureLockoutActionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdelayBindResponseFailureLockoutActionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdelay-bind-response-failure-lockout-actionSchemaUrn value
func (v EnumdelayBindResponseFailureLockoutActionSchemaUrn) Ptr() *EnumdelayBindResponseFailureLockoutActionSchemaUrn {
	return &v
}

type NullableEnumdelayBindResponseFailureLockoutActionSchemaUrn struct {
	value *EnumdelayBindResponseFailureLockoutActionSchemaUrn
	isSet bool
}

func (v NullableEnumdelayBindResponseFailureLockoutActionSchemaUrn) Get() *EnumdelayBindResponseFailureLockoutActionSchemaUrn {
	return v.value
}

func (v *NullableEnumdelayBindResponseFailureLockoutActionSchemaUrn) Set(val *EnumdelayBindResponseFailureLockoutActionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdelayBindResponseFailureLockoutActionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdelayBindResponseFailureLockoutActionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdelayBindResponseFailureLockoutActionSchemaUrn(val *EnumdelayBindResponseFailureLockoutActionSchemaUrn) *NullableEnumdelayBindResponseFailureLockoutActionSchemaUrn {
	return &NullableEnumdelayBindResponseFailureLockoutActionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdelayBindResponseFailureLockoutActionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdelayBindResponseFailureLockoutActionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

