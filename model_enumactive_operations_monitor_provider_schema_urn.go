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

// EnumactiveOperationsMonitorProviderSchemaUrn the model 'EnumactiveOperationsMonitorProviderSchemaUrn'
type EnumactiveOperationsMonitorProviderSchemaUrn string

// List of Enumactive-operations-monitor-providerSchemaUrn
const (
	ENUMACTIVEOPERATIONSMONITORPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0MONITOR_PROVIDERACTIVE_OPERATIONS EnumactiveOperationsMonitorProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:monitor-provider:active-operations"
)

// All allowed values of EnumactiveOperationsMonitorProviderSchemaUrn enum
var AllowedEnumactiveOperationsMonitorProviderSchemaUrnEnumValues = []EnumactiveOperationsMonitorProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:monitor-provider:active-operations",
}

func (v *EnumactiveOperationsMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumactiveOperationsMonitorProviderSchemaUrn(value)
	for _, existing := range AllowedEnumactiveOperationsMonitorProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumactiveOperationsMonitorProviderSchemaUrn", value)
}

// NewEnumactiveOperationsMonitorProviderSchemaUrnFromValue returns a pointer to a valid EnumactiveOperationsMonitorProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumactiveOperationsMonitorProviderSchemaUrnFromValue(v string) (*EnumactiveOperationsMonitorProviderSchemaUrn, error) {
	ev := EnumactiveOperationsMonitorProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumactiveOperationsMonitorProviderSchemaUrn: valid values are %v", v, AllowedEnumactiveOperationsMonitorProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumactiveOperationsMonitorProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumactiveOperationsMonitorProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumactive-operations-monitor-providerSchemaUrn value
func (v EnumactiveOperationsMonitorProviderSchemaUrn) Ptr() *EnumactiveOperationsMonitorProviderSchemaUrn {
	return &v
}

type NullableEnumactiveOperationsMonitorProviderSchemaUrn struct {
	value *EnumactiveOperationsMonitorProviderSchemaUrn
	isSet bool
}

func (v NullableEnumactiveOperationsMonitorProviderSchemaUrn) Get() *EnumactiveOperationsMonitorProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumactiveOperationsMonitorProviderSchemaUrn) Set(val *EnumactiveOperationsMonitorProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumactiveOperationsMonitorProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumactiveOperationsMonitorProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumactiveOperationsMonitorProviderSchemaUrn(val *EnumactiveOperationsMonitorProviderSchemaUrn) *NullableEnumactiveOperationsMonitorProviderSchemaUrn {
	return &NullableEnumactiveOperationsMonitorProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumactiveOperationsMonitorProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumactiveOperationsMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

