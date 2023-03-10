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

// EnumvelocityHttpServletExtensionSchemaUrn the model 'EnumvelocityHttpServletExtensionSchemaUrn'
type EnumvelocityHttpServletExtensionSchemaUrn string

// List of Enumvelocity-http-servlet-extensionSchemaUrn
const (
	ENUMVELOCITYHTTPSERVLETEXTENSIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0HTTP_SERVLET_EXTENSIONVELOCITY EnumvelocityHttpServletExtensionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:velocity"
)

// All allowed values of EnumvelocityHttpServletExtensionSchemaUrn enum
var AllowedEnumvelocityHttpServletExtensionSchemaUrnEnumValues = []EnumvelocityHttpServletExtensionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:velocity",
}

func (v *EnumvelocityHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumvelocityHttpServletExtensionSchemaUrn(value)
	for _, existing := range AllowedEnumvelocityHttpServletExtensionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumvelocityHttpServletExtensionSchemaUrn", value)
}

// NewEnumvelocityHttpServletExtensionSchemaUrnFromValue returns a pointer to a valid EnumvelocityHttpServletExtensionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumvelocityHttpServletExtensionSchemaUrnFromValue(v string) (*EnumvelocityHttpServletExtensionSchemaUrn, error) {
	ev := EnumvelocityHttpServletExtensionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumvelocityHttpServletExtensionSchemaUrn: valid values are %v", v, AllowedEnumvelocityHttpServletExtensionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumvelocityHttpServletExtensionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumvelocityHttpServletExtensionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumvelocity-http-servlet-extensionSchemaUrn value
func (v EnumvelocityHttpServletExtensionSchemaUrn) Ptr() *EnumvelocityHttpServletExtensionSchemaUrn {
	return &v
}

type NullableEnumvelocityHttpServletExtensionSchemaUrn struct {
	value *EnumvelocityHttpServletExtensionSchemaUrn
	isSet bool
}

func (v NullableEnumvelocityHttpServletExtensionSchemaUrn) Get() *EnumvelocityHttpServletExtensionSchemaUrn {
	return v.value
}

func (v *NullableEnumvelocityHttpServletExtensionSchemaUrn) Set(val *EnumvelocityHttpServletExtensionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumvelocityHttpServletExtensionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumvelocityHttpServletExtensionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumvelocityHttpServletExtensionSchemaUrn(val *EnumvelocityHttpServletExtensionSchemaUrn) *NullableEnumvelocityHttpServletExtensionSchemaUrn {
	return &NullableEnumvelocityHttpServletExtensionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumvelocityHttpServletExtensionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumvelocityHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
