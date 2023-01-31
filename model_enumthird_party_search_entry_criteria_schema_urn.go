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

// EnumthirdPartySearchEntryCriteriaSchemaUrn the model 'EnumthirdPartySearchEntryCriteriaSchemaUrn'
type EnumthirdPartySearchEntryCriteriaSchemaUrn string

// List of Enumthird-party-search-entry-criteriaSchemaUrn
const (
	ENUMTHIRDPARTYSEARCHENTRYCRITERIASCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0SEARCH_ENTRY_CRITERIATHIRD_PARTY EnumthirdPartySearchEntryCriteriaSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:search-entry-criteria:third-party"
)

// All allowed values of EnumthirdPartySearchEntryCriteriaSchemaUrn enum
var AllowedEnumthirdPartySearchEntryCriteriaSchemaUrnEnumValues = []EnumthirdPartySearchEntryCriteriaSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:search-entry-criteria:third-party",
}

func (v *EnumthirdPartySearchEntryCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumthirdPartySearchEntryCriteriaSchemaUrn(value)
	for _, existing := range AllowedEnumthirdPartySearchEntryCriteriaSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumthirdPartySearchEntryCriteriaSchemaUrn", value)
}

// NewEnumthirdPartySearchEntryCriteriaSchemaUrnFromValue returns a pointer to a valid EnumthirdPartySearchEntryCriteriaSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumthirdPartySearchEntryCriteriaSchemaUrnFromValue(v string) (*EnumthirdPartySearchEntryCriteriaSchemaUrn, error) {
	ev := EnumthirdPartySearchEntryCriteriaSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumthirdPartySearchEntryCriteriaSchemaUrn: valid values are %v", v, AllowedEnumthirdPartySearchEntryCriteriaSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumthirdPartySearchEntryCriteriaSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumthirdPartySearchEntryCriteriaSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumthird-party-search-entry-criteriaSchemaUrn value
func (v EnumthirdPartySearchEntryCriteriaSchemaUrn) Ptr() *EnumthirdPartySearchEntryCriteriaSchemaUrn {
	return &v
}

type NullableEnumthirdPartySearchEntryCriteriaSchemaUrn struct {
	value *EnumthirdPartySearchEntryCriteriaSchemaUrn
	isSet bool
}

func (v NullableEnumthirdPartySearchEntryCriteriaSchemaUrn) Get() *EnumthirdPartySearchEntryCriteriaSchemaUrn {
	return v.value
}

func (v *NullableEnumthirdPartySearchEntryCriteriaSchemaUrn) Set(val *EnumthirdPartySearchEntryCriteriaSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumthirdPartySearchEntryCriteriaSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumthirdPartySearchEntryCriteriaSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumthirdPartySearchEntryCriteriaSchemaUrn(val *EnumthirdPartySearchEntryCriteriaSchemaUrn) *NullableEnumthirdPartySearchEntryCriteriaSchemaUrn {
	return &NullableEnumthirdPartySearchEntryCriteriaSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumthirdPartySearchEntryCriteriaSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumthirdPartySearchEntryCriteriaSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
