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

// EnumtrustedCertificateSchemaUrn the model 'EnumtrustedCertificateSchemaUrn'
type EnumtrustedCertificateSchemaUrn string

// List of Enumtrusted-certificateSchemaUrn
const (
	ENUMTRUSTEDCERTIFICATESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0TRUSTED_CERTIFICATE EnumtrustedCertificateSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:trusted-certificate"
)

// All allowed values of EnumtrustedCertificateSchemaUrn enum
var AllowedEnumtrustedCertificateSchemaUrnEnumValues = []EnumtrustedCertificateSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:trusted-certificate",
}

func (v *EnumtrustedCertificateSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumtrustedCertificateSchemaUrn(value)
	for _, existing := range AllowedEnumtrustedCertificateSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumtrustedCertificateSchemaUrn", value)
}

// NewEnumtrustedCertificateSchemaUrnFromValue returns a pointer to a valid EnumtrustedCertificateSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumtrustedCertificateSchemaUrnFromValue(v string) (*EnumtrustedCertificateSchemaUrn, error) {
	ev := EnumtrustedCertificateSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumtrustedCertificateSchemaUrn: valid values are %v", v, AllowedEnumtrustedCertificateSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumtrustedCertificateSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumtrustedCertificateSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumtrusted-certificateSchemaUrn value
func (v EnumtrustedCertificateSchemaUrn) Ptr() *EnumtrustedCertificateSchemaUrn {
	return &v
}

type NullableEnumtrustedCertificateSchemaUrn struct {
	value *EnumtrustedCertificateSchemaUrn
	isSet bool
}

func (v NullableEnumtrustedCertificateSchemaUrn) Get() *EnumtrustedCertificateSchemaUrn {
	return v.value
}

func (v *NullableEnumtrustedCertificateSchemaUrn) Set(val *EnumtrustedCertificateSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumtrustedCertificateSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumtrustedCertificateSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumtrustedCertificateSchemaUrn(val *EnumtrustedCertificateSchemaUrn) *NullableEnumtrustedCertificateSchemaUrn {
	return &NullableEnumtrustedCertificateSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumtrustedCertificateSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumtrustedCertificateSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

