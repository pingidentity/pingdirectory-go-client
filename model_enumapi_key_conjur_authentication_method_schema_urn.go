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

// EnumapiKeyConjurAuthenticationMethodSchemaUrn the model 'EnumapiKeyConjurAuthenticationMethodSchemaUrn'
type EnumapiKeyConjurAuthenticationMethodSchemaUrn string

// List of Enumapi-key-conjur-authentication-methodSchemaUrn
const (
	URNPINGIDENTITYSCHEMASCONFIGURATION2_0CONJUR_AUTHENTICATION_METHODAPI_KEY EnumapiKeyConjurAuthenticationMethodSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:conjur-authentication-method:api-key"
)

// All allowed values of EnumapiKeyConjurAuthenticationMethodSchemaUrn enum
var AllowedEnumapiKeyConjurAuthenticationMethodSchemaUrnEnumValues = []EnumapiKeyConjurAuthenticationMethodSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:conjur-authentication-method:api-key",
}

func (v *EnumapiKeyConjurAuthenticationMethodSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumapiKeyConjurAuthenticationMethodSchemaUrn(value)
	for _, existing := range AllowedEnumapiKeyConjurAuthenticationMethodSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumapiKeyConjurAuthenticationMethodSchemaUrn", value)
}

// NewEnumapiKeyConjurAuthenticationMethodSchemaUrnFromValue returns a pointer to a valid EnumapiKeyConjurAuthenticationMethodSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumapiKeyConjurAuthenticationMethodSchemaUrnFromValue(v string) (*EnumapiKeyConjurAuthenticationMethodSchemaUrn, error) {
	ev := EnumapiKeyConjurAuthenticationMethodSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumapiKeyConjurAuthenticationMethodSchemaUrn: valid values are %v", v, AllowedEnumapiKeyConjurAuthenticationMethodSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumapiKeyConjurAuthenticationMethodSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumapiKeyConjurAuthenticationMethodSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumapi-key-conjur-authentication-methodSchemaUrn value
func (v EnumapiKeyConjurAuthenticationMethodSchemaUrn) Ptr() *EnumapiKeyConjurAuthenticationMethodSchemaUrn {
	return &v
}

type NullableEnumapiKeyConjurAuthenticationMethodSchemaUrn struct {
	value *EnumapiKeyConjurAuthenticationMethodSchemaUrn
	isSet bool
}

func (v NullableEnumapiKeyConjurAuthenticationMethodSchemaUrn) Get() *EnumapiKeyConjurAuthenticationMethodSchemaUrn {
	return v.value
}

func (v *NullableEnumapiKeyConjurAuthenticationMethodSchemaUrn) Set(val *EnumapiKeyConjurAuthenticationMethodSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumapiKeyConjurAuthenticationMethodSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumapiKeyConjurAuthenticationMethodSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumapiKeyConjurAuthenticationMethodSchemaUrn(val *EnumapiKeyConjurAuthenticationMethodSchemaUrn) *NullableEnumapiKeyConjurAuthenticationMethodSchemaUrn {
	return &NullableEnumapiKeyConjurAuthenticationMethodSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumapiKeyConjurAuthenticationMethodSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumapiKeyConjurAuthenticationMethodSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

