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

// EnumexternalServerScimHostnameVerificationMethodProp The mechanism for checking if the service provider's hostname matches the name(s) stored inside the server's X.509 certificate. This is only applicable if SSL is being used for connection security.
type EnumexternalServerScimHostnameVerificationMethodProp string

// List of Enumexternal-server-scim-hostnameVerificationMethodProp
const (
	ENUMEXTERNALSERVERSCIMHOSTNAMEVERIFICATIONMETHODPROP_ALLOW_ALL          EnumexternalServerScimHostnameVerificationMethodProp = "allow-all"
	ENUMEXTERNALSERVERSCIMHOSTNAMEVERIFICATIONMETHODPROP_BROWSER_COMPATIBLE EnumexternalServerScimHostnameVerificationMethodProp = "browser-compatible"
	ENUMEXTERNALSERVERSCIMHOSTNAMEVERIFICATIONMETHODPROP_STRICT             EnumexternalServerScimHostnameVerificationMethodProp = "strict"
)

// All allowed values of EnumexternalServerScimHostnameVerificationMethodProp enum
var AllowedEnumexternalServerScimHostnameVerificationMethodPropEnumValues = []EnumexternalServerScimHostnameVerificationMethodProp{
	"allow-all",
	"browser-compatible",
	"strict",
}

func (v *EnumexternalServerScimHostnameVerificationMethodProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerScimHostnameVerificationMethodProp(value)
	for _, existing := range AllowedEnumexternalServerScimHostnameVerificationMethodPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerScimHostnameVerificationMethodProp", value)
}

// NewEnumexternalServerScimHostnameVerificationMethodPropFromValue returns a pointer to a valid EnumexternalServerScimHostnameVerificationMethodProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerScimHostnameVerificationMethodPropFromValue(v string) (*EnumexternalServerScimHostnameVerificationMethodProp, error) {
	ev := EnumexternalServerScimHostnameVerificationMethodProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerScimHostnameVerificationMethodProp: valid values are %v", v, AllowedEnumexternalServerScimHostnameVerificationMethodPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerScimHostnameVerificationMethodProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerScimHostnameVerificationMethodPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-scim-hostnameVerificationMethodProp value
func (v EnumexternalServerScimHostnameVerificationMethodProp) Ptr() *EnumexternalServerScimHostnameVerificationMethodProp {
	return &v
}

type NullableEnumexternalServerScimHostnameVerificationMethodProp struct {
	value *EnumexternalServerScimHostnameVerificationMethodProp
	isSet bool
}

func (v NullableEnumexternalServerScimHostnameVerificationMethodProp) Get() *EnumexternalServerScimHostnameVerificationMethodProp {
	return v.value
}

func (v *NullableEnumexternalServerScimHostnameVerificationMethodProp) Set(val *EnumexternalServerScimHostnameVerificationMethodProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerScimHostnameVerificationMethodProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerScimHostnameVerificationMethodProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerScimHostnameVerificationMethodProp(val *EnumexternalServerScimHostnameVerificationMethodProp) *NullableEnumexternalServerScimHostnameVerificationMethodProp {
	return &NullableEnumexternalServerScimHostnameVerificationMethodProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerScimHostnameVerificationMethodProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerScimHostnameVerificationMethodProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
