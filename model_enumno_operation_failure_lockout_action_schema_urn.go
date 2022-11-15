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

// EnumnoOperationFailureLockoutActionSchemaUrn the model 'EnumnoOperationFailureLockoutActionSchemaUrn'
type EnumnoOperationFailureLockoutActionSchemaUrn string

// List of Enumno-operation-failure-lockout-actionSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0FAILURE_LOCKOUT_ACTIONNO_OPERATION EnumnoOperationFailureLockoutActionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:failure-lockout-action:no-operation"
)

// All allowed values of EnumnoOperationFailureLockoutActionSchemaUrn enum
var AllowedEnumnoOperationFailureLockoutActionSchemaUrnEnumValues = []EnumnoOperationFailureLockoutActionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:failure-lockout-action:no-operation",
}

func (v *EnumnoOperationFailureLockoutActionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumnoOperationFailureLockoutActionSchemaUrn(value)
	for _, existing := range AllowedEnumnoOperationFailureLockoutActionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumnoOperationFailureLockoutActionSchemaUrn", value)
}

// NewEnumnoOperationFailureLockoutActionSchemaUrnFromValue returns a pointer to a valid EnumnoOperationFailureLockoutActionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumnoOperationFailureLockoutActionSchemaUrnFromValue(v string) (*EnumnoOperationFailureLockoutActionSchemaUrn, error) {
	ev := EnumnoOperationFailureLockoutActionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumnoOperationFailureLockoutActionSchemaUrn: valid values are %v", v, AllowedEnumnoOperationFailureLockoutActionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumnoOperationFailureLockoutActionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumnoOperationFailureLockoutActionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumno-operation-failure-lockout-actionSchemaUrn value
func (v EnumnoOperationFailureLockoutActionSchemaUrn) Ptr() *EnumnoOperationFailureLockoutActionSchemaUrn {
	return &v
}

type NullableEnumnoOperationFailureLockoutActionSchemaUrn struct {
	value *EnumnoOperationFailureLockoutActionSchemaUrn
	isSet bool
}

func (v NullableEnumnoOperationFailureLockoutActionSchemaUrn) Get() *EnumnoOperationFailureLockoutActionSchemaUrn {
	return v.value
}

func (v *NullableEnumnoOperationFailureLockoutActionSchemaUrn) Set(val *EnumnoOperationFailureLockoutActionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumnoOperationFailureLockoutActionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumnoOperationFailureLockoutActionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumnoOperationFailureLockoutActionSchemaUrn(val *EnumnoOperationFailureLockoutActionSchemaUrn) *NullableEnumnoOperationFailureLockoutActionSchemaUrn {
	return &NullableEnumnoOperationFailureLockoutActionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumnoOperationFailureLockoutActionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumnoOperationFailureLockoutActionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

