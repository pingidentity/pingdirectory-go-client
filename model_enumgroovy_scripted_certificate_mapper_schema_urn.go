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

// EnumgroovyScriptedCertificateMapperSchemaUrn the model 'EnumgroovyScriptedCertificateMapperSchemaUrn'
type EnumgroovyScriptedCertificateMapperSchemaUrn string

// List of Enumgroovy-scripted-certificate-mapperSchemaUrn
const (
	ENUMGROOVYSCRIPTEDCERTIFICATEMAPPERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CERTIFICATE_MAPPERGROOVY_SCRIPTED EnumgroovyScriptedCertificateMapperSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:certificate-mapper:groovy-scripted"
)

// All allowed values of EnumgroovyScriptedCertificateMapperSchemaUrn enum
var AllowedEnumgroovyScriptedCertificateMapperSchemaUrnEnumValues = []EnumgroovyScriptedCertificateMapperSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:certificate-mapper:groovy-scripted",
}

func (v *EnumgroovyScriptedCertificateMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgroovyScriptedCertificateMapperSchemaUrn(value)
	for _, existing := range AllowedEnumgroovyScriptedCertificateMapperSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgroovyScriptedCertificateMapperSchemaUrn", value)
}

// NewEnumgroovyScriptedCertificateMapperSchemaUrnFromValue returns a pointer to a valid EnumgroovyScriptedCertificateMapperSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgroovyScriptedCertificateMapperSchemaUrnFromValue(v string) (*EnumgroovyScriptedCertificateMapperSchemaUrn, error) {
	ev := EnumgroovyScriptedCertificateMapperSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgroovyScriptedCertificateMapperSchemaUrn: valid values are %v", v, AllowedEnumgroovyScriptedCertificateMapperSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgroovyScriptedCertificateMapperSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgroovyScriptedCertificateMapperSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgroovy-scripted-certificate-mapperSchemaUrn value
func (v EnumgroovyScriptedCertificateMapperSchemaUrn) Ptr() *EnumgroovyScriptedCertificateMapperSchemaUrn {
	return &v
}

type NullableEnumgroovyScriptedCertificateMapperSchemaUrn struct {
	value *EnumgroovyScriptedCertificateMapperSchemaUrn
	isSet bool
}

func (v NullableEnumgroovyScriptedCertificateMapperSchemaUrn) Get() *EnumgroovyScriptedCertificateMapperSchemaUrn {
	return v.value
}

func (v *NullableEnumgroovyScriptedCertificateMapperSchemaUrn) Set(val *EnumgroovyScriptedCertificateMapperSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgroovyScriptedCertificateMapperSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgroovyScriptedCertificateMapperSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgroovyScriptedCertificateMapperSchemaUrn(val *EnumgroovyScriptedCertificateMapperSchemaUrn) *NullableEnumgroovyScriptedCertificateMapperSchemaUrn {
	return &NullableEnumgroovyScriptedCertificateMapperSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgroovyScriptedCertificateMapperSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgroovyScriptedCertificateMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
