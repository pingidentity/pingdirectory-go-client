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

// EnumpingIdentityDsExternalServerSchemaUrn the model 'EnumpingIdentityDsExternalServerSchemaUrn'
type EnumpingIdentityDsExternalServerSchemaUrn string

// List of Enumping-identity-ds-external-serverSchemaUrn
const (
	ENUMPINGIDENTITYDSEXTERNALSERVERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0EXTERNAL_SERVERPING_IDENTITY_DS EnumpingIdentityDsExternalServerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:external-server:ping-identity-ds"
)

// All allowed values of EnumpingIdentityDsExternalServerSchemaUrn enum
var AllowedEnumpingIdentityDsExternalServerSchemaUrnEnumValues = []EnumpingIdentityDsExternalServerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:external-server:ping-identity-ds",
}

func (v *EnumpingIdentityDsExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpingIdentityDsExternalServerSchemaUrn(value)
	for _, existing := range AllowedEnumpingIdentityDsExternalServerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpingIdentityDsExternalServerSchemaUrn", value)
}

// NewEnumpingIdentityDsExternalServerSchemaUrnFromValue returns a pointer to a valid EnumpingIdentityDsExternalServerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpingIdentityDsExternalServerSchemaUrnFromValue(v string) (*EnumpingIdentityDsExternalServerSchemaUrn, error) {
	ev := EnumpingIdentityDsExternalServerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpingIdentityDsExternalServerSchemaUrn: valid values are %v", v, AllowedEnumpingIdentityDsExternalServerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpingIdentityDsExternalServerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumpingIdentityDsExternalServerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumping-identity-ds-external-serverSchemaUrn value
func (v EnumpingIdentityDsExternalServerSchemaUrn) Ptr() *EnumpingIdentityDsExternalServerSchemaUrn {
	return &v
}

type NullableEnumpingIdentityDsExternalServerSchemaUrn struct {
	value *EnumpingIdentityDsExternalServerSchemaUrn
	isSet bool
}

func (v NullableEnumpingIdentityDsExternalServerSchemaUrn) Get() *EnumpingIdentityDsExternalServerSchemaUrn {
	return v.value
}

func (v *NullableEnumpingIdentityDsExternalServerSchemaUrn) Set(val *EnumpingIdentityDsExternalServerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpingIdentityDsExternalServerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpingIdentityDsExternalServerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpingIdentityDsExternalServerSchemaUrn(val *EnumpingIdentityDsExternalServerSchemaUrn) *NullableEnumpingIdentityDsExternalServerSchemaUrn {
	return &NullableEnumpingIdentityDsExternalServerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumpingIdentityDsExternalServerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpingIdentityDsExternalServerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
