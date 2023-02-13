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

// EnumconsentDefinitionLocalizationSchemaUrn the model 'EnumconsentDefinitionLocalizationSchemaUrn'
type EnumconsentDefinitionLocalizationSchemaUrn string

// List of Enumconsent-definition-localizationSchemaUrn
const (
	ENUMCONSENTDEFINITIONLOCALIZATIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CONSENT_DEFINITION_LOCALIZATION EnumconsentDefinitionLocalizationSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:consent-definition-localization"
)

// All allowed values of EnumconsentDefinitionLocalizationSchemaUrn enum
var AllowedEnumconsentDefinitionLocalizationSchemaUrnEnumValues = []EnumconsentDefinitionLocalizationSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:consent-definition-localization",
}

func (v *EnumconsentDefinitionLocalizationSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumconsentDefinitionLocalizationSchemaUrn(value)
	for _, existing := range AllowedEnumconsentDefinitionLocalizationSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumconsentDefinitionLocalizationSchemaUrn", value)
}

// NewEnumconsentDefinitionLocalizationSchemaUrnFromValue returns a pointer to a valid EnumconsentDefinitionLocalizationSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumconsentDefinitionLocalizationSchemaUrnFromValue(v string) (*EnumconsentDefinitionLocalizationSchemaUrn, error) {
	ev := EnumconsentDefinitionLocalizationSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumconsentDefinitionLocalizationSchemaUrn: valid values are %v", v, AllowedEnumconsentDefinitionLocalizationSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumconsentDefinitionLocalizationSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumconsentDefinitionLocalizationSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumconsent-definition-localizationSchemaUrn value
func (v EnumconsentDefinitionLocalizationSchemaUrn) Ptr() *EnumconsentDefinitionLocalizationSchemaUrn {
	return &v
}

type NullableEnumconsentDefinitionLocalizationSchemaUrn struct {
	value *EnumconsentDefinitionLocalizationSchemaUrn
	isSet bool
}

func (v NullableEnumconsentDefinitionLocalizationSchemaUrn) Get() *EnumconsentDefinitionLocalizationSchemaUrn {
	return v.value
}

func (v *NullableEnumconsentDefinitionLocalizationSchemaUrn) Set(val *EnumconsentDefinitionLocalizationSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumconsentDefinitionLocalizationSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumconsentDefinitionLocalizationSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumconsentDefinitionLocalizationSchemaUrn(val *EnumconsentDefinitionLocalizationSchemaUrn) *NullableEnumconsentDefinitionLocalizationSchemaUrn {
	return &NullableEnumconsentDefinitionLocalizationSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumconsentDefinitionLocalizationSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumconsentDefinitionLocalizationSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
