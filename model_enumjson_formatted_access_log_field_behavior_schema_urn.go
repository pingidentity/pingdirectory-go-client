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

// EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn the model 'EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn'
type EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn string

// List of Enumjson-formatted-access-log-field-behaviorSchemaUrn
const (
	ENUMJSONFORMATTEDACCESSLOGFIELDBEHAVIORSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_FIELD_BEHAVIORJSON_FORMATTED_ACCESS EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-field-behavior:json-formatted-access"
)

// All allowed values of EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn enum
var AllowedEnumjsonFormattedAccessLogFieldBehaviorSchemaUrnEnumValues = []EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-field-behavior:json-formatted-access",
}

func (v *EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn(value)
	for _, existing := range AllowedEnumjsonFormattedAccessLogFieldBehaviorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn", value)
}

// NewEnumjsonFormattedAccessLogFieldBehaviorSchemaUrnFromValue returns a pointer to a valid EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumjsonFormattedAccessLogFieldBehaviorSchemaUrnFromValue(v string) (*EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn, error) {
	ev := EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn: valid values are %v", v, AllowedEnumjsonFormattedAccessLogFieldBehaviorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumjsonFormattedAccessLogFieldBehaviorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumjson-formatted-access-log-field-behaviorSchemaUrn value
func (v EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) Ptr() *EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn {
	return &v
}

type NullableEnumjsonFormattedAccessLogFieldBehaviorSchemaUrn struct {
	value *EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn
	isSet bool
}

func (v NullableEnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) Get() *EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn {
	return v.value
}

func (v *NullableEnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) Set(val *EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumjsonFormattedAccessLogFieldBehaviorSchemaUrn(val *EnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) *NullableEnumjsonFormattedAccessLogFieldBehaviorSchemaUrn {
	return &NullableEnumjsonFormattedAccessLogFieldBehaviorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumjsonFormattedAccessLogFieldBehaviorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
