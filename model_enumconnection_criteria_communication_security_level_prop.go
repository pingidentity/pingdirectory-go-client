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

// EnumconnectionCriteriaCommunicationSecurityLevelProp Indicates whether this Simple Connection Criteria should require or allow clients using a secure communication channel.
type EnumconnectionCriteriaCommunicationSecurityLevelProp string

// List of Enumconnection-criteria-communicationSecurityLevelProp
const (
	ENUMCONNECTIONCRITERIACOMMUNICATIONSECURITYLEVELPROP_ANY EnumconnectionCriteriaCommunicationSecurityLevelProp = "any"
	ENUMCONNECTIONCRITERIACOMMUNICATIONSECURITYLEVELPROP_SECURE_ONLY EnumconnectionCriteriaCommunicationSecurityLevelProp = "secure-only"
	ENUMCONNECTIONCRITERIACOMMUNICATIONSECURITYLEVELPROP_INSECURE_ONLY EnumconnectionCriteriaCommunicationSecurityLevelProp = "insecure-only"
)

// All allowed values of EnumconnectionCriteriaCommunicationSecurityLevelProp enum
var AllowedEnumconnectionCriteriaCommunicationSecurityLevelPropEnumValues = []EnumconnectionCriteriaCommunicationSecurityLevelProp{
	"any",
	"secure-only",
	"insecure-only",
}

func (v *EnumconnectionCriteriaCommunicationSecurityLevelProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumconnectionCriteriaCommunicationSecurityLevelProp(value)
	for _, existing := range AllowedEnumconnectionCriteriaCommunicationSecurityLevelPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumconnectionCriteriaCommunicationSecurityLevelProp", value)
}

// NewEnumconnectionCriteriaCommunicationSecurityLevelPropFromValue returns a pointer to a valid EnumconnectionCriteriaCommunicationSecurityLevelProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumconnectionCriteriaCommunicationSecurityLevelPropFromValue(v string) (*EnumconnectionCriteriaCommunicationSecurityLevelProp, error) {
	ev := EnumconnectionCriteriaCommunicationSecurityLevelProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumconnectionCriteriaCommunicationSecurityLevelProp: valid values are %v", v, AllowedEnumconnectionCriteriaCommunicationSecurityLevelPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumconnectionCriteriaCommunicationSecurityLevelProp) IsValid() bool {
	for _, existing := range AllowedEnumconnectionCriteriaCommunicationSecurityLevelPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumconnection-criteria-communicationSecurityLevelProp value
func (v EnumconnectionCriteriaCommunicationSecurityLevelProp) Ptr() *EnumconnectionCriteriaCommunicationSecurityLevelProp {
	return &v
}

type NullableEnumconnectionCriteriaCommunicationSecurityLevelProp struct {
	value *EnumconnectionCriteriaCommunicationSecurityLevelProp
	isSet bool
}

func (v NullableEnumconnectionCriteriaCommunicationSecurityLevelProp) Get() *EnumconnectionCriteriaCommunicationSecurityLevelProp {
	return v.value
}

func (v *NullableEnumconnectionCriteriaCommunicationSecurityLevelProp) Set(val *EnumconnectionCriteriaCommunicationSecurityLevelProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumconnectionCriteriaCommunicationSecurityLevelProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumconnectionCriteriaCommunicationSecurityLevelProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumconnectionCriteriaCommunicationSecurityLevelProp(val *EnumconnectionCriteriaCommunicationSecurityLevelProp) *NullableEnumconnectionCriteriaCommunicationSecurityLevelProp {
	return &NullableEnumconnectionCriteriaCommunicationSecurityLevelProp{value: val, isSet: true}
}

func (v NullableEnumconnectionCriteriaCommunicationSecurityLevelProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumconnectionCriteriaCommunicationSecurityLevelProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

