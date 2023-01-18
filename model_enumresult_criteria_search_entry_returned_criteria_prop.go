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

// EnumresultCriteriaSearchEntryReturnedCriteriaProp Indicates whether the number of entries returned should be considered when determining whether a search operation should be included in this Simple Result Criteria. This will be ignored for all operations other than search.
type EnumresultCriteriaSearchEntryReturnedCriteriaProp string

// List of Enumresult-criteria-searchEntryReturnedCriteriaProp
const (
	ENUMRESULTCRITERIASEARCHENTRYRETURNEDCRITERIAPROP_ANY EnumresultCriteriaSearchEntryReturnedCriteriaProp = "any"
	ENUMRESULTCRITERIASEARCHENTRYRETURNEDCRITERIAPROP_EQUAL_TO EnumresultCriteriaSearchEntryReturnedCriteriaProp = "equal-to"
	ENUMRESULTCRITERIASEARCHENTRYRETURNEDCRITERIAPROP_NOT_EQUAL_TO EnumresultCriteriaSearchEntryReturnedCriteriaProp = "not-equal-to"
	ENUMRESULTCRITERIASEARCHENTRYRETURNEDCRITERIAPROP_LESS_THAN_OR_EQUAL_TO EnumresultCriteriaSearchEntryReturnedCriteriaProp = "less-than-or-equal-to"
	ENUMRESULTCRITERIASEARCHENTRYRETURNEDCRITERIAPROP_GREATER_THAN_OR_EQUAL_TO EnumresultCriteriaSearchEntryReturnedCriteriaProp = "greater-than-or-equal-to"
)

// All allowed values of EnumresultCriteriaSearchEntryReturnedCriteriaProp enum
var AllowedEnumresultCriteriaSearchEntryReturnedCriteriaPropEnumValues = []EnumresultCriteriaSearchEntryReturnedCriteriaProp{
	"any",
	"equal-to",
	"not-equal-to",
	"less-than-or-equal-to",
	"greater-than-or-equal-to",
}

func (v *EnumresultCriteriaSearchEntryReturnedCriteriaProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumresultCriteriaSearchEntryReturnedCriteriaProp(value)
	for _, existing := range AllowedEnumresultCriteriaSearchEntryReturnedCriteriaPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumresultCriteriaSearchEntryReturnedCriteriaProp", value)
}

// NewEnumresultCriteriaSearchEntryReturnedCriteriaPropFromValue returns a pointer to a valid EnumresultCriteriaSearchEntryReturnedCriteriaProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumresultCriteriaSearchEntryReturnedCriteriaPropFromValue(v string) (*EnumresultCriteriaSearchEntryReturnedCriteriaProp, error) {
	ev := EnumresultCriteriaSearchEntryReturnedCriteriaProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumresultCriteriaSearchEntryReturnedCriteriaProp: valid values are %v", v, AllowedEnumresultCriteriaSearchEntryReturnedCriteriaPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumresultCriteriaSearchEntryReturnedCriteriaProp) IsValid() bool {
	for _, existing := range AllowedEnumresultCriteriaSearchEntryReturnedCriteriaPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumresult-criteria-searchEntryReturnedCriteriaProp value
func (v EnumresultCriteriaSearchEntryReturnedCriteriaProp) Ptr() *EnumresultCriteriaSearchEntryReturnedCriteriaProp {
	return &v
}

type NullableEnumresultCriteriaSearchEntryReturnedCriteriaProp struct {
	value *EnumresultCriteriaSearchEntryReturnedCriteriaProp
	isSet bool
}

func (v NullableEnumresultCriteriaSearchEntryReturnedCriteriaProp) Get() *EnumresultCriteriaSearchEntryReturnedCriteriaProp {
	return v.value
}

func (v *NullableEnumresultCriteriaSearchEntryReturnedCriteriaProp) Set(val *EnumresultCriteriaSearchEntryReturnedCriteriaProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumresultCriteriaSearchEntryReturnedCriteriaProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumresultCriteriaSearchEntryReturnedCriteriaProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumresultCriteriaSearchEntryReturnedCriteriaProp(val *EnumresultCriteriaSearchEntryReturnedCriteriaProp) *NullableEnumresultCriteriaSearchEntryReturnedCriteriaProp {
	return &NullableEnumresultCriteriaSearchEntryReturnedCriteriaProp{value: val, isSet: true}
}

func (v NullableEnumresultCriteriaSearchEntryReturnedCriteriaProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumresultCriteriaSearchEntryReturnedCriteriaProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
