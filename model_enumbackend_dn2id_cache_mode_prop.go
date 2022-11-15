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

// EnumbackendDn2idCacheModeProp Specifies the cache mode that should be used when accessing the records in the dn2id database, which provides a mapping between normalized entry DNs and the corresponding entry IDs.
type EnumbackendDn2idCacheModeProp string

// List of Enumbackend-dn2idCacheModeProp
const (
	ENUMBACKENDDN2IDCACHEMODEPROP_CACHE_KEYS_AND_VALUES EnumbackendDn2idCacheModeProp = "cache-keys-and-values"
	ENUMBACKENDDN2IDCACHEMODEPROP_CACHE_KEYS_ONLY EnumbackendDn2idCacheModeProp = "cache-keys-only"
	ENUMBACKENDDN2IDCACHEMODEPROP_NO_CACHING EnumbackendDn2idCacheModeProp = "no-caching"
	ENUMBACKENDDN2IDCACHEMODEPROP_KEEP_HOT EnumbackendDn2idCacheModeProp = "keep-hot"
	ENUMBACKENDDN2IDCACHEMODEPROP_DEFAULT EnumbackendDn2idCacheModeProp = "default"
	ENUMBACKENDDN2IDCACHEMODEPROP_MAKE_COLD EnumbackendDn2idCacheModeProp = "make-cold"
	ENUMBACKENDDN2IDCACHEMODEPROP_EVICT_LEAF_IMMEDIATELY EnumbackendDn2idCacheModeProp = "evict-leaf-immediately"
	ENUMBACKENDDN2IDCACHEMODEPROP_EVICT_BIN_IMMEDIATELY EnumbackendDn2idCacheModeProp = "evict-bin-immediately"
)

// All allowed values of EnumbackendDn2idCacheModeProp enum
var AllowedEnumbackendDn2idCacheModePropEnumValues = []EnumbackendDn2idCacheModeProp{
	"cache-keys-and-values",
	"cache-keys-only",
	"no-caching",
	"keep-hot",
	"default",
	"make-cold",
	"evict-leaf-immediately",
	"evict-bin-immediately",
}

func (v *EnumbackendDn2idCacheModeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumbackendDn2idCacheModeProp(value)
	for _, existing := range AllowedEnumbackendDn2idCacheModePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumbackendDn2idCacheModeProp", value)
}

// NewEnumbackendDn2idCacheModePropFromValue returns a pointer to a valid EnumbackendDn2idCacheModeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumbackendDn2idCacheModePropFromValue(v string) (*EnumbackendDn2idCacheModeProp, error) {
	ev := EnumbackendDn2idCacheModeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumbackendDn2idCacheModeProp: valid values are %v", v, AllowedEnumbackendDn2idCacheModePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumbackendDn2idCacheModeProp) IsValid() bool {
	for _, existing := range AllowedEnumbackendDn2idCacheModePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumbackend-dn2idCacheModeProp value
func (v EnumbackendDn2idCacheModeProp) Ptr() *EnumbackendDn2idCacheModeProp {
	return &v
}

type NullableEnumbackendDn2idCacheModeProp struct {
	value *EnumbackendDn2idCacheModeProp
	isSet bool
}

func (v NullableEnumbackendDn2idCacheModeProp) Get() *EnumbackendDn2idCacheModeProp {
	return v.value
}

func (v *NullableEnumbackendDn2idCacheModeProp) Set(val *EnumbackendDn2idCacheModeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumbackendDn2idCacheModeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumbackendDn2idCacheModeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumbackendDn2idCacheModeProp(val *EnumbackendDn2idCacheModeProp) *NullableEnumbackendDn2idCacheModeProp {
	return &NullableEnumbackendDn2idCacheModeProp{value: val, isSet: true}
}

func (v NullableEnumbackendDn2idCacheModeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumbackendDn2idCacheModeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

