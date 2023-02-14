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

// EnumlastModPluginSchemaUrn the model 'EnumlastModPluginSchemaUrn'
type EnumlastModPluginSchemaUrn string

// List of Enumlast-mod-pluginSchemaUrn
const (
	ENUMLASTMODPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGINLAST_MOD EnumlastModPluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:last-mod"
)

// All allowed values of EnumlastModPluginSchemaUrn enum
var AllowedEnumlastModPluginSchemaUrnEnumValues = []EnumlastModPluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:last-mod",
}

func (v *EnumlastModPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlastModPluginSchemaUrn(value)
	for _, existing := range AllowedEnumlastModPluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlastModPluginSchemaUrn", value)
}

// NewEnumlastModPluginSchemaUrnFromValue returns a pointer to a valid EnumlastModPluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlastModPluginSchemaUrnFromValue(v string) (*EnumlastModPluginSchemaUrn, error) {
	ev := EnumlastModPluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlastModPluginSchemaUrn: valid values are %v", v, AllowedEnumlastModPluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlastModPluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumlastModPluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlast-mod-pluginSchemaUrn value
func (v EnumlastModPluginSchemaUrn) Ptr() *EnumlastModPluginSchemaUrn {
	return &v
}

type NullableEnumlastModPluginSchemaUrn struct {
	value *EnumlastModPluginSchemaUrn
	isSet bool
}

func (v NullableEnumlastModPluginSchemaUrn) Get() *EnumlastModPluginSchemaUrn {
	return v.value
}

func (v *NullableEnumlastModPluginSchemaUrn) Set(val *EnumlastModPluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlastModPluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlastModPluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlastModPluginSchemaUrn(val *EnumlastModPluginSchemaUrn) *NullableEnumlastModPluginSchemaUrn {
	return &NullableEnumlastModPluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumlastModPluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlastModPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
