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

// EnumnumSubordinatesVirtualAttributeSchemaUrn the model 'EnumnumSubordinatesVirtualAttributeSchemaUrn'
type EnumnumSubordinatesVirtualAttributeSchemaUrn string

// List of Enumnum-subordinates-virtual-attributeSchemaUrn
const (
	ENUMNUMSUBORDINATESVIRTUALATTRIBUTESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0VIRTUAL_ATTRIBUTENUM_SUBORDINATES EnumnumSubordinatesVirtualAttributeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:virtual-attribute:num-subordinates"
)

// All allowed values of EnumnumSubordinatesVirtualAttributeSchemaUrn enum
var AllowedEnumnumSubordinatesVirtualAttributeSchemaUrnEnumValues = []EnumnumSubordinatesVirtualAttributeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:virtual-attribute:num-subordinates",
}

func (v *EnumnumSubordinatesVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumnumSubordinatesVirtualAttributeSchemaUrn(value)
	for _, existing := range AllowedEnumnumSubordinatesVirtualAttributeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumnumSubordinatesVirtualAttributeSchemaUrn", value)
}

// NewEnumnumSubordinatesVirtualAttributeSchemaUrnFromValue returns a pointer to a valid EnumnumSubordinatesVirtualAttributeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumnumSubordinatesVirtualAttributeSchemaUrnFromValue(v string) (*EnumnumSubordinatesVirtualAttributeSchemaUrn, error) {
	ev := EnumnumSubordinatesVirtualAttributeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumnumSubordinatesVirtualAttributeSchemaUrn: valid values are %v", v, AllowedEnumnumSubordinatesVirtualAttributeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumnumSubordinatesVirtualAttributeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumnumSubordinatesVirtualAttributeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumnum-subordinates-virtual-attributeSchemaUrn value
func (v EnumnumSubordinatesVirtualAttributeSchemaUrn) Ptr() *EnumnumSubordinatesVirtualAttributeSchemaUrn {
	return &v
}

type NullableEnumnumSubordinatesVirtualAttributeSchemaUrn struct {
	value *EnumnumSubordinatesVirtualAttributeSchemaUrn
	isSet bool
}

func (v NullableEnumnumSubordinatesVirtualAttributeSchemaUrn) Get() *EnumnumSubordinatesVirtualAttributeSchemaUrn {
	return v.value
}

func (v *NullableEnumnumSubordinatesVirtualAttributeSchemaUrn) Set(val *EnumnumSubordinatesVirtualAttributeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumnumSubordinatesVirtualAttributeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumnumSubordinatesVirtualAttributeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumnumSubordinatesVirtualAttributeSchemaUrn(val *EnumnumSubordinatesVirtualAttributeSchemaUrn) *NullableEnumnumSubordinatesVirtualAttributeSchemaUrn {
	return &NullableEnumnumSubordinatesVirtualAttributeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumnumSubordinatesVirtualAttributeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumnumSubordinatesVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
