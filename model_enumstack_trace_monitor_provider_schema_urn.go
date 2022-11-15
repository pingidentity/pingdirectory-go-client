/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// EnumstackTraceMonitorProviderSchemaUrn the model 'EnumstackTraceMonitorProviderSchemaUrn'
type EnumstackTraceMonitorProviderSchemaUrn string

// List of Enumstack-trace-monitor-providerSchemaUrn
const (
	ENUMSTACKTRACEMONITORPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0MONITOR_PROVIDERSTACK_TRACE EnumstackTraceMonitorProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:monitor-provider:stack-trace"
)

// All allowed values of EnumstackTraceMonitorProviderSchemaUrn enum
var AllowedEnumstackTraceMonitorProviderSchemaUrnEnumValues = []EnumstackTraceMonitorProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:monitor-provider:stack-trace",
}

func (v *EnumstackTraceMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumstackTraceMonitorProviderSchemaUrn(value)
	for _, existing := range AllowedEnumstackTraceMonitorProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumstackTraceMonitorProviderSchemaUrn", value)
}

// NewEnumstackTraceMonitorProviderSchemaUrnFromValue returns a pointer to a valid EnumstackTraceMonitorProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumstackTraceMonitorProviderSchemaUrnFromValue(v string) (*EnumstackTraceMonitorProviderSchemaUrn, error) {
	ev := EnumstackTraceMonitorProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumstackTraceMonitorProviderSchemaUrn: valid values are %v", v, AllowedEnumstackTraceMonitorProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumstackTraceMonitorProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumstackTraceMonitorProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumstack-trace-monitor-providerSchemaUrn value
func (v EnumstackTraceMonitorProviderSchemaUrn) Ptr() *EnumstackTraceMonitorProviderSchemaUrn {
	return &v
}

type NullableEnumstackTraceMonitorProviderSchemaUrn struct {
	value *EnumstackTraceMonitorProviderSchemaUrn
	isSet bool
}

func (v NullableEnumstackTraceMonitorProviderSchemaUrn) Get() *EnumstackTraceMonitorProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumstackTraceMonitorProviderSchemaUrn) Set(val *EnumstackTraceMonitorProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumstackTraceMonitorProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumstackTraceMonitorProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumstackTraceMonitorProviderSchemaUrn(val *EnumstackTraceMonitorProviderSchemaUrn) *NullableEnumstackTraceMonitorProviderSchemaUrn {
	return &NullableEnumstackTraceMonitorProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumstackTraceMonitorProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumstackTraceMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

