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

// EnumsyslogExternalServerSchemaUrn the model 'EnumsyslogExternalServerSchemaUrn'
type EnumsyslogExternalServerSchemaUrn string

// List of Enumsyslog-external-serverSchemaUrn
const (
	ENUMSYSLOGEXTERNALSERVERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERSYSLOG EnumsyslogExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:syslog"
)

// All allowed values of EnumsyslogExternalServerSchemaUrn enum
var AllowedEnumsyslogExternalServerSchemaUrnEnumValues = []EnumsyslogExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:syslog",
}

func (v *EnumsyslogExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsyslogExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumsyslogExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsyslogExternalServerSchemaUrn", value)
}

// NewEnumsyslogExternalServerSchemaUrnFromValue returns a pointer to a valid EnumsyslogExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsyslogExternalServerSchemaUrnFromValue(v string) (*EnumsyslogExternalServerSchemaUrn, error) {
	ev := EnumsyslogExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsyslogExternalServerSchemaUrn: valid values are %v", v, AllowedEnumsyslogExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsyslogExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsyslogExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsyslog-external-serverSchemaUrn value
func (v EnumsyslogExternalServerSchemaUrn) Ptr() *EnumsyslogExternalServerSchemaUrn {
	return &v
}

type NullableEnumsyslogExternalServerSchemaUrn struct {
	value *EnumsyslogExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumsyslogExternalServerSchemaUrn) Get() *EnumsyslogExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumsyslogExternalServerSchemaUrn) Set(val *EnumsyslogExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsyslogExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsyslogExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsyslogExternalServerSchemaUrn(val *EnumsyslogExternalServerSchemaUrn) *NullableEnumsyslogExternalServerSchemaUrn {
	return &NullableEnumsyslogExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsyslogExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsyslogExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
