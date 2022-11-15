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

// EnumdisabledAccountDataSecurityAuditorSchemaUrn the model 'EnumdisabledAccountDataSecurityAuditorSchemaUrn'
type EnumdisabledAccountDataSecurityAuditorSchemaUrn string

// List of Enumdisabled-account-data-security-auditorSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0DATA_SECURITY_AUDITORDISABLED_ACCOUNT EnumdisabledAccountDataSecurityAuditorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:data-security-auditor:disabled-account"
)

// All allowed values of EnumdisabledAccountDataSecurityAuditorSchemaUrn enum
var AllowedEnumdisabledAccountDataSecurityAuditorSchemaUrnEnumValues = []EnumdisabledAccountDataSecurityAuditorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:data-security-auditor:disabled-account",
}

func (v *EnumdisabledAccountDataSecurityAuditorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdisabledAccountDataSecurityAuditorSchemaUrn(value)
	for _, existing := range AllowedEnumdisabledAccountDataSecurityAuditorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdisabledAccountDataSecurityAuditorSchemaUrn", value)
}

// NewEnumdisabledAccountDataSecurityAuditorSchemaUrnFromValue returns a pointer to a valid EnumdisabledAccountDataSecurityAuditorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdisabledAccountDataSecurityAuditorSchemaUrnFromValue(v string) (*EnumdisabledAccountDataSecurityAuditorSchemaUrn, error) {
	ev := EnumdisabledAccountDataSecurityAuditorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdisabledAccountDataSecurityAuditorSchemaUrn: valid values are %v", v, AllowedEnumdisabledAccountDataSecurityAuditorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdisabledAccountDataSecurityAuditorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdisabledAccountDataSecurityAuditorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdisabled-account-data-security-auditorSchemaUrn value
func (v EnumdisabledAccountDataSecurityAuditorSchemaUrn) Ptr() *EnumdisabledAccountDataSecurityAuditorSchemaUrn {
	return &v
}

type NullableEnumdisabledAccountDataSecurityAuditorSchemaUrn struct {
	value *EnumdisabledAccountDataSecurityAuditorSchemaUrn
	isSet bool
}

func (v NullableEnumdisabledAccountDataSecurityAuditorSchemaUrn) Get() *EnumdisabledAccountDataSecurityAuditorSchemaUrn {
	return v.value
}

func (v *NullableEnumdisabledAccountDataSecurityAuditorSchemaUrn) Set(val *EnumdisabledAccountDataSecurityAuditorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdisabledAccountDataSecurityAuditorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdisabledAccountDataSecurityAuditorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdisabledAccountDataSecurityAuditorSchemaUrn(val *EnumdisabledAccountDataSecurityAuditorSchemaUrn) *NullableEnumdisabledAccountDataSecurityAuditorSchemaUrn {
	return &NullableEnumdisabledAccountDataSecurityAuditorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdisabledAccountDataSecurityAuditorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdisabledAccountDataSecurityAuditorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

