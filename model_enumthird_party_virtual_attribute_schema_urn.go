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

// EnumthirdPartyVirtualAttributeSchemaUrn the model 'EnumthirdPartyVirtualAttributeSchemaUrn'
type EnumthirdPartyVirtualAttributeSchemaUrn string

// List of Enumthird-party-virtual-attributeSchemaUrn
const (
	ENUMTHIRDPARTYVIRTUALATTRIBUTESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0VIRTUAL_ATTRIBUTETHIRD_PARTY EnumthirdPartyVirtualAttributeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:virtual-attribute:third-party"
)

// All allowed values of EnumthirdPartyVirtualAttributeSchemaUrn enum
var AllowedEnumthirdPartyVirtualAttributeSchemaUrnEnumValues = []EnumthirdPartyVirtualAttributeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:virtual-attribute:third-party",
}

func (v *EnumthirdPartyVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumthirdPartyVirtualAttributeSchemaUrn(value)
	for _, existing := range AllowedEnumthirdPartyVirtualAttributeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumthirdPartyVirtualAttributeSchemaUrn", value)
}

// NewEnumthirdPartyVirtualAttributeSchemaUrnFromValue returns a pointer to a valid EnumthirdPartyVirtualAttributeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumthirdPartyVirtualAttributeSchemaUrnFromValue(v string) (*EnumthirdPartyVirtualAttributeSchemaUrn, error) {
	ev := EnumthirdPartyVirtualAttributeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumthirdPartyVirtualAttributeSchemaUrn: valid values are %v", v, AllowedEnumthirdPartyVirtualAttributeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumthirdPartyVirtualAttributeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumthirdPartyVirtualAttributeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumthird-party-virtual-attributeSchemaUrn value
func (v EnumthirdPartyVirtualAttributeSchemaUrn) Ptr() *EnumthirdPartyVirtualAttributeSchemaUrn {
	return &v
}

type NullableEnumthirdPartyVirtualAttributeSchemaUrn struct {
	value *EnumthirdPartyVirtualAttributeSchemaUrn
	isSet bool
}

func (v NullableEnumthirdPartyVirtualAttributeSchemaUrn) Get() *EnumthirdPartyVirtualAttributeSchemaUrn {
	return v.value
}

func (v *NullableEnumthirdPartyVirtualAttributeSchemaUrn) Set(val *EnumthirdPartyVirtualAttributeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumthirdPartyVirtualAttributeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumthirdPartyVirtualAttributeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumthirdPartyVirtualAttributeSchemaUrn(val *EnumthirdPartyVirtualAttributeSchemaUrn) *NullableEnumthirdPartyVirtualAttributeSchemaUrn {
	return &NullableEnumthirdPartyVirtualAttributeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumthirdPartyVirtualAttributeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumthirdPartyVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

