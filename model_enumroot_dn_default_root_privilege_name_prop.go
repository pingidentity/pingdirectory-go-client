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

// EnumrootDnDefaultRootPrivilegeNameProp Specifies the names of the privileges that root users will be granted by default.
type EnumrootDnDefaultRootPrivilegeNameProp string

// List of Enumroot-dn-defaultRootPrivilegeNameProp
const (
	AUDIT_DATA_SECURITY EnumrootDnDefaultRootPrivilegeNameProp = "audit-data-security"
	BYPASS_ACL EnumrootDnDefaultRootPrivilegeNameProp = "bypass-acl"
	BYPASS_READ_ACL EnumrootDnDefaultRootPrivilegeNameProp = "bypass-read-acl"
	MODIFY_ACL EnumrootDnDefaultRootPrivilegeNameProp = "modify-acl"
	CONFIG_READ EnumrootDnDefaultRootPrivilegeNameProp = "config-read"
	CONFIG_WRITE EnumrootDnDefaultRootPrivilegeNameProp = "config-write"
	JMX_READ EnumrootDnDefaultRootPrivilegeNameProp = "jmx-read"
	JMX_WRITE EnumrootDnDefaultRootPrivilegeNameProp = "jmx-write"
	JMX_NOTIFY EnumrootDnDefaultRootPrivilegeNameProp = "jmx-notify"
	LDIF_IMPORT EnumrootDnDefaultRootPrivilegeNameProp = "ldif-import"
	LDIF_EXPORT EnumrootDnDefaultRootPrivilegeNameProp = "ldif-export"
	BACKEND_BACKUP EnumrootDnDefaultRootPrivilegeNameProp = "backend-backup"
	BACKEND_RESTORE EnumrootDnDefaultRootPrivilegeNameProp = "backend-restore"
	SERVER_SHUTDOWN EnumrootDnDefaultRootPrivilegeNameProp = "server-shutdown"
	SERVER_RESTART EnumrootDnDefaultRootPrivilegeNameProp = "server-restart"
	PROXIED_AUTH EnumrootDnDefaultRootPrivilegeNameProp = "proxied-auth"
	DISCONNECT_CLIENT EnumrootDnDefaultRootPrivilegeNameProp = "disconnect-client"
	PASSWORD_RESET EnumrootDnDefaultRootPrivilegeNameProp = "password-reset"
	UPDATE_SCHEMA EnumrootDnDefaultRootPrivilegeNameProp = "update-schema"
	PRIVILEGE_CHANGE EnumrootDnDefaultRootPrivilegeNameProp = "privilege-change"
	UNINDEXED_SEARCH EnumrootDnDefaultRootPrivilegeNameProp = "unindexed-search"
	UNINDEXED_SEARCH_WITH_CONTROL EnumrootDnDefaultRootPrivilegeNameProp = "unindexed-search-with-control"
	BYPASS_PW_POLICY EnumrootDnDefaultRootPrivilegeNameProp = "bypass-pw-policy"
	LOCKDOWN_MODE EnumrootDnDefaultRootPrivilegeNameProp = "lockdown-mode"
	STREAM_VALUES EnumrootDnDefaultRootPrivilegeNameProp = "stream-values"
	THIRD_PARTY_TASK EnumrootDnDefaultRootPrivilegeNameProp = "third-party-task"
	USE_ADMIN_SESSION EnumrootDnDefaultRootPrivilegeNameProp = "use-admin-session"
	SOFT_DELETE_READ EnumrootDnDefaultRootPrivilegeNameProp = "soft-delete-read"
	METRICS_READ EnumrootDnDefaultRootPrivilegeNameProp = "metrics-read"
	REMOTE_LOG_READ EnumrootDnDefaultRootPrivilegeNameProp = "remote-log-read"
	MANAGE_TOPOLOGY EnumrootDnDefaultRootPrivilegeNameProp = "manage-topology"
	PERMIT_GET_PASSWORD_POLICY_STATE_ISSUES EnumrootDnDefaultRootPrivilegeNameProp = "permit-get-password-policy-state-issues"
	PERMIT_PROXIED_MSCHAPV2_DETAILS EnumrootDnDefaultRootPrivilegeNameProp = "permit-proxied-mschapv2-details"
	PERMIT_EXTERNALLY_PROCESSED_AUTHENTICATION EnumrootDnDefaultRootPrivilegeNameProp = "permit-externally-processed-authentication"
	PERMIT_EXPORT_REVERSIBLE_PASSWORDS EnumrootDnDefaultRootPrivilegeNameProp = "permit-export-reversible-passwords"
	PERMIT_FORWARDING_CLIENT_CONNECTION_POLICY EnumrootDnDefaultRootPrivilegeNameProp = "permit-forwarding-client-connection-policy"
	EXEC_TASK EnumrootDnDefaultRootPrivilegeNameProp = "exec-task"
	COLLECT_SUPPORT_DATA EnumrootDnDefaultRootPrivilegeNameProp = "collect-support-data"
	FILE_SERVLET_ACCESS EnumrootDnDefaultRootPrivilegeNameProp = "file-servlet-access"
	PERMIT_REPLACE_CERTIFICATE_REQUEST EnumrootDnDefaultRootPrivilegeNameProp = "permit-replace-certificate-request"
)

// All allowed values of EnumrootDnDefaultRootPrivilegeNameProp enum
var AllowedEnumrootDnDefaultRootPrivilegeNamePropEnumValues = []EnumrootDnDefaultRootPrivilegeNameProp{
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

func (v *EnumrootDnDefaultRootPrivilegeNameProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumrootDnDefaultRootPrivilegeNameProp(value)
	for _, existing := range AllowedEnumrootDnDefaultRootPrivilegeNamePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumrootDnDefaultRootPrivilegeNameProp", value)
}

// NewEnumrootDnDefaultRootPrivilegeNamePropFromValue returns a pointer to a valid EnumrootDnDefaultRootPrivilegeNameProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumrootDnDefaultRootPrivilegeNamePropFromValue(v string) (*EnumrootDnDefaultRootPrivilegeNameProp, error) {
	ev := EnumrootDnDefaultRootPrivilegeNameProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumrootDnDefaultRootPrivilegeNameProp: valid values are %v", v, AllowedEnumrootDnDefaultRootPrivilegeNamePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumrootDnDefaultRootPrivilegeNameProp) IsValid() bool {
	for _, existing := range AllowedEnumrootDnDefaultRootPrivilegeNamePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumroot-dn-defaultRootPrivilegeNameProp value
func (v EnumrootDnDefaultRootPrivilegeNameProp) Ptr() *EnumrootDnDefaultRootPrivilegeNameProp {
	return &v
}

type NullableEnumrootDnDefaultRootPrivilegeNameProp struct {
	value *EnumrootDnDefaultRootPrivilegeNameProp
	isSet bool
}

func (v NullableEnumrootDnDefaultRootPrivilegeNameProp) Get() *EnumrootDnDefaultRootPrivilegeNameProp {
	return v.value
}

func (v *NullableEnumrootDnDefaultRootPrivilegeNameProp) Set(val *EnumrootDnDefaultRootPrivilegeNameProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumrootDnDefaultRootPrivilegeNameProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumrootDnDefaultRootPrivilegeNameProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumrootDnDefaultRootPrivilegeNameProp(val *EnumrootDnDefaultRootPrivilegeNameProp) *NullableEnumrootDnDefaultRootPrivilegeNameProp {
	return &NullableEnumrootDnDefaultRootPrivilegeNameProp{value: val, isSet: true}
}

func (v NullableEnumrootDnDefaultRootPrivilegeNameProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumrootDnDefaultRootPrivilegeNameProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

