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

// EnumgenericWebApplicationExtensionSchemaUrn the model 'EnumgenericWebApplicationExtensionSchemaUrn'
type EnumgenericWebApplicationExtensionSchemaUrn string

// List of Enumgeneric-web-application-extensionSchemaUrn
const (
	ENUMGENERICWEBAPPLICATIONEXTENSIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0WEB_APPLICATION_EXTENSIONGENERIC EnumgenericWebApplicationExtensionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:web-application-extension:generic"
)

// All allowed values of EnumgenericWebApplicationExtensionSchemaUrn enum
var AllowedEnumgenericWebApplicationExtensionSchemaUrnEnumValues = []EnumgenericWebApplicationExtensionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:web-application-extension:generic",
}

func (v *EnumgenericWebApplicationExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgenericWebApplicationExtensionSchemaUrn(value)
	for _, existing := range AllowedEnumgenericWebApplicationExtensionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgenericWebApplicationExtensionSchemaUrn", value)
}

// NewEnumgenericWebApplicationExtensionSchemaUrnFromValue returns a pointer to a valid EnumgenericWebApplicationExtensionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgenericWebApplicationExtensionSchemaUrnFromValue(v string) (*EnumgenericWebApplicationExtensionSchemaUrn, error) {
	ev := EnumgenericWebApplicationExtensionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgenericWebApplicationExtensionSchemaUrn: valid values are %v", v, AllowedEnumgenericWebApplicationExtensionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgenericWebApplicationExtensionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgenericWebApplicationExtensionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgeneric-web-application-extensionSchemaUrn value
func (v EnumgenericWebApplicationExtensionSchemaUrn) Ptr() *EnumgenericWebApplicationExtensionSchemaUrn {
	return &v
}

type NullableEnumgenericWebApplicationExtensionSchemaUrn struct {
	value *EnumgenericWebApplicationExtensionSchemaUrn
	isSet bool
}

func (v NullableEnumgenericWebApplicationExtensionSchemaUrn) Get() *EnumgenericWebApplicationExtensionSchemaUrn {
	return v.value
}

func (v *NullableEnumgenericWebApplicationExtensionSchemaUrn) Set(val *EnumgenericWebApplicationExtensionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgenericWebApplicationExtensionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgenericWebApplicationExtensionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgenericWebApplicationExtensionSchemaUrn(val *EnumgenericWebApplicationExtensionSchemaUrn) *NullableEnumgenericWebApplicationExtensionSchemaUrn {
	return &NullableEnumgenericWebApplicationExtensionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgenericWebApplicationExtensionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgenericWebApplicationExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
