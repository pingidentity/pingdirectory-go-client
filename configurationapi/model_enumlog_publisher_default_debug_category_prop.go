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

// EnumlogPublisherDefaultDebugCategoryProp The debug message categories to be logged when none of the defined targets match the message.
type EnumlogPublisherDefaultDebugCategoryProp string

// List of Enumlog-publisher-defaultDebugCategoryProp
const (
	ENUMLOGPUBLISHERDEFAULTDEBUGCATEGORYPROP_CAUGHT          EnumlogPublisherDefaultDebugCategoryProp = "caught"
	ENUMLOGPUBLISHERDEFAULTDEBUGCATEGORYPROP_CONSTRUCTOR     EnumlogPublisherDefaultDebugCategoryProp = "constructor"
	ENUMLOGPUBLISHERDEFAULTDEBUGCATEGORYPROP_DATA            EnumlogPublisherDefaultDebugCategoryProp = "data"
	ENUMLOGPUBLISHERDEFAULTDEBUGCATEGORYPROP_DATABASE_ACCESS EnumlogPublisherDefaultDebugCategoryProp = "database-access"
	ENUMLOGPUBLISHERDEFAULTDEBUGCATEGORYPROP_ENTER           EnumlogPublisherDefaultDebugCategoryProp = "enter"
	ENUMLOGPUBLISHERDEFAULTDEBUGCATEGORYPROP_EXIT            EnumlogPublisherDefaultDebugCategoryProp = "exit"
	ENUMLOGPUBLISHERDEFAULTDEBUGCATEGORYPROP_MESSAGE         EnumlogPublisherDefaultDebugCategoryProp = "message"
	ENUMLOGPUBLISHERDEFAULTDEBUGCATEGORYPROP_PROTOCOL        EnumlogPublisherDefaultDebugCategoryProp = "protocol"
	ENUMLOGPUBLISHERDEFAULTDEBUGCATEGORYPROP_THROWN          EnumlogPublisherDefaultDebugCategoryProp = "thrown"
)

// All allowed values of EnumlogPublisherDefaultDebugCategoryProp enum
var AllowedEnumlogPublisherDefaultDebugCategoryPropEnumValues = []EnumlogPublisherDefaultDebugCategoryProp{
	"caught",
	"constructor",
	"data",
	"database-access",
	"enter",
	"exit",
	"message",
	"protocol",
	"thrown",
}

func (v *EnumlogPublisherDefaultDebugCategoryProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherDefaultDebugCategoryProp(value)
	for _, existing := range AllowedEnumlogPublisherDefaultDebugCategoryPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherDefaultDebugCategoryProp", value)
}

// NewEnumlogPublisherDefaultDebugCategoryPropFromValue returns a pointer to a valid EnumlogPublisherDefaultDebugCategoryProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherDefaultDebugCategoryPropFromValue(v string) (*EnumlogPublisherDefaultDebugCategoryProp, error) {
	ev := EnumlogPublisherDefaultDebugCategoryProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherDefaultDebugCategoryProp: valid values are %v", v, AllowedEnumlogPublisherDefaultDebugCategoryPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherDefaultDebugCategoryProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherDefaultDebugCategoryPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-defaultDebugCategoryProp value
func (v EnumlogPublisherDefaultDebugCategoryProp) Ptr() *EnumlogPublisherDefaultDebugCategoryProp {
	return &v
}

type NullableEnumlogPublisherDefaultDebugCategoryProp struct {
	value *EnumlogPublisherDefaultDebugCategoryProp
	isSet bool
}

func (v NullableEnumlogPublisherDefaultDebugCategoryProp) Get() *EnumlogPublisherDefaultDebugCategoryProp {
	return v.value
}

func (v *NullableEnumlogPublisherDefaultDebugCategoryProp) Set(val *EnumlogPublisherDefaultDebugCategoryProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherDefaultDebugCategoryProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherDefaultDebugCategoryProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherDefaultDebugCategoryProp(val *EnumlogPublisherDefaultDebugCategoryProp) *NullableEnumlogPublisherDefaultDebugCategoryProp {
	return &NullableEnumlogPublisherDefaultDebugCategoryProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherDefaultDebugCategoryProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherDefaultDebugCategoryProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
