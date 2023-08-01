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

// EnumbackendMirroredSubtreePreferredMasterTypeProp Specifies whether the master selection algorithm should favor particular types of server.
type EnumbackendMirroredSubtreePreferredMasterTypeProp string

// List of Enumbackend-mirroredSubtreePreferredMasterTypeProp
const (
	ENUMBACKENDMIRROREDSUBTREEPREFERREDMASTERTYPEPROP_DS      EnumbackendMirroredSubtreePreferredMasterTypeProp = "ds"
	ENUMBACKENDMIRROREDSUBTREEPREFERREDMASTERTYPEPROP_PROXY   EnumbackendMirroredSubtreePreferredMasterTypeProp = "proxy"
	ENUMBACKENDMIRROREDSUBTREEPREFERREDMASTERTYPEPROP_BROKER  EnumbackendMirroredSubtreePreferredMasterTypeProp = "broker"
	ENUMBACKENDMIRROREDSUBTREEPREFERREDMASTERTYPEPROP_METRICS EnumbackendMirroredSubtreePreferredMasterTypeProp = "metrics"
	ENUMBACKENDMIRROREDSUBTREEPREFERREDMASTERTYPEPROP_SYNC    EnumbackendMirroredSubtreePreferredMasterTypeProp = "sync"
)

// All allowed values of EnumbackendMirroredSubtreePreferredMasterTypeProp enum
var AllowedEnumbackendMirroredSubtreePreferredMasterTypePropEnumValues = []EnumbackendMirroredSubtreePreferredMasterTypeProp{
	"ds",
	"proxy",
	"broker",
	"metrics",
	"sync",
}

func (v *EnumbackendMirroredSubtreePreferredMasterTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumbackendMirroredSubtreePreferredMasterTypeProp(value)
	for _, existing := range AllowedEnumbackendMirroredSubtreePreferredMasterTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumbackendMirroredSubtreePreferredMasterTypeProp", value)
}

// NewEnumbackendMirroredSubtreePreferredMasterTypePropFromValue returns a pointer to a valid EnumbackendMirroredSubtreePreferredMasterTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumbackendMirroredSubtreePreferredMasterTypePropFromValue(v string) (*EnumbackendMirroredSubtreePreferredMasterTypeProp, error) {
	ev := EnumbackendMirroredSubtreePreferredMasterTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumbackendMirroredSubtreePreferredMasterTypeProp: valid values are %v", v, AllowedEnumbackendMirroredSubtreePreferredMasterTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumbackendMirroredSubtreePreferredMasterTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumbackendMirroredSubtreePreferredMasterTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumbackend-mirroredSubtreePreferredMasterTypeProp value
func (v EnumbackendMirroredSubtreePreferredMasterTypeProp) Ptr() *EnumbackendMirroredSubtreePreferredMasterTypeProp {
	return &v
}

type NullableEnumbackendMirroredSubtreePreferredMasterTypeProp struct {
	value *EnumbackendMirroredSubtreePreferredMasterTypeProp
	isSet bool
}

func (v NullableEnumbackendMirroredSubtreePreferredMasterTypeProp) Get() *EnumbackendMirroredSubtreePreferredMasterTypeProp {
	return v.value
}

func (v *NullableEnumbackendMirroredSubtreePreferredMasterTypeProp) Set(val *EnumbackendMirroredSubtreePreferredMasterTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumbackendMirroredSubtreePreferredMasterTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumbackendMirroredSubtreePreferredMasterTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumbackendMirroredSubtreePreferredMasterTypeProp(val *EnumbackendMirroredSubtreePreferredMasterTypeProp) *NullableEnumbackendMirroredSubtreePreferredMasterTypeProp {
	return &NullableEnumbackendMirroredSubtreePreferredMasterTypeProp{value: val, isSet: true}
}

func (v NullableEnumbackendMirroredSubtreePreferredMasterTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumbackendMirroredSubtreePreferredMasterTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
