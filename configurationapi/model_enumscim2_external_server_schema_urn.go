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

// Enumscim2ExternalServerSchemaUrn the model 'Enumscim2ExternalServerSchemaUrn'
type Enumscim2ExternalServerSchemaUrn string

// List of Enumscim2-external-serverSchemaUrn
const (
	ENUMSCIM2EXTERNALSERVERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERSCIM2 Enumscim2ExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:scim2"
)

// All allowed values of Enumscim2ExternalServerSchemaUrn enum
var AllowedEnumscim2ExternalServerSchemaUrnEnumValues = []Enumscim2ExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:scim2",
}

func (v *Enumscim2ExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Enumscim2ExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumscim2ExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Enumscim2ExternalServerSchemaUrn", value)
}

// NewEnumscim2ExternalServerSchemaUrnFromValue returns a pointer to a valid Enumscim2ExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumscim2ExternalServerSchemaUrnFromValue(v string) (*Enumscim2ExternalServerSchemaUrn, error) {
	ev := Enumscim2ExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Enumscim2ExternalServerSchemaUrn: valid values are %v", v, AllowedEnumscim2ExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Enumscim2ExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumscim2ExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumscim2-external-serverSchemaUrn value
func (v Enumscim2ExternalServerSchemaUrn) Ptr() *Enumscim2ExternalServerSchemaUrn {
	return &v
}

type NullableEnumscim2ExternalServerSchemaUrn struct {
	value *Enumscim2ExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumscim2ExternalServerSchemaUrn) Get() *Enumscim2ExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumscim2ExternalServerSchemaUrn) Set(val *Enumscim2ExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumscim2ExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumscim2ExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumscim2ExternalServerSchemaUrn(val *Enumscim2ExternalServerSchemaUrn) *NullableEnumscim2ExternalServerSchemaUrn {
	return &NullableEnumscim2ExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumscim2ExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumscim2ExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
