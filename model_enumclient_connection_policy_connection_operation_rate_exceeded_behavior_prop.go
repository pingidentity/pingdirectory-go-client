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

// EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp Specifies the behavior that the Directory Server should exhibit if a client connection attempts to exceed a rate defined in the maximum-connection-operation-rate property. If the configured behavior is one that will reject requested operations, then that behavior will persist until the end of the corresponding interval. The server will resume allowing that client to perform operations when that interval expires, as long as no other operation rate limits have been exceeded.
type EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp string

// List of Enumclient-connection-policy-connectionOperationRateExceededBehaviorProp
const (
	ENUMCLIENTCONNECTIONPOLICYCONNECTIONOPERATIONRATEEXCEEDEDBEHAVIORPROP_DISCONNECT EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp = "disconnect"
	ENUMCLIENTCONNECTIONPOLICYCONNECTIONOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_ADMIN_LIMIT_EXCEEDED EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp = "reject-admin-limit-exceeded"
	ENUMCLIENTCONNECTIONPOLICYCONNECTIONOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_CONSTRAINT_VIOLATION EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp = "reject-constraint-violation"
	ENUMCLIENTCONNECTIONPOLICYCONNECTIONOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_BUSY EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp = "reject-busy"
	ENUMCLIENTCONNECTIONPOLICYCONNECTIONOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_UNAVAILABLE EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp = "reject-unavailable"
	ENUMCLIENTCONNECTIONPOLICYCONNECTIONOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_UNWILLING_TO_PERFORM EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp = "reject-unwilling-to-perform"
	ENUMCLIENTCONNECTIONPOLICYCONNECTIONOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_OTHER EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp = "reject-other"
)

// All allowed values of EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp enum
var AllowedEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorPropEnumValues = []EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp{
	"disconnect",
	"reject-admin-limit-exceeded",
	"reject-constraint-violation",
	"reject-busy",
	"reject-unavailable",
	"reject-unwilling-to-perform",
	"reject-other",
}

func (v *EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp(value)
	for _, existing := range AllowedEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp", value)
}

// NewEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorPropFromValue returns a pointer to a valid EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorPropFromValue(v string) (*EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp, error) {
	ev := EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp: valid values are %v", v, AllowedEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumclient-connection-policy-connectionOperationRateExceededBehaviorProp value
func (v EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp) Ptr() *EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp {
	return &v
}

type NullableEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp struct {
	value *EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp
	isSet bool
}

func (v NullableEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp) Get() *EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp {
	return v.value
}

func (v *NullableEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp) Set(val *EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp(val *EnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp) *NullableEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp {
	return &NullableEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumclientConnectionPolicyConnectionOperationRateExceededBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
