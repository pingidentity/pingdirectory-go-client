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

// EnumstringTokenClaimValidationSchemaUrn the model 'EnumstringTokenClaimValidationSchemaUrn'
type EnumstringTokenClaimValidationSchemaUrn string

// List of Enumstring-token-claim-validationSchemaUrn
const (
	ENUMSTRINGTOKENCLAIMVALIDATIONSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0TOKEN_CLAIM_VALIDATIONSTRING EnumstringTokenClaimValidationSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:token-claim-validation:string"
)

// All allowed values of EnumstringTokenClaimValidationSchemaUrn enum
var AllowedEnumstringTokenClaimValidationSchemaUrnEnumValues = []EnumstringTokenClaimValidationSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:token-claim-validation:string",
}

func (v *EnumstringTokenClaimValidationSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumstringTokenClaimValidationSchemaUrn(value)
	for _, existing := range AllowedEnumstringTokenClaimValidationSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumstringTokenClaimValidationSchemaUrn", value)
}

// NewEnumstringTokenClaimValidationSchemaUrnFromValue returns a pointer to a valid EnumstringTokenClaimValidationSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumstringTokenClaimValidationSchemaUrnFromValue(v string) (*EnumstringTokenClaimValidationSchemaUrn, error) {
	ev := EnumstringTokenClaimValidationSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumstringTokenClaimValidationSchemaUrn: valid values are %v", v, AllowedEnumstringTokenClaimValidationSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumstringTokenClaimValidationSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumstringTokenClaimValidationSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumstring-token-claim-validationSchemaUrn value
func (v EnumstringTokenClaimValidationSchemaUrn) Ptr() *EnumstringTokenClaimValidationSchemaUrn {
	return &v
}

type NullableEnumstringTokenClaimValidationSchemaUrn struct {
	value *EnumstringTokenClaimValidationSchemaUrn
	isSet bool
}

func (v NullableEnumstringTokenClaimValidationSchemaUrn) Get() *EnumstringTokenClaimValidationSchemaUrn {
	return v.value
}

func (v *NullableEnumstringTokenClaimValidationSchemaUrn) Set(val *EnumstringTokenClaimValidationSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumstringTokenClaimValidationSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumstringTokenClaimValidationSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumstringTokenClaimValidationSchemaUrn(val *EnumstringTokenClaimValidationSchemaUrn) *NullableEnumstringTokenClaimValidationSchemaUrn {
	return &NullableEnumstringTokenClaimValidationSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumstringTokenClaimValidationSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumstringTokenClaimValidationSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}