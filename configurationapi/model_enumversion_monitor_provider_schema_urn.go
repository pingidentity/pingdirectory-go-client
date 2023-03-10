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

// EnumversionMonitorProviderSchemaUrn the model 'EnumversionMonitorProviderSchemaUrn'
type EnumversionMonitorProviderSchemaUrn string

// List of Enumversion-monitor-providerSchemaUrn
const (
	ENUMVERSIONMONITORPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0MONITOR_PROVIDERVERSION EnumversionMonitorProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:monitor-provider:version"
)

// All allowed values of EnumversionMonitorProviderSchemaUrn enum
var AllowedEnumversionMonitorProviderSchemaUrnEnumValues = []EnumversionMonitorProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:monitor-provider:version",
}

func (v *EnumversionMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumversionMonitorProviderSchemaUrn(value)
	for _, existing := range AllowedEnumversionMonitorProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumversionMonitorProviderSchemaUrn", value)
}

// NewEnumversionMonitorProviderSchemaUrnFromValue returns a pointer to a valid EnumversionMonitorProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumversionMonitorProviderSchemaUrnFromValue(v string) (*EnumversionMonitorProviderSchemaUrn, error) {
	ev := EnumversionMonitorProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumversionMonitorProviderSchemaUrn: valid values are %v", v, AllowedEnumversionMonitorProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumversionMonitorProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumversionMonitorProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumversion-monitor-providerSchemaUrn value
func (v EnumversionMonitorProviderSchemaUrn) Ptr() *EnumversionMonitorProviderSchemaUrn {
	return &v
}

type NullableEnumversionMonitorProviderSchemaUrn struct {
	value *EnumversionMonitorProviderSchemaUrn
	isSet bool
}

func (v NullableEnumversionMonitorProviderSchemaUrn) Get() *EnumversionMonitorProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumversionMonitorProviderSchemaUrn) Set(val *EnumversionMonitorProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumversionMonitorProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumversionMonitorProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumversionMonitorProviderSchemaUrn(val *EnumversionMonitorProviderSchemaUrn) *NullableEnumversionMonitorProviderSchemaUrn {
	return &NullableEnumversionMonitorProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumversionMonitorProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumversionMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
