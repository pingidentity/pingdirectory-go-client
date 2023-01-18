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

// EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp Specifies an allow list of JWT key encryption algorithms that will be accepted by the JWT Access Token Validator. This setting is only used if encryption-key-pair is set.
type EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp string

// List of Enumaccess-token-validator-allowedKeyEncryptionAlgorithmProp
const (
	ENUMACCESSTOKENVALIDATORALLOWEDKEYENCRYPTIONALGORITHMPROP_RSA_OAEP EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp = "RSA_OAEP"
	ENUMACCESSTOKENVALIDATORALLOWEDKEYENCRYPTIONALGORITHMPROP_ECDH_ES EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp = "ECDH_ES"
	ENUMACCESSTOKENVALIDATORALLOWEDKEYENCRYPTIONALGORITHMPROP_ECDH_ES_A128_KW EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp = "ECDH_ES_A128KW"
	ENUMACCESSTOKENVALIDATORALLOWEDKEYENCRYPTIONALGORITHMPROP_ECDH_ES_A192_KW EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp = "ECDH_ES_A192KW"
	ENUMACCESSTOKENVALIDATORALLOWEDKEYENCRYPTIONALGORITHMPROP_ECDH_ES_A256_KW EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp = "ECDH_ES_A256KW"
)

// All allowed values of EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp enum
var AllowedEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmPropEnumValues = []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp{
	"RSA_OAEP",
	"ECDH_ES",
	"ECDH_ES_A128KW",
	"ECDH_ES_A192KW",
	"ECDH_ES_A256KW",
}

func (v *EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp(value)
	for _, existing := range AllowedEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp", value)
}

// NewEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmPropFromValue returns a pointer to a valid EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmPropFromValue(v string) (*EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp, error) {
	ev := EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp: valid values are %v", v, AllowedEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) IsValid() bool {
	for _, existing := range AllowedEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumaccess-token-validator-allowedKeyEncryptionAlgorithmProp value
func (v EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) Ptr() *EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp {
	return &v
}

type NullableEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp struct {
	value *EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp
	isSet bool
}

func (v NullableEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) Get() *EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp {
	return v.value
}

func (v *NullableEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) Set(val *EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp(val *EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) *NullableEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp {
	return &NullableEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp{value: val, isSet: true}
}

func (v NullableEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
