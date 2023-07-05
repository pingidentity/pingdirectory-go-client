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

// EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn the model 'EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn'
type EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn string

// List of Enumencryption-settings-database-accessibility-monitor-providerSchemaUrn
const (
	ENUMENCRYPTIONSETTINGSDATABASEACCESSIBILITYMONITORPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0MONITOR_PROVIDERENCRYPTION_SETTINGS_DATABASE_ACCESSIBILITY EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:monitor-provider:encryption-settings-database-accessibility"
)

// All allowed values of EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn enum
var AllowedEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrnEnumValues = []EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:monitor-provider:encryption-settings-database-accessibility",
}

func (v *EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn(value)
	for _, existing := range AllowedEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn", value)
}

// NewEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrnFromValue returns a pointer to a valid EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrnFromValue(v string) (*EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn, error) {
	ev := EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn: valid values are %v", v, AllowedEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumencryption-settings-database-accessibility-monitor-providerSchemaUrn value
func (v EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) Ptr() *EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn {
	return &v
}

type NullableEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn struct {
	value *EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn
	isSet bool
}

func (v NullableEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) Get() *EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) Set(val *EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn(val *EnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) *NullableEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn {
	return &NullableEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumencryptionSettingsDatabaseAccessibilityMonitorProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
