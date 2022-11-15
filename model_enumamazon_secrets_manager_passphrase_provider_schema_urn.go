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

// EnumamazonSecretsManagerPassphraseProviderSchemaUrn the model 'EnumamazonSecretsManagerPassphraseProviderSchemaUrn'
type EnumamazonSecretsManagerPassphraseProviderSchemaUrn string

// List of Enumamazon-secrets-manager-passphrase-providerSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASSPHRASE_PROVIDERAMAZON_SECRETS_MANAGER EnumamazonSecretsManagerPassphraseProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:passphrase-provider:amazon-secrets-manager"
)

// All allowed values of EnumamazonSecretsManagerPassphraseProviderSchemaUrn enum
var AllowedEnumamazonSecretsManagerPassphraseProviderSchemaUrnEnumValues = []EnumamazonSecretsManagerPassphraseProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:passphrase-provider:amazon-secrets-manager",
}

func (v *EnumamazonSecretsManagerPassphraseProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumamazonSecretsManagerPassphraseProviderSchemaUrn(value)
	for _, existing := range AllowedEnumamazonSecretsManagerPassphraseProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumamazonSecretsManagerPassphraseProviderSchemaUrn", value)
}

// NewEnumamazonSecretsManagerPassphraseProviderSchemaUrnFromValue returns a pointer to a valid EnumamazonSecretsManagerPassphraseProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumamazonSecretsManagerPassphraseProviderSchemaUrnFromValue(v string) (*EnumamazonSecretsManagerPassphraseProviderSchemaUrn, error) {
	ev := EnumamazonSecretsManagerPassphraseProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumamazonSecretsManagerPassphraseProviderSchemaUrn: valid values are %v", v, AllowedEnumamazonSecretsManagerPassphraseProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumamazonSecretsManagerPassphraseProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumamazonSecretsManagerPassphraseProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumamazon-secrets-manager-passphrase-providerSchemaUrn value
func (v EnumamazonSecretsManagerPassphraseProviderSchemaUrn) Ptr() *EnumamazonSecretsManagerPassphraseProviderSchemaUrn {
	return &v
}

type NullableEnumamazonSecretsManagerPassphraseProviderSchemaUrn struct {
	value *EnumamazonSecretsManagerPassphraseProviderSchemaUrn
	isSet bool
}

func (v NullableEnumamazonSecretsManagerPassphraseProviderSchemaUrn) Get() *EnumamazonSecretsManagerPassphraseProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumamazonSecretsManagerPassphraseProviderSchemaUrn) Set(val *EnumamazonSecretsManagerPassphraseProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumamazonSecretsManagerPassphraseProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumamazonSecretsManagerPassphraseProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumamazonSecretsManagerPassphraseProviderSchemaUrn(val *EnumamazonSecretsManagerPassphraseProviderSchemaUrn) *NullableEnumamazonSecretsManagerPassphraseProviderSchemaUrn {
	return &NullableEnumamazonSecretsManagerPassphraseProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumamazonSecretsManagerPassphraseProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumamazonSecretsManagerPassphraseProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

