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

// EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn the model 'EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn'
type EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn string

// List of Enumget-symmetric-key-extended-operation-handlerSchemaUrn
const (
	ENUMGETSYMMETRICKEYEXTENDEDOPERATIONHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTENDED_OPERATION_HANDLERGET_SYMMETRIC_KEY EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:get-symmetric-key"
)

// All allowed values of EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn enum
var AllowedEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrnEnumValues = []EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:get-symmetric-key",
}

func (v *EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn", value)
}

// NewEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrnFromValue returns a pointer to a valid EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrnFromValue(v string) (*EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn, error) {
	ev := EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn: valid values are %v", v, AllowedEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumget-symmetric-key-extended-operation-handlerSchemaUrn value
func (v EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn) Ptr() *EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn {
	return &v
}

type NullableEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn struct {
	value *EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn) Get() *EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn) Set(val *EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn(val *EnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn) *NullableEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn {
	return &NullableEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgetSymmetricKeyExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
