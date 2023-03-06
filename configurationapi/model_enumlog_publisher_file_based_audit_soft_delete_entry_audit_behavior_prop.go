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

// EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp Specifies the audit behavior for delete and modify operations on soft-deleted entries.
type EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp string

// List of Enumlog-publisher-file-based-audit-softDeleteEntryAuditBehaviorProp
const (
	ENUMLOGPUBLISHERFILEBASEDAUDITSOFTDELETEENTRYAUDITBEHAVIORPROP_COMMENTED   EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp = "commented"
	ENUMLOGPUBLISHERFILEBASEDAUDITSOFTDELETEENTRYAUDITBEHAVIORPROP_UNCOMMENTED EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp = "uncommented"
)

// All allowed values of EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp enum
var AllowedEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorPropEnumValues = []EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp{
	"commented",
	"uncommented",
}

func (v *EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp(value)
	for _, existing := range AllowedEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp", value)
}

// NewEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorPropFromValue returns a pointer to a valid EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorPropFromValue(v string) (*EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp, error) {
	ev := EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp: valid values are %v", v, AllowedEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-file-based-audit-softDeleteEntryAuditBehaviorProp value
func (v EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp) Ptr() *EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp {
	return &v
}

type NullableEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp struct {
	value *EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp
	isSet bool
}

func (v NullableEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp) Get() *EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp {
	return v.value
}

func (v *NullableEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp) Set(val *EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp(val *EnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp) *NullableEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp {
	return &NullableEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherFileBasedAuditSoftDeleteEntryAuditBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}