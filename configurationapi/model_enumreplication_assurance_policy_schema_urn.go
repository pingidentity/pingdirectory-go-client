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

// EnumreplicationAssurancePolicySchemaUrn the model 'EnumreplicationAssurancePolicySchemaUrn'
type EnumreplicationAssurancePolicySchemaUrn string

// List of Enumreplication-assurance-policySchemaUrn
const (
	ENUMREPLICATIONASSURANCEPOLICYSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0REPLICATION_ASSURANCE_POLICY EnumreplicationAssurancePolicySchemaUrn = "urn:pingidentity:schemas:configuration:2.0:replication-assurance-policy"
)

// All allowed values of EnumreplicationAssurancePolicySchemaUrn enum
var AllowedEnumreplicationAssurancePolicySchemaUrnEnumValues = []EnumreplicationAssurancePolicySchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:replication-assurance-policy",
}

func (v *EnumreplicationAssurancePolicySchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumreplicationAssurancePolicySchemaUrn(value)
	for _, existing := range AllowedEnumreplicationAssurancePolicySchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumreplicationAssurancePolicySchemaUrn", value)
}

// NewEnumreplicationAssurancePolicySchemaUrnFromValue returns a pointer to a valid EnumreplicationAssurancePolicySchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumreplicationAssurancePolicySchemaUrnFromValue(v string) (*EnumreplicationAssurancePolicySchemaUrn, error) {
	ev := EnumreplicationAssurancePolicySchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumreplicationAssurancePolicySchemaUrn: valid values are %v", v, AllowedEnumreplicationAssurancePolicySchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumreplicationAssurancePolicySchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumreplicationAssurancePolicySchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumreplication-assurance-policySchemaUrn value
func (v EnumreplicationAssurancePolicySchemaUrn) Ptr() *EnumreplicationAssurancePolicySchemaUrn {
	return &v
}

type NullableEnumreplicationAssurancePolicySchemaUrn struct {
	value *EnumreplicationAssurancePolicySchemaUrn
	isSet bool
}

func (v NullableEnumreplicationAssurancePolicySchemaUrn) Get() *EnumreplicationAssurancePolicySchemaUrn {
	return v.value
}

func (v *NullableEnumreplicationAssurancePolicySchemaUrn) Set(val *EnumreplicationAssurancePolicySchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumreplicationAssurancePolicySchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumreplicationAssurancePolicySchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumreplicationAssurancePolicySchemaUrn(val *EnumreplicationAssurancePolicySchemaUrn) *NullableEnumreplicationAssurancePolicySchemaUrn {
	return &NullableEnumreplicationAssurancePolicySchemaUrn{value: val, isSet: true}
}

func (v NullableEnumreplicationAssurancePolicySchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumreplicationAssurancePolicySchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
