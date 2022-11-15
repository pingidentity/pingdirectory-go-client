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

// EnumfileBasedPassphraseProviderSchemaUrn the model 'EnumfileBasedPassphraseProviderSchemaUrn'
type EnumfileBasedPassphraseProviderSchemaUrn string

// List of Enumfile-based-passphrase-providerSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSPHRASE_PROVIDERFILE_BASED EnumfileBasedPassphraseProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:passphrase-provider:file-based"
)

// All allowed values of EnumfileBasedPassphraseProviderSchemaUrn enum
var AllowedEnumfileBasedPassphraseProviderSchemaUrnEnumValues = []EnumfileBasedPassphraseProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:passphrase-provider:file-based",
}

func (v *EnumfileBasedPassphraseProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumfileBasedPassphraseProviderSchemaUrn(value)
	for _, existing := range AllowedEnumfileBasedPassphraseProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumfileBasedPassphraseProviderSchemaUrn", value)
}

// NewEnumfileBasedPassphraseProviderSchemaUrnFromValue returns a pointer to a valid EnumfileBasedPassphraseProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumfileBasedPassphraseProviderSchemaUrnFromValue(v string) (*EnumfileBasedPassphraseProviderSchemaUrn, error) {
	ev := EnumfileBasedPassphraseProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumfileBasedPassphraseProviderSchemaUrn: valid values are %v", v, AllowedEnumfileBasedPassphraseProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumfileBasedPassphraseProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumfileBasedPassphraseProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumfile-based-passphrase-providerSchemaUrn value
func (v EnumfileBasedPassphraseProviderSchemaUrn) Ptr() *EnumfileBasedPassphraseProviderSchemaUrn {
	return &v
}

type NullableEnumfileBasedPassphraseProviderSchemaUrn struct {
	value *EnumfileBasedPassphraseProviderSchemaUrn
	isSet bool
}

func (v NullableEnumfileBasedPassphraseProviderSchemaUrn) Get() *EnumfileBasedPassphraseProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumfileBasedPassphraseProviderSchemaUrn) Set(val *EnumfileBasedPassphraseProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumfileBasedPassphraseProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumfileBasedPassphraseProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumfileBasedPassphraseProviderSchemaUrn(val *EnumfileBasedPassphraseProviderSchemaUrn) *NullableEnumfileBasedPassphraseProviderSchemaUrn {
	return &NullableEnumfileBasedPassphraseProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumfileBasedPassphraseProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumfileBasedPassphraseProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

