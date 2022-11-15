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

// EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp Indicates whether operations being processed using a dedicated administrative session worker thread should be included in this Simple Request Criteria.
type EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp string

// List of Enumrequest-criteria-usingAdministrativeSessionWorkerThreadProp
const (
	ENUMREQUESTCRITERIAUSINGADMINISTRATIVESESSIONWORKERTHREADPROP_TRUE EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp = "true"
	ENUMREQUESTCRITERIAUSINGADMINISTRATIVESESSIONWORKERTHREADPROP_FALSE EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp = "false"
	ENUMREQUESTCRITERIAUSINGADMINISTRATIVESESSIONWORKERTHREADPROP_ANY EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp = "any"
)

// All allowed values of EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp enum
var AllowedEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadPropEnumValues = []EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp{
	"true",
	"false",
	"any",
}

func (v *EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp(value)
	for _, existing := range AllowedEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp", value)
}

// NewEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadPropFromValue returns a pointer to a valid EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadPropFromValue(v string) (*EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp, error) {
	ev := EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp: valid values are %v", v, AllowedEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp) IsValid() bool {
	for _, existing := range AllowedEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumrequest-criteria-usingAdministrativeSessionWorkerThreadProp value
func (v EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp) Ptr() *EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp {
	return &v
}

type NullableEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp struct {
	value *EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp
	isSet bool
}

func (v NullableEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp) Get() *EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp {
	return v.value
}

func (v *NullableEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp) Set(val *EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp(val *EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp) *NullableEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp {
	return &NullableEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp{value: val, isSet: true}
}

func (v NullableEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

