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

// EnumlogPublisherSyslogSeverityProp The syslog severity to use for the messages that are logged by this Syslog JSON Audit Log Publisher.
type EnumlogPublisherSyslogSeverityProp string

// List of Enumlog-publisher-syslogSeverityProp
const (
	ENUMLOGPUBLISHERSYSLOGSEVERITYPROP_EMERGENCY     EnumlogPublisherSyslogSeverityProp = "emergency"
	ENUMLOGPUBLISHERSYSLOGSEVERITYPROP_ALERT         EnumlogPublisherSyslogSeverityProp = "alert"
	ENUMLOGPUBLISHERSYSLOGSEVERITYPROP_CRITICAL      EnumlogPublisherSyslogSeverityProp = "critical"
	ENUMLOGPUBLISHERSYSLOGSEVERITYPROP_ERROR         EnumlogPublisherSyslogSeverityProp = "error"
	ENUMLOGPUBLISHERSYSLOGSEVERITYPROP_WARNING       EnumlogPublisherSyslogSeverityProp = "warning"
	ENUMLOGPUBLISHERSYSLOGSEVERITYPROP_NOTICE        EnumlogPublisherSyslogSeverityProp = "notice"
	ENUMLOGPUBLISHERSYSLOGSEVERITYPROP_INFORMATIONAL EnumlogPublisherSyslogSeverityProp = "informational"
	ENUMLOGPUBLISHERSYSLOGSEVERITYPROP_DEBUG         EnumlogPublisherSyslogSeverityProp = "debug"
)

// All allowed values of EnumlogPublisherSyslogSeverityProp enum
var AllowedEnumlogPublisherSyslogSeverityPropEnumValues = []EnumlogPublisherSyslogSeverityProp{
	"emergency",
	"alert",
	"critical",
	"error",
	"warning",
	"notice",
	"informational",
	"debug",
}

func (v *EnumlogPublisherSyslogSeverityProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherSyslogSeverityProp(value)
	for _, existing := range AllowedEnumlogPublisherSyslogSeverityPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherSyslogSeverityProp", value)
}

// NewEnumlogPublisherSyslogSeverityPropFromValue returns a pointer to a valid EnumlogPublisherSyslogSeverityProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherSyslogSeverityPropFromValue(v string) (*EnumlogPublisherSyslogSeverityProp, error) {
	ev := EnumlogPublisherSyslogSeverityProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherSyslogSeverityProp: valid values are %v", v, AllowedEnumlogPublisherSyslogSeverityPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherSyslogSeverityProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherSyslogSeverityPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-syslogSeverityProp value
func (v EnumlogPublisherSyslogSeverityProp) Ptr() *EnumlogPublisherSyslogSeverityProp {
	return &v
}

type NullableEnumlogPublisherSyslogSeverityProp struct {
	value *EnumlogPublisherSyslogSeverityProp
	isSet bool
}

func (v NullableEnumlogPublisherSyslogSeverityProp) Get() *EnumlogPublisherSyslogSeverityProp {
	return v.value
}

func (v *NullableEnumlogPublisherSyslogSeverityProp) Set(val *EnumlogPublisherSyslogSeverityProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherSyslogSeverityProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherSyslogSeverityProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherSyslogSeverityProp(val *EnumlogPublisherSyslogSeverityProp) *NullableEnumlogPublisherSyslogSeverityProp {
	return &NullableEnumlogPublisherSyslogSeverityProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherSyslogSeverityProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherSyslogSeverityProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
