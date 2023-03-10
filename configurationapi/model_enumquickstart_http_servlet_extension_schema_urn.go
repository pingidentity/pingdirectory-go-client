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

// EnumquickstartHttpServletExtensionSchemaUrn the model 'EnumquickstartHttpServletExtensionSchemaUrn'
type EnumquickstartHttpServletExtensionSchemaUrn string

// List of Enumquickstart-http-servlet-extensionSchemaUrn
const (
	ENUMQUICKSTARTHTTPSERVLETEXTENSIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0HTTP_SERVLET_EXTENSIONQUICKSTART EnumquickstartHttpServletExtensionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:quickstart"
)

// All allowed values of EnumquickstartHttpServletExtensionSchemaUrn enum
var AllowedEnumquickstartHttpServletExtensionSchemaUrnEnumValues = []EnumquickstartHttpServletExtensionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:quickstart",
}

func (v *EnumquickstartHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumquickstartHttpServletExtensionSchemaUrn(value)
	for _, existing := range AllowedEnumquickstartHttpServletExtensionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumquickstartHttpServletExtensionSchemaUrn", value)
}

// NewEnumquickstartHttpServletExtensionSchemaUrnFromValue returns a pointer to a valid EnumquickstartHttpServletExtensionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumquickstartHttpServletExtensionSchemaUrnFromValue(v string) (*EnumquickstartHttpServletExtensionSchemaUrn, error) {
	ev := EnumquickstartHttpServletExtensionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumquickstartHttpServletExtensionSchemaUrn: valid values are %v", v, AllowedEnumquickstartHttpServletExtensionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumquickstartHttpServletExtensionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumquickstartHttpServletExtensionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumquickstart-http-servlet-extensionSchemaUrn value
func (v EnumquickstartHttpServletExtensionSchemaUrn) Ptr() *EnumquickstartHttpServletExtensionSchemaUrn {
	return &v
}

type NullableEnumquickstartHttpServletExtensionSchemaUrn struct {
	value *EnumquickstartHttpServletExtensionSchemaUrn
	isSet bool
}

func (v NullableEnumquickstartHttpServletExtensionSchemaUrn) Get() *EnumquickstartHttpServletExtensionSchemaUrn {
	return v.value
}

func (v *NullableEnumquickstartHttpServletExtensionSchemaUrn) Set(val *EnumquickstartHttpServletExtensionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumquickstartHttpServletExtensionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumquickstartHttpServletExtensionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumquickstartHttpServletExtensionSchemaUrn(val *EnumquickstartHttpServletExtensionSchemaUrn) *NullableEnumquickstartHttpServletExtensionSchemaUrn {
	return &NullableEnumquickstartHttpServletExtensionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumquickstartHttpServletExtensionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumquickstartHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
