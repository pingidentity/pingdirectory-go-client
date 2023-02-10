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

// EnumrootDnUserPrivilegeProp Privileges that are either explicitly granted or revoked from the root user. Privileges can be revoked by including a minus sign (-) before the privilege name. This is stored in the ds-privilege-name LDAP attribute.
type EnumrootDnUserPrivilegeProp string

// List of Enumroot-dn-user-privilegeProp
const (
	ENUMROOTDNUSERPRIVILEGEPROP_AUDIT_DATA_SECURITY                               EnumrootDnUserPrivilegeProp = "audit-data-security"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_AUDIT_DATA_SECURITY                        EnumrootDnUserPrivilegeProp = "-audit-data-security"
	ENUMROOTDNUSERPRIVILEGEPROP_BYPASS_ACL                                        EnumrootDnUserPrivilegeProp = "bypass-acl"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_BYPASS_ACL                                 EnumrootDnUserPrivilegeProp = "-bypass-acl"
	ENUMROOTDNUSERPRIVILEGEPROP_BYPASS_READ_ACL                                   EnumrootDnUserPrivilegeProp = "bypass-read-acl"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_BYPASS_READ_ACL                            EnumrootDnUserPrivilegeProp = "-bypass-read-acl"
	ENUMROOTDNUSERPRIVILEGEPROP_MODIFY_ACL                                        EnumrootDnUserPrivilegeProp = "modify-acl"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_MODIFY_ACL                                 EnumrootDnUserPrivilegeProp = "-modify-acl"
	ENUMROOTDNUSERPRIVILEGEPROP_CONFIG_READ                                       EnumrootDnUserPrivilegeProp = "config-read"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_CONFIG_READ                                EnumrootDnUserPrivilegeProp = "-config-read"
	ENUMROOTDNUSERPRIVILEGEPROP_CONFIG_WRITE                                      EnumrootDnUserPrivilegeProp = "config-write"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_CONFIG_WRITE                               EnumrootDnUserPrivilegeProp = "-config-write"
	ENUMROOTDNUSERPRIVILEGEPROP_JMX_READ                                          EnumrootDnUserPrivilegeProp = "jmx-read"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_JMX_READ                                   EnumrootDnUserPrivilegeProp = "-jmx-read"
	ENUMROOTDNUSERPRIVILEGEPROP_JMX_WRITE                                         EnumrootDnUserPrivilegeProp = "jmx-write"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_JMX_WRITE                                  EnumrootDnUserPrivilegeProp = "-jmx-write"
	ENUMROOTDNUSERPRIVILEGEPROP_JMX_NOTIFY                                        EnumrootDnUserPrivilegeProp = "jmx-notify"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_JMX_NOTIFY                                 EnumrootDnUserPrivilegeProp = "-jmx-notify"
	ENUMROOTDNUSERPRIVILEGEPROP_LDIF_IMPORT                                       EnumrootDnUserPrivilegeProp = "ldif-import"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_LDIF_IMPORT                                EnumrootDnUserPrivilegeProp = "-ldif-import"
	ENUMROOTDNUSERPRIVILEGEPROP_LDIF_EXPORT                                       EnumrootDnUserPrivilegeProp = "ldif-export"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_LDIF_EXPORT                                EnumrootDnUserPrivilegeProp = "-ldif-export"
	ENUMROOTDNUSERPRIVILEGEPROP_BACKEND_BACKUP                                    EnumrootDnUserPrivilegeProp = "backend-backup"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_BACKEND_BACKUP                             EnumrootDnUserPrivilegeProp = "-backend-backup"
	ENUMROOTDNUSERPRIVILEGEPROP_BACKEND_RESTORE                                   EnumrootDnUserPrivilegeProp = "backend-restore"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_BACKEND_RESTORE                            EnumrootDnUserPrivilegeProp = "-backend-restore"
	ENUMROOTDNUSERPRIVILEGEPROP_SERVER_SHUTDOWN                                   EnumrootDnUserPrivilegeProp = "server-shutdown"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_SERVER_SHUTDOWN                            EnumrootDnUserPrivilegeProp = "-server-shutdown"
	ENUMROOTDNUSERPRIVILEGEPROP_SERVER_RESTART                                    EnumrootDnUserPrivilegeProp = "server-restart"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_SERVER_RESTART                             EnumrootDnUserPrivilegeProp = "-server-restart"
	ENUMROOTDNUSERPRIVILEGEPROP_PROXIED_AUTH                                      EnumrootDnUserPrivilegeProp = "proxied-auth"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_PROXIED_AUTH                               EnumrootDnUserPrivilegeProp = "-proxied-auth"
	ENUMROOTDNUSERPRIVILEGEPROP_DISCONNECT_CLIENT                                 EnumrootDnUserPrivilegeProp = "disconnect-client"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_DISCONNECT_CLIENT                          EnumrootDnUserPrivilegeProp = "-disconnect-client"
	ENUMROOTDNUSERPRIVILEGEPROP_PASSWORD_RESET                                    EnumrootDnUserPrivilegeProp = "password-reset"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_PASSWORD_RESET                             EnumrootDnUserPrivilegeProp = "-password-reset"
	ENUMROOTDNUSERPRIVILEGEPROP_UPDATE_SCHEMA                                     EnumrootDnUserPrivilegeProp = "update-schema"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_UPDATE_SCHEMA                              EnumrootDnUserPrivilegeProp = "-update-schema"
	ENUMROOTDNUSERPRIVILEGEPROP_PRIVILEGE_CHANGE                                  EnumrootDnUserPrivilegeProp = "privilege-change"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_PRIVILEGE_CHANGE                           EnumrootDnUserPrivilegeProp = "-privilege-change"
	ENUMROOTDNUSERPRIVILEGEPROP_UNINDEXED_SEARCH                                  EnumrootDnUserPrivilegeProp = "unindexed-search"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_UNINDEXED_SEARCH                           EnumrootDnUserPrivilegeProp = "-unindexed-search"
	ENUMROOTDNUSERPRIVILEGEPROP_UNINDEXED_SEARCH_WITH_CONTROL                     EnumrootDnUserPrivilegeProp = "unindexed-search-with-control"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_UNINDEXED_SEARCH_WITH_CONTROL              EnumrootDnUserPrivilegeProp = "-unindexed-search-with-control"
	ENUMROOTDNUSERPRIVILEGEPROP_BYPASS_PW_POLICY                                  EnumrootDnUserPrivilegeProp = "bypass-pw-policy"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_BYPASS_PW_POLICY                           EnumrootDnUserPrivilegeProp = "-bypass-pw-policy"
	ENUMROOTDNUSERPRIVILEGEPROP_LOCKDOWN_MODE                                     EnumrootDnUserPrivilegeProp = "lockdown-mode"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_LOCKDOWN_MODE                              EnumrootDnUserPrivilegeProp = "-lockdown-mode"
	ENUMROOTDNUSERPRIVILEGEPROP_STREAM_VALUES                                     EnumrootDnUserPrivilegeProp = "stream-values"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_STREAM_VALUES                              EnumrootDnUserPrivilegeProp = "-stream-values"
	ENUMROOTDNUSERPRIVILEGEPROP_THIRD_PARTY_TASK                                  EnumrootDnUserPrivilegeProp = "third-party-task"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_THIRD_PARTY_TASK                           EnumrootDnUserPrivilegeProp = "-third-party-task"
	ENUMROOTDNUSERPRIVILEGEPROP_USE_ADMIN_SESSION                                 EnumrootDnUserPrivilegeProp = "use-admin-session"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_USE_ADMIN_SESSION                          EnumrootDnUserPrivilegeProp = "-use-admin-session"
	ENUMROOTDNUSERPRIVILEGEPROP_SOFT_DELETE_READ                                  EnumrootDnUserPrivilegeProp = "soft-delete-read"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_SOFT_DELETE_READ                           EnumrootDnUserPrivilegeProp = "-soft-delete-read"
	ENUMROOTDNUSERPRIVILEGEPROP_METRICS_READ                                      EnumrootDnUserPrivilegeProp = "metrics-read"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_METRICS_READ                               EnumrootDnUserPrivilegeProp = "-metrics-read"
	ENUMROOTDNUSERPRIVILEGEPROP_REMOTE_LOG_READ                                   EnumrootDnUserPrivilegeProp = "remote-log-read"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_REMOTE_LOG_READ                            EnumrootDnUserPrivilegeProp = "-remote-log-read"
	ENUMROOTDNUSERPRIVILEGEPROP_MANAGE_TOPOLOGY                                   EnumrootDnUserPrivilegeProp = "manage-topology"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_MANAGE_TOPOLOGY                            EnumrootDnUserPrivilegeProp = "-manage-topology"
	ENUMROOTDNUSERPRIVILEGEPROP_PERMIT_GET_PASSWORD_POLICY_STATE_ISSUES           EnumrootDnUserPrivilegeProp = "permit-get-password-policy-state-issues"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_PERMIT_GET_PASSWORD_POLICY_STATE_ISSUES    EnumrootDnUserPrivilegeProp = "-permit-get-password-policy-state-issues"
	ENUMROOTDNUSERPRIVILEGEPROP_PERMIT_PROXIED_MSCHAPV2_DETAILS                   EnumrootDnUserPrivilegeProp = "permit-proxied-mschapv2-details"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_PERMIT_PROXIED_MSCHAPV2_DETAILS            EnumrootDnUserPrivilegeProp = "-permit-proxied-mschapv2-details"
	ENUMROOTDNUSERPRIVILEGEPROP_PERMIT_EXTERNALLY_PROCESSED_AUTHENTICATION        EnumrootDnUserPrivilegeProp = "permit-externally-processed-authentication"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_PERMIT_EXTERNALLY_PROCESSED_AUTHENTICATION EnumrootDnUserPrivilegeProp = "-permit-externally-processed-authentication"
	ENUMROOTDNUSERPRIVILEGEPROP_PERMIT_EXPORT_REVERSIBLE_PASSWORDS                EnumrootDnUserPrivilegeProp = "permit-export-reversible-passwords"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_PERMIT_EXPORT_REVERSIBLE_PASSWORDS         EnumrootDnUserPrivilegeProp = "-permit-export-reversible-passwords"
	ENUMROOTDNUSERPRIVILEGEPROP_PERMIT_FORWARDING_CLIENT_CONNECTION_POLICY        EnumrootDnUserPrivilegeProp = "permit-forwarding-client-connection-policy"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_PERMIT_FORWARDING_CLIENT_CONNECTION_POLICY EnumrootDnUserPrivilegeProp = "-permit-forwarding-client-connection-policy"
	ENUMROOTDNUSERPRIVILEGEPROP_EXEC_TASK                                         EnumrootDnUserPrivilegeProp = "exec-task"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_EXEC_TASK                                  EnumrootDnUserPrivilegeProp = "-exec-task"
	ENUMROOTDNUSERPRIVILEGEPROP_COLLECT_SUPPORT_DATA                              EnumrootDnUserPrivilegeProp = "collect-support-data"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_COLLECT_SUPPORT_DATA                       EnumrootDnUserPrivilegeProp = "-collect-support-data"
	ENUMROOTDNUSERPRIVILEGEPROP_FILE_SERVLET_ACCESS                               EnumrootDnUserPrivilegeProp = "file-servlet-access"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_FILE_SERVLET_ACCESS                        EnumrootDnUserPrivilegeProp = "-file-servlet-access"
	ENUMROOTDNUSERPRIVILEGEPROP_PERMIT_REPLACE_CERTIFICATE_REQUEST                EnumrootDnUserPrivilegeProp = "permit-replace-certificate-request"
	ENUMROOTDNUSERPRIVILEGEPROP_REVOKE_PERMIT_REPLACE_CERTIFICATE_REQUEST         EnumrootDnUserPrivilegeProp = "-permit-replace-certificate-request"
)

// All allowed values of EnumrootDnUserPrivilegeProp enum
var AllowedEnumrootDnUserPrivilegePropEnumValues = []EnumrootDnUserPrivilegeProp{
	"audit-data-security",
	"-audit-data-security",
	"bypass-acl",
	"-bypass-acl",
	"bypass-read-acl",
	"-bypass-read-acl",
	"modify-acl",
	"-modify-acl",
	"config-read",
	"-config-read",
	"config-write",
	"-config-write",
	"jmx-read",
	"-jmx-read",
	"jmx-write",
	"-jmx-write",
	"jmx-notify",
	"-jmx-notify",
	"ldif-import",
	"-ldif-import",
	"ldif-export",
	"-ldif-export",
	"backend-backup",
	"-backend-backup",
	"backend-restore",
	"-backend-restore",
	"server-shutdown",
	"-server-shutdown",
	"server-restart",
	"-server-restart",
	"proxied-auth",
	"-proxied-auth",
	"disconnect-client",
	"-disconnect-client",
	"password-reset",
	"-password-reset",
	"update-schema",
	"-update-schema",
	"privilege-change",
	"-privilege-change",
	"unindexed-search",
	"-unindexed-search",
	"unindexed-search-with-control",
	"-unindexed-search-with-control",
	"bypass-pw-policy",
	"-bypass-pw-policy",
	"lockdown-mode",
	"-lockdown-mode",
	"stream-values",
	"-stream-values",
	"third-party-task",
	"-third-party-task",
	"use-admin-session",
	"-use-admin-session",
	"soft-delete-read",
	"-soft-delete-read",
	"metrics-read",
	"-metrics-read",
	"remote-log-read",
	"-remote-log-read",
	"manage-topology",
	"-manage-topology",
	"permit-get-password-policy-state-issues",
	"-permit-get-password-policy-state-issues",
	"permit-proxied-mschapv2-details",
	"-permit-proxied-mschapv2-details",
	"permit-externally-processed-authentication",
	"-permit-externally-processed-authentication",
	"permit-export-reversible-passwords",
	"-permit-export-reversible-passwords",
	"permit-forwarding-client-connection-policy",
	"-permit-forwarding-client-connection-policy",
	"exec-task",
	"-exec-task",
	"collect-support-data",
	"-collect-support-data",
	"file-servlet-access",
	"-file-servlet-access",
	"permit-replace-certificate-request",
	"-permit-replace-certificate-request",
}

func (v *EnumrootDnUserPrivilegeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumrootDnUserPrivilegeProp(value)
	for _, existing := range AllowedEnumrootDnUserPrivilegePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumrootDnUserPrivilegeProp", value)
}

// NewEnumrootDnUserPrivilegePropFromValue returns a pointer to a valid EnumrootDnUserPrivilegeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumrootDnUserPrivilegePropFromValue(v string) (*EnumrootDnUserPrivilegeProp, error) {
	ev := EnumrootDnUserPrivilegeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumrootDnUserPrivilegeProp: valid values are %v", v, AllowedEnumrootDnUserPrivilegePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumrootDnUserPrivilegeProp) IsValid() bool {
	for _, existing := range AllowedEnumrootDnUserPrivilegePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumroot-dn-user-privilegeProp value
func (v EnumrootDnUserPrivilegeProp) Ptr() *EnumrootDnUserPrivilegeProp {
	return &v
}

type NullableEnumrootDnUserPrivilegeProp struct {
	value *EnumrootDnUserPrivilegeProp
	isSet bool
}

func (v NullableEnumrootDnUserPrivilegeProp) Get() *EnumrootDnUserPrivilegeProp {
	return v.value
}

func (v *NullableEnumrootDnUserPrivilegeProp) Set(val *EnumrootDnUserPrivilegeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumrootDnUserPrivilegeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumrootDnUserPrivilegeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumrootDnUserPrivilegeProp(val *EnumrootDnUserPrivilegeProp) *NullableEnumrootDnUserPrivilegeProp {
	return &NullableEnumrootDnUserPrivilegeProp{value: val, isSet: true}
}

func (v NullableEnumrootDnUserPrivilegeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumrootDnUserPrivilegeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
