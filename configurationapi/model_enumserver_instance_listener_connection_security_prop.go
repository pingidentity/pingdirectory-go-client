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

// EnumserverInstanceListenerConnectionSecurityProp Specifies the mechanism to use for securing connections to the server.
type EnumserverInstanceListenerConnectionSecurityProp string

// List of Enumserver-instance-listener-connectionSecurityProp
const (
	ENUMSERVERINSTANCELISTENERCONNECTIONSECURITYPROP_NONE     EnumserverInstanceListenerConnectionSecurityProp = "none"
	ENUMSERVERINSTANCELISTENERCONNECTIONSECURITYPROP_SSL      EnumserverInstanceListenerConnectionSecurityProp = "ssl"
	ENUMSERVERINSTANCELISTENERCONNECTIONSECURITYPROP_STARTTLS EnumserverInstanceListenerConnectionSecurityProp = "starttls"
)

// All allowed values of EnumserverInstanceListenerConnectionSecurityProp enum
var AllowedEnumserverInstanceListenerConnectionSecurityPropEnumValues = []EnumserverInstanceListenerConnectionSecurityProp{
	"none",
	"ssl",
	"starttls",
}

func (v *EnumserverInstanceListenerConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumserverInstanceListenerConnectionSecurityProp(value)
	for _, existing := range AllowedEnumserverInstanceListenerConnectionSecurityPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumserverInstanceListenerConnectionSecurityProp", value)
}

// NewEnumserverInstanceListenerConnectionSecurityPropFromValue returns a pointer to a valid EnumserverInstanceListenerConnectionSecurityProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumserverInstanceListenerConnectionSecurityPropFromValue(v string) (*EnumserverInstanceListenerConnectionSecurityProp, error) {
	ev := EnumserverInstanceListenerConnectionSecurityProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumserverInstanceListenerConnectionSecurityProp: valid values are %v", v, AllowedEnumserverInstanceListenerConnectionSecurityPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumserverInstanceListenerConnectionSecurityProp) IsValid() bool {
	for _, existing := range AllowedEnumserverInstanceListenerConnectionSecurityPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumserver-instance-listener-connectionSecurityProp value
func (v EnumserverInstanceListenerConnectionSecurityProp) Ptr() *EnumserverInstanceListenerConnectionSecurityProp {
	return &v
}

type NullableEnumserverInstanceListenerConnectionSecurityProp struct {
	value *EnumserverInstanceListenerConnectionSecurityProp
	isSet bool
}

func (v NullableEnumserverInstanceListenerConnectionSecurityProp) Get() *EnumserverInstanceListenerConnectionSecurityProp {
	return v.value
}

func (v *NullableEnumserverInstanceListenerConnectionSecurityProp) Set(val *EnumserverInstanceListenerConnectionSecurityProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumserverInstanceListenerConnectionSecurityProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumserverInstanceListenerConnectionSecurityProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumserverInstanceListenerConnectionSecurityProp(val *EnumserverInstanceListenerConnectionSecurityProp) *NullableEnumserverInstanceListenerConnectionSecurityProp {
	return &NullableEnumserverInstanceListenerConnectionSecurityProp{value: val, isSet: true}
}

func (v NullableEnumserverInstanceListenerConnectionSecurityProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumserverInstanceListenerConnectionSecurityProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}