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

// EnumerrorLogFieldMappingSchemaUrn the model 'EnumerrorLogFieldMappingSchemaUrn'
type EnumerrorLogFieldMappingSchemaUrn string

// List of Enumerror-log-field-mappingSchemaUrn
const (
	ENUMERRORLOGFIELDMAPPINGSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_FIELD_MAPPINGERROR EnumerrorLogFieldMappingSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-field-mapping:error"
)

// All allowed values of EnumerrorLogFieldMappingSchemaUrn enum
var AllowedEnumerrorLogFieldMappingSchemaUrnEnumValues = []EnumerrorLogFieldMappingSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-field-mapping:error",
}

func (v *EnumerrorLogFieldMappingSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumerrorLogFieldMappingSchemaUrn(value)
	for _, existing := range AllowedEnumerrorLogFieldMappingSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumerrorLogFieldMappingSchemaUrn", value)
}

// NewEnumerrorLogFieldMappingSchemaUrnFromValue returns a pointer to a valid EnumerrorLogFieldMappingSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumerrorLogFieldMappingSchemaUrnFromValue(v string) (*EnumerrorLogFieldMappingSchemaUrn, error) {
	ev := EnumerrorLogFieldMappingSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumerrorLogFieldMappingSchemaUrn: valid values are %v", v, AllowedEnumerrorLogFieldMappingSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumerrorLogFieldMappingSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumerrorLogFieldMappingSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumerror-log-field-mappingSchemaUrn value
func (v EnumerrorLogFieldMappingSchemaUrn) Ptr() *EnumerrorLogFieldMappingSchemaUrn {
	return &v
}

type NullableEnumerrorLogFieldMappingSchemaUrn struct {
	value *EnumerrorLogFieldMappingSchemaUrn
	isSet bool
}

func (v NullableEnumerrorLogFieldMappingSchemaUrn) Get() *EnumerrorLogFieldMappingSchemaUrn {
	return v.value
}

func (v *NullableEnumerrorLogFieldMappingSchemaUrn) Set(val *EnumerrorLogFieldMappingSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumerrorLogFieldMappingSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumerrorLogFieldMappingSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumerrorLogFieldMappingSchemaUrn(val *EnumerrorLogFieldMappingSchemaUrn) *NullableEnumerrorLogFieldMappingSchemaUrn {
	return &NullableEnumerrorLogFieldMappingSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumerrorLogFieldMappingSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumerrorLogFieldMappingSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
