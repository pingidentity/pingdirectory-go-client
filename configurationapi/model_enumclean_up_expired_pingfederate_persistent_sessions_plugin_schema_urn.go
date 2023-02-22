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

// EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn the model 'EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn'
type EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn string

// List of Enumclean-up-expired-pingfederate-persistent-sessions-pluginSchemaUrn
const (
	ENUMCLEANUPEXPIREDPINGFEDERATEPERSISTENTSESSIONSPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGINCLEAN_UP_EXPIRED_PINGFEDERATE_PERSISTENT_SESSIONS EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:clean-up-expired-pingfederate-persistent-sessions"
)

// All allowed values of EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn enum
var AllowedEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrnEnumValues = []EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:clean-up-expired-pingfederate-persistent-sessions",
}

func (v *EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn(value)
	for _, existing := range AllowedEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn", value)
}

// NewEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrnFromValue returns a pointer to a valid EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrnFromValue(v string) (*EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn, error) {
	ev := EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn: valid values are %v", v, AllowedEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumclean-up-expired-pingfederate-persistent-sessions-pluginSchemaUrn value
func (v EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) Ptr() *EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn {
	return &v
}

type NullableEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn struct {
	value *EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn
	isSet bool
}

func (v NullableEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) Get() *EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn {
	return v.value
}

func (v *NullableEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) Set(val *EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn(val *EnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) *NullableEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn {
	return &NullableEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcleanUpExpiredPingfederatePersistentSessionsPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}