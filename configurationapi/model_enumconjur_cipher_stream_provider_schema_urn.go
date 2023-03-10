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

// EnumconjurCipherStreamProviderSchemaUrn the model 'EnumconjurCipherStreamProviderSchemaUrn'
type EnumconjurCipherStreamProviderSchemaUrn string

// List of Enumconjur-cipher-stream-providerSchemaUrn
const (
	ENUMCONJURCIPHERSTREAMPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0CIPHER_STREAM_PROVIDERCONJUR EnumconjurCipherStreamProviderSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:cipher-stream-provider:conjur"
)

// All allowed values of EnumconjurCipherStreamProviderSchemaUrn enum
var AllowedEnumconjurCipherStreamProviderSchemaUrnEnumValues = []EnumconjurCipherStreamProviderSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:cipher-stream-provider:conjur",
}

func (v *EnumconjurCipherStreamProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumconjurCipherStreamProviderSchemaUrn(value)
	for _, existing := range AllowedEnumconjurCipherStreamProviderSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumconjurCipherStreamProviderSchemaUrn", value)
}

// NewEnumconjurCipherStreamProviderSchemaUrnFromValue returns a pointer to a valid EnumconjurCipherStreamProviderSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumconjurCipherStreamProviderSchemaUrnFromValue(v string) (*EnumconjurCipherStreamProviderSchemaUrn, error) {
	ev := EnumconjurCipherStreamProviderSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumconjurCipherStreamProviderSchemaUrn: valid values are %v", v, AllowedEnumconjurCipherStreamProviderSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumconjurCipherStreamProviderSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumconjurCipherStreamProviderSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumconjur-cipher-stream-providerSchemaUrn value
func (v EnumconjurCipherStreamProviderSchemaUrn) Ptr() *EnumconjurCipherStreamProviderSchemaUrn {
	return &v
}

type NullableEnumconjurCipherStreamProviderSchemaUrn struct {
	value *EnumconjurCipherStreamProviderSchemaUrn
	isSet bool
}

func (v NullableEnumconjurCipherStreamProviderSchemaUrn) Get() *EnumconjurCipherStreamProviderSchemaUrn {
	return v.value
}

func (v *NullableEnumconjurCipherStreamProviderSchemaUrn) Set(val *EnumconjurCipherStreamProviderSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumconjurCipherStreamProviderSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumconjurCipherStreamProviderSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumconjurCipherStreamProviderSchemaUrn(val *EnumconjurCipherStreamProviderSchemaUrn) *NullableEnumconjurCipherStreamProviderSchemaUrn {
	return &NullableEnumconjurCipherStreamProviderSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumconjurCipherStreamProviderSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumconjurCipherStreamProviderSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
