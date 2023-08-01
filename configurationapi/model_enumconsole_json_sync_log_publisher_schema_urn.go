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

// EnumconsoleJsonSyncLogPublisherSchemaUrn the model 'EnumconsoleJsonSyncLogPublisherSchemaUrn'
type EnumconsoleJsonSyncLogPublisherSchemaUrn string

// List of Enumconsole-json-sync-log-publisherSchemaUrn
const (
	ENUMCONSOLEJSONSYNCLOGPUBLISHERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_PUBLISHERCONSOLE_JSON_SYNC EnumconsoleJsonSyncLogPublisherSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-publisher:console-json-sync"
)

// All allowed values of EnumconsoleJsonSyncLogPublisherSchemaUrn enum
var AllowedEnumconsoleJsonSyncLogPublisherSchemaUrnEnumValues = []EnumconsoleJsonSyncLogPublisherSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-publisher:console-json-sync",
}

func (v *EnumconsoleJsonSyncLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumconsoleJsonSyncLogPublisherSchemaUrn(value)
	for _, existing := range AllowedEnumconsoleJsonSyncLogPublisherSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumconsoleJsonSyncLogPublisherSchemaUrn", value)
}

// NewEnumconsoleJsonSyncLogPublisherSchemaUrnFromValue returns a pointer to a valid EnumconsoleJsonSyncLogPublisherSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumconsoleJsonSyncLogPublisherSchemaUrnFromValue(v string) (*EnumconsoleJsonSyncLogPublisherSchemaUrn, error) {
	ev := EnumconsoleJsonSyncLogPublisherSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumconsoleJsonSyncLogPublisherSchemaUrn: valid values are %v", v, AllowedEnumconsoleJsonSyncLogPublisherSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumconsoleJsonSyncLogPublisherSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumconsoleJsonSyncLogPublisherSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumconsole-json-sync-log-publisherSchemaUrn value
func (v EnumconsoleJsonSyncLogPublisherSchemaUrn) Ptr() *EnumconsoleJsonSyncLogPublisherSchemaUrn {
	return &v
}

type NullableEnumconsoleJsonSyncLogPublisherSchemaUrn struct {
	value *EnumconsoleJsonSyncLogPublisherSchemaUrn
	isSet bool
}

func (v NullableEnumconsoleJsonSyncLogPublisherSchemaUrn) Get() *EnumconsoleJsonSyncLogPublisherSchemaUrn {
	return v.value
}

func (v *NullableEnumconsoleJsonSyncLogPublisherSchemaUrn) Set(val *EnumconsoleJsonSyncLogPublisherSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumconsoleJsonSyncLogPublisherSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumconsoleJsonSyncLogPublisherSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumconsoleJsonSyncLogPublisherSchemaUrn(val *EnumconsoleJsonSyncLogPublisherSchemaUrn) *NullableEnumconsoleJsonSyncLogPublisherSchemaUrn {
	return &NullableEnumconsoleJsonSyncLogPublisherSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumconsoleJsonSyncLogPublisherSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumconsoleJsonSyncLogPublisherSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
