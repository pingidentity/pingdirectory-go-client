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

// EnumbackendPrimeMethodProp Specifies the method that should be used to prime caches with data for this backend.
type EnumbackendPrimeMethodProp string

// List of Enumbackend-primeMethodProp
const (
	ENUMBACKENDPRIMEMETHODPROP_NONE                                     EnumbackendPrimeMethodProp = "none"
	ENUMBACKENDPRIMEMETHODPROP_PRELOAD                                  EnumbackendPrimeMethodProp = "preload"
	ENUMBACKENDPRIMEMETHODPROP_PRELOAD_INTERNAL_NODES_ONLY              EnumbackendPrimeMethodProp = "preload-internal-nodes-only"
	ENUMBACKENDPRIMEMETHODPROP_CURSOR_ACROSS_INDEXES                    EnumbackendPrimeMethodProp = "cursor-across-indexes"
	ENUMBACKENDPRIMEMETHODPROP_PRIME_TO_FILESYSTEM_CACHE                EnumbackendPrimeMethodProp = "prime-to-filesystem-cache"
	ENUMBACKENDPRIMEMETHODPROP_PRIME_TO_FILESYSTEM_CACHE_NON_SEQUENTIAL EnumbackendPrimeMethodProp = "prime-to-filesystem-cache-non-sequential"
)

// All allowed values of EnumbackendPrimeMethodProp enum
var AllowedEnumbackendPrimeMethodPropEnumValues = []EnumbackendPrimeMethodProp{
	"none",
	"preload",
	"preload-internal-nodes-only",
	"cursor-across-indexes",
	"prime-to-filesystem-cache",
	"prime-to-filesystem-cache-non-sequential",
}

func (v *EnumbackendPrimeMethodProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumbackendPrimeMethodProp(value)
	for _, existing := range AllowedEnumbackendPrimeMethodPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumbackendPrimeMethodProp", value)
}

// NewEnumbackendPrimeMethodPropFromValue returns a pointer to a valid EnumbackendPrimeMethodProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumbackendPrimeMethodPropFromValue(v string) (*EnumbackendPrimeMethodProp, error) {
	ev := EnumbackendPrimeMethodProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumbackendPrimeMethodProp: valid values are %v", v, AllowedEnumbackendPrimeMethodPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumbackendPrimeMethodProp) IsValid() bool {
	for _, existing := range AllowedEnumbackendPrimeMethodPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumbackend-primeMethodProp value
func (v EnumbackendPrimeMethodProp) Ptr() *EnumbackendPrimeMethodProp {
	return &v
}

type NullableEnumbackendPrimeMethodProp struct {
	value *EnumbackendPrimeMethodProp
	isSet bool
}

func (v NullableEnumbackendPrimeMethodProp) Get() *EnumbackendPrimeMethodProp {
	return v.value
}

func (v *NullableEnumbackendPrimeMethodProp) Set(val *EnumbackendPrimeMethodProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumbackendPrimeMethodProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumbackendPrimeMethodProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumbackendPrimeMethodProp(val *EnumbackendPrimeMethodProp) *NullableEnumbackendPrimeMethodProp {
	return &NullableEnumbackendPrimeMethodProp{value: val, isSet: true}
}

func (v NullableEnumbackendPrimeMethodProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumbackendPrimeMethodProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}