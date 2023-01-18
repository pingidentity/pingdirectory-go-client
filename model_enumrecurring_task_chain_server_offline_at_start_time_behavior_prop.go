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

// EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp Specifies the behavior that the server should exhibit if it is offline when the start time arrives for the tasks in this Recurring Task Chain.
type EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp string

// List of Enumrecurring-task-chain-serverOfflineAtStartTimeBehaviorProp
const (
	ENUMRECURRINGTASKCHAINSERVEROFFLINEATSTARTTIMEBEHAVIORPROP_RUN_IMMEDIATELY_UPON_SERVER_STARTUP EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp = "run-immediately-upon-server-startup"
	ENUMRECURRINGTASKCHAINSERVEROFFLINEATSTARTTIMEBEHAVIORPROP_CANCEL_ITERATION_AND_WAIT_FOR_NEXT_SCHEDULED_START_TIME EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp = "cancel-iteration-and-wait-for-next-scheduled-start-time"
)

// All allowed values of EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp enum
var AllowedEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorPropEnumValues = []EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp{
	"run-immediately-upon-server-startup",
	"cancel-iteration-and-wait-for-next-scheduled-start-time",
}

func (v *EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp(value)
	for _, existing := range AllowedEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp", value)
}

// NewEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorPropFromValue returns a pointer to a valid EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorPropFromValue(v string) (*EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp, error) {
	ev := EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp: valid values are %v", v, AllowedEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumrecurring-task-chain-serverOfflineAtStartTimeBehaviorProp value
func (v EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) Ptr() *EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp {
	return &v
}

type NullableEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp struct {
	value *EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp
	isSet bool
}

func (v NullableEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) Get() *EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp {
	return v.value
}

func (v *NullableEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) Set(val *EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp(val *EnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) *NullableEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp {
	return &NullableEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumrecurringTaskChainServerOfflineAtStartTimeBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
