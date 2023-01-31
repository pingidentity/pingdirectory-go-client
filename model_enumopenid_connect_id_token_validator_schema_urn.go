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

// EnumopenidConnectIdTokenValidatorSchemaUrn the model 'EnumopenidConnectIdTokenValidatorSchemaUrn'
type EnumopenidConnectIdTokenValidatorSchemaUrn string

// List of Enumopenid-connect-id-token-validatorSchemaUrn
const (
	ENUMOPENIDCONNECTIDTOKENVALIDATORSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0ID_TOKEN_VALIDATOROPENID_CONNECT EnumopenidConnectIdTokenValidatorSchemaUrn = "urn:pingidentity:schemas:configuration:2.0:id-token-validator:openid-connect"
)

// All allowed values of EnumopenidConnectIdTokenValidatorSchemaUrn enum
var AllowedEnumopenidConnectIdTokenValidatorSchemaUrnEnumValues = []EnumopenidConnectIdTokenValidatorSchemaUrn{
	"urn:pingidentity:schemas:configuration:2.0:id-token-validator:openid-connect",
}

func (v *EnumopenidConnectIdTokenValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumopenidConnectIdTokenValidatorSchemaUrn(value)
	for _, existing := range AllowedEnumopenidConnectIdTokenValidatorSchemaUrnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumopenidConnectIdTokenValidatorSchemaUrn", value)
}

// NewEnumopenidConnectIdTokenValidatorSchemaUrnFromValue returns a pointer to a valid EnumopenidConnectIdTokenValidatorSchemaUrn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumopenidConnectIdTokenValidatorSchemaUrnFromValue(v string) (*EnumopenidConnectIdTokenValidatorSchemaUrn, error) {
	ev := EnumopenidConnectIdTokenValidatorSchemaUrn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumopenidConnectIdTokenValidatorSchemaUrn: valid values are %v", v, AllowedEnumopenidConnectIdTokenValidatorSchemaUrnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumopenidConnectIdTokenValidatorSchemaUrn) IsValid() bool {
	for _, existing := range AllowedEnumopenidConnectIdTokenValidatorSchemaUrnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumopenid-connect-id-token-validatorSchemaUrn value
func (v EnumopenidConnectIdTokenValidatorSchemaUrn) Ptr() *EnumopenidConnectIdTokenValidatorSchemaUrn {
	return &v
}

type NullableEnumopenidConnectIdTokenValidatorSchemaUrn struct {
	value *EnumopenidConnectIdTokenValidatorSchemaUrn
	isSet bool
}

func (v NullableEnumopenidConnectIdTokenValidatorSchemaUrn) Get() *EnumopenidConnectIdTokenValidatorSchemaUrn {
	return v.value
}

func (v *NullableEnumopenidConnectIdTokenValidatorSchemaUrn) Set(val *EnumopenidConnectIdTokenValidatorSchemaUrn) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumopenidConnectIdTokenValidatorSchemaUrn) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumopenidConnectIdTokenValidatorSchemaUrn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumopenidConnectIdTokenValidatorSchemaUrn(val *EnumopenidConnectIdTokenValidatorSchemaUrn) *NullableEnumopenidConnectIdTokenValidatorSchemaUrn {
	return &NullableEnumopenidConnectIdTokenValidatorSchemaUrn{value: val, isSet: true}
}

func (v NullableEnumopenidConnectIdTokenValidatorSchemaUrn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumopenidConnectIdTokenValidatorSchemaUrn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
