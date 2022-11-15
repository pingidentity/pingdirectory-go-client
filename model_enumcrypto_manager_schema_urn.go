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

// EnumcryptoManagerSchemaUrn the model 'EnumcryptoManagerSchemaUrn'
type EnumcryptoManagerSchemaUrn string

// List of Enumcrypto-managerSchemaUrn
const (
	ENUMCRYPTOMANAGERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CRYPTO_MANAGER EnumcryptoManagerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:crypto-manager"
)

// All allowed values of EnumcryptoManagerSchemaUrn enum
var AllowedEnumcryptoManagerSchemaUrnEnumValues = []EnumcryptoManagerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:crypto-manager",
}

func (v *EnumcryptoManagerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumcryptoManagerSchemaUrn(value)
	for _, existing := range AllowedEnumcryptoManagerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumcryptoManagerSchemaUrn", value)
}

// NewEnumcryptoManagerSchemaUrnFromValue returns a pointer to a valid EnumcryptoManagerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumcryptoManagerSchemaUrnFromValue(v string) (*EnumcryptoManagerSchemaUrn, error) {
	ev := EnumcryptoManagerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumcryptoManagerSchemaUrn: valid values are %v", v, AllowedEnumcryptoManagerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumcryptoManagerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumcryptoManagerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumcrypto-managerSchemaUrn value
func (v EnumcryptoManagerSchemaUrn) Ptr() *EnumcryptoManagerSchemaUrn {
	return &v
}

type NullableEnumcryptoManagerSchemaUrn struct {
	value *EnumcryptoManagerSchemaUrn
	isSet bool
}

func (v NullableEnumcryptoManagerSchemaUrn) Get() *EnumcryptoManagerSchemaUrn {
	return v.value
}

func (v *NullableEnumcryptoManagerSchemaUrn) Set(val *EnumcryptoManagerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumcryptoManagerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumcryptoManagerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumcryptoManagerSchemaUrn(val *EnumcryptoManagerSchemaUrn) *NullableEnumcryptoManagerSchemaUrn {
	return &NullableEnumcryptoManagerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumcryptoManagerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumcryptoManagerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

