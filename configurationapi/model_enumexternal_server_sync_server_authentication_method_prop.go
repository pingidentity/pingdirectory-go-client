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

// EnumexternalServerSyncServerAuthenticationMethodProp The mechanism to use to authenticate to the target server.
type EnumexternalServerSyncServerAuthenticationMethodProp string

// List of Enumexternal-server-sync-server-authenticationMethodProp
const (
	ENUMEXTERNALSERVERSYNCSERVERAUTHENTICATIONMETHODPROP_NONE         EnumexternalServerSyncServerAuthenticationMethodProp = "none"
	ENUMEXTERNALSERVERSYNCSERVERAUTHENTICATIONMETHODPROP_SIMPLE       EnumexternalServerSyncServerAuthenticationMethodProp = "simple"
	ENUMEXTERNALSERVERSYNCSERVERAUTHENTICATIONMETHODPROP_EXTERNAL     EnumexternalServerSyncServerAuthenticationMethodProp = "external"
	ENUMEXTERNALSERVERSYNCSERVERAUTHENTICATIONMETHODPROP_INTER_SERVER EnumexternalServerSyncServerAuthenticationMethodProp = "inter-server"
)

// All allowed values of EnumexternalServerSyncServerAuthenticationMethodProp enum
var AllowedEnumexternalServerSyncServerAuthenticationMethodPropEnumValues = []EnumexternalServerSyncServerAuthenticationMethodProp{
	"none",
	"simple",
	"external",
	"inter-server",
}

func (v *EnumexternalServerSyncServerAuthenticationMethodProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerSyncServerAuthenticationMethodProp(value)
	for _, existing := range AllowedEnumexternalServerSyncServerAuthenticationMethodPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerSyncServerAuthenticationMethodProp", value)
}

// NewEnumexternalServerSyncServerAuthenticationMethodPropFromValue returns a pointer to a valid EnumexternalServerSyncServerAuthenticationMethodProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerSyncServerAuthenticationMethodPropFromValue(v string) (*EnumexternalServerSyncServerAuthenticationMethodProp, error) {
	ev := EnumexternalServerSyncServerAuthenticationMethodProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerSyncServerAuthenticationMethodProp: valid values are %v", v, AllowedEnumexternalServerSyncServerAuthenticationMethodPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerSyncServerAuthenticationMethodProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerSyncServerAuthenticationMethodPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-sync-server-authenticationMethodProp value
func (v EnumexternalServerSyncServerAuthenticationMethodProp) Ptr() *EnumexternalServerSyncServerAuthenticationMethodProp {
	return &v
}

type NullableEnumexternalServerSyncServerAuthenticationMethodProp struct {
	value *EnumexternalServerSyncServerAuthenticationMethodProp
	isSet bool
}

func (v NullableEnumexternalServerSyncServerAuthenticationMethodProp) Get() *EnumexternalServerSyncServerAuthenticationMethodProp {
	return v.value
}

func (v *NullableEnumexternalServerSyncServerAuthenticationMethodProp) Set(val *EnumexternalServerSyncServerAuthenticationMethodProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerSyncServerAuthenticationMethodProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerSyncServerAuthenticationMethodProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerSyncServerAuthenticationMethodProp(val *EnumexternalServerSyncServerAuthenticationMethodProp) *NullableEnumexternalServerSyncServerAuthenticationMethodProp {
	return &NullableEnumexternalServerSyncServerAuthenticationMethodProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerSyncServerAuthenticationMethodProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerSyncServerAuthenticationMethodProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
