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

// EnumnokiaDsExternalServerSchemaUrn the model 'EnumnokiaDsExternalServerSchemaUrn'
type EnumnokiaDsExternalServerSchemaUrn string

// List of Enumnokia-ds-external-serverSchemaUrn
const (
	ENUMNOKIADSEXTERNALSERVERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERNOKIA_DS EnumnokiaDsExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:nokia-ds"
)

// All allowed values of EnumnokiaDsExternalServerSchemaUrn enum
var AllowedEnumnokiaDsExternalServerSchemaUrnEnumValues = []EnumnokiaDsExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:nokia-ds",
}

func (v *EnumnokiaDsExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumnokiaDsExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumnokiaDsExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumnokiaDsExternalServerSchemaUrn", value)
}

// NewEnumnokiaDsExternalServerSchemaUrnFromValue returns a pointer to a valid EnumnokiaDsExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumnokiaDsExternalServerSchemaUrnFromValue(v string) (*EnumnokiaDsExternalServerSchemaUrn, error) {
	ev := EnumnokiaDsExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumnokiaDsExternalServerSchemaUrn: valid values are %v", v, AllowedEnumnokiaDsExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumnokiaDsExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumnokiaDsExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumnokia-ds-external-serverSchemaUrn value
func (v EnumnokiaDsExternalServerSchemaUrn) Ptr() *EnumnokiaDsExternalServerSchemaUrn {
	return &v
}

type NullableEnumnokiaDsExternalServerSchemaUrn struct {
	value *EnumnokiaDsExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumnokiaDsExternalServerSchemaUrn) Get() *EnumnokiaDsExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumnokiaDsExternalServerSchemaUrn) Set(val *EnumnokiaDsExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumnokiaDsExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumnokiaDsExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumnokiaDsExternalServerSchemaUrn(val *EnumnokiaDsExternalServerSchemaUrn) *NullableEnumnokiaDsExternalServerSchemaUrn {
	return &NullableEnumnokiaDsExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumnokiaDsExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumnokiaDsExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
