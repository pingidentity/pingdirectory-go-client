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

// EnumidentifyReferencesVirtualAttributeSchemaUrn the model 'EnumidentifyReferencesVirtualAttributeSchemaUrn'
type EnumidentifyReferencesVirtualAttributeSchemaUrn string

// List of Enumidentify-references-virtual-attributeSchemaUrn
const (
	ENUMIDENTIFYREFERENCESVIRTUALATTRIBUTESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0VIRTUAL_ATTRIBUTEIDENTIFY_REFERENCES EnumidentifyReferencesVirtualAttributeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:virtual-attribute:identify-references"
)

// All allowed values of EnumidentifyReferencesVirtualAttributeSchemaUrn enum
var AllowedEnumidentifyReferencesVirtualAttributeSchemaUrnEnumValues = []EnumidentifyReferencesVirtualAttributeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:virtual-attribute:identify-references",
}

func (v *EnumidentifyReferencesVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumidentifyReferencesVirtualAttributeSchemaUrn(value)
	for _, existing := range AllowedEnumidentifyReferencesVirtualAttributeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumidentifyReferencesVirtualAttributeSchemaUrn", value)
}

// NewEnumidentifyReferencesVirtualAttributeSchemaUrnFromValue returns a pointer to a valid EnumidentifyReferencesVirtualAttributeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumidentifyReferencesVirtualAttributeSchemaUrnFromValue(v string) (*EnumidentifyReferencesVirtualAttributeSchemaUrn, error) {
	ev := EnumidentifyReferencesVirtualAttributeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumidentifyReferencesVirtualAttributeSchemaUrn: valid values are %v", v, AllowedEnumidentifyReferencesVirtualAttributeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumidentifyReferencesVirtualAttributeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumidentifyReferencesVirtualAttributeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumidentify-references-virtual-attributeSchemaUrn value
func (v EnumidentifyReferencesVirtualAttributeSchemaUrn) Ptr() *EnumidentifyReferencesVirtualAttributeSchemaUrn {
	return &v
}

type NullableEnumidentifyReferencesVirtualAttributeSchemaUrn struct {
	value *EnumidentifyReferencesVirtualAttributeSchemaUrn
	isSet bool
}

func (v NullableEnumidentifyReferencesVirtualAttributeSchemaUrn) Get() *EnumidentifyReferencesVirtualAttributeSchemaUrn {
	return v.value
}

func (v *NullableEnumidentifyReferencesVirtualAttributeSchemaUrn) Set(val *EnumidentifyReferencesVirtualAttributeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumidentifyReferencesVirtualAttributeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumidentifyReferencesVirtualAttributeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumidentifyReferencesVirtualAttributeSchemaUrn(val *EnumidentifyReferencesVirtualAttributeSchemaUrn) *NullableEnumidentifyReferencesVirtualAttributeSchemaUrn {
	return &NullableEnumidentifyReferencesVirtualAttributeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumidentifyReferencesVirtualAttributeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumidentifyReferencesVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

