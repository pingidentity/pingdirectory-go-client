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

// EnumorderingMatchingRuleSchemaUrn the model 'EnumorderingMatchingRuleSchemaUrn'
type EnumorderingMatchingRuleSchemaUrn string

// List of Enumordering-matching-ruleSchemaUrn
const (
	ENUMORDERINGMATCHINGRULESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0MATCHING_RULEORDERING EnumorderingMatchingRuleSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:matching-rule:ordering"
)

// All allowed values of EnumorderingMatchingRuleSchemaUrn enum
var AllowedEnumorderingMatchingRuleSchemaUrnEnumValues = []EnumorderingMatchingRuleSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:matching-rule:ordering",
}

func (v *EnumorderingMatchingRuleSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumorderingMatchingRuleSchemaUrn(value)
	for _, existing := range AllowedEnumorderingMatchingRuleSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumorderingMatchingRuleSchemaUrn", value)
}

// NewEnumorderingMatchingRuleSchemaUrnFromValue returns a pointer to a valid EnumorderingMatchingRuleSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumorderingMatchingRuleSchemaUrnFromValue(v string) (*EnumorderingMatchingRuleSchemaUrn, error) {
	ev := EnumorderingMatchingRuleSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumorderingMatchingRuleSchemaUrn: valid values are %v", v, AllowedEnumorderingMatchingRuleSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumorderingMatchingRuleSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumorderingMatchingRuleSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumordering-matching-ruleSchemaUrn value
func (v EnumorderingMatchingRuleSchemaUrn) Ptr() *EnumorderingMatchingRuleSchemaUrn {
	return &v
}

type NullableEnumorderingMatchingRuleSchemaUrn struct {
	value *EnumorderingMatchingRuleSchemaUrn
	isSet bool
}

func (v NullableEnumorderingMatchingRuleSchemaUrn) Get() *EnumorderingMatchingRuleSchemaUrn {
	return v.value
}

func (v *NullableEnumorderingMatchingRuleSchemaUrn) Set(val *EnumorderingMatchingRuleSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumorderingMatchingRuleSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumorderingMatchingRuleSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumorderingMatchingRuleSchemaUrn(val *EnumorderingMatchingRuleSchemaUrn) *NullableEnumorderingMatchingRuleSchemaUrn {
	return &NullableEnumorderingMatchingRuleSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumorderingMatchingRuleSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumorderingMatchingRuleSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
