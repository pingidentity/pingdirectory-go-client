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

// EnumthirdPartyHttpServletExtensionSchemaUrn the model 'EnumthirdPartyHttpServletExtensionSchemaUrn'
type EnumthirdPartyHttpServletExtensionSchemaUrn string

// List of Enumthird-party-http-servlet-extensionSchemaUrn
const (
	ENUMTHIRDPARTYHTTPSERVLETEXTENSIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0HTTP_SERVLET_EXTENSIONTHIRD_PARTY EnumthirdPartyHttpServletExtensionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:third-party"
)

// All allowed values of EnumthirdPartyHttpServletExtensionSchemaUrn enum
var AllowedEnumthirdPartyHttpServletExtensionSchemaUrnEnumValues = []EnumthirdPartyHttpServletExtensionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:third-party",
}

func (v *EnumthirdPartyHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumthirdPartyHttpServletExtensionSchemaUrn(value)
	for _, existing := range AllowedEnumthirdPartyHttpServletExtensionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumthirdPartyHttpServletExtensionSchemaUrn", value)
}

// NewEnumthirdPartyHttpServletExtensionSchemaUrnFromValue returns a pointer to a valid EnumthirdPartyHttpServletExtensionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumthirdPartyHttpServletExtensionSchemaUrnFromValue(v string) (*EnumthirdPartyHttpServletExtensionSchemaUrn, error) {
	ev := EnumthirdPartyHttpServletExtensionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumthirdPartyHttpServletExtensionSchemaUrn: valid values are %v", v, AllowedEnumthirdPartyHttpServletExtensionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumthirdPartyHttpServletExtensionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumthirdPartyHttpServletExtensionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumthird-party-http-servlet-extensionSchemaUrn value
func (v EnumthirdPartyHttpServletExtensionSchemaUrn) Ptr() *EnumthirdPartyHttpServletExtensionSchemaUrn {
	return &v
}

type NullableEnumthirdPartyHttpServletExtensionSchemaUrn struct {
	value *EnumthirdPartyHttpServletExtensionSchemaUrn
	isSet bool
}

func (v NullableEnumthirdPartyHttpServletExtensionSchemaUrn) Get() *EnumthirdPartyHttpServletExtensionSchemaUrn {
	return v.value
}

func (v *NullableEnumthirdPartyHttpServletExtensionSchemaUrn) Set(val *EnumthirdPartyHttpServletExtensionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumthirdPartyHttpServletExtensionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumthirdPartyHttpServletExtensionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumthirdPartyHttpServletExtensionSchemaUrn(val *EnumthirdPartyHttpServletExtensionSchemaUrn) *NullableEnumthirdPartyHttpServletExtensionSchemaUrn {
	return &NullableEnumthirdPartyHttpServletExtensionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumthirdPartyHttpServletExtensionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumthirdPartyHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}