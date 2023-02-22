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

// EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn the model 'EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn'
type EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn string

// List of Enumget-changelog-batch-extended-operation-handlerSchemaUrn
const (
	ENUMGETCHANGELOGBATCHEXTENDEDOPERATIONHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTENDED_OPERATION_HANDLERGET_CHANGELOG_BATCH EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:get-changelog-batch"
)

// All allowed values of EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn enum
var AllowedEnumgetChangelogBatchExtendedOperationHandlerSchemaUrnEnumValues = []EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:get-changelog-batch",
}

func (v *EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumgetChangelogBatchExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn", value)
}

// NewEnumgetChangelogBatchExtendedOperationHandlerSchemaUrnFromValue returns a pointer to a valid EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgetChangelogBatchExtendedOperationHandlerSchemaUrnFromValue(v string) (*EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn, error) {
	ev := EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn: valid values are %v", v, AllowedEnumgetChangelogBatchExtendedOperationHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgetChangelogBatchExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumget-changelog-batch-extended-operation-handlerSchemaUrn value
func (v EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn) Ptr() *EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn {
	return &v
}

type NullableEnumgetChangelogBatchExtendedOperationHandlerSchemaUrn struct {
	value *EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumgetChangelogBatchExtendedOperationHandlerSchemaUrn) Get() *EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumgetChangelogBatchExtendedOperationHandlerSchemaUrn) Set(val *EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgetChangelogBatchExtendedOperationHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgetChangelogBatchExtendedOperationHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgetChangelogBatchExtendedOperationHandlerSchemaUrn(val *EnumgetChangelogBatchExtendedOperationHandlerSchemaUrn) *NullableEnumgetChangelogBatchExtendedOperationHandlerSchemaUrn {
	return &NullableEnumgetChangelogBatchExtendedOperationHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgetChangelogBatchExtendedOperationHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgetChangelogBatchExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}