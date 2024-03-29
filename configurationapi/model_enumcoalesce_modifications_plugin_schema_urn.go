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

// EnumcoalesceModificationsPluginSchemaUrn the model 'EnumcoalesceModificationsPluginSchemaUrn'
type EnumcoalesceModificationsPluginSchemaUrn string

// List of Enumcoalesce-modifications-pluginSchemaUrn
const (
	ENUMCOALESCEMODIFICATIONSPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGINCOALESCE_MODIFICATIONS EnumcoalesceModificationsPluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:coalesce-modifications"
)

// All allowed values of EnumcoalesceModificationsPluginSchemaUrn enum
var AllowedEnumcoalesceModificationsPluginSchemaUrnEnumValues = []EnumcoalesceModificationsPluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:coalesce-modifications",
}

func (v *EnumcoalesceModificationsPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcoalesceModificationsPluginSchemaUrn(value)
	for _, existing := range AllowedEnumcoalesceModificationsPluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcoalesceModificationsPluginSchemaUrn", value)
}

// NewEnumcoalesceModificationsPluginSchemaUrnFromValue returns a pointer to a valid EnumcoalesceModificationsPluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcoalesceModificationsPluginSchemaUrnFromValue(v string) (*EnumcoalesceModificationsPluginSchemaUrn, error) {
	ev := EnumcoalesceModificationsPluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcoalesceModificationsPluginSchemaUrn: valid values are %v", v, AllowedEnumcoalesceModificationsPluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcoalesceModificationsPluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcoalesceModificationsPluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumcoalesce-modifications-pluginSchemaUrn value
func (v EnumcoalesceModificationsPluginSchemaUrn) Ptr() *EnumcoalesceModificationsPluginSchemaUrn {
	return &v
}

type NullableEnumcoalesceModificationsPluginSchemaUrn struct {
	value *EnumcoalesceModificationsPluginSchemaUrn
	isSet bool
}

func (v NullableEnumcoalesceModificationsPluginSchemaUrn) Get() *EnumcoalesceModificationsPluginSchemaUrn {
	return v.value
}

func (v *NullableEnumcoalesceModificationsPluginSchemaUrn) Set(val *EnumcoalesceModificationsPluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcoalesceModificationsPluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcoalesceModificationsPluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcoalesceModificationsPluginSchemaUrn(val *EnumcoalesceModificationsPluginSchemaUrn) *NullableEnumcoalesceModificationsPluginSchemaUrn {
	return &NullableEnumcoalesceModificationsPluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcoalesceModificationsPluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcoalesceModificationsPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
