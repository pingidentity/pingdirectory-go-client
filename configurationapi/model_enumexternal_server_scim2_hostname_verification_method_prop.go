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

// EnumexternalServerScim2HostnameVerificationMethodProp The method that should be used to validate the hostname in the server certificate presented during TLS negotiation.
type EnumexternalServerScim2HostnameVerificationMethodProp string

// List of Enumexternal-server-scim2-hostnameVerificationMethodProp
const (
	ENUMEXTERNALSERVERSCIM2HOSTNAMEVERIFICATIONMETHODPROP_STRICT    EnumexternalServerScim2HostnameVerificationMethodProp = "strict"
	ENUMEXTERNALSERVERSCIM2HOSTNAMEVERIFICATIONMETHODPROP_ALLOW_ALL EnumexternalServerScim2HostnameVerificationMethodProp = "allow-all"
)

// All allowed values of EnumexternalServerScim2HostnameVerificationMethodProp enum
var AllowedEnumexternalServerScim2HostnameVerificationMethodPropEnumValues = []EnumexternalServerScim2HostnameVerificationMethodProp{
	"strict",
	"allow-all",
}

func (v *EnumexternalServerScim2HostnameVerificationMethodProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerScim2HostnameVerificationMethodProp(value)
	for _, existing := range AllowedEnumexternalServerScim2HostnameVerificationMethodPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerScim2HostnameVerificationMethodProp", value)
}

// NewEnumexternalServerScim2HostnameVerificationMethodPropFromValue returns a pointer to a valid EnumexternalServerScim2HostnameVerificationMethodProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerScim2HostnameVerificationMethodPropFromValue(v string) (*EnumexternalServerScim2HostnameVerificationMethodProp, error) {
	ev := EnumexternalServerScim2HostnameVerificationMethodProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerScim2HostnameVerificationMethodProp: valid values are %v", v, AllowedEnumexternalServerScim2HostnameVerificationMethodPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerScim2HostnameVerificationMethodProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerScim2HostnameVerificationMethodPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-scim2-hostnameVerificationMethodProp value
func (v EnumexternalServerScim2HostnameVerificationMethodProp) Ptr() *EnumexternalServerScim2HostnameVerificationMethodProp {
	return &v
}

type NullableEnumexternalServerScim2HostnameVerificationMethodProp struct {
	value *EnumexternalServerScim2HostnameVerificationMethodProp
	isSet bool
}

func (v NullableEnumexternalServerScim2HostnameVerificationMethodProp) Get() *EnumexternalServerScim2HostnameVerificationMethodProp {
	return v.value
}

func (v *NullableEnumexternalServerScim2HostnameVerificationMethodProp) Set(val *EnumexternalServerScim2HostnameVerificationMethodProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerScim2HostnameVerificationMethodProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerScim2HostnameVerificationMethodProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerScim2HostnameVerificationMethodProp(val *EnumexternalServerScim2HostnameVerificationMethodProp) *NullableEnumexternalServerScim2HostnameVerificationMethodProp {
	return &NullableEnumexternalServerScim2HostnameVerificationMethodProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerScim2HostnameVerificationMethodProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerScim2HostnameVerificationMethodProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
