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

// EnumaccessControlDataSecurityAuditorSchemaUrn the model 'EnumaccessControlDataSecurityAuditorSchemaUrn'
type EnumaccessControlDataSecurityAuditorSchemaUrn string

// List of Enumaccess-control-data-security-auditorSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0DATA_SECURITY_AUDITORACCESS_CONTROL EnumaccessControlDataSecurityAuditorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:data-security-auditor:access-control"
)

// All allowed values of EnumaccessControlDataSecurityAuditorSchemaUrn enum
var AllowedEnumaccessControlDataSecurityAuditorSchemaUrnEnumValues = []EnumaccessControlDataSecurityAuditorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:data-security-auditor:access-control",
}

func (v *EnumaccessControlDataSecurityAuditorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumaccessControlDataSecurityAuditorSchemaUrn(value)
	for _, existing := range AllowedEnumaccessControlDataSecurityAuditorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumaccessControlDataSecurityAuditorSchemaUrn", value)
}

// NewEnumaccessControlDataSecurityAuditorSchemaUrnFromValue returns a pointer to a valid EnumaccessControlDataSecurityAuditorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumaccessControlDataSecurityAuditorSchemaUrnFromValue(v string) (*EnumaccessControlDataSecurityAuditorSchemaUrn, error) {
	ev := EnumaccessControlDataSecurityAuditorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumaccessControlDataSecurityAuditorSchemaUrn: valid values are %v", v, AllowedEnumaccessControlDataSecurityAuditorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumaccessControlDataSecurityAuditorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumaccessControlDataSecurityAuditorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumaccess-control-data-security-auditorSchemaUrn value
func (v EnumaccessControlDataSecurityAuditorSchemaUrn) Ptr() *EnumaccessControlDataSecurityAuditorSchemaUrn {
	return &v
}

type NullableEnumaccessControlDataSecurityAuditorSchemaUrn struct {
	value *EnumaccessControlDataSecurityAuditorSchemaUrn
	isSet bool
}

func (v NullableEnumaccessControlDataSecurityAuditorSchemaUrn) Get() *EnumaccessControlDataSecurityAuditorSchemaUrn {
	return v.value
}

func (v *NullableEnumaccessControlDataSecurityAuditorSchemaUrn) Set(val *EnumaccessControlDataSecurityAuditorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumaccessControlDataSecurityAuditorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumaccessControlDataSecurityAuditorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumaccessControlDataSecurityAuditorSchemaUrn(val *EnumaccessControlDataSecurityAuditorSchemaUrn) *NullableEnumaccessControlDataSecurityAuditorSchemaUrn {
	return &NullableEnumaccessControlDataSecurityAuditorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumaccessControlDataSecurityAuditorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumaccessControlDataSecurityAuditorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

