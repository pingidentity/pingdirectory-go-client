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

// EnumjsonErrorLogPublisherSchemaUrn the model 'EnumjsonErrorLogPublisherSchemaUrn'
type EnumjsonErrorLogPublisherSchemaUrn string

// List of Enumjson-error-log-publisherSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERJSON_ERROR EnumjsonErrorLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:json-error"
)

// All allowed values of EnumjsonErrorLogPublisherSchemaUrn enum
var AllowedEnumjsonErrorLogPublisherSchemaUrnEnumValues = []EnumjsonErrorLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:json-error",
}

func (v *EnumjsonErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumjsonErrorLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumjsonErrorLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumjsonErrorLogPublisherSchemaUrn", value)
}

// NewEnumjsonErrorLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumjsonErrorLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumjsonErrorLogPublisherSchemaUrnFromValue(v string) (*EnumjsonErrorLogPublisherSchemaUrn, error) {
	ev := EnumjsonErrorLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumjsonErrorLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumjsonErrorLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumjsonErrorLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumjsonErrorLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumjson-error-log-publisherSchemaUrn value
func (v EnumjsonErrorLogPublisherSchemaUrn) Ptr() *EnumjsonErrorLogPublisherSchemaUrn {
	return &v
}

type NullableEnumjsonErrorLogPublisherSchemaUrn struct {
	value *EnumjsonErrorLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumjsonErrorLogPublisherSchemaUrn) Get() *EnumjsonErrorLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumjsonErrorLogPublisherSchemaUrn) Set(val *EnumjsonErrorLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumjsonErrorLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumjsonErrorLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumjsonErrorLogPublisherSchemaUrn(val *EnumjsonErrorLogPublisherSchemaUrn) *NullableEnumjsonErrorLogPublisherSchemaUrn {
	return &NullableEnumjsonErrorLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumjsonErrorLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumjsonErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

