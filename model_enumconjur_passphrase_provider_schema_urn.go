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

// EnumconjurPassphraseProviderSchemaUrn the model 'EnumconjurPassphraseProviderSchemaUrn'
type EnumconjurPassphraseProviderSchemaUrn string

// List of Enumconjur-passphrase-providerSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSPHRASE_PROVIDERCONJUR EnumconjurPassphraseProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:passphrase-provider:conjur"
)

// All allowed values of EnumconjurPassphraseProviderSchemaUrn enum
var AllowedEnumconjurPassphraseProviderSchemaUrnEnumValues = []EnumconjurPassphraseProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:passphrase-provider:conjur",
}

func (v *EnumconjurPassphraseProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumconjurPassphraseProviderSchemaUrn(value)
	for _, existing := range AllowedEnumconjurPassphraseProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumconjurPassphraseProviderSchemaUrn", value)
}

// NewEnumconjurPassphraseProviderSchemaUrnFromValue returns a pointer to a valid EnumconjurPassphraseProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumconjurPassphraseProviderSchemaUrnFromValue(v string) (*EnumconjurPassphraseProviderSchemaUrn, error) {
	ev := EnumconjurPassphraseProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumconjurPassphraseProviderSchemaUrn: valid values are %v", v, AllowedEnumconjurPassphraseProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumconjurPassphraseProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumconjurPassphraseProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumconjur-passphrase-providerSchemaUrn value
func (v EnumconjurPassphraseProviderSchemaUrn) Ptr() *EnumconjurPassphraseProviderSchemaUrn {
	return &v
}

type NullableEnumconjurPassphraseProviderSchemaUrn struct {
	value *EnumconjurPassphraseProviderSchemaUrn
	isSet bool
}

func (v NullableEnumconjurPassphraseProviderSchemaUrn) Get() *EnumconjurPassphraseProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumconjurPassphraseProviderSchemaUrn) Set(val *EnumconjurPassphraseProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumconjurPassphraseProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumconjurPassphraseProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumconjurPassphraseProviderSchemaUrn(val *EnumconjurPassphraseProviderSchemaUrn) *NullableEnumconjurPassphraseProviderSchemaUrn {
	return &NullableEnumconjurPassphraseProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumconjurPassphraseProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumconjurPassphraseProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

