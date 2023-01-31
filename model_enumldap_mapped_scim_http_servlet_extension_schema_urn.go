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

// EnumldapMappedScimHttpServletExtensionSchemaUrn the model 'EnumldapMappedScimHttpServletExtensionSchemaUrn'
type EnumldapMappedScimHttpServletExtensionSchemaUrn string

// List of Enumldap-mapped-scim-http-servlet-extensionSchemaUrn
const (
	ENUMLDAPMAPPEDSCIMHTTPSERVLETEXTENSIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0HTTP_SERVLET_EXTENSIONLDAP_MAPPED_SCIM EnumldapMappedScimHttpServletExtensionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:ldap-mapped-scim"
)

// All allowed values of EnumldapMappedScimHttpServletExtensionSchemaUrn enum
var AllowedEnumldapMappedScimHttpServletExtensionSchemaUrnEnumValues = []EnumldapMappedScimHttpServletExtensionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:ldap-mapped-scim",
}

func (v *EnumldapMappedScimHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumldapMappedScimHttpServletExtensionSchemaUrn(value)
	for _, existing := range AllowedEnumldapMappedScimHttpServletExtensionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumldapMappedScimHttpServletExtensionSchemaUrn", value)
}

// NewEnumldapMappedScimHttpServletExtensionSchemaUrnFromValue returns a pointer to a valid EnumldapMappedScimHttpServletExtensionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumldapMappedScimHttpServletExtensionSchemaUrnFromValue(v string) (*EnumldapMappedScimHttpServletExtensionSchemaUrn, error) {
	ev := EnumldapMappedScimHttpServletExtensionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumldapMappedScimHttpServletExtensionSchemaUrn: valid values are %v", v, AllowedEnumldapMappedScimHttpServletExtensionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumldapMappedScimHttpServletExtensionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumldapMappedScimHttpServletExtensionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumldap-mapped-scim-http-servlet-extensionSchemaUrn value
func (v EnumldapMappedScimHttpServletExtensionSchemaUrn) Ptr() *EnumldapMappedScimHttpServletExtensionSchemaUrn {
	return &v
}

type NullableEnumldapMappedScimHttpServletExtensionSchemaUrn struct {
	value *EnumldapMappedScimHttpServletExtensionSchemaUrn
	isSet bool
}

func (v NullableEnumldapMappedScimHttpServletExtensionSchemaUrn) Get() *EnumldapMappedScimHttpServletExtensionSchemaUrn {
	return v.value
}

func (v *NullableEnumldapMappedScimHttpServletExtensionSchemaUrn) Set(val *EnumldapMappedScimHttpServletExtensionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumldapMappedScimHttpServletExtensionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumldapMappedScimHttpServletExtensionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumldapMappedScimHttpServletExtensionSchemaUrn(val *EnumldapMappedScimHttpServletExtensionSchemaUrn) *NullableEnumldapMappedScimHttpServletExtensionSchemaUrn {
	return &NullableEnumldapMappedScimHttpServletExtensionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumldapMappedScimHttpServletExtensionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumldapMappedScimHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
