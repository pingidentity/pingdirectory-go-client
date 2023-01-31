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

// EnumconsentDefinitionSchemaUrn the model 'EnumconsentDefinitionSchemaUrn'
type EnumconsentDefinitionSchemaUrn string

// List of Enumconsent-definitionSchemaUrn
const (
	ENUMCONSENTDEFINITIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CONSENT_DEFINITION EnumconsentDefinitionSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:consent-definition"
)

// All allowed values of EnumconsentDefinitionSchemaUrn enum
var AllowedEnumconsentDefinitionSchemaUrnEnumValues = []EnumconsentDefinitionSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:consent-definition",
}

func (v *EnumconsentDefinitionSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumconsentDefinitionSchemaUrn(value)
	for _, existing := range AllowedEnumconsentDefinitionSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumconsentDefinitionSchemaUrn", value)
}

// NewEnumconsentDefinitionSchemaUrnFromValue returns a pointer to a valid EnumconsentDefinitionSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumconsentDefinitionSchemaUrnFromValue(v string) (*EnumconsentDefinitionSchemaUrn, error) {
	ev := EnumconsentDefinitionSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumconsentDefinitionSchemaUrn: valid values are %v", v, AllowedEnumconsentDefinitionSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumconsentDefinitionSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumconsentDefinitionSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumconsent-definitionSchemaUrn value
func (v EnumconsentDefinitionSchemaUrn) Ptr() *EnumconsentDefinitionSchemaUrn {
	return &v
}

type NullableEnumconsentDefinitionSchemaUrn struct {
	value *EnumconsentDefinitionSchemaUrn
	isSet bool
}

func (v NullableEnumconsentDefinitionSchemaUrn) Get() *EnumconsentDefinitionSchemaUrn {
	return v.value
}

func (v *NullableEnumconsentDefinitionSchemaUrn) Set(val *EnumconsentDefinitionSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumconsentDefinitionSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumconsentDefinitionSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumconsentDefinitionSchemaUrn(val *EnumconsentDefinitionSchemaUrn) *NullableEnumconsentDefinitionSchemaUrn {
	return &NullableEnumconsentDefinitionSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumconsentDefinitionSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumconsentDefinitionSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
