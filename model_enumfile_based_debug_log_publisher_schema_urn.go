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

// EnumfileBasedDebugLogPublisherSchemaUrn the model 'EnumfileBasedDebugLogPublisherSchemaUrn'
type EnumfileBasedDebugLogPublisherSchemaUrn string

// List of Enumfile-based-debug-log-publisherSchemaUrn
const (
	ENUMFILEBASEDDEBUGLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERFILE_BASED_DEBUG EnumfileBasedDebugLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:file-based-debug"
)

// All allowed values of EnumfileBasedDebugLogPublisherSchemaUrn enum
var AllowedEnumfileBasedDebugLogPublisherSchemaUrnEnumValues = []EnumfileBasedDebugLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:file-based-debug",
}

func (v *EnumfileBasedDebugLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumfileBasedDebugLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumfileBasedDebugLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumfileBasedDebugLogPublisherSchemaUrn", value)
}

// NewEnumfileBasedDebugLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumfileBasedDebugLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumfileBasedDebugLogPublisherSchemaUrnFromValue(v string) (*EnumfileBasedDebugLogPublisherSchemaUrn, error) {
	ev := EnumfileBasedDebugLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumfileBasedDebugLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumfileBasedDebugLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumfileBasedDebugLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumfileBasedDebugLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumfile-based-debug-log-publisherSchemaUrn value
func (v EnumfileBasedDebugLogPublisherSchemaUrn) Ptr() *EnumfileBasedDebugLogPublisherSchemaUrn {
	return &v
}

type NullableEnumfileBasedDebugLogPublisherSchemaUrn struct {
	value *EnumfileBasedDebugLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumfileBasedDebugLogPublisherSchemaUrn) Get() *EnumfileBasedDebugLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumfileBasedDebugLogPublisherSchemaUrn) Set(val *EnumfileBasedDebugLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumfileBasedDebugLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumfileBasedDebugLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumfileBasedDebugLogPublisherSchemaUrn(val *EnumfileBasedDebugLogPublisherSchemaUrn) *NullableEnumfileBasedDebugLogPublisherSchemaUrn {
	return &NullableEnumfileBasedDebugLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumfileBasedDebugLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumfileBasedDebugLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
