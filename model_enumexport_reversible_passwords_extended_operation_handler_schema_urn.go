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

// EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn the model 'EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn'
type EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn string

// List of Enumexport-reversible-passwords-extended-operation-handlerSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTENDED_OPERATION_HANDLEREXPORT_REVERSIBLE_PASSWORDS EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:export-reversible-passwords"
)

// All allowed values of EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn enum
var AllowedEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrnEnumValues = []EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:export-reversible-passwords",
}

func (v *EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn", value)
}

// NewEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrnFromValue returns a pointer to a valid EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrnFromValue(v string) (*EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn, error) {
	ev := EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn: valid values are %v", v, AllowedEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexport-reversible-passwords-extended-operation-handlerSchemaUrn value
func (v EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) Ptr() *EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn {
	return &v
}

type NullableEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn struct {
	value *EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) Get() *EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) Set(val *EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn(val *EnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) *NullableEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn {
	return &NullableEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexportReversiblePasswordsExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

