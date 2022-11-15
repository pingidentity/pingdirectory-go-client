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

// EnumfingerprintCertificateMapperSchemaUrn the model 'EnumfingerprintCertificateMapperSchemaUrn'
type EnumfingerprintCertificateMapperSchemaUrn string

// List of Enumfingerprint-certificate-mapperSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0CERTIFICATE_MAPPERFINGERPRINT EnumfingerprintCertificateMapperSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:certificate-mapper:fingerprint"
)

// All allowed values of EnumfingerprintCertificateMapperSchemaUrn enum
var AllowedEnumfingerprintCertificateMapperSchemaUrnEnumValues = []EnumfingerprintCertificateMapperSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:certificate-mapper:fingerprint",
}

func (v *EnumfingerprintCertificateMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumfingerprintCertificateMapperSchemaUrn(value)
	for _, existing := range AllowedEnumfingerprintCertificateMapperSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumfingerprintCertificateMapperSchemaUrn", value)
}

// NewEnumfingerprintCertificateMapperSchemaUrnFromValue returns a pointer to a valid EnumfingerprintCertificateMapperSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumfingerprintCertificateMapperSchemaUrnFromValue(v string) (*EnumfingerprintCertificateMapperSchemaUrn, error) {
	ev := EnumfingerprintCertificateMapperSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumfingerprintCertificateMapperSchemaUrn: valid values are %v", v, AllowedEnumfingerprintCertificateMapperSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumfingerprintCertificateMapperSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumfingerprintCertificateMapperSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumfingerprint-certificate-mapperSchemaUrn value
func (v EnumfingerprintCertificateMapperSchemaUrn) Ptr() *EnumfingerprintCertificateMapperSchemaUrn {
	return &v
}

type NullableEnumfingerprintCertificateMapperSchemaUrn struct {
	value *EnumfingerprintCertificateMapperSchemaUrn
	isSet bool
}

func (v NullableEnumfingerprintCertificateMapperSchemaUrn) Get() *EnumfingerprintCertificateMapperSchemaUrn {
	return v.value
}

func (v *NullableEnumfingerprintCertificateMapperSchemaUrn) Set(val *EnumfingerprintCertificateMapperSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumfingerprintCertificateMapperSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumfingerprintCertificateMapperSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumfingerprintCertificateMapperSchemaUrn(val *EnumfingerprintCertificateMapperSchemaUrn) *NullableEnumfingerprintCertificateMapperSchemaUrn {
	return &NullableEnumfingerprintCertificateMapperSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumfingerprintCertificateMapperSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumfingerprintCertificateMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

