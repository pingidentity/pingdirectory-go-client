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

// EnumamazonSecretsManagerCipherStreamProviderSchemaUrn the model 'EnumamazonSecretsManagerCipherStreamProviderSchemaUrn'
type EnumamazonSecretsManagerCipherStreamProviderSchemaUrn string

// List of Enumamazon-secrets-manager-cipher-stream-providerSchemaUrn
const (
	ENUMAMAZONSECRETSMANAGERCIPHERSTREAMPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CIPHER_STREAM_PROVIDERAMAZON_SECRETS_MANAGER EnumamazonSecretsManagerCipherStreamProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:cipher-stream-provider:amazon-secrets-manager"
)

// All allowed values of EnumamazonSecretsManagerCipherStreamProviderSchemaUrn enum
var AllowedEnumamazonSecretsManagerCipherStreamProviderSchemaUrnEnumValues = []EnumamazonSecretsManagerCipherStreamProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:cipher-stream-provider:amazon-secrets-manager",
}

func (v *EnumamazonSecretsManagerCipherStreamProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumamazonSecretsManagerCipherStreamProviderSchemaUrn(value)
	for _, existing := range AllowedEnumamazonSecretsManagerCipherStreamProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumamazonSecretsManagerCipherStreamProviderSchemaUrn", value)
}

// NewEnumamazonSecretsManagerCipherStreamProviderSchemaUrnFromValue returns a pointer to a valid EnumamazonSecretsManagerCipherStreamProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumamazonSecretsManagerCipherStreamProviderSchemaUrnFromValue(v string) (*EnumamazonSecretsManagerCipherStreamProviderSchemaUrn, error) {
	ev := EnumamazonSecretsManagerCipherStreamProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumamazonSecretsManagerCipherStreamProviderSchemaUrn: valid values are %v", v, AllowedEnumamazonSecretsManagerCipherStreamProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumamazonSecretsManagerCipherStreamProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumamazonSecretsManagerCipherStreamProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumamazon-secrets-manager-cipher-stream-providerSchemaUrn value
func (v EnumamazonSecretsManagerCipherStreamProviderSchemaUrn) Ptr() *EnumamazonSecretsManagerCipherStreamProviderSchemaUrn {
	return &v
}

type NullableEnumamazonSecretsManagerCipherStreamProviderSchemaUrn struct {
	value *EnumamazonSecretsManagerCipherStreamProviderSchemaUrn
	isSet bool
}

func (v NullableEnumamazonSecretsManagerCipherStreamProviderSchemaUrn) Get() *EnumamazonSecretsManagerCipherStreamProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumamazonSecretsManagerCipherStreamProviderSchemaUrn) Set(val *EnumamazonSecretsManagerCipherStreamProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumamazonSecretsManagerCipherStreamProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumamazonSecretsManagerCipherStreamProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumamazonSecretsManagerCipherStreamProviderSchemaUrn(val *EnumamazonSecretsManagerCipherStreamProviderSchemaUrn) *NullableEnumamazonSecretsManagerCipherStreamProviderSchemaUrn {
	return &NullableEnumamazonSecretsManagerCipherStreamProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumamazonSecretsManagerCipherStreamProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumamazonSecretsManagerCipherStreamProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
