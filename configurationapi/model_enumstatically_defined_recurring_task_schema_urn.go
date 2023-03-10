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

// EnumstaticallyDefinedRecurringTaskSchemaUrn the model 'EnumstaticallyDefinedRecurringTaskSchemaUrn'
type EnumstaticallyDefinedRecurringTaskSchemaUrn string

// List of Enumstatically-defined-recurring-taskSchemaUrn
const (
	ENUMSTATICALLYDEFINEDRECURRINGTASKSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0RECURRING_TASKSTATICALLY_DEFINED EnumstaticallyDefinedRecurringTaskSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:recurring-task:statically-defined"
)

// All allowed values of EnumstaticallyDefinedRecurringTaskSchemaUrn enum
var AllowedEnumstaticallyDefinedRecurringTaskSchemaUrnEnumValues = []EnumstaticallyDefinedRecurringTaskSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:recurring-task:statically-defined",
}

func (v *EnumstaticallyDefinedRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumstaticallyDefinedRecurringTaskSchemaUrn(value)
	for _, existing := range AllowedEnumstaticallyDefinedRecurringTaskSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumstaticallyDefinedRecurringTaskSchemaUrn", value)
}

// NewEnumstaticallyDefinedRecurringTaskSchemaUrnFromValue returns a pointer to a valid EnumstaticallyDefinedRecurringTaskSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumstaticallyDefinedRecurringTaskSchemaUrnFromValue(v string) (*EnumstaticallyDefinedRecurringTaskSchemaUrn, error) {
	ev := EnumstaticallyDefinedRecurringTaskSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumstaticallyDefinedRecurringTaskSchemaUrn: valid values are %v", v, AllowedEnumstaticallyDefinedRecurringTaskSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumstaticallyDefinedRecurringTaskSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumstaticallyDefinedRecurringTaskSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumstatically-defined-recurring-taskSchemaUrn value
func (v EnumstaticallyDefinedRecurringTaskSchemaUrn) Ptr() *EnumstaticallyDefinedRecurringTaskSchemaUrn {
	return &v
}

type NullableEnumstaticallyDefinedRecurringTaskSchemaUrn struct {
	value *EnumstaticallyDefinedRecurringTaskSchemaUrn
	isSet bool
}

func (v NullableEnumstaticallyDefinedRecurringTaskSchemaUrn) Get() *EnumstaticallyDefinedRecurringTaskSchemaUrn {
	return v.value
}

func (v *NullableEnumstaticallyDefinedRecurringTaskSchemaUrn) Set(val *EnumstaticallyDefinedRecurringTaskSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumstaticallyDefinedRecurringTaskSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumstaticallyDefinedRecurringTaskSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumstaticallyDefinedRecurringTaskSchemaUrn(val *EnumstaticallyDefinedRecurringTaskSchemaUrn) *NullableEnumstaticallyDefinedRecurringTaskSchemaUrn {
	return &NullableEnumstaticallyDefinedRecurringTaskSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumstaticallyDefinedRecurringTaskSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumstaticallyDefinedRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
