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

// EnumentryChecksumVirtualAttributeSchemaUrn the model 'EnumentryChecksumVirtualAttributeSchemaUrn'
type EnumentryChecksumVirtualAttributeSchemaUrn string

// List of Enumentry-checksum-virtual-attributeSchemaUrn
const (
	ENUMENTRYCHECKSUMVIRTUALATTRIBUTESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0VIRTUAL_ATTRIBUTEENTRY_CHECKSUM EnumentryChecksumVirtualAttributeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:virtual-attribute:entry-checksum"
)

// All allowed values of EnumentryChecksumVirtualAttributeSchemaUrn enum
var AllowedEnumentryChecksumVirtualAttributeSchemaUrnEnumValues = []EnumentryChecksumVirtualAttributeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:virtual-attribute:entry-checksum",
}

func (v *EnumentryChecksumVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumentryChecksumVirtualAttributeSchemaUrn(value)
	for _, existing := range AllowedEnumentryChecksumVirtualAttributeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumentryChecksumVirtualAttributeSchemaUrn", value)
}

// NewEnumentryChecksumVirtualAttributeSchemaUrnFromValue returns a pointer to a valid EnumentryChecksumVirtualAttributeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumentryChecksumVirtualAttributeSchemaUrnFromValue(v string) (*EnumentryChecksumVirtualAttributeSchemaUrn, error) {
	ev := EnumentryChecksumVirtualAttributeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumentryChecksumVirtualAttributeSchemaUrn: valid values are %v", v, AllowedEnumentryChecksumVirtualAttributeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumentryChecksumVirtualAttributeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumentryChecksumVirtualAttributeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumentry-checksum-virtual-attributeSchemaUrn value
func (v EnumentryChecksumVirtualAttributeSchemaUrn) Ptr() *EnumentryChecksumVirtualAttributeSchemaUrn {
	return &v
}

type NullableEnumentryChecksumVirtualAttributeSchemaUrn struct {
	value *EnumentryChecksumVirtualAttributeSchemaUrn
	isSet bool
}

func (v NullableEnumentryChecksumVirtualAttributeSchemaUrn) Get() *EnumentryChecksumVirtualAttributeSchemaUrn {
	return v.value
}

func (v *NullableEnumentryChecksumVirtualAttributeSchemaUrn) Set(val *EnumentryChecksumVirtualAttributeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumentryChecksumVirtualAttributeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumentryChecksumVirtualAttributeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumentryChecksumVirtualAttributeSchemaUrn(val *EnumentryChecksumVirtualAttributeSchemaUrn) *NullableEnumentryChecksumVirtualAttributeSchemaUrn {
	return &NullableEnumentryChecksumVirtualAttributeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumentryChecksumVirtualAttributeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumentryChecksumVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

