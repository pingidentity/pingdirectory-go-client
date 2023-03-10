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

// EnumdefaultUncachedEntryCriteriaSchemaUrn the model 'EnumdefaultUncachedEntryCriteriaSchemaUrn'
type EnumdefaultUncachedEntryCriteriaSchemaUrn string

// List of Enumdefault-uncached-entry-criteriaSchemaUrn
const (
	ENUMDEFAULTUNCACHEDENTRYCRITERIASCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0UNCACHED_ENTRY_CRITERIADEFAULT EnumdefaultUncachedEntryCriteriaSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:uncached-entry-criteria:default"
)

// All allowed values of EnumdefaultUncachedEntryCriteriaSchemaUrn enum
var AllowedEnumdefaultUncachedEntryCriteriaSchemaUrnEnumValues = []EnumdefaultUncachedEntryCriteriaSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:uncached-entry-criteria:default",
}

func (v *EnumdefaultUncachedEntryCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdefaultUncachedEntryCriteriaSchemaUrn(value)
	for _, existing := range AllowedEnumdefaultUncachedEntryCriteriaSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdefaultUncachedEntryCriteriaSchemaUrn", value)
}

// NewEnumdefaultUncachedEntryCriteriaSchemaUrnFromValue returns a pointer to a valid EnumdefaultUncachedEntryCriteriaSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdefaultUncachedEntryCriteriaSchemaUrnFromValue(v string) (*EnumdefaultUncachedEntryCriteriaSchemaUrn, error) {
	ev := EnumdefaultUncachedEntryCriteriaSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdefaultUncachedEntryCriteriaSchemaUrn: valid values are %v", v, AllowedEnumdefaultUncachedEntryCriteriaSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdefaultUncachedEntryCriteriaSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdefaultUncachedEntryCriteriaSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdefault-uncached-entry-criteriaSchemaUrn value
func (v EnumdefaultUncachedEntryCriteriaSchemaUrn) Ptr() *EnumdefaultUncachedEntryCriteriaSchemaUrn {
	return &v
}

type NullableEnumdefaultUncachedEntryCriteriaSchemaUrn struct {
	value *EnumdefaultUncachedEntryCriteriaSchemaUrn
	isSet bool
}

func (v NullableEnumdefaultUncachedEntryCriteriaSchemaUrn) Get() *EnumdefaultUncachedEntryCriteriaSchemaUrn {
	return v.value
}

func (v *NullableEnumdefaultUncachedEntryCriteriaSchemaUrn) Set(val *EnumdefaultUncachedEntryCriteriaSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdefaultUncachedEntryCriteriaSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdefaultUncachedEntryCriteriaSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdefaultUncachedEntryCriteriaSchemaUrn(val *EnumdefaultUncachedEntryCriteriaSchemaUrn) *NullableEnumdefaultUncachedEntryCriteriaSchemaUrn {
	return &NullableEnumdefaultUncachedEntryCriteriaSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdefaultUncachedEntryCriteriaSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdefaultUncachedEntryCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
