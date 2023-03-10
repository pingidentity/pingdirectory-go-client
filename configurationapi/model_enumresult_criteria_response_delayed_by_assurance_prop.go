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

// EnumresultCriteriaResponseDelayedByAssuranceProp Indicates whether this Replication Assurance Result Criteria should match operations based on whether the response to the client was delayed by assurance processing.
type EnumresultCriteriaResponseDelayedByAssuranceProp string

// List of Enumresult-criteria-responseDelayedByAssuranceProp
const (
	ENUMRESULTCRITERIARESPONSEDELAYEDBYASSURANCEPROP_ANY   EnumresultCriteriaResponseDelayedByAssuranceProp = "any"
	ENUMRESULTCRITERIARESPONSEDELAYEDBYASSURANCEPROP_TRUE  EnumresultCriteriaResponseDelayedByAssuranceProp = "true"
	ENUMRESULTCRITERIARESPONSEDELAYEDBYASSURANCEPROP_FALSE EnumresultCriteriaResponseDelayedByAssuranceProp = "false"
)

// All allowed values of EnumresultCriteriaResponseDelayedByAssuranceProp enum
var AllowedEnumresultCriteriaResponseDelayedByAssurancePropEnumValues = []EnumresultCriteriaResponseDelayedByAssuranceProp{
	"any",
	"true",
	"false",
}

func (v *EnumresultCriteriaResponseDelayedByAssuranceProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumresultCriteriaResponseDelayedByAssuranceProp(value)
	for _, existing := range AllowedEnumresultCriteriaResponseDelayedByAssurancePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumresultCriteriaResponseDelayedByAssuranceProp", value)
}

// NewEnumresultCriteriaResponseDelayedByAssurancePropFromValue returns a pointer to a valid EnumresultCriteriaResponseDelayedByAssuranceProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumresultCriteriaResponseDelayedByAssurancePropFromValue(v string) (*EnumresultCriteriaResponseDelayedByAssuranceProp, error) {
	ev := EnumresultCriteriaResponseDelayedByAssuranceProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumresultCriteriaResponseDelayedByAssuranceProp: valid values are %v", v, AllowedEnumresultCriteriaResponseDelayedByAssurancePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumresultCriteriaResponseDelayedByAssuranceProp) IsValid() bool {
	for _, existing := range AllowedEnumresultCriteriaResponseDelayedByAssurancePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumresult-criteria-responseDelayedByAssuranceProp value
func (v EnumresultCriteriaResponseDelayedByAssuranceProp) Ptr() *EnumresultCriteriaResponseDelayedByAssuranceProp {
	return &v
}

type NullableEnumresultCriteriaResponseDelayedByAssuranceProp struct {
	value *EnumresultCriteriaResponseDelayedByAssuranceProp
	isSet bool
}

func (v NullableEnumresultCriteriaResponseDelayedByAssuranceProp) Get() *EnumresultCriteriaResponseDelayedByAssuranceProp {
	return v.value
}

func (v *NullableEnumresultCriteriaResponseDelayedByAssuranceProp) Set(val *EnumresultCriteriaResponseDelayedByAssuranceProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumresultCriteriaResponseDelayedByAssuranceProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumresultCriteriaResponseDelayedByAssuranceProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumresultCriteriaResponseDelayedByAssuranceProp(val *EnumresultCriteriaResponseDelayedByAssuranceProp) *NullableEnumresultCriteriaResponseDelayedByAssuranceProp {
	return &NullableEnumresultCriteriaResponseDelayedByAssuranceProp{value: val, isSet: true}
}

func (v NullableEnumresultCriteriaResponseDelayedByAssuranceProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumresultCriteriaResponseDelayedByAssuranceProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
