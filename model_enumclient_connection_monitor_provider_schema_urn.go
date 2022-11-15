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

// EnumclientConnectionMonitorProviderSchemaUrn the model 'EnumclientConnectionMonitorProviderSchemaUrn'
type EnumclientConnectionMonitorProviderSchemaUrn string

// List of Enumclient-connection-monitor-providerSchemaUrn
const (
	ENUMCLIENTCONNECTIONMONITORPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0MONITOR_PROVIDERCLIENT_CONNECTION EnumclientConnectionMonitorProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:monitor-provider:client-connection"
)

// All allowed values of EnumclientConnectionMonitorProviderSchemaUrn enum
var AllowedEnumclientConnectionMonitorProviderSchemaUrnEnumValues = []EnumclientConnectionMonitorProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:monitor-provider:client-connection",
}

func (v *EnumclientConnectionMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumclientConnectionMonitorProviderSchemaUrn(value)
	for _, existing := range AllowedEnumclientConnectionMonitorProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumclientConnectionMonitorProviderSchemaUrn", value)
}

// NewEnumclientConnectionMonitorProviderSchemaUrnFromValue returns a pointer to a valid EnumclientConnectionMonitorProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumclientConnectionMonitorProviderSchemaUrnFromValue(v string) (*EnumclientConnectionMonitorProviderSchemaUrn, error) {
	ev := EnumclientConnectionMonitorProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumclientConnectionMonitorProviderSchemaUrn: valid values are %v", v, AllowedEnumclientConnectionMonitorProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumclientConnectionMonitorProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumclientConnectionMonitorProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumclient-connection-monitor-providerSchemaUrn value
func (v EnumclientConnectionMonitorProviderSchemaUrn) Ptr() *EnumclientConnectionMonitorProviderSchemaUrn {
	return &v
}

type NullableEnumclientConnectionMonitorProviderSchemaUrn struct {
	value *EnumclientConnectionMonitorProviderSchemaUrn
	isSet bool
}

func (v NullableEnumclientConnectionMonitorProviderSchemaUrn) Get() *EnumclientConnectionMonitorProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumclientConnectionMonitorProviderSchemaUrn) Set(val *EnumclientConnectionMonitorProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumclientConnectionMonitorProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumclientConnectionMonitorProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumclientConnectionMonitorProviderSchemaUrn(val *EnumclientConnectionMonitorProviderSchemaUrn) *NullableEnumclientConnectionMonitorProviderSchemaUrn {
	return &NullableEnumclientConnectionMonitorProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumclientConnectionMonitorProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumclientConnectionMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

