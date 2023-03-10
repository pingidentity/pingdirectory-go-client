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

// EnumsyslogJsonAuditLogPublisherSchemaUrn the model 'EnumsyslogJsonAuditLogPublisherSchemaUrn'
type EnumsyslogJsonAuditLogPublisherSchemaUrn string

// List of Enumsyslog-json-audit-log-publisherSchemaUrn
const (
	ENUMSYSLOGJSONAUDITLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERSYSLOG_JSON_AUDIT EnumsyslogJsonAuditLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:syslog-json-audit"
)

// All allowed values of EnumsyslogJsonAuditLogPublisherSchemaUrn enum
var AllowedEnumsyslogJsonAuditLogPublisherSchemaUrnEnumValues = []EnumsyslogJsonAuditLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:syslog-json-audit",
}

func (v *EnumsyslogJsonAuditLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsyslogJsonAuditLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumsyslogJsonAuditLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsyslogJsonAuditLogPublisherSchemaUrn", value)
}

// NewEnumsyslogJsonAuditLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumsyslogJsonAuditLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsyslogJsonAuditLogPublisherSchemaUrnFromValue(v string) (*EnumsyslogJsonAuditLogPublisherSchemaUrn, error) {
	ev := EnumsyslogJsonAuditLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsyslogJsonAuditLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumsyslogJsonAuditLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsyslogJsonAuditLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsyslogJsonAuditLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsyslog-json-audit-log-publisherSchemaUrn value
func (v EnumsyslogJsonAuditLogPublisherSchemaUrn) Ptr() *EnumsyslogJsonAuditLogPublisherSchemaUrn {
	return &v
}

type NullableEnumsyslogJsonAuditLogPublisherSchemaUrn struct {
	value *EnumsyslogJsonAuditLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumsyslogJsonAuditLogPublisherSchemaUrn) Get() *EnumsyslogJsonAuditLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumsyslogJsonAuditLogPublisherSchemaUrn) Set(val *EnumsyslogJsonAuditLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsyslogJsonAuditLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsyslogJsonAuditLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsyslogJsonAuditLogPublisherSchemaUrn(val *EnumsyslogJsonAuditLogPublisherSchemaUrn) *NullableEnumsyslogJsonAuditLogPublisherSchemaUrn {
	return &NullableEnumsyslogJsonAuditLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsyslogJsonAuditLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsyslogJsonAuditLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
