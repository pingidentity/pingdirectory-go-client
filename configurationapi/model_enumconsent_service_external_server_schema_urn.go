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

// EnumconsentServiceExternalServerSchemaUrn the model 'EnumconsentServiceExternalServerSchemaUrn'
type EnumconsentServiceExternalServerSchemaUrn string

// List of Enumconsent-service-external-serverSchemaUrn
const (
	ENUMCONSENTSERVICEEXTERNALSERVERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERCONSENT_SERVICE EnumconsentServiceExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:consent-service"
)

// All allowed values of EnumconsentServiceExternalServerSchemaUrn enum
var AllowedEnumconsentServiceExternalServerSchemaUrnEnumValues = []EnumconsentServiceExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:consent-service",
}

func (v *EnumconsentServiceExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumconsentServiceExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumconsentServiceExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumconsentServiceExternalServerSchemaUrn", value)
}

// NewEnumconsentServiceExternalServerSchemaUrnFromValue returns a pointer to a valid EnumconsentServiceExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumconsentServiceExternalServerSchemaUrnFromValue(v string) (*EnumconsentServiceExternalServerSchemaUrn, error) {
	ev := EnumconsentServiceExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumconsentServiceExternalServerSchemaUrn: valid values are %v", v, AllowedEnumconsentServiceExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumconsentServiceExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumconsentServiceExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumconsent-service-external-serverSchemaUrn value
func (v EnumconsentServiceExternalServerSchemaUrn) Ptr() *EnumconsentServiceExternalServerSchemaUrn {
	return &v
}

type NullableEnumconsentServiceExternalServerSchemaUrn struct {
	value *EnumconsentServiceExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumconsentServiceExternalServerSchemaUrn) Get() *EnumconsentServiceExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumconsentServiceExternalServerSchemaUrn) Set(val *EnumconsentServiceExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumconsentServiceExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumconsentServiceExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumconsentServiceExternalServerSchemaUrn(val *EnumconsentServiceExternalServerSchemaUrn) *NullableEnumconsentServiceExternalServerSchemaUrn {
	return &NullableEnumconsentServiceExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumconsentServiceExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumconsentServiceExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}