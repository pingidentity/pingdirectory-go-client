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

// EnumresultCriteriaUsedPrivilegeProp Specifies the name of a privilege that must have been used during the processing for operations included in this Simple Result Criteria. If any privilege names are provided, then the associated operation must have used at least one of those privileges. If no privilege names were provided, then the set of privileges used will not be considered when determining whether an operation should be included in this Simple Result Criteria.
type EnumresultCriteriaUsedPrivilegeProp string

// List of Enumresult-criteria-usedPrivilegeProp
const (
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_AUDIT_DATA_SECURITY EnumresultCriteriaUsedPrivilegeProp = "audit-data-security"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_BYPASS_ACL EnumresultCriteriaUsedPrivilegeProp = "bypass-acl"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_BYPASS_READ_ACL EnumresultCriteriaUsedPrivilegeProp = "bypass-read-acl"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_MODIFY_ACL EnumresultCriteriaUsedPrivilegeProp = "modify-acl"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_CONFIG_READ EnumresultCriteriaUsedPrivilegeProp = "config-read"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_CONFIG_WRITE EnumresultCriteriaUsedPrivilegeProp = "config-write"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_JMX_READ EnumresultCriteriaUsedPrivilegeProp = "jmx-read"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_JMX_WRITE EnumresultCriteriaUsedPrivilegeProp = "jmx-write"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_JMX_NOTIFY EnumresultCriteriaUsedPrivilegeProp = "jmx-notify"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_LDIF_IMPORT EnumresultCriteriaUsedPrivilegeProp = "ldif-import"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_LDIF_EXPORT EnumresultCriteriaUsedPrivilegeProp = "ldif-export"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_BACKEND_BACKUP EnumresultCriteriaUsedPrivilegeProp = "backend-backup"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_BACKEND_RESTORE EnumresultCriteriaUsedPrivilegeProp = "backend-restore"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_SERVER_SHUTDOWN EnumresultCriteriaUsedPrivilegeProp = "server-shutdown"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_SERVER_RESTART EnumresultCriteriaUsedPrivilegeProp = "server-restart"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_PROXIED_AUTH EnumresultCriteriaUsedPrivilegeProp = "proxied-auth"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_DISCONNECT_CLIENT EnumresultCriteriaUsedPrivilegeProp = "disconnect-client"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_PASSWORD_RESET EnumresultCriteriaUsedPrivilegeProp = "password-reset"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_UPDATE_SCHEMA EnumresultCriteriaUsedPrivilegeProp = "update-schema"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_PRIVILEGE_CHANGE EnumresultCriteriaUsedPrivilegeProp = "privilege-change"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_UNINDEXED_SEARCH EnumresultCriteriaUsedPrivilegeProp = "unindexed-search"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_UNINDEXED_SEARCH_WITH_CONTROL EnumresultCriteriaUsedPrivilegeProp = "unindexed-search-with-control"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_BYPASS_PW_POLICY EnumresultCriteriaUsedPrivilegeProp = "bypass-pw-policy"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_LOCKDOWN_MODE EnumresultCriteriaUsedPrivilegeProp = "lockdown-mode"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_STREAM_VALUES EnumresultCriteriaUsedPrivilegeProp = "stream-values"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_THIRD_PARTY_TASK EnumresultCriteriaUsedPrivilegeProp = "third-party-task"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_USE_ADMIN_SESSION EnumresultCriteriaUsedPrivilegeProp = "use-admin-session"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_SOFT_DELETE_READ EnumresultCriteriaUsedPrivilegeProp = "soft-delete-read"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_METRICS_READ EnumresultCriteriaUsedPrivilegeProp = "metrics-read"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_REMOTE_LOG_READ EnumresultCriteriaUsedPrivilegeProp = "remote-log-read"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_MANAGE_TOPOLOGY EnumresultCriteriaUsedPrivilegeProp = "manage-topology"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_PERMIT_GET_PASSWORD_POLICY_STATE_ISSUES EnumresultCriteriaUsedPrivilegeProp = "permit-get-password-policy-state-issues"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_PERMIT_PROXIED_MSCHAPV2_DETAILS EnumresultCriteriaUsedPrivilegeProp = "permit-proxied-mschapv2-details"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_PERMIT_EXTERNALLY_PROCESSED_AUTHENTICATION EnumresultCriteriaUsedPrivilegeProp = "permit-externally-processed-authentication"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_PERMIT_EXPORT_REVERSIBLE_PASSWORDS EnumresultCriteriaUsedPrivilegeProp = "permit-export-reversible-passwords"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_PERMIT_FORWARDING_CLIENT_CONNECTION_POLICY EnumresultCriteriaUsedPrivilegeProp = "permit-forwarding-client-connection-policy"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_EXEC_TASK EnumresultCriteriaUsedPrivilegeProp = "exec-task"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_COLLECT_SUPPORT_DATA EnumresultCriteriaUsedPrivilegeProp = "collect-support-data"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_FILE_SERVLET_ACCESS EnumresultCriteriaUsedPrivilegeProp = "file-servlet-access"
	ENUMRESULTCRITERIAUSEDPRIVILEGEPROP_PERMIT_REPLACE_CERTIFICATE_REQUEST EnumresultCriteriaUsedPrivilegeProp = "permit-replace-certificate-request"
)

// All allowed values of EnumresultCriteriaUsedPrivilegeProp enum
var AllowedEnumresultCriteriaUsedPrivilegePropEnumValues = []EnumresultCriteriaUsedPrivilegeProp{
	"audit-data-security",
	"bypass-acl",
	"bypass-read-acl",
	"modify-acl",
	"config-read",
	"config-write",
	"jmx-read",
	"jmx-write",
	"jmx-notify",
	"ldif-import",
	"ldif-export",
	"backend-backup",
	"backend-restore",
	"server-shutdown",
	"server-restart",
	"proxied-auth",
	"disconnect-client",
	"password-reset",
	"update-schema",
	"privilege-change",
	"unindexed-search",
	"unindexed-search-with-control",
	"bypass-pw-policy",
	"lockdown-mode",
	"stream-values",
	"third-party-task",
	"use-admin-session",
	"soft-delete-read",
	"metrics-read",
	"remote-log-read",
	"manage-topology",
	"permit-get-password-policy-state-issues",
	"permit-proxied-mschapv2-details",
	"permit-externally-processed-authentication",
	"permit-export-reversible-passwords",
	"permit-forwarding-client-connection-policy",
	"exec-task",
	"collect-support-data",
	"file-servlet-access",
	"permit-replace-certificate-request",
}

func (v *EnumresultCriteriaUsedPrivilegeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumresultCriteriaUsedPrivilegeProp(value)
	for _, existing := range AllowedEnumresultCriteriaUsedPrivilegePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumresultCriteriaUsedPrivilegeProp", value)
}

// NewEnumresultCriteriaUsedPrivilegePropFromValue returns a pointer to a valid EnumresultCriteriaUsedPrivilegeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumresultCriteriaUsedPrivilegePropFromValue(v string) (*EnumresultCriteriaUsedPrivilegeProp, error) {
	ev := EnumresultCriteriaUsedPrivilegeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumresultCriteriaUsedPrivilegeProp: valid values are %v", v, AllowedEnumresultCriteriaUsedPrivilegePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumresultCriteriaUsedPrivilegeProp) IsValid() bool {
	for _, existing := range AllowedEnumresultCriteriaUsedPrivilegePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumresult-criteria-usedPrivilegeProp value
func (v EnumresultCriteriaUsedPrivilegeProp) Ptr() *EnumresultCriteriaUsedPrivilegeProp {
	return &v
}

type NullableEnumresultCriteriaUsedPrivilegeProp struct {
	value *EnumresultCriteriaUsedPrivilegeProp
	isSet bool
}

func (v NullableEnumresultCriteriaUsedPrivilegeProp) Get() *EnumresultCriteriaUsedPrivilegeProp {
	return v.value
}

func (v *NullableEnumresultCriteriaUsedPrivilegeProp) Set(val *EnumresultCriteriaUsedPrivilegeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumresultCriteriaUsedPrivilegeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumresultCriteriaUsedPrivilegeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumresultCriteriaUsedPrivilegeProp(val *EnumresultCriteriaUsedPrivilegeProp) *NullableEnumresultCriteriaUsedPrivilegeProp {
	return &NullableEnumresultCriteriaUsedPrivilegeProp{value: val, isSet: true}
}

func (v NullableEnumresultCriteriaUsedPrivilegeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumresultCriteriaUsedPrivilegeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
