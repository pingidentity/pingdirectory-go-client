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

// EnumexternalServerPolicyHostnameVerificationMethodProp The mechanism for checking if the hostname of the Policy External Server matches the name(s) stored inside the server's X.509 certificate. This is only applicable if SSL is being used for connection security.
type EnumexternalServerPolicyHostnameVerificationMethodProp string

// List of Enumexternal-server-policy-hostnameVerificationMethodProp
const (
	ENUMEXTERNALSERVERPOLICYHOSTNAMEVERIFICATIONMETHODPROP_ALLOW_ALL EnumexternalServerPolicyHostnameVerificationMethodProp = "allow-all"
	ENUMEXTERNALSERVERPOLICYHOSTNAMEVERIFICATIONMETHODPROP_STRICT    EnumexternalServerPolicyHostnameVerificationMethodProp = "strict"
)

// All allowed values of EnumexternalServerPolicyHostnameVerificationMethodProp enum
var AllowedEnumexternalServerPolicyHostnameVerificationMethodPropEnumValues = []EnumexternalServerPolicyHostnameVerificationMethodProp{
	"allow-all",
	"strict",
}

func (v *EnumexternalServerPolicyHostnameVerificationMethodProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerPolicyHostnameVerificationMethodProp(value)
	for _, existing := range AllowedEnumexternalServerPolicyHostnameVerificationMethodPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerPolicyHostnameVerificationMethodProp", value)
}

// NewEnumexternalServerPolicyHostnameVerificationMethodPropFromValue returns a pointer to a valid EnumexternalServerPolicyHostnameVerificationMethodProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerPolicyHostnameVerificationMethodPropFromValue(v string) (*EnumexternalServerPolicyHostnameVerificationMethodProp, error) {
	ev := EnumexternalServerPolicyHostnameVerificationMethodProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerPolicyHostnameVerificationMethodProp: valid values are %v", v, AllowedEnumexternalServerPolicyHostnameVerificationMethodPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerPolicyHostnameVerificationMethodProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerPolicyHostnameVerificationMethodPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-policy-hostnameVerificationMethodProp value
func (v EnumexternalServerPolicyHostnameVerificationMethodProp) Ptr() *EnumexternalServerPolicyHostnameVerificationMethodProp {
	return &v
}

type NullableEnumexternalServerPolicyHostnameVerificationMethodProp struct {
	value *EnumexternalServerPolicyHostnameVerificationMethodProp
	isSet bool
}

func (v NullableEnumexternalServerPolicyHostnameVerificationMethodProp) Get() *EnumexternalServerPolicyHostnameVerificationMethodProp {
	return v.value
}

func (v *NullableEnumexternalServerPolicyHostnameVerificationMethodProp) Set(val *EnumexternalServerPolicyHostnameVerificationMethodProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerPolicyHostnameVerificationMethodProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerPolicyHostnameVerificationMethodProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerPolicyHostnameVerificationMethodProp(val *EnumexternalServerPolicyHostnameVerificationMethodProp) *NullableEnumexternalServerPolicyHostnameVerificationMethodProp {
	return &NullableEnumexternalServerPolicyHostnameVerificationMethodProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerPolicyHostnameVerificationMethodProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerPolicyHostnameVerificationMethodProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
