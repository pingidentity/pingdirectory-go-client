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

// EnumexternalServerHostnameVerificationMethodProp The mechanism for checking if the hostname in the PingOne ID Token Validator's base-url value matches the name(s) stored inside the X.509 certificate presented by PingOne.
type EnumexternalServerHostnameVerificationMethodProp string

// List of Enumexternal-server-hostnameVerificationMethodProp
const (
	ENUMEXTERNALSERVERHOSTNAMEVERIFICATIONMETHODPROP_ALLOW_ALL EnumexternalServerHostnameVerificationMethodProp = "allow-all"
	ENUMEXTERNALSERVERHOSTNAMEVERIFICATIONMETHODPROP_STRICT    EnumexternalServerHostnameVerificationMethodProp = "strict"
)

// All allowed values of EnumexternalServerHostnameVerificationMethodProp enum
var AllowedEnumexternalServerHostnameVerificationMethodPropEnumValues = []EnumexternalServerHostnameVerificationMethodProp{
	"allow-all",
	"strict",
}

func (v *EnumexternalServerHostnameVerificationMethodProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerHostnameVerificationMethodProp(value)
	for _, existing := range AllowedEnumexternalServerHostnameVerificationMethodPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerHostnameVerificationMethodProp", value)
}

// NewEnumexternalServerHostnameVerificationMethodPropFromValue returns a pointer to a valid EnumexternalServerHostnameVerificationMethodProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerHostnameVerificationMethodPropFromValue(v string) (*EnumexternalServerHostnameVerificationMethodProp, error) {
	ev := EnumexternalServerHostnameVerificationMethodProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerHostnameVerificationMethodProp: valid values are %v", v, AllowedEnumexternalServerHostnameVerificationMethodPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerHostnameVerificationMethodProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerHostnameVerificationMethodPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-hostnameVerificationMethodProp value
func (v EnumexternalServerHostnameVerificationMethodProp) Ptr() *EnumexternalServerHostnameVerificationMethodProp {
	return &v
}

type NullableEnumexternalServerHostnameVerificationMethodProp struct {
	value *EnumexternalServerHostnameVerificationMethodProp
	isSet bool
}

func (v NullableEnumexternalServerHostnameVerificationMethodProp) Get() *EnumexternalServerHostnameVerificationMethodProp {
	return v.value
}

func (v *NullableEnumexternalServerHostnameVerificationMethodProp) Set(val *EnumexternalServerHostnameVerificationMethodProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerHostnameVerificationMethodProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerHostnameVerificationMethodProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerHostnameVerificationMethodProp(val *EnumexternalServerHostnameVerificationMethodProp) *NullableEnumexternalServerHostnameVerificationMethodProp {
	return &NullableEnumexternalServerHostnameVerificationMethodProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerHostnameVerificationMethodProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerHostnameVerificationMethodProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
