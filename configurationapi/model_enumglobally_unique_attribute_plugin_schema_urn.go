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

// EnumgloballyUniqueAttributePluginSchemaUrn the model 'EnumgloballyUniqueAttributePluginSchemaUrn'
type EnumgloballyUniqueAttributePluginSchemaUrn string

// List of Enumglobally-unique-attribute-pluginSchemaUrn
const (
	ENUMGLOBALLYUNIQUEATTRIBUTEPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGINGLOBALLY_UNIQUE_ATTRIBUTE EnumgloballyUniqueAttributePluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:globally-unique-attribute"
)

// All allowed values of EnumgloballyUniqueAttributePluginSchemaUrn enum
var AllowedEnumgloballyUniqueAttributePluginSchemaUrnEnumValues = []EnumgloballyUniqueAttributePluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:globally-unique-attribute",
}

func (v *EnumgloballyUniqueAttributePluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgloballyUniqueAttributePluginSchemaUrn(value)
	for _, existing := range AllowedEnumgloballyUniqueAttributePluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgloballyUniqueAttributePluginSchemaUrn", value)
}

// NewEnumgloballyUniqueAttributePluginSchemaUrnFromValue returns a pointer to a valid EnumgloballyUniqueAttributePluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgloballyUniqueAttributePluginSchemaUrnFromValue(v string) (*EnumgloballyUniqueAttributePluginSchemaUrn, error) {
	ev := EnumgloballyUniqueAttributePluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgloballyUniqueAttributePluginSchemaUrn: valid values are %v", v, AllowedEnumgloballyUniqueAttributePluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgloballyUniqueAttributePluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgloballyUniqueAttributePluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumglobally-unique-attribute-pluginSchemaUrn value
func (v EnumgloballyUniqueAttributePluginSchemaUrn) Ptr() *EnumgloballyUniqueAttributePluginSchemaUrn {
	return &v
}

type NullableEnumgloballyUniqueAttributePluginSchemaUrn struct {
	value *EnumgloballyUniqueAttributePluginSchemaUrn
	isSet bool
}

func (v NullableEnumgloballyUniqueAttributePluginSchemaUrn) Get() *EnumgloballyUniqueAttributePluginSchemaUrn {
	return v.value
}

func (v *NullableEnumgloballyUniqueAttributePluginSchemaUrn) Set(val *EnumgloballyUniqueAttributePluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgloballyUniqueAttributePluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgloballyUniqueAttributePluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgloballyUniqueAttributePluginSchemaUrn(val *EnumgloballyUniqueAttributePluginSchemaUrn) *NullableEnumgloballyUniqueAttributePluginSchemaUrn {
	return &NullableEnumgloballyUniqueAttributePluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgloballyUniqueAttributePluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgloballyUniqueAttributePluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
