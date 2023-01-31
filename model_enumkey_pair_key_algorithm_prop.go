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

// EnumkeyPairKeyAlgorithmProp The algorithm name and the length in bits of the key, e.g. RSA_2048.
type EnumkeyPairKeyAlgorithmProp string

// List of Enumkey-pair-keyAlgorithmProp
const (
	ENUMKEYPAIRKEYALGORITHMPROP_RSA_2048 EnumkeyPairKeyAlgorithmProp = "RSA_2048"
	ENUMKEYPAIRKEYALGORITHMPROP_RSA_3072 EnumkeyPairKeyAlgorithmProp = "RSA_3072"
	ENUMKEYPAIRKEYALGORITHMPROP_RSA_4096 EnumkeyPairKeyAlgorithmProp = "RSA_4096"
	ENUMKEYPAIRKEYALGORITHMPROP_EC_256   EnumkeyPairKeyAlgorithmProp = "EC_256"
	ENUMKEYPAIRKEYALGORITHMPROP_EC_384   EnumkeyPairKeyAlgorithmProp = "EC_384"
	ENUMKEYPAIRKEYALGORITHMPROP_EC_521   EnumkeyPairKeyAlgorithmProp = "EC_521"
)

// All allowed values of EnumkeyPairKeyAlgorithmProp enum
var AllowedEnumkeyPairKeyAlgorithmPropEnumValues = []EnumkeyPairKeyAlgorithmProp{
	"RSA_2048",
	"RSA_3072",
	"RSA_4096",
	"EC_256",
	"EC_384",
	"EC_521",
}

func (v *EnumkeyPairKeyAlgorithmProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumkeyPairKeyAlgorithmProp(value)
	for _, existing := range AllowedEnumkeyPairKeyAlgorithmPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumkeyPairKeyAlgorithmProp", value)
}

// NewEnumkeyPairKeyAlgorithmPropFromValue returns a pointer to a valid EnumkeyPairKeyAlgorithmProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumkeyPairKeyAlgorithmPropFromValue(v string) (*EnumkeyPairKeyAlgorithmProp, error) {
	ev := EnumkeyPairKeyAlgorithmProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumkeyPairKeyAlgorithmProp: valid values are %v", v, AllowedEnumkeyPairKeyAlgorithmPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumkeyPairKeyAlgorithmProp) IsValid() bool {
	for _, existing := range AllowedEnumkeyPairKeyAlgorithmPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumkey-pair-keyAlgorithmProp value
func (v EnumkeyPairKeyAlgorithmProp) Ptr() *EnumkeyPairKeyAlgorithmProp {
	return &v
}

type NullableEnumkeyPairKeyAlgorithmProp struct {
	value *EnumkeyPairKeyAlgorithmProp
	isSet bool
}

func (v NullableEnumkeyPairKeyAlgorithmProp) Get() *EnumkeyPairKeyAlgorithmProp {
	return v.value
}

func (v *NullableEnumkeyPairKeyAlgorithmProp) Set(val *EnumkeyPairKeyAlgorithmProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumkeyPairKeyAlgorithmProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumkeyPairKeyAlgorithmProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumkeyPairKeyAlgorithmProp(val *EnumkeyPairKeyAlgorithmProp) *NullableEnumkeyPairKeyAlgorithmProp {
	return &NullableEnumkeyPairKeyAlgorithmProp{value: val, isSet: true}
}

func (v NullableEnumkeyPairKeyAlgorithmProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumkeyPairKeyAlgorithmProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
