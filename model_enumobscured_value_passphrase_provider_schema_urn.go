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

// EnumobscuredValuePassphraseProviderSchemaUrn the model 'EnumobscuredValuePassphraseProviderSchemaUrn'
type EnumobscuredValuePassphraseProviderSchemaUrn string

// List of Enumobscured-value-passphrase-providerSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSPHRASE_PROVIDEROBSCURED_VALUE EnumobscuredValuePassphraseProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:passphrase-provider:obscured-value"
)

// All allowed values of EnumobscuredValuePassphraseProviderSchemaUrn enum
var AllowedEnumobscuredValuePassphraseProviderSchemaUrnEnumValues = []EnumobscuredValuePassphraseProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:passphrase-provider:obscured-value",
}

func (v *EnumobscuredValuePassphraseProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumobscuredValuePassphraseProviderSchemaUrn(value)
	for _, existing := range AllowedEnumobscuredValuePassphraseProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumobscuredValuePassphraseProviderSchemaUrn", value)
}

// NewEnumobscuredValuePassphraseProviderSchemaUrnFromValue returns a pointer to a valid EnumobscuredValuePassphraseProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumobscuredValuePassphraseProviderSchemaUrnFromValue(v string) (*EnumobscuredValuePassphraseProviderSchemaUrn, error) {
	ev := EnumobscuredValuePassphraseProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumobscuredValuePassphraseProviderSchemaUrn: valid values are %v", v, AllowedEnumobscuredValuePassphraseProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumobscuredValuePassphraseProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumobscuredValuePassphraseProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumobscured-value-passphrase-providerSchemaUrn value
func (v EnumobscuredValuePassphraseProviderSchemaUrn) Ptr() *EnumobscuredValuePassphraseProviderSchemaUrn {
	return &v
}

type NullableEnumobscuredValuePassphraseProviderSchemaUrn struct {
	value *EnumobscuredValuePassphraseProviderSchemaUrn
	isSet bool
}

func (v NullableEnumobscuredValuePassphraseProviderSchemaUrn) Get() *EnumobscuredValuePassphraseProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumobscuredValuePassphraseProviderSchemaUrn) Set(val *EnumobscuredValuePassphraseProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumobscuredValuePassphraseProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumobscuredValuePassphraseProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumobscuredValuePassphraseProviderSchemaUrn(val *EnumobscuredValuePassphraseProviderSchemaUrn) *NullableEnumobscuredValuePassphraseProviderSchemaUrn {
	return &NullableEnumobscuredValuePassphraseProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumobscuredValuePassphraseProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumobscuredValuePassphraseProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

