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

// EnumchangeSubscriptionSchemaUrn the model 'EnumchangeSubscriptionSchemaUrn'
type EnumchangeSubscriptionSchemaUrn string

// List of Enumchange-subscriptionSchemaUrn
const (
	ENUMCHANGESUBSCRIPTIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CHANGE_SUBSCRIPTION EnumchangeSubscriptionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:change-subscription"
)

// All allowed values of EnumchangeSubscriptionSchemaUrn enum
var AllowedEnumchangeSubscriptionSchemaUrnEnumValues = []EnumchangeSubscriptionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:change-subscription",
}

func (v *EnumchangeSubscriptionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumchangeSubscriptionSchemaUrn(value)
	for _, existing := range AllowedEnumchangeSubscriptionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumchangeSubscriptionSchemaUrn", value)
}

// NewEnumchangeSubscriptionSchemaUrnFromValue returns a pointer to a valid EnumchangeSubscriptionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumchangeSubscriptionSchemaUrnFromValue(v string) (*EnumchangeSubscriptionSchemaUrn, error) {
	ev := EnumchangeSubscriptionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumchangeSubscriptionSchemaUrn: valid values are %v", v, AllowedEnumchangeSubscriptionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumchangeSubscriptionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumchangeSubscriptionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumchange-subscriptionSchemaUrn value
func (v EnumchangeSubscriptionSchemaUrn) Ptr() *EnumchangeSubscriptionSchemaUrn {
	return &v
}

type NullableEnumchangeSubscriptionSchemaUrn struct {
	value *EnumchangeSubscriptionSchemaUrn
	isSet bool
}

func (v NullableEnumchangeSubscriptionSchemaUrn) Get() *EnumchangeSubscriptionSchemaUrn {
	return v.value
}

func (v *NullableEnumchangeSubscriptionSchemaUrn) Set(val *EnumchangeSubscriptionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumchangeSubscriptionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumchangeSubscriptionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumchangeSubscriptionSchemaUrn(val *EnumchangeSubscriptionSchemaUrn) *NullableEnumchangeSubscriptionSchemaUrn {
	return &NullableEnumchangeSubscriptionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumchangeSubscriptionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumchangeSubscriptionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
