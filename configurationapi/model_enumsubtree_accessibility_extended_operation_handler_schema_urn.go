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

// EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn the model 'EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn'
type EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn string

// List of Enumsubtree-accessibility-extended-operation-handlerSchemaUrn
const (
	ENUMSUBTREEACCESSIBILITYEXTENDEDOPERATIONHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTENDED_OPERATION_HANDLERSUBTREE_ACCESSIBILITY EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:subtree-accessibility"
)

// All allowed values of EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn enum
var AllowedEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrnEnumValues = []EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:extended-operation-handler:subtree-accessibility",
}

func (v *EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn", value)
}

// NewEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrnFromValue returns a pointer to a valid EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrnFromValue(v string) (*EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn, error) {
	ev := EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn: valid values are %v", v, AllowedEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsubtree-accessibility-extended-operation-handlerSchemaUrn value
func (v EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn) Ptr() *EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn {
	return &v
}

type NullableEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn struct {
	value *EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn) Get() *EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn) Set(val *EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn(val *EnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn) *NullableEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn {
	return &NullableEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsubtreeAccessibilityExtendedOperationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
