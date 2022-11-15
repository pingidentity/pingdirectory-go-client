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

// EnumsubschemaSubentryVirtualAttributeSchemaUrn the model 'EnumsubschemaSubentryVirtualAttributeSchemaUrn'
type EnumsubschemaSubentryVirtualAttributeSchemaUrn string

// List of Enumsubschema-subentry-virtual-attributeSchemaUrn
const (
	ENUMSUBSCHEMASUBENTRYVIRTUALATTRIBUTESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0VIRTUAL_ATTRIBUTESUBSCHEMA_SUBENTRY EnumsubschemaSubentryVirtualAttributeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:virtual-attribute:subschema-subentry"
)

// All allowed values of EnumsubschemaSubentryVirtualAttributeSchemaUrn enum
var AllowedEnumsubschemaSubentryVirtualAttributeSchemaUrnEnumValues = []EnumsubschemaSubentryVirtualAttributeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:virtual-attribute:subschema-subentry",
}

func (v *EnumsubschemaSubentryVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsubschemaSubentryVirtualAttributeSchemaUrn(value)
	for _, existing := range AllowedEnumsubschemaSubentryVirtualAttributeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsubschemaSubentryVirtualAttributeSchemaUrn", value)
}

// NewEnumsubschemaSubentryVirtualAttributeSchemaUrnFromValue returns a pointer to a valid EnumsubschemaSubentryVirtualAttributeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsubschemaSubentryVirtualAttributeSchemaUrnFromValue(v string) (*EnumsubschemaSubentryVirtualAttributeSchemaUrn, error) {
	ev := EnumsubschemaSubentryVirtualAttributeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsubschemaSubentryVirtualAttributeSchemaUrn: valid values are %v", v, AllowedEnumsubschemaSubentryVirtualAttributeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsubschemaSubentryVirtualAttributeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumsubschemaSubentryVirtualAttributeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsubschema-subentry-virtual-attributeSchemaUrn value
func (v EnumsubschemaSubentryVirtualAttributeSchemaUrn) Ptr() *EnumsubschemaSubentryVirtualAttributeSchemaUrn {
	return &v
}

type NullableEnumsubschemaSubentryVirtualAttributeSchemaUrn struct {
	value *EnumsubschemaSubentryVirtualAttributeSchemaUrn
	isSet bool
}

func (v NullableEnumsubschemaSubentryVirtualAttributeSchemaUrn) Get() *EnumsubschemaSubentryVirtualAttributeSchemaUrn {
	return v.value
}

func (v *NullableEnumsubschemaSubentryVirtualAttributeSchemaUrn) Set(val *EnumsubschemaSubentryVirtualAttributeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsubschemaSubentryVirtualAttributeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsubschemaSubentryVirtualAttributeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsubschemaSubentryVirtualAttributeSchemaUrn(val *EnumsubschemaSubentryVirtualAttributeSchemaUrn) *NullableEnumsubschemaSubentryVirtualAttributeSchemaUrn {
	return &NullableEnumsubschemaSubentryVirtualAttributeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumsubschemaSubentryVirtualAttributeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsubschemaSubentryVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

