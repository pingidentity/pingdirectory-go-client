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

// EnumaggregateSearchReferenceCriteriaSchemaUrn the model 'EnumaggregateSearchReferenceCriteriaSchemaUrn'
type EnumaggregateSearchReferenceCriteriaSchemaUrn string

// List of Enumaggregate-search-reference-criteriaSchemaUrn
const (
	ENUMAGGREGATESEARCHREFERENCECRITERIASCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0SEARCH_REFERENCE_CRITERIAAGGREGATE EnumaggregateSearchReferenceCriteriaSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:search-reference-criteria:aggregate"
)

// All allowed values of EnumaggregateSearchReferenceCriteriaSchemaUrn enum
var AllowedEnumaggregateSearchReferenceCriteriaSchemaUrnEnumValues = []EnumaggregateSearchReferenceCriteriaSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:search-reference-criteria:aggregate",
}

func (v *EnumaggregateSearchReferenceCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumaggregateSearchReferenceCriteriaSchemaUrn(value)
	for _, existing := range AllowedEnumaggregateSearchReferenceCriteriaSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumaggregateSearchReferenceCriteriaSchemaUrn", value)
}

// NewEnumaggregateSearchReferenceCriteriaSchemaUrnFromValue returns a pointer to a valid EnumaggregateSearchReferenceCriteriaSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumaggregateSearchReferenceCriteriaSchemaUrnFromValue(v string) (*EnumaggregateSearchReferenceCriteriaSchemaUrn, error) {
	ev := EnumaggregateSearchReferenceCriteriaSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumaggregateSearchReferenceCriteriaSchemaUrn: valid values are %v", v, AllowedEnumaggregateSearchReferenceCriteriaSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumaggregateSearchReferenceCriteriaSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumaggregateSearchReferenceCriteriaSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumaggregate-search-reference-criteriaSchemaUrn value
func (v EnumaggregateSearchReferenceCriteriaSchemaUrn) Ptr() *EnumaggregateSearchReferenceCriteriaSchemaUrn {
	return &v
}

type NullableEnumaggregateSearchReferenceCriteriaSchemaUrn struct {
	value *EnumaggregateSearchReferenceCriteriaSchemaUrn
	isSet bool
}

func (v NullableEnumaggregateSearchReferenceCriteriaSchemaUrn) Get() *EnumaggregateSearchReferenceCriteriaSchemaUrn {
	return v.value
}

func (v *NullableEnumaggregateSearchReferenceCriteriaSchemaUrn) Set(val *EnumaggregateSearchReferenceCriteriaSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumaggregateSearchReferenceCriteriaSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumaggregateSearchReferenceCriteriaSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumaggregateSearchReferenceCriteriaSchemaUrn(val *EnumaggregateSearchReferenceCriteriaSchemaUrn) *NullableEnumaggregateSearchReferenceCriteriaSchemaUrn {
	return &NullableEnumaggregateSearchReferenceCriteriaSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumaggregateSearchReferenceCriteriaSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumaggregateSearchReferenceCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

