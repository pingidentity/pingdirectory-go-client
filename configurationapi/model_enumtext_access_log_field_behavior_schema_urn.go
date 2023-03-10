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

// EnumtextAccessLogFieldBehaviorSchemaUrn the model 'EnumtextAccessLogFieldBehaviorSchemaUrn'
type EnumtextAccessLogFieldBehaviorSchemaUrn string

// List of Enumtext-access-log-field-behaviorSchemaUrn
const (
	ENUMTEXTACCESSLOGFIELDBEHAVIORSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_FIELD_BEHAVIORTEXT_ACCESS EnumtextAccessLogFieldBehaviorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-field-behavior:text-access"
)

// All allowed values of EnumtextAccessLogFieldBehaviorSchemaUrn enum
var AllowedEnumtextAccessLogFieldBehaviorSchemaUrnEnumValues = []EnumtextAccessLogFieldBehaviorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-field-behavior:text-access",
}

func (v *EnumtextAccessLogFieldBehaviorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumtextAccessLogFieldBehaviorSchemaUrn(value)
	for _, existing := range AllowedEnumtextAccessLogFieldBehaviorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumtextAccessLogFieldBehaviorSchemaUrn", value)
}

// NewEnumtextAccessLogFieldBehaviorSchemaUrnFromValue returns a pointer to a valid EnumtextAccessLogFieldBehaviorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumtextAccessLogFieldBehaviorSchemaUrnFromValue(v string) (*EnumtextAccessLogFieldBehaviorSchemaUrn, error) {
	ev := EnumtextAccessLogFieldBehaviorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumtextAccessLogFieldBehaviorSchemaUrn: valid values are %v", v, AllowedEnumtextAccessLogFieldBehaviorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumtextAccessLogFieldBehaviorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumtextAccessLogFieldBehaviorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumtext-access-log-field-behaviorSchemaUrn value
func (v EnumtextAccessLogFieldBehaviorSchemaUrn) Ptr() *EnumtextAccessLogFieldBehaviorSchemaUrn {
	return &v
}

type NullableEnumtextAccessLogFieldBehaviorSchemaUrn struct {
	value *EnumtextAccessLogFieldBehaviorSchemaUrn
	isSet bool
}

func (v NullableEnumtextAccessLogFieldBehaviorSchemaUrn) Get() *EnumtextAccessLogFieldBehaviorSchemaUrn {
	return v.value
}

func (v *NullableEnumtextAccessLogFieldBehaviorSchemaUrn) Set(val *EnumtextAccessLogFieldBehaviorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumtextAccessLogFieldBehaviorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumtextAccessLogFieldBehaviorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumtextAccessLogFieldBehaviorSchemaUrn(val *EnumtextAccessLogFieldBehaviorSchemaUrn) *NullableEnumtextAccessLogFieldBehaviorSchemaUrn {
	return &NullableEnumtextAccessLogFieldBehaviorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumtextAccessLogFieldBehaviorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumtextAccessLogFieldBehaviorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
