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

// EnumaggregateConnectionCriteriaSchemaUrn the model 'EnumaggregateConnectionCriteriaSchemaUrn'
type EnumaggregateConnectionCriteriaSchemaUrn string

// List of Enumaggregate-connection-criteriaSchemaUrn
const (
	ENUMAGGREGATECONNECTIONCRITERIASCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CONNECTION_CRITERIAAGGREGATE EnumaggregateConnectionCriteriaSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:connection-criteria:aggregate"
)

// All allowed values of EnumaggregateConnectionCriteriaSchemaUrn enum
var AllowedEnumaggregateConnectionCriteriaSchemaUrnEnumValues = []EnumaggregateConnectionCriteriaSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:connection-criteria:aggregate",
}

func (v *EnumaggregateConnectionCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumaggregateConnectionCriteriaSchemaUrn(value)
	for _, existing := range AllowedEnumaggregateConnectionCriteriaSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumaggregateConnectionCriteriaSchemaUrn", value)
}

// NewEnumaggregateConnectionCriteriaSchemaUrnFromValue returns a pointer to a valid EnumaggregateConnectionCriteriaSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumaggregateConnectionCriteriaSchemaUrnFromValue(v string) (*EnumaggregateConnectionCriteriaSchemaUrn, error) {
	ev := EnumaggregateConnectionCriteriaSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumaggregateConnectionCriteriaSchemaUrn: valid values are %v", v, AllowedEnumaggregateConnectionCriteriaSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumaggregateConnectionCriteriaSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumaggregateConnectionCriteriaSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumaggregate-connection-criteriaSchemaUrn value
func (v EnumaggregateConnectionCriteriaSchemaUrn) Ptr() *EnumaggregateConnectionCriteriaSchemaUrn {
	return &v
}

type NullableEnumaggregateConnectionCriteriaSchemaUrn struct {
	value *EnumaggregateConnectionCriteriaSchemaUrn
	isSet bool
}

func (v NullableEnumaggregateConnectionCriteriaSchemaUrn) Get() *EnumaggregateConnectionCriteriaSchemaUrn {
	return v.value
}

func (v *NullableEnumaggregateConnectionCriteriaSchemaUrn) Set(val *EnumaggregateConnectionCriteriaSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumaggregateConnectionCriteriaSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumaggregateConnectionCriteriaSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumaggregateConnectionCriteriaSchemaUrn(val *EnumaggregateConnectionCriteriaSchemaUrn) *NullableEnumaggregateConnectionCriteriaSchemaUrn {
	return &NullableEnumaggregateConnectionCriteriaSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumaggregateConnectionCriteriaSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumaggregateConnectionCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
