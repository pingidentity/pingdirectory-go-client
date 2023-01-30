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

// EnumuserDefinedVirtualAttributeSchemaUrn the model 'EnumuserDefinedVirtualAttributeSchemaUrn'
type EnumuserDefinedVirtualAttributeSchemaUrn string

// List of Enumuser-defined-virtual-attributeSchemaUrn
const (
	ENUMUSERDEFINEDVIRTUALATTRIBUTESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0VIRTUAL_ATTRIBUTEUSER_DEFINED EnumuserDefinedVirtualAttributeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:virtual-attribute:user-defined"
)

// All allowed values of EnumuserDefinedVirtualAttributeSchemaUrn enum
var AllowedEnumuserDefinedVirtualAttributeSchemaUrnEnumValues = []EnumuserDefinedVirtualAttributeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:virtual-attribute:user-defined",
}

func (v *EnumuserDefinedVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumuserDefinedVirtualAttributeSchemaUrn(value)
	for _, existing := range AllowedEnumuserDefinedVirtualAttributeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumuserDefinedVirtualAttributeSchemaUrn", value)
}

// NewEnumuserDefinedVirtualAttributeSchemaUrnFromValue returns a pointer to a valid EnumuserDefinedVirtualAttributeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumuserDefinedVirtualAttributeSchemaUrnFromValue(v string) (*EnumuserDefinedVirtualAttributeSchemaUrn, error) {
	ev := EnumuserDefinedVirtualAttributeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumuserDefinedVirtualAttributeSchemaUrn: valid values are %v", v, AllowedEnumuserDefinedVirtualAttributeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumuserDefinedVirtualAttributeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumuserDefinedVirtualAttributeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumuser-defined-virtual-attributeSchemaUrn value
func (v EnumuserDefinedVirtualAttributeSchemaUrn) Ptr() *EnumuserDefinedVirtualAttributeSchemaUrn {
	return &v
}

type NullableEnumuserDefinedVirtualAttributeSchemaUrn struct {
	value *EnumuserDefinedVirtualAttributeSchemaUrn
	isSet bool
}

func (v NullableEnumuserDefinedVirtualAttributeSchemaUrn) Get() *EnumuserDefinedVirtualAttributeSchemaUrn {
	return v.value
}

func (v *NullableEnumuserDefinedVirtualAttributeSchemaUrn) Set(val *EnumuserDefinedVirtualAttributeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumuserDefinedVirtualAttributeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumuserDefinedVirtualAttributeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumuserDefinedVirtualAttributeSchemaUrn(val *EnumuserDefinedVirtualAttributeSchemaUrn) *NullableEnumuserDefinedVirtualAttributeSchemaUrn {
	return &NullableEnumuserDefinedVirtualAttributeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumuserDefinedVirtualAttributeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumuserDefinedVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
