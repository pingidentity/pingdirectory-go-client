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

// EnumperiodicGcPluginSchemaUrn the model 'EnumperiodicGcPluginSchemaUrn'
type EnumperiodicGcPluginSchemaUrn string

// List of Enumperiodic-gc-pluginSchemaUrn
const (
	ENUMPERIODICGCPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGINPERIODIC_GC EnumperiodicGcPluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:periodic-gc"
)

// All allowed values of EnumperiodicGcPluginSchemaUrn enum
var AllowedEnumperiodicGcPluginSchemaUrnEnumValues = []EnumperiodicGcPluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:periodic-gc",
}

func (v *EnumperiodicGcPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumperiodicGcPluginSchemaUrn(value)
	for _, existing := range AllowedEnumperiodicGcPluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumperiodicGcPluginSchemaUrn", value)
}

// NewEnumperiodicGcPluginSchemaUrnFromValue returns a pointer to a valid EnumperiodicGcPluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumperiodicGcPluginSchemaUrnFromValue(v string) (*EnumperiodicGcPluginSchemaUrn, error) {
	ev := EnumperiodicGcPluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumperiodicGcPluginSchemaUrn: valid values are %v", v, AllowedEnumperiodicGcPluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumperiodicGcPluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumperiodicGcPluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumperiodic-gc-pluginSchemaUrn value
func (v EnumperiodicGcPluginSchemaUrn) Ptr() *EnumperiodicGcPluginSchemaUrn {
	return &v
}

type NullableEnumperiodicGcPluginSchemaUrn struct {
	value *EnumperiodicGcPluginSchemaUrn
	isSet bool
}

func (v NullableEnumperiodicGcPluginSchemaUrn) Get() *EnumperiodicGcPluginSchemaUrn {
	return v.value
}

func (v *NullableEnumperiodicGcPluginSchemaUrn) Set(val *EnumperiodicGcPluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumperiodicGcPluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumperiodicGcPluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumperiodicGcPluginSchemaUrn(val *EnumperiodicGcPluginSchemaUrn) *NullableEnumperiodicGcPluginSchemaUrn {
	return &NullableEnumperiodicGcPluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumperiodicGcPluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumperiodicGcPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
