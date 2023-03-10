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

// EnumldapPassThroughAuthenticationHandlerSchemaUrn the model 'EnumldapPassThroughAuthenticationHandlerSchemaUrn'
type EnumldapPassThroughAuthenticationHandlerSchemaUrn string

// List of Enumldap-pass-through-authentication-handlerSchemaUrn
const (
	ENUMLDAPPASSTHROUGHAUTHENTICATIONHANDLERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0PASS_THROUGH_AUTHENTICATION_HANDLERLDAP EnumldapPassThroughAuthenticationHandlerSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:pass-through-authentication-handler:ldap"
)

// All allowed values of EnumldapPassThroughAuthenticationHandlerSchemaUrn enum
var AllowedEnumldapPassThroughAuthenticationHandlerSchemaUrnEnumValues = []EnumldapPassThroughAuthenticationHandlerSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:pass-through-authentication-handler:ldap",
}

func (v *EnumldapPassThroughAuthenticationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumldapPassThroughAuthenticationHandlerSchemaUrn(value)
	for _, existing := range AllowedEnumldapPassThroughAuthenticationHandlerSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumldapPassThroughAuthenticationHandlerSchemaUrn", value)
}

// NewEnumldapPassThroughAuthenticationHandlerSchemaUrnFromValue returns a pointer to a valid EnumldapPassThroughAuthenticationHandlerSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumldapPassThroughAuthenticationHandlerSchemaUrnFromValue(v string) (*EnumldapPassThroughAuthenticationHandlerSchemaUrn, error) {
	ev := EnumldapPassThroughAuthenticationHandlerSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumldapPassThroughAuthenticationHandlerSchemaUrn: valid values are %v", v, AllowedEnumldapPassThroughAuthenticationHandlerSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumldapPassThroughAuthenticationHandlerSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumldapPassThroughAuthenticationHandlerSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumldap-pass-through-authentication-handlerSchemaUrn value
func (v EnumldapPassThroughAuthenticationHandlerSchemaUrn) Ptr() *EnumldapPassThroughAuthenticationHandlerSchemaUrn {
	return &v
}

type NullableEnumldapPassThroughAuthenticationHandlerSchemaUrn struct {
	value *EnumldapPassThroughAuthenticationHandlerSchemaUrn
	isSet bool
}

func (v NullableEnumldapPassThroughAuthenticationHandlerSchemaUrn) Get() *EnumldapPassThroughAuthenticationHandlerSchemaUrn {
	return v.value
}

func (v *NullableEnumldapPassThroughAuthenticationHandlerSchemaUrn) Set(val *EnumldapPassThroughAuthenticationHandlerSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumldapPassThroughAuthenticationHandlerSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumldapPassThroughAuthenticationHandlerSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumldapPassThroughAuthenticationHandlerSchemaUrn(val *EnumldapPassThroughAuthenticationHandlerSchemaUrn) *NullableEnumldapPassThroughAuthenticationHandlerSchemaUrn {
	return &NullableEnumldapPassThroughAuthenticationHandlerSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumldapPassThroughAuthenticationHandlerSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumldapPassThroughAuthenticationHandlerSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
