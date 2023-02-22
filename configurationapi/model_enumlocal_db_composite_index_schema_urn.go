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

// EnumlocalDbCompositeIndexSchemaUrn the model 'EnumlocalDbCompositeIndexSchemaUrn'
type EnumlocalDbCompositeIndexSchemaUrn string

// List of Enumlocal-db-composite-indexSchemaUrn
const (
	ENUMLOCALDBCOMPOSITEINDEXSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOCAL_DB_COMPOSITE_INDEX EnumlocalDbCompositeIndexSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:local-db-composite-index"
)

// All allowed values of EnumlocalDbCompositeIndexSchemaUrn enum
var AllowedEnumlocalDbCompositeIndexSchemaUrnEnumValues = []EnumlocalDbCompositeIndexSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:local-db-composite-index",
}

func (v *EnumlocalDbCompositeIndexSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlocalDbCompositeIndexSchemaUrn(value)
	for _, existing := range AllowedEnumlocalDbCompositeIndexSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlocalDbCompositeIndexSchemaUrn", value)
}

// NewEnumlocalDbCompositeIndexSchemaUrnFromValue returns a pointer to a valid EnumlocalDbCompositeIndexSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlocalDbCompositeIndexSchemaUrnFromValue(v string) (*EnumlocalDbCompositeIndexSchemaUrn, error) {
	ev := EnumlocalDbCompositeIndexSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlocalDbCompositeIndexSchemaUrn: valid values are %v", v, AllowedEnumlocalDbCompositeIndexSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlocalDbCompositeIndexSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumlocalDbCompositeIndexSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlocal-db-composite-indexSchemaUrn value
func (v EnumlocalDbCompositeIndexSchemaUrn) Ptr() *EnumlocalDbCompositeIndexSchemaUrn {
	return &v
}

type NullableEnumlocalDbCompositeIndexSchemaUrn struct {
	value *EnumlocalDbCompositeIndexSchemaUrn
	isSet bool
}

func (v NullableEnumlocalDbCompositeIndexSchemaUrn) Get() *EnumlocalDbCompositeIndexSchemaUrn {
	return v.value
}

func (v *NullableEnumlocalDbCompositeIndexSchemaUrn) Set(val *EnumlocalDbCompositeIndexSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlocalDbCompositeIndexSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlocalDbCompositeIndexSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlocalDbCompositeIndexSchemaUrn(val *EnumlocalDbCompositeIndexSchemaUrn) *NullableEnumlocalDbCompositeIndexSchemaUrn {
	return &NullableEnumlocalDbCompositeIndexSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumlocalDbCompositeIndexSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlocalDbCompositeIndexSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}