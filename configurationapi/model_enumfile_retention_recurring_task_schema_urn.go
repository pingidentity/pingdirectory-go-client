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

// EnumfileRetentionRecurringTaskSchemaUrn the model 'EnumfileRetentionRecurringTaskSchemaUrn'
type EnumfileRetentionRecurringTaskSchemaUrn string

// List of Enumfile-retention-recurring-taskSchemaUrn
const (
	ENUMFILERETENTIONRECURRINGTASKSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0RECURRING_TASKFILE_RETENTION EnumfileRetentionRecurringTaskSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:recurring-task:file-retention"
)

// All allowed values of EnumfileRetentionRecurringTaskSchemaUrn enum
var AllowedEnumfileRetentionRecurringTaskSchemaUrnEnumValues = []EnumfileRetentionRecurringTaskSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:recurring-task:file-retention",
}

func (v *EnumfileRetentionRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumfileRetentionRecurringTaskSchemaUrn(value)
	for _, existing := range AllowedEnumfileRetentionRecurringTaskSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumfileRetentionRecurringTaskSchemaUrn", value)
}

// NewEnumfileRetentionRecurringTaskSchemaUrnFromValue returns a pointer to a valid EnumfileRetentionRecurringTaskSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumfileRetentionRecurringTaskSchemaUrnFromValue(v string) (*EnumfileRetentionRecurringTaskSchemaUrn, error) {
	ev := EnumfileRetentionRecurringTaskSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumfileRetentionRecurringTaskSchemaUrn: valid values are %v", v, AllowedEnumfileRetentionRecurringTaskSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumfileRetentionRecurringTaskSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumfileRetentionRecurringTaskSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumfile-retention-recurring-taskSchemaUrn value
func (v EnumfileRetentionRecurringTaskSchemaUrn) Ptr() *EnumfileRetentionRecurringTaskSchemaUrn {
	return &v
}

type NullableEnumfileRetentionRecurringTaskSchemaUrn struct {
	value *EnumfileRetentionRecurringTaskSchemaUrn
	isSet bool
}

func (v NullableEnumfileRetentionRecurringTaskSchemaUrn) Get() *EnumfileRetentionRecurringTaskSchemaUrn {
	return v.value
}

func (v *NullableEnumfileRetentionRecurringTaskSchemaUrn) Set(val *EnumfileRetentionRecurringTaskSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumfileRetentionRecurringTaskSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumfileRetentionRecurringTaskSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumfileRetentionRecurringTaskSchemaUrn(val *EnumfileRetentionRecurringTaskSchemaUrn) *NullableEnumfileRetentionRecurringTaskSchemaUrn {
	return &NullableEnumfileRetentionRecurringTaskSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumfileRetentionRecurringTaskSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumfileRetentionRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}