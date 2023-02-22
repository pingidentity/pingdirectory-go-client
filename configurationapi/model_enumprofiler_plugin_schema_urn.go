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

// EnumprofilerPluginSchemaUrn the model 'EnumprofilerPluginSchemaUrn'
type EnumprofilerPluginSchemaUrn string

// List of Enumprofiler-pluginSchemaUrn
const (
	ENUMPROFILERPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGINPROFILER EnumprofilerPluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:profiler"
)

// All allowed values of EnumprofilerPluginSchemaUrn enum
var AllowedEnumprofilerPluginSchemaUrnEnumValues = []EnumprofilerPluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:profiler",
}

func (v *EnumprofilerPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumprofilerPluginSchemaUrn(value)
	for _, existing := range AllowedEnumprofilerPluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumprofilerPluginSchemaUrn", value)
}

// NewEnumprofilerPluginSchemaUrnFromValue returns a pointer to a valid EnumprofilerPluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumprofilerPluginSchemaUrnFromValue(v string) (*EnumprofilerPluginSchemaUrn, error) {
	ev := EnumprofilerPluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumprofilerPluginSchemaUrn: valid values are %v", v, AllowedEnumprofilerPluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumprofilerPluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumprofilerPluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumprofiler-pluginSchemaUrn value
func (v EnumprofilerPluginSchemaUrn) Ptr() *EnumprofilerPluginSchemaUrn {
	return &v
}

type NullableEnumprofilerPluginSchemaUrn struct {
	value *EnumprofilerPluginSchemaUrn
	isSet bool
}

func (v NullableEnumprofilerPluginSchemaUrn) Get() *EnumprofilerPluginSchemaUrn {
	return v.value
}

func (v *NullableEnumprofilerPluginSchemaUrn) Set(val *EnumprofilerPluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumprofilerPluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumprofilerPluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumprofilerPluginSchemaUrn(val *EnumprofilerPluginSchemaUrn) *NullableEnumprofilerPluginSchemaUrn {
	return &NullableEnumprofilerPluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumprofilerPluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumprofilerPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}