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

// EnumbackendId2subtreeCacheModeProp Specifies the cache mode that should be used when accessing the records in the id2subtree database, which provides a mapping between the entry ID of a particular entry and the entry IDs of all of its children to any depth. This index may be used when performing searches with a whole-subtree or subordinate-subtree scope if the search filter cannot be resolved to a small enough candidate list. The size of this database directly depends on the number of entries that have children.
type EnumbackendId2subtreeCacheModeProp string

// List of Enumbackend-id2subtreeCacheModeProp
const (
	ENUMBACKENDID2SUBTREECACHEMODEPROP_CACHE_KEYS_AND_VALUES EnumbackendId2subtreeCacheModeProp = "cache-keys-and-values"
	ENUMBACKENDID2SUBTREECACHEMODEPROP_CACHE_KEYS_ONLY EnumbackendId2subtreeCacheModeProp = "cache-keys-only"
	ENUMBACKENDID2SUBTREECACHEMODEPROP_NO_CACHING EnumbackendId2subtreeCacheModeProp = "no-caching"
	ENUMBACKENDID2SUBTREECACHEMODEPROP_KEEP_HOT EnumbackendId2subtreeCacheModeProp = "keep-hot"
	ENUMBACKENDID2SUBTREECACHEMODEPROP_DEFAULT EnumbackendId2subtreeCacheModeProp = "default"
	ENUMBACKENDID2SUBTREECACHEMODEPROP_MAKE_COLD EnumbackendId2subtreeCacheModeProp = "make-cold"
	ENUMBACKENDID2SUBTREECACHEMODEPROP_EVICT_LEAF_IMMEDIATELY EnumbackendId2subtreeCacheModeProp = "evict-leaf-immediately"
	ENUMBACKENDID2SUBTREECACHEMODEPROP_EVICT_BIN_IMMEDIATELY EnumbackendId2subtreeCacheModeProp = "evict-bin-immediately"
)

// All allowed values of EnumbackendId2subtreeCacheModeProp enum
var AllowedEnumbackendId2subtreeCacheModePropEnumValues = []EnumbackendId2subtreeCacheModeProp{
	"cache-keys-and-values",
	"cache-keys-only",
	"no-caching",
	"keep-hot",
	"default",
	"make-cold",
	"evict-leaf-immediately",
	"evict-bin-immediately",
}

func (v *EnumbackendId2subtreeCacheModeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumbackendId2subtreeCacheModeProp(value)
	for _, existing := range AllowedEnumbackendId2subtreeCacheModePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumbackendId2subtreeCacheModeProp", value)
}

// NewEnumbackendId2subtreeCacheModePropFromValue returns a pointer to a valid EnumbackendId2subtreeCacheModeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumbackendId2subtreeCacheModePropFromValue(v string) (*EnumbackendId2subtreeCacheModeProp, error) {
	ev := EnumbackendId2subtreeCacheModeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumbackendId2subtreeCacheModeProp: valid values are %v", v, AllowedEnumbackendId2subtreeCacheModePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumbackendId2subtreeCacheModeProp) IsValid() bool {
	for _, existing := range AllowedEnumbackendId2subtreeCacheModePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumbackend-id2subtreeCacheModeProp value
func (v EnumbackendId2subtreeCacheModeProp) Ptr() *EnumbackendId2subtreeCacheModeProp {
	return &v
}

type NullableEnumbackendId2subtreeCacheModeProp struct {
	value *EnumbackendId2subtreeCacheModeProp
	isSet bool
}

func (v NullableEnumbackendId2subtreeCacheModeProp) Get() *EnumbackendId2subtreeCacheModeProp {
	return v.value
}

func (v *NullableEnumbackendId2subtreeCacheModeProp) Set(val *EnumbackendId2subtreeCacheModeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumbackendId2subtreeCacheModeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumbackendId2subtreeCacheModeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumbackendId2subtreeCacheModeProp(val *EnumbackendId2subtreeCacheModeProp) *NullableEnumbackendId2subtreeCacheModeProp {
	return &NullableEnumbackendId2subtreeCacheModeProp{value: val, isSet: true}
}

func (v NullableEnumbackendId2subtreeCacheModeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumbackendId2subtreeCacheModeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
