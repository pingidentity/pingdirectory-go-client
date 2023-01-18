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

// EnumdseeCompatAccessControlHandlerSchemaUrn the model 'EnumdseeCompatAccessControlHandlerSchemaUrn'
type EnumdseeCompatAccessControlHandlerSchemaUrn string

// List of Enumdsee-compat-access-control-handlerSchemaUrn
const (
	ENUMDSEECOMPATACCESSCONTROLHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ACCESS_CONTROL_HANDLERDSEE_COMPAT EnumdseeCompatAccessControlHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:access-control-handler:dsee-compat"
)

// All allowed values of EnumdseeCompatAccessControlHandlerSchemaUrn enum
var AllowedEnumdseeCompatAccessControlHandlerSchemaUrnEnumValues = []EnumdseeCompatAccessControlHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:access-control-handler:dsee-compat",
}

func (v *EnumdseeCompatAccessControlHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdseeCompatAccessControlHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumdseeCompatAccessControlHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdseeCompatAccessControlHandlerSchemaUrn", value)
}

// NewEnumdseeCompatAccessControlHandlerSchemaUrnFromValue returns a pointer to a valid EnumdseeCompatAccessControlHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdseeCompatAccessControlHandlerSchemaUrnFromValue(v string) (*EnumdseeCompatAccessControlHandlerSchemaUrn, error) {
	ev := EnumdseeCompatAccessControlHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdseeCompatAccessControlHandlerSchemaUrn: valid values are %v", v, AllowedEnumdseeCompatAccessControlHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdseeCompatAccessControlHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdseeCompatAccessControlHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdsee-compat-access-control-handlerSchemaUrn value
func (v EnumdseeCompatAccessControlHandlerSchemaUrn) Ptr() *EnumdseeCompatAccessControlHandlerSchemaUrn {
	return &v
}

type NullableEnumdseeCompatAccessControlHandlerSchemaUrn struct {
	value *EnumdseeCompatAccessControlHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumdseeCompatAccessControlHandlerSchemaUrn) Get() *EnumdseeCompatAccessControlHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumdseeCompatAccessControlHandlerSchemaUrn) Set(val *EnumdseeCompatAccessControlHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdseeCompatAccessControlHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdseeCompatAccessControlHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdseeCompatAccessControlHandlerSchemaUrn(val *EnumdseeCompatAccessControlHandlerSchemaUrn) *NullableEnumdseeCompatAccessControlHandlerSchemaUrn {
	return &NullableEnumdseeCompatAccessControlHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdseeCompatAccessControlHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdseeCompatAccessControlHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
