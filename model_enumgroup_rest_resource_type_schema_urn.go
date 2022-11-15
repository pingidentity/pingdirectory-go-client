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

// EnumgroupRestResourceTypeSchemaUrn the model 'EnumgroupRestResourceTypeSchemaUrn'
type EnumgroupRestResourceTypeSchemaUrn string

// List of Enumgroup-rest-resource-typeSchemaUrn
const (
	ENUMGROUPRESTRESOURCETYPESCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0REST_RESOURCE_TYPEGROUP EnumgroupRestResourceTypeSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:rest-resource-type:group"
)

// All allowed values of EnumgroupRestResourceTypeSchemaUrn enum
var AllowedEnumgroupRestResourceTypeSchemaUrnEnumValues = []EnumgroupRestResourceTypeSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:rest-resource-type:group",
}

func (v *EnumgroupRestResourceTypeSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumgroupRestResourceTypeSchemaUrn(value)
	for _, existing := range AllowedEnumgroupRestResourceTypeSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumgroupRestResourceTypeSchemaUrn", value)
}

// NewEnumgroupRestResourceTypeSchemaUrnFromValue returns a pointer to a valid EnumgroupRestResourceTypeSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumgroupRestResourceTypeSchemaUrnFromValue(v string) (*EnumgroupRestResourceTypeSchemaUrn, error) {
	ev := EnumgroupRestResourceTypeSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumgroupRestResourceTypeSchemaUrn: valid values are %v", v, AllowedEnumgroupRestResourceTypeSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumgroupRestResourceTypeSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumgroupRestResourceTypeSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumgroup-rest-resource-typeSchemaUrn value
func (v EnumgroupRestResourceTypeSchemaUrn) Ptr() *EnumgroupRestResourceTypeSchemaUrn {
	return &v
}

type NullableEnumgroupRestResourceTypeSchemaUrn struct {
	value *EnumgroupRestResourceTypeSchemaUrn
	isSet bool
}

func (v NullableEnumgroupRestResourceTypeSchemaUrn) Get() *EnumgroupRestResourceTypeSchemaUrn {
	return v.value
}

func (v *NullableEnumgroupRestResourceTypeSchemaUrn) Set(val *EnumgroupRestResourceTypeSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumgroupRestResourceTypeSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumgroupRestResourceTypeSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumgroupRestResourceTypeSchemaUrn(val *EnumgroupRestResourceTypeSchemaUrn) *NullableEnumgroupRestResourceTypeSchemaUrn {
	return &NullableEnumgroupRestResourceTypeSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumgroupRestResourceTypeSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumgroupRestResourceTypeSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

