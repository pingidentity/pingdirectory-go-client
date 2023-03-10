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

// EnumexternalServerPingIdentityDsAuthenticationMethodProp The mechanism to use to authenticate to the target server.
type EnumexternalServerPingIdentityDsAuthenticationMethodProp string

// List of Enumexternal-server-ping-identity-ds-authenticationMethodProp
const (
	ENUMEXTERNALSERVERPINGIDENTITYDSAUTHENTICATIONMETHODPROP_NONE         EnumexternalServerPingIdentityDsAuthenticationMethodProp = "none"
	ENUMEXTERNALSERVERPINGIDENTITYDSAUTHENTICATIONMETHODPROP_SIMPLE       EnumexternalServerPingIdentityDsAuthenticationMethodProp = "simple"
	ENUMEXTERNALSERVERPINGIDENTITYDSAUTHENTICATIONMETHODPROP_EXTERNAL     EnumexternalServerPingIdentityDsAuthenticationMethodProp = "external"
	ENUMEXTERNALSERVERPINGIDENTITYDSAUTHENTICATIONMETHODPROP_INTER_SERVER EnumexternalServerPingIdentityDsAuthenticationMethodProp = "inter-server"
)

// All allowed values of EnumexternalServerPingIdentityDsAuthenticationMethodProp enum
var AllowedEnumexternalServerPingIdentityDsAuthenticationMethodPropEnumValues = []EnumexternalServerPingIdentityDsAuthenticationMethodProp{
	"none",
	"simple",
	"external",
	"inter-server",
}

func (v *EnumexternalServerPingIdentityDsAuthenticationMethodProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerPingIdentityDsAuthenticationMethodProp(value)
	for _, existing := range AllowedEnumexternalServerPingIdentityDsAuthenticationMethodPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerPingIdentityDsAuthenticationMethodProp", value)
}

// NewEnumexternalServerPingIdentityDsAuthenticationMethodPropFromValue returns a pointer to a valid EnumexternalServerPingIdentityDsAuthenticationMethodProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerPingIdentityDsAuthenticationMethodPropFromValue(v string) (*EnumexternalServerPingIdentityDsAuthenticationMethodProp, error) {
	ev := EnumexternalServerPingIdentityDsAuthenticationMethodProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerPingIdentityDsAuthenticationMethodProp: valid values are %v", v, AllowedEnumexternalServerPingIdentityDsAuthenticationMethodPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerPingIdentityDsAuthenticationMethodProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerPingIdentityDsAuthenticationMethodPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-ping-identity-ds-authenticationMethodProp value
func (v EnumexternalServerPingIdentityDsAuthenticationMethodProp) Ptr() *EnumexternalServerPingIdentityDsAuthenticationMethodProp {
	return &v
}

type NullableEnumexternalServerPingIdentityDsAuthenticationMethodProp struct {
	value *EnumexternalServerPingIdentityDsAuthenticationMethodProp
	isSet bool
}

func (v NullableEnumexternalServerPingIdentityDsAuthenticationMethodProp) Get() *EnumexternalServerPingIdentityDsAuthenticationMethodProp {
	return v.value
}

func (v *NullableEnumexternalServerPingIdentityDsAuthenticationMethodProp) Set(val *EnumexternalServerPingIdentityDsAuthenticationMethodProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerPingIdentityDsAuthenticationMethodProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerPingIdentityDsAuthenticationMethodProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerPingIdentityDsAuthenticationMethodProp(val *EnumexternalServerPingIdentityDsAuthenticationMethodProp) *NullableEnumexternalServerPingIdentityDsAuthenticationMethodProp {
	return &NullableEnumexternalServerPingIdentityDsAuthenticationMethodProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerPingIdentityDsAuthenticationMethodProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerPingIdentityDsAuthenticationMethodProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
