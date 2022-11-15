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

// EnumaggregateResultCriteriaSchemaUrn the model 'EnumaggregateResultCriteriaSchemaUrn'
type EnumaggregateResultCriteriaSchemaUrn string

// List of Enumaggregate-result-criteriaSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0RESULT_CRITERIAAGGREGATE EnumaggregateResultCriteriaSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:result-criteria:aggregate"
)

// All allowed values of EnumaggregateResultCriteriaSchemaUrn enum
var AllowedEnumaggregateResultCriteriaSchemaUrnEnumValues = []EnumaggregateResultCriteriaSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:result-criteria:aggregate",
}

func (v *EnumaggregateResultCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumaggregateResultCriteriaSchemaUrn(value)
	for _, existing := range AllowedEnumaggregateResultCriteriaSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumaggregateResultCriteriaSchemaUrn", value)
}

// NewEnumaggregateResultCriteriaSchemaUrnFromValue returns a pointer to a valid EnumaggregateResultCriteriaSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumaggregateResultCriteriaSchemaUrnFromValue(v string) (*EnumaggregateResultCriteriaSchemaUrn, error) {
	ev := EnumaggregateResultCriteriaSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumaggregateResultCriteriaSchemaUrn: valid values are %v", v, AllowedEnumaggregateResultCriteriaSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumaggregateResultCriteriaSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumaggregateResultCriteriaSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumaggregate-result-criteriaSchemaUrn value
func (v EnumaggregateResultCriteriaSchemaUrn) Ptr() *EnumaggregateResultCriteriaSchemaUrn {
	return &v
}

type NullableEnumaggregateResultCriteriaSchemaUrn struct {
	value *EnumaggregateResultCriteriaSchemaUrn
	isSet bool
}

func (v NullableEnumaggregateResultCriteriaSchemaUrn) Get() *EnumaggregateResultCriteriaSchemaUrn {
	return v.value
}

func (v *NullableEnumaggregateResultCriteriaSchemaUrn) Set(val *EnumaggregateResultCriteriaSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumaggregateResultCriteriaSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumaggregateResultCriteriaSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumaggregateResultCriteriaSchemaUrn(val *EnumaggregateResultCriteriaSchemaUrn) *NullableEnumaggregateResultCriteriaSchemaUrn {
	return &NullableEnumaggregateResultCriteriaSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumaggregateResultCriteriaSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumaggregateResultCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

