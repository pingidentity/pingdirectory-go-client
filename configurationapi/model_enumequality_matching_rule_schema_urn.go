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

// EnumequalityMatchingRuleSchemaUrn the model 'EnumequalityMatchingRuleSchemaUrn'
type EnumequalityMatchingRuleSchemaUrn string

// List of Enumequality-matching-ruleSchemaUrn
const (
	ENUMEQUALITYMATCHINGRULESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0MATCHING_RULEEQUALITY EnumequalityMatchingRuleSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:matching-rule:equality"
)

// All allowed values of EnumequalityMatchingRuleSchemaUrn enum
var AllowedEnumequalityMatchingRuleSchemaUrnEnumValues = []EnumequalityMatchingRuleSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:matching-rule:equality",
}

func (v *EnumequalityMatchingRuleSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumequalityMatchingRuleSchemaUrn(value)
	for _, existing := range AllowedEnumequalityMatchingRuleSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumequalityMatchingRuleSchemaUrn", value)
}

// NewEnumequalityMatchingRuleSchemaUrnFromValue returns a pointer to a valid EnumequalityMatchingRuleSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumequalityMatchingRuleSchemaUrnFromValue(v string) (*EnumequalityMatchingRuleSchemaUrn, error) {
	ev := EnumequalityMatchingRuleSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumequalityMatchingRuleSchemaUrn: valid values are %v", v, AllowedEnumequalityMatchingRuleSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumequalityMatchingRuleSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumequalityMatchingRuleSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumequality-matching-ruleSchemaUrn value
func (v EnumequalityMatchingRuleSchemaUrn) Ptr() *EnumequalityMatchingRuleSchemaUrn {
	return &v
}

type NullableEnumequalityMatchingRuleSchemaUrn struct {
	value *EnumequalityMatchingRuleSchemaUrn
	isSet bool
}

func (v NullableEnumequalityMatchingRuleSchemaUrn) Get() *EnumequalityMatchingRuleSchemaUrn {
	return v.value
}

func (v *NullableEnumequalityMatchingRuleSchemaUrn) Set(val *EnumequalityMatchingRuleSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumequalityMatchingRuleSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumequalityMatchingRuleSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumequalityMatchingRuleSchemaUrn(val *EnumequalityMatchingRuleSchemaUrn) *NullableEnumequalityMatchingRuleSchemaUrn {
	return &NullableEnumequalityMatchingRuleSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumequalityMatchingRuleSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumequalityMatchingRuleSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
