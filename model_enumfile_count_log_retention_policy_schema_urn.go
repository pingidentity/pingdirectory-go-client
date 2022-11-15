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

// EnumfileCountLogRetentionPolicySchemaUrn the model 'EnumfileCountLogRetentionPolicySchemaUrn'
type EnumfileCountLogRetentionPolicySchemaUrn string

// List of Enumfile-count-log-retention-policySchemaUrn
const (
	ENUMFILECOUNTLOGRETENTIONPOLICYSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_RETENTION_POLICYFILE_COUNT EnumfileCountLogRetentionPolicySchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-retention-policy:file-count"
)

// All allowed values of EnumfileCountLogRetentionPolicySchemaUrn enum
var AllowedEnumfileCountLogRetentionPolicySchemaUrnEnumValues = []EnumfileCountLogRetentionPolicySchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-retention-policy:file-count",
}

func (v *EnumfileCountLogRetentionPolicySchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumfileCountLogRetentionPolicySchemaUrn(value)
	for _, existing := range AllowedEnumfileCountLogRetentionPolicySchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumfileCountLogRetentionPolicySchemaUrn", value)
}

// NewEnumfileCountLogRetentionPolicySchemaUrnFromValue returns a pointer to a valid EnumfileCountLogRetentionPolicySchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumfileCountLogRetentionPolicySchemaUrnFromValue(v string) (*EnumfileCountLogRetentionPolicySchemaUrn, error) {
	ev := EnumfileCountLogRetentionPolicySchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumfileCountLogRetentionPolicySchemaUrn: valid values are %v", v, AllowedEnumfileCountLogRetentionPolicySchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumfileCountLogRetentionPolicySchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumfileCountLogRetentionPolicySchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumfile-count-log-retention-policySchemaUrn value
func (v EnumfileCountLogRetentionPolicySchemaUrn) Ptr() *EnumfileCountLogRetentionPolicySchemaUrn {
	return &v
}

type NullableEnumfileCountLogRetentionPolicySchemaUrn struct {
	value *EnumfileCountLogRetentionPolicySchemaUrn
	isSet bool
}

func (v NullableEnumfileCountLogRetentionPolicySchemaUrn) Get() *EnumfileCountLogRetentionPolicySchemaUrn {
	return v.value
}

func (v *NullableEnumfileCountLogRetentionPolicySchemaUrn) Set(val *EnumfileCountLogRetentionPolicySchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumfileCountLogRetentionPolicySchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumfileCountLogRetentionPolicySchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumfileCountLogRetentionPolicySchemaUrn(val *EnumfileCountLogRetentionPolicySchemaUrn) *NullableEnumfileCountLogRetentionPolicySchemaUrn {
	return &NullableEnumfileCountLogRetentionPolicySchemaUrn{value: val, isSet: true}
}

func (v NullableEnumfileCountLogRetentionPolicySchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumfileCountLogRetentionPolicySchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

