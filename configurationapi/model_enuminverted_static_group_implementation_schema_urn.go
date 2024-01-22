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

// EnuminvertedStaticGroupImplementationSchemaUrn the model 'EnuminvertedStaticGroupImplementationSchemaUrn'
type EnuminvertedStaticGroupImplementationSchemaUrn string

// List of Enuminverted-static-group-implementationSchemaUrn
const (
	ENUMINVERTEDSTATICGROUPIMPLEMENTATIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0GROUP_IMPLEMENTATIONINVERTED_STATIC EnuminvertedStaticGroupImplementationSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:group-implementation:inverted-static"
)

// All allowed values of EnuminvertedStaticGroupImplementationSchemaUrn enum
var AllowedEnuminvertedStaticGroupImplementationSchemaUrnEnumValues = []EnuminvertedStaticGroupImplementationSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:group-implementation:inverted-static",
}

func (v *EnuminvertedStaticGroupImplementationSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnuminvertedStaticGroupImplementationSchemaUrn(value)
	for _, existing := range AllowedEnuminvertedStaticGroupImplementationSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnuminvertedStaticGroupImplementationSchemaUrn", value)
}

// NewEnuminvertedStaticGroupImplementationSchemaUrnFromValue returns a pointer to a valid EnuminvertedStaticGroupImplementationSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnuminvertedStaticGroupImplementationSchemaUrnFromValue(v string) (*EnuminvertedStaticGroupImplementationSchemaUrn, error) {
	ev := EnuminvertedStaticGroupImplementationSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnuminvertedStaticGroupImplementationSchemaUrn: valid values are %v", v, AllowedEnuminvertedStaticGroupImplementationSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnuminvertedStaticGroupImplementationSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnuminvertedStaticGroupImplementationSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enuminverted-static-group-implementationSchemaUrn value
func (v EnuminvertedStaticGroupImplementationSchemaUrn) Ptr() *EnuminvertedStaticGroupImplementationSchemaUrn {
	return &v
}

type NullableEnuminvertedStaticGroupImplementationSchemaUrn struct {
	value *EnuminvertedStaticGroupImplementationSchemaUrn
	isSet bool
}

func (v NullableEnuminvertedStaticGroupImplementationSchemaUrn) Get() *EnuminvertedStaticGroupImplementationSchemaUrn {
	return v.value
}

func (v *NullableEnuminvertedStaticGroupImplementationSchemaUrn) Set(val *EnuminvertedStaticGroupImplementationSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnuminvertedStaticGroupImplementationSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnuminvertedStaticGroupImplementationSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnuminvertedStaticGroupImplementationSchemaUrn(val *EnuminvertedStaticGroupImplementationSchemaUrn) *NullableEnuminvertedStaticGroupImplementationSchemaUrn {
	return &NullableEnuminvertedStaticGroupImplementationSchemaUrn{value: val, isSet: true}
}

func (v NullableEnuminvertedStaticGroupImplementationSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnuminvertedStaticGroupImplementationSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
