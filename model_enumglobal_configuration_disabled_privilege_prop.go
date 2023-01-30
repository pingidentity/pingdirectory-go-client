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

// EnumglobalConfigurationDisabledPrivilegeProp Specifies the name of a privilege that should not be evaluated by the server.
type EnumglobalConfigurationDisabledPrivilegeProp string

// List of Enumglobal-configuration-disabledPrivilegeProp
const (
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_AUDIT_DATA_SECURITY                        EnumglobalConfigurationDisabledPrivilegeProp = "audit-data-security"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_BYPASS_ACL                                 EnumglobalConfigurationDisabledPrivilegeProp = "bypass-acl"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_BYPASS_READ_ACL                            EnumglobalConfigurationDisabledPrivilegeProp = "bypass-read-acl"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_MODIFY_ACL                                 EnumglobalConfigurationDisabledPrivilegeProp = "modify-acl"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_CONFIG_READ                                EnumglobalConfigurationDisabledPrivilegeProp = "config-read"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_CONFIG_WRITE                               EnumglobalConfigurationDisabledPrivilegeProp = "config-write"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_JMX_READ                                   EnumglobalConfigurationDisabledPrivilegeProp = "jmx-read"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_JMX_WRITE                                  EnumglobalConfigurationDisabledPrivilegeProp = "jmx-write"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_JMX_NOTIFY                                 EnumglobalConfigurationDisabledPrivilegeProp = "jmx-notify"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_LDIF_IMPORT                                EnumglobalConfigurationDisabledPrivilegeProp = "ldif-import"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_LDIF_EXPORT                                EnumglobalConfigurationDisabledPrivilegeProp = "ldif-export"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_BACKEND_BACKUP                             EnumglobalConfigurationDisabledPrivilegeProp = "backend-backup"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_BACKEND_RESTORE                            EnumglobalConfigurationDisabledPrivilegeProp = "backend-restore"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_SERVER_SHUTDOWN                            EnumglobalConfigurationDisabledPrivilegeProp = "server-shutdown"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_SERVER_RESTART                             EnumglobalConfigurationDisabledPrivilegeProp = "server-restart"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_PROXIED_AUTH                               EnumglobalConfigurationDisabledPrivilegeProp = "proxied-auth"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_DISCONNECT_CLIENT                          EnumglobalConfigurationDisabledPrivilegeProp = "disconnect-client"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_PASSWORD_RESET                             EnumglobalConfigurationDisabledPrivilegeProp = "password-reset"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_UPDATE_SCHEMA                              EnumglobalConfigurationDisabledPrivilegeProp = "update-schema"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_PRIVILEGE_CHANGE                           EnumglobalConfigurationDisabledPrivilegeProp = "privilege-change"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_UNINDEXED_SEARCH                           EnumglobalConfigurationDisabledPrivilegeProp = "unindexed-search"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_UNINDEXED_SEARCH_WITH_CONTROL              EnumglobalConfigurationDisabledPrivilegeProp = "unindexed-search-with-control"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_BYPASS_PW_POLICY                           EnumglobalConfigurationDisabledPrivilegeProp = "bypass-pw-policy"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_LOCKDOWN_MODE                              EnumglobalConfigurationDisabledPrivilegeProp = "lockdown-mode"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_STREAM_VALUES                              EnumglobalConfigurationDisabledPrivilegeProp = "stream-values"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_THIRD_PARTY_TASK                           EnumglobalConfigurationDisabledPrivilegeProp = "third-party-task"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_USE_ADMIN_SESSION                          EnumglobalConfigurationDisabledPrivilegeProp = "use-admin-session"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_SOFT_DELETE_READ                           EnumglobalConfigurationDisabledPrivilegeProp = "soft-delete-read"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_METRICS_READ                               EnumglobalConfigurationDisabledPrivilegeProp = "metrics-read"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_REMOTE_LOG_READ                            EnumglobalConfigurationDisabledPrivilegeProp = "remote-log-read"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_MANAGE_TOPOLOGY                            EnumglobalConfigurationDisabledPrivilegeProp = "manage-topology"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_PERMIT_GET_PASSWORD_POLICY_STATE_ISSUES    EnumglobalConfigurationDisabledPrivilegeProp = "permit-get-password-policy-state-issues"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_PERMIT_PROXIED_MSCHAPV2_DETAILS            EnumglobalConfigurationDisabledPrivilegeProp = "permit-proxied-mschapv2-details"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_PERMIT_EXTERNALLY_PROCESSED_AUTHENTICATION EnumglobalConfigurationDisabledPrivilegeProp = "permit-externally-processed-authentication"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_PERMIT_EXPORT_REVERSIBLE_PASSWORDS         EnumglobalConfigurationDisabledPrivilegeProp = "permit-export-reversible-passwords"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_PERMIT_FORWARDING_CLIENT_CONNECTION_POLICY EnumglobalConfigurationDisabledPrivilegeProp = "permit-forwarding-client-connection-policy"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_EXEC_TASK                                  EnumglobalConfigurationDisabledPrivilegeProp = "exec-task"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_COLLECT_SUPPORT_DATA                       EnumglobalConfigurationDisabledPrivilegeProp = "collect-support-data"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_FILE_SERVLET_ACCESS                        EnumglobalConfigurationDisabledPrivilegeProp = "file-servlet-access"
	ENUMGLOBALCONFIGURATIONDISABLEDPRIVILEGEPROP_PERMIT_REPLACE_CERTIFICATE_REQUEST         EnumglobalConfigurationDisabledPrivilegeProp = "permit-replace-certificate-request"
)

// All allowed values of EnumglobalConfigurationDisabledPrivilegeProp enum
var AllowedEnumglobalConfigurationDisabledPrivilegePropEnumValues = []EnumglobalConfigurationDisabledPrivilegeProp{
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

func (v *EnumglobalConfigurationDisabledPrivilegeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumglobalConfigurationDisabledPrivilegeProp(value)
	for _, existing := range AllowedEnumglobalConfigurationDisabledPrivilegePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumglobalConfigurationDisabledPrivilegeProp", value)
}

// NewEnumglobalConfigurationDisabledPrivilegePropFromValue returns a pointer to a valid EnumglobalConfigurationDisabledPrivilegeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumglobalConfigurationDisabledPrivilegePropFromValue(v string) (*EnumglobalConfigurationDisabledPrivilegeProp, error) {
	ev := EnumglobalConfigurationDisabledPrivilegeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumglobalConfigurationDisabledPrivilegeProp: valid values are %v", v, AllowedEnumglobalConfigurationDisabledPrivilegePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumglobalConfigurationDisabledPrivilegeProp) IsValid() bool {
	for _, existing := range AllowedEnumglobalConfigurationDisabledPrivilegePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumglobal-configuration-disabledPrivilegeProp value
func (v EnumglobalConfigurationDisabledPrivilegeProp) Ptr() *EnumglobalConfigurationDisabledPrivilegeProp {
	return &v
}

type NullableEnumglobalConfigurationDisabledPrivilegeProp struct {
	value *EnumglobalConfigurationDisabledPrivilegeProp
	isSet bool
}

func (v NullableEnumglobalConfigurationDisabledPrivilegeProp) Get() *EnumglobalConfigurationDisabledPrivilegeProp {
	return v.value
}

func (v *NullableEnumglobalConfigurationDisabledPrivilegeProp) Set(val *EnumglobalConfigurationDisabledPrivilegeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumglobalConfigurationDisabledPrivilegeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumglobalConfigurationDisabledPrivilegeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumglobalConfigurationDisabledPrivilegeProp(val *EnumglobalConfigurationDisabledPrivilegeProp) *NullableEnumglobalConfigurationDisabledPrivilegeProp {
	return &NullableEnumglobalConfigurationDisabledPrivilegeProp{value: val, isSet: true}
}

func (v NullableEnumglobalConfigurationDisabledPrivilegeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumglobalConfigurationDisabledPrivilegeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
