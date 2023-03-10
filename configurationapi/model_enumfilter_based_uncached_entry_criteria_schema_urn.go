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

// EnumfilterBasedUncachedEntryCriteriaSchemaUrn the model 'EnumfilterBasedUncachedEntryCriteriaSchemaUrn'
type EnumfilterBasedUncachedEntryCriteriaSchemaUrn string

// List of Enumfilter-based-uncached-entry-criteriaSchemaUrn
const (
	ENUMFILTERBASEDUNCACHEDENTRYCRITERIASCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0UNCACHED_ENTRY_CRITERIAFILTER_BASED EnumfilterBasedUncachedEntryCriteriaSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:uncached-entry-criteria:filter-based"
)

// All allowed values of EnumfilterBasedUncachedEntryCriteriaSchemaUrn enum
var AllowedEnumfilterBasedUncachedEntryCriteriaSchemaUrnEnumValues = []EnumfilterBasedUncachedEntryCriteriaSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:uncached-entry-criteria:filter-based",
}

func (v *EnumfilterBasedUncachedEntryCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumfilterBasedUncachedEntryCriteriaSchemaUrn(value)
	for _, existing := range AllowedEnumfilterBasedUncachedEntryCriteriaSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumfilterBasedUncachedEntryCriteriaSchemaUrn", value)
}

// NewEnumfilterBasedUncachedEntryCriteriaSchemaUrnFromValue returns a pointer to a valid EnumfilterBasedUncachedEntryCriteriaSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumfilterBasedUncachedEntryCriteriaSchemaUrnFromValue(v string) (*EnumfilterBasedUncachedEntryCriteriaSchemaUrn, error) {
	ev := EnumfilterBasedUncachedEntryCriteriaSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumfilterBasedUncachedEntryCriteriaSchemaUrn: valid values are %v", v, AllowedEnumfilterBasedUncachedEntryCriteriaSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumfilterBasedUncachedEntryCriteriaSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumfilterBasedUncachedEntryCriteriaSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumfilter-based-uncached-entry-criteriaSchemaUrn value
func (v EnumfilterBasedUncachedEntryCriteriaSchemaUrn) Ptr() *EnumfilterBasedUncachedEntryCriteriaSchemaUrn {
	return &v
}

type NullableEnumfilterBasedUncachedEntryCriteriaSchemaUrn struct {
	value *EnumfilterBasedUncachedEntryCriteriaSchemaUrn
	isSet bool
}

func (v NullableEnumfilterBasedUncachedEntryCriteriaSchemaUrn) Get() *EnumfilterBasedUncachedEntryCriteriaSchemaUrn {
	return v.value
}

func (v *NullableEnumfilterBasedUncachedEntryCriteriaSchemaUrn) Set(val *EnumfilterBasedUncachedEntryCriteriaSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumfilterBasedUncachedEntryCriteriaSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumfilterBasedUncachedEntryCriteriaSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumfilterBasedUncachedEntryCriteriaSchemaUrn(val *EnumfilterBasedUncachedEntryCriteriaSchemaUrn) *NullableEnumfilterBasedUncachedEntryCriteriaSchemaUrn {
	return &NullableEnumfilterBasedUncachedEntryCriteriaSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumfilterBasedUncachedEntryCriteriaSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumfilterBasedUncachedEntryCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
