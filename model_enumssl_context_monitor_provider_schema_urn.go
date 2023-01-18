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

// EnumsslContextMonitorProviderSchemaUrn the model 'EnumsslContextMonitorProviderSchemaUrn'
type EnumsslContextMonitorProviderSchemaUrn string

// List of Enumssl-context-monitor-providerSchemaUrn
const (
	ENUMSSLCONTEXTMONITORPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0MONITOR_PROVIDERSSL_CONTEXT EnumsslContextMonitorProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:monitor-provider:ssl-context"
)

// All allowed values of EnumsslContextMonitorProviderSchemaUrn enum
var AllowedEnumsslContextMonitorProviderSchemaUrnEnumValues = []EnumsslContextMonitorProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:monitor-provider:ssl-context",
}

func (v *EnumsslContextMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsslContextMonitorProviderSchemaUrn(value)
	for _, existing := range AllowedEnumsslContextMonitorProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsslContextMonitorProviderSchemaUrn", value)
}

// NewEnumsslContextMonitorProviderSchemaUrnFromValue returns a pointer to a valid EnumsslContextMonitorProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsslContextMonitorProviderSchemaUrnFromValue(v string) (*EnumsslContextMonitorProviderSchemaUrn, error) {
	ev := EnumsslContextMonitorProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsslContextMonitorProviderSchemaUrn: valid values are %v", v, AllowedEnumsslContextMonitorProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsslContextMonitorProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsslContextMonitorProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumssl-context-monitor-providerSchemaUrn value
func (v EnumsslContextMonitorProviderSchemaUrn) Ptr() *EnumsslContextMonitorProviderSchemaUrn {
	return &v
}

type NullableEnumsslContextMonitorProviderSchemaUrn struct {
	value *EnumsslContextMonitorProviderSchemaUrn
	isSet bool
}

func (v NullableEnumsslContextMonitorProviderSchemaUrn) Get() *EnumsslContextMonitorProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumsslContextMonitorProviderSchemaUrn) Set(val *EnumsslContextMonitorProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsslContextMonitorProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsslContextMonitorProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsslContextMonitorProviderSchemaUrn(val *EnumsslContextMonitorProviderSchemaUrn) *NullableEnumsslContextMonitorProviderSchemaUrn {
	return &NullableEnumsslContextMonitorProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsslContextMonitorProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsslContextMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
