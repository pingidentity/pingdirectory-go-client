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

// EnumbackupRecurringTaskSchemaUrn the model 'EnumbackupRecurringTaskSchemaUrn'
type EnumbackupRecurringTaskSchemaUrn string

// List of Enumbackup-recurring-taskSchemaUrn
const (
	ENUMBACKUPRECURRINGTASKSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0RECURRING_TASKBACKUP EnumbackupRecurringTaskSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:recurring-task:backup"
)

// All allowed values of EnumbackupRecurringTaskSchemaUrn enum
var AllowedEnumbackupRecurringTaskSchemaUrnEnumValues = []EnumbackupRecurringTaskSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:recurring-task:backup",
}

func (v *EnumbackupRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumbackupRecurringTaskSchemaUrn(value)
	for _, existing := range AllowedEnumbackupRecurringTaskSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumbackupRecurringTaskSchemaUrn", value)
}

// NewEnumbackupRecurringTaskSchemaUrnFromValue returns a pointer to a valid EnumbackupRecurringTaskSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumbackupRecurringTaskSchemaUrnFromValue(v string) (*EnumbackupRecurringTaskSchemaUrn, error) {
	ev := EnumbackupRecurringTaskSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumbackupRecurringTaskSchemaUrn: valid values are %v", v, AllowedEnumbackupRecurringTaskSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumbackupRecurringTaskSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumbackupRecurringTaskSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumbackup-recurring-taskSchemaUrn value
func (v EnumbackupRecurringTaskSchemaUrn) Ptr() *EnumbackupRecurringTaskSchemaUrn {
	return &v
}

type NullableEnumbackupRecurringTaskSchemaUrn struct {
	value *EnumbackupRecurringTaskSchemaUrn
	isSet bool
}

func (v NullableEnumbackupRecurringTaskSchemaUrn) Get() *EnumbackupRecurringTaskSchemaUrn {
	return v.value
}

func (v *NullableEnumbackupRecurringTaskSchemaUrn) Set(val *EnumbackupRecurringTaskSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumbackupRecurringTaskSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumbackupRecurringTaskSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumbackupRecurringTaskSchemaUrn(val *EnumbackupRecurringTaskSchemaUrn) *NullableEnumbackupRecurringTaskSchemaUrn {
	return &NullableEnumbackupRecurringTaskSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumbackupRecurringTaskSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumbackupRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
