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

// EnumsimpleResultCriteriaSchemaUrn the model 'EnumsimpleResultCriteriaSchemaUrn'
type EnumsimpleResultCriteriaSchemaUrn string

// List of Enumsimple-result-criteriaSchemaUrn
const (
	ENUMSIMPLERESULTCRITERIASCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0RESULT_CRITERIASIMPLE EnumsimpleResultCriteriaSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:result-criteria:simple"
)

// All allowed values of EnumsimpleResultCriteriaSchemaUrn enum
var AllowedEnumsimpleResultCriteriaSchemaUrnEnumValues = []EnumsimpleResultCriteriaSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:result-criteria:simple",
}

func (v *EnumsimpleResultCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsimpleResultCriteriaSchemaUrn(value)
	for _, existing := range AllowedEnumsimpleResultCriteriaSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsimpleResultCriteriaSchemaUrn", value)
}

// NewEnumsimpleResultCriteriaSchemaUrnFromValue returns a pointer to a valid EnumsimpleResultCriteriaSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsimpleResultCriteriaSchemaUrnFromValue(v string) (*EnumsimpleResultCriteriaSchemaUrn, error) {
	ev := EnumsimpleResultCriteriaSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsimpleResultCriteriaSchemaUrn: valid values are %v", v, AllowedEnumsimpleResultCriteriaSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsimpleResultCriteriaSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsimpleResultCriteriaSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsimple-result-criteriaSchemaUrn value
func (v EnumsimpleResultCriteriaSchemaUrn) Ptr() *EnumsimpleResultCriteriaSchemaUrn {
	return &v
}

type NullableEnumsimpleResultCriteriaSchemaUrn struct {
	value *EnumsimpleResultCriteriaSchemaUrn
	isSet bool
}

func (v NullableEnumsimpleResultCriteriaSchemaUrn) Get() *EnumsimpleResultCriteriaSchemaUrn {
	return v.value
}

func (v *NullableEnumsimpleResultCriteriaSchemaUrn) Set(val *EnumsimpleResultCriteriaSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsimpleResultCriteriaSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsimpleResultCriteriaSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsimpleResultCriteriaSchemaUrn(val *EnumsimpleResultCriteriaSchemaUrn) *NullableEnumsimpleResultCriteriaSchemaUrn {
	return &NullableEnumsimpleResultCriteriaSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsimpleResultCriteriaSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsimpleResultCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
