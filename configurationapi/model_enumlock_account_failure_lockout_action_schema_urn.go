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

// EnumlockAccountFailureLockoutActionSchemaUrn the model 'EnumlockAccountFailureLockoutActionSchemaUrn'
type EnumlockAccountFailureLockoutActionSchemaUrn string

// List of Enumlock-account-failure-lockout-actionSchemaUrn
const (
	ENUMLOCKACCOUNTFAILURELOCKOUTACTIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0FAILURE_LOCKOUT_ACTIONLOCK_ACCOUNT EnumlockAccountFailureLockoutActionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:failure-lockout-action:lock-account"
)

// All allowed values of EnumlockAccountFailureLockoutActionSchemaUrn enum
var AllowedEnumlockAccountFailureLockoutActionSchemaUrnEnumValues = []EnumlockAccountFailureLockoutActionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:failure-lockout-action:lock-account",
}

func (v *EnumlockAccountFailureLockoutActionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlockAccountFailureLockoutActionSchemaUrn(value)
	for _, existing := range AllowedEnumlockAccountFailureLockoutActionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlockAccountFailureLockoutActionSchemaUrn", value)
}

// NewEnumlockAccountFailureLockoutActionSchemaUrnFromValue returns a pointer to a valid EnumlockAccountFailureLockoutActionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlockAccountFailureLockoutActionSchemaUrnFromValue(v string) (*EnumlockAccountFailureLockoutActionSchemaUrn, error) {
	ev := EnumlockAccountFailureLockoutActionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlockAccountFailureLockoutActionSchemaUrn: valid values are %v", v, AllowedEnumlockAccountFailureLockoutActionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlockAccountFailureLockoutActionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumlockAccountFailureLockoutActionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlock-account-failure-lockout-actionSchemaUrn value
func (v EnumlockAccountFailureLockoutActionSchemaUrn) Ptr() *EnumlockAccountFailureLockoutActionSchemaUrn {
	return &v
}

type NullableEnumlockAccountFailureLockoutActionSchemaUrn struct {
	value *EnumlockAccountFailureLockoutActionSchemaUrn
	isSet bool
}

func (v NullableEnumlockAccountFailureLockoutActionSchemaUrn) Get() *EnumlockAccountFailureLockoutActionSchemaUrn {
	return v.value
}

func (v *NullableEnumlockAccountFailureLockoutActionSchemaUrn) Set(val *EnumlockAccountFailureLockoutActionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlockAccountFailureLockoutActionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlockAccountFailureLockoutActionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlockAccountFailureLockoutActionSchemaUrn(val *EnumlockAccountFailureLockoutActionSchemaUrn) *NullableEnumlockAccountFailureLockoutActionSchemaUrn {
	return &NullableEnumlockAccountFailureLockoutActionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumlockAccountFailureLockoutActionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlockAccountFailureLockoutActionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
