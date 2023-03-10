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

// EnumcertificateInterServerAuthenticationInfoSchemaUrn the model 'EnumcertificateInterServerAuthenticationInfoSchemaUrn'
type EnumcertificateInterServerAuthenticationInfoSchemaUrn string

// List of Enumcertificate-inter-server-authentication-infoSchemaUrn
const (
	ENUMCERTIFICATEINTERSERVERAUTHENTICATIONINFOSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0INTER_SERVER_AUTHENTICATION_INFOCERTIFICATE EnumcertificateInterServerAuthenticationInfoSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:inter-server-authentication-info:certificate"
)

// All allowed values of EnumcertificateInterServerAuthenticationInfoSchemaUrn enum
var AllowedEnumcertificateInterServerAuthenticationInfoSchemaUrnEnumValues = []EnumcertificateInterServerAuthenticationInfoSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:inter-server-authentication-info:certificate",
}

func (v *EnumcertificateInterServerAuthenticationInfoSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcertificateInterServerAuthenticationInfoSchemaUrn(value)
	for _, existing := range AllowedEnumcertificateInterServerAuthenticationInfoSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcertificateInterServerAuthenticationInfoSchemaUrn", value)
}

// NewEnumcertificateInterServerAuthenticationInfoSchemaUrnFromValue returns a pointer to a valid EnumcertificateInterServerAuthenticationInfoSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcertificateInterServerAuthenticationInfoSchemaUrnFromValue(v string) (*EnumcertificateInterServerAuthenticationInfoSchemaUrn, error) {
	ev := EnumcertificateInterServerAuthenticationInfoSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcertificateInterServerAuthenticationInfoSchemaUrn: valid values are %v", v, AllowedEnumcertificateInterServerAuthenticationInfoSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcertificateInterServerAuthenticationInfoSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcertificateInterServerAuthenticationInfoSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumcertificate-inter-server-authentication-infoSchemaUrn value
func (v EnumcertificateInterServerAuthenticationInfoSchemaUrn) Ptr() *EnumcertificateInterServerAuthenticationInfoSchemaUrn {
	return &v
}

type NullableEnumcertificateInterServerAuthenticationInfoSchemaUrn struct {
	value *EnumcertificateInterServerAuthenticationInfoSchemaUrn
	isSet bool
}

func (v NullableEnumcertificateInterServerAuthenticationInfoSchemaUrn) Get() *EnumcertificateInterServerAuthenticationInfoSchemaUrn {
	return v.value
}

func (v *NullableEnumcertificateInterServerAuthenticationInfoSchemaUrn) Set(val *EnumcertificateInterServerAuthenticationInfoSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcertificateInterServerAuthenticationInfoSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcertificateInterServerAuthenticationInfoSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcertificateInterServerAuthenticationInfoSchemaUrn(val *EnumcertificateInterServerAuthenticationInfoSchemaUrn) *NullableEnumcertificateInterServerAuthenticationInfoSchemaUrn {
	return &NullableEnumcertificateInterServerAuthenticationInfoSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcertificateInterServerAuthenticationInfoSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcertificateInterServerAuthenticationInfoSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
