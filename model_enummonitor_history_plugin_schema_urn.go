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

// EnummonitorHistoryPluginSchemaUrn the model 'EnummonitorHistoryPluginSchemaUrn'
type EnummonitorHistoryPluginSchemaUrn string

// List of Enummonitor-history-pluginSchemaUrn
const (
	ENUMMONITORHISTORYPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGINMONITOR_HISTORY EnummonitorHistoryPluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:monitor-history"
)

// All allowed values of EnummonitorHistoryPluginSchemaUrn enum
var AllowedEnummonitorHistoryPluginSchemaUrnEnumValues = []EnummonitorHistoryPluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:monitor-history",
}

func (v *EnummonitorHistoryPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnummonitorHistoryPluginSchemaUrn(value)
	for _, existing := range AllowedEnummonitorHistoryPluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnummonitorHistoryPluginSchemaUrn", value)
}

// NewEnummonitorHistoryPluginSchemaUrnFromValue returns a pointer to a valid EnummonitorHistoryPluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnummonitorHistoryPluginSchemaUrnFromValue(v string) (*EnummonitorHistoryPluginSchemaUrn, error) {
	ev := EnummonitorHistoryPluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnummonitorHistoryPluginSchemaUrn: valid values are %v", v, AllowedEnummonitorHistoryPluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnummonitorHistoryPluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnummonitorHistoryPluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enummonitor-history-pluginSchemaUrn value
func (v EnummonitorHistoryPluginSchemaUrn) Ptr() *EnummonitorHistoryPluginSchemaUrn {
	return &v
}

type NullableEnummonitorHistoryPluginSchemaUrn struct {
	value *EnummonitorHistoryPluginSchemaUrn
	isSet bool
}

func (v NullableEnummonitorHistoryPluginSchemaUrn) Get() *EnummonitorHistoryPluginSchemaUrn {
	return v.value
}

func (v *NullableEnummonitorHistoryPluginSchemaUrn) Set(val *EnummonitorHistoryPluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnummonitorHistoryPluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnummonitorHistoryPluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnummonitorHistoryPluginSchemaUrn(val *EnummonitorHistoryPluginSchemaUrn) *NullableEnummonitorHistoryPluginSchemaUrn {
	return &NullableEnummonitorHistoryPluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnummonitorHistoryPluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnummonitorHistoryPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
