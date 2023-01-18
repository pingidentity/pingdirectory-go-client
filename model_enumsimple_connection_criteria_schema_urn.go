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

// EnumsimpleConnectionCriteriaSchemaUrn the model 'EnumsimpleConnectionCriteriaSchemaUrn'
type EnumsimpleConnectionCriteriaSchemaUrn string

// List of Enumsimple-connection-criteriaSchemaUrn
const (
	ENUMSIMPLECONNECTIONCRITERIASCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CONNECTION_CRITERIASIMPLE EnumsimpleConnectionCriteriaSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:connection-criteria:simple"
)

// All allowed values of EnumsimpleConnectionCriteriaSchemaUrn enum
var AllowedEnumsimpleConnectionCriteriaSchemaUrnEnumValues = []EnumsimpleConnectionCriteriaSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:connection-criteria:simple",
}

func (v *EnumsimpleConnectionCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsimpleConnectionCriteriaSchemaUrn(value)
	for _, existing := range AllowedEnumsimpleConnectionCriteriaSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsimpleConnectionCriteriaSchemaUrn", value)
}

// NewEnumsimpleConnectionCriteriaSchemaUrnFromValue returns a pointer to a valid EnumsimpleConnectionCriteriaSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsimpleConnectionCriteriaSchemaUrnFromValue(v string) (*EnumsimpleConnectionCriteriaSchemaUrn, error) {
	ev := EnumsimpleConnectionCriteriaSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsimpleConnectionCriteriaSchemaUrn: valid values are %v", v, AllowedEnumsimpleConnectionCriteriaSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsimpleConnectionCriteriaSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsimpleConnectionCriteriaSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsimple-connection-criteriaSchemaUrn value
func (v EnumsimpleConnectionCriteriaSchemaUrn) Ptr() *EnumsimpleConnectionCriteriaSchemaUrn {
	return &v
}

type NullableEnumsimpleConnectionCriteriaSchemaUrn struct {
	value *EnumsimpleConnectionCriteriaSchemaUrn
	isSet bool
}

func (v NullableEnumsimpleConnectionCriteriaSchemaUrn) Get() *EnumsimpleConnectionCriteriaSchemaUrn {
	return v.value
}

func (v *NullableEnumsimpleConnectionCriteriaSchemaUrn) Set(val *EnumsimpleConnectionCriteriaSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsimpleConnectionCriteriaSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsimpleConnectionCriteriaSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsimpleConnectionCriteriaSchemaUrn(val *EnumsimpleConnectionCriteriaSchemaUrn) *NullableEnumsimpleConnectionCriteriaSchemaUrn {
	return &NullableEnumsimpleConnectionCriteriaSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsimpleConnectionCriteriaSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsimpleConnectionCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
