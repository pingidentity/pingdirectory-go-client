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

// EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn the model 'EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn'
type EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn string

// List of Enuminteractive-transactions-extended-operation-handlerSchemaUrn
const (
	ENUMINTERACTIVETRANSACTIONSEXTENDEDOPERATIONHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTENDED_OPERATION_HANDLERINTERACTIVE_TRANSACTIONS EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:interactive-transactions"
)

// All allowed values of EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn enum
var AllowedEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrnEnumValues = []EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:interactive-transactions",
}

func (v *EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn(value)
	for _, existing := range AllowedEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn", value)
}

// NewEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrnFromValue returns a pointer to a valid EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrnFromValue(v string) (*EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn, error) {
	ev := EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn: valid values are %v", v, AllowedEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enuminteractive-transactions-extended-operation-handlerSchemaUrn value
func (v EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn) Ptr() *EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn {
	return &v
}

type NullableEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn struct {
	value *EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn
	isSet bool
}

func (v NullableEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn) Get() *EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn) Set(val *EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn(val *EnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn) *NullableEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn {
	return &NullableEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnuminteractiveTransactionsExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}