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

// EnumvaultPassphraseProviderSchemaUrn the model 'EnumvaultPassphraseProviderSchemaUrn'
type EnumvaultPassphraseProviderSchemaUrn string

// List of Enumvault-passphrase-providerSchemaUrn
const (
	ENUMVAULTPASSPHRASEPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSPHRASE_PROVIDERVAULT EnumvaultPassphraseProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:passphrase-provider:vault"
)

// All allowed values of EnumvaultPassphraseProviderSchemaUrn enum
var AllowedEnumvaultPassphraseProviderSchemaUrnEnumValues = []EnumvaultPassphraseProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:passphrase-provider:vault",
}

func (v *EnumvaultPassphraseProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumvaultPassphraseProviderSchemaUrn(value)
	for _, existing := range AllowedEnumvaultPassphraseProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumvaultPassphraseProviderSchemaUrn", value)
}

// NewEnumvaultPassphraseProviderSchemaUrnFromValue returns a pointer to a valid EnumvaultPassphraseProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumvaultPassphraseProviderSchemaUrnFromValue(v string) (*EnumvaultPassphraseProviderSchemaUrn, error) {
	ev := EnumvaultPassphraseProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumvaultPassphraseProviderSchemaUrn: valid values are %v", v, AllowedEnumvaultPassphraseProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumvaultPassphraseProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumvaultPassphraseProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumvault-passphrase-providerSchemaUrn value
func (v EnumvaultPassphraseProviderSchemaUrn) Ptr() *EnumvaultPassphraseProviderSchemaUrn {
	return &v
}

type NullableEnumvaultPassphraseProviderSchemaUrn struct {
	value *EnumvaultPassphraseProviderSchemaUrn
	isSet bool
}

func (v NullableEnumvaultPassphraseProviderSchemaUrn) Get() *EnumvaultPassphraseProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumvaultPassphraseProviderSchemaUrn) Set(val *EnumvaultPassphraseProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumvaultPassphraseProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumvaultPassphraseProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumvaultPassphraseProviderSchemaUrn(val *EnumvaultPassphraseProviderSchemaUrn) *NullableEnumvaultPassphraseProviderSchemaUrn {
	return &NullableEnumvaultPassphraseProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumvaultPassphraseProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumvaultPassphraseProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
