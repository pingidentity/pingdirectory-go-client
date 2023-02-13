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

// EnumdnMapperPluginSchemaUrn the model 'EnumdnMapperPluginSchemaUrn'
type EnumdnMapperPluginSchemaUrn string

// List of Enumdn-mapper-pluginSchemaUrn
const (
	ENUMDNMAPPERPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGINDN_MAPPER EnumdnMapperPluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:dn-mapper"
)

// All allowed values of EnumdnMapperPluginSchemaUrn enum
var AllowedEnumdnMapperPluginSchemaUrnEnumValues = []EnumdnMapperPluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:dn-mapper",
}

func (v *EnumdnMapperPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdnMapperPluginSchemaUrn(value)
	for _, existing := range AllowedEnumdnMapperPluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdnMapperPluginSchemaUrn", value)
}

// NewEnumdnMapperPluginSchemaUrnFromValue returns a pointer to a valid EnumdnMapperPluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdnMapperPluginSchemaUrnFromValue(v string) (*EnumdnMapperPluginSchemaUrn, error) {
	ev := EnumdnMapperPluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdnMapperPluginSchemaUrn: valid values are %v", v, AllowedEnumdnMapperPluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdnMapperPluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdnMapperPluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdn-mapper-pluginSchemaUrn value
func (v EnumdnMapperPluginSchemaUrn) Ptr() *EnumdnMapperPluginSchemaUrn {
	return &v
}

type NullableEnumdnMapperPluginSchemaUrn struct {
	value *EnumdnMapperPluginSchemaUrn
	isSet bool
}

func (v NullableEnumdnMapperPluginSchemaUrn) Get() *EnumdnMapperPluginSchemaUrn {
	return v.value
}

func (v *NullableEnumdnMapperPluginSchemaUrn) Set(val *EnumdnMapperPluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdnMapperPluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdnMapperPluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdnMapperPluginSchemaUrn(val *EnumdnMapperPluginSchemaUrn) *NullableEnumdnMapperPluginSchemaUrn {
	return &NullableEnumdnMapperPluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdnMapperPluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdnMapperPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
