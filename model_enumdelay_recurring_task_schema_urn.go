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

// EnumdelayRecurringTaskSchemaUrn the model 'EnumdelayRecurringTaskSchemaUrn'
type EnumdelayRecurringTaskSchemaUrn string

// List of Enumdelay-recurring-taskSchemaUrn
const (
	ENUMDELAYRECURRINGTASKSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0RECURRING_TASKDELAY EnumdelayRecurringTaskSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:recurring-task:delay"
)

// All allowed values of EnumdelayRecurringTaskSchemaUrn enum
var AllowedEnumdelayRecurringTaskSchemaUrnEnumValues = []EnumdelayRecurringTaskSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:recurring-task:delay",
}

func (v *EnumdelayRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumdelayRecurringTaskSchemaUrn(value)
	for _, existing := range AllowedEnumdelayRecurringTaskSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumdelayRecurringTaskSchemaUrn", value)
}

// NewEnumdelayRecurringTaskSchemaUrnFromValue returns a pointer to a valid EnumdelayRecurringTaskSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumdelayRecurringTaskSchemaUrnFromValue(v string) (*EnumdelayRecurringTaskSchemaUrn, error) {
	ev := EnumdelayRecurringTaskSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumdelayRecurringTaskSchemaUrn: valid values are %v", v, AllowedEnumdelayRecurringTaskSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumdelayRecurringTaskSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumdelayRecurringTaskSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumdelay-recurring-taskSchemaUrn value
func (v EnumdelayRecurringTaskSchemaUrn) Ptr() *EnumdelayRecurringTaskSchemaUrn {
	return &v
}

type NullableEnumdelayRecurringTaskSchemaUrn struct {
	value *EnumdelayRecurringTaskSchemaUrn
	isSet bool
}

func (v NullableEnumdelayRecurringTaskSchemaUrn) Get() *EnumdelayRecurringTaskSchemaUrn {
	return v.value
}

func (v *NullableEnumdelayRecurringTaskSchemaUrn) Set(val *EnumdelayRecurringTaskSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumdelayRecurringTaskSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumdelayRecurringTaskSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumdelayRecurringTaskSchemaUrn(val *EnumdelayRecurringTaskSchemaUrn) *NullableEnumdelayRecurringTaskSchemaUrn {
	return &NullableEnumdelayRecurringTaskSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumdelayRecurringTaskSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumdelayRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
