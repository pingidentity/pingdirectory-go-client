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

// EnumthirdPartyCertificateMapperSchemaUrn the model 'EnumthirdPartyCertificateMapperSchemaUrn'
type EnumthirdPartyCertificateMapperSchemaUrn string

// List of Enumthird-party-certificate-mapperSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0CERTIFICATE_MAPPERTHIRD_PARTY EnumthirdPartyCertificateMapperSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:certificate-mapper:third-party"
)

// All allowed values of EnumthirdPartyCertificateMapperSchemaUrn enum
var AllowedEnumthirdPartyCertificateMapperSchemaUrnEnumValues = []EnumthirdPartyCertificateMapperSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:certificate-mapper:third-party",
}

func (v *EnumthirdPartyCertificateMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumthirdPartyCertificateMapperSchemaUrn(value)
	for _, existing := range AllowedEnumthirdPartyCertificateMapperSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumthirdPartyCertificateMapperSchemaUrn", value)
}

// NewEnumthirdPartyCertificateMapperSchemaUrnFromValue returns a pointer to a valid EnumthirdPartyCertificateMapperSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumthirdPartyCertificateMapperSchemaUrnFromValue(v string) (*EnumthirdPartyCertificateMapperSchemaUrn, error) {
	ev := EnumthirdPartyCertificateMapperSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumthirdPartyCertificateMapperSchemaUrn: valid values are %v", v, AllowedEnumthirdPartyCertificateMapperSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumthirdPartyCertificateMapperSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumthirdPartyCertificateMapperSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumthird-party-certificate-mapperSchemaUrn value
func (v EnumthirdPartyCertificateMapperSchemaUrn) Ptr() *EnumthirdPartyCertificateMapperSchemaUrn {
	return &v
}

type NullableEnumthirdPartyCertificateMapperSchemaUrn struct {
	value *EnumthirdPartyCertificateMapperSchemaUrn
	isSet bool
}

func (v NullableEnumthirdPartyCertificateMapperSchemaUrn) Get() *EnumthirdPartyCertificateMapperSchemaUrn {
	return v.value
}

func (v *NullableEnumthirdPartyCertificateMapperSchemaUrn) Set(val *EnumthirdPartyCertificateMapperSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumthirdPartyCertificateMapperSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumthirdPartyCertificateMapperSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumthirdPartyCertificateMapperSchemaUrn(val *EnumthirdPartyCertificateMapperSchemaUrn) *NullableEnumthirdPartyCertificateMapperSchemaUrn {
	return &NullableEnumthirdPartyCertificateMapperSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumthirdPartyCertificateMapperSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumthirdPartyCertificateMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

