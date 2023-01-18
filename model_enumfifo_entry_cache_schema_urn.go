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

// EnumfifoEntryCacheSchemaUrn the model 'EnumfifoEntryCacheSchemaUrn'
type EnumfifoEntryCacheSchemaUrn string

// List of Enumfifo-entry-cacheSchemaUrn
const (
	ENUMFIFOENTRYCACHESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ENTRY_CACHEFIFO EnumfifoEntryCacheSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:entry-cache:fifo"
)

// All allowed values of EnumfifoEntryCacheSchemaUrn enum
var AllowedEnumfifoEntryCacheSchemaUrnEnumValues = []EnumfifoEntryCacheSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:entry-cache:fifo",
}

func (v *EnumfifoEntryCacheSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumfifoEntryCacheSchemaUrn(value)
	for _, existing := range AllowedEnumfifoEntryCacheSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumfifoEntryCacheSchemaUrn", value)
}

// NewEnumfifoEntryCacheSchemaUrnFromValue returns a pointer to a valid EnumfifoEntryCacheSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumfifoEntryCacheSchemaUrnFromValue(v string) (*EnumfifoEntryCacheSchemaUrn, error) {
	ev := EnumfifoEntryCacheSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumfifoEntryCacheSchemaUrn: valid values are %v", v, AllowedEnumfifoEntryCacheSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumfifoEntryCacheSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumfifoEntryCacheSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumfifo-entry-cacheSchemaUrn value
func (v EnumfifoEntryCacheSchemaUrn) Ptr() *EnumfifoEntryCacheSchemaUrn {
	return &v
}

type NullableEnumfifoEntryCacheSchemaUrn struct {
	value *EnumfifoEntryCacheSchemaUrn
	isSet bool
}

func (v NullableEnumfifoEntryCacheSchemaUrn) Get() *EnumfifoEntryCacheSchemaUrn {
	return v.value
}

func (v *NullableEnumfifoEntryCacheSchemaUrn) Set(val *EnumfifoEntryCacheSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumfifoEntryCacheSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumfifoEntryCacheSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumfifoEntryCacheSchemaUrn(val *EnumfifoEntryCacheSchemaUrn) *NullableEnumfifoEntryCacheSchemaUrn {
	return &NullableEnumfifoEntryCacheSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumfifoEntryCacheSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumfifoEntryCacheSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
