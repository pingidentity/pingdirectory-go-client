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

// EnumpluginAgentxTransportProp Indicates whether the SNMP agent should use TCP or UDP as the transport layer for sub-agent sessions.
type EnumpluginAgentxTransportProp string

// List of Enumplugin-agentxTransportProp
const (
	ENUMPLUGINAGENTXTRANSPORTPROP_TCP EnumpluginAgentxTransportProp = "tcp"
	ENUMPLUGINAGENTXTRANSPORTPROP_UDP EnumpluginAgentxTransportProp = "udp"
)

// All allowed values of EnumpluginAgentxTransportProp enum
var AllowedEnumpluginAgentxTransportPropEnumValues = []EnumpluginAgentxTransportProp{
	"tcp",
	"udp",
}

func (v *EnumpluginAgentxTransportProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginAgentxTransportProp(value)
	for _, existing := range AllowedEnumpluginAgentxTransportPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginAgentxTransportProp", value)
}

// NewEnumpluginAgentxTransportPropFromValue returns a pointer to a valid EnumpluginAgentxTransportProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginAgentxTransportPropFromValue(v string) (*EnumpluginAgentxTransportProp, error) {
	ev := EnumpluginAgentxTransportProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginAgentxTransportProp: valid values are %v", v, AllowedEnumpluginAgentxTransportPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginAgentxTransportProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginAgentxTransportPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-agentxTransportProp value
func (v EnumpluginAgentxTransportProp) Ptr() *EnumpluginAgentxTransportProp {
	return &v
}

type NullableEnumpluginAgentxTransportProp struct {
	value *EnumpluginAgentxTransportProp
	isSet bool
}

func (v NullableEnumpluginAgentxTransportProp) Get() *EnumpluginAgentxTransportProp {
	return v.value
}

func (v *NullableEnumpluginAgentxTransportProp) Set(val *EnumpluginAgentxTransportProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginAgentxTransportProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginAgentxTransportProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginAgentxTransportProp(val *EnumpluginAgentxTransportProp) *NullableEnumpluginAgentxTransportProp {
	return &NullableEnumpluginAgentxTransportProp{value: val, isSet: true}
}

func (v NullableEnumpluginAgentxTransportProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginAgentxTransportProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
