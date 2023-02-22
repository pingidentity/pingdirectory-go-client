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

// EnumaccessLogFieldMappingSchemaUrn the model 'EnumaccessLogFieldMappingSchemaUrn'
type EnumaccessLogFieldMappingSchemaUrn string

// List of Enumaccess-log-field-mappingSchemaUrn
const (
	ENUMACCESSLOGFIELDMAPPINGSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0LOG_FIELD_MAPPINGACCESS EnumaccessLogFieldMappingSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:log-field-mapping:access"
)

// All allowed values of EnumaccessLogFieldMappingSchemaUrn enum
var AllowedEnumaccessLogFieldMappingSchemaUrnEnumValues = []EnumaccessLogFieldMappingSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:log-field-mapping:access",
}

func (v *EnumaccessLogFieldMappingSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumaccessLogFieldMappingSchemaUrn(value)
	for _, existing := range AllowedEnumaccessLogFieldMappingSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumaccessLogFieldMappingSchemaUrn", value)
}

// NewEnumaccessLogFieldMappingSchemaUrnFromValue returns a pointer to a valid EnumaccessLogFieldMappingSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumaccessLogFieldMappingSchemaUrnFromValue(v string) (*EnumaccessLogFieldMappingSchemaUrn, error) {
	ev := EnumaccessLogFieldMappingSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumaccessLogFieldMappingSchemaUrn: valid values are %v", v, AllowedEnumaccessLogFieldMappingSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumaccessLogFieldMappingSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumaccessLogFieldMappingSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumaccess-log-field-mappingSchemaUrn value
func (v EnumaccessLogFieldMappingSchemaUrn) Ptr() *EnumaccessLogFieldMappingSchemaUrn {
	return &v
}

type NullableEnumaccessLogFieldMappingSchemaUrn struct {
	value *EnumaccessLogFieldMappingSchemaUrn
	isSet bool
}

func (v NullableEnumaccessLogFieldMappingSchemaUrn) Get() *EnumaccessLogFieldMappingSchemaUrn {
	return v.value
}

func (v *NullableEnumaccessLogFieldMappingSchemaUrn) Set(val *EnumaccessLogFieldMappingSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumaccessLogFieldMappingSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumaccessLogFieldMappingSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumaccessLogFieldMappingSchemaUrn(val *EnumaccessLogFieldMappingSchemaUrn) *NullableEnumaccessLogFieldMappingSchemaUrn {
	return &NullableEnumaccessLogFieldMappingSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumaccessLogFieldMappingSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumaccessLogFieldMappingSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}