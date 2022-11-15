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

// EnumbackendDefaultCacheModeProp Specifies the cache mode that should be used for any database for which the cache mode is not explicitly specified. This includes the id2entry database, which stores encoded entries, and all attribute indexes.
type EnumbackendDefaultCacheModeProp string

// List of Enumbackend-defaultCacheModeProp
const (
	ENUMBACKENDDEFAULTCACHEMODEPROP_CACHE_KEYS_AND_VALUES EnumbackendDefaultCacheModeProp = "cache-keys-and-values"
	ENUMBACKENDDEFAULTCACHEMODEPROP_CACHE_KEYS_ONLY EnumbackendDefaultCacheModeProp = "cache-keys-only"
	ENUMBACKENDDEFAULTCACHEMODEPROP_NO_CACHING EnumbackendDefaultCacheModeProp = "no-caching"
	ENUMBACKENDDEFAULTCACHEMODEPROP_KEEP_HOT EnumbackendDefaultCacheModeProp = "keep-hot"
	ENUMBACKENDDEFAULTCACHEMODEPROP_DEFAULT EnumbackendDefaultCacheModeProp = "default"
	ENUMBACKENDDEFAULTCACHEMODEPROP_MAKE_COLD EnumbackendDefaultCacheModeProp = "make-cold"
	ENUMBACKENDDEFAULTCACHEMODEPROP_EVICT_LEAF_IMMEDIATELY EnumbackendDefaultCacheModeProp = "evict-leaf-immediately"
	ENUMBACKENDDEFAULTCACHEMODEPROP_EVICT_BIN_IMMEDIATELY EnumbackendDefaultCacheModeProp = "evict-bin-immediately"
)

// All allowed values of EnumbackendDefaultCacheModeProp enum
var AllowedEnumbackendDefaultCacheModePropEnumValues = []EnumbackendDefaultCacheModeProp{
	"cache-keys-and-values",
	"cache-keys-only",
	"no-caching",
	"keep-hot",
	"default",
	"make-cold",
	"evict-leaf-immediately",
	"evict-bin-immediately",
}

func (v *EnumbackendDefaultCacheModeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumbackendDefaultCacheModeProp(value)
	for _, existing := range AllowedEnumbackendDefaultCacheModePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumbackendDefaultCacheModeProp", value)
}

// NewEnumbackendDefaultCacheModePropFromValue returns a pointer to a valid EnumbackendDefaultCacheModeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumbackendDefaultCacheModePropFromValue(v string) (*EnumbackendDefaultCacheModeProp, error) {
	ev := EnumbackendDefaultCacheModeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumbackendDefaultCacheModeProp: valid values are %v", v, AllowedEnumbackendDefaultCacheModePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumbackendDefaultCacheModeProp) IsValid() bool {
	for _, existing := range AllowedEnumbackendDefaultCacheModePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumbackend-defaultCacheModeProp value
func (v EnumbackendDefaultCacheModeProp) Ptr() *EnumbackendDefaultCacheModeProp {
	return &v
}

type NullableEnumbackendDefaultCacheModeProp struct {
	value *EnumbackendDefaultCacheModeProp
	isSet bool
}

func (v NullableEnumbackendDefaultCacheModeProp) Get() *EnumbackendDefaultCacheModeProp {
	return v.value
}

func (v *NullableEnumbackendDefaultCacheModeProp) Set(val *EnumbackendDefaultCacheModeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumbackendDefaultCacheModeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumbackendDefaultCacheModeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumbackendDefaultCacheModeProp(val *EnumbackendDefaultCacheModeProp) *NullableEnumbackendDefaultCacheModeProp {
	return &NullableEnumbackendDefaultCacheModeProp{value: val, isSet: true}
}

func (v NullableEnumbackendDefaultCacheModeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumbackendDefaultCacheModeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

