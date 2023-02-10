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

// EnumlocalDbVlvIndexSchemaUrn the model 'EnumlocalDbVlvIndexSchemaUrn'
type EnumlocalDbVlvIndexSchemaUrn string

// List of Enumlocal-db-vlv-indexSchemaUrn
const (
	ENUMLOCALDBVLVINDEXSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOCAL_DB_VLV_INDEX EnumlocalDbVlvIndexSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:local-db-vlv-index"
)

// All allowed values of EnumlocalDbVlvIndexSchemaUrn enum
var AllowedEnumlocalDbVlvIndexSchemaUrnEnumValues = []EnumlocalDbVlvIndexSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:local-db-vlv-index",
}

func (v *EnumlocalDbVlvIndexSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlocalDbVlvIndexSchemaUrn(value)
	for _, existing := range AllowedEnumlocalDbVlvIndexSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlocalDbVlvIndexSchemaUrn", value)
}

// NewEnumlocalDbVlvIndexSchemaUrnFromValue returns a pointer to a valid EnumlocalDbVlvIndexSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlocalDbVlvIndexSchemaUrnFromValue(v string) (*EnumlocalDbVlvIndexSchemaUrn, error) {
	ev := EnumlocalDbVlvIndexSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlocalDbVlvIndexSchemaUrn: valid values are %v", v, AllowedEnumlocalDbVlvIndexSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlocalDbVlvIndexSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumlocalDbVlvIndexSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlocal-db-vlv-indexSchemaUrn value
func (v EnumlocalDbVlvIndexSchemaUrn) Ptr() *EnumlocalDbVlvIndexSchemaUrn {
	return &v
}

type NullableEnumlocalDbVlvIndexSchemaUrn struct {
	value *EnumlocalDbVlvIndexSchemaUrn
	isSet bool
}

func (v NullableEnumlocalDbVlvIndexSchemaUrn) Get() *EnumlocalDbVlvIndexSchemaUrn {
	return v.value
}

func (v *NullableEnumlocalDbVlvIndexSchemaUrn) Set(val *EnumlocalDbVlvIndexSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlocalDbVlvIndexSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlocalDbVlvIndexSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlocalDbVlvIndexSchemaUrn(val *EnumlocalDbVlvIndexSchemaUrn) *NullableEnumlocalDbVlvIndexSchemaUrn {
	return &NullableEnumlocalDbVlvIndexSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumlocalDbVlvIndexSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlocalDbVlvIndexSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
