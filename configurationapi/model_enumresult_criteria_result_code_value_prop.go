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

// EnumresultCriteriaResultCodeValueProp Specifies the operation result code values for results included in this Simple Result Criteria. This will only be taken into account if the \"result-code-criteria\" property has a value of \"selected-result-codes\".
type EnumresultCriteriaResultCodeValueProp string

// List of Enumresult-criteria-resultCodeValueProp
const (
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_SUCCESS                              EnumresultCriteriaResultCodeValueProp = "success"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_OPERATIONS_ERROR                     EnumresultCriteriaResultCodeValueProp = "operations-error"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_PROTOCOL_ERROR                       EnumresultCriteriaResultCodeValueProp = "protocol-error"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_TIME_LIMIT_EXCEEDED                  EnumresultCriteriaResultCodeValueProp = "time-limit-exceeded"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_SIZE_LIMIT_EXCEEDED                  EnumresultCriteriaResultCodeValueProp = "size-limit-exceeded"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_COMPARE_FALSE                        EnumresultCriteriaResultCodeValueProp = "compare-false"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_COMPARE_TRUE                         EnumresultCriteriaResultCodeValueProp = "compare-true"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_AUTH_METHOD_NOT_SUPPORTED            EnumresultCriteriaResultCodeValueProp = "auth-method-not-supported"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_STRONG_AUTH_REQUIRED                 EnumresultCriteriaResultCodeValueProp = "strong-auth-required"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_REFERRAL                             EnumresultCriteriaResultCodeValueProp = "referral"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_ADMIN_LIMIT_EXCEEDED                 EnumresultCriteriaResultCodeValueProp = "admin-limit-exceeded"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_UNAVAILABLE_CRITICAL_EXTENSION       EnumresultCriteriaResultCodeValueProp = "unavailable-critical-extension"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_CONFIDENTIALITY_REQUIRED             EnumresultCriteriaResultCodeValueProp = "confidentiality-required"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_SASL_BIND_IN_PROGRESS                EnumresultCriteriaResultCodeValueProp = "sasl-bind-in-progress"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_NO_SUCH_ATTRIBUTE                    EnumresultCriteriaResultCodeValueProp = "no-such-attribute"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_UNDEFINED_ATTRIBUTE_TYPE             EnumresultCriteriaResultCodeValueProp = "undefined-attribute-type"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_INAPPROPRIATE_MATCHING               EnumresultCriteriaResultCodeValueProp = "inappropriate-matching"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_CONSTRAINT_VIOLATION                 EnumresultCriteriaResultCodeValueProp = "constraint-violation"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_ATTRIBUTE_OR_VALUE_EXISTS            EnumresultCriteriaResultCodeValueProp = "attribute-or-value-exists"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_INVALID_ATTRIBUTE_SYNTAX             EnumresultCriteriaResultCodeValueProp = "invalid-attribute-syntax"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_NO_SUCH_OBJECT                       EnumresultCriteriaResultCodeValueProp = "no-such-object"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_ALIAS_PROBLEM                        EnumresultCriteriaResultCodeValueProp = "alias-problem"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_INVALID_DN_SYNTAX                    EnumresultCriteriaResultCodeValueProp = "invalid-dn-syntax"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_ALIAS_DEREFERENCING_PROBLEM          EnumresultCriteriaResultCodeValueProp = "alias-dereferencing-problem"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_INAPPROPRIATE_AUTHENTICATION         EnumresultCriteriaResultCodeValueProp = "inappropriate-authentication"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_INVALID_CREDENTIALS                  EnumresultCriteriaResultCodeValueProp = "invalid-credentials"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_INSUFFICIENT_ACCESS_RIGHTS           EnumresultCriteriaResultCodeValueProp = "insufficient-access-rights"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_BUSY                                 EnumresultCriteriaResultCodeValueProp = "busy"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_UNAVAILABLE                          EnumresultCriteriaResultCodeValueProp = "unavailable"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_UNWILLING_TO_PERFORM                 EnumresultCriteriaResultCodeValueProp = "unwilling-to-perform"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_LOOP_DETECT                          EnumresultCriteriaResultCodeValueProp = "loop-detect"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_SORT_CONTROL_MISSING                 EnumresultCriteriaResultCodeValueProp = "sort-control-missing"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_OFFSET_RANGE_ERROR                   EnumresultCriteriaResultCodeValueProp = "offset-range-error"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_NAMING_VIOLATION                     EnumresultCriteriaResultCodeValueProp = "naming-violation"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_OBJECTCLASS_VIOLATION                EnumresultCriteriaResultCodeValueProp = "objectclass-violation"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_NOT_ALLOWED_ON_NONLEAF               EnumresultCriteriaResultCodeValueProp = "not-allowed-on-nonleaf"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_NOT_ALLOWED_ON_RDN                   EnumresultCriteriaResultCodeValueProp = "not-allowed-on-rdn"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_ENTRY_ALREADY_EXISTS                 EnumresultCriteriaResultCodeValueProp = "entry-already-exists"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_OBJECTCLASS_MODS_PROHIBITED          EnumresultCriteriaResultCodeValueProp = "objectclass-mods-prohibited"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_AFFECTS_MULTIPLE_DSAS                EnumresultCriteriaResultCodeValueProp = "affects-multiple-dsas"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_VIRTUAL_LIST_VIEW_ERROR              EnumresultCriteriaResultCodeValueProp = "virtual-list-view-error"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_OTHER                                EnumresultCriteriaResultCodeValueProp = "other"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_CANCELED                             EnumresultCriteriaResultCodeValueProp = "canceled"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_NO_SUCH_OPERATION                    EnumresultCriteriaResultCodeValueProp = "no-such-operation"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_TOO_LATE                             EnumresultCriteriaResultCodeValueProp = "too-late"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_CANNOT_CANCEL                        EnumresultCriteriaResultCodeValueProp = "cannot-cancel"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_ASSERTION_FAILED                     EnumresultCriteriaResultCodeValueProp = "assertion-failed"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_AUTHORIZATION_DENIED                 EnumresultCriteriaResultCodeValueProp = "authorization-denied"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_NO_OPERATION                         EnumresultCriteriaResultCodeValueProp = "no-operation"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_INTERACTIVE_TRANSACTION_ABORTED      EnumresultCriteriaResultCodeValueProp = "interactive-transaction-aborted"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_DATABASE_LOCK_CONFLICT               EnumresultCriteriaResultCodeValueProp = "database-lock-conflict"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_MIRRORED_SUBTREE_DIGEST_MISMATCH     EnumresultCriteriaResultCodeValueProp = "mirrored-subtree-digest-mismatch"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_TOKEN_DELIVERY_MECHANISM_UNAVAILABLE EnumresultCriteriaResultCodeValueProp = "token-delivery-mechanism-unavailable"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_TOKEN_DELIVERY_ATTEMPT_FAILED        EnumresultCriteriaResultCodeValueProp = "token-delivery-attempt-failed"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_TOKEN_DELIVERY_INVALID_RECIPIENT_ID  EnumresultCriteriaResultCodeValueProp = "token-delivery-invalid-recipient-id"
	ENUMRESULTCRITERIARESULTCODEVALUEPROP_TOKEN_DELIVERY_INVALID_ACCOUNT_STATE EnumresultCriteriaResultCodeValueProp = "token-delivery-invalid-account-state"
)

// All allowed values of EnumresultCriteriaResultCodeValueProp enum
var AllowedEnumresultCriteriaResultCodeValuePropEnumValues = []EnumresultCriteriaResultCodeValueProp{
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
	"objectclass-violation",
	"not-allowed-on-nonleaf",
	"not-allowed-on-rdn",
	"entry-already-exists",
	"objectclass-mods-prohibited",
	"affects-multiple-dsas",
	"virtual-list-view-error",
	"other",
	"canceled",
	"no-such-operation",
	"too-late",
	"cannot-cancel",
	"assertion-failed",
	"authorization-denied",
	"no-operation",
	"interactive-transaction-aborted",
	"database-lock-conflict",
	"mirrored-subtree-digest-mismatch",
	"token-delivery-mechanism-unavailable",
	"token-delivery-attempt-failed",
	"token-delivery-invalid-recipient-id",
	"token-delivery-invalid-account-state",
}

func (v *EnumresultCriteriaResultCodeValueProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumresultCriteriaResultCodeValueProp(value)
	for _, existing := range AllowedEnumresultCriteriaResultCodeValuePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumresultCriteriaResultCodeValueProp", value)
}

// NewEnumresultCriteriaResultCodeValuePropFromValue returns a pointer to a valid EnumresultCriteriaResultCodeValueProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumresultCriteriaResultCodeValuePropFromValue(v string) (*EnumresultCriteriaResultCodeValueProp, error) {
	ev := EnumresultCriteriaResultCodeValueProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumresultCriteriaResultCodeValueProp: valid values are %v", v, AllowedEnumresultCriteriaResultCodeValuePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumresultCriteriaResultCodeValueProp) IsValid() bool {
	for _, existing := range AllowedEnumresultCriteriaResultCodeValuePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumresult-criteria-resultCodeValueProp value
func (v EnumresultCriteriaResultCodeValueProp) Ptr() *EnumresultCriteriaResultCodeValueProp {
	return &v
}

type NullableEnumresultCriteriaResultCodeValueProp struct {
	value *EnumresultCriteriaResultCodeValueProp
	isSet bool
}

func (v NullableEnumresultCriteriaResultCodeValueProp) Get() *EnumresultCriteriaResultCodeValueProp {
	return v.value
}

func (v *NullableEnumresultCriteriaResultCodeValueProp) Set(val *EnumresultCriteriaResultCodeValueProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumresultCriteriaResultCodeValueProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumresultCriteriaResultCodeValueProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumresultCriteriaResultCodeValueProp(val *EnumresultCriteriaResultCodeValueProp) *NullableEnumresultCriteriaResultCodeValueProp {
	return &NullableEnumresultCriteriaResultCodeValueProp{value: val, isSet: true}
}

func (v NullableEnumresultCriteriaResultCodeValueProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumresultCriteriaResultCodeValueProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
