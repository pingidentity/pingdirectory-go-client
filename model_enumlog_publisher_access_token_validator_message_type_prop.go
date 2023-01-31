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

// EnumlogPublisherAccessTokenValidatorMessageTypeProp Specifies the access token validator message types that can be logged.
type EnumlogPublisherAccessTokenValidatorMessageTypeProp string

// List of Enumlog-publisher-accessTokenValidatorMessageTypeProp
const (
	ENUMLOGPUBLISHERACCESSTOKENVALIDATORMESSAGETYPEPROP_SUBJECT_LOOKUP           EnumlogPublisherAccessTokenValidatorMessageTypeProp = "subject-lookup"
	ENUMLOGPUBLISHERACCESSTOKENVALIDATORMESSAGETYPEPROP_EXTERNAL_SERVER_REQUEST  EnumlogPublisherAccessTokenValidatorMessageTypeProp = "external-server-request"
	ENUMLOGPUBLISHERACCESSTOKENVALIDATORMESSAGETYPEPROP_EXTERNAL_SERVER_RESPONSE EnumlogPublisherAccessTokenValidatorMessageTypeProp = "external-server-response"
	ENUMLOGPUBLISHERACCESSTOKENVALIDATORMESSAGETYPEPROP_VALIDATION               EnumlogPublisherAccessTokenValidatorMessageTypeProp = "validation"
	ENUMLOGPUBLISHERACCESSTOKENVALIDATORMESSAGETYPEPROP_ERROR                    EnumlogPublisherAccessTokenValidatorMessageTypeProp = "error"
)

// All allowed values of EnumlogPublisherAccessTokenValidatorMessageTypeProp enum
var AllowedEnumlogPublisherAccessTokenValidatorMessageTypePropEnumValues = []EnumlogPublisherAccessTokenValidatorMessageTypeProp{
	"subject-lookup",
	"external-server-request",
	"external-server-response",
	"validation",
	"error",
}

func (v *EnumlogPublisherAccessTokenValidatorMessageTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumlogPublisherAccessTokenValidatorMessageTypeProp(value)
	for _, existing := range AllowedEnumlogPublisherAccessTokenValidatorMessageTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumlogPublisherAccessTokenValidatorMessageTypeProp", value)
}

// NewEnumlogPublisherAccessTokenValidatorMessageTypePropFromValue returns a pointer to a valid EnumlogPublisherAccessTokenValidatorMessageTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumlogPublisherAccessTokenValidatorMessageTypePropFromValue(v string) (*EnumlogPublisherAccessTokenValidatorMessageTypeProp, error) {
	ev := EnumlogPublisherAccessTokenValidatorMessageTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumlogPublisherAccessTokenValidatorMessageTypeProp: valid values are %v", v, AllowedEnumlogPublisherAccessTokenValidatorMessageTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumlogPublisherAccessTokenValidatorMessageTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumlogPublisherAccessTokenValidatorMessageTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumlog-publisher-accessTokenValidatorMessageTypeProp value
func (v EnumlogPublisherAccessTokenValidatorMessageTypeProp) Ptr() *EnumlogPublisherAccessTokenValidatorMessageTypeProp {
	return &v
}

type NullableEnumlogPublisherAccessTokenValidatorMessageTypeProp struct {
	value *EnumlogPublisherAccessTokenValidatorMessageTypeProp
	isSet bool
}

func (v NullableEnumlogPublisherAccessTokenValidatorMessageTypeProp) Get() *EnumlogPublisherAccessTokenValidatorMessageTypeProp {
	return v.value
}

func (v *NullableEnumlogPublisherAccessTokenValidatorMessageTypeProp) Set(val *EnumlogPublisherAccessTokenValidatorMessageTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumlogPublisherAccessTokenValidatorMessageTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumlogPublisherAccessTokenValidatorMessageTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumlogPublisherAccessTokenValidatorMessageTypeProp(val *EnumlogPublisherAccessTokenValidatorMessageTypeProp) *NullableEnumlogPublisherAccessTokenValidatorMessageTypeProp {
	return &NullableEnumlogPublisherAccessTokenValidatorMessageTypeProp{value: val, isSet: true}
}

func (v NullableEnumlogPublisherAccessTokenValidatorMessageTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumlogPublisherAccessTokenValidatorMessageTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
