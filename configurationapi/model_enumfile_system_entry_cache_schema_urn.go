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

// EnumfileSystemEntryCacheSchemaUrn the model 'EnumfileSystemEntryCacheSchemaUrn'
type EnumfileSystemEntryCacheSchemaUrn string

// List of Enumfile-system-entry-cacheSchemaUrn
const (
	ENUMFILESYSTEMENTRYCACHESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ENTRY_CACHEFILE_SYSTEM EnumfileSystemEntryCacheSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:entry-cache:file-system"
)

// All allowed values of EnumfileSystemEntryCacheSchemaUrn enum
var AllowedEnumfileSystemEntryCacheSchemaUrnEnumValues = []EnumfileSystemEntryCacheSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:entry-cache:file-system",
}

func (v *EnumfileSystemEntryCacheSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumfileSystemEntryCacheSchemaUrn(value)
	for _, existing := range AllowedEnumfileSystemEntryCacheSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumfileSystemEntryCacheSchemaUrn", value)
}

// NewEnumfileSystemEntryCacheSchemaUrnFromValue returns a pointer to a valid EnumfileSystemEntryCacheSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumfileSystemEntryCacheSchemaUrnFromValue(v string) (*EnumfileSystemEntryCacheSchemaUrn, error) {
	ev := EnumfileSystemEntryCacheSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumfileSystemEntryCacheSchemaUrn: valid values are %v", v, AllowedEnumfileSystemEntryCacheSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumfileSystemEntryCacheSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumfileSystemEntryCacheSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumfile-system-entry-cacheSchemaUrn value
func (v EnumfileSystemEntryCacheSchemaUrn) Ptr() *EnumfileSystemEntryCacheSchemaUrn {
	return &v
}

type NullableEnumfileSystemEntryCacheSchemaUrn struct {
	value *EnumfileSystemEntryCacheSchemaUrn
	isSet bool
}

func (v NullableEnumfileSystemEntryCacheSchemaUrn) Get() *EnumfileSystemEntryCacheSchemaUrn {
	return v.value
}

func (v *NullableEnumfileSystemEntryCacheSchemaUrn) Set(val *EnumfileSystemEntryCacheSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumfileSystemEntryCacheSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumfileSystemEntryCacheSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumfileSystemEntryCacheSchemaUrn(val *EnumfileSystemEntryCacheSchemaUrn) *NullableEnumfileSystemEntryCacheSchemaUrn {
	return &NullableEnumfileSystemEntryCacheSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumfileSystemEntryCacheSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumfileSystemEntryCacheSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
