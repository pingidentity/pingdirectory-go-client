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

// EnumconnectionCriteriaAllIncludedUserPrivilegeProp Specifies the name of a privilege that must be held by the authenticated user for clients included in this Simple Connection Criteria. If any privilege names are provided, then the authenticated user must have all of those privileges. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections.
type EnumconnectionCriteriaAllIncludedUserPrivilegeProp string

// List of Enumconnection-criteria-allIncludedUserPrivilegeProp
const (
	AUDIT_DATA_SECURITY EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "audit-data-security"
	BYPASS_ACL EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "bypass-acl"
	BYPASS_READ_ACL EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "bypass-read-acl"
	MODIFY_ACL EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "modify-acl"
	CONFIG_READ EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "config-read"
	CONFIG_WRITE EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "config-write"
	JMX_READ EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "jmx-read"
	JMX_WRITE EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "jmx-write"
	JMX_NOTIFY EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "jmx-notify"
	LDIF_IMPORT EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "ldif-import"
	LDIF_EXPORT EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "ldif-export"
	BACKEND_BACKUP EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "backend-backup"
	BACKEND_RESTORE EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "backend-restore"
	SERVER_SHUTDOWN EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "server-shutdown"
	SERVER_RESTART EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "server-restart"
	PROXIED_AUTH EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "proxied-auth"
	DISCONNECT_CLIENT EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "disconnect-client"
	PASSWORD_RESET EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "password-reset"
	UPDATE_SCHEMA EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "update-schema"
	PRIVILEGE_CHANGE EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "privilege-change"
	UNINDEXED_SEARCH EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "unindexed-search"
	UNINDEXED_SEARCH_WITH_CONTROL EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "unindexed-search-with-control"
	BYPASS_PW_POLICY EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "bypass-pw-policy"
	LOCKDOWN_MODE EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "lockdown-mode"
	STREAM_VALUES EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "stream-values"
	THIRD_PARTY_TASK EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "third-party-task"
	USE_ADMIN_SESSION EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "use-admin-session"
	SOFT_DELETE_READ EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "soft-delete-read"
	METRICS_READ EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "metrics-read"
	REMOTE_LOG_READ EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "remote-log-read"
	MANAGE_TOPOLOGY EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "manage-topology"
	PERMIT_GET_PASSWORD_POLICY_STATE_ISSUES EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "permit-get-password-policy-state-issues"
	PERMIT_PROXIED_MSCHAPV2_DETAILS EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "permit-proxied-mschapv2-details"
	PERMIT_EXTERNALLY_PROCESSED_AUTHENTICATION EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "permit-externally-processed-authentication"
	PERMIT_EXPORT_REVERSIBLE_PASSWORDS EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "permit-export-reversible-passwords"
	PERMIT_FORWARDING_CLIENT_CONNECTION_POLICY EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "permit-forwarding-client-connection-policy"
	EXEC_TASK EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "exec-task"
	COLLECT_SUPPORT_DATA EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "collect-support-data"
	FILE_SERVLET_ACCESS EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "file-servlet-access"
	PERMIT_REPLACE_CERTIFICATE_REQUEST EnumconnectionCriteriaAllIncludedUserPrivilegeProp = "permit-replace-certificate-request"
)

// All allowed values of EnumconnectionCriteriaAllIncludedUserPrivilegeProp enum
var AllowedEnumconnectionCriteriaAllIncludedUserPrivilegePropEnumValues = []EnumconnectionCriteriaAllIncludedUserPrivilegeProp{
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

func (v *EnumconnectionCriteriaAllIncludedUserPrivilegeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumconnectionCriteriaAllIncludedUserPrivilegeProp(value)
	for _, existing := range AllowedEnumconnectionCriteriaAllIncludedUserPrivilegePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumconnectionCriteriaAllIncludedUserPrivilegeProp", value)
}

// NewEnumconnectionCriteriaAllIncludedUserPrivilegePropFromValue returns a pointer to a valid EnumconnectionCriteriaAllIncludedUserPrivilegeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumconnectionCriteriaAllIncludedUserPrivilegePropFromValue(v string) (*EnumconnectionCriteriaAllIncludedUserPrivilegeProp, error) {
	ev := EnumconnectionCriteriaAllIncludedUserPrivilegeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumconnectionCriteriaAllIncludedUserPrivilegeProp: valid values are %v", v, AllowedEnumconnectionCriteriaAllIncludedUserPrivilegePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumconnectionCriteriaAllIncludedUserPrivilegeProp) IsValid() bool {
	for _, existing := range AllowedEnumconnectionCriteriaAllIncludedUserPrivilegePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumconnection-criteria-allIncludedUserPrivilegeProp value
func (v EnumconnectionCriteriaAllIncludedUserPrivilegeProp) Ptr() *EnumconnectionCriteriaAllIncludedUserPrivilegeProp {
	return &v
}

type NullableEnumconnectionCriteriaAllIncludedUserPrivilegeProp struct {
	value *EnumconnectionCriteriaAllIncludedUserPrivilegeProp
	isSet bool
}

func (v NullableEnumconnectionCriteriaAllIncludedUserPrivilegeProp) Get() *EnumconnectionCriteriaAllIncludedUserPrivilegeProp {
	return v.value
}

func (v *NullableEnumconnectionCriteriaAllIncludedUserPrivilegeProp) Set(val *EnumconnectionCriteriaAllIncludedUserPrivilegeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumconnectionCriteriaAllIncludedUserPrivilegeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumconnectionCriteriaAllIncludedUserPrivilegeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumconnectionCriteriaAllIncludedUserPrivilegeProp(val *EnumconnectionCriteriaAllIncludedUserPrivilegeProp) *NullableEnumconnectionCriteriaAllIncludedUserPrivilegeProp {
	return &NullableEnumconnectionCriteriaAllIncludedUserPrivilegeProp{value: val, isSet: true}
}

func (v NullableEnumconnectionCriteriaAllIncludedUserPrivilegeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumconnectionCriteriaAllIncludedUserPrivilegeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

