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

// EnumcollectSupportDataRecurringTaskSchemaUrn the model 'EnumcollectSupportDataRecurringTaskSchemaUrn'
type EnumcollectSupportDataRecurringTaskSchemaUrn string

// List of Enumcollect-support-data-recurring-taskSchemaUrn
const (
	ENUMCOLLECTSUPPORTDATARECURRINGTASKSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0RECURRING_TASKCOLLECT_SUPPORT_DATA EnumcollectSupportDataRecurringTaskSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:recurring-task:collect-support-data"
)

// All allowed values of EnumcollectSupportDataRecurringTaskSchemaUrn enum
var AllowedEnumcollectSupportDataRecurringTaskSchemaUrnEnumValues = []EnumcollectSupportDataRecurringTaskSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:recurring-task:collect-support-data",
}

func (v *EnumcollectSupportDataRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcollectSupportDataRecurringTaskSchemaUrn(value)
	for _, existing := range AllowedEnumcollectSupportDataRecurringTaskSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcollectSupportDataRecurringTaskSchemaUrn", value)
}

// NewEnumcollectSupportDataRecurringTaskSchemaUrnFromValue returns a pointer to a valid EnumcollectSupportDataRecurringTaskSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcollectSupportDataRecurringTaskSchemaUrnFromValue(v string) (*EnumcollectSupportDataRecurringTaskSchemaUrn, error) {
	ev := EnumcollectSupportDataRecurringTaskSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcollectSupportDataRecurringTaskSchemaUrn: valid values are %v", v, AllowedEnumcollectSupportDataRecurringTaskSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcollectSupportDataRecurringTaskSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcollectSupportDataRecurringTaskSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumcollect-support-data-recurring-taskSchemaUrn value
func (v EnumcollectSupportDataRecurringTaskSchemaUrn) Ptr() *EnumcollectSupportDataRecurringTaskSchemaUrn {
	return &v
}

type NullableEnumcollectSupportDataRecurringTaskSchemaUrn struct {
	value *EnumcollectSupportDataRecurringTaskSchemaUrn
	isSet bool
}

func (v NullableEnumcollectSupportDataRecurringTaskSchemaUrn) Get() *EnumcollectSupportDataRecurringTaskSchemaUrn {
	return v.value
}

func (v *NullableEnumcollectSupportDataRecurringTaskSchemaUrn) Set(val *EnumcollectSupportDataRecurringTaskSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcollectSupportDataRecurringTaskSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcollectSupportDataRecurringTaskSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcollectSupportDataRecurringTaskSchemaUrn(val *EnumcollectSupportDataRecurringTaskSchemaUrn) *NullableEnumcollectSupportDataRecurringTaskSchemaUrn {
	return &NullableEnumcollectSupportDataRecurringTaskSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcollectSupportDataRecurringTaskSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcollectSupportDataRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

