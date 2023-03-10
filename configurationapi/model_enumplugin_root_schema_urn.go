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

// EnumpluginRootSchemaUrn the model 'EnumpluginRootSchemaUrn'
type EnumpluginRootSchemaUrn string

// List of Enumplugin-rootSchemaUrn
const (
	ENUMPLUGINROOTSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGIN_ROOT EnumpluginRootSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin-root"
)

// All allowed values of EnumpluginRootSchemaUrn enum
var AllowedEnumpluginRootSchemaUrnEnumValues = []EnumpluginRootSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin-root",
}

func (v *EnumpluginRootSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginRootSchemaUrn(value)
	for _, existing := range AllowedEnumpluginRootSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginRootSchemaUrn", value)
}

// NewEnumpluginRootSchemaUrnFromValue returns a pointer to a valid EnumpluginRootSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginRootSchemaUrnFromValue(v string) (*EnumpluginRootSchemaUrn, error) {
	ev := EnumpluginRootSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginRootSchemaUrn: valid values are %v", v, AllowedEnumpluginRootSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginRootSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumpluginRootSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-rootSchemaUrn value
func (v EnumpluginRootSchemaUrn) Ptr() *EnumpluginRootSchemaUrn {
	return &v
}

type NullableEnumpluginRootSchemaUrn struct {
	value *EnumpluginRootSchemaUrn
	isSet bool
}

func (v NullableEnumpluginRootSchemaUrn) Get() *EnumpluginRootSchemaUrn {
	return v.value
}

func (v *NullableEnumpluginRootSchemaUrn) Set(val *EnumpluginRootSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginRootSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginRootSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginRootSchemaUrn(val *EnumpluginRootSchemaUrn) *NullableEnumpluginRootSchemaUrn {
	return &NullableEnumpluginRootSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumpluginRootSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginRootSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
