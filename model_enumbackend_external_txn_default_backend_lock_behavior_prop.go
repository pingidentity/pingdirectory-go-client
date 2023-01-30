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

// EnumbackendExternalTxnDefaultBackendLockBehaviorProp Specifies the default behavior that should be exhibited by external transactions (e.g., an LDAP transaction or an atomic multi-update operation) with regard to acquiring an exclusive lock in this backend.
type EnumbackendExternalTxnDefaultBackendLockBehaviorProp string

// List of Enumbackend-externalTxnDefaultBackendLockBehaviorProp
const (
	ENUMBACKENDEXTERNALTXNDEFAULTBACKENDLOCKBEHAVIORPROP_DO_NOT_ACQUIRE                 EnumbackendExternalTxnDefaultBackendLockBehaviorProp = "do-not-acquire"
	ENUMBACKENDEXTERNALTXNDEFAULTBACKENDLOCKBEHAVIORPROP_ACQUIRE_AFTER_RETRIES          EnumbackendExternalTxnDefaultBackendLockBehaviorProp = "acquire-after-retries"
	ENUMBACKENDEXTERNALTXNDEFAULTBACKENDLOCKBEHAVIORPROP_ACQUIRE_BEFORE_RETRIES         EnumbackendExternalTxnDefaultBackendLockBehaviorProp = "acquire-before-retries"
	ENUMBACKENDEXTERNALTXNDEFAULTBACKENDLOCKBEHAVIORPROP_ACQUIRE_BEFORE_INITIAL_ATTEMPT EnumbackendExternalTxnDefaultBackendLockBehaviorProp = "acquire-before-initial-attempt"
)

// All allowed values of EnumbackendExternalTxnDefaultBackendLockBehaviorProp enum
var AllowedEnumbackendExternalTxnDefaultBackendLockBehaviorPropEnumValues = []EnumbackendExternalTxnDefaultBackendLockBehaviorProp{
	"do-not-acquire",
	"acquire-after-retries",
	"acquire-before-retries",
	"acquire-before-initial-attempt",
}

func (v *EnumbackendExternalTxnDefaultBackendLockBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumbackendExternalTxnDefaultBackendLockBehaviorProp(value)
	for _, existing := range AllowedEnumbackendExternalTxnDefaultBackendLockBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumbackendExternalTxnDefaultBackendLockBehaviorProp", value)
}

// NewEnumbackendExternalTxnDefaultBackendLockBehaviorPropFromValue returns a pointer to a valid EnumbackendExternalTxnDefaultBackendLockBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumbackendExternalTxnDefaultBackendLockBehaviorPropFromValue(v string) (*EnumbackendExternalTxnDefaultBackendLockBehaviorProp, error) {
	ev := EnumbackendExternalTxnDefaultBackendLockBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumbackendExternalTxnDefaultBackendLockBehaviorProp: valid values are %v", v, AllowedEnumbackendExternalTxnDefaultBackendLockBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumbackendExternalTxnDefaultBackendLockBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumbackendExternalTxnDefaultBackendLockBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumbackend-externalTxnDefaultBackendLockBehaviorProp value
func (v EnumbackendExternalTxnDefaultBackendLockBehaviorProp) Ptr() *EnumbackendExternalTxnDefaultBackendLockBehaviorProp {
	return &v
}

type NullableEnumbackendExternalTxnDefaultBackendLockBehaviorProp struct {
	value *EnumbackendExternalTxnDefaultBackendLockBehaviorProp
	isSet bool
}

func (v NullableEnumbackendExternalTxnDefaultBackendLockBehaviorProp) Get() *EnumbackendExternalTxnDefaultBackendLockBehaviorProp {
	return v.value
}

func (v *NullableEnumbackendExternalTxnDefaultBackendLockBehaviorProp) Set(val *EnumbackendExternalTxnDefaultBackendLockBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumbackendExternalTxnDefaultBackendLockBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumbackendExternalTxnDefaultBackendLockBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumbackendExternalTxnDefaultBackendLockBehaviorProp(val *EnumbackendExternalTxnDefaultBackendLockBehaviorProp) *NullableEnumbackendExternalTxnDefaultBackendLockBehaviorProp {
	return &NullableEnumbackendExternalTxnDefaultBackendLockBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumbackendExternalTxnDefaultBackendLockBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumbackendExternalTxnDefaultBackendLockBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
