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

// EnumentryDnVirtualAttributeSchemaUrn the model 'EnumentryDnVirtualAttributeSchemaUrn'
type EnumentryDnVirtualAttributeSchemaUrn string

// List of Enumentry-dn-virtual-attributeSchemaUrn
const (
	ENUMENTRYDNVIRTUALATTRIBUTESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0VIRTUAL_ATTRIBUTEENTRY_DN EnumentryDnVirtualAttributeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:virtual-attribute:entry-dn"
)

// All allowed values of EnumentryDnVirtualAttributeSchemaUrn enum
var AllowedEnumentryDnVirtualAttributeSchemaUrnEnumValues = []EnumentryDnVirtualAttributeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:virtual-attribute:entry-dn",
}

func (v *EnumentryDnVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumentryDnVirtualAttributeSchemaUrn(value)
	for _, existing := range AllowedEnumentryDnVirtualAttributeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumentryDnVirtualAttributeSchemaUrn", value)
}

// NewEnumentryDnVirtualAttributeSchemaUrnFromValue returns a pointer to a valid EnumentryDnVirtualAttributeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumentryDnVirtualAttributeSchemaUrnFromValue(v string) (*EnumentryDnVirtualAttributeSchemaUrn, error) {
	ev := EnumentryDnVirtualAttributeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumentryDnVirtualAttributeSchemaUrn: valid values are %v", v, AllowedEnumentryDnVirtualAttributeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumentryDnVirtualAttributeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumentryDnVirtualAttributeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumentry-dn-virtual-attributeSchemaUrn value
func (v EnumentryDnVirtualAttributeSchemaUrn) Ptr() *EnumentryDnVirtualAttributeSchemaUrn {
	return &v
}

type NullableEnumentryDnVirtualAttributeSchemaUrn struct {
	value *EnumentryDnVirtualAttributeSchemaUrn
	isSet bool
}

func (v NullableEnumentryDnVirtualAttributeSchemaUrn) Get() *EnumentryDnVirtualAttributeSchemaUrn {
	return v.value
}

func (v *NullableEnumentryDnVirtualAttributeSchemaUrn) Set(val *EnumentryDnVirtualAttributeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumentryDnVirtualAttributeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumentryDnVirtualAttributeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumentryDnVirtualAttributeSchemaUrn(val *EnumentryDnVirtualAttributeSchemaUrn) *NullableEnumentryDnVirtualAttributeSchemaUrn {
	return &NullableEnumentryDnVirtualAttributeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumentryDnVirtualAttributeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumentryDnVirtualAttributeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
