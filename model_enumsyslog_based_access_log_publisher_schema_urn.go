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

// EnumsyslogBasedAccessLogPublisherSchemaUrn the model 'EnumsyslogBasedAccessLogPublisherSchemaUrn'
type EnumsyslogBasedAccessLogPublisherSchemaUrn string

// List of Enumsyslog-based-access-log-publisherSchemaUrn
const (
	ENUMSYSLOGBASEDACCESSLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERSYSLOG_BASED_ACCESS EnumsyslogBasedAccessLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:syslog-based-access"
)

// All allowed values of EnumsyslogBasedAccessLogPublisherSchemaUrn enum
var AllowedEnumsyslogBasedAccessLogPublisherSchemaUrnEnumValues = []EnumsyslogBasedAccessLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:syslog-based-access",
}

func (v *EnumsyslogBasedAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsyslogBasedAccessLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumsyslogBasedAccessLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsyslogBasedAccessLogPublisherSchemaUrn", value)
}

// NewEnumsyslogBasedAccessLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumsyslogBasedAccessLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsyslogBasedAccessLogPublisherSchemaUrnFromValue(v string) (*EnumsyslogBasedAccessLogPublisherSchemaUrn, error) {
	ev := EnumsyslogBasedAccessLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsyslogBasedAccessLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumsyslogBasedAccessLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsyslogBasedAccessLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsyslogBasedAccessLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsyslog-based-access-log-publisherSchemaUrn value
func (v EnumsyslogBasedAccessLogPublisherSchemaUrn) Ptr() *EnumsyslogBasedAccessLogPublisherSchemaUrn {
	return &v
}

type NullableEnumsyslogBasedAccessLogPublisherSchemaUrn struct {
	value *EnumsyslogBasedAccessLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumsyslogBasedAccessLogPublisherSchemaUrn) Get() *EnumsyslogBasedAccessLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumsyslogBasedAccessLogPublisherSchemaUrn) Set(val *EnumsyslogBasedAccessLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsyslogBasedAccessLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsyslogBasedAccessLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsyslogBasedAccessLogPublisherSchemaUrn(val *EnumsyslogBasedAccessLogPublisherSchemaUrn) *NullableEnumsyslogBasedAccessLogPublisherSchemaUrn {
	return &NullableEnumsyslogBasedAccessLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsyslogBasedAccessLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsyslogBasedAccessLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
