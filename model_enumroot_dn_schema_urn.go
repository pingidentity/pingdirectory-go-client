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

// EnumrootDnSchemaUrn the model 'EnumrootDnSchemaUrn'
type EnumrootDnSchemaUrn string

// List of Enumroot-dnSchemaUrn
const (
	ENUMROOTDNSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ROOT_DN EnumrootDnSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:root-dn"
)

// All allowed values of EnumrootDnSchemaUrn enum
var AllowedEnumrootDnSchemaUrnEnumValues = []EnumrootDnSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:root-dn",
}

func (v *EnumrootDnSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumrootDnSchemaUrn(value)
	for _, existing := range AllowedEnumrootDnSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumrootDnSchemaUrn", value)
}

// NewEnumrootDnSchemaUrnFromValue returns a pointer to a valid EnumrootDnSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumrootDnSchemaUrnFromValue(v string) (*EnumrootDnSchemaUrn, error) {
	ev := EnumrootDnSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumrootDnSchemaUrn: valid values are %v", v, AllowedEnumrootDnSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumrootDnSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumrootDnSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumroot-dnSchemaUrn value
func (v EnumrootDnSchemaUrn) Ptr() *EnumrootDnSchemaUrn {
	return &v
}

type NullableEnumrootDnSchemaUrn struct {
	value *EnumrootDnSchemaUrn
	isSet bool
}

func (v NullableEnumrootDnSchemaUrn) Get() *EnumrootDnSchemaUrn {
	return v.value
}

func (v *NullableEnumrootDnSchemaUrn) Set(val *EnumrootDnSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumrootDnSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumrootDnSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumrootDnSchemaUrn(val *EnumrootDnSchemaUrn) *NullableEnumrootDnSchemaUrn {
	return &NullableEnumrootDnSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumrootDnSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumrootDnSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
