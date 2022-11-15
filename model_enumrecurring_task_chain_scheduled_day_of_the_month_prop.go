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

// EnumrecurringTaskChainScheduledDayOfTheMonthProp The specific days of the month on which instances of this Recurring Task Chain may be scheduled to start. If the scheduled-day-selection-type property has a value of selected-days-of-the-month, then this property must have one or more values; otherwise, it must be left undefined.
type EnumrecurringTaskChainScheduledDayOfTheMonthProp string

// List of Enumrecurring-task-chain-scheduledDayOfTheMonthProp
const (
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_01 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-01"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_02 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-02"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_03 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-03"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_04 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-04"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_05 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-05"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_06 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-06"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_07 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-07"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_08 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-08"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_09 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-09"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_10 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-10"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_11 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-11"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_12 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-12"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_13 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-13"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_14 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-14"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_15 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-15"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_16 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-16"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_17 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-17"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_18 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-18"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_19 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-19"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_20 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-20"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_21 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-21"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_22 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-22"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_23 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-23"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_24 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-24"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_25 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-25"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_26 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-26"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_27 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-27"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_28 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-28"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_29 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-29"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_30 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-30"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_DAY_31 EnumrecurringTaskChainScheduledDayOfTheMonthProp = "day-31"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_NEXT_TO_LAST_DAY_OF_THE_MONTH EnumrecurringTaskChainScheduledDayOfTheMonthProp = "next-to-last-day-of-the-month"
	ENUMRECURRINGTASKCHAINSCHEDULEDDAYOFTHEMONTHPROP_LAST_DAY_OF_THE_MONTH EnumrecurringTaskChainScheduledDayOfTheMonthProp = "last-day-of-the-month"
)

// All allowed values of EnumrecurringTaskChainScheduledDayOfTheMonthProp enum
var AllowedEnumrecurringTaskChainScheduledDayOfTheMonthPropEnumValues = []EnumrecurringTaskChainScheduledDayOfTheMonthProp{
	"day-01",
	"day-02",
	"day-03",
	"day-04",
	"day-05",
	"day-06",
	"day-07",
	"day-08",
	"day-09",
	"day-10",
	"day-11",
	"day-12",
	"day-13",
	"day-14",
	"day-15",
	"day-16",
	"day-17",
	"day-18",
	"day-19",
	"day-20",
	"day-21",
	"day-22",
	"day-23",
	"day-24",
	"day-25",
	"day-26",
	"day-27",
	"day-28",
	"day-29",
	"day-30",
	"day-31",
	"next-to-last-day-of-the-month",
	"last-day-of-the-month",
}

func (v *EnumrecurringTaskChainScheduledDayOfTheMonthProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumrecurringTaskChainScheduledDayOfTheMonthProp(value)
	for _, existing := range AllowedEnumrecurringTaskChainScheduledDayOfTheMonthPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumrecurringTaskChainScheduledDayOfTheMonthProp", value)
}

// NewEnumrecurringTaskChainScheduledDayOfTheMonthPropFromValue returns a pointer to a valid EnumrecurringTaskChainScheduledDayOfTheMonthProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumrecurringTaskChainScheduledDayOfTheMonthPropFromValue(v string) (*EnumrecurringTaskChainScheduledDayOfTheMonthProp, error) {
	ev := EnumrecurringTaskChainScheduledDayOfTheMonthProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumrecurringTaskChainScheduledDayOfTheMonthProp: valid values are %v", v, AllowedEnumrecurringTaskChainScheduledDayOfTheMonthPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumrecurringTaskChainScheduledDayOfTheMonthProp) IsValid() bool {
	for _, existing := range AllowedEnumrecurringTaskChainScheduledDayOfTheMonthPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumrecurring-task-chain-scheduledDayOfTheMonthProp value
func (v EnumrecurringTaskChainScheduledDayOfTheMonthProp) Ptr() *EnumrecurringTaskChainScheduledDayOfTheMonthProp {
	return &v
}

type NullableEnumrecurringTaskChainScheduledDayOfTheMonthProp struct {
	value *EnumrecurringTaskChainScheduledDayOfTheMonthProp
	isSet bool
}

func (v NullableEnumrecurringTaskChainScheduledDayOfTheMonthProp) Get() *EnumrecurringTaskChainScheduledDayOfTheMonthProp {
	return v.value
}

func (v *NullableEnumrecurringTaskChainScheduledDayOfTheMonthProp) Set(val *EnumrecurringTaskChainScheduledDayOfTheMonthProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumrecurringTaskChainScheduledDayOfTheMonthProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumrecurringTaskChainScheduledDayOfTheMonthProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumrecurringTaskChainScheduledDayOfTheMonthProp(val *EnumrecurringTaskChainScheduledDayOfTheMonthProp) *NullableEnumrecurringTaskChainScheduledDayOfTheMonthProp {
	return &NullableEnumrecurringTaskChainScheduledDayOfTheMonthProp{value: val, isSet: true}
}

func (v NullableEnumrecurringTaskChainScheduledDayOfTheMonthProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumrecurringTaskChainScheduledDayOfTheMonthProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

