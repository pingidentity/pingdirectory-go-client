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

// EnumexternalServerPingOneHttpHostnameVerificationMethodProp The mechanism for checking if the hostname in the PingOne ID Token Validator's base-url value matches the name(s) stored inside the X.509 certificate presented by PingOne.
type EnumexternalServerPingOneHttpHostnameVerificationMethodProp string

// List of Enumexternal-server-ping-one-http-hostnameVerificationMethodProp
const (
	ENUMEXTERNALSERVERPINGONEHTTPHOSTNAMEVERIFICATIONMETHODPROP_ALLOW_ALL EnumexternalServerPingOneHttpHostnameVerificationMethodProp = "allow-all"
	ENUMEXTERNALSERVERPINGONEHTTPHOSTNAMEVERIFICATIONMETHODPROP_STRICT    EnumexternalServerPingOneHttpHostnameVerificationMethodProp = "strict"
)

// All allowed values of EnumexternalServerPingOneHttpHostnameVerificationMethodProp enum
var AllowedEnumexternalServerPingOneHttpHostnameVerificationMethodPropEnumValues = []EnumexternalServerPingOneHttpHostnameVerificationMethodProp{
	"allow-all",
	"strict",
}

func (v *EnumexternalServerPingOneHttpHostnameVerificationMethodProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerPingOneHttpHostnameVerificationMethodProp(value)
	for _, existing := range AllowedEnumexternalServerPingOneHttpHostnameVerificationMethodPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerPingOneHttpHostnameVerificationMethodProp", value)
}

// NewEnumexternalServerPingOneHttpHostnameVerificationMethodPropFromValue returns a pointer to a valid EnumexternalServerPingOneHttpHostnameVerificationMethodProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerPingOneHttpHostnameVerificationMethodPropFromValue(v string) (*EnumexternalServerPingOneHttpHostnameVerificationMethodProp, error) {
	ev := EnumexternalServerPingOneHttpHostnameVerificationMethodProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerPingOneHttpHostnameVerificationMethodProp: valid values are %v", v, AllowedEnumexternalServerPingOneHttpHostnameVerificationMethodPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerPingOneHttpHostnameVerificationMethodProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerPingOneHttpHostnameVerificationMethodPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-ping-one-http-hostnameVerificationMethodProp value
func (v EnumexternalServerPingOneHttpHostnameVerificationMethodProp) Ptr() *EnumexternalServerPingOneHttpHostnameVerificationMethodProp {
	return &v
}

type NullableEnumexternalServerPingOneHttpHostnameVerificationMethodProp struct {
	value *EnumexternalServerPingOneHttpHostnameVerificationMethodProp
	isSet bool
}

func (v NullableEnumexternalServerPingOneHttpHostnameVerificationMethodProp) Get() *EnumexternalServerPingOneHttpHostnameVerificationMethodProp {
	return v.value
}

func (v *NullableEnumexternalServerPingOneHttpHostnameVerificationMethodProp) Set(val *EnumexternalServerPingOneHttpHostnameVerificationMethodProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerPingOneHttpHostnameVerificationMethodProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerPingOneHttpHostnameVerificationMethodProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerPingOneHttpHostnameVerificationMethodProp(val *EnumexternalServerPingOneHttpHostnameVerificationMethodProp) *NullableEnumexternalServerPingOneHttpHostnameVerificationMethodProp {
	return &NullableEnumexternalServerPingOneHttpHostnameVerificationMethodProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerPingOneHttpHostnameVerificationMethodProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerPingOneHttpHostnameVerificationMethodProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
