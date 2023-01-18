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

// EnumidTokenValidatorAllowedSigningAlgorithmProp Specifies an allow list of JWT signing algorithms that will be accepted by the OpenID Connect ID Token Validator.
type EnumidTokenValidatorAllowedSigningAlgorithmProp string

// List of Enumid-token-validator-allowedSigningAlgorithmProp
const (
	ENUMIDTOKENVALIDATORALLOWEDSIGNINGALGORITHMPROP_RS256 EnumidTokenValidatorAllowedSigningAlgorithmProp = "RS256"
	ENUMIDTOKENVALIDATORALLOWEDSIGNINGALGORITHMPROP_RS384 EnumidTokenValidatorAllowedSigningAlgorithmProp = "RS384"
	ENUMIDTOKENVALIDATORALLOWEDSIGNINGALGORITHMPROP_RS512 EnumidTokenValidatorAllowedSigningAlgorithmProp = "RS512"
	ENUMIDTOKENVALIDATORALLOWEDSIGNINGALGORITHMPROP_ES256 EnumidTokenValidatorAllowedSigningAlgorithmProp = "ES256"
	ENUMIDTOKENVALIDATORALLOWEDSIGNINGALGORITHMPROP_ES384 EnumidTokenValidatorAllowedSigningAlgorithmProp = "ES384"
	ENUMIDTOKENVALIDATORALLOWEDSIGNINGALGORITHMPROP_ES512 EnumidTokenValidatorAllowedSigningAlgorithmProp = "ES512"
)

// All allowed values of EnumidTokenValidatorAllowedSigningAlgorithmProp enum
var AllowedEnumidTokenValidatorAllowedSigningAlgorithmPropEnumValues = []EnumidTokenValidatorAllowedSigningAlgorithmProp{
	"RS256",
	"RS384",
	"RS512",
	"ES256",
	"ES384",
	"ES512",
}

func (v *EnumidTokenValidatorAllowedSigningAlgorithmProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumidTokenValidatorAllowedSigningAlgorithmProp(value)
	for _, existing := range AllowedEnumidTokenValidatorAllowedSigningAlgorithmPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumidTokenValidatorAllowedSigningAlgorithmProp", value)
}

// NewEnumidTokenValidatorAllowedSigningAlgorithmPropFromValue returns a pointer to a valid EnumidTokenValidatorAllowedSigningAlgorithmProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumidTokenValidatorAllowedSigningAlgorithmPropFromValue(v string) (*EnumidTokenValidatorAllowedSigningAlgorithmProp, error) {
	ev := EnumidTokenValidatorAllowedSigningAlgorithmProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumidTokenValidatorAllowedSigningAlgorithmProp: valid values are %v", v, AllowedEnumidTokenValidatorAllowedSigningAlgorithmPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumidTokenValidatorAllowedSigningAlgorithmProp) IsValid() bool {
	for _, existing := range AllowedEnumidTokenValidatorAllowedSigningAlgorithmPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumid-token-validator-allowedSigningAlgorithmProp value
func (v EnumidTokenValidatorAllowedSigningAlgorithmProp) Ptr() *EnumidTokenValidatorAllowedSigningAlgorithmProp {
	return &v
}

type NullableEnumidTokenValidatorAllowedSigningAlgorithmProp struct {
	value *EnumidTokenValidatorAllowedSigningAlgorithmProp
	isSet bool
}

func (v NullableEnumidTokenValidatorAllowedSigningAlgorithmProp) Get() *EnumidTokenValidatorAllowedSigningAlgorithmProp {
	return v.value
}

func (v *NullableEnumidTokenValidatorAllowedSigningAlgorithmProp) Set(val *EnumidTokenValidatorAllowedSigningAlgorithmProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumidTokenValidatorAllowedSigningAlgorithmProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumidTokenValidatorAllowedSigningAlgorithmProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumidTokenValidatorAllowedSigningAlgorithmProp(val *EnumidTokenValidatorAllowedSigningAlgorithmProp) *NullableEnumidTokenValidatorAllowedSigningAlgorithmProp {
	return &NullableEnumidTokenValidatorAllowedSigningAlgorithmProp{value: val, isSet: true}
}

func (v NullableEnumidTokenValidatorAllowedSigningAlgorithmProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumidTokenValidatorAllowedSigningAlgorithmProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
