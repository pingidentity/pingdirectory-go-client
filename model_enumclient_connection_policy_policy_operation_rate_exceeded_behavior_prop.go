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

// EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp Specifies the behavior that the Directory Server should exhibit if a client connection attempts to exceed a rate defined in the maximum-policy-operation-rate property. If the configured behavior is one that will reject requested operations, then that behavior will persist until the end of the corresponding interval. The server will resume allowing clients associated with this Client Connection Policy to perform operations when that interval expires, as long as no other operation rate limits have been exceeded.
type EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp string

// List of Enumclient-connection-policy-policyOperationRateExceededBehaviorProp
const (
	ENUMCLIENTCONNECTIONPOLICYPOLICYOPERATIONRATEEXCEEDEDBEHAVIORPROP_DISCONNECT                  EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp = "disconnect"
	ENUMCLIENTCONNECTIONPOLICYPOLICYOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_ADMIN_LIMIT_EXCEEDED EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp = "reject-admin-limit-exceeded"
	ENUMCLIENTCONNECTIONPOLICYPOLICYOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_CONSTRAINT_VIOLATION EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp = "reject-constraint-violation"
	ENUMCLIENTCONNECTIONPOLICYPOLICYOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_BUSY                 EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp = "reject-busy"
	ENUMCLIENTCONNECTIONPOLICYPOLICYOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_UNAVAILABLE          EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp = "reject-unavailable"
	ENUMCLIENTCONNECTIONPOLICYPOLICYOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_UNWILLING_TO_PERFORM EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp = "reject-unwilling-to-perform"
	ENUMCLIENTCONNECTIONPOLICYPOLICYOPERATIONRATEEXCEEDEDBEHAVIORPROP_REJECT_OTHER                EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp = "reject-other"
)

// All allowed values of EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp enum
var AllowedEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorPropEnumValues = []EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp{
	"disconnect",
	"reject-admin-limit-exceeded",
	"reject-constraint-violation",
	"reject-busy",
	"reject-unavailable",
	"reject-unwilling-to-perform",
	"reject-other",
}

func (v *EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp(value)
	for _, existing := range AllowedEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp", value)
}

// NewEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorPropFromValue returns a pointer to a valid EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorPropFromValue(v string) (*EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp, error) {
	ev := EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp: valid values are %v", v, AllowedEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp) IsValid() bool {
	for _, existing := range AllowedEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumclient-connection-policy-policyOperationRateExceededBehaviorProp value
func (v EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp) Ptr() *EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp {
	return &v
}

type NullableEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp struct {
	value *EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp
	isSet bool
}

func (v NullableEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp) Get() *EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp {
	return v.value
}

func (v *NullableEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp) Set(val *EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp(val *EnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp) *NullableEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp {
	return &NullableEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp{value: val, isSet: true}
}

func (v NullableEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumclientConnectionPolicyPolicyOperationRateExceededBehaviorProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
