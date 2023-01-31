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

// EnumsubjectEqualsDnCertificateMapperSchemaUrn the model 'EnumsubjectEqualsDnCertificateMapperSchemaUrn'
type EnumsubjectEqualsDnCertificateMapperSchemaUrn string

// List of Enumsubject-equals-dn-certificate-mapperSchemaUrn
const (
	ENUMSUBJECTEQUALSDNCERTIFICATEMAPPERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CERTIFICATE_MAPPERSUBJECT_EQUALS_DN EnumsubjectEqualsDnCertificateMapperSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:certificate-mapper:subject-equals-dn"
)

// All allowed values of EnumsubjectEqualsDnCertificateMapperSchemaUrn enum
var AllowedEnumsubjectEqualsDnCertificateMapperSchemaUrnEnumValues = []EnumsubjectEqualsDnCertificateMapperSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:certificate-mapper:subject-equals-dn",
}

func (v *EnumsubjectEqualsDnCertificateMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsubjectEqualsDnCertificateMapperSchemaUrn(value)
	for _, existing := range AllowedEnumsubjectEqualsDnCertificateMapperSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsubjectEqualsDnCertificateMapperSchemaUrn", value)
}

// NewEnumsubjectEqualsDnCertificateMapperSchemaUrnFromValue returns a pointer to a valid EnumsubjectEqualsDnCertificateMapperSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsubjectEqualsDnCertificateMapperSchemaUrnFromValue(v string) (*EnumsubjectEqualsDnCertificateMapperSchemaUrn, error) {
	ev := EnumsubjectEqualsDnCertificateMapperSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsubjectEqualsDnCertificateMapperSchemaUrn: valid values are %v", v, AllowedEnumsubjectEqualsDnCertificateMapperSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsubjectEqualsDnCertificateMapperSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsubjectEqualsDnCertificateMapperSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsubject-equals-dn-certificate-mapperSchemaUrn value
func (v EnumsubjectEqualsDnCertificateMapperSchemaUrn) Ptr() *EnumsubjectEqualsDnCertificateMapperSchemaUrn {
	return &v
}

type NullableEnumsubjectEqualsDnCertificateMapperSchemaUrn struct {
	value *EnumsubjectEqualsDnCertificateMapperSchemaUrn
	isSet bool
}

func (v NullableEnumsubjectEqualsDnCertificateMapperSchemaUrn) Get() *EnumsubjectEqualsDnCertificateMapperSchemaUrn {
	return v.value
}

func (v *NullableEnumsubjectEqualsDnCertificateMapperSchemaUrn) Set(val *EnumsubjectEqualsDnCertificateMapperSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsubjectEqualsDnCertificateMapperSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsubjectEqualsDnCertificateMapperSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsubjectEqualsDnCertificateMapperSchemaUrn(val *EnumsubjectEqualsDnCertificateMapperSchemaUrn) *NullableEnumsubjectEqualsDnCertificateMapperSchemaUrn {
	return &NullableEnumsubjectEqualsDnCertificateMapperSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsubjectEqualsDnCertificateMapperSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsubjectEqualsDnCertificateMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
