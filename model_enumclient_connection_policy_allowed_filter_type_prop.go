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

// EnumclientConnectionPolicyAllowedFilterTypeProp Specifies the types of filter components that may be included in search requests from clients associated with this Client Connection Policy which have a non-baseObject scope.
type EnumclientConnectionPolicyAllowedFilterTypeProp string

// List of Enumclient-connection-policy-allowedFilterTypeProp
const (
	AND EnumclientConnectionPolicyAllowedFilterTypeProp = "and"
	OR EnumclientConnectionPolicyAllowedFilterTypeProp = "or"
	NOT EnumclientConnectionPolicyAllowedFilterTypeProp = "not"
	EQUALITY EnumclientConnectionPolicyAllowedFilterTypeProp = "equality"
	SUB_INITIAL EnumclientConnectionPolicyAllowedFilterTypeProp = "sub-initial"
	SUB_ANY EnumclientConnectionPolicyAllowedFilterTypeProp = "sub-any"
	SUB_FINAL EnumclientConnectionPolicyAllowedFilterTypeProp = "sub-final"
	GREATER_OR_EQUAL EnumclientConnectionPolicyAllowedFilterTypeProp = "greater-or-equal"
	LESS_OR_EQUAL EnumclientConnectionPolicyAllowedFilterTypeProp = "less-or-equal"
	PRESENT EnumclientConnectionPolicyAllowedFilterTypeProp = "present"
	APPROXIMATE_MATCH EnumclientConnectionPolicyAllowedFilterTypeProp = "approximate-match"
	EXTENSIBLE_MATCH EnumclientConnectionPolicyAllowedFilterTypeProp = "extensible-match"
)

// All allowed values of EnumclientConnectionPolicyAllowedFilterTypeProp enum
var AllowedEnumclientConnectionPolicyAllowedFilterTypePropEnumValues = []EnumclientConnectionPolicyAllowedFilterTypeProp{
	"and",
	"or",
	"not",
	"equality",
	"sub-initial",
	"sub-any",
	"sub-final",
	"greater-or-equal",
	"less-or-equal",
	"present",
	"approximate-match",
	"extensible-match",
}

func (v *EnumclientConnectionPolicyAllowedFilterTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumclientConnectionPolicyAllowedFilterTypeProp(value)
	for _, existing := range AllowedEnumclientConnectionPolicyAllowedFilterTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumclientConnectionPolicyAllowedFilterTypeProp", value)
}

// NewEnumclientConnectionPolicyAllowedFilterTypePropFromValue returns a pointer to a valid EnumclientConnectionPolicyAllowedFilterTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumclientConnectionPolicyAllowedFilterTypePropFromValue(v string) (*EnumclientConnectionPolicyAllowedFilterTypeProp, error) {
	ev := EnumclientConnectionPolicyAllowedFilterTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumclientConnectionPolicyAllowedFilterTypeProp: valid values are %v", v, AllowedEnumclientConnectionPolicyAllowedFilterTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumclientConnectionPolicyAllowedFilterTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumclientConnectionPolicyAllowedFilterTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumclient-connection-policy-allowedFilterTypeProp value
func (v EnumclientConnectionPolicyAllowedFilterTypeProp) Ptr() *EnumclientConnectionPolicyAllowedFilterTypeProp {
	return &v
}

type NullableEnumclientConnectionPolicyAllowedFilterTypeProp struct {
	value *EnumclientConnectionPolicyAllowedFilterTypeProp
	isSet bool
}

func (v NullableEnumclientConnectionPolicyAllowedFilterTypeProp) Get() *EnumclientConnectionPolicyAllowedFilterTypeProp {
	return v.value
}

func (v *NullableEnumclientConnectionPolicyAllowedFilterTypeProp) Set(val *EnumclientConnectionPolicyAllowedFilterTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumclientConnectionPolicyAllowedFilterTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumclientConnectionPolicyAllowedFilterTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumclientConnectionPolicyAllowedFilterTypeProp(val *EnumclientConnectionPolicyAllowedFilterTypeProp) *NullableEnumclientConnectionPolicyAllowedFilterTypeProp {
	return &NullableEnumclientConnectionPolicyAllowedFilterTypeProp{value: val, isSet: true}
}

func (v NullableEnumclientConnectionPolicyAllowedFilterTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumclientConnectionPolicyAllowedFilterTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

