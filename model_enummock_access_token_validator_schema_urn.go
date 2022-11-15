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

// EnummockAccessTokenValidatorSchemaUrn the model 'EnummockAccessTokenValidatorSchemaUrn'
type EnummockAccessTokenValidatorSchemaUrn string

// List of Enummock-access-token-validatorSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0ACCESS_TOKEN_VALIDATORMOCK EnummockAccessTokenValidatorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:access-token-validator:mock"
)

// All allowed values of EnummockAccessTokenValidatorSchemaUrn enum
var AllowedEnummockAccessTokenValidatorSchemaUrnEnumValues = []EnummockAccessTokenValidatorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:access-token-validator:mock",
}

func (v *EnummockAccessTokenValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnummockAccessTokenValidatorSchemaUrn(value)
	for _, existing := range AllowedEnummockAccessTokenValidatorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnummockAccessTokenValidatorSchemaUrn", value)
}

// NewEnummockAccessTokenValidatorSchemaUrnFromValue returns a pointer to a valid EnummockAccessTokenValidatorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnummockAccessTokenValidatorSchemaUrnFromValue(v string) (*EnummockAccessTokenValidatorSchemaUrn, error) {
	ev := EnummockAccessTokenValidatorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnummockAccessTokenValidatorSchemaUrn: valid values are %v", v, AllowedEnummockAccessTokenValidatorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnummockAccessTokenValidatorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnummockAccessTokenValidatorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enummock-access-token-validatorSchemaUrn value
func (v EnummockAccessTokenValidatorSchemaUrn) Ptr() *EnummockAccessTokenValidatorSchemaUrn {
	return &v
}

type NullableEnummockAccessTokenValidatorSchemaUrn struct {
	value *EnummockAccessTokenValidatorSchemaUrn
	isSet bool
}

func (v NullableEnummockAccessTokenValidatorSchemaUrn) Get() *EnummockAccessTokenValidatorSchemaUrn {
	return v.value
}

func (v *NullableEnummockAccessTokenValidatorSchemaUrn) Set(val *EnummockAccessTokenValidatorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnummockAccessTokenValidatorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnummockAccessTokenValidatorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnummockAccessTokenValidatorSchemaUrn(val *EnummockAccessTokenValidatorSchemaUrn) *NullableEnummockAccessTokenValidatorSchemaUrn {
	return &NullableEnummockAccessTokenValidatorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnummockAccessTokenValidatorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnummockAccessTokenValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

