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

// EnumgenerateServerProfileRecurringTaskSchemaUrn the model 'EnumgenerateServerProfileRecurringTaskSchemaUrn'
type EnumgenerateServerProfileRecurringTaskSchemaUrn string

// List of Enumgenerate-server-profile-recurring-taskSchemaUrn
const (
	ENUMGENERATESERVERPROFILERECURRINGTASKSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0RECURRING_TASKGENERATE_SERVER_PROFILE EnumgenerateServerProfileRecurringTaskSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:recurring-task:generate-server-profile"
)

// All allowed values of EnumgenerateServerProfileRecurringTaskSchemaUrn enum
var AllowedEnumgenerateServerProfileRecurringTaskSchemaUrnEnumValues = []EnumgenerateServerProfileRecurringTaskSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:recurring-task:generate-server-profile",
}

func (v *EnumgenerateServerProfileRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgenerateServerProfileRecurringTaskSchemaUrn(value)
	for _, existing := range AllowedEnumgenerateServerProfileRecurringTaskSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgenerateServerProfileRecurringTaskSchemaUrn", value)
}

// NewEnumgenerateServerProfileRecurringTaskSchemaUrnFromValue returns a pointer to a valid EnumgenerateServerProfileRecurringTaskSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgenerateServerProfileRecurringTaskSchemaUrnFromValue(v string) (*EnumgenerateServerProfileRecurringTaskSchemaUrn, error) {
	ev := EnumgenerateServerProfileRecurringTaskSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgenerateServerProfileRecurringTaskSchemaUrn: valid values are %v", v, AllowedEnumgenerateServerProfileRecurringTaskSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgenerateServerProfileRecurringTaskSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgenerateServerProfileRecurringTaskSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgenerate-server-profile-recurring-taskSchemaUrn value
func (v EnumgenerateServerProfileRecurringTaskSchemaUrn) Ptr() *EnumgenerateServerProfileRecurringTaskSchemaUrn {
	return &v
}

type NullableEnumgenerateServerProfileRecurringTaskSchemaUrn struct {
	value *EnumgenerateServerProfileRecurringTaskSchemaUrn
	isSet bool
}

func (v NullableEnumgenerateServerProfileRecurringTaskSchemaUrn) Get() *EnumgenerateServerProfileRecurringTaskSchemaUrn {
	return v.value
}

func (v *NullableEnumgenerateServerProfileRecurringTaskSchemaUrn) Set(val *EnumgenerateServerProfileRecurringTaskSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgenerateServerProfileRecurringTaskSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgenerateServerProfileRecurringTaskSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgenerateServerProfileRecurringTaskSchemaUrn(val *EnumgenerateServerProfileRecurringTaskSchemaUrn) *NullableEnumgenerateServerProfileRecurringTaskSchemaUrn {
	return &NullableEnumgenerateServerProfileRecurringTaskSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgenerateServerProfileRecurringTaskSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgenerateServerProfileRecurringTaskSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
