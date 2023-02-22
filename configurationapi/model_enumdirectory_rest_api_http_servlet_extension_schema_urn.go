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

// EnumdirectoryRestApiHttpServletExtensionSchemaUrn the model 'EnumdirectoryRestApiHttpServletExtensionSchemaUrn'
type EnumdirectoryRestApiHttpServletExtensionSchemaUrn string

// List of Enumdirectory-rest-api-http-servlet-extensionSchemaUrn
const (
	ENUMDIRECTORYRESTAPIHTTPSERVLETEXTENSIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0HTTP_SERVLET_EXTENSIONDIRECTORY_REST_API EnumdirectoryRestApiHttpServletExtensionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:directory-rest-api"
)

// All allowed values of EnumdirectoryRestApiHttpServletExtensionSchemaUrn enum
var AllowedEnumdirectoryRestApiHttpServletExtensionSchemaUrnEnumValues = []EnumdirectoryRestApiHttpServletExtensionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:http-servlet-extension:directory-rest-api",
}

func (v *EnumdirectoryRestApiHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdirectoryRestApiHttpServletExtensionSchemaUrn(value)
	for _, existing := range AllowedEnumdirectoryRestApiHttpServletExtensionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdirectoryRestApiHttpServletExtensionSchemaUrn", value)
}

// NewEnumdirectoryRestApiHttpServletExtensionSchemaUrnFromValue returns a pointer to a valid EnumdirectoryRestApiHttpServletExtensionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdirectoryRestApiHttpServletExtensionSchemaUrnFromValue(v string) (*EnumdirectoryRestApiHttpServletExtensionSchemaUrn, error) {
	ev := EnumdirectoryRestApiHttpServletExtensionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdirectoryRestApiHttpServletExtensionSchemaUrn: valid values are %v", v, AllowedEnumdirectoryRestApiHttpServletExtensionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdirectoryRestApiHttpServletExtensionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdirectoryRestApiHttpServletExtensionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdirectory-rest-api-http-servlet-extensionSchemaUrn value
func (v EnumdirectoryRestApiHttpServletExtensionSchemaUrn) Ptr() *EnumdirectoryRestApiHttpServletExtensionSchemaUrn {
	return &v
}

type NullableEnumdirectoryRestApiHttpServletExtensionSchemaUrn struct {
	value *EnumdirectoryRestApiHttpServletExtensionSchemaUrn
	isSet bool
}

func (v NullableEnumdirectoryRestApiHttpServletExtensionSchemaUrn) Get() *EnumdirectoryRestApiHttpServletExtensionSchemaUrn {
	return v.value
}

func (v *NullableEnumdirectoryRestApiHttpServletExtensionSchemaUrn) Set(val *EnumdirectoryRestApiHttpServletExtensionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdirectoryRestApiHttpServletExtensionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdirectoryRestApiHttpServletExtensionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdirectoryRestApiHttpServletExtensionSchemaUrn(val *EnumdirectoryRestApiHttpServletExtensionSchemaUrn) *NullableEnumdirectoryRestApiHttpServletExtensionSchemaUrn {
	return &NullableEnumdirectoryRestApiHttpServletExtensionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdirectoryRestApiHttpServletExtensionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdirectoryRestApiHttpServletExtensionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}