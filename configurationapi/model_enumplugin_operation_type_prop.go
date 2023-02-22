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

// EnumpluginOperationTypeProp Specifies the types of operations that should result in access time updates.
type EnumpluginOperationTypeProp string

// List of Enumplugin-operationTypeProp
const (
	ENUMPLUGINOPERATIONTYPEPROP_ADD       EnumpluginOperationTypeProp = "add"
	ENUMPLUGINOPERATIONTYPEPROP_BIND      EnumpluginOperationTypeProp = "bind"
	ENUMPLUGINOPERATIONTYPEPROP_COMPARE   EnumpluginOperationTypeProp = "compare"
	ENUMPLUGINOPERATIONTYPEPROP_MODIFY    EnumpluginOperationTypeProp = "modify"
	ENUMPLUGINOPERATIONTYPEPROP_MODIFY_DN EnumpluginOperationTypeProp = "modify-dn"
	ENUMPLUGINOPERATIONTYPEPROP_SEARCH    EnumpluginOperationTypeProp = "search"
)

// All allowed values of EnumpluginOperationTypeProp enum
var AllowedEnumpluginOperationTypePropEnumValues = []EnumpluginOperationTypeProp{
	"add",
	"bind",
	"compare",
	"modify",
	"modify-dn",
	"search",
}

func (v *EnumpluginOperationTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumpluginOperationTypeProp(value)
	for _, existing := range AllowedEnumpluginOperationTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumpluginOperationTypeProp", value)
}

// NewEnumpluginOperationTypePropFromValue returns a pointer to a valid EnumpluginOperationTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumpluginOperationTypePropFromValue(v string) (*EnumpluginOperationTypeProp, error) {
	ev := EnumpluginOperationTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumpluginOperationTypeProp: valid values are %v", v, AllowedEnumpluginOperationTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumpluginOperationTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumpluginOperationTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumplugin-operationTypeProp value
func (v EnumpluginOperationTypeProp) Ptr() *EnumpluginOperationTypeProp {
	return &v
}

type NullableEnumpluginOperationTypeProp struct {
	value *EnumpluginOperationTypeProp
	isSet bool
}

func (v NullableEnumpluginOperationTypeProp) Get() *EnumpluginOperationTypeProp {
	return v.value
}

func (v *NullableEnumpluginOperationTypeProp) Set(val *EnumpluginOperationTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumpluginOperationTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumpluginOperationTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumpluginOperationTypeProp(val *EnumpluginOperationTypeProp) *NullableEnumpluginOperationTypeProp {
	return &NullableEnumpluginOperationTypeProp{value: val, isSet: true}
}

func (v NullableEnumpluginOperationTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumpluginOperationTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}