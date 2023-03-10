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

// EnumexternalServerAmazonAwsAuthenticationMethodProp The mechanism to use to authenticate to AWS.
type EnumexternalServerAmazonAwsAuthenticationMethodProp string

// List of Enumexternal-server-amazon-aws-authenticationMethodProp
const (
	ENUMEXTERNALSERVERAMAZONAWSAUTHENTICATIONMETHODPROP_DEFAULT_PROVIDER_CHAIN EnumexternalServerAmazonAwsAuthenticationMethodProp = "default-provider-chain"
	ENUMEXTERNALSERVERAMAZONAWSAUTHENTICATIONMETHODPROP_ACCESS_KEY             EnumexternalServerAmazonAwsAuthenticationMethodProp = "access-key"
	ENUMEXTERNALSERVERAMAZONAWSAUTHENTICATIONMETHODPROP_IAM_ROLE               EnumexternalServerAmazonAwsAuthenticationMethodProp = "iam-role"
	ENUMEXTERNALSERVERAMAZONAWSAUTHENTICATIONMETHODPROP_IRSA_ROLE              EnumexternalServerAmazonAwsAuthenticationMethodProp = "irsa-role"
)

// All allowed values of EnumexternalServerAmazonAwsAuthenticationMethodProp enum
var AllowedEnumexternalServerAmazonAwsAuthenticationMethodPropEnumValues = []EnumexternalServerAmazonAwsAuthenticationMethodProp{
	"default-provider-chain",
	"access-key",
	"iam-role",
	"irsa-role",
}

func (v *EnumexternalServerAmazonAwsAuthenticationMethodProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumexternalServerAmazonAwsAuthenticationMethodProp(value)
	for _, existing := range AllowedEnumexternalServerAmazonAwsAuthenticationMethodPropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumexternalServerAmazonAwsAuthenticationMethodProp", value)
}

// NewEnumexternalServerAmazonAwsAuthenticationMethodPropFromValue returns a pointer to a valid EnumexternalServerAmazonAwsAuthenticationMethodProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumexternalServerAmazonAwsAuthenticationMethodPropFromValue(v string) (*EnumexternalServerAmazonAwsAuthenticationMethodProp, error) {
	ev := EnumexternalServerAmazonAwsAuthenticationMethodProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumexternalServerAmazonAwsAuthenticationMethodProp: valid values are %v", v, AllowedEnumexternalServerAmazonAwsAuthenticationMethodPropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumexternalServerAmazonAwsAuthenticationMethodProp) IsValid() bool {
	for _, existing := range AllowedEnumexternalServerAmazonAwsAuthenticationMethodPropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumexternal-server-amazon-aws-authenticationMethodProp value
func (v EnumexternalServerAmazonAwsAuthenticationMethodProp) Ptr() *EnumexternalServerAmazonAwsAuthenticationMethodProp {
	return &v
}

type NullableEnumexternalServerAmazonAwsAuthenticationMethodProp struct {
	value *EnumexternalServerAmazonAwsAuthenticationMethodProp
	isSet bool
}

func (v NullableEnumexternalServerAmazonAwsAuthenticationMethodProp) Get() *EnumexternalServerAmazonAwsAuthenticationMethodProp {
	return v.value
}

func (v *NullableEnumexternalServerAmazonAwsAuthenticationMethodProp) Set(val *EnumexternalServerAmazonAwsAuthenticationMethodProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumexternalServerAmazonAwsAuthenticationMethodProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumexternalServerAmazonAwsAuthenticationMethodProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumexternalServerAmazonAwsAuthenticationMethodProp(val *EnumexternalServerAmazonAwsAuthenticationMethodProp) *NullableEnumexternalServerAmazonAwsAuthenticationMethodProp {
	return &NullableEnumexternalServerAmazonAwsAuthenticationMethodProp{value: val, isSet: true}
}

func (v NullableEnumexternalServerAmazonAwsAuthenticationMethodProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumexternalServerAmazonAwsAuthenticationMethodProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
