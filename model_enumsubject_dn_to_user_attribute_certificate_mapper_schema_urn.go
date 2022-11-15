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

// EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn the model 'EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn'
type EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn string

// List of Enumsubject-dn-to-user-attribute-certificate-mapperSchemaUrn
const (
	ENUMSUBJECTDNTOUSERATTRIBUTECERTIFICATEMAPPERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CERTIFICATE_MAPPERSUBJECT_DN_TO_USER_ATTRIBUTE EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:certificate-mapper:subject-dn-to-user-attribute"
)

// All allowed values of EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn enum
var AllowedEnumsubjectDnToUserAttributeCertificateMapperSchemaUrnEnumValues = []EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:certificate-mapper:subject-dn-to-user-attribute",
}

func (v *EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn(value)
	for _, existing := range AllowedEnumsubjectDnToUserAttributeCertificateMapperSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn", value)
}

// NewEnumsubjectDnToUserAttributeCertificateMapperSchemaUrnFromValue returns a pointer to a valid EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsubjectDnToUserAttributeCertificateMapperSchemaUrnFromValue(v string) (*EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn, error) {
	ev := EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn: valid values are %v", v, AllowedEnumsubjectDnToUserAttributeCertificateMapperSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsubjectDnToUserAttributeCertificateMapperSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsubject-dn-to-user-attribute-certificate-mapperSchemaUrn value
func (v EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) Ptr() *EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn {
	return &v
}

type NullableEnumsubjectDnToUserAttributeCertificateMapperSchemaUrn struct {
	value *EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn
	isSet bool
}

func (v NullableEnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) Get() *EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn {
	return v.value
}

func (v *NullableEnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) Set(val *EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsubjectDnToUserAttributeCertificateMapperSchemaUrn(val *EnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) *NullableEnumsubjectDnToUserAttributeCertificateMapperSchemaUrn {
	return &NullableEnumsubjectDnToUserAttributeCertificateMapperSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsubjectDnToUserAttributeCertificateMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

