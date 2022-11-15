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

// EnumsyslogTextErrorLogPublisherSchemaUrn the model 'EnumsyslogTextErrorLogPublisherSchemaUrn'
type EnumsyslogTextErrorLogPublisherSchemaUrn string

// List of Enumsyslog-text-error-log-publisherSchemaUrn
const (
	ENUMSYSLOGTEXTERRORLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERSYSLOG_TEXT_ERROR EnumsyslogTextErrorLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:syslog-text-error"
)

// All allowed values of EnumsyslogTextErrorLogPublisherSchemaUrn enum
var AllowedEnumsyslogTextErrorLogPublisherSchemaUrnEnumValues = []EnumsyslogTextErrorLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:syslog-text-error",
}

func (v *EnumsyslogTextErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsyslogTextErrorLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumsyslogTextErrorLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsyslogTextErrorLogPublisherSchemaUrn", value)
}

// NewEnumsyslogTextErrorLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumsyslogTextErrorLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsyslogTextErrorLogPublisherSchemaUrnFromValue(v string) (*EnumsyslogTextErrorLogPublisherSchemaUrn, error) {
	ev := EnumsyslogTextErrorLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsyslogTextErrorLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumsyslogTextErrorLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsyslogTextErrorLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsyslogTextErrorLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsyslog-text-error-log-publisherSchemaUrn value
func (v EnumsyslogTextErrorLogPublisherSchemaUrn) Ptr() *EnumsyslogTextErrorLogPublisherSchemaUrn {
	return &v
}

type NullableEnumsyslogTextErrorLogPublisherSchemaUrn struct {
	value *EnumsyslogTextErrorLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumsyslogTextErrorLogPublisherSchemaUrn) Get() *EnumsyslogTextErrorLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumsyslogTextErrorLogPublisherSchemaUrn) Set(val *EnumsyslogTextErrorLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsyslogTextErrorLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsyslogTextErrorLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsyslogTextErrorLogPublisherSchemaUrn(val *EnumsyslogTextErrorLogPublisherSchemaUrn) *NullableEnumsyslogTextErrorLogPublisherSchemaUrn {
	return &NullableEnumsyslogTextErrorLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsyslogTextErrorLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsyslogTextErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

