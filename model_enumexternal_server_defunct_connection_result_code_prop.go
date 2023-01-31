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

// EnumexternalServerDefunctConnectionResultCodeProp Specifies the operation result code values that should cause the associated connection should be considered defunct. If an operation fails with one of these result codes, then it will be terminated and an attempt will be made to establish a new connection in its place.
type EnumexternalServerDefunctConnectionResultCodeProp string

// List of Enumexternal-server-defunctConnectionResultCodeProp
const (
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_SUCCESS                         EnumexternalServerDefunctConnectionResultCodeProp = "success"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_OPERATIONS_ERROR                EnumexternalServerDefunctConnectionResultCodeProp = "operations-error"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_PROTOCOL_ERROR                  EnumexternalServerDefunctConnectionResultCodeProp = "protocol-error"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_TIME_LIMIT_EXCEEDED             EnumexternalServerDefunctConnectionResultCodeProp = "time-limit-exceeded"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_SIZE_LIMIT_EXCEEDED             EnumexternalServerDefunctConnectionResultCodeProp = "size-limit-exceeded"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_COMPARE_FALSE                   EnumexternalServerDefunctConnectionResultCodeProp = "compare-false"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_COMPARE_TRUE                    EnumexternalServerDefunctConnectionResultCodeProp = "compare-true"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_AUTH_METHOD_NOT_SUPPORTED       EnumexternalServerDefunctConnectionResultCodeProp = "auth-method-not-supported"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_STRONG_AUTH_REQUIRED            EnumexternalServerDefunctConnectionResultCodeProp = "strong-auth-required"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_REFERRAL                        EnumexternalServerDefunctConnectionResultCodeProp = "referral"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_ADMIN_LIMIT_EXCEEDED            EnumexternalServerDefunctConnectionResultCodeProp = "admin-limit-exceeded"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_UNAVAILABLE_CRITICAL_EXTENSION  EnumexternalServerDefunctConnectionResultCodeProp = "unavailable-critical-extension"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_CONFIDENTIALITY_REQUIRED        EnumexternalServerDefunctConnectionResultCodeProp = "confidentiality-required"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_SASL_BIND_IN_PROGRESS           EnumexternalServerDefunctConnectionResultCodeProp = "sasl-bind-in-progress"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_NO_SUCH_ATTRIBUTE               EnumexternalServerDefunctConnectionResultCodeProp = "no-such-attribute"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_UNDEFINED_ATTRIBUTE_TYPE        EnumexternalServerDefunctConnectionResultCodeProp = "undefined-attribute-type"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_INAPPROPRIATE_MATCHING          EnumexternalServerDefunctConnectionResultCodeProp = "inappropriate-matching"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_CONSTRAINT_VIOLATION            EnumexternalServerDefunctConnectionResultCodeProp = "constraint-violation"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_ATTRIBUTE_OR_VALUE_EXISTS       EnumexternalServerDefunctConnectionResultCodeProp = "attribute-or-value-exists"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_INVALID_ATTRIBUTE_SYNTAX        EnumexternalServerDefunctConnectionResultCodeProp = "invalid-attribute-syntax"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_NO_SUCH_OBJECT                  EnumexternalServerDefunctConnectionResultCodeProp = "no-such-object"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_ALIAS_PROBLEM                   EnumexternalServerDefunctConnectionResultCodeProp = "alias-problem"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_INVALID_DN_SYNTAX               EnumexternalServerDefunctConnectionResultCodeProp = "invalid-dn-syntax"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_ALIAS_DEREFERENCING_PROBLEM     EnumexternalServerDefunctConnectionResultCodeProp = "alias-dereferencing-problem"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_INAPPROPRIATE_AUTHENTICATION    EnumexternalServerDefunctConnectionResultCodeProp = "inappropriate-authentication"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_INVALID_CREDENTIALS             EnumexternalServerDefunctConnectionResultCodeProp = "invalid-credentials"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_INSUFFICIENT_ACCESS_RIGHTS      EnumexternalServerDefunctConnectionResultCodeProp = "insufficient-access-rights"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_BUSY                            EnumexternalServerDefunctConnectionResultCodeProp = "busy"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_UNAVAILABLE                     EnumexternalServerDefunctConnectionResultCodeProp = "unavailable"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_UNWILLING_TO_PERFORM            EnumexternalServerDefunctConnectionResultCodeProp = "unwilling-to-perform"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_LOOP_DETECT                     EnumexternalServerDefunctConnectionResultCodeProp = "loop-detect"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_SORT_CONTROL_MISSING            EnumexternalServerDefunctConnectionResultCodeProp = "sort-control-missing"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_OFFSET_RANGE_ERROR              EnumexternalServerDefunctConnectionResultCodeProp = "offset-range-error"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_NAMING_VIOLATION                EnumexternalServerDefunctConnectionResultCodeProp = "naming-violation"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_OBJECT_CLASS_VIOLATION          EnumexternalServerDefunctConnectionResultCodeProp = "object-class-violation"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_NOT_ALLOWED_ON_NONLEAF          EnumexternalServerDefunctConnectionResultCodeProp = "not-allowed-on-nonleaf"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_NOT_ALLOWED_ON_RDN              EnumexternalServerDefunctConnectionResultCodeProp = "not-allowed-on-rdn"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_ENTRY_ALREADY_EXISTS            EnumexternalServerDefunctConnectionResultCodeProp = "entry-already-exists"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_OBJECT_CLASS_MODS_PROHIBITED    EnumexternalServerDefunctConnectionResultCodeProp = "object-class-mods-prohibited"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_AFFECTS_MULTIPLE_DSAS           EnumexternalServerDefunctConnectionResultCodeProp = "affects-multiple-dsas"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_VIRTUAL_LIST_VIEW_ERROR         EnumexternalServerDefunctConnectionResultCodeProp = "virtual-list-view-error"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_OTHER                           EnumexternalServerDefunctConnectionResultCodeProp = "other"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_SERVER_DOWN                     EnumexternalServerDefunctConnectionResultCodeProp = "server-down"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_LOCAL_ERROR                     EnumexternalServerDefunctConnectionResultCodeProp = "local-error"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_ENCODING_ERROR                  EnumexternalServerDefunctConnectionResultCodeProp = "encoding-error"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_DECODING_ERROR                  EnumexternalServerDefunctConnectionResultCodeProp = "decoding-error"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_TIMEOUT                         EnumexternalServerDefunctConnectionResultCodeProp = "timeout"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_AUTH_UNKNOWN                    EnumexternalServerDefunctConnectionResultCodeProp = "auth-unknown"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_FILTER_ERROR                    EnumexternalServerDefunctConnectionResultCodeProp = "filter-error"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_USER_CANCELED                   EnumexternalServerDefunctConnectionResultCodeProp = "user-canceled"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_PARAM_ERROR                     EnumexternalServerDefunctConnectionResultCodeProp = "param-error"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_NO_MEMORY                       EnumexternalServerDefunctConnectionResultCodeProp = "no-memory"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_CONNECT_ERROR                   EnumexternalServerDefunctConnectionResultCodeProp = "connect-error"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_NOT_SUPPORTED                   EnumexternalServerDefunctConnectionResultCodeProp = "not-supported"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_CONTROL_NOT_FOUND               EnumexternalServerDefunctConnectionResultCodeProp = "control-not-found"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_NO_RESULTS_RETURNED             EnumexternalServerDefunctConnectionResultCodeProp = "no-results-returned"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_MORE_RESULTS_TO_RETURN          EnumexternalServerDefunctConnectionResultCodeProp = "more-results-to-return"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_CLIENT_LOOP                     EnumexternalServerDefunctConnectionResultCodeProp = "client-loop"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_REFERRAL_LIMIT_EXCEEDED         EnumexternalServerDefunctConnectionResultCodeProp = "referral-limit-exceeded"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_CANCELED                        EnumexternalServerDefunctConnectionResultCodeProp = "canceled"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_NO_SUCH_OPERATION               EnumexternalServerDefunctConnectionResultCodeProp = "no-such-operation"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_TOO_LATE                        EnumexternalServerDefunctConnectionResultCodeProp = "too-late"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_CANNOT_CANCEL                   EnumexternalServerDefunctConnectionResultCodeProp = "cannot-cancel"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_ASSERTION_FAILED                EnumexternalServerDefunctConnectionResultCodeProp = "assertion-failed"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_AUTHORIZATION_DENIED            EnumexternalServerDefunctConnectionResultCodeProp = "authorization-denied"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_NO_OPERATION                    EnumexternalServerDefunctConnectionResultCodeProp = "no-operation"
	ENUMEXTERNALSERVERDEFUNCTCONNECTIONRESULTCODEPROP_INTERACTIVE_TRANSACTION_ABORTED EnumexternalServerDefunctConnectionResultCodeProp = "interactive-transaction-aborted"
)

// All allowed values of EnumexternalServerDefunctConnectionResultCodeProp enum
var AllowedEnumexternalServerDefunctConnectionResultCodePropEnumValues = []EnumexternalServerDefunctConnectionResultCodeProp{
	"success",
	"operations-error",
	"protocol-error",
	"time-limit-exceeded",
	"size-limit-exceeded",
	"compare-false",
	"compare-true",
	"auth-method-not-supported",
	"strong-auth-required",
	"referral",
	"admin-limit-exceeded",
	"unavailable-critical-extension",
	"confidentiality-required",
	"sasl-bind-in-progress",
	"no-such-attribute",
	"undefined-attribute-type",
	"inappropriate-matching",
	"constraint-violation",
	"attribute-or-value-exists",
	"invalid-attribute-syntax",
	"no-such-object",
	"alias-problem",
	"invalid-dn-syntax",
	"alias-dereferencing-problem",
	"inappropriate-authentication",
	"invalid-credentials",
	"insufficient-access-rights",
	"busy",
	"unavailable",
	"unwilling-to-perform",
	"loop-detect",
	"sort-control-missing",
	"offset-range-error",
	"naming-violation",
	"object-class-violation",
	"not-allowed-on-nonleaf",
	"not-allowed-on-rdn",
	"entry-already-exists",
	"object-class-mods-prohibited",
	"affects-multiple-dsas",
	"virtual-list-view-error",
	"other",
	"server-down",
	"local-error",
	"encoding-error",
	"decoding-error",
	"timeout",
	"auth-unknown",
	"filter-error",
	"user-canceled",
	"param-error",
	"no-memory",
	"connect-error",
	"not-supported",
	"control-not-found",
	"no-results-returned",
	"more-results-to-return",
	"client-loop",
	"referral-limit-exceeded",
	"canceled",
	"no-such-operation",
	"too-late",
	"cannot-cancel",
	"assertion-failed",
	"authorization-denied",
	"no-operation",
	"interactive-transaction-aborted",
}

func (v *EnumexternalServerDefunctConnectionResultCodeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerDefunctConnectionResultCodeProp(value)
	for _, existing := range AllowedEnumexternalServerDefunctConnectionResultCodePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerDefunctConnectionResultCodeProp", value)
}

// NewEnumexternalServerDefunctConnectionResultCodePropFromValue returns a pointer to a valid EnumexternalServerDefunctConnectionResultCodeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerDefunctConnectionResultCodePropFromValue(v string) (*EnumexternalServerDefunctConnectionResultCodeProp, error) {
	ev := EnumexternalServerDefunctConnectionResultCodeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerDefunctConnectionResultCodeProp: valid values are %v", v, AllowedEnumexternalServerDefunctConnectionResultCodePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerDefunctConnectionResultCodeProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerDefunctConnectionResultCodePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-defunctConnectionResultCodeProp value
func (v EnumexternalServerDefunctConnectionResultCodeProp) Ptr() *EnumexternalServerDefunctConnectionResultCodeProp {
	return &v
}

type NullableEnumexternalServerDefunctConnectionResultCodeProp struct {
	value *EnumexternalServerDefunctConnectionResultCodeProp
	isSet bool
}

func (v NullableEnumexternalServerDefunctConnectionResultCodeProp) Get() *EnumexternalServerDefunctConnectionResultCodeProp {
	return v.value
}

func (v *NullableEnumexternalServerDefunctConnectionResultCodeProp) Set(val *EnumexternalServerDefunctConnectionResultCodeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerDefunctConnectionResultCodeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerDefunctConnectionResultCodeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerDefunctConnectionResultCodeProp(val *EnumexternalServerDefunctConnectionResultCodeProp) *NullableEnumexternalServerDefunctConnectionResultCodeProp {
	return &NullableEnumexternalServerDefunctConnectionResultCodeProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerDefunctConnectionResultCodeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerDefunctConnectionResultCodeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
