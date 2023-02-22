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

// EnumloggingChangeSubscriptionHandlerSchemaUrn the model 'EnumloggingChangeSubscriptionHandlerSchemaUrn'
type EnumloggingChangeSubscriptionHandlerSchemaUrn string

// List of Enumlogging-change-subscription-handlerSchemaUrn
const (
	ENUMLOGGINGCHANGESUBSCRIPTIONHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CHANGE_SUBSCRIPTION_HANDLERLOGGING EnumloggingChangeSubscriptionHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:change-subscription-handler:logging"
)

// All allowed values of EnumloggingChangeSubscriptionHandlerSchemaUrn enum
var AllowedEnumloggingChangeSubscriptionHandlerSchemaUrnEnumValues = []EnumloggingChangeSubscriptionHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:change-subscription-handler:logging",
}

func (v *EnumloggingChangeSubscriptionHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumloggingChangeSubscriptionHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumloggingChangeSubscriptionHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumloggingChangeSubscriptionHandlerSchemaUrn", value)
}

// NewEnumloggingChangeSubscriptionHandlerSchemaUrnFromValue returns a pointer to a valid EnumloggingChangeSubscriptionHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumloggingChangeSubscriptionHandlerSchemaUrnFromValue(v string) (*EnumloggingChangeSubscriptionHandlerSchemaUrn, error) {
	ev := EnumloggingChangeSubscriptionHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumloggingChangeSubscriptionHandlerSchemaUrn: valid values are %v", v, AllowedEnumloggingChangeSubscriptionHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumloggingChangeSubscriptionHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumloggingChangeSubscriptionHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlogging-change-subscription-handlerSchemaUrn value
func (v EnumloggingChangeSubscriptionHandlerSchemaUrn) Ptr() *EnumloggingChangeSubscriptionHandlerSchemaUrn {
	return &v
}

type NullableEnumloggingChangeSubscriptionHandlerSchemaUrn struct {
	value *EnumloggingChangeSubscriptionHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumloggingChangeSubscriptionHandlerSchemaUrn) Get() *EnumloggingChangeSubscriptionHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumloggingChangeSubscriptionHandlerSchemaUrn) Set(val *EnumloggingChangeSubscriptionHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumloggingChangeSubscriptionHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumloggingChangeSubscriptionHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumloggingChangeSubscriptionHandlerSchemaUrn(val *EnumloggingChangeSubscriptionHandlerSchemaUrn) *NullableEnumloggingChangeSubscriptionHandlerSchemaUrn {
	return &NullableEnumloggingChangeSubscriptionHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumloggingChangeSubscriptionHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumloggingChangeSubscriptionHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}