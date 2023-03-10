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

// EnumcustomVirtualAttributeSchemaUrn the model 'EnumcustomVirtualAttributeSchemaUrn'
type EnumcustomVirtualAttributeSchemaUrn string

// List of Enumcustom-virtual-attributeSchemaUrn
const (
	ENUMCUSTOMVIRTUALATTRIBUTESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0VIRTUAL_ATTRIBUTECUSTOM EnumcustomVirtualAttributeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:virtual-attribute:custom"
)

// All allowed values of EnumcustomVirtualAttributeSchemaUrn enum
var AllowedEnumcustomVirtualAttributeSchemaUrnEnumValues = []EnumcustomVirtualAttributeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:virtual-attribute:custom",
}

func (v *EnumcustomVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcustomVirtualAttributeSchemaUrn(value)
	for _, existing := range AllowedEnumcustomVirtualAttributeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcustomVirtualAttributeSchemaUrn", value)
}

// NewEnumcustomVirtualAttributeSchemaUrnFromValue returns a pointer to a valid EnumcustomVirtualAttributeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcustomVirtualAttributeSchemaUrnFromValue(v string) (*EnumcustomVirtualAttributeSchemaUrn, error) {
	ev := EnumcustomVirtualAttributeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcustomVirtualAttributeSchemaUrn: valid values are %v", v, AllowedEnumcustomVirtualAttributeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcustomVirtualAttributeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcustomVirtualAttributeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumcustom-virtual-attributeSchemaUrn value
func (v EnumcustomVirtualAttributeSchemaUrn) Ptr() *EnumcustomVirtualAttributeSchemaUrn {
	return &v
}

type NullableEnumcustomVirtualAttributeSchemaUrn struct {
	value *EnumcustomVirtualAttributeSchemaUrn
	isSet bool
}

func (v NullableEnumcustomVirtualAttributeSchemaUrn) Get() *EnumcustomVirtualAttributeSchemaUrn {
	return v.value
}

func (v *NullableEnumcustomVirtualAttributeSchemaUrn) Set(val *EnumcustomVirtualAttributeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcustomVirtualAttributeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcustomVirtualAttributeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcustomVirtualAttributeSchemaUrn(val *EnumcustomVirtualAttributeSchemaUrn) *NullableEnumcustomVirtualAttributeSchemaUrn {
	return &NullableEnumcustomVirtualAttributeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcustomVirtualAttributeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcustomVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
