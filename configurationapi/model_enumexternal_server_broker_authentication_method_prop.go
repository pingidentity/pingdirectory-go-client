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

// EnumexternalServerBrokerAuthenticationMethodProp The mechanism to use to authenticate to the target server.
type EnumexternalServerBrokerAuthenticationMethodProp string

// List of Enumexternal-server-broker-authenticationMethodProp
const (
	ENUMEXTERNALSERVERBROKERAUTHENTICATIONMETHODPROP_NONE         EnumexternalServerBrokerAuthenticationMethodProp = "none"
	ENUMEXTERNALSERVERBROKERAUTHENTICATIONMETHODPROP_SIMPLE       EnumexternalServerBrokerAuthenticationMethodProp = "simple"
	ENUMEXTERNALSERVERBROKERAUTHENTICATIONMETHODPROP_EXTERNAL     EnumexternalServerBrokerAuthenticationMethodProp = "external"
	ENUMEXTERNALSERVERBROKERAUTHENTICATIONMETHODPROP_INTER_SERVER EnumexternalServerBrokerAuthenticationMethodProp = "inter-server"
)

// All allowed values of EnumexternalServerBrokerAuthenticationMethodProp enum
var AllowedEnumexternalServerBrokerAuthenticationMethodPropEnumValues = []EnumexternalServerBrokerAuthenticationMethodProp{
	"none",
	"simple",
	"external",
	"inter-server",
}

func (v *EnumexternalServerBrokerAuthenticationMethodProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerBrokerAuthenticationMethodProp(value)
	for _, existing := range AllowedEnumexternalServerBrokerAuthenticationMethodPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerBrokerAuthenticationMethodProp", value)
}

// NewEnumexternalServerBrokerAuthenticationMethodPropFromValue returns a pointer to a valid EnumexternalServerBrokerAuthenticationMethodProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerBrokerAuthenticationMethodPropFromValue(v string) (*EnumexternalServerBrokerAuthenticationMethodProp, error) {
	ev := EnumexternalServerBrokerAuthenticationMethodProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerBrokerAuthenticationMethodProp: valid values are %v", v, AllowedEnumexternalServerBrokerAuthenticationMethodPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerBrokerAuthenticationMethodProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerBrokerAuthenticationMethodPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-broker-authenticationMethodProp value
func (v EnumexternalServerBrokerAuthenticationMethodProp) Ptr() *EnumexternalServerBrokerAuthenticationMethodProp {
	return &v
}

type NullableEnumexternalServerBrokerAuthenticationMethodProp struct {
	value *EnumexternalServerBrokerAuthenticationMethodProp
	isSet bool
}

func (v NullableEnumexternalServerBrokerAuthenticationMethodProp) Get() *EnumexternalServerBrokerAuthenticationMethodProp {
	return v.value
}

func (v *NullableEnumexternalServerBrokerAuthenticationMethodProp) Set(val *EnumexternalServerBrokerAuthenticationMethodProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerBrokerAuthenticationMethodProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerBrokerAuthenticationMethodProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerBrokerAuthenticationMethodProp(val *EnumexternalServerBrokerAuthenticationMethodProp) *NullableEnumexternalServerBrokerAuthenticationMethodProp {
	return &NullableEnumexternalServerBrokerAuthenticationMethodProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerBrokerAuthenticationMethodProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerBrokerAuthenticationMethodProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
