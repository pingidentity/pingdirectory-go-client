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

// EnumsensitiveAttributeAllowInReturnedEntriesProp Indicates whether sensitive attributes should be included in entries returned to the client. This includes not only search result entries, but also other forms including in the values of controls like the pre-read, post-read, get authorization entry, and LDAP join response controls.
type EnumsensitiveAttributeAllowInReturnedEntriesProp string

// List of Enumsensitive-attribute-allowInReturnedEntriesProp
const (
	ENUMSENSITIVEATTRIBUTEALLOWINRETURNEDENTRIESPROP_ALLOW       EnumsensitiveAttributeAllowInReturnedEntriesProp = "allow"
	ENUMSENSITIVEATTRIBUTEALLOWINRETURNEDENTRIESPROP_SUPPRESS    EnumsensitiveAttributeAllowInReturnedEntriesProp = "suppress"
	ENUMSENSITIVEATTRIBUTEALLOWINRETURNEDENTRIESPROP_SECURE_ONLY EnumsensitiveAttributeAllowInReturnedEntriesProp = "secure-only"
)

// All allowed values of EnumsensitiveAttributeAllowInReturnedEntriesProp enum
var AllowedEnumsensitiveAttributeAllowInReturnedEntriesPropEnumValues = []EnumsensitiveAttributeAllowInReturnedEntriesProp{
	"allow",
	"suppress",
	"secure-only",
}

func (v *EnumsensitiveAttributeAllowInReturnedEntriesProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsensitiveAttributeAllowInReturnedEntriesProp(value)
	for _, existing := range AllowedEnumsensitiveAttributeAllowInReturnedEntriesPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsensitiveAttributeAllowInReturnedEntriesProp", value)
}

// NewEnumsensitiveAttributeAllowInReturnedEntriesPropFromValue returns a pointer to a valid EnumsensitiveAttributeAllowInReturnedEntriesProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsensitiveAttributeAllowInReturnedEntriesPropFromValue(v string) (*EnumsensitiveAttributeAllowInReturnedEntriesProp, error) {
	ev := EnumsensitiveAttributeAllowInReturnedEntriesProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsensitiveAttributeAllowInReturnedEntriesProp: valid values are %v", v, AllowedEnumsensitiveAttributeAllowInReturnedEntriesPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsensitiveAttributeAllowInReturnedEntriesProp) IsValid() bool {
	for _, existing := range AllowedEnumsensitiveAttributeAllowInReturnedEntriesPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumsensitive-attribute-allowInReturnedEntriesProp value
func (v EnumsensitiveAttributeAllowInReturnedEntriesProp) Ptr() *EnumsensitiveAttributeAllowInReturnedEntriesProp {
	return &v
}

type NullableEnumsensitiveAttributeAllowInReturnedEntriesProp struct {
	value *EnumsensitiveAttributeAllowInReturnedEntriesProp
	isSet bool
}

func (v NullableEnumsensitiveAttributeAllowInReturnedEntriesProp) Get() *EnumsensitiveAttributeAllowInReturnedEntriesProp {
	return v.value
}

func (v *NullableEnumsensitiveAttributeAllowInReturnedEntriesProp) Set(val *EnumsensitiveAttributeAllowInReturnedEntriesProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsensitiveAttributeAllowInReturnedEntriesProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsensitiveAttributeAllowInReturnedEntriesProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsensitiveAttributeAllowInReturnedEntriesProp(val *EnumsensitiveAttributeAllowInReturnedEntriesProp) *NullableEnumsensitiveAttributeAllowInReturnedEntriesProp {
	return &NullableEnumsensitiveAttributeAllowInReturnedEntriesProp{value: val, isSet: true}
}

func (v NullableEnumsensitiveAttributeAllowInReturnedEntriesProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsensitiveAttributeAllowInReturnedEntriesProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
