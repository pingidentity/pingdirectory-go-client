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

// EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp Specifies the audit behavior for delete and modify operations on soft-deleted entries.
type EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp string

// List of Enumlog-publisher-file-based-json-audit-softDeleteEntryAuditBehaviorProp
const (
	ENUMLOGPUBLISHERFILEBASEDJSONAUDITSOFTDELETEENTRYAUDITBEHAVIORPROP_INCLUDED EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp = "included"
	ENUMLOGPUBLISHERFILEBASEDJSONAUDITSOFTDELETEENTRYAUDITBEHAVIORPROP_EXCLUDED EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp = "excluded"
)

// All allowed values of EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp enum
var AllowedEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorPropEnumValues = []EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp{
	"included",
	"excluded",
}

func (v *EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp(value)
	for _, existing := range AllowedEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp", value)
}

// NewEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorPropFromValue returns a pointer to a valid EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorPropFromValue(v string) (*EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp, error) {
	ev := EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp: valid values are %v", v, AllowedEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-file-based-json-audit-softDeleteEntryAuditBehaviorProp value
func (v EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp) Ptr() *EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp {
	return &v
}

type NullableEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp struct {
	value *EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp
	isSet bool
}

func (v NullableEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp) Get() *EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp {
	return v.value
}

func (v *NullableEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp) Set(val *EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp(val *EnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp) *NullableEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp {
	return &NullableEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherFileBasedJsonAuditSoftDeleteEntryAuditBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
