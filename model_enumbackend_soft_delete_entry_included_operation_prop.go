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

// EnumbackendSoftDeleteEntryIncludedOperationProp Specifies which operations performed on soft-deleted entries will appear in the changelog.
type EnumbackendSoftDeleteEntryIncludedOperationProp string

// List of Enumbackend-softDeleteEntryIncludedOperationProp
const (
	ENUMBACKENDSOFTDELETEENTRYINCLUDEDOPERATIONPROP_MODIFY EnumbackendSoftDeleteEntryIncludedOperationProp = "modify"
	ENUMBACKENDSOFTDELETEENTRYINCLUDEDOPERATIONPROP_DELETE EnumbackendSoftDeleteEntryIncludedOperationProp = "delete"
)

// All allowed values of EnumbackendSoftDeleteEntryIncludedOperationProp enum
var AllowedEnumbackendSoftDeleteEntryIncludedOperationPropEnumValues = []EnumbackendSoftDeleteEntryIncludedOperationProp{
	"modify",
	"delete",
}

func (v *EnumbackendSoftDeleteEntryIncludedOperationProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumbackendSoftDeleteEntryIncludedOperationProp(value)
	for _, existing := range AllowedEnumbackendSoftDeleteEntryIncludedOperationPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumbackendSoftDeleteEntryIncludedOperationProp", value)
}

// NewEnumbackendSoftDeleteEntryIncludedOperationPropFromValue returns a pointer to a valid EnumbackendSoftDeleteEntryIncludedOperationProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumbackendSoftDeleteEntryIncludedOperationPropFromValue(v string) (*EnumbackendSoftDeleteEntryIncludedOperationProp, error) {
	ev := EnumbackendSoftDeleteEntryIncludedOperationProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumbackendSoftDeleteEntryIncludedOperationProp: valid values are %v", v, AllowedEnumbackendSoftDeleteEntryIncludedOperationPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumbackendSoftDeleteEntryIncludedOperationProp) IsValid() bool {
	for _, existing := range AllowedEnumbackendSoftDeleteEntryIncludedOperationPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumbackend-softDeleteEntryIncludedOperationProp value
func (v EnumbackendSoftDeleteEntryIncludedOperationProp) Ptr() *EnumbackendSoftDeleteEntryIncludedOperationProp {
	return &v
}

type NullableEnumbackendSoftDeleteEntryIncludedOperationProp struct {
	value *EnumbackendSoftDeleteEntryIncludedOperationProp
	isSet bool
}

func (v NullableEnumbackendSoftDeleteEntryIncludedOperationProp) Get() *EnumbackendSoftDeleteEntryIncludedOperationProp {
	return v.value
}

func (v *NullableEnumbackendSoftDeleteEntryIncludedOperationProp) Set(val *EnumbackendSoftDeleteEntryIncludedOperationProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumbackendSoftDeleteEntryIncludedOperationProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumbackendSoftDeleteEntryIncludedOperationProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumbackendSoftDeleteEntryIncludedOperationProp(val *EnumbackendSoftDeleteEntryIncludedOperationProp) *NullableEnumbackendSoftDeleteEntryIncludedOperationProp {
	return &NullableEnumbackendSoftDeleteEntryIncludedOperationProp{value: val, isSet: true}
}

func (v NullableEnumbackendSoftDeleteEntryIncludedOperationProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumbackendSoftDeleteEntryIncludedOperationProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

