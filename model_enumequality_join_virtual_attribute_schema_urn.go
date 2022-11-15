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

// EnumequalityJoinVirtualAttributeSchemaUrn the model 'EnumequalityJoinVirtualAttributeSchemaUrn'
type EnumequalityJoinVirtualAttributeSchemaUrn string

// List of Enumequality-join-virtual-attributeSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0VIRTUAL_ATTRIBUTEEQUALITY_JOIN EnumequalityJoinVirtualAttributeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:virtual-attribute:equality-join"
)

// All allowed values of EnumequalityJoinVirtualAttributeSchemaUrn enum
var AllowedEnumequalityJoinVirtualAttributeSchemaUrnEnumValues = []EnumequalityJoinVirtualAttributeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:virtual-attribute:equality-join",
}

func (v *EnumequalityJoinVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumequalityJoinVirtualAttributeSchemaUrn(value)
	for _, existing := range AllowedEnumequalityJoinVirtualAttributeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumequalityJoinVirtualAttributeSchemaUrn", value)
}

// NewEnumequalityJoinVirtualAttributeSchemaUrnFromValue returns a pointer to a valid EnumequalityJoinVirtualAttributeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumequalityJoinVirtualAttributeSchemaUrnFromValue(v string) (*EnumequalityJoinVirtualAttributeSchemaUrn, error) {
	ev := EnumequalityJoinVirtualAttributeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumequalityJoinVirtualAttributeSchemaUrn: valid values are %v", v, AllowedEnumequalityJoinVirtualAttributeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumequalityJoinVirtualAttributeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumequalityJoinVirtualAttributeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumequality-join-virtual-attributeSchemaUrn value
func (v EnumequalityJoinVirtualAttributeSchemaUrn) Ptr() *EnumequalityJoinVirtualAttributeSchemaUrn {
	return &v
}

type NullableEnumequalityJoinVirtualAttributeSchemaUrn struct {
	value *EnumequalityJoinVirtualAttributeSchemaUrn
	isSet bool
}

func (v NullableEnumequalityJoinVirtualAttributeSchemaUrn) Get() *EnumequalityJoinVirtualAttributeSchemaUrn {
	return v.value
}

func (v *NullableEnumequalityJoinVirtualAttributeSchemaUrn) Set(val *EnumequalityJoinVirtualAttributeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumequalityJoinVirtualAttributeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumequalityJoinVirtualAttributeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumequalityJoinVirtualAttributeSchemaUrn(val *EnumequalityJoinVirtualAttributeSchemaUrn) *NullableEnumequalityJoinVirtualAttributeSchemaUrn {
	return &NullableEnumequalityJoinVirtualAttributeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumequalityJoinVirtualAttributeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumequalityJoinVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

