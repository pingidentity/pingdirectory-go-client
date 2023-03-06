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

// EnumcustomPluginSchemaUrn the model 'EnumcustomPluginSchemaUrn'
type EnumcustomPluginSchemaUrn string

// List of Enumcustom-pluginSchemaUrn
const (
	ENUMCUSTOMPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGINCUSTOM EnumcustomPluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:custom"
)

// All allowed values of EnumcustomPluginSchemaUrn enum
var AllowedEnumcustomPluginSchemaUrnEnumValues = []EnumcustomPluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:custom",
}

func (v *EnumcustomPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcustomPluginSchemaUrn(value)
	for _, existing := range AllowedEnumcustomPluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcustomPluginSchemaUrn", value)
}

// NewEnumcustomPluginSchemaUrnFromValue returns a pointer to a valid EnumcustomPluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcustomPluginSchemaUrnFromValue(v string) (*EnumcustomPluginSchemaUrn, error) {
	ev := EnumcustomPluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcustomPluginSchemaUrn: valid values are %v", v, AllowedEnumcustomPluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcustomPluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcustomPluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumcustom-pluginSchemaUrn value
func (v EnumcustomPluginSchemaUrn) Ptr() *EnumcustomPluginSchemaUrn {
	return &v
}

type NullableEnumcustomPluginSchemaUrn struct {
	value *EnumcustomPluginSchemaUrn
	isSet bool
}

func (v NullableEnumcustomPluginSchemaUrn) Get() *EnumcustomPluginSchemaUrn {
	return v.value
}

func (v *NullableEnumcustomPluginSchemaUrn) Set(val *EnumcustomPluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcustomPluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcustomPluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcustomPluginSchemaUrn(val *EnumcustomPluginSchemaUrn) *NullableEnumcustomPluginSchemaUrn {
	return &NullableEnumcustomPluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcustomPluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcustomPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
