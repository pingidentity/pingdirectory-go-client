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

// EnummultiplePasswordDataSecurityAuditorSchemaUrn the model 'EnummultiplePasswordDataSecurityAuditorSchemaUrn'
type EnummultiplePasswordDataSecurityAuditorSchemaUrn string

// List of Enummultiple-password-data-security-auditorSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0DATA_SECURITY_AUDITORMULTIPLE_PASSWORD EnummultiplePasswordDataSecurityAuditorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:data-security-auditor:multiple-password"
)

// All allowed values of EnummultiplePasswordDataSecurityAuditorSchemaUrn enum
var AllowedEnummultiplePasswordDataSecurityAuditorSchemaUrnEnumValues = []EnummultiplePasswordDataSecurityAuditorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:data-security-auditor:multiple-password",
}

func (v *EnummultiplePasswordDataSecurityAuditorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnummultiplePasswordDataSecurityAuditorSchemaUrn(value)
	for _, existing := range AllowedEnummultiplePasswordDataSecurityAuditorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnummultiplePasswordDataSecurityAuditorSchemaUrn", value)
}

// NewEnummultiplePasswordDataSecurityAuditorSchemaUrnFromValue returns a pointer to a valid EnummultiplePasswordDataSecurityAuditorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnummultiplePasswordDataSecurityAuditorSchemaUrnFromValue(v string) (*EnummultiplePasswordDataSecurityAuditorSchemaUrn, error) {
	ev := EnummultiplePasswordDataSecurityAuditorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnummultiplePasswordDataSecurityAuditorSchemaUrn: valid values are %v", v, AllowedEnummultiplePasswordDataSecurityAuditorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnummultiplePasswordDataSecurityAuditorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnummultiplePasswordDataSecurityAuditorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enummultiple-password-data-security-auditorSchemaUrn value
func (v EnummultiplePasswordDataSecurityAuditorSchemaUrn) Ptr() *EnummultiplePasswordDataSecurityAuditorSchemaUrn {
	return &v
}

type NullableEnummultiplePasswordDataSecurityAuditorSchemaUrn struct {
	value *EnummultiplePasswordDataSecurityAuditorSchemaUrn
	isSet bool
}

func (v NullableEnummultiplePasswordDataSecurityAuditorSchemaUrn) Get() *EnummultiplePasswordDataSecurityAuditorSchemaUrn {
	return v.value
}

func (v *NullableEnummultiplePasswordDataSecurityAuditorSchemaUrn) Set(val *EnummultiplePasswordDataSecurityAuditorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnummultiplePasswordDataSecurityAuditorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnummultiplePasswordDataSecurityAuditorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnummultiplePasswordDataSecurityAuditorSchemaUrn(val *EnummultiplePasswordDataSecurityAuditorSchemaUrn) *NullableEnummultiplePasswordDataSecurityAuditorSchemaUrn {
	return &NullableEnummultiplePasswordDataSecurityAuditorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnummultiplePasswordDataSecurityAuditorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnummultiplePasswordDataSecurityAuditorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

