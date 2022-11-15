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

// EnumconnectionCriteriaAuthenticationSecurityLevelProp Indicates whether this Simple Connection Criteria should require or allow clients that authenticated using a secure manner. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections.
type EnumconnectionCriteriaAuthenticationSecurityLevelProp string

// List of Enumconnection-criteria-authenticationSecurityLevelProp
const (
	ANY EnumconnectionCriteriaAuthenticationSecurityLevelProp = "any"
	SECURE_ONLY EnumconnectionCriteriaAuthenticationSecurityLevelProp = "secure-only"
	INSECURE_ONLY EnumconnectionCriteriaAuthenticationSecurityLevelProp = "insecure-only"
)

// All allowed values of EnumconnectionCriteriaAuthenticationSecurityLevelProp enum
var AllowedEnumconnectionCriteriaAuthenticationSecurityLevelPropEnumValues = []EnumconnectionCriteriaAuthenticationSecurityLevelProp{
	"any",
	"secure-only",
	"insecure-only",
}

func (v *EnumconnectionCriteriaAuthenticationSecurityLevelProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumconnectionCriteriaAuthenticationSecurityLevelProp(value)
	for _, existing := range AllowedEnumconnectionCriteriaAuthenticationSecurityLevelPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumconnectionCriteriaAuthenticationSecurityLevelProp", value)
}

// NewEnumconnectionCriteriaAuthenticationSecurityLevelPropFromValue returns a pointer to a valid EnumconnectionCriteriaAuthenticationSecurityLevelProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumconnectionCriteriaAuthenticationSecurityLevelPropFromValue(v string) (*EnumconnectionCriteriaAuthenticationSecurityLevelProp, error) {
	ev := EnumconnectionCriteriaAuthenticationSecurityLevelProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumconnectionCriteriaAuthenticationSecurityLevelProp: valid values are %v", v, AllowedEnumconnectionCriteriaAuthenticationSecurityLevelPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumconnectionCriteriaAuthenticationSecurityLevelProp) IsValid() bool {
	for _, existing := range AllowedEnumconnectionCriteriaAuthenticationSecurityLevelPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumconnection-criteria-authenticationSecurityLevelProp value
func (v EnumconnectionCriteriaAuthenticationSecurityLevelProp) Ptr() *EnumconnectionCriteriaAuthenticationSecurityLevelProp {
	return &v
}

type NullableEnumconnectionCriteriaAuthenticationSecurityLevelProp struct {
	value *EnumconnectionCriteriaAuthenticationSecurityLevelProp
	isSet bool
}

func (v NullableEnumconnectionCriteriaAuthenticationSecurityLevelProp) Get() *EnumconnectionCriteriaAuthenticationSecurityLevelProp {
	return v.value
}

func (v *NullableEnumconnectionCriteriaAuthenticationSecurityLevelProp) Set(val *EnumconnectionCriteriaAuthenticationSecurityLevelProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumconnectionCriteriaAuthenticationSecurityLevelProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumconnectionCriteriaAuthenticationSecurityLevelProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumconnectionCriteriaAuthenticationSecurityLevelProp(val *EnumconnectionCriteriaAuthenticationSecurityLevelProp) *NullableEnumconnectionCriteriaAuthenticationSecurityLevelProp {
	return &NullableEnumconnectionCriteriaAuthenticationSecurityLevelProp{value: val, isSet: true}
}

func (v NullableEnumconnectionCriteriaAuthenticationSecurityLevelProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumconnectionCriteriaAuthenticationSecurityLevelProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

