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

// EnumpluginTraditionalStaticGroupObjectClassProp The object class that defines the type of traditional static group that this plugin will attempt to emulate for inverted static groups.
type EnumpluginTraditionalStaticGroupObjectClassProp string

// List of Enumplugin-traditionalStaticGroupObjectClassProp
const (
	ENUMPLUGINTRADITIONALSTATICGROUPOBJECTCLASSPROP_GROUP_OF_NAMES        EnumpluginTraditionalStaticGroupObjectClassProp = "groupOfNames"
	ENUMPLUGINTRADITIONALSTATICGROUPOBJECTCLASSPROP_GROUP_OF_UNIQUE_NAMES EnumpluginTraditionalStaticGroupObjectClassProp = "groupOfUniqueNames"
	ENUMPLUGINTRADITIONALSTATICGROUPOBJECTCLASSPROP_GROUP_OF_ENTRIES      EnumpluginTraditionalStaticGroupObjectClassProp = "groupOfEntries"
)

// All allowed values of EnumpluginTraditionalStaticGroupObjectClassProp enum
var AllowedEnumpluginTraditionalStaticGroupObjectClassPropEnumValues = []EnumpluginTraditionalStaticGroupObjectClassProp{
	"groupOfNames",
	"groupOfUniqueNames",
	"groupOfEntries",
}

func (v *EnumpluginTraditionalStaticGroupObjectClassProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginTraditionalStaticGroupObjectClassProp(value)
	for _, existing := range AllowedEnumpluginTraditionalStaticGroupObjectClassPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginTraditionalStaticGroupObjectClassProp", value)
}

// NewEnumpluginTraditionalStaticGroupObjectClassPropFromValue returns a pointer to a valid EnumpluginTraditionalStaticGroupObjectClassProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginTraditionalStaticGroupObjectClassPropFromValue(v string) (*EnumpluginTraditionalStaticGroupObjectClassProp, error) {
	ev := EnumpluginTraditionalStaticGroupObjectClassProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginTraditionalStaticGroupObjectClassProp: valid values are %v", v, AllowedEnumpluginTraditionalStaticGroupObjectClassPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginTraditionalStaticGroupObjectClassProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginTraditionalStaticGroupObjectClassPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-traditionalStaticGroupObjectClassProp value
func (v EnumpluginTraditionalStaticGroupObjectClassProp) Ptr() *EnumpluginTraditionalStaticGroupObjectClassProp {
	return &v
}

type NullableEnumpluginTraditionalStaticGroupObjectClassProp struct {
	value *EnumpluginTraditionalStaticGroupObjectClassProp
	isSet bool
}

func (v NullableEnumpluginTraditionalStaticGroupObjectClassProp) Get() *EnumpluginTraditionalStaticGroupObjectClassProp {
	return v.value
}

func (v *NullableEnumpluginTraditionalStaticGroupObjectClassProp) Set(val *EnumpluginTraditionalStaticGroupObjectClassProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginTraditionalStaticGroupObjectClassProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginTraditionalStaticGroupObjectClassProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginTraditionalStaticGroupObjectClassProp(val *EnumpluginTraditionalStaticGroupObjectClassProp) *NullableEnumpluginTraditionalStaticGroupObjectClassProp {
	return &NullableEnumpluginTraditionalStaticGroupObjectClassProp{value: val, isSet: true}
}

func (v NullableEnumpluginTraditionalStaticGroupObjectClassProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginTraditionalStaticGroupObjectClassProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
