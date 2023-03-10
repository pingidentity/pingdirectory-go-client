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

// EnumaggregateIdentityMapperSchemaUrn the model 'EnumaggregateIdentityMapperSchemaUrn'
type EnumaggregateIdentityMapperSchemaUrn string

// List of Enumaggregate-identity-mapperSchemaUrn
const (
	ENUMAGGREGATEIDENTITYMAPPERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0IDENTITY_MAPPERAGGREGATE EnumaggregateIdentityMapperSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:identity-mapper:aggregate"
)

// All allowed values of EnumaggregateIdentityMapperSchemaUrn enum
var AllowedEnumaggregateIdentityMapperSchemaUrnEnumValues = []EnumaggregateIdentityMapperSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:identity-mapper:aggregate",
}

func (v *EnumaggregateIdentityMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumaggregateIdentityMapperSchemaUrn(value)
	for _, existing := range AllowedEnumaggregateIdentityMapperSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumaggregateIdentityMapperSchemaUrn", value)
}

// NewEnumaggregateIdentityMapperSchemaUrnFromValue returns a pointer to a valid EnumaggregateIdentityMapperSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumaggregateIdentityMapperSchemaUrnFromValue(v string) (*EnumaggregateIdentityMapperSchemaUrn, error) {
	ev := EnumaggregateIdentityMapperSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumaggregateIdentityMapperSchemaUrn: valid values are %v", v, AllowedEnumaggregateIdentityMapperSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumaggregateIdentityMapperSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumaggregateIdentityMapperSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumaggregate-identity-mapperSchemaUrn value
func (v EnumaggregateIdentityMapperSchemaUrn) Ptr() *EnumaggregateIdentityMapperSchemaUrn {
	return &v
}

type NullableEnumaggregateIdentityMapperSchemaUrn struct {
	value *EnumaggregateIdentityMapperSchemaUrn
	isSet bool
}

func (v NullableEnumaggregateIdentityMapperSchemaUrn) Get() *EnumaggregateIdentityMapperSchemaUrn {
	return v.value
}

func (v *NullableEnumaggregateIdentityMapperSchemaUrn) Set(val *EnumaggregateIdentityMapperSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumaggregateIdentityMapperSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumaggregateIdentityMapperSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumaggregateIdentityMapperSchemaUrn(val *EnumaggregateIdentityMapperSchemaUrn) *NullableEnumaggregateIdentityMapperSchemaUrn {
	return &NullableEnumaggregateIdentityMapperSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumaggregateIdentityMapperSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumaggregateIdentityMapperSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
