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

// EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp The final task state that a task instance should have if the task executes the specified command and that command completes with a nonzero exit code, which generally means that the command did not complete successfully.
type EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp string

// List of Enumrecurring-task-taskCompletionStateForNonzeroExitCodeProp
const (
	ENUMRECURRINGTASKTASKCOMPLETIONSTATEFORNONZEROEXITCODEPROP_STOPPED_BY_ERROR       EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp = "stopped-by-error"
	ENUMRECURRINGTASKTASKCOMPLETIONSTATEFORNONZEROEXITCODEPROP_COMPLETED_WITH_ERRORS  EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp = "completed-with-errors"
	ENUMRECURRINGTASKTASKCOMPLETIONSTATEFORNONZEROEXITCODEPROP_COMPLETED_SUCCESSFULLY EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp = "completed-successfully"
)

// All allowed values of EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp enum
var AllowedEnumrecurringTaskTaskCompletionStateForNonzeroExitCodePropEnumValues = []EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp{
	"stopped-by-error",
	"completed-with-errors",
	"completed-successfully",
}

func (v *EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp(value)
	for _, existing := range AllowedEnumrecurringTaskTaskCompletionStateForNonzeroExitCodePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp", value)
}

// NewEnumrecurringTaskTaskCompletionStateForNonzeroExitCodePropFromValue returns a pointer to a valid EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumrecurringTaskTaskCompletionStateForNonzeroExitCodePropFromValue(v string) (*EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp, error) {
	ev := EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp: valid values are %v", v, AllowedEnumrecurringTaskTaskCompletionStateForNonzeroExitCodePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) IsValid() bool {
	for _, existing := range AllowedEnumrecurringTaskTaskCompletionStateForNonzeroExitCodePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumrecurring-task-taskCompletionStateForNonzeroExitCodeProp value
func (v EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) Ptr() *EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp {
	return &v
}

type NullableEnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp struct {
	value *EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp
	isSet bool
}

func (v NullableEnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) Get() *EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp {
	return v.value
}

func (v *NullableEnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) Set(val *EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp(val *EnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) *NullableEnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp {
	return &NullableEnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp{value: val, isSet: true}
}

func (v NullableEnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumrecurringTaskTaskCompletionStateForNonzeroExitCodeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
