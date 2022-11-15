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

// EnumresultCriteriaReferralReturnedProp Indicates whether operation results which include one or more referral URLs should be included in this Simple Result Criteria. If no value is provided, then whether an operation includes any referral URLs will not be considered when determining whether it matches this Simple Result Criteria.
type EnumresultCriteriaReferralReturnedProp string

// List of Enumresult-criteria-referralReturnedProp
const (
	REQUIRED EnumresultCriteriaReferralReturnedProp = "required"
	PROHIBITED EnumresultCriteriaReferralReturnedProp = "prohibited"
	OPTIONAL EnumresultCriteriaReferralReturnedProp = "optional"
)

// All allowed values of EnumresultCriteriaReferralReturnedProp enum
var AllowedEnumresultCriteriaReferralReturnedPropEnumValues = []EnumresultCriteriaReferralReturnedProp{
	"required",
	"prohibited",
	"optional",
}

func (v *EnumresultCriteriaReferralReturnedProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumresultCriteriaReferralReturnedProp(value)
	for _, existing := range AllowedEnumresultCriteriaReferralReturnedPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumresultCriteriaReferralReturnedProp", value)
}

// NewEnumresultCriteriaReferralReturnedPropFromValue returns a pointer to a valid EnumresultCriteriaReferralReturnedProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumresultCriteriaReferralReturnedPropFromValue(v string) (*EnumresultCriteriaReferralReturnedProp, error) {
	ev := EnumresultCriteriaReferralReturnedProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumresultCriteriaReferralReturnedProp: valid values are %v", v, AllowedEnumresultCriteriaReferralReturnedPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumresultCriteriaReferralReturnedProp) IsValid() bool {
	for _, existing := range AllowedEnumresultCriteriaReferralReturnedPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumresult-criteria-referralReturnedProp value
func (v EnumresultCriteriaReferralReturnedProp) Ptr() *EnumresultCriteriaReferralReturnedProp {
	return &v
}

type NullableEnumresultCriteriaReferralReturnedProp struct {
	value *EnumresultCriteriaReferralReturnedProp
	isSet bool
}

func (v NullableEnumresultCriteriaReferralReturnedProp) Get() *EnumresultCriteriaReferralReturnedProp {
	return v.value
}

func (v *NullableEnumresultCriteriaReferralReturnedProp) Set(val *EnumresultCriteriaReferralReturnedProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumresultCriteriaReferralReturnedProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumresultCriteriaReferralReturnedProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumresultCriteriaReferralReturnedProp(val *EnumresultCriteriaReferralReturnedProp) *NullableEnumresultCriteriaReferralReturnedProp {
	return &NullableEnumresultCriteriaReferralReturnedProp{value: val, isSet: true}
}

func (v NullableEnumresultCriteriaReferralReturnedProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumresultCriteriaReferralReturnedProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

