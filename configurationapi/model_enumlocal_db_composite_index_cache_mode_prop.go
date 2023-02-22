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

// EnumlocalDbCompositeIndexCacheModeProp The behavior that the server should exhibit when storing information from this index in the database cache.
type EnumlocalDbCompositeIndexCacheModeProp string

// List of Enumlocal-db-composite-index-cacheModeProp
const (
	ENUMLOCALDBCOMPOSITEINDEXCACHEMODEPROP_CACHE_KEYS_AND_VALUES EnumlocalDbCompositeIndexCacheModeProp = "cache-keys-and-values"
	ENUMLOCALDBCOMPOSITEINDEXCACHEMODEPROP_CACHE_KEYS_ONLY       EnumlocalDbCompositeIndexCacheModeProp = "cache-keys-only"
	ENUMLOCALDBCOMPOSITEINDEXCACHEMODEPROP_NO_CACHING            EnumlocalDbCompositeIndexCacheModeProp = "no-caching"
)

// All allowed values of EnumlocalDbCompositeIndexCacheModeProp enum
var AllowedEnumlocalDbCompositeIndexCacheModePropEnumValues = []EnumlocalDbCompositeIndexCacheModeProp{
	"cache-keys-and-values",
	"cache-keys-only",
	"no-caching",
}

func (v *EnumlocalDbCompositeIndexCacheModeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlocalDbCompositeIndexCacheModeProp(value)
	for _, existing := range AllowedEnumlocalDbCompositeIndexCacheModePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlocalDbCompositeIndexCacheModeProp", value)
}

// NewEnumlocalDbCompositeIndexCacheModePropFromValue returns a pointer to a valid EnumlocalDbCompositeIndexCacheModeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlocalDbCompositeIndexCacheModePropFromValue(v string) (*EnumlocalDbCompositeIndexCacheModeProp, error) {
	ev := EnumlocalDbCompositeIndexCacheModeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlocalDbCompositeIndexCacheModeProp: valid values are %v", v, AllowedEnumlocalDbCompositeIndexCacheModePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlocalDbCompositeIndexCacheModeProp) IsValid() bool {
	for _, existing := range AllowedEnumlocalDbCompositeIndexCacheModePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlocal-db-composite-index-cacheModeProp value
func (v EnumlocalDbCompositeIndexCacheModeProp) Ptr() *EnumlocalDbCompositeIndexCacheModeProp {
	return &v
}

type NullableEnumlocalDbCompositeIndexCacheModeProp struct {
	value *EnumlocalDbCompositeIndexCacheModeProp
	isSet bool
}

func (v NullableEnumlocalDbCompositeIndexCacheModeProp) Get() *EnumlocalDbCompositeIndexCacheModeProp {
	return v.value
}

func (v *NullableEnumlocalDbCompositeIndexCacheModeProp) Set(val *EnumlocalDbCompositeIndexCacheModeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlocalDbCompositeIndexCacheModeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlocalDbCompositeIndexCacheModeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlocalDbCompositeIndexCacheModeProp(val *EnumlocalDbCompositeIndexCacheModeProp) *NullableEnumlocalDbCompositeIndexCacheModeProp {
	return &NullableEnumlocalDbCompositeIndexCacheModeProp{value: val, isSet: true}
}

func (v NullableEnumlocalDbCompositeIndexCacheModeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlocalDbCompositeIndexCacheModeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}