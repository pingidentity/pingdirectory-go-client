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

// EnumfileBasedTraceLogPublisherSchemaUrn the model 'EnumfileBasedTraceLogPublisherSchemaUrn'
type EnumfileBasedTraceLogPublisherSchemaUrn string

// List of Enumfile-based-trace-log-publisherSchemaUrn
const (
	ENUMFILEBASEDTRACELOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERFILE_BASED_TRACE EnumfileBasedTraceLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:file-based-trace"
)

// All allowed values of EnumfileBasedTraceLogPublisherSchemaUrn enum
var AllowedEnumfileBasedTraceLogPublisherSchemaUrnEnumValues = []EnumfileBasedTraceLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:file-based-trace",
}

func (v *EnumfileBasedTraceLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumfileBasedTraceLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumfileBasedTraceLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumfileBasedTraceLogPublisherSchemaUrn", value)
}

// NewEnumfileBasedTraceLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumfileBasedTraceLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumfileBasedTraceLogPublisherSchemaUrnFromValue(v string) (*EnumfileBasedTraceLogPublisherSchemaUrn, error) {
	ev := EnumfileBasedTraceLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumfileBasedTraceLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumfileBasedTraceLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumfileBasedTraceLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumfileBasedTraceLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumfile-based-trace-log-publisherSchemaUrn value
func (v EnumfileBasedTraceLogPublisherSchemaUrn) Ptr() *EnumfileBasedTraceLogPublisherSchemaUrn {
	return &v
}

type NullableEnumfileBasedTraceLogPublisherSchemaUrn struct {
	value *EnumfileBasedTraceLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumfileBasedTraceLogPublisherSchemaUrn) Get() *EnumfileBasedTraceLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumfileBasedTraceLogPublisherSchemaUrn) Set(val *EnumfileBasedTraceLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumfileBasedTraceLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumfileBasedTraceLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumfileBasedTraceLogPublisherSchemaUrn(val *EnumfileBasedTraceLogPublisherSchemaUrn) *NullableEnumfileBasedTraceLogPublisherSchemaUrn {
	return &NullableEnumfileBasedTraceLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumfileBasedTraceLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumfileBasedTraceLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
