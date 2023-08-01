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

// EnumapiExternalServerSchemaUrn the model 'EnumapiExternalServerSchemaUrn'
type EnumapiExternalServerSchemaUrn string

// List of Enumapi-external-serverSchemaUrn
const (
	ENUMAPIEXTERNALSERVERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERAPI EnumapiExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:api"
)

// All allowed values of EnumapiExternalServerSchemaUrn enum
var AllowedEnumapiExternalServerSchemaUrnEnumValues = []EnumapiExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:api",
}

func (v *EnumapiExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumapiExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumapiExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumapiExternalServerSchemaUrn", value)
}

// NewEnumapiExternalServerSchemaUrnFromValue returns a pointer to a valid EnumapiExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumapiExternalServerSchemaUrnFromValue(v string) (*EnumapiExternalServerSchemaUrn, error) {
	ev := EnumapiExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumapiExternalServerSchemaUrn: valid values are %v", v, AllowedEnumapiExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumapiExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumapiExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumapi-external-serverSchemaUrn value
func (v EnumapiExternalServerSchemaUrn) Ptr() *EnumapiExternalServerSchemaUrn {
	return &v
}

type NullableEnumapiExternalServerSchemaUrn struct {
	value *EnumapiExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumapiExternalServerSchemaUrn) Get() *EnumapiExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumapiExternalServerSchemaUrn) Set(val *EnumapiExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumapiExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumapiExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumapiExternalServerSchemaUrn(val *EnumapiExternalServerSchemaUrn) *NullableEnumapiExternalServerSchemaUrn {
	return &NullableEnumapiExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumapiExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumapiExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
