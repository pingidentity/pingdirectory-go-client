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

// EnumbackendDn2uriCacheModeProp Specifies the cache mode that should be used when accessing the records in the dn2uri database, which provides a mapping between a normalized entry DN and a set of referral URLs contained in the associated smart referral entry.
type EnumbackendDn2uriCacheModeProp string

// List of Enumbackend-dn2uriCacheModeProp
const (
	ENUMBACKENDDN2URICACHEMODEPROP_CACHE_KEYS_AND_VALUES  EnumbackendDn2uriCacheModeProp = "cache-keys-and-values"
	ENUMBACKENDDN2URICACHEMODEPROP_CACHE_KEYS_ONLY        EnumbackendDn2uriCacheModeProp = "cache-keys-only"
	ENUMBACKENDDN2URICACHEMODEPROP_NO_CACHING             EnumbackendDn2uriCacheModeProp = "no-caching"
	ENUMBACKENDDN2URICACHEMODEPROP_KEEP_HOT               EnumbackendDn2uriCacheModeProp = "keep-hot"
	ENUMBACKENDDN2URICACHEMODEPROP_DEFAULT                EnumbackendDn2uriCacheModeProp = "default"
	ENUMBACKENDDN2URICACHEMODEPROP_MAKE_COLD              EnumbackendDn2uriCacheModeProp = "make-cold"
	ENUMBACKENDDN2URICACHEMODEPROP_EVICT_LEAF_IMMEDIATELY EnumbackendDn2uriCacheModeProp = "evict-leaf-immediately"
	ENUMBACKENDDN2URICACHEMODEPROP_EVICT_BIN_IMMEDIATELY  EnumbackendDn2uriCacheModeProp = "evict-bin-immediately"
)

// All allowed values of EnumbackendDn2uriCacheModeProp enum
var AllowedEnumbackendDn2uriCacheModePropEnumValues = []EnumbackendDn2uriCacheModeProp{
	"cache-keys-and-values",
	"cache-keys-only",
	"no-caching",
	"keep-hot",
	"default",
	"make-cold",
	"evict-leaf-immediately",
	"evict-bin-immediately",
}

func (v *EnumbackendDn2uriCacheModeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumbackendDn2uriCacheModeProp(value)
	for _, existing := range AllowedEnumbackendDn2uriCacheModePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumbackendDn2uriCacheModeProp", value)
}

// NewEnumbackendDn2uriCacheModePropFromValue returns a pointer to a valid EnumbackendDn2uriCacheModeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumbackendDn2uriCacheModePropFromValue(v string) (*EnumbackendDn2uriCacheModeProp, error) {
	ev := EnumbackendDn2uriCacheModeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumbackendDn2uriCacheModeProp: valid values are %v", v, AllowedEnumbackendDn2uriCacheModePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumbackendDn2uriCacheModeProp) IsValid() bool {
	for _, existing := range AllowedEnumbackendDn2uriCacheModePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumbackend-dn2uriCacheModeProp value
func (v EnumbackendDn2uriCacheModeProp) Ptr() *EnumbackendDn2uriCacheModeProp {
	return &v
}

type NullableEnumbackendDn2uriCacheModeProp struct {
	value *EnumbackendDn2uriCacheModeProp
	isSet bool
}

func (v NullableEnumbackendDn2uriCacheModeProp) Get() *EnumbackendDn2uriCacheModeProp {
	return v.value
}

func (v *NullableEnumbackendDn2uriCacheModeProp) Set(val *EnumbackendDn2uriCacheModeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumbackendDn2uriCacheModeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumbackendDn2uriCacheModeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumbackendDn2uriCacheModeProp(val *EnumbackendDn2uriCacheModeProp) *NullableEnumbackendDn2uriCacheModeProp {
	return &NullableEnumbackendDn2uriCacheModeProp{value: val, isSet: true}
}

func (v NullableEnumbackendDn2uriCacheModeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumbackendDn2uriCacheModeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}