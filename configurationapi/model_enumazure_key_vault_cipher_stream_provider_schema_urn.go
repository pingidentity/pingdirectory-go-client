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

// EnumazureKeyVaultCipherStreamProviderSchemaUrn the model 'EnumazureKeyVaultCipherStreamProviderSchemaUrn'
type EnumazureKeyVaultCipherStreamProviderSchemaUrn string

// List of Enumazure-key-vault-cipher-stream-providerSchemaUrn
const (
	ENUMAZUREKEYVAULTCIPHERSTREAMPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CIPHER_STREAM_PROVIDERAZURE_KEY_VAULT EnumazureKeyVaultCipherStreamProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:cipher-stream-provider:azure-key-vault"
)

// All allowed values of EnumazureKeyVaultCipherStreamProviderSchemaUrn enum
var AllowedEnumazureKeyVaultCipherStreamProviderSchemaUrnEnumValues = []EnumazureKeyVaultCipherStreamProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:cipher-stream-provider:azure-key-vault",
}

func (v *EnumazureKeyVaultCipherStreamProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumazureKeyVaultCipherStreamProviderSchemaUrn(value)
	for _, existing := range AllowedEnumazureKeyVaultCipherStreamProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumazureKeyVaultCipherStreamProviderSchemaUrn", value)
}

// NewEnumazureKeyVaultCipherStreamProviderSchemaUrnFromValue returns a pointer to a valid EnumazureKeyVaultCipherStreamProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumazureKeyVaultCipherStreamProviderSchemaUrnFromValue(v string) (*EnumazureKeyVaultCipherStreamProviderSchemaUrn, error) {
	ev := EnumazureKeyVaultCipherStreamProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumazureKeyVaultCipherStreamProviderSchemaUrn: valid values are %v", v, AllowedEnumazureKeyVaultCipherStreamProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumazureKeyVaultCipherStreamProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumazureKeyVaultCipherStreamProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumazure-key-vault-cipher-stream-providerSchemaUrn value
func (v EnumazureKeyVaultCipherStreamProviderSchemaUrn) Ptr() *EnumazureKeyVaultCipherStreamProviderSchemaUrn {
	return &v
}

type NullableEnumazureKeyVaultCipherStreamProviderSchemaUrn struct {
	value *EnumazureKeyVaultCipherStreamProviderSchemaUrn
	isSet bool
}

func (v NullableEnumazureKeyVaultCipherStreamProviderSchemaUrn) Get() *EnumazureKeyVaultCipherStreamProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumazureKeyVaultCipherStreamProviderSchemaUrn) Set(val *EnumazureKeyVaultCipherStreamProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumazureKeyVaultCipherStreamProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumazureKeyVaultCipherStreamProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumazureKeyVaultCipherStreamProviderSchemaUrn(val *EnumazureKeyVaultCipherStreamProviderSchemaUrn) *NullableEnumazureKeyVaultCipherStreamProviderSchemaUrn {
	return &NullableEnumazureKeyVaultCipherStreamProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumazureKeyVaultCipherStreamProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumazureKeyVaultCipherStreamProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}