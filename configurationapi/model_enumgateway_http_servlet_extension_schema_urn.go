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

// EnumgatewayHttpServletExtensionSchemaUrn the model 'EnumgatewayHttpServletExtensionSchemaUrn'
type EnumgatewayHttpServletExtensionSchemaUrn string

// List of Enumgateway-http-servlet-extensionSchemaUrn
const (
	ENUMGATEWAYHTTPSERVLETEXTENSIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0HTTP_SERVLET_EXTENSIONGATEWAY EnumgatewayHttpServletExtensionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:gateway"
)

// All allowed values of EnumgatewayHttpServletExtensionSchemaUrn enum
var AllowedEnumgatewayHttpServletExtensionSchemaUrnEnumValues = []EnumgatewayHttpServletExtensionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:gateway",
}

func (v *EnumgatewayHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgatewayHttpServletExtensionSchemaUrn(value)
	for _, existing := range AllowedEnumgatewayHttpServletExtensionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgatewayHttpServletExtensionSchemaUrn", value)
}

// NewEnumgatewayHttpServletExtensionSchemaUrnFromValue returns a pointer to a valid EnumgatewayHttpServletExtensionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgatewayHttpServletExtensionSchemaUrnFromValue(v string) (*EnumgatewayHttpServletExtensionSchemaUrn, error) {
	ev := EnumgatewayHttpServletExtensionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgatewayHttpServletExtensionSchemaUrn: valid values are %v", v, AllowedEnumgatewayHttpServletExtensionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgatewayHttpServletExtensionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgatewayHttpServletExtensionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgateway-http-servlet-extensionSchemaUrn value
func (v EnumgatewayHttpServletExtensionSchemaUrn) Ptr() *EnumgatewayHttpServletExtensionSchemaUrn {
	return &v
}

type NullableEnumgatewayHttpServletExtensionSchemaUrn struct {
	value *EnumgatewayHttpServletExtensionSchemaUrn
	isSet bool
}

func (v NullableEnumgatewayHttpServletExtensionSchemaUrn) Get() *EnumgatewayHttpServletExtensionSchemaUrn {
	return v.value
}

func (v *NullableEnumgatewayHttpServletExtensionSchemaUrn) Set(val *EnumgatewayHttpServletExtensionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgatewayHttpServletExtensionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgatewayHttpServletExtensionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgatewayHttpServletExtensionSchemaUrn(val *EnumgatewayHttpServletExtensionSchemaUrn) *NullableEnumgatewayHttpServletExtensionSchemaUrn {
	return &NullableEnumgatewayHttpServletExtensionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgatewayHttpServletExtensionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgatewayHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
