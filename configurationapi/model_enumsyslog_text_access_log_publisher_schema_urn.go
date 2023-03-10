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

// EnumsyslogTextAccessLogPublisherSchemaUrn the model 'EnumsyslogTextAccessLogPublisherSchemaUrn'
type EnumsyslogTextAccessLogPublisherSchemaUrn string

// List of Enumsyslog-text-access-log-publisherSchemaUrn
const (
	ENUMSYSLOGTEXTACCESSLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERSYSLOG_TEXT_ACCESS EnumsyslogTextAccessLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:syslog-text-access"
)

// All allowed values of EnumsyslogTextAccessLogPublisherSchemaUrn enum
var AllowedEnumsyslogTextAccessLogPublisherSchemaUrnEnumValues = []EnumsyslogTextAccessLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:syslog-text-access",
}

func (v *EnumsyslogTextAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsyslogTextAccessLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumsyslogTextAccessLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsyslogTextAccessLogPublisherSchemaUrn", value)
}

// NewEnumsyslogTextAccessLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumsyslogTextAccessLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsyslogTextAccessLogPublisherSchemaUrnFromValue(v string) (*EnumsyslogTextAccessLogPublisherSchemaUrn, error) {
	ev := EnumsyslogTextAccessLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsyslogTextAccessLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumsyslogTextAccessLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsyslogTextAccessLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsyslogTextAccessLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsyslog-text-access-log-publisherSchemaUrn value
func (v EnumsyslogTextAccessLogPublisherSchemaUrn) Ptr() *EnumsyslogTextAccessLogPublisherSchemaUrn {
	return &v
}

type NullableEnumsyslogTextAccessLogPublisherSchemaUrn struct {
	value *EnumsyslogTextAccessLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumsyslogTextAccessLogPublisherSchemaUrn) Get() *EnumsyslogTextAccessLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumsyslogTextAccessLogPublisherSchemaUrn) Set(val *EnumsyslogTextAccessLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsyslogTextAccessLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsyslogTextAccessLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsyslogTextAccessLogPublisherSchemaUrn(val *EnumsyslogTextAccessLogPublisherSchemaUrn) *NullableEnumsyslogTextAccessLogPublisherSchemaUrn {
	return &NullableEnumsyslogTextAccessLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsyslogTextAccessLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsyslogTextAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
