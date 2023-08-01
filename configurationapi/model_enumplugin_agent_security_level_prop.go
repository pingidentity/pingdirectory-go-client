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

// EnumpluginAgentSecurityLevelProp The type of security to use for SNMPv3 communication.
type EnumpluginAgentSecurityLevelProp string

// List of Enumplugin-agentSecurityLevelProp
const (
	ENUMPLUGINAGENTSECURITYLEVELPROP_NOAUTH_NOPRIV EnumpluginAgentSecurityLevelProp = "noauth-nopriv"
	ENUMPLUGINAGENTSECURITYLEVELPROP_AUTH_NOPRIV   EnumpluginAgentSecurityLevelProp = "auth-nopriv"
	ENUMPLUGINAGENTSECURITYLEVELPROP_AUTH_PRIV     EnumpluginAgentSecurityLevelProp = "auth-priv"
)

// All allowed values of EnumpluginAgentSecurityLevelProp enum
var AllowedEnumpluginAgentSecurityLevelPropEnumValues = []EnumpluginAgentSecurityLevelProp{
	"noauth-nopriv",
	"auth-nopriv",
	"auth-priv",
}

func (v *EnumpluginAgentSecurityLevelProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginAgentSecurityLevelProp(value)
	for _, existing := range AllowedEnumpluginAgentSecurityLevelPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginAgentSecurityLevelProp", value)
}

// NewEnumpluginAgentSecurityLevelPropFromValue returns a pointer to a valid EnumpluginAgentSecurityLevelProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginAgentSecurityLevelPropFromValue(v string) (*EnumpluginAgentSecurityLevelProp, error) {
	ev := EnumpluginAgentSecurityLevelProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginAgentSecurityLevelProp: valid values are %v", v, AllowedEnumpluginAgentSecurityLevelPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginAgentSecurityLevelProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginAgentSecurityLevelPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-agentSecurityLevelProp value
func (v EnumpluginAgentSecurityLevelProp) Ptr() *EnumpluginAgentSecurityLevelProp {
	return &v
}

type NullableEnumpluginAgentSecurityLevelProp struct {
	value *EnumpluginAgentSecurityLevelProp
	isSet bool
}

func (v NullableEnumpluginAgentSecurityLevelProp) Get() *EnumpluginAgentSecurityLevelProp {
	return v.value
}

func (v *NullableEnumpluginAgentSecurityLevelProp) Set(val *EnumpluginAgentSecurityLevelProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginAgentSecurityLevelProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginAgentSecurityLevelProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginAgentSecurityLevelProp(val *EnumpluginAgentSecurityLevelProp) *NullableEnumpluginAgentSecurityLevelProp {
	return &NullableEnumpluginAgentSecurityLevelProp{value: val, isSet: true}
}

func (v NullableEnumpluginAgentSecurityLevelProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginAgentSecurityLevelProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
