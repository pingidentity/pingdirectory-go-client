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

// EnumldapSdkDebugLoggerDebugTypeProp The types of debug messages that should be logged.
type EnumldapSdkDebugLoggerDebugTypeProp string

// List of Enumldap-sdk-debug-logger-debugTypeProp
const (
	ENUMLDAPSDKDEBUGLOGGERDEBUGTYPEPROP_ASN1         EnumldapSdkDebugLoggerDebugTypeProp = "asn1"
	ENUMLDAPSDKDEBUGLOGGERDEBUGTYPEPROP_CODING_ERROR EnumldapSdkDebugLoggerDebugTypeProp = "coding-error"
	ENUMLDAPSDKDEBUGLOGGERDEBUGTYPEPROP_CONNECT      EnumldapSdkDebugLoggerDebugTypeProp = "connect"
	ENUMLDAPSDKDEBUGLOGGERDEBUGTYPEPROP_EXCEPTION    EnumldapSdkDebugLoggerDebugTypeProp = "exception"
	ENUMLDAPSDKDEBUGLOGGERDEBUGTYPEPROP_LDAP         EnumldapSdkDebugLoggerDebugTypeProp = "ldap"
	ENUMLDAPSDKDEBUGLOGGERDEBUGTYPEPROP_LDIF         EnumldapSdkDebugLoggerDebugTypeProp = "ldif"
	ENUMLDAPSDKDEBUGLOGGERDEBUGTYPEPROP_MONITOR      EnumldapSdkDebugLoggerDebugTypeProp = "monitor"
	ENUMLDAPSDKDEBUGLOGGERDEBUGTYPEPROP_OTHER        EnumldapSdkDebugLoggerDebugTypeProp = "other"
)

// All allowed values of EnumldapSdkDebugLoggerDebugTypeProp enum
var AllowedEnumldapSdkDebugLoggerDebugTypePropEnumValues = []EnumldapSdkDebugLoggerDebugTypeProp{
	"asn1",
	"coding-error",
	"connect",
	"exception",
	"ldap",
	"ldif",
	"monitor",
	"other",
}

func (v *EnumldapSdkDebugLoggerDebugTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumldapSdkDebugLoggerDebugTypeProp(value)
	for _, existing := range AllowedEnumldapSdkDebugLoggerDebugTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumldapSdkDebugLoggerDebugTypeProp", value)
}

// NewEnumldapSdkDebugLoggerDebugTypePropFromValue returns a pointer to a valid EnumldapSdkDebugLoggerDebugTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumldapSdkDebugLoggerDebugTypePropFromValue(v string) (*EnumldapSdkDebugLoggerDebugTypeProp, error) {
	ev := EnumldapSdkDebugLoggerDebugTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumldapSdkDebugLoggerDebugTypeProp: valid values are %v", v, AllowedEnumldapSdkDebugLoggerDebugTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumldapSdkDebugLoggerDebugTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumldapSdkDebugLoggerDebugTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumldap-sdk-debug-logger-debugTypeProp value
func (v EnumldapSdkDebugLoggerDebugTypeProp) Ptr() *EnumldapSdkDebugLoggerDebugTypeProp {
	return &v
}

type NullableEnumldapSdkDebugLoggerDebugTypeProp struct {
	value *EnumldapSdkDebugLoggerDebugTypeProp
	isSet bool
}

func (v NullableEnumldapSdkDebugLoggerDebugTypeProp) Get() *EnumldapSdkDebugLoggerDebugTypeProp {
	return v.value
}

func (v *NullableEnumldapSdkDebugLoggerDebugTypeProp) Set(val *EnumldapSdkDebugLoggerDebugTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumldapSdkDebugLoggerDebugTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumldapSdkDebugLoggerDebugTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumldapSdkDebugLoggerDebugTypeProp(val *EnumldapSdkDebugLoggerDebugTypeProp) *NullableEnumldapSdkDebugLoggerDebugTypeProp {
	return &NullableEnumldapSdkDebugLoggerDebugTypeProp{value: val, isSet: true}
}

func (v NullableEnumldapSdkDebugLoggerDebugTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumldapSdkDebugLoggerDebugTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
