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

// EnumlockedAccountDataSecurityAuditorSchemaUrn the model 'EnumlockedAccountDataSecurityAuditorSchemaUrn'
type EnumlockedAccountDataSecurityAuditorSchemaUrn string

// List of Enumlocked-account-data-security-auditorSchemaUrn
const (
	ENUMLOCKEDACCOUNTDATASECURITYAUDITORSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0DATA_SECURITY_AUDITORLOCKED_ACCOUNT EnumlockedAccountDataSecurityAuditorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:data-security-auditor:locked-account"
)

// All allowed values of EnumlockedAccountDataSecurityAuditorSchemaUrn enum
var AllowedEnumlockedAccountDataSecurityAuditorSchemaUrnEnumValues = []EnumlockedAccountDataSecurityAuditorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:data-security-auditor:locked-account",
}

func (v *EnumlockedAccountDataSecurityAuditorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlockedAccountDataSecurityAuditorSchemaUrn(value)
	for _, existing := range AllowedEnumlockedAccountDataSecurityAuditorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlockedAccountDataSecurityAuditorSchemaUrn", value)
}

// NewEnumlockedAccountDataSecurityAuditorSchemaUrnFromValue returns a pointer to a valid EnumlockedAccountDataSecurityAuditorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlockedAccountDataSecurityAuditorSchemaUrnFromValue(v string) (*EnumlockedAccountDataSecurityAuditorSchemaUrn, error) {
	ev := EnumlockedAccountDataSecurityAuditorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlockedAccountDataSecurityAuditorSchemaUrn: valid values are %v", v, AllowedEnumlockedAccountDataSecurityAuditorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlockedAccountDataSecurityAuditorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumlockedAccountDataSecurityAuditorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlocked-account-data-security-auditorSchemaUrn value
func (v EnumlockedAccountDataSecurityAuditorSchemaUrn) Ptr() *EnumlockedAccountDataSecurityAuditorSchemaUrn {
	return &v
}

type NullableEnumlockedAccountDataSecurityAuditorSchemaUrn struct {
	value *EnumlockedAccountDataSecurityAuditorSchemaUrn
	isSet bool
}

func (v NullableEnumlockedAccountDataSecurityAuditorSchemaUrn) Get() *EnumlockedAccountDataSecurityAuditorSchemaUrn {
	return v.value
}

func (v *NullableEnumlockedAccountDataSecurityAuditorSchemaUrn) Set(val *EnumlockedAccountDataSecurityAuditorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlockedAccountDataSecurityAuditorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlockedAccountDataSecurityAuditorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlockedAccountDataSecurityAuditorSchemaUrn(val *EnumlockedAccountDataSecurityAuditorSchemaUrn) *NullableEnumlockedAccountDataSecurityAuditorSchemaUrn {
	return &NullableEnumlockedAccountDataSecurityAuditorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumlockedAccountDataSecurityAuditorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlockedAccountDataSecurityAuditorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
