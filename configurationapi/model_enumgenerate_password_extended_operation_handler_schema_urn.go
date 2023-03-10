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

// EnumgeneratePasswordExtendedOperationHandlerSchemaUrn the model 'EnumgeneratePasswordExtendedOperationHandlerSchemaUrn'
type EnumgeneratePasswordExtendedOperationHandlerSchemaUrn string

// List of Enumgenerate-password-extended-operation-handlerSchemaUrn
const (
	ENUMGENERATEPASSWORDEXTENDEDOPERATIONHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTENDED_OPERATION_HANDLERGENERATE_PASSWORD EnumgeneratePasswordExtendedOperationHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:generate-password"
)

// All allowed values of EnumgeneratePasswordExtendedOperationHandlerSchemaUrn enum
var AllowedEnumgeneratePasswordExtendedOperationHandlerSchemaUrnEnumValues = []EnumgeneratePasswordExtendedOperationHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:generate-password",
}

func (v *EnumgeneratePasswordExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgeneratePasswordExtendedOperationHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumgeneratePasswordExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgeneratePasswordExtendedOperationHandlerSchemaUrn", value)
}

// NewEnumgeneratePasswordExtendedOperationHandlerSchemaUrnFromValue returns a pointer to a valid EnumgeneratePasswordExtendedOperationHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgeneratePasswordExtendedOperationHandlerSchemaUrnFromValue(v string) (*EnumgeneratePasswordExtendedOperationHandlerSchemaUrn, error) {
	ev := EnumgeneratePasswordExtendedOperationHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgeneratePasswordExtendedOperationHandlerSchemaUrn: valid values are %v", v, AllowedEnumgeneratePasswordExtendedOperationHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgeneratePasswordExtendedOperationHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgeneratePasswordExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgenerate-password-extended-operation-handlerSchemaUrn value
func (v EnumgeneratePasswordExtendedOperationHandlerSchemaUrn) Ptr() *EnumgeneratePasswordExtendedOperationHandlerSchemaUrn {
	return &v
}

type NullableEnumgeneratePasswordExtendedOperationHandlerSchemaUrn struct {
	value *EnumgeneratePasswordExtendedOperationHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumgeneratePasswordExtendedOperationHandlerSchemaUrn) Get() *EnumgeneratePasswordExtendedOperationHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumgeneratePasswordExtendedOperationHandlerSchemaUrn) Set(val *EnumgeneratePasswordExtendedOperationHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgeneratePasswordExtendedOperationHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgeneratePasswordExtendedOperationHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgeneratePasswordExtendedOperationHandlerSchemaUrn(val *EnumgeneratePasswordExtendedOperationHandlerSchemaUrn) *NullableEnumgeneratePasswordExtendedOperationHandlerSchemaUrn {
	return &NullableEnumgeneratePasswordExtendedOperationHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgeneratePasswordExtendedOperationHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgeneratePasswordExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
