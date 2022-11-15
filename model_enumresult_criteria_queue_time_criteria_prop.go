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

// EnumresultCriteriaQueueTimeCriteriaProp Indicates whether the time the operation was required to wait on the work queue should be taken into consideration when determining whether to include the operation in this Simple Result Criteria. If the queue time should be taken into account, then the \"queue-time-value\" property should contain the boundary value. This property should only be given a value other than \"any\" if the work queue has been configured to monitor the time operations have spent on the work queue.
type EnumresultCriteriaQueueTimeCriteriaProp string

// List of Enumresult-criteria-queueTimeCriteriaProp
const (
	ANY EnumresultCriteriaQueueTimeCriteriaProp = "any"
	LESS_THAN_OR_EQUAL_TO EnumresultCriteriaQueueTimeCriteriaProp = "less-than-or-equal-to"
	GREATER_THAN_OR_EQUAL_TO EnumresultCriteriaQueueTimeCriteriaProp = "greater-than-or-equal-to"
)

// All allowed values of EnumresultCriteriaQueueTimeCriteriaProp enum
var AllowedEnumresultCriteriaQueueTimeCriteriaPropEnumValues = []EnumresultCriteriaQueueTimeCriteriaProp{
	"any",
	"less-than-or-equal-to",
	"greater-than-or-equal-to",
}

func (v *EnumresultCriteriaQueueTimeCriteriaProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumresultCriteriaQueueTimeCriteriaProp(value)
	for _, existing := range AllowedEnumresultCriteriaQueueTimeCriteriaPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumresultCriteriaQueueTimeCriteriaProp", value)
}

// NewEnumresultCriteriaQueueTimeCriteriaPropFromValue returns a pointer to a valid EnumresultCriteriaQueueTimeCriteriaProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumresultCriteriaQueueTimeCriteriaPropFromValue(v string) (*EnumresultCriteriaQueueTimeCriteriaProp, error) {
	ev := EnumresultCriteriaQueueTimeCriteriaProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumresultCriteriaQueueTimeCriteriaProp: valid values are %v", v, AllowedEnumresultCriteriaQueueTimeCriteriaPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumresultCriteriaQueueTimeCriteriaProp) IsValid() bool {
	for _, existing := range AllowedEnumresultCriteriaQueueTimeCriteriaPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumresult-criteria-queueTimeCriteriaProp value
func (v EnumresultCriteriaQueueTimeCriteriaProp) Ptr() *EnumresultCriteriaQueueTimeCriteriaProp {
	return &v
}

type NullableEnumresultCriteriaQueueTimeCriteriaProp struct {
	value *EnumresultCriteriaQueueTimeCriteriaProp
	isSet bool
}

func (v NullableEnumresultCriteriaQueueTimeCriteriaProp) Get() *EnumresultCriteriaQueueTimeCriteriaProp {
	return v.value
}

func (v *NullableEnumresultCriteriaQueueTimeCriteriaProp) Set(val *EnumresultCriteriaQueueTimeCriteriaProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumresultCriteriaQueueTimeCriteriaProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumresultCriteriaQueueTimeCriteriaProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumresultCriteriaQueueTimeCriteriaProp(val *EnumresultCriteriaQueueTimeCriteriaProp) *NullableEnumresultCriteriaQueueTimeCriteriaProp {
	return &NullableEnumresultCriteriaQueueTimeCriteriaProp{value: val, isSet: true}
}

func (v NullableEnumresultCriteriaQueueTimeCriteriaProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumresultCriteriaQueueTimeCriteriaProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

