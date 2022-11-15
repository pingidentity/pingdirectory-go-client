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

// EnumrecurringTaskChainInterruptedByShutdownBehaviorProp Specifies the behavior that the server should exhibit if it is shut down or abnormally terminated while an instance of this Recurring Task Chain is running.
type EnumrecurringTaskChainInterruptedByShutdownBehaviorProp string

// List of Enumrecurring-task-chain-interruptedByShutdownBehaviorProp
const (
	ENUMRECURRINGTASKCHAININTERRUPTEDBYSHUTDOWNBEHAVIORPROP_INTERRUPTED_TASK_AND_DEPENDENCIES EnumrecurringTaskChainInterruptedByShutdownBehaviorProp = "cancel-interrupted-task-and-dependencies"
	ENUMRECURRINGTASKCHAININTERRUPTEDBYSHUTDOWNBEHAVIORPROP_ONLY_INTERRUPTED_TASK_BUT_PRESERVE_DEPENDENCIES EnumrecurringTaskChainInterruptedByShutdownBehaviorProp = "cancel-only-interrupted-task-but-preserve-dependencies"
)

// All allowed values of EnumrecurringTaskChainInterruptedByShutdownBehaviorProp enum
var AllowedEnumrecurringTaskChainInterruptedByShutdownBehaviorPropEnumValues = []EnumrecurringTaskChainInterruptedByShutdownBehaviorProp{
	"cancel-interrupted-task-and-dependencies",
	"cancel-only-interrupted-task-but-preserve-dependencies",
}

func (v *EnumrecurringTaskChainInterruptedByShutdownBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumrecurringTaskChainInterruptedByShutdownBehaviorProp(value)
	for _, existing := range AllowedEnumrecurringTaskChainInterruptedByShutdownBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumrecurringTaskChainInterruptedByShutdownBehaviorProp", value)
}

// NewEnumrecurringTaskChainInterruptedByShutdownBehaviorPropFromValue returns a pointer to a valid EnumrecurringTaskChainInterruptedByShutdownBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumrecurringTaskChainInterruptedByShutdownBehaviorPropFromValue(v string) (*EnumrecurringTaskChainInterruptedByShutdownBehaviorProp, error) {
	ev := EnumrecurringTaskChainInterruptedByShutdownBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumrecurringTaskChainInterruptedByShutdownBehaviorProp: valid values are %v", v, AllowedEnumrecurringTaskChainInterruptedByShutdownBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumrecurringTaskChainInterruptedByShutdownBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumrecurringTaskChainInterruptedByShutdownBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumrecurring-task-chain-interruptedByShutdownBehaviorProp value
func (v EnumrecurringTaskChainInterruptedByShutdownBehaviorProp) Ptr() *EnumrecurringTaskChainInterruptedByShutdownBehaviorProp {
	return &v
}

type NullableEnumrecurringTaskChainInterruptedByShutdownBehaviorProp struct {
	value *EnumrecurringTaskChainInterruptedByShutdownBehaviorProp
	isSet bool
}

func (v NullableEnumrecurringTaskChainInterruptedByShutdownBehaviorProp) Get() *EnumrecurringTaskChainInterruptedByShutdownBehaviorProp {
	return v.value
}

func (v *NullableEnumrecurringTaskChainInterruptedByShutdownBehaviorProp) Set(val *EnumrecurringTaskChainInterruptedByShutdownBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumrecurringTaskChainInterruptedByShutdownBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumrecurringTaskChainInterruptedByShutdownBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumrecurringTaskChainInterruptedByShutdownBehaviorProp(val *EnumrecurringTaskChainInterruptedByShutdownBehaviorProp) *NullableEnumrecurringTaskChainInterruptedByShutdownBehaviorProp {
	return &NullableEnumrecurringTaskChainInterruptedByShutdownBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumrecurringTaskChainInterruptedByShutdownBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumrecurringTaskChainInterruptedByShutdownBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

