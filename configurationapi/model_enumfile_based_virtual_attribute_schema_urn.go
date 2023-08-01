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

// EnumfileBasedVirtualAttributeSchemaUrn the model 'EnumfileBasedVirtualAttributeSchemaUrn'
type EnumfileBasedVirtualAttributeSchemaUrn string

// List of Enumfile-based-virtual-attributeSchemaUrn
const (
	ENUMFILEBASEDVIRTUALATTRIBUTESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0VIRTUAL_ATTRIBUTEFILE_BASED EnumfileBasedVirtualAttributeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:virtual-attribute:file-based"
)

// All allowed values of EnumfileBasedVirtualAttributeSchemaUrn enum
var AllowedEnumfileBasedVirtualAttributeSchemaUrnEnumValues = []EnumfileBasedVirtualAttributeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:virtual-attribute:file-based",
}

func (v *EnumfileBasedVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumfileBasedVirtualAttributeSchemaUrn(value)
	for _, existing := range AllowedEnumfileBasedVirtualAttributeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumfileBasedVirtualAttributeSchemaUrn", value)
}

// NewEnumfileBasedVirtualAttributeSchemaUrnFromValue returns a pointer to a valid EnumfileBasedVirtualAttributeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumfileBasedVirtualAttributeSchemaUrnFromValue(v string) (*EnumfileBasedVirtualAttributeSchemaUrn, error) {
	ev := EnumfileBasedVirtualAttributeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumfileBasedVirtualAttributeSchemaUrn: valid values are %v", v, AllowedEnumfileBasedVirtualAttributeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumfileBasedVirtualAttributeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumfileBasedVirtualAttributeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumfile-based-virtual-attributeSchemaUrn value
func (v EnumfileBasedVirtualAttributeSchemaUrn) Ptr() *EnumfileBasedVirtualAttributeSchemaUrn {
	return &v
}

type NullableEnumfileBasedVirtualAttributeSchemaUrn struct {
	value *EnumfileBasedVirtualAttributeSchemaUrn
	isSet bool
}

func (v NullableEnumfileBasedVirtualAttributeSchemaUrn) Get() *EnumfileBasedVirtualAttributeSchemaUrn {
	return v.value
}

func (v *NullableEnumfileBasedVirtualAttributeSchemaUrn) Set(val *EnumfileBasedVirtualAttributeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumfileBasedVirtualAttributeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumfileBasedVirtualAttributeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumfileBasedVirtualAttributeSchemaUrn(val *EnumfileBasedVirtualAttributeSchemaUrn) *NullableEnumfileBasedVirtualAttributeSchemaUrn {
	return &NullableEnumfileBasedVirtualAttributeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumfileBasedVirtualAttributeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumfileBasedVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
