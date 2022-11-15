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

// EnumsyslogBasedErrorLogPublisherSchemaUrn the model 'EnumsyslogBasedErrorLogPublisherSchemaUrn'
type EnumsyslogBasedErrorLogPublisherSchemaUrn string

// List of Enumsyslog-based-error-log-publisherSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERSYSLOG_BASED_ERROR EnumsyslogBasedErrorLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:syslog-based-error"
)

// All allowed values of EnumsyslogBasedErrorLogPublisherSchemaUrn enum
var AllowedEnumsyslogBasedErrorLogPublisherSchemaUrnEnumValues = []EnumsyslogBasedErrorLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:syslog-based-error",
}

func (v *EnumsyslogBasedErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsyslogBasedErrorLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumsyslogBasedErrorLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsyslogBasedErrorLogPublisherSchemaUrn", value)
}

// NewEnumsyslogBasedErrorLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumsyslogBasedErrorLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsyslogBasedErrorLogPublisherSchemaUrnFromValue(v string) (*EnumsyslogBasedErrorLogPublisherSchemaUrn, error) {
	ev := EnumsyslogBasedErrorLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsyslogBasedErrorLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumsyslogBasedErrorLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsyslogBasedErrorLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsyslogBasedErrorLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsyslog-based-error-log-publisherSchemaUrn value
func (v EnumsyslogBasedErrorLogPublisherSchemaUrn) Ptr() *EnumsyslogBasedErrorLogPublisherSchemaUrn {
	return &v
}

type NullableEnumsyslogBasedErrorLogPublisherSchemaUrn struct {
	value *EnumsyslogBasedErrorLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumsyslogBasedErrorLogPublisherSchemaUrn) Get() *EnumsyslogBasedErrorLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumsyslogBasedErrorLogPublisherSchemaUrn) Set(val *EnumsyslogBasedErrorLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsyslogBasedErrorLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsyslogBasedErrorLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsyslogBasedErrorLogPublisherSchemaUrn(val *EnumsyslogBasedErrorLogPublisherSchemaUrn) *NullableEnumsyslogBasedErrorLogPublisherSchemaUrn {
	return &NullableEnumsyslogBasedErrorLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsyslogBasedErrorLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsyslogBasedErrorLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

