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

// EnumhttpExternalServerSchemaUrn the model 'EnumhttpExternalServerSchemaUrn'
type EnumhttpExternalServerSchemaUrn string

// List of Enumhttp-external-serverSchemaUrn
const (
	ENUMHTTPEXTERNALSERVERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERHTTP EnumhttpExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:http"
)

// All allowed values of EnumhttpExternalServerSchemaUrn enum
var AllowedEnumhttpExternalServerSchemaUrnEnumValues = []EnumhttpExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:http",
}

func (v *EnumhttpExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumhttpExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumhttpExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumhttpExternalServerSchemaUrn", value)
}

// NewEnumhttpExternalServerSchemaUrnFromValue returns a pointer to a valid EnumhttpExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumhttpExternalServerSchemaUrnFromValue(v string) (*EnumhttpExternalServerSchemaUrn, error) {
	ev := EnumhttpExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumhttpExternalServerSchemaUrn: valid values are %v", v, AllowedEnumhttpExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumhttpExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumhttpExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumhttp-external-serverSchemaUrn value
func (v EnumhttpExternalServerSchemaUrn) Ptr() *EnumhttpExternalServerSchemaUrn {
	return &v
}

type NullableEnumhttpExternalServerSchemaUrn struct {
	value *EnumhttpExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumhttpExternalServerSchemaUrn) Get() *EnumhttpExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumhttpExternalServerSchemaUrn) Set(val *EnumhttpExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumhttpExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumhttpExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumhttpExternalServerSchemaUrn(val *EnumhttpExternalServerSchemaUrn) *NullableEnumhttpExternalServerSchemaUrn {
	return &NullableEnumhttpExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumhttpExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumhttpExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
