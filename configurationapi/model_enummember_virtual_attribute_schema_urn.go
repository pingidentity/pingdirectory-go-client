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

// EnummemberVirtualAttributeSchemaUrn the model 'EnummemberVirtualAttributeSchemaUrn'
type EnummemberVirtualAttributeSchemaUrn string

// List of Enummember-virtual-attributeSchemaUrn
const (
	ENUMMEMBERVIRTUALATTRIBUTESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0VIRTUAL_ATTRIBUTEMEMBER EnummemberVirtualAttributeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:virtual-attribute:member"
)

// All allowed values of EnummemberVirtualAttributeSchemaUrn enum
var AllowedEnummemberVirtualAttributeSchemaUrnEnumValues = []EnummemberVirtualAttributeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:virtual-attribute:member",
}

func (v *EnummemberVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnummemberVirtualAttributeSchemaUrn(value)
	for _, existing := range AllowedEnummemberVirtualAttributeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnummemberVirtualAttributeSchemaUrn", value)
}

// NewEnummemberVirtualAttributeSchemaUrnFromValue returns a pointer to a valid EnummemberVirtualAttributeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnummemberVirtualAttributeSchemaUrnFromValue(v string) (*EnummemberVirtualAttributeSchemaUrn, error) {
	ev := EnummemberVirtualAttributeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnummemberVirtualAttributeSchemaUrn: valid values are %v", v, AllowedEnummemberVirtualAttributeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnummemberVirtualAttributeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnummemberVirtualAttributeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enummember-virtual-attributeSchemaUrn value
func (v EnummemberVirtualAttributeSchemaUrn) Ptr() *EnummemberVirtualAttributeSchemaUrn {
	return &v
}

type NullableEnummemberVirtualAttributeSchemaUrn struct {
	value *EnummemberVirtualAttributeSchemaUrn
	isSet bool
}

func (v NullableEnummemberVirtualAttributeSchemaUrn) Get() *EnummemberVirtualAttributeSchemaUrn {
	return v.value
}

func (v *NullableEnummemberVirtualAttributeSchemaUrn) Set(val *EnummemberVirtualAttributeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnummemberVirtualAttributeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnummemberVirtualAttributeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnummemberVirtualAttributeSchemaUrn(val *EnummemberVirtualAttributeSchemaUrn) *NullableEnummemberVirtualAttributeSchemaUrn {
	return &NullableEnummemberVirtualAttributeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnummemberVirtualAttributeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnummemberVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
