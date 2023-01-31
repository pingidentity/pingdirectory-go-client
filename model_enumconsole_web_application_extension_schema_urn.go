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

// EnumconsoleWebApplicationExtensionSchemaUrn the model 'EnumconsoleWebApplicationExtensionSchemaUrn'
type EnumconsoleWebApplicationExtensionSchemaUrn string

// List of Enumconsole-web-application-extensionSchemaUrn
const (
	ENUMCONSOLEWEBAPPLICATIONEXTENSIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0WEB_APPLICATION_EXTENSIONCONSOLE EnumconsoleWebApplicationExtensionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:web-application-extension:console"
)

// All allowed values of EnumconsoleWebApplicationExtensionSchemaUrn enum
var AllowedEnumconsoleWebApplicationExtensionSchemaUrnEnumValues = []EnumconsoleWebApplicationExtensionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:web-application-extension:console",
}

func (v *EnumconsoleWebApplicationExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumconsoleWebApplicationExtensionSchemaUrn(value)
	for _, existing := range AllowedEnumconsoleWebApplicationExtensionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumconsoleWebApplicationExtensionSchemaUrn", value)
}

// NewEnumconsoleWebApplicationExtensionSchemaUrnFromValue returns a pointer to a valid EnumconsoleWebApplicationExtensionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumconsoleWebApplicationExtensionSchemaUrnFromValue(v string) (*EnumconsoleWebApplicationExtensionSchemaUrn, error) {
	ev := EnumconsoleWebApplicationExtensionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumconsoleWebApplicationExtensionSchemaUrn: valid values are %v", v, AllowedEnumconsoleWebApplicationExtensionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumconsoleWebApplicationExtensionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumconsoleWebApplicationExtensionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumconsole-web-application-extensionSchemaUrn value
func (v EnumconsoleWebApplicationExtensionSchemaUrn) Ptr() *EnumconsoleWebApplicationExtensionSchemaUrn {
	return &v
}

type NullableEnumconsoleWebApplicationExtensionSchemaUrn struct {
	value *EnumconsoleWebApplicationExtensionSchemaUrn
	isSet bool
}

func (v NullableEnumconsoleWebApplicationExtensionSchemaUrn) Get() *EnumconsoleWebApplicationExtensionSchemaUrn {
	return v.value
}

func (v *NullableEnumconsoleWebApplicationExtensionSchemaUrn) Set(val *EnumconsoleWebApplicationExtensionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumconsoleWebApplicationExtensionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumconsoleWebApplicationExtensionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumconsoleWebApplicationExtensionSchemaUrn(val *EnumconsoleWebApplicationExtensionSchemaUrn) *NullableEnumconsoleWebApplicationExtensionSchemaUrn {
	return &NullableEnumconsoleWebApplicationExtensionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumconsoleWebApplicationExtensionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumconsoleWebApplicationExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
