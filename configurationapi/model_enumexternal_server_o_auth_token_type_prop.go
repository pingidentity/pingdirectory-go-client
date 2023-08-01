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

// EnumexternalServerOAuthTokenTypeProp The type of OAuth token to use in conjunction with the OAuth authentication-method.
type EnumexternalServerOAuthTokenTypeProp string

// List of Enumexternal-server-OAuthTokenTypeProp
const (
	ENUMEXTERNALSERVEROAUTHTOKENTYPEPROP_BEARER EnumexternalServerOAuthTokenTypeProp = "bearer"
	ENUMEXTERNALSERVEROAUTHTOKENTYPEPROP_OAUTH  EnumexternalServerOAuthTokenTypeProp = "oauth"
)

// All allowed values of EnumexternalServerOAuthTokenTypeProp enum
var AllowedEnumexternalServerOAuthTokenTypePropEnumValues = []EnumexternalServerOAuthTokenTypeProp{
	"bearer",
	"oauth",
}

func (v *EnumexternalServerOAuthTokenTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerOAuthTokenTypeProp(value)
	for _, existing := range AllowedEnumexternalServerOAuthTokenTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerOAuthTokenTypeProp", value)
}

// NewEnumexternalServerOAuthTokenTypePropFromValue returns a pointer to a valid EnumexternalServerOAuthTokenTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerOAuthTokenTypePropFromValue(v string) (*EnumexternalServerOAuthTokenTypeProp, error) {
	ev := EnumexternalServerOAuthTokenTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerOAuthTokenTypeProp: valid values are %v", v, AllowedEnumexternalServerOAuthTokenTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerOAuthTokenTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerOAuthTokenTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-OAuthTokenTypeProp value
func (v EnumexternalServerOAuthTokenTypeProp) Ptr() *EnumexternalServerOAuthTokenTypeProp {
	return &v
}

type NullableEnumexternalServerOAuthTokenTypeProp struct {
	value *EnumexternalServerOAuthTokenTypeProp
	isSet bool
}

func (v NullableEnumexternalServerOAuthTokenTypeProp) Get() *EnumexternalServerOAuthTokenTypeProp {
	return v.value
}

func (v *NullableEnumexternalServerOAuthTokenTypeProp) Set(val *EnumexternalServerOAuthTokenTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerOAuthTokenTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerOAuthTokenTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerOAuthTokenTypeProp(val *EnumexternalServerOAuthTokenTypeProp) *NullableEnumexternalServerOAuthTokenTypeProp {
	return &NullableEnumexternalServerOAuthTokenTypeProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerOAuthTokenTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerOAuthTokenTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
