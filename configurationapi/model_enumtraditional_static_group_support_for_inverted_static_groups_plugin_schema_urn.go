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

// EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn the model 'EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn'
type EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn string

// List of Enumtraditional-static-group-support-for-inverted-static-groups-pluginSchemaUrn
const (
	ENUMTRADITIONALSTATICGROUPSUPPORTFORINVERTEDSTATICGROUPSPLUGINSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PLUGINTRADITIONAL_STATIC_GROUP_SUPPORT_FOR_INVERTED_STATIC_GROUPS EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:plugin:traditional-static-group-support-for-inverted-static-groups"
)

// All allowed values of EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn enum
var AllowedEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrnEnumValues = []EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:plugin:traditional-static-group-support-for-inverted-static-groups",
}

func (v *EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn(value)
	for _, existing := range AllowedEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn", value)
}

// NewEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrnFromValue returns a pointer to a valid EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrnFromValue(v string) (*EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn, error) {
	ev := EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn: valid values are %v", v, AllowedEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumtraditional-static-group-support-for-inverted-static-groups-pluginSchemaUrn value
func (v EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn) Ptr() *EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn {
	return &v
}

type NullableEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn struct {
	value *EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn
	isSet bool
}

func (v NullableEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn) Get() *EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn {
	return v.value
}

func (v *NullableEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn) Set(val *EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn(val *EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn) *NullableEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn {
	return &NullableEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
